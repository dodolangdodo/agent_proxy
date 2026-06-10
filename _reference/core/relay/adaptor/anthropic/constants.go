package anthropic

import (
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/mode"
)

var ModelList = []model.ModelConfig{
	{
		Model: "claude-opus-4-7",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
		Price: model.Price{
			InputPrice:         0.005,
			OutputPrice:        0.025,
			CachedPrice:        0.0005,
			CacheCreationPrice: 0.00625,
		},
		RetryTimes: 5,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(200000),
			model.WithModelConfigMaxOutputTokens(64000),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "claude-opus-4-6",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
		Price: model.Price{
			InputPrice:         0.005,
			OutputPrice:        0.025,
			CachedPrice:        0.0005,
			CacheCreationPrice: 0.00625,
		},
		RetryTimes: 5,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(200000),
			model.WithModelConfigMaxOutputTokens(64000),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "claude-opus-4-5",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
		Price: model.Price{
			InputPrice:         0.005,
			OutputPrice:        0.025,
			CachedPrice:        0.0005,
			CacheCreationPrice: 0.00625,
		},
		RetryTimes: 5,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(200000),
			model.WithModelConfigMaxOutputTokens(64000),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "claude-opus-4-1-20250805",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
		Price: model.Price{
			InputPrice:         0.015,
			OutputPrice:        0.075,
			CachedPrice:        0.0015,
			CacheCreationPrice: 0.01875,
		},
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(200000),
			model.WithModelConfigMaxOutputTokens(32000),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "claude-opus-4-20250514",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
		Price: model.Price{
			InputPrice:         0.015,
			OutputPrice:        0.075,
			CachedPrice:        0.0015,
			CacheCreationPrice: 0.01875,
		},
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(200000),
			model.WithModelConfigMaxOutputTokens(32000),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "claude-sonnet-4-6",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
		Price: model.Price{
			InputPrice:         0.003,
			OutputPrice:        0.015,
			CachedPrice:        0.0003,
			CacheCreationPrice: 0.00375,
		},
		RetryTimes: 5,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(200000),
			model.WithModelConfigMaxOutputTokens(64000),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "claude-sonnet-4-5",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
		Price: model.Price{
			InputPrice:         0.003,
			OutputPrice:        0.015,
			CachedPrice:        0.0003,
			CacheCreationPrice: 0.00375,
		},
		RetryTimes: 5,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(200000),
			model.WithModelConfigMaxOutputTokens(64000),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "claude-sonnet-4-5-20250929",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
		Price: model.Price{
			InputPrice:         0.003,
			OutputPrice:        0.015,
			CachedPrice:        0.0003,
			CacheCreationPrice: 0.00375,
		},
		RetryTimes: 5,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(200000),
			model.WithModelConfigMaxOutputTokens(64000),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "claude-sonnet-4-20250514",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
		Price: model.Price{
			InputPrice:         0.003,
			OutputPrice:        0.015,
			CachedPrice:        0.0003,
			CacheCreationPrice: 0.00375,
		},
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(200000),
			model.WithModelConfigMaxOutputTokens(64000),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "claude-haiku-4-5",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
		Price: model.Price{
			InputPrice:         0.001,
			OutputPrice:        0.005,
			CachedPrice:        0.0001,
			CacheCreationPrice: 0.00125,
		},
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(200000),
			model.WithModelConfigMaxOutputTokens(32000),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "claude-haiku-4-5-20251001",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
		Price: model.Price{
			InputPrice:         0.001,
			OutputPrice:        0.005,
			CachedPrice:        0.0001,
			CacheCreationPrice: 0.00125,
		},
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(200000),
			model.WithModelConfigMaxOutputTokens(32000),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "claude-3-haiku-20240307",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
		Price: model.Price{
			InputPrice:         0.0025,
			OutputPrice:        0.0125,
			CachedPrice:        0.00025,
			CacheCreationPrice: 0.003125,
		},
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(200000),
			model.WithModelConfigMaxOutputTokens(4096),
		),
	},
	{
		Model: "claude-3-opus-20240229",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
		Price: model.Price{
			InputPrice:         0.015,
			OutputPrice:        0.075,
			CachedPrice:        0.0015,
			CacheCreationPrice: 0.01875,
		},
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(200000),
			model.WithModelConfigMaxOutputTokens(4096),
		),
	},
	{
		Model: "claude-3-5-haiku-20241022",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
		Price: model.Price{
			InputPrice:         0.0008,
			OutputPrice:        0.004,
			CachedPrice:        0.00008,
			CacheCreationPrice: 0.001,
		},
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(200000),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "claude-3-5-sonnet-20240620",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
		Price: model.Price{
			InputPrice:         0.003,
			OutputPrice:        0.015,
			CachedPrice:        0.0003,
			CacheCreationPrice: 0.00375,
		},
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(200000),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "claude-3-5-sonnet-20241022",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
		Price: model.Price{
			InputPrice:         0.003,
			OutputPrice:        0.015,
			CachedPrice:        0.0003,
			CacheCreationPrice: 0.00375,
		},
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(200000),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "claude-3-5-sonnet-latest",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
		Price: model.Price{
			InputPrice:         0.003,
			OutputPrice:        0.015,
			CachedPrice:        0.0003,
			CacheCreationPrice: 0.00375,
		},
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(200000),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
}
