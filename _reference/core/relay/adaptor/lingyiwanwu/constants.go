package lingyiwanwu

import (
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/mode"
)

// https://platform.lingyiwanwu.com/docs

var ModelList = []model.ModelConfig{
	{
		Model: "yi-lightning",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerLingyiWanwu,
		Price: model.Price{
			InputPrice:  0.00099,
			OutputPrice: 0.00099,
		},
		RPM: 60,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(16384),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "yi-vision-v2",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerLingyiWanwu,
		Price: model.Price{
			InputPrice:  0.006,
			OutputPrice: 0.006,
		},
		RPM: 60,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(16384),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
}
