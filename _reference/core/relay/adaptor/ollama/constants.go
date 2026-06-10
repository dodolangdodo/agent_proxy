package ollama

import (
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/mode"
)

var ModelList = []model.ModelConfig{
	{
		Model: "codellama:7b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMeta,
	},
	{
		Model: "llama2:7b",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMeta,
	},
	{
		Model: "llama2:latest",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMeta,
	},
	{
		Model: "llama3:latest",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMeta,
	},
	{
		Model: "phi3:latest",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMicrosoft,
	},
	{
		Model: "qwen:0.5b-chat",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
	},
	{
		Model: "qwen:7b",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
	},
}
