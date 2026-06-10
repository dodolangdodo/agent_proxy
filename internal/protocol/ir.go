package protocol

import "encoding/json"

// Canonical Intermediate Representation for all LLM protocol formats.
// Superset struct that carries all features; adapters drop/approximate unsupported fields.

type RequestIR struct {
	Model       string           `json:"model"`
	System      string           `json:"system,omitempty"`
	Messages    []MessageIR      `json:"messages"`
	Tools       []ToolDefIR      `json:"tools,omitempty"`
	ToolChoice  any              `json:"tool_choice,omitempty"`
	MaxTokens   int              `json:"max_tokens,omitempty"`
	Temperature float64          `json:"temperature,omitempty"`
	TopP        float64          `json:"top_p,omitempty"`
	StopSeq     []string         `json:"stop,omitempty"`
	Stream      bool             `json:"stream"`
	Extensions  map[string]any   `json:"-"` // protocol-specific extras
}

type MessageIR struct {
	Role    string          `json:"role"`
	Content []ContentBlock  `json:"content"`
	Name    string          `json:"name,omitempty"`
}

type ContentType string

const (
	ContentTypeText       ContentType = "text"
	ContentTypeImage      ContentType = "image"
	ContentTypeToolUse    ContentType = "tool_use"
	ContentTypeToolResult ContentType = "tool_result"
)

type ContentBlock struct {
	Type     ContentType     `json:"type"`
	Text     string          `json:"text,omitempty"`
	ToolID   string          `json:"tool_id,omitempty"`
	ToolName string          `json:"tool_name,omitempty"`
	ToolArgs json.RawMessage `json:"tool_args,omitempty"`
	ImageURL string          `json:"image_url,omitempty"`
	ImageMediaType string    `json:"image_media_type,omitempty"`
	ImageData      string    `json:"image_data,omitempty"`
}

type ToolDefIR struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Parameters  json.RawMessage `json:"parameters"` // JSON Schema
}

type ResponseIR struct {
	ID        string        `json:"id"`
	Model     string        `json:"model"`
	Choices   []ChoiceIR    `json:"choices,omitempty"`
	Content   []ContentBlock `json:"content,omitempty"`
	Usage     UsageIR       `json:"usage"`
	StopReason string       `json:"stop_reason,omitempty"`
	Error     *ErrorIR      `json:"error,omitempty"`
}

type ChoiceIR struct {
	Index        int           `json:"index"`
	Message      *MessageIR    `json:"message,omitempty"`
	Delta        *MessageIR    `json:"delta,omitempty"`
	FinishReason string        `json:"finish_reason,omitempty"`
}

type UsageIR struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ErrorIR struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    string `json:"code,omitempty"`
}

// StreamEvent represents a single SSE event in the canonical form.
type StreamEvent struct {
	Event string // "message_start", "content_block_start", "content_block_delta", etc.
	Data  any
}
