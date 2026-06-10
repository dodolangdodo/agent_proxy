package tencent

import (
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/mode"
)

// https://cloud.tencent.com/document/product/1729/104753

var ModelList = []model.ModelConfig{
	{
		Model: "hunyuan-lite",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerTencent,
		Price: model.Price{
			InputPrice:  0.0001,
			OutputPrice: 0.0001,
		},
		RPM: 300,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(262144),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(6144),
		),
	},
	{
		Model: "hunyuan-turbo-latest",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerTencent,
		Price: model.Price{
			InputPrice:  0.015,
			OutputPrice: 0.05,
		},
		RPM: 300,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32768),
			model.WithModelConfigMaxInputTokens(28672),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "hunyuan-turbo",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerTencent,
		Price: model.Price{
			InputPrice:  0.015,
			OutputPrice: 0.05,
		},
		RPM: 300,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32768),
			model.WithModelConfigMaxInputTokens(28672),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigToolChoice(true),
		),
	},

	{
		Model: "hunyuan-pro",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerTencent,
		Price: model.Price{
			InputPrice:  0.03,
			OutputPrice: 0.10,
		},
		RPM: 300,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32768),
			model.WithModelConfigMaxInputTokens(28672),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "hunyuan-large",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerTencent,
		Price: model.Price{
			InputPrice:  0.004,
			OutputPrice: 0.012,
		},
		RPM: 300,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32768),
			model.WithModelConfigMaxInputTokens(28672),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "hunyuan-large-longcontext",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerTencent,
		Price: model.Price{
			InputPrice:  0.006,
			OutputPrice: 0.018,
		},
		RPM: 300,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxOutputTokens(6144),
		),
	},
	{
		Model: "hunyuan-standard",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerTencent,
		Price: model.Price{
			InputPrice:  0.0008,
			OutputPrice: 0.002,
		},
		RPM: 300,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32768),
			model.WithModelConfigMaxOutputTokens(2048),
		),
	},
	// {
	// 	Model:       "hunyuan-standard-256K",
	// 	Type:        mode.ChatCompletions,
	// 	Owner:       model.ModelOwnerTencent,
	// 	InputPrice:  0.0005,
	// 	OutputPrice: 0.002,
	// },
	{
		Model: "hunyuan-role",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerTencent,
		Price: model.Price{
			InputPrice:  0.004,
			OutputPrice: 0.008,
		},
		RPM: 300,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32768),
			model.WithModelConfigMaxOutputTokens(4096),
		),
	},
	{
		Model: "hunyuan-functioncall",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerTencent,
		Price: model.Price{
			InputPrice:  0.004,
			OutputPrice: 0.008,
		},
		RPM: 300,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32768),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "hunyuan-code",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerTencent,
		Price: model.Price{
			InputPrice:  0.004,
			OutputPrice: 0.008,
		},
		RPM: 300,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(8192),
			model.WithModelConfigMaxInputTokens(4096),
			model.WithModelConfigMaxOutputTokens(4096),
		),
	},
	{
		Model: "hunyuan-turbo-vision",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerTencent,
		Price: model.Price{
			InputPrice:  0.08,
			OutputPrice: 0.08,
		},
		RPM: 300,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(8192),
			model.WithModelConfigMaxInputTokens(6144),
			model.WithModelConfigMaxOutputTokens(2048),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "hunyuan-vision",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerTencent,
		Price: model.Price{
			InputPrice:  0.018,
			OutputPrice: 0.018,
		},
		RPM: 300,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(8192),
			model.WithModelConfigMaxInputTokens(6144),
			model.WithModelConfigMaxOutputTokens(2048),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},

	{
		Model: "hunyuan-embedding",
		Type:  mode.Embeddings,
		Owner: model.ModelOwnerTencent,
		Price: model.Price{
			InputPrice: 0.0007,
		},
		RPM: 300,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(1024),
		),
	},
}
