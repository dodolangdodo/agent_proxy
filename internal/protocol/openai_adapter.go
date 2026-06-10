package protocol

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// OpenAIAdapter translates IR to/from OpenAI API format.
type OpenAIAdapter struct {
	BaseURL string
	APIKey  string
}

func (a *OpenAIAdapter) ProviderName() string { return "openai" }

func (a *OpenAIAdapter) BuildRequest(ir *RequestIR) (string, string, []byte, http.Header, error) {
	req := openaiChatRequest{
		Model:      ir.Model,
		Stream:     ir.Stream,
		ToolChoice: ir.ToolChoice,
	}
	if ir.MaxTokens > 0 {
		req.MaxTokens = &ir.MaxTokens
	}
	if ir.Temperature != 0 {
		req.Temperature = &ir.Temperature
	}
	if ir.TopP != 0 {
		req.TopP = &ir.TopP
	}
	if len(ir.StopSeq) == 1 {
		req.Stop = ir.StopSeq[0]
	} else if len(ir.StopSeq) > 1 {
		req.Stop = ir.StopSeq
	}

	// Messages: add system as first message if present
	messages := make([]openaiMessage, 0, len(ir.Messages)+1)
	if ir.System != "" {
		messages = append(messages, openaiMessage{
			Role:    "system",
			Content: ir.System,
		})
	}
	for _, m := range ir.Messages {
		om := openaiMessage{Role: m.Role, Name: m.Name}
		var textContent []string
		for _, b := range m.Content {
			switch b.Type {
			case ContentTypeText:
				textContent = append(textContent, b.Text)
			case ContentTypeImage:
				om.Content = []openaiContentBlock{
					{Type: "text", Text: strings.Join(textContent, "\n")},
					{Type: "image_url", ImageURL: &openaiImageURL{URL: b.ImageURL}},
				}
				textContent = nil
			case ContentTypeToolUse:
				if len(textContent) > 0 {
					om.Content = strings.Join(textContent, "\n")
					textContent = nil
				}
				om.ToolCalls = append(om.ToolCalls, openaiToolCall{
					ID:   b.ToolID,
					Type: "function",
					Function: openaiToolCallFunc{
						Name:      b.ToolName,
						Arguments: string(b.ToolArgs),
					},
				})
			case ContentTypeToolResult:
				om.ToolCallID = b.ToolID
				om.Content = b.Text
			}
		}
		if len(textContent) > 0 && om.Content == nil {
			om.Content = strings.Join(textContent, "\n")
		}
		messages = append(messages, om)
	}
	req.Messages = messages

	// Tools
	for _, t := range ir.Tools {
		req.Tools = append(req.Tools, openaiTool{
			Type: "function",
			Function: openaiFunction{
				Name:        t.Name,
				Description: t.Description,
				Parameters:  t.Parameters,
			},
		})
	}

	body, err := json.Marshal(req)
	if err != nil {
		return "", "", nil, nil, fmt.Errorf("openai marshal: %w", err)
	}

	h := http.Header{}
	h.Set("Content-Type", "application/json")
	if a.APIKey != "" {
		h.Set("Authorization", "Bearer "+a.APIKey)
	}

	return "POST", "/v1/chat/completions", body, h, nil
}

func (a *OpenAIAdapter) ParseResponse(httpStatus int, body []byte) (*ResponseIR, error) {
	if httpStatus >= 400 {
		var oerr openaiError
		if json.Unmarshal(body, &oerr) == nil {
			return &ResponseIR{Error: &ErrorIR{Message: oerr.Message, Type: oerr.Type}}, nil
		}
		bodyStr := string(body)
		if len(bodyStr) > 500 {
			bodyStr = bodyStr[:500] + "..."
		}
		return &ResponseIR{Error: &ErrorIR{Message: bodyStr, Type: "http_error"}}, nil
	}

	var resp openaiResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		bodyStr := string(body)
		if len(bodyStr) > 500 {
			bodyStr = bodyStr[:500] + "..."
		}
		return nil, fmt.Errorf("openai parse response (status=%d body=%q): %w", httpStatus, bodyStr, err)
	}

	ir := &ResponseIR{
		ID:    resp.ID,
		Model: resp.Model,
		Usage: UsageIR{
			PromptTokens:     resp.Usage.PromptTokens,
			CompletionTokens: resp.Usage.CompletionTokens,
			TotalTokens:      resp.Usage.TotalTokens,
		},
	}
	for _, c := range resp.Choices {
		ch := ChoiceIR{Index: c.Index}
		if c.FinishReason != nil {
			ch.FinishReason = *c.FinishReason
		}
		if c.Message != nil {
			ch.Message = openaiMsgToIR(c.Message)
		}
		ir.Choices = append(ir.Choices, ch)
	}
	return ir, nil
}

func openaiMsgToIR(m *openaiRespMsg) *MessageIR {
	msg := &MessageIR{Role: m.Role}
	if s, ok := m.Content.(string); ok && s != "" {
		msg.Content = append(msg.Content, ContentBlock{Type: ContentTypeText, Text: s})
	}
	for _, tc := range m.ToolCalls {
		msg.Content = append(msg.Content, ContentBlock{
			Type:     ContentTypeToolUse,
			ToolID:   tc.ID,
			ToolName: tc.Function.Name,
			ToolArgs: json.RawMessage(tc.Function.Arguments),
		})
	}
	return msg
}

func (a *OpenAIAdapter) SupportsStreaming() bool { return true }

func (a *OpenAIAdapter) ParseStreamEvent(data []byte) (*StreamEvent, error) {
	if string(data) == "[DONE]" {
		return nil, nil // stream complete
	}
	var chunk struct {
		ID     string `json:"id"`
		Model  string `json:"model"`
		Object string `json:"object"`
		Choices []struct {
			Index int `json:"index"`
			Delta struct {
				Role    string `json:"role,omitempty"`
				Content any    `json:"content,omitempty"`
				ToolCalls []struct {
					Index    int    `json:"index"`
					ID       string `json:"id,omitempty"`
					Type     string `json:"type,omitempty"`
					Function struct {
						Name      string `json:"name,omitempty"`
						Arguments string `json:"arguments,omitempty"`
					} `json:"function,omitempty"`
				} `json:"tool_calls,omitempty"`
			} `json:"delta"`
			FinishReason *string `json:"finish_reason"`
		} `json:"choices"`
	}
	if err := json.Unmarshal(data, &chunk); err != nil {
		return &StreamEvent{}, nil // ignore unparseable chunks
	}
	if len(chunk.Choices) == 0 {
		return &StreamEvent{}, nil
	}
	delta := chunk.Choices[0].Delta
	ev := &StreamEvent{Event: "delta"}
	if delta.Content != nil {
		s, _ := delta.Content.(string)
		ev.Data = ContentBlock{Type: ContentTypeText, Text: s}
	}
	for _, tc := range delta.ToolCalls {
		ev.Data = ContentBlock{
			Type:     ContentTypeToolUse,
			ToolID:   tc.ID,
			ToolName: tc.Function.Name,
			ToolArgs: json.RawMessage(tc.Function.Arguments),
		}
	}
	return ev, nil
}
