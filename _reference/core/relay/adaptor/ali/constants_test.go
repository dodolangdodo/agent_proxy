//nolint:testpackage
package ali

import (
	"testing"

	"github.com/bytedance/sonic"
	"github.com/labring/aiproxy/core/common/consume"
	"github.com/labring/aiproxy/core/model"
)

func aliModelPriceForTest(t *testing.T, modelName string) model.Price {
	t.Helper()

	for _, mc := range ModelList {
		if mc.Model == modelName {
			return mc.Price
		}
	}

	t.Fatalf("model %s not found", modelName)

	return model.Price{}
}

func TestAliModelListPricesAreValid(t *testing.T) {
	t.Parallel()

	for _, mc := range ModelList {
		t.Run(mc.Model, func(t *testing.T) {
			t.Parallel()

			if err := mc.Price.ValidateConditionalPrices(); err != nil {
				t.Fatalf("invalid conditional prices: %v", err)
			}
		})
	}
}

func TestAliImagePriceUsesSuccessfulImageCount(t *testing.T) {
	t.Parallel()

	amount := consume.CalculateAmountDetail(
		200,
		model.Usage{
			OutputTokens:      3,
			ImageOutputTokens: 3,
			TotalTokens:       3,
		},
		model.UsageContext{},
		aliModelPriceForTest(t, "qwen-image-plus"),
	)

	if amount.UsedAmount != 0.6 {
		t.Fatalf("expected 0.6 image amount, got %#v", amount)
	}
}

func TestAliVideoPriceUsesOutputSecondsAndResolution(t *testing.T) {
	t.Parallel()

	amount := consume.CalculateAmountDetail(
		200,
		model.Usage{
			OutputTokens: 5,
			TotalTokens:  5,
		},
		model.UsageContext{
			Resolution: "720p",
		},
		aliModelPriceForTest(t, "wan2.5-t2v-preview"),
	)

	if amount.UsedAmount != 3 {
		t.Fatalf("expected 3 video amount, got %#v", amount)
	}
}

func TestAliVideoPriceFallsBackToMostExpensiveResolution(t *testing.T) {
	t.Parallel()

	amount := consume.CalculateAmountDetail(
		200,
		model.Usage{
			OutputTokens: 5,
			TotalTokens:  5,
		},
		model.UsageContext{},
		aliModelPriceForTest(t, "wan2.5-t2v-preview"),
	)

	if amount.UsedAmount != 5 {
		t.Fatalf("expected 5 video amount, got %#v", amount)
	}
}

func TestAliVideoEditPriceUsesInputAndOutputSeconds(t *testing.T) {
	t.Parallel()

	amount := consume.CalculateAmountDetail(
		200,
		model.Usage{
			VideoInputTokens: 2,
			OutputTokens:     5,
			TotalTokens:      7,
		},
		model.UsageContext{
			Resolution: "1080p",
		},
		aliModelPriceForTest(t, "happyhorse-1.0-video-edit"),
	)

	if amount.UsedAmount != 11.2 {
		t.Fatalf("expected 11.2 video edit amount, got %#v", amount)
	}
}

func TestAliQwenOmniAudioOutputOnlyChargesAudioOutput(t *testing.T) {
	t.Parallel()

	usage := model.Usage{
		InputTokens:       1000,
		AudioInputTokens:  250,
		OutputTokens:      1000,
		AudioOutputTokens: 800,
		TotalTokens:       2000,
	}

	amount := consume.CalculateAmountDetail(
		200,
		usage,
		aliChatUsageContext(usage),
		aliModelPriceForTest(t, "qwen3-omni-flash"),
	)

	if amount.OutputAmount != 0 {
		t.Fatalf("expected no text output amount for audio output, got %#v", amount)
	}

	if amount.AudioOutputAmount != 0.05008 {
		t.Fatalf("expected audio output amount 0.05008, got %#v", amount)
	}
}

func TestAliQwenOmniTextOutputUsesInputMediaContext(t *testing.T) {
	t.Parallel()

	textUsage := model.Usage{
		InputTokens:  1000,
		OutputTokens: 1000,
		TotalTokens:  2000,
	}
	textAmount := consume.CalculateAmountDetail(
		200,
		textUsage,
		aliChatUsageContextWithDefaults(aliChatUsageContext(textUsage)),
		aliModelPriceForTest(t, "qwen-omni-turbo"),
	)

	if textAmount.OutputAmount != 0.0016 {
		t.Fatalf("expected pure text output amount 0.0016, got %#v", textAmount)
	}

	imageUsage := model.Usage{
		InputTokens:      1000,
		ImageInputTokens: 300,
		OutputTokens:     1000,
		TotalTokens:      2000,
	}
	imageAmount := consume.CalculateAmountDetail(
		200,
		imageUsage,
		aliChatUsageContext(imageUsage),
		aliModelPriceForTest(t, "qwen-omni-turbo"),
	)

	if imageAmount.OutputAmount != 0.0045 {
		t.Fatalf("expected multimodal output amount 0.0045, got %#v", imageAmount)
	}
}

func TestAliChatRequestUsageContextDetectsMediaAndAudioOutput(t *testing.T) {
	t.Parallel()

	imageContext := aliChatRequestUsageContextForTest(t, `{
		"messages": [
			{
				"role": "user",
				"content": [
					{"type": "image_url", "image_url": {"url": "https://example.com/a.png"}},
					{"type": "text", "text": "describe"}
				]
			}
		]
	}`)

	if imageContext.InputMedia == nil || !*imageContext.InputMedia {
		t.Fatalf("expected image request to set input media, got %#v", imageContext)
	}

	if imageContext.InputVideo != nil && *imageContext.InputVideo {
		t.Fatalf("expected image request not to set input video, got %#v", imageContext)
	}

	videoContext := aliChatRequestUsageContextForTest(t, `{
		"messages": [
			{
				"role": "user",
				"content": [
					{"type": "video_url", "video_url": {"url": "https://example.com/a.mp4"}}
				]
			}
		]
	}`)

	if videoContext.InputMedia == nil || !*videoContext.InputMedia {
		t.Fatalf("expected video request to set input media, got %#v", videoContext)
	}

	if videoContext.InputVideo == nil || !*videoContext.InputVideo {
		t.Fatalf("expected video request to set input video, got %#v", videoContext)
	}

	audioOutputContext := aliChatRequestUsageContextForTest(t, `{
		"messages": [{"role": "user", "content": "say hello"}],
		"modalities": ["text", "audio"],
		"audio": {"voice": "Cherry", "format": "wav"}
	}`)

	if audioOutputContext.OutputAudio == nil || !*audioOutputContext.OutputAudio {
		t.Fatalf("expected audio output request to set output audio, got %#v", audioOutputContext)
	}
}

func aliChatRequestUsageContextForTest(t *testing.T, body string) model.UsageContext {
	t.Helper()

	node, err := sonic.GetFromString(body)
	if err != nil {
		t.Fatalf("parse request body: %v", err)
	}

	return aliChatRequestUsageContext(&node)
}
