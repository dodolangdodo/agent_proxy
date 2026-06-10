package moonshot

import (
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/mode"
)

var ModelList = []model.ModelConfig{
	{
		Model: "moonshot-v1-8k",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMoonshot,
		Price: model.Price{
			InputPrice:  0.012,
			OutputPrice: 0.012,
		},
		RPM: 500,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "moonshot-v1-32k",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMoonshot,
		Price: model.Price{
			InputPrice:  0.024,
			OutputPrice: 0.024,
		},
		RPM: 500,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(32768),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "moonshot-v1-128k",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMoonshot,
		Price: model.Price{
			InputPrice:  0.06,
			OutputPrice: 0.06,
		},
		RPM: 500,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(131072),
			model.WithModelConfigToolChoice(true),
		),
	},

	{
		Model: "moonshot-v1-8k-vision-preview",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMoonshot,
		Price: model.Price{
			InputPrice:  0.012,
			OutputPrice: 0.012,
		},
		RPM: 500,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(8192),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "moonshot-v1-32k-vision-preview",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMoonshot,
		Price: model.Price{
			InputPrice:  0.024,
			OutputPrice: 0.024,
		},
		RPM: 500,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(32768),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "moonshot-v1-128k-vision-preview",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerMoonshot,
		Price: model.Price{
			InputPrice:  0.06,
			OutputPrice: 0.06,
		},
		RPM: 500,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(131072),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
}
