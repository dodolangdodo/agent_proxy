package model

import "github.com/labring/aiproxy/core/model"

type TextToSpeechRequest struct {
	Model          string  `json:"model"           binding:"required"`
	Input          string  `json:"input"           binding:"required"`
	Voice          string  `json:"voice"           binding:"required"`
	ResponseFormat string  `json:"response_format"`
	Speed          float64 `json:"speed"`
	StreamFormat   string  `json:"stream_format"`
}

const (
	TextToSpeechSSEResponseTypeDelta = "speech.audio.delta"
	TextToSpeechSSEResponseTypeDone  = "speech.audio.done"
)

type TextToSpeechSSEResponse struct {
	Type  string             `json:"type"` // constant: TextToSpeechSSEResponseType
	Audio string             `json:"audio,omitempty"`
	Usage *TextToSpeechUsage `json:"usage,omitempty"`
}

type TextToSpeechUsage struct {
	InputTokens  int64 `json:"input_tokens"`
	OutputTokens int64 `json:"output_tokens"`
	TotalTokens  int64 `json:"total_tokens"`
}

func (u *TextToSpeechUsage) ToModelUsage() model.Usage {
	return model.Usage{
		InputTokens:       model.ZeroNullInt64(u.InputTokens),
		OutputTokens:      model.ZeroNullInt64(u.OutputTokens),
		AudioOutputTokens: model.ZeroNullInt64(u.OutputTokens),
		TotalTokens:       model.ZeroNullInt64(u.TotalTokens),
	}
}
