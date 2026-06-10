package proxy

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"agent-proxy-v2/internal/backend"
	"agent-proxy-v2/internal/protocol"
	"agent-proxy-v2/internal/strategy"
)

// StrategyFn returns the strategy for a given model (hot-reloadable).
type StrategyFn func(model string) strategy.LoadBalanceStrategy

// Handler is the main proxy request handler — format detection, routing, translation.
type Handler struct {
	Registry        *backend.Registry
	GetStrategy     StrategyFn // hot-reloadable strategy lookup
	DefaultStrategy strategy.LoadBalanceStrategy
	HTTPClient      *http.Client // shared client with connection pooling
}

// HandleRequest processes an incoming request and proxies it to a backend.
func (h *Handler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "read body: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Detect protocol format from URL path or headers
	clientFormat := h.detectFormat(r)

	// Parse request into canonical IR
	var ir *protocol.RequestIR
	switch clientFormat {
	case "anthropic":
		ir, err = protocol.ParseAnthropicRequest(body)
	case "openai":
		ir, err = protocol.ParseOpenAIRequest(body)
	default:
		http.Error(w, "unknown format", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, fmt.Sprintf("parse request: %v", err), http.StatusBadRequest)
		return
	}

	// Resolve model -> backends
	candidates := h.Registry.GetForModel(ir.Model)

	// Also try fuzzy matching: if model contains "gpt-", search for "gpt-" backends
	if len(candidates) == 0 && strings.Contains(ir.Model, "gpt") {
		candidates = h.Registry.GetForModel("gpt-*")
	}
	if len(candidates) == 0 && strings.Contains(ir.Model, "claude") {
		candidates = h.Registry.GetForModel("claude-*")
	}

	if len(candidates) == 0 {
		log.Printf("[proxy] no backends for model=%s, falling back to auto", ir.Model)
		ir.Model = "auto"
		candidates = h.Registry.GetForModel(ir.Model)
	}
	if len(candidates) == 0 {
		writeError(w, clientFormat, "no backends available for model "+ir.Model)
		return
	}

	// Context-length filter pipeline
	estTokens := estimateInputTokens(ir)
	candidates = strategy.FilterByContextWindow(candidates, estTokens)
	if len(candidates) == 0 {
		writeError(w, clientFormat, "all backends unhealthy or exhausted for model "+ir.Model)
		return
	}

	// Split candidates by cost tier: prepaid first, then pay-per-token
	prepaid, payPerToken := splitByCostTier(candidates)
	groups := make([][]*backend.Backend, 0, 2)
	groupNames := make([]string, 0, 2)
	if len(prepaid) > 0 {
		groups = append(groups, prepaid)
		groupNames = append(groupNames, "prepaid")
	}
	if len(payPerToken) > 0 {
		groups = append(groups, payPerToken)
		groupNames = append(groupNames, "pay_per_token")
	}

	// Get strategy for this model (hot-reloadable)
	st := h.DefaultStrategy
	if h.GetStrategy != nil {
		if s := h.GetStrategy(ir.Model); s != nil {
			st = s
		}
	}

	originalModel := ir.Model
	var lastErrMsg string
	for i, group := range groups {
		ir.Model = originalModel
		selected := st.Select(r.Context(), group, h.Registry)
		if selected == nil {
			continue
		}
		log.Printf("[proxy] model=%s strategy=%s tier=%s backend=%s(%s)", ir.Model, st.Name(), groupNames[i], selected.Name, selected.ID)
		ir.Model = selected.ResolveModel(ir.Model)

		ok, retryable, errMsg := h.executeBackend(w, r, ir, clientFormat, selected, startTime, body)
		if ok {
			return
		}
		if !retryable {
			writeError(w, clientFormat, errMsg)
			return
		}
		lastErrMsg = errMsg
		log.Printf("[proxy] retryable error on %s (%s), trying next tier", selected.Name, errMsg)
	}

	if lastErrMsg != "" {
		writeError(w, clientFormat, lastErrMsg)
	} else {
		writeError(w, clientFormat, "all backends unavailable for model "+ir.Model)
	}
}

// executeBackend sends the request to a single backend and handles the response.
// Returns (success, retryable, errorMessage).
func (h *Handler) executeBackend(
	w http.ResponseWriter,
	r *http.Request,
	ir *protocol.RequestIR,
	clientFormat string,
	selected *backend.Backend,
	startTime time.Time,
	rawBody []byte,
) (bool, bool, string) {
	var method string
	var urlPath string
	var reqBody []byte
	var headers http.Header
	var err error
	var adapter protocol.ProviderAdapter

	if formatsMatch(clientFormat, selected.Provider) {
		// Fast-path: replace model in raw JSON body, skip full IR translation
		reqBody, err = replaceModelInBody(rawBody, ir.Model)
		if err != nil {
			return false, false, "fast-path model replacement: " + err.Error()
		}
		method = "POST"
		if clientFormat == "anthropic" {
			urlPath = "/v1/messages"
		} else {
			urlPath = "/v1/chat/completions"
		}
		headers = http.Header{}
		headers.Set("Content-Type", "application/json")
		if selected.Provider == "anthropic" {
			headers.Set("x-api-key", selected.APIKey)
			headers.Set("anthropic-version", "2023-06-01")
		} else {
			headers.Set("Authorization", "Bearer "+selected.APIKey)
		}
	} else {
		adapter = h.adapterFor(selected)
		method, urlPath, reqBody, headers, err = adapter.BuildRequest(ir)
		if err != nil {
			return false, false, "build provider request: " + err.Error()
		}
	}

	backendURL := strings.TrimRight(selected.BaseURL, "/") + urlPath
	req, err := http.NewRequestWithContext(r.Context(), method, backendURL, bytes.NewReader(reqBody))
	if err != nil {
		return false, false, "create backend request: " + err.Error()
	}
	for k, vs := range headers {
		for _, v := range vs {
			req.Header.Add(k, v)
		}
	}

	timeout := time.Duration(selected.Timeout) * time.Second
	if selected.Timeout == 0 {
		timeout = 120 * time.Second
	}
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()
	req = req.WithContext(ctx)

	client := h.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		h.Registry.RecordResult(selected.ID, time.Since(startTime), true)
		return false, true, "backend request failed: " + err.Error()
	}
	defer resp.Body.Close()

	// Check for retryable HTTP status codes BEFORE writing to response writer
	if isRetryableStatus(resp.StatusCode) {
		h.Registry.RecordResult(selected.ID, time.Since(startTime), true)
		io.Copy(io.Discard, resp.Body)
		return false, true, fmt.Sprintf("backend %s returned status %d", selected.Name, resp.StatusCode)
	}

	// From this point on we are committed — we may write to w
	if ir.Stream {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		for _, hname := range []string{"X-Request-Id", "Anthropic-Ratelimit-Requests-Limit", "Anthropic-Ratelimit-Requests-Remaining", "Anthropic-Ratelimit-Tokens-Limit", "Anthropic-Ratelimit-Tokens-Remaining"} {
			if v := resp.Header.Get(hname); v != "" {
				w.Header().Set(hname, v)
			}
		}
		w.WriteHeader(resp.StatusCode)
		flusher, ok := w.(http.Flusher)
		if clientFormat == "anthropic" {
			h.streamAnthropicFiltered(w, resp.Body, flusher, ok)
		} else {
			buf := make([]byte, 4096)
			for {
				n, err := resp.Body.Read(buf)
				if n > 0 {
					w.Write(buf[:n])
					if ok {
						flusher.Flush()
					}
				}
				if err != nil {
					break
				}
			}
		}
		h.Registry.RecordResult(selected.ID, time.Since(startTime), resp.StatusCode >= 400)
		return true, false, ""
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		h.Registry.RecordResult(selected.ID, time.Since(startTime), true)
		return false, false, "read backend response: " + err.Error()
	}

	h.Registry.RecordResult(selected.ID, time.Since(startTime), resp.StatusCode >= 400)
	h.Registry.CaptureRateLimits(selected.ID, resp)
	log.Printf("[proxy] backend=%s status=%d len=%d", selected.Name, resp.StatusCode, len(respBody))

	// Fast-path: when client format matches backend format, skip full IR translation
	// and forward the raw response body directly.
	if resp.StatusCode < 400 && formatsMatch(clientFormat, selected.Provider) {
		// Minimal usage extraction for token tracking (much cheaper than full parse)
		var minimal struct {
			Usage struct {
				InputTokens  int `json:"input_tokens"`
				OutputTokens int `json:"output_tokens"`
			} `json:"usage"`
		}
		if json.Unmarshal(respBody, &minimal) == nil {
			total := minimal.Usage.InputTokens + minimal.Usage.OutputTokens
			if total > 0 {
				h.Registry.RecordTokens(selected.ID, int64(total))
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(respBody)
		return true, false, ""
	}

	responseIR, err := adapter.ParseResponse(resp.StatusCode, respBody)
	if err != nil {
		bodyPreview := string(respBody)
		if len(bodyPreview) > 500 {
			bodyPreview = bodyPreview[:500] + "..."
		}
		log.Printf("[proxy] parse error backend=%s status=%d body=%q err=%v", selected.Name, resp.StatusCode, bodyPreview, err)
		return false, false, "parse provider response: " + err.Error()
	}
	if responseIR.Error != nil {
		return false, false, responseIR.Error.Message
	}

	if responseIR.Usage.TotalTokens > 0 {
		h.Registry.RecordTokens(selected.ID, int64(responseIR.Usage.TotalTokens))
	} else if responseIR.Usage.PromptTokens+responseIR.Usage.CompletionTokens > 0 {
		h.Registry.RecordTokens(selected.ID, int64(responseIR.Usage.PromptTokens+responseIR.Usage.CompletionTokens))
	}

	if clientFormat == "openai" && len(responseIR.Content) > 0 && len(responseIR.Choices) == 0 {
		msgIR := &protocol.MessageIR{Role: "assistant", Content: responseIR.Content}
		responseIR.Choices = []protocol.ChoiceIR{{
			Index:        0,
			Message:      msgIR,
			FinishReason: responseIR.StopReason,
		}}
	}

	var outBody []byte
	switch clientFormat {
	case "openai":
		outBody, err = protocol.SerializeOpenAIResponse(responseIR, ir.Stream)
	case "anthropic":
		outBody, err = protocol.SerializeAnthropicResponse(responseIR)
	}
	if err != nil {
		return false, false, "serialize response: " + err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(outBody)
	return true, false, ""
}

func formatsMatch(clientFormat, provider string) bool {
	if clientFormat == "anthropic" && provider == "anthropic" {
		return true
	}
	if clientFormat == "openai" && (provider == "openai" || provider == "custom") {
		return true
	}
	return false
}

func (h *Handler) detectFormat(r *http.Request) string {
	path := r.URL.Path
	if strings.Contains(path, "/v1/messages") {
		return "anthropic"
	}
	// Check content-type for Anthropic-version header
	if r.Header.Get("anthropic-version") != "" {
		return "anthropic"
	}
	return "openai"
}

func (h *Handler) adapterFor(b *backend.Backend) protocol.ProviderAdapter {
	switch b.Provider {
	case "anthropic":
		return &protocol.AnthropicAdapter{
			BaseURL: b.BaseURL,
			APIKey:  b.APIKey,
		}
	default: // "openai", "custom"
		return &protocol.OpenAIAdapter{
			BaseURL: b.BaseURL,
			APIKey:  b.APIKey,
		}
	}
}

// HandleModels returns available models in OpenAI-compatible format.
func (h *Handler) HandleModels(w http.ResponseWriter, r *http.Request) {
	backends := h.Registry.List()
	modelSet := make(map[string]bool)
	for _, b := range backends {
		if b.Status == backend.StatusActive {
			for _, m := range b.Models {
				modelSet[m] = true
			}
		}
	}
	type modelEntry struct {
		ID     string `json:"id"`
		Object string `json:"object"`
	}
	models := make([]modelEntry, 0, len(modelSet))
	for m := range modelSet {
		models = append(models, modelEntry{ID: m, Object: "model"})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"object": "list",
		"data":   models,
	})
}

// estimateInputTokens computes a rough token count from the IR using a
// characters/3 heuristic. Over-estimates for safety.
func estimateInputTokens(ir *protocol.RequestIR) int {
	chars := len(ir.System)
	for _, msg := range ir.Messages {
		for _, block := range msg.Content {
			if block.Type == protocol.ContentTypeText {
				chars += len(block.Text)
			}
		}
	}
	if chars == 0 {
		return 0
	}
	return chars / 3
}

func writeError(w http.ResponseWriter, clientFormat, msg string) {
	errIR := &protocol.ResponseIR{
		Error: &protocol.ErrorIR{Message: msg, Type: "proxy_error"},
	}
	if clientFormat == "anthropic" {
		body, _ := protocol.SerializeAnthropicResponse(errIR)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
		w.Write(body)
	} else {
		resp := map[string]any{
			"error": map[string]any{
				"message": msg,
				"type":    "proxy_error",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(resp)
	}
}

// streamAnthropicFiltered reads Anthropic-format SSE from backend, strips
// thinking blocks, and writes the remaining events to the client.
func (h *Handler) streamAnthropicFiltered(w http.ResponseWriter, r io.Reader, flusher http.Flusher, canFlush bool) {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 4096), 1024*1024)

	thinkingIdx := make(map[int]bool) // indices that are thinking blocks
	var eventLines []string

	flushEvent := func(lines []string) {
		for _, line := range lines {
			w.Write([]byte(line + "\n"))
		}
		w.Write([]byte("\n"))
		if canFlush {
			flusher.Flush()
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// End of event — decide whether to forward or drop
			if len(eventLines) == 0 {
				continue
			}
			var dataLine string
			for _, l := range eventLines {
				if strings.HasPrefix(l, "data: ") {
					dataLine = strings.TrimPrefix(l, "data: ")
					break
				}
			}
			skip := false
			if dataLine != "" {
				var base struct {
					Type  string `json:"type"`
					Index int    `json:"index"`
				}
				if json.Unmarshal([]byte(dataLine), &base) == nil {
					switch base.Type {
					case "content_block_start":
						var detail struct {
							ContentBlock struct {
								Type string `json:"type"`
							} `json:"content_block"`
						}
						if json.Unmarshal([]byte(dataLine), &detail) == nil && detail.ContentBlock.Type == "thinking" {
							thinkingIdx[base.Index] = true
							skip = true
						}
					case "content_block_delta":
						var detail struct {
							Delta struct {
								Type string `json:"type"`
							} `json:"delta"`
						}
						if json.Unmarshal([]byte(dataLine), &detail) == nil && detail.Delta.Type == "thinking_delta" {
							skip = true
						}
					case "content_block_stop":
						if thinkingIdx[base.Index] {
							delete(thinkingIdx, base.Index)
							skip = true
						}
					}
				}
			}
			if !skip {
				flushEvent(eventLines)
			}
			eventLines = nil
		} else {
			eventLines = append(eventLines, line)
		}
	}
	// Flush any remaining event without trailing blank line
	if len(eventLines) > 0 {
		flushEvent(eventLines)
	}
}

func splitByCostTier(backends []*backend.Backend) (prepaid, payPerToken []*backend.Backend) {
	for _, b := range backends {
		if b.CostTier == "prepaid" {
			prepaid = append(prepaid, b)
		} else {
			payPerToken = append(payPerToken, b)
		}
	}
	return
}

func isRetryableStatus(code int) bool {
	return code == 408 || code == 429 || code >= 502
}

// replaceModelInBody replaces the "model" field in a JSON body with the given model name.
// Used for fast-path when client and backend formats match.
func replaceModelInBody(body []byte, model string) ([]byte, error) {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, err
	}
	modelJSON, _ := json.Marshal(model)
	raw["model"] = modelJSON
	return json.Marshal(raw)
}
