package vertexai

import (
	relaymodel "github.com/labring/aiproxy/core/relay/model"
)

type Request struct {
	AnthropicVersion string `json:"anthropic_version"`
	*relaymodel.ClaudeRequest
}
