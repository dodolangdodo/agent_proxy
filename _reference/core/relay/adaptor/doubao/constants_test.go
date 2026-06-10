//nolint:testpackage
package doubao

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/common/consume"
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/controller"
	"github.com/labring/aiproxy/core/relay/mode"
)

func doubaoModelPriceForTest(t *testing.T, modelName string) model.Price {
	t.Helper()

	for _, mc := range ModelList {
		if mc.Model == modelName {
			return mc.Price
		}
	}

	t.Fatalf("model %s not found", modelName)

	return model.Price{}
}

func TestDoubaoSeedancePriceUsesReturnedCompletionTokens(t *testing.T) {
	t.Parallel()

	amount := consume.CalculateAmountDetail(
		200,
		model.Usage{
			OutputTokens: 1000,
			TotalTokens:  1000,
		},
		model.UsageContext{
			Resolution: "1080p",
		},
		doubaoModelPriceForTest(t, "doubao-seedance-2-0-260128"),
	)

	if amount.UsedAmount != 0.051 {
		t.Fatalf("expected 0.051 token amount, got %#v", amount)
	}
}

func TestDoubaoSeedancePriceFallsBackToMostExpensiveResolution(t *testing.T) {
	t.Parallel()

	amount := consume.CalculateAmountDetail(
		200,
		model.Usage{
			OutputTokens: 1000,
			TotalTokens:  1000,
		},
		model.UsageContext{},
		doubaoModelPriceForTest(t, "doubao-seedance-2-0-260128"),
	)

	if amount.UsedAmount != 0.051 {
		t.Fatalf("expected 0.051 token amount, got %#v", amount)
	}
}

func TestDoubaoSeedanceConditionalPricesValidate(t *testing.T) {
	t.Parallel()

	for _, mc := range ModelList {
		if mc.Type != mode.DoubaoVideo {
			continue
		}

		if err := mc.Price.ValidateConditionalPrices(); err != nil {
			t.Fatalf("model %s has invalid conditional prices: %v", mc.Model, err)
		}
	}
}

func TestDoubaoSeedanceConditionalPriceUsesInputVideoContext(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		priceModel   string
		usageContext model.UsageContext
		want         float64
	}{
		{
			name:       "seedance 2 720p without input video",
			priceModel: "doubao-seedance-2-0-260128",
			usageContext: model.UsageContext{
				Resolution: "1280x720",
				InputVideo: new(false),
			},
			want: 0.046,
		},
		{
			name:       "seedance 2 720p with input video",
			priceModel: "doubao-seedance-2-0-260128",
			usageContext: model.UsageContext{
				Resolution: "1280x720",
				InputVideo: new(true),
			},
			want: 0.028,
		},
		{
			name:       "seedance 2 1080p with input video",
			priceModel: "doubao-seedance-2-0-260128",
			usageContext: model.UsageContext{
				NativeResolution: "1080p",
				InputVideo:       new(true),
			},
			want: 0.031,
		},
		{
			name:       "seedance 2 fast without input video",
			priceModel: "doubao-seedance-2-0-fast-260128",
			usageContext: model.UsageContext{
				Resolution: "720p",
				InputVideo: new(false),
			},
			want: 0.037,
		},
		{
			name:       "seedance 2 fast with input video",
			priceModel: "doubao-seedance-2-0-fast-260128",
			usageContext: model.UsageContext{
				Resolution: "720p",
				InputVideo: new(true),
			},
			want: 0.022,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			amount := consume.CalculateAmountDetail(
				200,
				model.Usage{OutputTokens: 1000, TotalTokens: 1000},
				tt.usageContext,
				doubaoModelPriceForTest(t, tt.priceModel),
			)

			if amount.UsedAmount != tt.want {
				t.Fatalf("expected %v token amount, got %#v", tt.want, amount)
			}
		})
	}
}

func TestDoubaoSeedance15ConditionalPriceUsesOutputAudioAndServiceTier(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		usageContext model.UsageContext
		want         float64
	}{
		{
			name: "online generated audio",
			usageContext: model.UsageContext{
				ServiceTier: "default",
				OutputAudio: new(true),
			},
			want: 0.016,
		},
		{
			name: "online silent",
			usageContext: model.UsageContext{
				ServiceTier: "default",
				OutputAudio: new(false),
			},
			want: 0.008,
		},
		{
			name: "offline generated audio",
			usageContext: model.UsageContext{
				ServiceTier: "flex",
				OutputAudio: new(true),
			},
			want: 0.008,
		},
		{
			name: "offline silent",
			usageContext: model.UsageContext{
				ServiceTier: "flex",
				OutputAudio: new(false),
			},
			want: 0.004,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			amount := consume.CalculateAmountDetail(
				200,
				model.Usage{OutputTokens: 1000, TotalTokens: 1000},
				tt.usageContext,
				doubaoModelPriceForTest(t, "doubao-seedance-1-5-pro-251215"),
			)

			if amount.UsedAmount != tt.want {
				t.Fatalf("expected %v token amount, got %#v", tt.want, amount)
			}
		})
	}
}

func TestDoubaoSeedance10ConditionalPriceUsesServiceTier(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		priceModel   string
		usageContext model.UsageContext
		want         float64
	}{
		{
			name:       "seedance 1 pro online",
			priceModel: "doubao-seedance-1-0-pro-250528",
			want:       0.015,
		},
		{
			name:       "seedance 1 pro offline",
			priceModel: "doubao-seedance-1-0-pro-250528",
			usageContext: model.UsageContext{
				ServiceTier: "flex",
			},
			want: 0.0075,
		},
		{
			name:       "seedance 1 pro fast online",
			priceModel: "doubao-seedance-1-0-pro-fast-251015",
			want:       0.0042,
		},
		{
			name:       "seedance 1 pro fast offline",
			priceModel: "doubao-seedance-1-0-pro-fast-251015",
			usageContext: model.UsageContext{
				ServiceTier: "flex",
			},
			want: 0.0021,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			amount := consume.CalculateAmountDetail(
				200,
				model.Usage{OutputTokens: 1000, TotalTokens: 1000},
				tt.usageContext,
				doubaoModelPriceForTest(t, tt.priceModel),
			)

			if amount.UsedAmount != tt.want {
				t.Fatalf("expected %v token amount, got %#v", tt.want, amount)
			}
		})
	}
}

func TestDoubaoSeedanceConditionalPriceKeepsTokenUnitAfterVideoController(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/video/generations/jobs",
		bytes.NewBufferString(`{
			"model":"doubao-seedance-2-0-260128",
			"prompt":"A city street",
			"n_seconds":5,
			"size":"1920x1080"
		}`),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	price, err := controller.GetVideoGenerationJobRequestPrice(ctx, model.ModelConfig{
		Price: doubaoModelPriceForTest(t, "doubao-seedance-2-0-260128"),
	})
	if err != nil {
		t.Fatalf("GetVideoGenerationJobRequestPrice returned error: %v", err)
	}

	amount := consume.CalculateAmountDetail(
		http.StatusOK,
		model.Usage{
			OutputTokens: 1000,
			TotalTokens:  1000,
		},
		model.UsageContext{
			Resolution: "1080p",
		},
		price,
	)

	if amount.UsedAmount != 0.051 {
		t.Fatalf("expected 0.051 token amount after controller price projection, got %#v", amount)
	}
}
