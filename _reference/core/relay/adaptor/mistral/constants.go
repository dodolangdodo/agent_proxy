package mistral

import (
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/mode"
)

var ModelList = []model.ModelConfig{
	{
		Model: "open-mistral-7b",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMistral,
	},
	{
		Model: "open-mixtral-8x7b",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMistral,
	},
	{
		Model: "mistral-small-latest",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMistral,
	},
	{
		Model: "mistral-medium-latest",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMistral,
	},
	{
		Model: "mistral-large-latest",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMistral,
	},
	{
		Model: "mistral-embed",
		Type:  mode.Embeddings,
		Owner: model.ModelOwnerMistral,
	},
}
