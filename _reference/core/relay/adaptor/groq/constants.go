package groq

import (
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/mode"
)

// https://console.groq.com/docs/models

var ModelList = []model.ModelConfig{
	{
		Model: "gemma-7b-it",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerGoogle,
	},
	{
		Model: "gemma2-9b-it",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerGoogle,
	},
	{
		Model: "llama-3.1-70b-versatile",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMeta,
	},
	{
		Model: "llama-3.1-8b-instant",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMeta,
	},
	{
		Model: "llama-3.2-11b-text-preview",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMeta,
	},
	{
		Model: "llama-3.2-11b-vision-preview",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMeta,
	},
	{
		Model: "llama-3.2-1b-preview",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMeta,
	},
	{
		Model: "llama-3.2-3b-preview",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMeta,
	},
	{
		Model: "llama-3.2-11b-vision-preview",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMeta,
	},
	{
		Model: "llama-3.2-90b-text-preview",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMeta,
	},
	{
		Model: "llama-3.2-90b-vision-preview",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMeta,
	},
	{
		Model: "llama-guard-3-8b",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMeta,
	},
	{
		Model: "llama3-70b-8192",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMeta,
	},
	{
		Model: "llama3-8b-8192",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMeta,
	},
	{
		Model: "llama3-groq-70b-8192-tool-use-preview",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMeta,
	},
	{
		Model: "llama3-groq-8b-8192-tool-use-preview",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMeta,
	},
	{
		Model: "llava-v1.5-7b-4096-preview",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMeta,
	},
	{
		Model: "mixtral-8x7b-32768",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMistral,
	},
	{
		Model: "distil-whisper-large-v3-en",
		Type:  mode.AudioTranscription,
		Owner: model.ModelOwnerHuggingFace,
	},
	{
		Model: "whisper-large-v3",
		Type:  mode.AudioTranscription,
		Owner: model.ModelOwnerOpenAI,
	},
	{
		Model: "whisper-large-v3-turbo",
		Type:  mode.AudioTranscription,
		Owner: model.ModelOwnerOpenAI,
	},
}
