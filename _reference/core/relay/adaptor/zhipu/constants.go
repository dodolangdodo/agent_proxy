package zhipu

import (
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/mode"
)

var ModelList = []model.ModelConfig{
	{
		Model: "glm-3-turbo",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerChatGLM,
		Price: model.Price{
			InputPrice:  0.001,
			OutputPrice: 0.001,
		},
		RPM: 300,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxOutputTokens(4096),
		),
	},
	{
		Model: "glm-4",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerChatGLM,
		Price: model.Price{
			InputPrice:  0.1,
			OutputPrice: 0.1,
		},
		RPM: 60,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "glm-4-plus",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerChatGLM,
		Price: model.Price{
			InputPrice:  0.05,
			OutputPrice: 0.05,
		},
		RPM: 600,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "glm-4-air",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerChatGLM,
		Price: model.Price{
			InputPrice:  0.001,
			OutputPrice: 0.001,
		},
		RPM: 900,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "glm-4-airx",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerChatGLM,
		Price: model.Price{
			InputPrice:  0.01,
			OutputPrice: 0.01,
		},
		RPM: 60,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(8192),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "glm-4-long",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerChatGLM,
		Price: model.Price{
			InputPrice:  0.001,
			OutputPrice: 0.001,
		},
		RPM: 60,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1024000),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "glm-4-flashx",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerChatGLM,
		Price: model.Price{
			InputPrice:  0.0001,
			OutputPrice: 0.0001,
		},
		RPM: 600,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "glm-4-flash",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerChatGLM,
		Price: model.Price{
			InputPrice:  0.0001,
			OutputPrice: 0.0001,
		},
		RPM: 1800,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "glm-4v-flash",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerChatGLM,
		Price: model.Price{
			InputPrice:  0.0001,
			OutputPrice: 0.0001,
		},
		RPM: 60,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(8192),
			model.WithModelConfigMaxOutputTokens(1024),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "glm-4v",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerChatGLM,
		Price: model.Price{
			InputPrice:  0.05,
			OutputPrice: 0.05,
		},
		RPM: 60,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(2048),
			model.WithModelConfigMaxOutputTokens(1024),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "glm-4v-plus",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerChatGLM,
		Price: model.Price{
			InputPrice:  0.01,
			OutputPrice: 0.01,
		},
		RPM: 60,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(8192),
			model.WithModelConfigMaxOutputTokens(1024),
			model.WithModelConfigVision(true),
		),
	},

	{
		Model: "charglm-4",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerChatGLM,
		Price: model.Price{
			InputPrice:  0.001,
			OutputPrice: 0.001,
		},
		RPM: 60,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(4096),
			model.WithModelConfigMaxOutputTokens(2048),
		),
	},
	{
		Model: "codegeex-4",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerChatGLM,
		Price: model.Price{
			InputPrice:  0.0001,
			OutputPrice: 0.0001,
		},
		RPM: 60,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxOutputTokens(4096),
		),
	},

	{
		Model: "embedding-2",
		Type:  mode.Embeddings,
		Owner: model.ModelOwnerChatGLM,
		Price: model.Price{
			InputPrice: 0.0005,
		},
		RPM: 60,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(8192),
		),
	},
	{
		Model: "embedding-3",
		Type:  mode.Embeddings,
		Owner: model.ModelOwnerChatGLM,
		Price: model.Price{
			InputPrice: 0.0005,
		},
		RPM: 600,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(8192),
		),
	},

	{
		Model: "cogview-3",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerChatGLM,
		Price: model.Price{
			OutputPriceUnit: 1,
			ConditionalPrices: []model.ConditionalPrice{
				imageOutputPrice("1024x1024", 0.1),
			},
		},
		RPM: 60,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxOutputTokens(1024),
		),
	},
	{
		Model: "cogview-3-plus",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerChatGLM,
		Price: model.Price{
			OutputPriceUnit: 1,
			ConditionalPrices: []model.ConditionalPrice{
				imageOutputPrice("1024x1024", 0.06),
				imageOutputPrice("768x1344", 0.06),
				imageOutputPrice("864x1152", 0.06),
				imageOutputPrice("1344x768", 0.06),
				imageOutputPrice("1152x864", 0.06),
				imageOutputPrice("1440x720", 0.06),
				imageOutputPrice("720x1440", 0.06),
			},
		},
		RPM: 60,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxOutputTokens(1024),
		),
	},
}

func imageOutputPrice(size string, price float64) model.ConditionalPrice {
	return model.ConditionalPrice{
		Condition: model.PriceCondition{Resolution: []string{size}},
		Price: model.Price{
			OutputPrice:     model.ZeroNullFloat64(price),
			OutputPriceUnit: 1,
		},
	}
}
