package gemini

import (
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/mode"
)

// https://ai.google.dev/models/gemini
// https://ai.google.dev/gemini-api/docs/pricing

var ModelList = []model.ModelConfig{
	{
		Model: "gemini-3.1-pro-preview",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerGoogle,
		Price: model.Price{
			InputPrice:         0.004,
			OutputPrice:        0.018,
			CachedPrice:        0.0004,
			WebSearchPrice:     0.014,
			WebSearchPriceUnit: 1,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 200000, ServiceTier: "priority"},
					Price: model.Price{
						InputPrice:         0.0036,
						OutputPrice:        0.0216,
						CachedPrice:        0.00036,
						WebSearchPrice:     0.014,
						WebSearchPriceUnit: 1,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMax: 200000},
					Price: model.Price{
						InputPrice:         0.002,
						OutputPrice:        0.012,
						CachedPrice:        0.0002,
						WebSearchPrice:     0.014,
						WebSearchPriceUnit: 1,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 200001, ServiceTier: "priority"},
					Price: model.Price{
						InputPrice:         0.0072,
						OutputPrice:        0.0324,
						CachedPrice:        0.00072,
						WebSearchPrice:     0.014,
						WebSearchPriceUnit: 1,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 200001},
					Price: model.Price{
						InputPrice:         0.004,
						OutputPrice:        0.018,
						CachedPrice:        0.0004,
						WebSearchPrice:     0.014,
						WebSearchPriceUnit: 1,
					},
				},
			},
		},
		SummaryServiceTier: true,
		RPM:                600,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1048576),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "gemini-3-flash-preview",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerGoogle,
		Price: model.Price{
			InputPrice:         0.0005,
			AudioInputPrice:    0.001,
			OutputPrice:        0.003,
			CachedPrice:        0.00005,
			WebSearchPrice:     0.014,
			WebSearchPriceUnit: 1,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{ServiceTier: "priority"},
					Price: model.Price{
						InputPrice:         0.0009,
						AudioInputPrice:    0.0018,
						OutputPrice:        0.0054,
						CachedPrice:        0.00009,
						WebSearchPrice:     0.014,
						WebSearchPriceUnit: 1,
					},
				},
			},
		},
		SummaryServiceTier: true,
		RPM:                600,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1048576),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "gemini-3.1-flash-lite-preview",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerGoogle,
		Price: model.Price{
			InputPrice:         0.00025,
			AudioInputPrice:    0.0005,
			OutputPrice:        0.0015,
			CachedPrice:        0.000025,
			WebSearchPrice:     0.014,
			WebSearchPriceUnit: 1,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{ServiceTier: "priority"},
					Price: model.Price{
						InputPrice:         0.00045,
						AudioInputPrice:    0.0009,
						OutputPrice:        0.0027,
						CachedPrice:        0.000045,
						WebSearchPrice:     0.014,
						WebSearchPriceUnit: 1,
					},
				},
			},
		},
		SummaryServiceTier: true,
		RPM:                600,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1048576),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "gemini-3.1-flash-image-preview",
		Type:  mode.GeminiImage,
		Owner: model.ModelOwnerGoogle,
		Price: model.Price{
			InputPrice:         0.0005,
			ImageInputPrice:    0.0005,
			OutputPrice:        0.003,
			ImageOutputPrice:   0.060,
			WebSearchPrice:     0.014,
			WebSearchPriceUnit: 1,
		},
	},
	{
		Model: "gemini-2.5-flash",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerGoogle,
		Price: model.Price{
			// Gemini usage metadata separates AUDIO prompt tokens; text/image/video
			// prompt tokens remain in InputTokens and use the base input price.
			InputPrice:      0.0003,
			AudioInputPrice: 0.001,
			OutputPrice:     0.0025,
			CachedPrice:     0.00003,
			// Grounding with Google Search is billed per grounded prompt, not per
			// returned query; Gemini usage records 1 when grounding metadata appears.
			WebSearchPrice:     0.035,
			WebSearchPriceUnit: 1,
		},
		RPM: 600,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1048576),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "gemini-2.5-flash-lite",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerGoogle,
		Price: model.Price{
			// Gemini usage metadata separates AUDIO prompt tokens; text/image/video
			// prompt tokens remain in InputTokens and use the base input price.
			InputPrice:      0.0001,
			AudioInputPrice: 0.0003,
			OutputPrice:     0.0004,
			CachedPrice:     0.00001,
			// Grounding with Google Search is billed per grounded prompt, not per
			// returned query; Gemini usage records 1 when grounding metadata appears.
			WebSearchPrice:     0.035,
			WebSearchPriceUnit: 1,
		},
		RPM: 600,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1048576),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "gemini-2.5-pro",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerGoogle,
		Price: model.Price{
			// Gemini usage metadata separates AUDIO prompt tokens; text/image/video
			// prompt tokens remain in InputTokens and use the base input price.
			InputPrice:      0.0025,
			AudioInputPrice: 0.001,
			OutputPrice:     0.015,
			CachedPrice:     0.000625,
			// Grounding with Google Search is billed per grounded prompt, not per
			// returned query; Gemini usage records 1 when grounding metadata appears.
			WebSearchPrice:     0.035,
			WebSearchPriceUnit: 1,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 200000},
					Price: model.Price{
						InputPrice:         0.00125,
						AudioInputPrice:    0.001,
						OutputPrice:        0.01,
						CachedPrice:        0.000125,
						WebSearchPrice:     0.035,
						WebSearchPriceUnit: 1,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 200001},
					Price: model.Price{
						InputPrice:         0.0025,
						AudioInputPrice:    0.001,
						OutputPrice:        0.015,
						CachedPrice:        0.000625,
						WebSearchPrice:     0.035,
						WebSearchPriceUnit: 1,
					},
				},
			},
		},
		RPM: 600,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1048576),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "gemini-2.5-flash-lite-preview-09-2025",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerGoogle,
		Price: model.Price{
			InputPrice:         0.0001,
			AudioInputPrice:    0.0003,
			OutputPrice:        0.0004,
			CachedPrice:        0.00001,
			WebSearchPrice:     0.035,
			WebSearchPriceUnit: 1,
		},
		RPM: 600,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1048576),
			model.WithModelConfigMaxOutputTokens(65536),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "gemini-2.5-flash-image",
		Type:  mode.GeminiImage,
		Owner: model.ModelOwnerGoogle,
		Price: model.Price{
			InputPrice:       0.0003,
			ImageInputPrice:  0.0003,
			ImageOutputPrice: 0.030,
		},
	},
	{
		Model: "gemini-2.5-flash-tts",
		Type:  mode.GeminiTTS,
		Owner: model.ModelOwnerGoogle,
		Price: model.Price{
			// Gemini TTS reports generated speech as AUDIO candidate tokens.
			InputPrice:       0.0005,
			AudioOutputPrice: 0.01,
		},
		RPM: 600,
	},
	{
		Model: "gemini-3-pro-image-preview",
		Type:  mode.GeminiImage,
		Owner: model.ModelOwnerGoogle,
		Price: model.Price{
			InputPrice:         0.004,
			ImageInputPrice:    0.002,
			OutputPrice:        0.012,
			ImageOutputPrice:   0.120,
			WebSearchPrice:     0.014,
			WebSearchPriceUnit: 1,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{InputTokenMax: 200000},
					Price: model.Price{
						InputPrice:         0.002,
						ImageInputPrice:    0.002,
						OutputPrice:        0.012,
						ImageOutputPrice:   0.120,
						WebSearchPrice:     0.014,
						WebSearchPriceUnit: 1,
					},
				},
				{
					Condition: model.PriceCondition{InputTokenMin: 200001},
					Price: model.Price{
						InputPrice:         0.004,
						ImageInputPrice:    0.002,
						OutputPrice:        0.012,
						ImageOutputPrice:   0.120,
						WebSearchPrice:     0.014,
						WebSearchPriceUnit: 1,
					},
				},
			},
		},
	},
	{
		Model: "veo-3.1-generate-preview",
		Type:  mode.GeminiVideo,
		Owner: model.ModelOwnerGoogle,
		Price: model.Price{
			// Veo billing is based on successfully generated video seconds.
			OutputPrice:     0.4,
			OutputPriceUnit: 1,
		},
	},
	{
		Model: "veo-3.1-fast-generate-preview",
		Type:  mode.GeminiVideo,
		Owner: model.ModelOwnerGoogle,
		Price: model.Price{
			// Veo async usage stores one output token per successfully generated
			// second. Resolution conditions select the official per-second tier.
			OutputPrice:     0.30,
			OutputPriceUnit: 1,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{Resolution: []string{"720p"}},
					Price: model.Price{
						OutputPrice:     model.ZeroNullFloat64(0.10),
						OutputPriceUnit: 1,
					},
				},
				{
					Condition: model.PriceCondition{Resolution: []string{"1080p"}},
					Price: model.Price{
						OutputPrice:     model.ZeroNullFloat64(0.12),
						OutputPriceUnit: 1,
					},
				},
				{
					Condition: model.PriceCondition{Resolution: []string{"4k"}},
					Price: model.Price{
						OutputPrice:     model.ZeroNullFloat64(0.30),
						OutputPriceUnit: 1,
					},
				},
			},
		},
	},
	{
		Model: "veo-3.1-lite-generate-preview",
		Type:  mode.GeminiVideo,
		Owner: model.ModelOwnerGoogle,
		Price: model.Price{
			// Veo async usage stores one output token per successfully generated
			// second. Resolution conditions select the official per-second tier.
			OutputPrice:     0.08,
			OutputPriceUnit: 1,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{Resolution: []string{"720p"}},
					Price: model.Price{
						OutputPrice:     model.ZeroNullFloat64(0.05),
						OutputPriceUnit: 1,
					},
				},
				{
					Condition: model.PriceCondition{Resolution: []string{"1080p"}},
					Price: model.Price{
						OutputPrice:     model.ZeroNullFloat64(0.08),
						OutputPriceUnit: 1,
					},
				},
			},
		},
	},
	{
		Model: "gemini-1.5-pro",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerGoogle,
		Price: model.Price{
			InputPrice:  0.0025,
			OutputPrice: 0.01,
		},
		RPM: 600,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(2097152),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "gemini-1.5-flash",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerGoogle,
		Price: model.Price{
			InputPrice:  0.00015,
			OutputPrice: 0.0006,
		},
		RPM: 600,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1048576),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "gemini-1.5-flash-8b",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerGoogle,
		Price: model.Price{
			InputPrice:  0.000075,
			OutputPrice: 0.0003,
		},
		RPM: 600,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1048576),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "gemini-2.0-flash",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerGoogle,
		Price: model.Price{
			InputPrice:  0.0001,
			OutputPrice: 0.0004,
		},
		RPM: 600,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1048576),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "gemini-2.0-flash-lite-preview",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerGoogle,
		Price: model.Price{
			InputPrice:  0.000075,
			OutputPrice: 0.0003,
		},
		RPM: 600,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1048576),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "gemini-2.0-flash-thinking-exp",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerGoogle,
		Price: model.Price{
			InputPrice:  0.0001,
			OutputPrice: 0.0004,
		},
		RPM: 600,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1048576),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model: "gemini-2.0-pro-exp",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerGoogle,
		Price: model.Price{
			InputPrice:  0.0025,
			OutputPrice: 0.01,
		},
		RPM: 600,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(2097152),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
			model.WithModelConfigVision(true),
		),
	},

	{
		Model: "text-embedding-004",
		Type:  mode.Embeddings,
		Owner: model.ModelOwnerGoogle,
		Price: model.Price{
			InputPrice: 0.0001,
		},
		RPM: 1500,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(2048),
			model.WithModelConfigMaxOutputTokens(768),
		),
	},
}
