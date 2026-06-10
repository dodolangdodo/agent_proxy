//nolint:testpackage
package gemini

import (
	"testing"

	"github.com/labring/aiproxy/core/common/consume"
	"github.com/labring/aiproxy/core/model"
)

func geminiModelPriceForTest(t *testing.T, modelName string) model.Price {
	t.Helper()

	for _, mc := range ModelList {
		if mc.Model == modelName {
			return mc.Price
		}
	}

	t.Fatalf("model %s not found", modelName)

	return model.Price{}
}

func TestGeminiTTSPriceUsesAudioOutputTokens(t *testing.T) {
	t.Parallel()

	amount := consume.CalculateAmountDetail(
		200,
		model.Usage{
			InputTokens:       1000,
			OutputTokens:      2000,
			AudioOutputTokens: 2000,
			TotalTokens:       3000,
		},
		model.UsageContext{},
		model.Price{
			InputPrice:       0.0005,
			AudioOutputPrice: 0.01,
		},
	)

	if amount.UsedAmount != 0.0205 {
		t.Fatalf("expected 0.0205 amount, got %#v", amount)
	}
}

func TestGeminiFlashPriceUsesSeparateAudioInputTokens(t *testing.T) {
	t.Parallel()

	amount := consume.CalculateAmountDetail(
		200,
		model.Usage{
			InputTokens:      1500,
			AudioInputTokens: 500,
			OutputTokens:     1000,
			TotalTokens:      2500,
		},
		model.UsageContext{},
		model.Price{
			InputPrice:      0.0003,
			AudioInputPrice: 0.001,
			OutputPrice:     0.0025,
		},
	)

	if amount.UsedAmount != 0.0033 {
		t.Fatalf("expected 0.0033 amount, got %#v", amount)
	}
}

func TestGeminiSearchGroundingPriceUsesOneGroundedPrompt(t *testing.T) {
	t.Parallel()

	amount := consume.CalculateAmountDetail(
		200,
		model.Usage{
			WebSearchCount: 1,
		},
		model.UsageContext{},
		model.Price{
			WebSearchPrice:     0.035,
			WebSearchPriceUnit: 1,
		},
	)

	if amount.UsedAmount != 0.035 {
		t.Fatalf("expected 0.035 amount, got %#v", amount)
	}
}

func TestGeminiImagePriceUsesSeparateImageInputTokens(t *testing.T) {
	t.Parallel()

	amount := consume.CalculateAmountDetail(
		200,
		model.Usage{
			InputTokens:       1000,
			ImageInputTokens:  400,
			OutputTokens:      1120,
			ImageOutputTokens: 1120,
			TotalTokens:       2120,
		},
		model.UsageContext{},
		geminiModelPriceForTest(t, "gemini-3-pro-image-preview"),
	)

	if amount.InputAmount != 0.0012 {
		t.Fatalf("expected text input amount 0.0012, got %#v", amount)
	}

	if amount.ImageInputAmount != 0.0008 {
		t.Fatalf("expected image input amount 0.0008, got %#v", amount)
	}

	if amount.UsedAmount != 0.1364 {
		t.Fatalf("expected 0.1364 amount, got %#v", amount)
	}
}

func TestGeminiVideoPriceUsesGeneratedSeconds(t *testing.T) {
	t.Parallel()

	amount := consume.CalculateAmountDetail(
		200,
		model.Usage{
			OutputTokens: 8,
			TotalTokens:  8,
		},
		model.UsageContext{},
		model.Price{
			OutputPrice:     0.4,
			OutputPriceUnit: 1,
		},
	)

	if amount.UsedAmount != 3.2 {
		t.Fatalf("expected 3.2 amount, got %#v", amount)
	}
}

func TestGeminiVeoFastPriceUsesResolutionTier(t *testing.T) {
	t.Parallel()

	amount := consume.CalculateAmountDetail(
		200,
		model.Usage{
			OutputTokens: 8,
			TotalTokens:  8,
		},
		model.UsageContext{
			Resolution: "1080p",
		},
		geminiModelPriceForTest(t, "veo-3.1-fast-generate-preview"),
	)

	if amount.UsedAmount != 0.96 {
		t.Fatalf("expected 0.96 amount, got %#v", amount)
	}
}

func TestGeminiVeoFastPriceFallsBackToMostExpensiveResolution(t *testing.T) {
	t.Parallel()

	amount := consume.CalculateAmountDetail(
		200,
		model.Usage{
			OutputTokens: 8,
			TotalTokens:  8,
		},
		model.UsageContext{},
		geminiModelPriceForTest(t, "veo-3.1-fast-generate-preview"),
	)

	if amount.UsedAmount != 2.4 {
		t.Fatalf("expected 2.4 amount, got %#v", amount)
	}
}
