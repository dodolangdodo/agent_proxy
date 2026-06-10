package protocol

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// AnthropicAdapter translates IR to/from Anthropic API format.
type AnthropicAdapter struct {
	BaseURL     string
	APIKey      string
	AnthropicVersion string
}

const defaultAnthropicVersion = "2023-06-01"

func (a *AnthropicAdapter) ProviderName() string { return "anthropic" }

func (a *AnthropicAdapter) BuildRequest(ir *RequestIR) (string, string, []byte, http.Header, error) {
	version := a.AnthropicVersion
	if version == "" {
		version = defaultAnthropicVersion
	}

	req := anthropicRequest{
		Model:     ir.Model,
		MaxTokens: ir.MaxTokens,
		Stream:    ir.Stream,
		StopSeq:   ir.StopSeq,
		ToolChoice: ir.ToolChoice,
	}
	if ir.MaxTokens == 0 {
		req.MaxTokens = 4096
	}
	if ir.Temperature != 0 {
		req.Temperature = &ir.Temperature
	}
	if ir.TopP != 0 {
		req.TopP = &ir.TopP
	}
	if ir.System != "" {
		req.System = ir.System
	}

	// Messages: filter system role (handled above), build Anthropic messages
	for _, m := range ir.Messages {
		if m.Role == "system" {
			if ir.System == "" {
				for _, b := range m.Content {
					if b.Type == ContentTypeText {
						req.System = b.Text
						break
					}
				}
			}
			continue
		}
		am := anthropicMessage{Role: m.Role}
		var blocks []anthropicContentBlock
		for _, b := range m.Content {
			switch b.Type {
			case ContentTypeText:
				blocks = append(blocks, anthropicContentBlock{Type: "text", Text: b.Text})
			case ContentTypeImage:
				blocks = append(blocks, anthropicContentBlock{
					Type: "image",
					Source: &anthropicSource{
						Type:      "base64",
						MediaType: b.ImageMediaType,
						Data:      b.ImageData,
					},
				})
			case ContentTypeToolUse:
				blocks = append(blocks, anthropicContentBlock{
					Type:  "tool_use",
					ID:    b.ToolID,
					Name:  b.ToolName,
					Input: b.ToolArgs,
				})
			case ContentTypeToolResult:
				blocks = append(blocks, anthropicContentBlock{
					Type:      "tool_result",
					ToolUseID: b.ToolID,
					Content:   b.Text,
				})
			}
		}
		if len(blocks) == 1 && blocks[0].Type == "text" {
			am.Content = blocks[0].Text
		} else {
			am.Content = blocks
		}
		req.Messages = append(req.Messages, am)
	}

	for _, t := range ir.Tools {
		req.Tools = append(req.Tools, anthropicTool{
			Name:        t.Name,
			Description: t.Description,
			InputSchema: t.Parameters,
		})
	}

	body, err := json.Marshal(req)
	if err != nil {
		return "", "", nil, nil, fmt.Errorf("anthropic marshal: %w", err)
	}

	h := http.Header{}
	h.Set("Content-Type", "application/json")
	h.Set("x-api-key", a.APIKey)
	h.Set("anthropic-version", version)

	return "POST", "/v1/messages", body, h, nil
}

func (a *AnthropicAdapter) ParseResponse(httpStatus int, body []byte) (*ResponseIR, error) {
	if httpStatus >= 400 {
		var wrapper struct {
			Error anthropicError `json:"error"`
		}
		if json.Unmarshal(body, &wrapper) == nil && wrapper.Error.Message != "" {
			return &ResponseIR{Error: &ErrorIR{Message: wrapper.Error.Message, Type: wrapper.Error.Type}}, nil
		}
		var aerr anthropicError
		if json.Unmarshal(body, &aerr) == nil && aerr.Message != "" {
			return &ResponseIR{Error: &ErrorIR{Message: aerr.Message, Type: aerr.Type}}, nil
		}
		bodyStr := string(body)
		if len(bodyStr) > 500 {
			bodyStr = bodyStr[:500] + "..."
		}
		return &ResponseIR{Error: &ErrorIR{Message: bodyStr, Type: "http_error"}}, nil
	}

	var resp anthropicResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		bodyStr := string(body)
		if len(bodyStr) > 500 {
			bodyStr = bodyStr[:500] + "..."
		}
		return nil, fmt.Errorf("anthropic parse response (status=%d body=%q): %w", httpStatus, bodyStr, err)
	}

	ir := &ResponseIR{
		ID:    resp.ID,
		Model: resp.Model,
		Usage: UsageIR{
			PromptTokens:     resp.Usage.InputTokens,
			CompletionTokens: resp.Usage.OutputTokens,
			TotalTokens:      resp.Usage.InputTokens + resp.Usage.OutputTokens,
		},
	}
	if resp.StopReason != nil {
		ir.StopReason = *resp.StopReason
	}
	for _, cb := range resp.Content {
		// Strip thinking blocks — clients don't expect them from non-Anthropic backends
		if cb.Type == "thinking" {
			continue
		}
		switch cb.Type {
		case "text":
			ir.Content = append(ir.Content, ContentBlock{Type: ContentTypeText, Text: cb.Text})
		case "tool_use":
			ir.Content = append(ir.Content, ContentBlock{
				Type:     ContentTypeToolUse,
				ToolID:   cb.ID,
				ToolName: cb.Name,
				ToolArgs: cb.Input,
			})
		}
	}
	return ir, nil
}

func (a *AnthropicAdapter) SupportsStreaming() bool { return true }

func (a *AnthropicAdapter) ParseStreamEvent(data []byte) (*StreamEvent, error) {
	// Anthropic SSE format: "data: {...}"
	content := string(data)
	if !strings.HasPrefix(content, "{") {
		return nil, nil
	}

	var base struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &base); err != nil {
		return &StreamEvent{}, nil
	}

	switch base.Type {
	case "message_start":
		var ev struct {
			Message struct {
				ID    string `json:"id"`
				Model string `json:"model"`
				Usage anthropicUsage `json:"usage"`
			} `json:"message"`
		}
		json.Unmarshal(data, &ev)
		return &StreamEvent{Event: "message_start", Data: ev}, nil

	case "content_block_start":
		var ev struct {
			Index        int                   `json:"index"`
			ContentBlock anthropicContentBlock `json:"content_block"`
		}
		json.Unmarshal(data, &ev)
		var cb ContentBlock
		switch ev.ContentBlock.Type {
		case "text":
			cb = ContentBlock{Type: ContentTypeText}
		case "tool_use":
			cb = ContentBlock{Type: ContentTypeToolUse, ToolID: ev.ContentBlock.ID, ToolName: ev.ContentBlock.Name}
		}
		return &StreamEvent{Event: "content_block_start", Data: cb}, nil

	case "content_block_delta":
		var ev struct {
			Index int             `json:"index"`
			Delta anthropicDelta `json:"delta"`
		}
		json.Unmarshal(data, &ev)
		return &StreamEvent{Event: "content_block_delta", Data: ContentBlock{
			Type: ContentTypeText, Text: ev.Delta.Text,
		}}, nil

	case "content_block_stop":
		return &StreamEvent{Event: "content_block_stop", Data: nil}, nil

	case "message_delta":
		var ev struct {
			Usage anthropicUsage     `json:"usage"`
			Delta anthropicStopDelta `json:"delta"`
		}
		json.Unmarshal(data, &ev)
		return &StreamEvent{Event: "message_delta", Data: map[string]string{
			"stop_reason": ev.Delta.StopReason,
		}}, nil

	case "message_stop":
		return nil, nil // stream complete

	case "error":
		var ev struct {
			Error anthropicError `json:"error"`
		}
		json.Unmarshal(data, &ev)
		return &StreamEvent{Event: "error", Data: ev.Error}, nil
	}
	return &StreamEvent{}, nil
}
