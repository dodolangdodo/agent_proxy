package utils

import (
	"bytes"
	"fmt"
	"io"

	"github.com/bytedance/sonic"
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/mode"
	relaymodel "github.com/labring/aiproxy/core/relay/model"
)

type UnsupportedModelTypeError struct {
	ModelType string
}

func (e *UnsupportedModelTypeError) Error() string {
	return fmt.Sprintf("model type '%s' not supported", e.ModelType)
}

func NewErrUnsupportedModelType(modelType string) *UnsupportedModelTypeError {
	return &UnsupportedModelTypeError{ModelType: modelType}
}

func BuildRequest(modelConfig model.ModelConfig) (io.Reader, mode.Mode, error) {
	switch modelConfig.Type {
	case mode.ChatCompletions:
		body, err := BuildChatCompletionRequest(modelConfig.Model)
		if err != nil {
			return nil, mode.Unknown, err
		}

		return body, mode.ChatCompletions, nil
	case mode.Completions:
		body, err := BuildCompletionsRequest(modelConfig.Model)
		if err != nil {
			return nil, mode.Unknown, err
		}

		return body, mode.Completions, nil
	case mode.Embeddings:
		body, err := BuildEmbeddingsRequest(modelConfig.Model)
		if err != nil {
			return nil, mode.Unknown, err
		}

		return body, mode.Embeddings, nil
	case mode.Moderations:
		body, err := BuildModerationsRequest(modelConfig.Model)
		if err != nil {
			return nil, mode.Unknown, err
		}

		return body, mode.Moderations, nil
	case mode.ImagesGenerations, mode.GeminiImage:
		body, err := BuildImagesGenerationsRequest(modelConfig)
		if err != nil {
			return nil, mode.Unknown, err
		}

		return body, mode.ImagesGenerations, nil
	case mode.ImagesEdits:
		return nil, mode.Unknown, NewErrUnsupportedModelType("edits")
	case mode.AudioSpeech, mode.GeminiTTS:
		body, err := BuildAudioSpeechRequest(modelConfig.Model)
		if err != nil {
			return nil, mode.Unknown, err
		}

		return body, mode.AudioSpeech, nil
	case mode.AudioTranscription:
		return nil, mode.Unknown, NewErrUnsupportedModelType("audio transcription")
	case mode.AudioTranslation:
		return nil, mode.Unknown, NewErrUnsupportedModelType("audio translation")
	case mode.Rerank:
		body, err := BuildRerankRequest(modelConfig.Model)
		if err != nil {
			return nil, mode.Unknown, err
		}

		return body, mode.Rerank, nil
	case mode.VideoGenerationsJobs:
		body, err := BuildVideoGenerationJobRequest(modelConfig.Model)
		if err != nil {
			return nil, mode.Unknown, err
		}

		return body, mode.VideoGenerationsJobs, nil
	case mode.Videos:
		body, err := BuildVideosRequest(modelConfig.Model)
		if err != nil {
			return nil, mode.Unknown, err
		}

		return body, mode.Videos, nil
	case mode.ParsePdf:
		return nil, mode.Unknown, NewErrUnsupportedModelType("parse pdf")
	case mode.GeminiVideo:
		body, err := BuildGeminiVideoRequest(modelConfig.Model)
		if err != nil {
			return nil, mode.Unknown, err
		}

		return body, mode.GeminiVideo, nil
	case mode.AliVideo:
		body, err := BuildAliVideoRequest(modelConfig.Model)
		if err != nil {
			return nil, mode.Unknown, err
		}

		return body, mode.AliVideo, nil
	case mode.DoubaoVideo:
		body, err := BuildDoubaoVideoRequest(modelConfig.Model)
		if err != nil {
			return nil, mode.Unknown, err
		}

		return body, mode.DoubaoVideo, nil
	default:
		return nil, mode.Unknown, NewErrUnsupportedModelType(modelConfig.Type.String())
	}
}

func BuildChatCompletionRequest(model string) (io.Reader, error) {
	testRequest := &relaymodel.GeneralOpenAIRequest{
		Model: model,
		Messages: []relaymodel.Message{
			{
				Role:    "user",
				Content: "hi",
			},
		},
	}

	jsonBytes, err := sonic.Marshal(testRequest)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func BuildCompletionsRequest(model string) (io.Reader, error) {
	completionsRequest := &relaymodel.GeneralOpenAIRequest{
		Model:  model,
		Prompt: "hi",
	}

	jsonBytes, err := sonic.Marshal(completionsRequest)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func BuildEmbeddingsRequest(model string) (io.Reader, error) {
	embeddingsRequest := &relaymodel.GeneralOpenAIRequest{
		Model: model,
		Input: "hi",
	}

	jsonBytes, err := sonic.Marshal(embeddingsRequest)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func BuildModerationsRequest(model string) (io.Reader, error) {
	moderationsRequest := &relaymodel.GeneralOpenAIRequest{
		Model: model,
		Input: "hi",
	}

	jsonBytes, err := sonic.Marshal(moderationsRequest)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func BuildImagesGenerationsRequest(modelConfig model.ModelConfig) (io.Reader, error) {
	imagesGenerationsRequest := &relaymodel.GeneralOpenAIRequest{
		Model:  modelConfig.Model,
		Prompt: "A simple red square icon on a white background.",
		Size:   "1024x1024",
	}

	jsonBytes, err := sonic.Marshal(imagesGenerationsRequest)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func BuildAudioSpeechRequest(model string) (io.Reader, error) {
	audioSpeechRequest := &relaymodel.GeneralOpenAIRequest{
		Model: model,
		Input: "hi",
	}

	jsonBytes, err := sonic.Marshal(audioSpeechRequest)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func BuildRerankRequest(model string) (io.Reader, error) {
	rerankRequest := &relaymodel.RerankRequest{
		Model:     model,
		Query:     "hi",
		Documents: []string{"hi"},
	}

	jsonBytes, err := sonic.Marshal(rerankRequest)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func BuildVideoGenerationJobRequest(model string) (io.Reader, error) {
	testRequest := map[string]any{
		"model":  model,
		"prompt": "A calm cinematic shot of clouds moving over a mountain.",
	}

	jsonBytes, err := sonic.Marshal(testRequest)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func BuildVideosRequest(model string) (io.Reader, error) {
	testRequest := &relaymodel.VideosRequest{
		Model:  model,
		Prompt: "A calm cinematic shot of clouds moving over a mountain.",
	}

	jsonBytes, err := sonic.Marshal(testRequest)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func BuildGeminiVideoRequest(_ string) (io.Reader, error) {
	testRequest := map[string]any{
		"instances": []map[string]any{
			{
				"prompt": "A calm cinematic shot of clouds moving over a mountain.",
			},
		},
	}

	jsonBytes, err := sonic.Marshal(testRequest)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func BuildAliVideoRequest(model string) (io.Reader, error) {
	testRequest := map[string]any{
		"model": model,
		"input": map[string]any{
			"prompt": "A calm cinematic shot of clouds moving over a mountain.",
		},
		"parameters": map[string]any{
			"duration": 5,
			"size":     "720P",
		},
	}

	jsonBytes, err := sonic.Marshal(testRequest)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func BuildDoubaoVideoRequest(model string) (io.Reader, error) {
	testRequest := map[string]any{
		"model": model,
		"content": []map[string]any{
			{
				"type": "text",
				"text": "A calm cinematic shot of clouds moving over a mountain.",
			},
		},
		"duration":   5,
		"resolution": "720p",
		"ratio":      "16:9",
	}

	jsonBytes, err := sonic.Marshal(testRequest)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}
