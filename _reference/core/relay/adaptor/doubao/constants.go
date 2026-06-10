package doubao

import (
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/mode"
)

// https://console.volcengine.com/ark/region:ark+cn-beijing/model
// https://www.volcengine.com/docs/82379/1544106

var ModelList = []model.ModelConfig{
	{
		Model: "doubao-seedream-5-0-lite",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerDoubao,
		Price: model.Price{
			// Seedream image generation bills by the number of successfully generated images.
			ImageOutputPrice:     0.22,
			ImageOutputPriceUnit: 1,
		},
	},
	{
		Model: "doubao-seedream-4-5",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerDoubao,
		Price: model.Price{
			// Seedream image generation bills by the number of successfully generated images.
			ImageOutputPrice:     0.25,
			ImageOutputPriceUnit: 1,
		},
	},
	{
		Model: "doubao-seedream-4-0",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerDoubao,
		Price: model.Price{
			// Seedream image generation bills by the number of successfully generated images.
			ImageOutputPrice:     0.2,
			ImageOutputPriceUnit: 1,
		},
	},
	{
		Model:                     "doubao-seedance-2-0-260128",
		Type:                      mode.DoubaoVideo,
		Owner:                     model.ModelOwnerDoubao,
		AllowedResolutions:        []string{"480p", "720p", "1080p"},
		MaxVideoGenerationSeconds: 15,
		Price: model.Price{
			// Seedance video billing uses the API response usage.completion_tokens.
			// Official prices are RMB per million tokens; aiproxy stores prices per
			// 1K tokens by default, so divide the official value by 1000.
			OutputPrice:     0.051,
			OutputPriceUnit: model.PriceUnit,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{
						Resolution: []string{"480p", "720p"},
						InputVideo: new(false),
					},
					Price: model.Price{
						OutputPrice:     0.046,
						OutputPriceUnit: model.PriceUnit,
					},
				},
				{
					Condition: model.PriceCondition{
						Resolution: []string{"480p", "720p"},
						InputVideo: new(true),
					},
					Price: model.Price{
						OutputPrice:     0.028,
						OutputPriceUnit: model.PriceUnit,
					},
				},
				{
					Condition: model.PriceCondition{
						Resolution: []string{"1080p"},
						InputVideo: new(false),
					},
					Price: model.Price{
						OutputPrice:     0.051,
						OutputPriceUnit: model.PriceUnit,
					},
				},
				{
					Condition: model.PriceCondition{
						Resolution: []string{"1080p"},
						InputVideo: new(true),
					},
					Price: model.Price{
						OutputPrice:     0.031,
						OutputPriceUnit: model.PriceUnit,
					},
				},
			},
		},
	},
	{
		Model:                     "doubao-seedance-2-0-fast-260128",
		Type:                      mode.DoubaoVideo,
		Owner:                     model.ModelOwnerDoubao,
		AllowedResolutions:        []string{"480p", "720p"},
		MaxVideoGenerationSeconds: 15,
		Price: model.Price{
			// Seedance video billing uses the API response usage.completion_tokens.
			OutputPrice:     0.037,
			OutputPriceUnit: model.PriceUnit,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputVideo: new(false)},
					Price: model.Price{
						OutputPrice:     0.037,
						OutputPriceUnit: model.PriceUnit,
					},
				},
				{
					Condition: model.PriceCondition{InputVideo: new(true)},
					Price: model.Price{
						OutputPrice:     0.022,
						OutputPriceUnit: model.PriceUnit,
					},
				},
			},
		},
	},
	{
		Model:                     "doubao-seedance-1-5-pro-251215",
		Type:                      mode.DoubaoVideo,
		Owner:                     model.ModelOwnerDoubao,
		AllowedResolutions:        []string{"480p", "720p", "1080p"},
		MaxVideoGenerationSeconds: 12,
		Price: model.Price{
			// Seedance 1.5 pro defaults generate_audio to true.
			OutputPrice:     0.016,
			OutputPriceUnit: model.PriceUnit,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{
						ServiceTier: "default",
						OutputAudio: new(false),
					},
					Price: model.Price{
						OutputPrice:     0.008,
						OutputPriceUnit: model.PriceUnit,
					},
				},
				{
					Condition: model.PriceCondition{
						ServiceTier: "flex",
						OutputAudio: new(true),
					},
					Price: model.Price{
						OutputPrice:     0.008,
						OutputPriceUnit: model.PriceUnit,
					},
				},
				{
					Condition: model.PriceCondition{
						ServiceTier: "flex",
						OutputAudio: new(false),
					},
					Price: model.Price{
						OutputPrice:     0.004,
						OutputPriceUnit: model.PriceUnit,
					},
				},
			},
		},
	},
	{
		Model:                     "doubao-seedance-1-0-pro-250528",
		Type:                      mode.DoubaoVideo,
		Owner:                     model.ModelOwnerDoubao,
		AllowedResolutions:        []string{"480p", "720p", "1080p"},
		MaxVideoGenerationSeconds: 12,
		Price: model.Price{
			// Seedance 1.0 pro bills by returned video completion tokens.
			OutputPrice:     0.015,
			OutputPriceUnit: model.PriceUnit,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{ServiceTier: "flex"},
					Price: model.Price{
						OutputPrice:     0.0075,
						OutputPriceUnit: model.PriceUnit,
					},
				},
			},
		},
	},
	{
		Model:                     "doubao-seedance-1-0-pro-fast-251015",
		Type:                      mode.DoubaoVideo,
		Owner:                     model.ModelOwnerDoubao,
		AllowedResolutions:        []string{"480p", "720p", "1080p"},
		MaxVideoGenerationSeconds: 12,
		Price: model.Price{
			// Seedance 1.0 pro fast bills by returned video completion tokens.
			OutputPrice:     0.0042,
			OutputPriceUnit: model.PriceUnit,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{ServiceTier: "flex"},
					Price: model.Price{
						OutputPrice:     0.0021,
						OutputPriceUnit: model.PriceUnit,
					},
				},
			},
		},
	},
	{
		Model: "doubao-seed-1-6-250615",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerDoubao,
		Price: model.Price{
			InputPrice:         0.0024,   // 2.40 per million tokens
			OutputPrice:        0.024,    // 24.00 per million tokens
			CacheCreationPrice: 0.000017, // 0.017 per million tokens per hour
			CachedPrice:        0.00016,  // 0.16 per million tokens

			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{
						InputTokenMin:  0,
						InputTokenMax:  32000,
						OutputTokenMin: 0,
						OutputTokenMax: 200,
					},
					Price: model.Price{
						InputPrice:         0.0008,   // 0.80 per million tokens
						OutputPrice:        0.002,    // 2.00 per million tokens
						CacheCreationPrice: 0.000017, // 0.017 per million tokens per hour
						CachedPrice:        0.00016,  // 0.16 per million tokens
					},
				},
				{
					Condition: model.PriceCondition{
						InputTokenMin:  0,
						InputTokenMax:  32000,
						OutputTokenMin: 201,
						OutputTokenMax: 16000,
					},
					Price: model.Price{
						InputPrice:         0.0008,   // 0.80 per million tokens
						OutputPrice:        0.008,    // 8.00 per million tokens
						CacheCreationPrice: 0.000017, // 0.017 per million tokens per hour
						CachedPrice:        0.00016,  // 0.16 per million tokens
					},
				},
				{
					Condition: model.PriceCondition{
						InputTokenMin: 32001,
						InputTokenMax: 128000,
					},
					Price: model.Price{
						InputPrice:         0.0012,   // 1.20 per million tokens
						OutputPrice:        0.016,    // 16.00 per million tokens
						CacheCreationPrice: 0.000017, // 0.017 per million tokens per hour
						CachedPrice:        0.00016,  // 0.16 per million tokens
					},
				},
				{
					Condition: model.PriceCondition{
						InputTokenMin: 128001,
						InputTokenMax: 256000,
					},
					Price: model.Price{
						InputPrice:         0.0024,   // 2.40 per million tokens
						OutputPrice:        0.024,    // 24.00 per million tokens
						CacheCreationPrice: 0.000017, // 0.017 per million tokens per hour
						CachedPrice:        0.00016,  // 0.16 per million tokens
					},
				},
			},
		},
		RPM: 30000,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxOutputTokens(16000),
			model.WithModelConfigMaxInputTokens(224000),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "doubao-seed-1-6-thinking-250615",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerDoubao,
		Price: model.Price{
			InputPrice:         0.0024,   // 2.40 per million tokens
			OutputPrice:        0.024,    // 24.00 per million tokens
			CacheCreationPrice: 0.000017, // 0.017 per million tokens per hour
			CachedPrice:        0.00016,  // 0.16 per million tokens

			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{
						InputTokenMin: 0,
						InputTokenMax: 32000,
					},
					Price: model.Price{
						InputPrice:         0.0008,   // 0.80 per million tokens
						OutputPrice:        0.008,    // 8.00 per million tokens
						CacheCreationPrice: 0.000017, // 0.017 per million tokens per hour
						CachedPrice:        0.00016,  // 0.16 per million tokens
					},
				},
				{
					Condition: model.PriceCondition{
						InputTokenMin: 32001,
						InputTokenMax: 128000,
					},
					Price: model.Price{
						InputPrice:         0.0012,   // 1.20 per million tokens
						OutputPrice:        0.016,    // 16.00 per million tokens
						CacheCreationPrice: 0.000017, // 0.017 per million tokens per hour
						CachedPrice:        0.00016,  // 0.16 per million tokens
					},
				},
				{
					Condition: model.PriceCondition{
						InputTokenMin: 128001,
						InputTokenMax: 256000,
					},
					Price: model.Price{
						InputPrice:         0.0024,   // 2.40 per million tokens
						OutputPrice:        0.024,    // 24.00 per million tokens
						CacheCreationPrice: 0.000017, // 0.017 per million tokens per hour
						CachedPrice:        0.00016,  // 0.16 per million tokens
					},
				},
			},
		},
		RPM: 30000,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxOutputTokens(16000),
			model.WithModelConfigMaxInputTokens(224000),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "doubao-seed-1-6-flash-250615",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerDoubao,
		Price: model.Price{
			InputPrice:         0.0006,   // 0.60 per million tokens
			OutputPrice:        0.006,    // 6.00 per million tokens
			CacheCreationPrice: 0.000017, // 0.017 per million tokens per hour
			CachedPrice:        0.00003,  // 0.03 per million tokens

			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{
						InputTokenMin: 0,
						InputTokenMax: 32000,
					},
					Price: model.Price{
						InputPrice:         0.00015,  // 0.15 per million tokens
						OutputPrice:        0.0015,   // 1.50 per million tokens
						CacheCreationPrice: 0.000017, // 0.017 per million tokens per hour
						CachedPrice:        0.00003,  // 0.03 per million tokens
					},
				},
				{
					Condition: model.PriceCondition{
						InputTokenMin: 32001,
						InputTokenMax: 128000,
					},
					Price: model.Price{
						InputPrice:         0.0003,   // 0.30 per million tokens
						OutputPrice:        0.003,    // 3.00 per million tokens
						CacheCreationPrice: 0.000017, // 0.017 per million tokens per hour
						CachedPrice:        0.00003,  // 0.03 per million tokens
					},
				},
				{
					Condition: model.PriceCondition{
						InputTokenMin: 128001,
						InputTokenMax: 256000,
					},
					Price: model.Price{
						InputPrice:         0.0006,   // 0.60 per million tokens
						OutputPrice:        0.006,    // 6.00 per million tokens
						CacheCreationPrice: 0.000017, // 0.017 per million tokens per hour
						CachedPrice:        0.00003,  // 0.03 per million tokens
					},
				},
			},
		},
		RPM: 30000,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxOutputTokens(16000),
			model.WithModelConfigMaxInputTokens(224000),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "Doubao-1.5-vision-pro-32k",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerDoubao,
		Price: model.Price{
			InputPrice:  0.003,
			OutputPrice: 0.009,
		},
		RPM: 15000,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(32768),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "Doubao-1.5-pro-32k",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerDoubao,
		Price: model.Price{
			InputPrice:  0.0008,
			OutputPrice: 0.0020,
		},
		RPM: 15000,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(32768),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "Doubao-1.5-pro-256k",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerDoubao,
		Price: model.Price{
			InputPrice:  0.005,
			OutputPrice: 0.009,
		},
		RPM: 2000,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxOutputTokens(12000),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "Doubao-1.5-lite-32k",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerDoubao,
		Price: model.Price{
			InputPrice:  0.0003,
			OutputPrice: 0.0006,
		},
		RPM: 15000,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32768),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigToolChoice(true),
		),
	},

	{
		Model: "Doubao-vision-lite-32k",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerDoubao,
		Price: model.Price{
			InputPrice:  0.008,
			OutputPrice: 0.008,
		},
		RPM: 15000,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(32768),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "Doubao-vision-pro-32k",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerDoubao,
		Price: model.Price{
			InputPrice:  0.02,
			OutputPrice: 0.02,
		},
		RPM: 15000,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(32768),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "Doubao-pro-256k",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerDoubao,
		Price: model.Price{
			InputPrice:  0.0050,
			OutputPrice: 0.0090,
		},
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(256000),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "Doubao-pro-128k",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerDoubao,
		Price: model.Price{
			InputPrice:  0.0050,
			OutputPrice: 0.0090,
		},
		RPM: 1000,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(128000),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "Doubao-pro-32k",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerDoubao,
		Price: model.Price{
			InputPrice:  0.0008,
			OutputPrice: 0.0020,
		},
		RPM: 15000,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32768),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "Doubao-pro-4k",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerDoubao,
		Price: model.Price{
			InputPrice:  0.0008,
			OutputPrice: 0.0020,
		},
		RPM: 10000,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(4096),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "Doubao-lite-128k",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerDoubao,
		Price: model.Price{
			InputPrice:  0.0008,
			OutputPrice: 0.0010,
		},
		RPM: 15000,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(128000),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "Doubao-lite-32k",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerDoubao,
		Price: model.Price{
			InputPrice:  0.0003,
			OutputPrice: 0.0006,
		},
		RPM: 15000,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32768),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model: "Doubao-lite-4k",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerDoubao,
		Price: model.Price{
			InputPrice:  0.0003,
			OutputPrice: 0.0006,
		},
		RPM: 10000,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(4096),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigToolChoice(true),
		),
	},

	{
		Model: "Doubao-embedding",
		Type:  mode.Embeddings,
		Owner: model.ModelOwnerDoubao,
		Price: model.Price{
			InputPrice: 0.0005,
		},
		RPM: 1200,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(4096),
		),
	},
	{
		Model: "Doubao-embedding-large",
		Type:  mode.Embeddings,
		Owner: model.ModelOwnerDoubao,
		Price: model.Price{
			InputPrice: 0.0007,
		},
		RPM: 1000,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(4096),
		),
	},
}
