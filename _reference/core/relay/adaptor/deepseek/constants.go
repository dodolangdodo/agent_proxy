package deepseek

import (
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/mode"
)

var ModelList = []model.ModelConfig{
	{
		Model: "deepseek-chat",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerDeepSeek,
		Price: model.Price{
			InputPrice:  0.001,
			OutputPrice: 0.002,
		},
		RPM: 10000,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(64000),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},

	{
		Model: "deepseek-reasoner",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerDeepSeek,
		Price: model.Price{
			InputPrice:  0.004,
			OutputPrice: 0.016,
		},
		RPM: 10000,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(64000),
			model.WithModelConfigMaxOutputTokens(8192),
		),
	},
}
