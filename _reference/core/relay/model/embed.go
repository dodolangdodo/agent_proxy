package model

import "github.com/labring/aiproxy/core/model"

type EmbeddingRequest struct {
	Input          string `json:"input"`
	Model          string `json:"model"`
	EncodingFormat string `json:"encoding_format"`
	Dimensions     int    `json:"dimensions"`
}

type EmbeddingResponseItem struct {
	Object    string    `json:"object"`
	Embedding []float64 `json:"embedding"`
	Index     int       `json:"index"`
}

type EmbeddingResponse struct {
	Object string                   `json:"object"`
	Model  string                   `json:"model"`
	Data   []*EmbeddingResponseItem `json:"data"`
	Usage  EmbeddingUsage           `json:"usage"`
}

type EmbeddingUsage struct {
	PromptTokens        int64                         `json:"prompt_tokens,omitempty"`
	TotalTokens         int64                         `json:"total_tokens"`
	PromptTokensDetails *EmbeddingPromptTokensDetails `json:"prompt_tokens_details,omitempty"`
}

type EmbeddingPromptTokensDetails struct {
	TextTokens  int64 `json:"text_tokens,omitempty"`
	ImageTokens int64 `json:"image_tokens,omitempty"`
	VideoTokens int64 `json:"video_tokens,omitempty"`
}

func (u EmbeddingUsage) ToModelUsage() model.Usage {
	usage := model.Usage{
		InputTokens: model.ZeroNullInt64(u.PromptTokens),
		TotalTokens: model.ZeroNullInt64(u.TotalTokens),
	}
	if u.PromptTokensDetails != nil {
		usage.ImageInputTokens = model.ZeroNullInt64(u.PromptTokensDetails.ImageTokens)
		usage.VideoInputTokens = model.ZeroNullInt64(u.PromptTokensDetails.VideoTokens)
	}

	return usage
}
