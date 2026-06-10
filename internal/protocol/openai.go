package protocol

import (
	"encoding/json"
	"fmt"
)

// --- OpenAI Request Parsing ---

type openaiChatRequest struct {
	Model       string             `json:"model"`
	Messages    []openaiMessage    `json:"messages"`
	Tools       []openaiTool       `json:"tools,omitempty"`
	ToolChoice  any                `json:"tool_choice,omitempty"`
	MaxTokens   *int               `json:"max_tokens,omitempty"`
	Temperature *float64           `json:"temperature,omitempty"`
	TopP        *float64           `json:"top_p,omitempty"`
	Stop        any                `json:"stop,omitempty"`
	Stream      bool               `json:"stream"`
}

type openaiMessage struct {
	Role       string              `json:"role"`
	Content    any                 `json:"content"` // string or []openaiContentBlock
	Name       string              `json:"name,omitempty"`
	ToolCalls  []openaiToolCall    `json:"tool_calls,omitempty"`
	ToolCallID string              `json:"tool_call_id,omitempty"`
}

type openaiContentBlock struct {
	Type     string              `json:"type"`
	Text     string              `json:"text,omitempty"`
	ImageURL *openaiImageURL     `json:"image_url,omitempty"`
}

type openaiImageURL struct {
	URL    string `json:"url"`
	Detail string `json:"detail,omitempty"`
}

type openaiTool struct {
	Type     string          `json:"type"`
	Function openaiFunction  `json:"function"`
}

type openaiFunction struct {
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	Parameters  json.RawMessage `json:"parameters,omitempty"`
}

type openaiToolCall struct {
	ID       string              `json:"id"`
	Type     string              `json:"type"`
	Function openaiToolCallFunc  `json:"function"`
}

type openaiToolCallFunc struct {
	Name      string          `json:"name"`
	Arguments string          `json:"arguments"`
}

// ParseOpenAIRequest parses an OpenAI-format JSON body into canonical IR.
func ParseOpenAIRequest(body []byte) (*RequestIR, error) {
	var req openaiChatRequest
	if err := json.Unmarshal(body, &req); err != nil {
		return nil, fmt.Errorf("parse openai request: %w", err)
	}
	ir := &RequestIR{
		Model:      req.Model,
		Stream:     req.Stream,
		ToolChoice: req.ToolChoice,
	}
	if req.MaxTokens != nil {
		ir.MaxTokens = *req.MaxTokens
	}
	if req.Temperature != nil {
		ir.Temperature = *req.Temperature
	}
	if req.TopP != nil {
		ir.TopP = *req.TopP
	}
	if s, ok := req.Stop.(string); ok {
		ir.StopSeq = []string{s}
	} else if ss, ok := req.Stop.([]string); ok {
		ir.StopSeq = ss
	}
	for _, t := range req.Tools {
		ir.Tools = append(ir.Tools, ToolDefIR{
			Name:        t.Function.Name,
			Description: t.Function.Description,
			Parameters:  t.Function.Parameters,
		})
	}
	for _, m := range req.Messages {
		msgIR := MessageIR{Role: m.Role, Name: m.Name}
		switch c := m.Content.(type) {
		case string:
			if c != "" {
				msgIR.Content = []ContentBlock{{Type: ContentTypeText, Text: c}}
			}
		case []any:
			for _, raw := range c {
				b, _ := json.Marshal(raw)
				var cb openaiContentBlock
				if json.Unmarshal(b, &cb) == nil && cb.Type == "image_url" && cb.ImageURL != nil {
					msgIR.Content = append(msgIR.Content, ContentBlock{
						Type:     ContentTypeImage,
						ImageURL: cb.ImageURL.URL,
					})
				} else if cb.Type == "text" {
					msgIR.Content = append(msgIR.Content, ContentBlock{
						Type: ContentTypeText,
						Text: cb.Text,
					})
				}
			}
		}
		for _, tc := range m.ToolCalls {
			msgIR.Content = append(msgIR.Content, ContentBlock{
				Type:     ContentTypeToolUse,
				ToolID:   tc.ID,
				ToolName: tc.Function.Name,
				ToolArgs: json.RawMessage(tc.Function.Arguments),
			})
		}
		if m.ToolCallID != "" {
			msgIR.Content = append(msgIR.Content, ContentBlock{
				Type:   ContentTypeToolResult,
				ToolID: m.ToolCallID,
			})
		}
		ir.Messages = append(ir.Messages, msgIR)
	}
	return ir, nil
}

// --- OpenAI Response Serialization ---

type openaiResponse struct {
	ID      string            `json:"id"`
	Object  string            `json:"object"`
	Model   string            `json:"model"`
	Choices []openaiChoice    `json:"choices"`
	Usage   openaiUsage       `json:"usage"`
	Error   *openaiError      `json:"error,omitempty"`
}

type openaiChoice struct {
	Index        int             `json:"index"`
	Message      *openaiRespMsg  `json:"message,omitempty"`
	Delta        *openaiRespMsg  `json:"delta,omitempty"`
	FinishReason *string         `json:"finish_reason,omitempty"`
}

type openaiRespMsg struct {
	Role      string           `json:"role,omitempty"`
	Content   any              `json:"content,omitempty"`
	ToolCalls []openaiToolCall `json:"tool_calls,omitempty"`
}

type openaiUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type openaiError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    string `json:"code,omitempty"`
}

// SerializeOpenAIResponse converts canonical IR to OpenAI-format JSON.
func SerializeOpenAIResponse(ir *ResponseIR, stream bool) ([]byte, error) {
	resp := openaiResponse{
		ID:     ir.ID,
		Object: "chat.completion",
		Model:  ir.Model,
		Usage: openaiUsage{
			PromptTokens:     ir.Usage.PromptTokens,
			CompletionTokens: ir.Usage.CompletionTokens,
			TotalTokens:      ir.Usage.TotalTokens,
		},
	}
	if ir.Error != nil {
		resp.Error = &openaiError{
			Message: ir.Error.Message,
			Type:    ir.Error.Type,
			Code:    ir.Error.Code,
		}
		return json.Marshal(resp)
	}
	for _, c := range ir.Choices {
		oc := openaiChoice{
			Index:        c.Index,
			FinishReason: strPtr(c.FinishReason),
		}
		if c.Message != nil {
			oc.Message = irMsgToOpenAI(c.Message)
		}
		if c.Delta != nil {
			oc.Delta = irMsgToOpenAI(c.Delta)
		}
		resp.Choices = append(resp.Choices, oc)
	}
	return json.Marshal(resp)
}

func irMsgToOpenAI(msg *MessageIR) *openaiRespMsg {
	om := &openaiRespMsg{Role: msg.Role}
	var textParts []string
	for _, b := range msg.Content {
		switch b.Type {
		case ContentTypeText:
			textParts = append(textParts, b.Text)
		case ContentTypeToolUse:
			om.ToolCalls = append(om.ToolCalls, openaiToolCall{
				ID:   b.ToolID,
				Type: "function",
				Function: openaiToolCallFunc{
					Name:      b.ToolName,
					Arguments: string(b.ToolArgs),
				},
			})
		}
	}
	if len(om.ToolCalls) == 0 && len(textParts) == 1 {
		om.Content = textParts[0]
	} else if len(textParts) > 0 {
		om.Content = textParts[0]
	}
	return om
}

func strPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
