package protocol

import (
	"encoding/json"
	"fmt"
	"strings"
)

// --- Anthropic Request Parsing ---

type anthropicRequest struct {
	Model       string               `json:"model"`
	System      any                  `json:"system,omitempty"` // string or []anthropicTextBlock
	Messages    []anthropicMessage   `json:"messages"`
	Tools       []anthropicTool      `json:"tools,omitempty"`
	ToolChoice  any                  `json:"tool_choice,omitempty"`
	MaxTokens   int                  `json:"max_tokens"`
	Temperature *float64             `json:"temperature,omitempty"`
	TopP        *float64             `json:"top_p,omitempty"`
	StopSeq     []string             `json:"stop_sequences,omitempty"`
	Stream      bool                 `json:"stream"`
}

type anthropicMessage struct {
	Role    string               `json:"role"`
	Content any                  `json:"content"` // string or []anthropicContentBlock
}

type anthropicContentBlock struct {
	Type         string          `json:"type"`
	Text         string          `json:"text,omitempty"`
	ID           string          `json:"id,omitempty"`
	Name         string          `json:"name,omitempty"`
	Input        json.RawMessage `json:"input,omitempty"`
	ToolUseID    string          `json:"tool_use_id,omitempty"`
	Content      any             `json:"content,omitempty"`
	Source       *anthropicSource `json:"source,omitempty"`
}

type anthropicSource struct {
	Type      string `json:"type"`
	MediaType string `json:"media_type"`
	Data      string `json:"data"`
}

type anthropicTool struct {
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	InputSchema json.RawMessage `json:"input_schema"`
}

// ParseAnthropicRequest parses an Anthropic-format JSON body into canonical IR.
func ParseAnthropicRequest(body []byte) (*RequestIR, error) {
	var req anthropicRequest
	if err := json.Unmarshal(body, &req); err != nil {
		return nil, fmt.Errorf("parse anthropic request: %w", err)
	}
	ir := &RequestIR{
		Model:     req.Model,
		MaxTokens: req.MaxTokens,
		Stream:    req.Stream,
		StopSeq:   req.StopSeq,
		ToolChoice: req.ToolChoice,
	}
	if req.Temperature != nil {
		ir.Temperature = *req.Temperature
	}
	if req.TopP != nil {
		ir.TopP = *req.TopP
	}
	// Parse system prompt
	switch s := req.System.(type) {
	case string:
		ir.System = s
	case []any:
		var parts []string
		for _, raw := range s {
			b, _ := json.Marshal(raw)
			var cb anthropicContentBlock
			if json.Unmarshal(b, &cb) == nil && cb.Type == "text" {
				parts = append(parts, cb.Text)
			}
		}
		ir.System = strings.Join(parts, "\n")
	}
	// Parse tools
	for _, t := range req.Tools {
		ir.Tools = append(ir.Tools, ToolDefIR{
			Name:        t.Name,
			Description: t.Description,
			Parameters:  t.InputSchema,
		})
	}
	// Parse messages
	for _, m := range req.Messages {
		msgIR := MessageIR{Role: m.Role}
		switch c := m.Content.(type) {
		case string:
			if c != "" {
				msgIR.Content = []ContentBlock{{Type: ContentTypeText, Text: c}}
			}
		case []any:
			for _, raw := range c {
				b, _ := json.Marshal(raw)
				var cb anthropicContentBlock
				if err := json.Unmarshal(b, &cb); err != nil {
					continue
				}
				switch cb.Type {
				case "text":
					msgIR.Content = append(msgIR.Content, ContentBlock{Type: ContentTypeText, Text: cb.Text})
				case "image":
					if cb.Source != nil {
						msgIR.Content = append(msgIR.Content, ContentBlock{
							Type:           ContentTypeImage,
							ImageMediaType: cb.Source.MediaType,
							ImageData:      cb.Source.Data,
						})
					}
				case "tool_use":
					msgIR.Content = append(msgIR.Content, ContentBlock{
						Type:     ContentTypeToolUse,
						ToolID:   cb.ID,
						ToolName: cb.Name,
						ToolArgs: cb.Input,
					})
				case "tool_result":
					text := ""
					if cstr, ok := cb.Content.(string); ok {
						text = cstr
					}
					msgIR.Content = append(msgIR.Content, ContentBlock{
						Type:   ContentTypeToolResult,
						ToolID: cb.ToolUseID,
						Text:   text,
					})
				}
			}
		}
		ir.Messages = append(ir.Messages, msgIR)
	}
	return ir, nil
}

// --- Anthropic Response Serialization ---

type anthropicResponse struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Role       string                 `json:"role"`
	Model      string                 `json:"model"`
	Content    []anthropicContentBlock `json:"content"`
	StopReason *string                `json:"stop_reason"`
	Usage      anthropicUsage         `json:"usage"`
	Error      *anthropicError        `json:"error,omitempty"`
}

type anthropicUsage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}

type anthropicError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// SerializeAnthropicResponse converts canonical IR to Anthropic-format JSON.
func SerializeAnthropicResponse(ir *ResponseIR) ([]byte, error) {
	resp := anthropicResponse{
		ID:    ir.ID,
		Type:  "message",
		Role:  "assistant",
		Model: ir.Model,
		Usage: anthropicUsage{
			InputTokens:  ir.Usage.PromptTokens,
			OutputTokens: ir.Usage.CompletionTokens,
		},
	}
	if ir.Error != nil {
		resp.Error = &anthropicError{Type: ir.Error.Type, Message: ir.Error.Message}
		return json.Marshal(resp)
	}
	if ir.StopReason != "" {
		resp.StopReason = &ir.StopReason
	}
	for _, c := range ir.Content {
		switch c.Type {
		case ContentTypeText:
			resp.Content = append(resp.Content, anthropicContentBlock{
				Type: "text",
				Text: c.Text,
			})
		case ContentTypeToolUse:
			resp.Content = append(resp.Content, anthropicContentBlock{
				Type:  "tool_use",
				ID:    c.ToolID,
				Name:  c.ToolName,
				Input: c.ToolArgs,
			})
		}
	}
	return json.Marshal(resp)
}

// --- Anthropic Streaming Serializer ---

// WriteAnthropicStreamEvent writes a single SSE event in Anthropic format.
// Handles: message_start, content_block_start, content_block_delta,
// content_block_stop, message_delta, message_stop, ping, error.
func WriteAnthropicStreamEvent(eventType string, data any) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("event: %s\ndata: %s\n\n", eventType, string(b)), nil
}

type anthropicMsgStart struct {
	Type    string                 `json:"type"`
	Message anthropicStreamMessage `json:"message"`
}

type anthropicStreamMessage struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Role   string `json:"role"`
	Model  string `json:"model"`
	Usage  anthropicUsage `json:"usage"`
}

type anthropicBlockStart struct {
	Type         string `json:"type"`
	Index        int    `json:"index"`
	ContentBlock anthropicContentBlock `json:"content_block"`
}

type anthropicBlockDelta struct {
	Type  string        `json:"type"`
	Index int           `json:"index"`
	Delta anthropicDelta `json:"delta"`
}

type anthropicDelta struct {
	Type         string `json:"type"`
	Text         string `json:"text,omitempty"`
	PartialJSON  string `json:"partial_json,omitempty"`
}

type anthropicBlockStop struct {
	Type  string `json:"type"`
	Index int    `json:"index"`
}

type anthropicMsgDelta struct {
	Type       string  `json:"type"`
	Delta      anthropicStopDelta `json:"delta"`
	Usage      anthropicUsage     `json:"usage"`
}

type anthropicStopDelta struct {
	StopReason string `json:"stop_reason"`
}
