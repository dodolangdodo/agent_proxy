//nolint:testpackage
package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/model"
	"github.com/stretchr/testify/require"
)

func TestGetGeminiVideoRequestUsageJSON(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"instances":[{"prompt":"A city street"}],
		"parameters":{"durationSeconds":6,"numberOfVideos":2}
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1beta/models/veo-3.1-generate-preview:predictLongRunning",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	usage, err := GetGeminiVideoRequestUsage(ctx, model.ModelConfig{})
	require.NoError(t, err)
	require.Zero(t, usage.Usage.OutputTokens)
	require.Zero(t, usage.Usage.TotalTokens)
	require.Equal(t, "720p", usage.Context.Resolution)
}

func TestGetGeminiVideoRequestUsageReadsTopLevelNativeParameters(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"instances":[{"prompt":"A city street"}],
		"durationSeconds":7,
		"numberOfVideos":3
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1beta/models/veo-3.1-generate-preview:predictLongRunning",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	usage, err := GetGeminiVideoRequestUsage(ctx, model.ModelConfig{})
	require.NoError(t, err)
	require.Zero(t, usage.Usage.OutputTokens)
	require.Zero(t, usage.Usage.TotalTokens)
}

func TestGetGeminiVideoRequestUsageTopLevelZeroFallsBackToParameters(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"instances":[{"prompt":"A city street"}],
		"durationSeconds":0,
		"numberOfVideos":0,
		"parameters":{"durationSeconds":6,"numberOfVideos":2}
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1beta/models/veo-3.1-generate-preview:predictLongRunning",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	usage, err := GetGeminiVideoRequestUsage(ctx, model.ModelConfig{})
	require.NoError(t, err)
	require.Zero(t, usage.Usage.OutputTokens)
	require.Zero(t, usage.Usage.TotalTokens)
}

func TestValidateGeminiVideoRequestRejectsTooLongDuration(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"instances":[{"prompt":"A city street"}],
		"parameters":{"durationSeconds":6,"numberOfVideos":1}
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1beta/models/veo-3.1-generate-preview:predictLongRunning",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	err := ValidateGeminiVideoRequest(ctx, model.ModelConfig{
		MaxVideoGenerationSeconds: 5,
	})
	require.Error(t, err)
	require.Equal(t, "seconds must be less than or equal to 5", err.Error())

	var requestParamErr *RequestParamError
	require.ErrorAs(t, err, &requestParamErr)
	require.Equal(t, 400, requestParamErr.StatusCode)
}

func TestValidateGeminiVideoRequestRejectsTooLongTopLevelDuration(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"instances":[{"prompt":"A city street"}],
		"durationSeconds":60,
		"numberOfVideos":1
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1beta/models/veo-3.1-generate-preview:predictLongRunning",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	err := ValidateGeminiVideoRequest(ctx, model.ModelConfig{
		MaxVideoGenerationSeconds: 15,
	})
	require.Error(t, err)
	require.Equal(t, "seconds must be less than or equal to 15", err.Error())
}

func TestValidateGeminiVideoRequestRejectsTooManyVideos(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"instances":[{"prompt":"A city street"}],
		"parameters":{"durationSeconds":6,"numberOfVideos":3}
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1beta/models/veo-3.1-generate-preview:predictLongRunning",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	err := ValidateGeminiVideoRequest(ctx, model.ModelConfig{
		MaxVideoGenerationCount: 2,
	})
	require.Error(t, err)
	require.Equal(t, "video count must be less than or equal to 2", err.Error())
}

func TestValidateGeminiVideoRequestRejectsUnsupportedResolution(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"instances":[{"prompt":"A city street"}],
		"parameters":{"durationSeconds":6,"resolution":"1080p"}
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1beta/models/veo-3.1-generate-preview:predictLongRunning",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	err := ValidateGeminiVideoRequest(ctx, model.ModelConfig{
		AllowedResolutions: []string{"720p"},
	})
	require.Error(t, err)
	require.Equal(
		t,
		"unsupported video resolution `1080p`, supported resolutions: 720p",
		err.Error(),
	)
}

func TestValidateGeminiVideoRequestRejectsInvalidResolutionFormat(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"instances":[{"prompt":"A city street"}],
		"parameters":{"durationSeconds":6,"resolution":"1280x720"}
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1beta/models/veo-3.1-generate-preview:predictLongRunning",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	err := ValidateGeminiVideoRequest(ctx, model.ModelConfig{})
	require.Error(t, err)
	require.Equal(
		t,
		"invalid gemini video resolution `1280x720`, supported resolutions: 720p, 1080p, 4k",
		err.Error(),
	)
}

func TestValidateGeminiVideoRequestResolutionMatrix(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		resolution string
		allowed    []string
		disable    bool
		wantErr    bool
		wantErrMsg string
	}{
		{
			name:       "exact native tier",
			resolution: "720p",
			allowed:    []string{"720p"},
		},
		{
			name:       "case-insensitive native tier",
			resolution: "720P",
			allowed:    []string{"720p"},
		},
		{
			name:       "allowed dimensions fuzzy match native tier",
			resolution: "720p",
			allowed:    []string{"1280x720"},
		},
		{
			name:       "disabled fuzzy rejects allowed dimensions",
			resolution: "720p",
			allowed:    []string{"1280x720"},
			disable:    true,
			wantErr:    true,
			wantErrMsg: "unsupported video resolution `720p`, supported resolutions: none",
		},
		{
			name:       "invalid native dimension request",
			resolution: "1280x720",
			allowed:    []string{"720p"},
			wantErr:    true,
			wantErrMsg: "invalid gemini video resolution `1280x720`, supported resolutions: 720p",
		},
		{
			name:       "blank allowed means unsupported",
			resolution: "720p",
			allowed:    []string{" ", "\t"},
			wantErr:    true,
			wantErrMsg: "unsupported video resolution `720p`, supported resolutions: none",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			body := `{
				"instances":[{"prompt":"A city street"}],
				"parameters":{"durationSeconds":6,"resolution":` + strconv.Quote(tt.resolution) + `}
			}`
			req := httptest.NewRequestWithContext(
				t.Context(),
				http.MethodPost,
				"/v1beta/models/veo-3.1-generate-preview:predictLongRunning",
				bytes.NewBufferString(body),
			)
			req.Header.Set("Content-Type", "application/json")

			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			ctx.Request = req

			err := ValidateGeminiVideoRequest(ctx, model.ModelConfig{
				AllowedResolutions:          tt.allowed,
				DisableResolutionFuzzyMatch: tt.disable,
			})
			if tt.wantErr {
				require.Error(t, err)
				require.Equal(t, tt.wantErrMsg, err.Error())
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestGetGeminiVideoRequestUsageSetsResolutionCondition(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"instances":[{"prompt":"A city street"}],
		"parameters":{"durationSeconds":6,"numberOfVideos":2,"resolution":"1080p"}
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1beta/models/veo-3.1-generate-preview:predictLongRunning",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	usage, err := GetGeminiVideoRequestUsage(ctx, model.ModelConfig{
		AllowedResolutions: []string{"1080p"},
	})
	require.NoError(t, err)
	require.Zero(t, usage.Usage.OutputTokens)
	require.Equal(t, "1080p", usage.Context.Resolution)
}

func TestGetGeminiVideoRequestUsageIgnoresAspectRatioForResolutionCondition(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"instances":[{"prompt":"A city street"}],
		"parameters":{"durationSeconds":6,"numberOfVideos":2,"aspectRatio":"16:9"}
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1beta/models/veo-3.1-generate-preview:predictLongRunning",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	usage, err := GetGeminiVideoRequestUsage(ctx, model.ModelConfig{
		AllowedResolutions: []string{"720p"},
	})
	require.NoError(t, err)
	require.Zero(t, usage.Usage.OutputTokens)
	require.Equal(t, "720p", usage.Context.Resolution)
}

func TestGetGeminiVideoRequestUsageRejectsWhenDefaultResolutionUnsupported(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"instances":[{"prompt":"A city street"}],
		"parameters":{"durationSeconds":6,"aspectRatio":"1280x720"}
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1beta/models/veo-3.1-generate-preview:predictLongRunning",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	_, err := GetGeminiVideoRequestUsage(ctx, model.ModelConfig{
		AllowedResolutions: []string{"1080p"},
	})
	require.Error(t, err)
	require.Equal(
		t,
		"unsupported video resolution `720p`, supported resolutions: 1080p",
		err.Error(),
	)
}

func TestGetGeminiVideoRequestUsageResolutionOverridesAspectRatio(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"instances":[{"prompt":"A city street"}],
		"parameters":{"durationSeconds":6,"resolution":"1080p","aspectRatio":"16:9"}
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1beta/models/veo-3.1-generate-preview:predictLongRunning",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	usage, err := GetGeminiVideoRequestUsage(ctx, model.ModelConfig{
		AllowedResolutions: []string{"1080p"},
	})
	require.NoError(t, err)
	require.Equal(t, "1080p", usage.Context.Resolution)
}

func TestGetGeminiVideoRequestUsageDefaultsWhenParametersMissing(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{"instances":[{"prompt":"A city street"}]}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1beta/models/veo-3.1-generate-preview:predictLongRunning",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	usage, err := GetGeminiVideoRequestUsage(ctx, model.ModelConfig{})
	require.NoError(t, err)
	require.Zero(t, usage.Usage.OutputTokens)
	require.Zero(t, usage.Usage.TotalTokens)
	require.Equal(t, "720p", usage.Context.Resolution)
}

func TestGetGeminiVideoRequestUsageIgnoresVertexSampleCount(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"instances":[{"prompt":"A city street"}],
		"parameters":{"durationSeconds":6,"sampleCount":2}
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1beta/models/veo-3.1-generate-preview:predictLongRunning",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	usage, err := GetGeminiVideoRequestUsage(ctx, model.ModelConfig{})
	require.NoError(t, err)
	require.Zero(t, usage.Usage.OutputTokens)
	require.Zero(t, usage.Usage.TotalTokens)
}

func TestGetGeminiVideoRequestUsageRejectsUnsupportedResolution(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"instances":[{"prompt":"A city street"}],
		"parameters":{"durationSeconds":6,"resolution":"1080p"}
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1beta/models/veo-3.1-generate-preview:predictLongRunning",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	_, err := GetGeminiVideoRequestUsage(ctx, model.ModelConfig{
		AllowedResolutions: []string{"720p"},
	})
	require.Error(t, err)
	require.Equal(
		t,
		"unsupported video resolution `1080p`, supported resolutions: 720p",
		err.Error(),
	)
}

func TestGetGeminiVideoRequestUsageRejectsTooLongDuration(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"instances":[{"prompt":"A city street"}],
		"parameters":{"durationSeconds":6,"numberOfVideos":1}
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1beta/models/veo-3.1-generate-preview:predictLongRunning",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	_, err := GetGeminiVideoRequestUsage(ctx, model.ModelConfig{
		MaxVideoGenerationSeconds: 5,
	})
	require.Error(t, err)
	require.Equal(t, "seconds must be less than or equal to 5", err.Error())

	var requestParamErr *RequestParamError
	require.ErrorAs(t, err, &requestParamErr)
	require.Equal(t, 400, requestParamErr.StatusCode)
}

func TestGetGeminiVideoRequestUsageRejectsNegativeNativeDuration(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"instances":[{"prompt":"A city street"}],
		"parameters":{"durationSeconds":-1}
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1beta/models/veo-3.1-generate-preview:predictLongRunning",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	_, err := GetGeminiVideoRequestUsage(ctx, model.ModelConfig{})
	require.Error(t, err)
	require.Equal(t, "invalid durationSeconds: must be non-negative", err.Error())
}

func TestGetGeminiVideoRequestPriceDoesNotParseOpenAIVideoFields(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"instances":[{"prompt":"A city street"}],
		"parameters":{"durationSeconds":6,"numberOfVideos":1}
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1beta/models/veo-3.1-generate-preview:predictLongRunning",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	price, err := GetGeminiVideoRequestPrice(ctx, model.ModelConfig{
		AllowedResolutions: []string{"1080p"},
		Price: model.Price{
			OutputPrice: 0.2,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{Resolution: []string{"1080p"}},
					Price:     model.Price{OutputPrice: 0.5},
				},
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, model.ZeroNullFloat64(0.2), price.OutputPrice)
	require.Equal(t, model.ZeroNullInt64(1), price.OutputPriceUnit)
	require.Equal(t, model.ZeroNullInt64(1), price.ConditionalPrices[0].Price.OutputPriceUnit)
}
