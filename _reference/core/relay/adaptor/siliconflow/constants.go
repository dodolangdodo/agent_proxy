package siliconflow

import (
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/mode"
)

// https://docs.siliconflow.cn/docs/getting-started

var ModelList = []model.ModelConfig{
	{
		Model: "BAAI/bge-reranker-v2-m3",
		Type:  mode.Rerank,
		Owner: model.ModelOwnerBAAI,
		RPM:   2000,
	},

	{
		Model: "BAAI/bge-large-zh-v1.5",
		Type:  mode.Embeddings,
		Owner: model.ModelOwnerBAAI,
		RPM:   2000,
	},
	{
		Model: "Qwen/Qwen3-VL-Embedding-8B",
		Type:  mode.Embeddings,
		Owner: model.ModelOwnerAlibaba,
		Config: model.NewModelConfig(
			model.WithModelConfigVision(true),
		),
		RPM: 2000,
	},

	{
		Model: "fishaudio/fish-speech-1.4",
		Type:  mode.AudioSpeech,
		Owner: model.ModelOwnerFishAudio,
		Config: map[model.ModelConfigKey]any{
			model.ModelConfigSupportVoicesKey: []string{
				"fishaudio/fish-speech-1.4:alex",
				"fishaudio/fish-speech-1.4:benjamin",
				"fishaudio/fish-speech-1.4:charles",
				"fishaudio/fish-speech-1.4:david",
				"fishaudio/fish-speech-1.4:anna",
				"fishaudio/fish-speech-1.4:bella",
				"fishaudio/fish-speech-1.4:claire",
				"fishaudio/fish-speech-1.4:diana",
			},
		},
	},

	{
		Model: "FunAudioLLM/SenseVoiceSmall",
		Type:  mode.AudioTranscription,
		Owner: model.ModelOwnerFunAudioLLM,
	},

	{
		Model: "stabilityai/stable-diffusion-3-5-large",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerStabilityAI,
	},
	{
		Model: "stabilityai/stable-diffusion-3-5-large-turbo",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerStabilityAI,
	},
	{
		Model: "black-forest-labs/FLUX.1-schnell",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerHuggingFace,
		// SiliconFlow image generation bills by successfully returned image count.
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.0014),
			ImageOutputPriceUnit: 1,
		},
	},
	{
		Model: "Tongyi-MAI/Z-Image-Turbo",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.005),
			ImageOutputPriceUnit: 1,
		},
	},
	{
		Model: "black-forest-labs/FLUX.1-dev",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerHuggingFace,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.014),
			ImageOutputPriceUnit: 1,
		},
	},
	{
		Model: "black-forest-labs/FLUX.1-Kontext-dev",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerHuggingFace,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.015),
			ImageOutputPriceUnit: 1,
		},
	},
	{
		Model: "Qwen/Qwen-Image",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.02),
			ImageOutputPriceUnit: 1,
		},
	},
	{
		Model: "black-forest-labs/FLUX.2-pro",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerHuggingFace,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.03),
			ImageOutputPriceUnit: 1,
		},
	},
	{
		Model: "black-forest-labs/FLUX.1-Kontext-pro",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerHuggingFace,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.04),
			ImageOutputPriceUnit: 1,
		},
	},
	{
		Model: "Qwen/Qwen-Image-Edit",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.04),
			ImageOutputPriceUnit: 1,
		},
	},
	{
		Model: "black-forest-labs/FLUX-1.1-pro",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerHuggingFace,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.04),
			ImageOutputPriceUnit: 1,
		},
	},
	{
		Model: "black-forest-labs/FLUX.2-flex",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerHuggingFace,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.06),
			ImageOutputPriceUnit: 1,
		},
	},
	{
		Model: "black-forest-labs/FLUX-1.1-pro-Ultra",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerHuggingFace,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.06),
			ImageOutputPriceUnit: 1,
		},
	},
	{
		Model: "black-forest-labs/FLUX.1-Kontext-max",
		Type:  mode.ImagesGenerations,
		Owner: model.ModelOwnerHuggingFace,
		Price: model.Price{
			ImageOutputPrice:     model.ZeroNullFloat64(0.08),
			ImageOutputPriceUnit: 1,
		},
	},
	{
		Model: "Wan-AI/Wan2.1-T2V-14B-Turbo",
		Type:  mode.VideoGenerationsJobs,
		Owner: model.ModelOwnerAlibaba,
		// SiliconFlow video generation bills by successfully returned video count.
		Price: model.Price{
			OutputPrice:     model.ZeroNullFloat64(0.21),
			OutputPriceUnit: 1,
		},
	},
	{
		Model: "Wan-AI/Wan2.1-I2V-14B-720P-Turbo",
		Type:  mode.VideoGenerationsJobs,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			OutputPrice:     model.ZeroNullFloat64(0.21),
			OutputPriceUnit: 1,
		},
	},
	{
		Model: "Wan-AI/Wan2.2-I2V-A14B",
		Type:  mode.VideoGenerationsJobs,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			OutputPrice:     model.ZeroNullFloat64(0.29),
			OutputPriceUnit: 1,
		},
	},
	{
		Model: "Wan-AI/Wan2.2-T2V-A14B",
		Type:  mode.VideoGenerationsJobs,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			OutputPrice:     model.ZeroNullFloat64(0.29),
			OutputPriceUnit: 1,
		},
	},
	{
		Model: "Wan-AI/Wan2.1-T2V-14B",
		Type:  mode.VideoGenerationsJobs,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			OutputPrice:     model.ZeroNullFloat64(0.29),
			OutputPriceUnit: 1,
		},
	},
	{
		Model: "Wan-AI/Wan2.1-I2V-14B-720P",
		Type:  mode.VideoGenerationsJobs,
		Owner: model.ModelOwnerAlibaba,
		Price: model.Price{
			OutputPrice:     model.ZeroNullFloat64(0.29),
			OutputPriceUnit: 1,
		},
	},
}
