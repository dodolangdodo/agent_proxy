package protocol

import "net/http"

// ProviderAdapter handles translation between canonical IR and a specific
// LLM provider's native API format (for calling downstream backends).
type ProviderAdapter interface {
	// ProviderName returns "openai", "anthropic", etc.
	ProviderName() string

	// BuildRequest converts IR to the provider's native request format.
	BuildRequest(ir *RequestIR) (method, urlPath string, body []byte, headers http.Header, err error)

	// ParseResponse converts the provider's HTTP response to canonical IR.
	ParseResponse(httpStatus int, body []byte) (*ResponseIR, error)

	// SupportsStreaming returns true if the provider supports SSE streaming.
	SupportsStreaming() bool

	// ParseStreamEvent converts one SSE data line from the provider into a
	// canonical StreamEvent. Returns nil when the stream is complete.
	ParseStreamEvent(data []byte) (*StreamEvent, error)
}
