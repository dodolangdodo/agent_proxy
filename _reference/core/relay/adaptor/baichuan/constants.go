package baichuan

import (
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/mode"
)

var ModelList = []model.ModelConfig{
	{
		Model: "Baichuan4-Turbo",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerBaichuan,
		Price: model.Price{
			InputPrice:  0.015,
			OutputPrice: 0.015,
		},
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32768),
		),
	},
	{
		Model: "Baichuan4-Air",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerBaichuan,
		Price: model.Price{
			InputPrice:  0.00098,
			OutputPrice: 0.00098,
		},
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32768),
		),
	},
	{
		Model: "Baichuan4",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerBaichuan,
		Price: model.Price{
			InputPrice:  0.1,
			OutputPrice: 0.1,
		},
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32768),
		),
	},
	{
		Model: "Baichuan3-Turbo",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerBaichuan,
		Price: model.Price{
			InputPrice:  0.012,
			OutputPrice: 0.012,
		},
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32768),
		),
	},
	{
		Model: "Baichuan3-Turbo-128k",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerBaichuan,
		Price: model.Price{
			InputPrice:  0.024,
			OutputPrice: 0.024,
		},
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
		),
	},

	{
		Model: "Baichuan-Text-Embedding",
		Type:  mode.Embeddings,
		Owner: model.ModelOwnerBaichuan,
		Price: model.Price{
			InputPrice: 0.0005,
		},
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(512),
		),
	},
}
