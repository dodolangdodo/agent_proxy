package ali

import (
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/mode"
)

// https://help.aliyun.com/zh/model-studio/getting-started/models?spm=a2c4g.11186623.0.i12#ced16cb6cdfsy
// https://help.aliyun.com/zh/model-studio/model-pricing

var ModelList = []model.ModelConfig{
	{
		Model: "qwen3.7-max",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.012,
			OutputPrice: 0.036,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3.7-max-2026-05-20",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.012,
			OutputPrice: 0.036,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3.7-max-preview",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.012,
			OutputPrice: 0.036,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3.7-2026-05-17",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.012,
			OutputPrice: 0.036,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3.6-max-preview",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.015,
			OutputPrice: 0.09,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 128000},
					Price: model.Price{
						InputPrice:  0.009,
						OutputPrice: 0.054,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 128001, InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:  0.015,
						OutputPrice: 0.09,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-max",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.007,
			OutputPrice: 0.028,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 32000},
					Price: model.Price{
						InputPrice:  0.0025,
						OutputPrice: 0.01,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 32001, InputTokenMax: 128000},
					Price: model.Price{
						InputPrice:  0.004,
						OutputPrice: 0.016,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 128001, InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:  0.007,
						OutputPrice: 0.028,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-max-2026-01-23",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.007,
			OutputPrice: 0.028,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 32000},
					Price: model.Price{
						InputPrice:  0.0025,
						OutputPrice: 0.01,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 32001, InputTokenMax: 128000},
					Price: model.Price{
						InputPrice:  0.004,
						OutputPrice: 0.016,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 128001, InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:  0.007,
						OutputPrice: 0.028,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-max-2025-09-23",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.015,
			OutputPrice: 0.06,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 32000},
					Price: model.Price{
						InputPrice:  0.006,
						OutputPrice: 0.024,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 32001, InputTokenMax: 128000},
					Price: model.Price{
						InputPrice:  0.01,
						OutputPrice: 0.04,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 128001, InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:  0.015,
						OutputPrice: 0.06,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-max-preview",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.015,
			OutputPrice: 0.06,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 32000},
					Price: model.Price{
						InputPrice:  0.006,
						OutputPrice: 0.024,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 32001, InputTokenMax: 128000},
					Price: model.Price{
						InputPrice:  0.01,
						OutputPrice: 0.04,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 128001, InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:  0.015,
						OutputPrice: 0.06,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	// 通义千问-Max
	{
		Model: "qwen-max",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0024,
			OutputPrice: 0.0096,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32768),
			model.WithModelConfigMaxInputTokens(30720),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-max-latest",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0024,
			OutputPrice: 0.0096,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32768),
			model.WithModelConfigMaxInputTokens(30720),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},

	// 通义千问-Plus
	{
		Model: "qwen3.6-plus",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.008,
			OutputPrice: 0.048,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:  0.002,
						OutputPrice: 0.012,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 256001, InputTokenMax: 1000000},
					Price: model.Price{
						InputPrice:  0.008,
						OutputPrice: 0.048,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3.6-plus-2026-04-02",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.008,
			OutputPrice: 0.048,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:  0.002,
						OutputPrice: 0.012,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 256001, InputTokenMax: 1000000},
					Price: model.Price{
						InputPrice:  0.008,
						OutputPrice: 0.048,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3.5-plus",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.004,
			OutputPrice: 0.024,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 128000},
					Price: model.Price{
						InputPrice:  0.0008,
						OutputPrice: 0.0048,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 128001, InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:  0.002,
						OutputPrice: 0.012,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 256001, InputTokenMax: 1000000},
					Price: model.Price{
						InputPrice:  0.004,
						OutputPrice: 0.024,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3.5-plus-2026-04-20",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.004,
			OutputPrice: 0.024,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 128000},
					Price: model.Price{
						InputPrice:  0.0008,
						OutputPrice: 0.0048,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 128001, InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:  0.002,
						OutputPrice: 0.012,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 256001, InputTokenMax: 1000000},
					Price: model.Price{
						InputPrice:  0.004,
						OutputPrice: 0.024,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3.5-plus-2026-02-15",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.004,
			OutputPrice: 0.024,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 128000},
					Price: model.Price{
						InputPrice:  0.0008,
						OutputPrice: 0.0048,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 128001, InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:  0.002,
						OutputPrice: 0.012,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 256001, InputTokenMax: 1000000},
					Price: model.Price{
						InputPrice:  0.004,
						OutputPrice: 0.024,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-plus",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:              0.0048,
			OutputPrice:             0.048,
			ThinkingModeOutputPrice: 0.064,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 128000},
					Price: model.Price{
						InputPrice:              0.0008,
						OutputPrice:             0.002,
						ThinkingModeOutputPrice: 0.008,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 128001, InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:              0.0024,
						OutputPrice:             0.02,
						ThinkingModeOutputPrice: 0.024,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 256001, InputTokenMax: 1000000},
					Price: model.Price{
						InputPrice:              0.0048,
						OutputPrice:             0.048,
						ThinkingModeOutputPrice: 0.064,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-plus-latest",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:              0.0048,
			OutputPrice:             0.048,
			ThinkingModeOutputPrice: 0.064,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 128000},
					Price: model.Price{
						InputPrice:              0.0008,
						OutputPrice:             0.002,
						ThinkingModeOutputPrice: 0.008,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 128001, InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:              0.0024,
						OutputPrice:             0.02,
						ThinkingModeOutputPrice: 0.024,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 256001, InputTokenMax: 1000000},
					Price: model.Price{
						InputPrice:              0.0048,
						OutputPrice:             0.048,
						ThinkingModeOutputPrice: 0.064,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32000),
			model.WithModelConfigMaxInputTokens(30000),
			model.WithModelConfigMaxOutputTokens(8000),
			model.WithModelConfigToolChoice(true),
		),
	},

	// 通义千问-Turbo
	{
		Model: "qwen3.6-flash",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0048,
			OutputPrice: 0.0288,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:  0.0012,
						OutputPrice: 0.0072,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 256001, InputTokenMax: 1000000},
					Price: model.Price{
						InputPrice:  0.0048,
						OutputPrice: 0.0288,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3.6-flash-2026-04-16",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0048,
			OutputPrice: 0.0288,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:  0.0012,
						OutputPrice: 0.0072,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 256001, InputTokenMax: 1000000},
					Price: model.Price{
						InputPrice:  0.0048,
						OutputPrice: 0.0288,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3.5-flash",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0012,
			OutputPrice: 0.012,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 128000},
					Price: model.Price{
						InputPrice:  0.0002,
						OutputPrice: 0.002,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 128001, InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:  0.0008,
						OutputPrice: 0.008,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 256001, InputTokenMax: 1000000},
					Price: model.Price{
						InputPrice:  0.0012,
						OutputPrice: 0.012,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3.5-flash-2026-02-23",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0012,
			OutputPrice: 0.012,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 128000},
					Price: model.Price{
						InputPrice:  0.0002,
						OutputPrice: 0.002,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 128001, InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:  0.0008,
						OutputPrice: 0.008,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 256001, InputTokenMax: 1000000},
					Price: model.Price{
						InputPrice:  0.0012,
						OutputPrice: 0.012,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-flash",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0012,
			OutputPrice: 0.012,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 128000},
					Price: model.Price{
						InputPrice:  0.00015,
						OutputPrice: 0.0015,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 128001, InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:  0.0006,
						OutputPrice: 0.006,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 256001, InputTokenMax: 1000000},
					Price: model.Price{
						InputPrice:  0.0012,
						OutputPrice: 0.012,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-flash-2025-07-28",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0012,
			OutputPrice: 0.012,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 128000},
					Price: model.Price{
						InputPrice:  0.00015,
						OutputPrice: 0.0015,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 128001, InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:  0.0006,
						OutputPrice: 0.006,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 256001, InputTokenMax: 1000000},
					Price: model.Price{
						InputPrice:  0.0012,
						OutputPrice: 0.012,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-turbo",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:              0.0003,
			OutputPrice:             0.0006,
			ThinkingModeOutputPrice: 0.003,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-turbo-latest",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:              0.0003,
			OutputPrice:             0.0006,
			ThinkingModeOutputPrice: 0.003,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-turbo-2025-07-15",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:              0.0003,
			OutputPrice:             0.0006,
			ThinkingModeOutputPrice: 0.003,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-turbo-2025-04-28",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:              0.0003,
			OutputPrice:             0.0006,
			ThinkingModeOutputPrice: 0.003,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},

	// Qwen-Long
	{
		Model: "qwen-long",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0005,
			OutputPrice: 0.002,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(6000),
			model.WithModelConfigToolChoice(true),
		),
	},

	// 通义千问Omni
	{
		Model: "qwen3.5-omni-plus",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:       0.007,
			ImageInputPrice:  0.007,
			AudioInputPrice:  0.053,
			VideoInputPrice:  0.007,
			OutputPrice:      0.04,
			AudioOutputPrice: 0.213,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{OutputAudio: new(true)},
					Price: model.Price{
						InputPrice:       0.007,
						ImageInputPrice:  0.007,
						AudioInputPrice:  0.053,
						VideoInputPrice:  0.007,
						OutputPrice:      0,
						AudioOutputPrice: 0.213,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3.5-omni-plus-2026-03-15",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:       0.007,
			ImageInputPrice:  0.007,
			AudioInputPrice:  0.053,
			VideoInputPrice:  0.007,
			OutputPrice:      0.04,
			AudioOutputPrice: 0.213,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{OutputAudio: new(true)},
					Price: model.Price{
						InputPrice:       0.007,
						ImageInputPrice:  0.007,
						AudioInputPrice:  0.053,
						VideoInputPrice:  0.007,
						OutputPrice:      0,
						AudioOutputPrice: 0.213,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3.5-omni-flash",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:       0.0022,
			ImageInputPrice:  0.0022,
			AudioInputPrice:  0.018,
			VideoInputPrice:  0.0022,
			OutputPrice:      0.0133,
			AudioOutputPrice: 0.072,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{OutputAudio: new(true)},
					Price: model.Price{
						InputPrice:       0.0022,
						ImageInputPrice:  0.0022,
						AudioInputPrice:  0.018,
						VideoInputPrice:  0.0022,
						OutputPrice:      0,
						AudioOutputPrice: 0.072,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3.5-omni-flash-2026-03-15",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:       0.0022,
			ImageInputPrice:  0.0022,
			AudioInputPrice:  0.018,
			VideoInputPrice:  0.0022,
			OutputPrice:      0.0133,
			AudioOutputPrice: 0.072,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{OutputAudio: new(true)},
					Price: model.Price{
						InputPrice:       0.0022,
						ImageInputPrice:  0.0022,
						AudioInputPrice:  0.018,
						VideoInputPrice:  0.0022,
						OutputPrice:      0,
						AudioOutputPrice: 0.072,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-omni-flash",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:       0.0018,
			ImageInputPrice:  0.0033,
			AudioInputPrice:  0.0158,
			VideoInputPrice:  0.0033,
			OutputPrice:      0.0127,
			AudioOutputPrice: 0.0626,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{
						InputMedia:  new(false),
						OutputAudio: new(false),
					},
					Price: model.Price{
						InputPrice:       0.0018,
						ImageInputPrice:  0.0033,
						AudioInputPrice:  0.0158,
						VideoInputPrice:  0.0033,
						OutputPrice:      0.0069,
						AudioOutputPrice: 0.0626,
					},
				},
				{
					Condition: model.PriceCondition{OutputAudio: new(true)},
					Price: model.Price{
						InputPrice:       0.0018,
						ImageInputPrice:  0.0033,
						AudioInputPrice:  0.0158,
						VideoInputPrice:  0.0033,
						OutputPrice:      0,
						AudioOutputPrice: 0.0626,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-omni-flash-2025-12-01",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:       0.0018,
			ImageInputPrice:  0.0033,
			AudioInputPrice:  0.0158,
			VideoInputPrice:  0.0033,
			OutputPrice:      0.0127,
			AudioOutputPrice: 0.0626,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{
						InputMedia:  new(false),
						OutputAudio: new(false),
					},
					Price: model.Price{
						InputPrice:       0.0018,
						ImageInputPrice:  0.0033,
						AudioInputPrice:  0.0158,
						VideoInputPrice:  0.0033,
						OutputPrice:      0.0069,
						AudioOutputPrice: 0.0626,
					},
				},
				{
					Condition: model.PriceCondition{OutputAudio: new(true)},
					Price: model.Price{
						InputPrice:       0.0018,
						ImageInputPrice:  0.0033,
						AudioInputPrice:  0.0158,
						VideoInputPrice:  0.0033,
						OutputPrice:      0,
						AudioOutputPrice: 0.0626,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-omni-flash-2025-09-15",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:       0.0018,
			ImageInputPrice:  0.0033,
			AudioInputPrice:  0.0158,
			VideoInputPrice:  0.0033,
			OutputPrice:      0.0127,
			AudioOutputPrice: 0.0626,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{
						InputMedia:  new(false),
						OutputAudio: new(false),
					},
					Price: model.Price{
						InputPrice:       0.0018,
						ImageInputPrice:  0.0033,
						AudioInputPrice:  0.0158,
						VideoInputPrice:  0.0033,
						OutputPrice:      0.0069,
						AudioOutputPrice: 0.0626,
					},
				},
				{
					Condition: model.PriceCondition{OutputAudio: new(true)},
					Price: model.Price{
						InputPrice:       0.0018,
						ImageInputPrice:  0.0033,
						AudioInputPrice:  0.0158,
						VideoInputPrice:  0.0033,
						OutputPrice:      0,
						AudioOutputPrice: 0.0626,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-omni-turbo",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:       0.0004,
			ImageInputPrice:  0.0015,
			AudioInputPrice:  0.025,
			VideoInputPrice:  0.0015,
			OutputPrice:      0.0045,
			AudioOutputPrice: 0.05,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{
						InputMedia:  new(false),
						OutputAudio: new(false),
					},
					Price: model.Price{
						InputPrice:       0.0004,
						ImageInputPrice:  0.0015,
						AudioInputPrice:  0.025,
						VideoInputPrice:  0.0015,
						OutputPrice:      0.0016,
						AudioOutputPrice: 0.05,
					},
				},
				{
					Condition: model.PriceCondition{OutputAudio: new(true)},
					Price: model.Price{
						InputPrice:       0.0004,
						ImageInputPrice:  0.0015,
						AudioInputPrice:  0.025,
						VideoInputPrice:  0.0015,
						OutputPrice:      0,
						AudioOutputPrice: 0.05,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-omni-turbo-latest",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:       0.0004,
			ImageInputPrice:  0.0015,
			AudioInputPrice:  0.025,
			VideoInputPrice:  0.0015,
			OutputPrice:      0.0045,
			AudioOutputPrice: 0.05,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{
						InputMedia:  new(false),
						OutputAudio: new(false),
					},
					Price: model.Price{
						InputPrice:       0.0004,
						ImageInputPrice:  0.0015,
						AudioInputPrice:  0.025,
						VideoInputPrice:  0.0015,
						OutputPrice:      0.0016,
						AudioOutputPrice: 0.05,
					},
				},
				{
					Condition: model.PriceCondition{OutputAudio: new(true)},
					Price: model.Price{
						InputPrice:       0.0004,
						ImageInputPrice:  0.0015,
						AudioInputPrice:  0.025,
						VideoInputPrice:  0.0015,
						OutputPrice:      0,
						AudioOutputPrice: 0.05,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-omni-turbo-2025-03-26",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:       0.0004,
			ImageInputPrice:  0.0015,
			AudioInputPrice:  0.025,
			VideoInputPrice:  0.0015,
			OutputPrice:      0.0045,
			AudioOutputPrice: 0.05,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{
						InputMedia:  new(false),
						OutputAudio: new(false),
					},
					Price: model.Price{
						InputPrice:       0.0004,
						ImageInputPrice:  0.0015,
						AudioInputPrice:  0.025,
						VideoInputPrice:  0.0015,
						OutputPrice:      0.0016,
						AudioOutputPrice: 0.05,
					},
				},
				{
					Condition: model.PriceCondition{OutputAudio: new(true)},
					Price: model.Price{
						InputPrice:       0.0004,
						ImageInputPrice:  0.0015,
						AudioInputPrice:  0.025,
						VideoInputPrice:  0.0015,
						OutputPrice:      0,
						AudioOutputPrice: 0.05,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-omni-turbo-2025-01-19",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:       0.0004,
			ImageInputPrice:  0.0015,
			AudioInputPrice:  0.025,
			VideoInputPrice:  0.0015,
			OutputPrice:      0.0045,
			AudioOutputPrice: 0.05,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{
						InputMedia:  new(false),
						OutputAudio: new(false),
					},
					Price: model.Price{
						InputPrice:       0.0004,
						ImageInputPrice:  0.0015,
						AudioInputPrice:  0.025,
						VideoInputPrice:  0.0015,
						OutputPrice:      0.0016,
						AudioOutputPrice: 0.05,
					},
				},
				{
					Condition: model.PriceCondition{OutputAudio: new(true)},
					Price: model.Price{
						InputPrice:       0.0004,
						ImageInputPrice:  0.0015,
						AudioInputPrice:  0.025,
						VideoInputPrice:  0.0015,
						OutputPrice:      0,
						AudioOutputPrice: 0.05,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen2.5-omni-7b",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:       0.0006,
			ImageInputPrice:  0.002,
			AudioInputPrice:  0.038,
			VideoInputPrice:  0.002,
			OutputPrice:      0.006,
			AudioOutputPrice: 0.076,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{
						InputMedia:  new(false),
						OutputAudio: new(false),
					},
					Price: model.Price{
						InputPrice:       0.0006,
						ImageInputPrice:  0.002,
						AudioInputPrice:  0.038,
						VideoInputPrice:  0.002,
						OutputPrice:      0.0024,
						AudioOutputPrice: 0.076,
					},
				},
				{
					Condition: model.PriceCondition{OutputAudio: new(true)},
					Price: model.Price{
						InputPrice:       0.0006,
						ImageInputPrice:  0.002,
						AudioInputPrice:  0.038,
						VideoInputPrice:  0.002,
						OutputPrice:      0,
						AudioOutputPrice: 0.076,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-omni-30b-a3b-captioner",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0158,
			OutputPrice: 0.0127,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},

	// 通义千问VL
	{
		Model: "qwen3-vl-plus",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.003,
			OutputPrice: 0.03,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 32000},
					Price: model.Price{
						InputPrice:  0.001,
						OutputPrice: 0.01,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 32001, InputTokenMax: 128000},
					Price: model.Price{
						InputPrice:  0.0015,
						OutputPrice: 0.015,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 128001, InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:  0.003,
						OutputPrice: 0.03,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-vl-plus-2025-12-19",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.003,
			OutputPrice: 0.03,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 32000},
					Price: model.Price{
						InputPrice:  0.001,
						OutputPrice: 0.01,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 32001, InputTokenMax: 128000},
					Price: model.Price{
						InputPrice:  0.0015,
						OutputPrice: 0.015,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 128001, InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:  0.003,
						OutputPrice: 0.03,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-vl-plus-2025-09-23",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.003,
			OutputPrice: 0.03,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 32000},
					Price: model.Price{
						InputPrice:  0.001,
						OutputPrice: 0.01,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 32001, InputTokenMax: 128000},
					Price: model.Price{
						InputPrice:  0.0015,
						OutputPrice: 0.015,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 128001, InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:  0.003,
						OutputPrice: 0.03,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-vl-flash",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0006,
			OutputPrice: 0.006,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 32000},
					Price: model.Price{
						InputPrice:  0.00015,
						OutputPrice: 0.0015,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 32001, InputTokenMax: 128000},
					Price: model.Price{
						InputPrice:  0.0003,
						OutputPrice: 0.003,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 128001, InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:  0.0006,
						OutputPrice: 0.006,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-vl-flash-2026-01-22",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0006,
			OutputPrice: 0.006,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 32000},
					Price: model.Price{
						InputPrice:  0.00015,
						OutputPrice: 0.0015,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 32001, InputTokenMax: 128000},
					Price: model.Price{
						InputPrice:  0.0003,
						OutputPrice: 0.003,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 128001, InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:  0.0006,
						OutputPrice: 0.006,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-vl-flash-2025-10-15",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0006,
			OutputPrice: 0.006,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 32000},
					Price: model.Price{
						InputPrice:  0.00015,
						OutputPrice: 0.0015,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 32001, InputTokenMax: 128000},
					Price: model.Price{
						InputPrice:  0.0003,
						OutputPrice: 0.003,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 128001, InputTokenMax: 256000},
					Price: model.Price{
						InputPrice:  0.0006,
						OutputPrice: 0.006,
					},
				},
			},
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-vl-235b-a22b-thinking",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.002,
			OutputPrice: 0.02,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-vl-235b-a22b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.002,
			OutputPrice: 0.008,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-vl-32b-thinking",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.002,
			OutputPrice: 0.02,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-vl-32b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.002,
			OutputPrice: 0.008,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-vl-30b-a3b-thinking",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.00075,
			OutputPrice: 0.0075,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-vl-30b-a3b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.00075,
			OutputPrice: 0.003,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-vl-8b-thinking",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0005,
			OutputPrice: 0.005,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-vl-8b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0005,
			OutputPrice: 0.002,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-next-80b-a3b-thinking",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.001,
			OutputPrice: 0.01,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-next-80b-a3b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.001,
			OutputPrice: 0.004,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-235b-a22b-thinking-2507",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.002,
			OutputPrice: 0.02,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-235b-a22b-instruct-2507",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.002,
			OutputPrice: 0.008,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-30b-a3b-thinking-2507",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.00075,
			OutputPrice: 0.0075,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-30b-a3b-instruct-2507",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.00075,
			OutputPrice: 0.003,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-235b-a22b",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:              0.002,
			OutputPrice:             0.008,
			ThinkingModeOutputPrice: 0.02,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-32b",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:              0.002,
			OutputPrice:             0.008,
			ThinkingModeOutputPrice: 0.02,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-30b-a3b",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:              0.00075,
			OutputPrice:             0.003,
			ThinkingModeOutputPrice: 0.0075,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-14b",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:              0.001,
			OutputPrice:             0.004,
			ThinkingModeOutputPrice: 0.01,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-8b",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:              0.0005,
			OutputPrice:             0.002,
			ThinkingModeOutputPrice: 0.005,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-4b",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:              0.0003,
			OutputPrice:             0.0012,
			ThinkingModeOutputPrice: 0.003,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-1.7b",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:              0.0003,
			OutputPrice:             0.0012,
			ThinkingModeOutputPrice: 0.003,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen3-0.6b",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:              0.0003,
			OutputPrice:             0.0012,
			ThinkingModeOutputPrice: 0.003,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxInputTokens(256000),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-vl-max",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0016,
			OutputPrice: 0.004,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32000),
			model.WithModelConfigMaxInputTokens(30000),
			model.WithModelConfigMaxOutputTokens(2000),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-vl-max-latest",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0016,
			OutputPrice: 0.004,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32000),
			model.WithModelConfigMaxInputTokens(30000),
			model.WithModelConfigMaxOutputTokens(2000),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-vl-plus",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0008,
			OutputPrice: 0.002,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(8000),
			model.WithModelConfigMaxInputTokens(6000),
			model.WithModelConfigMaxOutputTokens(2000),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-vl-plus-latest",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0008,
			OutputPrice: 0.002,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32000),
			model.WithModelConfigMaxInputTokens(30000),
			model.WithModelConfigMaxOutputTokens(2000),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},

	// 通义千问OCR
	{
		Model: "qwen-vl-ocr",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0003,
			OutputPrice: 0.0005,
		},
		RPM:              600,
		ExcludeFromTests: true,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(34096),
			model.WithModelConfigMaxInputTokens(30000),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "qwen-vl-ocr-latest",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0003,
			OutputPrice: 0.0005,
		},
		RPM:              600,
		ExcludeFromTests: true,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(34096),
			model.WithModelConfigMaxInputTokens(30000),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigVision(true),
		),
	},

	// 通义千问Math
	{
		Model: "qwen-math-plus",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.004,
			OutputPrice: 0.012,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(4096),
			model.WithModelConfigMaxInputTokens(3072),
			model.WithModelConfigMaxOutputTokens(3072),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-math-plus-latest",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.004,
			OutputPrice: 0.012,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(4096),
			model.WithModelConfigMaxInputTokens(3072),
			model.WithModelConfigMaxOutputTokens(3072),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-math-turbo",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.002,
			OutputPrice: 0.006,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(4096),
			model.WithModelConfigMaxInputTokens(3072),
			model.WithModelConfigMaxOutputTokens(3072),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-math-turbo-latest",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.002,
			OutputPrice: 0.006,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(4096),
			model.WithModelConfigMaxInputTokens(3072),
			model.WithModelConfigMaxOutputTokens(3072),
			model.WithModelConfigToolChoice(true),
		),
	},

	// 通义千问Coder
	{
		Model: "qwen-coder-plus",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0035,
			OutputPrice: 0.007,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-coder-plus-latest",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0035,
			OutputPrice: 0.007,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-coder-turbo",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.002,
			OutputPrice: 0.006,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-coder-turbo-latest",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.002,
			OutputPrice: 0.006,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},

	// 通义千问2.5
	{
		Model: "qwen2.5-72b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.004,
			OutputPrice: 0.012,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen2.5-32b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0035,
			OutputPrice: 0.007,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen2.5-14b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.002,
			OutputPrice: 0.006,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen2.5-7b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.001,
			OutputPrice: 0.002,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen2.5-vl-72b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.016,
			OutputPrice: 0.048,
		},
		RPM: 60,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen2.5-vl-32b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.008,
			OutputPrice: 0.024,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen2.5-vl-7b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.002,
			OutputPrice: 0.005,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen2.5-vl-3b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0012,
			OutputPrice: 0.0036,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},

	// 通义千问2
	{
		Model: "qwen2-72b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.004,
			OutputPrice: 0.012,
		},
		RPM: 60,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(128000),
			model.WithModelConfigMaxOutputTokens(6144),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen2-57b-a14b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0035,
			OutputPrice: 0.007,
		},
		RPM: 60,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(65536),
			model.WithModelConfigMaxInputTokens(63488),
			model.WithModelConfigMaxOutputTokens(6144),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen2-7b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.001,
			OutputPrice: 0.002,
		},
		RPM: 60,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(128000),
			model.WithModelConfigMaxOutputTokens(6144),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen2-vl-72b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.016,
			OutputPrice: 0.048,
		},
		RPM: 60,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32768),
			model.WithModelConfigMaxInputTokens(30720),
			model.WithModelConfigMaxOutputTokens(2048),
			model.WithModelConfigVision(true),
		),
	},

	// 通义千问1.5
	{
		Model: "qwen1.5-110b-chat",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.007,
			OutputPrice: 0.014,
		},
		RPM: 10,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32000),
			model.WithModelConfigMaxInputTokens(30000),
			model.WithModelConfigMaxOutputTokens(8000),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen1.5-72b-chat",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.005,
			OutputPrice: 0.01,
		},
		RPM: 120,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32000),
			model.WithModelConfigMaxInputTokens(30000),
			model.WithModelConfigMaxOutputTokens(8000),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen1.5-32b-chat",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0035,
			OutputPrice: 0.007,
		},
		RPM: 10,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32000),
			model.WithModelConfigMaxInputTokens(30000),
			model.WithModelConfigMaxOutputTokens(8000),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen1.5-14b-chat",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.002,
			OutputPrice: 0.004,
		},
		RPM: 120,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(8000),
			model.WithModelConfigMaxInputTokens(6000),
			model.WithModelConfigMaxOutputTokens(2000),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen1.5-7b-chat",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.001,
			OutputPrice: 0.002,
		},
		RPM: 120,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(8000),
			model.WithModelConfigMaxInputTokens(6000),
			model.WithModelConfigMaxOutputTokens(2000),
			model.WithModelConfigToolChoice(true),
		),
	},

	// 通义千问
	{
		Model: "qwen-72b-chat",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.02,
			OutputPrice: 0.02,
		},
		RPM: 80,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32000),
			model.WithModelConfigMaxInputTokens(30000),
			model.WithModelConfigMaxOutputTokens(2000),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-14b-chat",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.008,
			OutputPrice: 0.008,
		},
		RPM: 300,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(8000),
			model.WithModelConfigMaxInputTokens(6000),
			model.WithModelConfigMaxOutputTokens(2000),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen-7b-chat",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.006,
			OutputPrice: 0.006,
		},
		RPM: 300,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(7500),
			model.WithModelConfigMaxInputTokens(6000),
			model.WithModelConfigMaxOutputTokens(1500),
			model.WithModelConfigToolChoice(true),
		),
	},

	// 通义千问数学模型
	{
		Model: "qwen2.5-math-72b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.004,
			OutputPrice: 0.012,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(4096),
			model.WithModelConfigMaxInputTokens(3072),
			model.WithModelConfigMaxOutputTokens(3072),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen2.5-math-7b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.001,
			OutputPrice: 0.002,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(4096),
			model.WithModelConfigMaxInputTokens(3072),
			model.WithModelConfigMaxOutputTokens(3072),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen2-math-72b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.004,
			OutputPrice: 0.012,
		},
		RPM: 10,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(4096),
			model.WithModelConfigMaxInputTokens(3072),
			model.WithModelConfigMaxOutputTokens(3072),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen2-math-7b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.001,
			OutputPrice: 0.002,
		},
		RPM: 10,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(4096),
			model.WithModelConfigMaxInputTokens(3072),
			model.WithModelConfigMaxOutputTokens(3072),
			model.WithModelConfigToolChoice(true),
		),
	},

	// 通义千问Coder
	{
		Model: "qwen2.5-coder-32b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0035,
			OutputPrice: 0.007,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen2.5-coder-14b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.002,
			OutputPrice: 0.006,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "qwen2.5-coder-7b-instruct",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.001,
			OutputPrice: 0.002,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},

	{
		Model: "qwq-32b-preview",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.0035,
			OutputPrice: 0.007,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32768),
			model.WithModelConfigMaxInputTokens(30720),
			model.WithModelConfigMaxOutputTokens(16384),
		),
	},
	{
		Model: "qvq-72b-preview",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.012,
			OutputPrice: 0.036,
		},
		RPM: 60,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32768),
			model.WithModelConfigMaxInputTokens(16384),
			model.WithModelConfigMaxOutputTokens(16384),
		),
	},

	{
		Model: "qwen-mt-plus",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.015,
			OutputPrice: 0.045,
		},
		RPM:              60,
		ExcludeFromTests: true,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(2048),
			model.WithModelConfigMaxInputTokens(1024),
			model.WithModelConfigMaxOutputTokens(1024),
		),
	},
	{
		Model: "qwen-mt-turbo",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:  0.001,
			OutputPrice: 0.003,
		},
		RPM:              60,
		ExcludeFromTests: true,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(2048),
			model.WithModelConfigMaxInputTokens(1024),
			model.WithModelConfigMaxOutputTokens(1024),
		),
	},

	// stable-diffusion
	{
		Model: "stable-diffusion-xl",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerStabilityAI,
		RPM:   2,
	},
	{
		Model: "stable-diffusion-v1.5",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerStabilityAI,
		RPM:   2,
	},
	{
		Model: "stable-diffusion-3.5-large",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerStabilityAI,
		RPM:   2,
	},
	{
		Model: "stable-diffusion-3.5-large-turbo",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerStabilityAI,
		RPM:   2,
	},
	{
		Model: "qwen-image",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		// Ali image generation/edit APIs bill by successful output image count.
		// aliImageUsageToOpenAI maps that count into ImageOutputTokens.
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.25),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "qwen-image-plus",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.2),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "qwen-image-plus-2026-01-09",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.2),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "qwen-image-max",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.5),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "qwen-image-max-2025-12-30",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.5),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "qwen-image-2.0",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.2),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "qwen-image-2.0-2026-03-03",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.2),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "qwen-image-2.0-pro",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.5),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "qwen-image-2.0-pro-2026-03-03",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.5),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "qwen-image-2.0-pro-2026-04-22",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.5),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "qwen-image-edit",
		Type:  mode.ImagesEdits,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.3),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "qwen-image-edit-plus",
		Type:  mode.ImagesEdits,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.2),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "qwen-image-edit-plus-2025-10-30",
		Type:  mode.ImagesEdits,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.2),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "qwen-image-edit-plus-2025-12-15",
		Type:  mode.ImagesEdits,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.2),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "qwen-image-edit-max",
		Type:  mode.ImagesEdits,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.5),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "qwen-image-edit-max-2026-01-16",
		Type:  mode.ImagesEdits,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.5),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "qwen-mt-image",
		Type:  mode.ImagesEdits,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.003),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "z-image-turbo",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.2),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "wan2.7-image",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.2),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "wan2.7-image-pro",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.5),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "wan2.6-t2i",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.2),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "wan2.5-t2i-preview",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.2),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "wan2.2-t2i-flash",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.14),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "wan2.2-t2i-plus",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.2),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "wanx2.1-t2i-turbo",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.14),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "wanx2.1-t2i-plus",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.2),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "wanx2.0-t2i-turbo",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.04),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "wanx-v1",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.16),
			ImageOutputPriceUnit: 1,
		},
		RPM: 60,
	},
	{
		Model: "wan2.5-t2v-preview",
		Type:  mode.AliVideo,
		Owner: model.ModelOwnerAlibaba,
		// Ali video APIs bill by successful output video seconds. Async usage
		// normalizes returned dimensions into a resolution condition.
		Price: model.Price{
			OutputPrice:     model.ZeroNullFloat64(1.0),
			OutputPriceUnit: 1,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{Resolution: []string{"480p"}},
					Price: model.Price{
						OutputPrice:     model.ZeroNullFloat64(0.3),
						OutputPriceUnit: 1,
					},
				},
				{
					Condition: model.PriceCondition{Resolution: []string{"720p"}},
					Price: model.Price{
						OutputPrice:     model.ZeroNullFloat64(0.6),
						OutputPriceUnit: 1,
					},
				},
				{
					Condition: model.PriceCondition{Resolution: []string{"1080p"}},
					Price: model.Price{
						OutputPrice:     model.ZeroNullFloat64(1.0),
						OutputPriceUnit: 1,
					},
				},
			},
		},
		RPM: 60,
	},
	{
		Model: "wan2.5-i2v-preview",
		Type:  mode.AliVideo,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			OutputPrice:     model.ZeroNullFloat64(1.0),
			OutputPriceUnit: 1,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{Resolution: []string{"480p"}},
					Price: model.Price{
						OutputPrice:     model.ZeroNullFloat64(0.3),
						OutputPriceUnit: 1,
					},
				},
				{
					Condition: model.PriceCondition{Resolution: []string{"720p"}},
					Price: model.Price{
						OutputPrice:     model.ZeroNullFloat64(0.6),
						OutputPriceUnit: 1,
					},
				},
				{
					Condition: model.PriceCondition{Resolution: []string{"1080p"}},
					Price: model.Price{
						OutputPrice:     model.ZeroNullFloat64(1.0),
						OutputPriceUnit: 1,
					},
				},
			},
		},
		RPM: 60,
	},
	{
		Model: "wan2.2-t2v-plus",
		Type:  mode.AliVideo,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			OutputPrice:     model.ZeroNullFloat64(0.70),
			OutputPriceUnit: 1,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{Resolution: []string{"480p"}},
					Price: model.Price{
						OutputPrice:     model.ZeroNullFloat64(0.14),
						OutputPriceUnit: 1,
					},
				},
				{
					Condition: model.PriceCondition{Resolution: []string{"1080p"}},
					Price: model.Price{
						OutputPrice:     model.ZeroNullFloat64(0.70),
						OutputPriceUnit: 1,
					},
				},
			},
		},
		RPM: 60,
	},
	{
		Model: "wan2.2-i2v-plus",
		Type:  mode.AliVideo,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			OutputPrice:     model.ZeroNullFloat64(0.70),
			OutputPriceUnit: 1,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{Resolution: []string{"480p"}},
					Price: model.Price{
						OutputPrice:     model.ZeroNullFloat64(0.14),
						OutputPriceUnit: 1,
					},
				},
				{
					Condition: model.PriceCondition{Resolution: []string{"1080p"}},
					Price: model.Price{
						OutputPrice:     model.ZeroNullFloat64(0.70),
						OutputPriceUnit: 1,
					},
				},
			},
		},
		RPM: 60,
	},
	{
		Model: "wan2.2-videoedit",
		Type:  mode.AliVideo,
		Owner: model.ModelOwnerAlibaba,
		RPM:   60,
	},
	{
		Model: "kling-v1",
		Type:  mode.AliVideo,
		Owner: model.ModelOwnerAlibaba,
		RPM:   60,
	},
	{
		Model: "vidu2.0",
		Type:  mode.AliVideo,
		Owner: model.ModelOwnerAlibaba,
		RPM:   60,
	},
	{
		Model: "pixverse-v4.5",
		Type:  mode.AliVideo,
		Owner: model.ModelOwnerAlibaba,
		RPM:   60,
	},
	{
		Model: "happyhorse-1.0-t2v",
		Type:  mode.AliVideo,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			OutputPrice:     model.ZeroNullFloat64(1.6),
			OutputPriceUnit: 1,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{Resolution: []string{"720p"}},
					Price: model.Price{
						OutputPrice:     model.ZeroNullFloat64(0.9),
						OutputPriceUnit: 1,
					},
				},
				{
					Condition: model.PriceCondition{Resolution: []string{"1080p"}},
					Price: model.Price{
						OutputPrice:     model.ZeroNullFloat64(1.6),
						OutputPriceUnit: 1,
					},
				},
			},
		},
		RPM: 60,
	},
	{
		Model: "happyhorse-1.0-i2v",
		Type:  mode.AliVideo,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			OutputPrice:     model.ZeroNullFloat64(1.6),
			OutputPriceUnit: 1,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{Resolution: []string{"720p"}},
					Price: model.Price{
						OutputPrice:     model.ZeroNullFloat64(0.9),
						OutputPriceUnit: 1,
					},
				},
				{
					Condition: model.PriceCondition{Resolution: []string{"1080p"}},
					Price: model.Price{
						OutputPrice:     model.ZeroNullFloat64(1.6),
						OutputPriceUnit: 1,
					},
				},
			},
		},
		RPM: 60,
	},
	{
		Model: "happyhorse-1.0-r2v",
		Type:  mode.AliVideo,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			OutputPrice:     model.ZeroNullFloat64(1.6),
			OutputPriceUnit: 1,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{Resolution: []string{"720p"}},
					Price: model.Price{
						OutputPrice:     model.ZeroNullFloat64(0.9),
						OutputPriceUnit: 1,
					},
				},
				{
					Condition: model.PriceCondition{Resolution: []string{"1080p"}},
					Price: model.Price{
						OutputPrice:     model.ZeroNullFloat64(1.6),
						OutputPriceUnit: 1,
					},
				},
			},
		},
		RPM: 60,
	},
	{
		Model: "happyhorse-1.0-video-edit",
		Type:  mode.AliVideo,
		Owner: model.ModelOwnerAlibaba,
		// HappyHorse video edit bills both input video seconds and output video
		// seconds. fetchAliVideoJobUsage maps those seconds into VideoInputTokens
		// and OutputTokens.
		Price: model.Price{
			VideoInputPrice:     model.ZeroNullFloat64(1.6),
			VideoInputPriceUnit: 1,
			OutputPrice:         model.ZeroNullFloat64(1.6),
			OutputPriceUnit:     1,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{Resolution: []string{"720p"}},
					Price: model.Price{
						VideoInputPrice:     model.ZeroNullFloat64(0.9),
						VideoInputPriceUnit: 1,
						OutputPrice:         model.ZeroNullFloat64(0.9),
						OutputPriceUnit:     1,
					},
				},
				{
					Condition: model.PriceCondition{Resolution: []string{"1080p"}},
					Price: model.Price{
						VideoInputPrice:     model.ZeroNullFloat64(1.6),
						VideoInputPriceUnit: 1,
						OutputPrice:         model.ZeroNullFloat64(1.6),
						OutputPriceUnit:     1,
					},
				},
			},
		},
		RPM: 60,
	},

	{
		Model: "sambert-v1",
		Type:  mode.AudioSpeech,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice: 0.1,
		},
		RPM: 20,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(10000),
			model.WithModelConfigSupportFormats([]string{"mp3", "wav", "pcm"}),
			model.WithModelConfigSupportVoices([]string{
				"zhinan",
				"zhiqi",
				"zhichu",
				"zhide",
				"zhijia",
				"zhiru",
				"zhiqian",
				"zhixiang",
				"zhiwei",
				"zhihao",
				"zhijing",
				"zhiming",
				"zhimo",
				"zhina",
				"zhishu",
				"zhistella",
				"zhiting",
				"zhixiao",
				"zhiya",
				"zhiye",
				"zhiying",
				"zhiyuan",
				"zhiyue",
				"zhigui",
				"zhishuo",
				"zhimiao-emo",
				"zhimao",
				"zhilun",
				"zhifei",
				"zhida",
				"indah",
				"clara",
				"hanna",
				"beth",
				"betty",
				"cally",
				"cindy",
				"eva",
				"donna",
				"brian",
				"waan",
			}),
		),
	},

	{
		Model: "paraformer-realtime-v2",
		Type:  mode.AudioTranscription,
		Owner: model.ModelOwnerAlibaba,
		RPM:   20,
		Price: model.Price{
			InputPrice: 0.24,
		},
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(10000),
			model.WithModelConfigSupportFormats(
				[]string{"pcm", "wav", "opus", "speex", "aac", "amr"},
			),
		),
	},

	{
		Model: "gte-rerank",
		Type:  mode.Rerank,
		Owner: model.ModelOwnerAlibaba,
		RPM:   300,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(4000),
			model.WithModelConfigMaxInputTokens(4000),
		),
	},
	{
		Model: "gte-rerank-v2",
		Type:  mode.Rerank,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice: 0.0008,
		},
		RPM: 300,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(8000),
			model.WithModelConfigMaxInputTokens(8000),
		),
	},
	{
		Model: "qwen3-vl-rerank",
		Type:  mode.Rerank,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:      0.0005,
			ImageInputPrice: 0.0005,
			VideoInputPrice: 0.0005,
		},
		RPM: 300,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32000),
			model.WithModelConfigMaxInputTokens(32000),
		),
	},

	{
		Model: "text-embedding-v1",
		Type:  mode.Embeddings,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice: 0.0007,
		},
		RPM: 1800,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(2048),
		),
	},
	{
		Model: "text-embedding-v2",
		Type:  mode.Embeddings,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice: 0.0007,
		},
		RPM: 1800,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(2048),
		),
	},
	{
		Model: "text-embedding-v3",
		Type:  mode.Embeddings,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice: 0.0005,
		},
		RPM: 1800,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(8192),
		),
	},
	{
		Model: "text-embedding-v4",
		Type:  mode.Embeddings,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice: 0.0005,
		},
		RPM: 1800,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(8192),
		),
	},
	{
		Model: "qwen3-vl-embedding",
		Type:  mode.Embeddings,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			InputPrice:      0.0007,
			ImageInputPrice: 0.0018,
			VideoInputPrice: 0.0018,
		},
		RPM: 1800,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(8192),
			model.WithModelConfigVision(true),
		),
	},
}
