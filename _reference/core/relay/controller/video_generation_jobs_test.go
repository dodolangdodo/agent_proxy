//nolint:testpackage
package controller

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/model"
	"github.com/stretchr/testify/require"
)

func TestGetVideoGenerationJobRequestUsageMultipart(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	var body bytes.Buffer

	writer := multipart.NewWriter(&body)
	require.NoError(t, writer.WriteField("model", "wan2.7-i2v"))
	require.NoError(t, writer.WriteField("prompt", "Animate the reference"))
	require.NoError(t, writer.WriteField("n_seconds", "5"))
	require.NoError(t, writer.Close())

	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/video/generations/jobs",
		&body,
	)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	usage, err := GetVideoGenerationJobRequestUsage(ctx, model.ModelConfig{})
	require.NoError(t, err)
	require.Zero(t, usage.Usage.OutputTokens)
	require.Zero(t, usage.Usage.TotalTokens)
}

func TestGetVideoGenerationJobRequestUsageJSON(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"model":"wan2.5-t2v-preview",
		"prompt":"A city street",
		"n_seconds":4,
		"n_variants":2
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/video/generations/jobs",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	usage, err := GetVideoGenerationJobRequestUsage(ctx, model.ModelConfig{})
	require.NoError(t, err)
	require.Zero(t, usage.Usage.OutputTokens)
	require.Zero(t, usage.Usage.TotalTokens)
}

func TestValidateVideoGenerationJobRequestRejectsTooManyVariants(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"model":"video-model",
		"prompt":"A city street",
		"n_seconds":6,
		"n_variants":3
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/video/generations/jobs",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	err := ValidateVideoGenerationJobRequest(ctx, model.ModelConfig{
		MaxVideoGenerationCount: 2,
	})
	require.Error(t, err)
	require.Equal(t, "video count must be less than or equal to 2", err.Error())
}

func TestValidateVideoGenerationJobRequestRejectsTooLongNSeconds(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"model":"video-model",
		"prompt":"A city street",
		"n_seconds":6
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/video/generations/jobs",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	err := ValidateVideoGenerationJobRequest(ctx, model.ModelConfig{
		MaxVideoGenerationSeconds: 5,
	})
	require.Error(t, err)
	require.Equal(t, "seconds must be less than or equal to 5", err.Error())

	var requestParamErr *RequestParamError
	require.ErrorAs(t, err, &requestParamErr)
	require.Equal(t, 400, requestParamErr.StatusCode)
}

func TestGetVideoGenerationJobRequestUsageSetsResolutionCondition(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"model":"video-model",
		"prompt":"A city street",
		"width":1280,
		"height":720,
		"n_seconds":5
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/video/generations/jobs",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	usage, err := GetVideoGenerationJobRequestUsage(ctx, model.ModelConfig{})
	require.NoError(t, err)
	require.Zero(t, usage.Usage.OutputTokens)
	require.Equal(t, "720p", usage.Context.Resolution)
}

func TestGetVideoGenerationJobRequestUsageRejectsUnsupportedResolution(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"model":"video-model",
		"prompt":"A city street",
		"width":1280,
		"height":720,
		"n_seconds":5
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/video/generations/jobs",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	_, err := GetVideoGenerationJobRequestUsage(ctx, model.ModelConfig{
		AllowedResolutions: []string{"1080p"},
	})
	require.Error(t, err)
	require.Equal(
		t,
		"unsupported video resolution `720p`, supported resolutions: 1920x1080",
		err.Error(),
	)
}

func TestGetVideoGenerationJobRequestUsageIgnoresSizeField(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"model":"video-model",
		"prompt":"A city street",
		"size":"720P",
		"n_seconds":5
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/video/generations/jobs",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	usage, err := GetVideoGenerationJobRequestUsage(ctx, model.ModelConfig{
		AllowedResolutions: []string{"720p"},
	})
	require.NoError(t, err)
	require.Empty(t, usage.Context.Resolution)
}

func TestGetVideoGenerationJobRequestUsageRejectsUnpairedWidthHeight(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"model":"video-model",
		"prompt":"A city street",
		"width":1280,
		"n_seconds":5
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/video/generations/jobs",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	_, err := GetVideoGenerationJobRequestUsage(ctx, model.ModelConfig{})
	require.Error(t, err)
	require.Equal(t, "width and height must be provided together", err.Error())
}

func TestGetVideoGenerationJobRequestUsageMatchesDimensionToTier(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"model":"video-model",
		"prompt":"A city street",
		"width":1280,
		"height":720,
		"n_seconds":5
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/video/generations/jobs",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	_, err := GetVideoGenerationJobRequestUsage(ctx, model.ModelConfig{
		AllowedResolutions: []string{"720p"},
	})
	require.NoError(t, err)
}

func TestGetVideoGenerationJobRequestUsageRejectsAllBlankAllowedResolutions(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"model":"video-model",
		"prompt":"A city street",
		"width":1280,
		"height":720,
		"n_seconds":5
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/video/generations/jobs",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	_, err := GetVideoGenerationJobRequestUsage(ctx, model.ModelConfig{
		AllowedResolutions: []string{" ", "\t"},
	})
	require.Error(t, err)
	require.Equal(
		t,
		"unsupported video resolution `720p`, supported resolutions: none",
		err.Error(),
	)
}

func TestGetVideoGenerationJobRequestUsageMatchesPortraitDimensionToTier(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"model":"video-model",
		"prompt":"A city street",
		"width":720,
		"height":1280,
		"n_seconds":5
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/video/generations/jobs",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	_, err := GetVideoGenerationJobRequestUsage(ctx, model.ModelConfig{
		AllowedResolutions: []string{"720p"},
	})
	require.NoError(t, err)
}

func TestGetVideoGenerationJobRequestUsageResolutionMatrix(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		width      int
		height     int
		allowed    []string
		wantCtx    string
		wantErr    bool
		wantErrMsg string
	}{
		{
			name:    "landscape dimensions match 720p tier",
			width:   1280,
			height:  720,
			allowed: []string{"720p"},
			wantCtx: "720p",
		},
		{
			name:    "portrait dimensions match 720p tier",
			width:   720,
			height:  1280,
			allowed: []string{"720p"},
			wantCtx: "720p",
		},
		{
			name:    "dimensions match exact allowed dimension by fuzzy tier",
			width:   1280,
			height:  720,
			allowed: []string{"1280x720"},
			wantCtx: "720p",
		},
		{
			name:       "width and height must be paired",
			width:      1280,
			allowed:    []string{"720p"},
			wantErr:    true,
			wantErrMsg: "width and height must be provided together",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			body := `{
				"model":"video-model",
				"prompt":"A city street",
				"width":` + strconv.Itoa(tt.width) + `,
				"height":` + strconv.Itoa(tt.height) + `,
				"n_seconds":5
			}`
			req := httptest.NewRequestWithContext(
				t.Context(),
				http.MethodPost,
				"/v1/video/generations/jobs",
				bytes.NewBufferString(body),
			)
			req.Header.Set("Content-Type", "application/json")

			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			ctx.Request = req

			usage, err := GetVideoGenerationJobRequestUsage(ctx, model.ModelConfig{
				AllowedResolutions: tt.allowed,
			})
			if tt.wantErr {
				require.Error(t, err)
				require.Equal(t, tt.wantErrMsg, err.Error())
				return
			}

			require.NoError(t, err)
			require.Zero(t, usage.Usage.OutputTokens)
			require.Equal(t, tt.wantCtx, usage.Context.Resolution)
		})
	}
}

func TestGetVideoGenerationJobRequestUsageDisableFuzzyRejectsDimensionToTier(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"model":"video-model",
		"prompt":"A city street",
		"width":1280,
		"height":720,
		"n_seconds":5
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/video/generations/jobs",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	_, err := GetVideoGenerationJobRequestUsage(ctx, model.ModelConfig{
		AllowedResolutions:          []string{"720p"},
		DisableResolutionFuzzyMatch: true,
	})
	require.NoError(t, err)
}

func TestGetVideoGenerationJobRequestUsageRejectsTooLongSeconds(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"model":"video-model",
		"prompt":"A city street",
		"n_seconds":6
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/video/generations/jobs",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	_, err := GetVideoGenerationJobRequestUsage(ctx, model.ModelConfig{
		MaxVideoGenerationSeconds: 5,
	})
	require.Error(t, err)
	require.Equal(t, "seconds must be less than or equal to 5", err.Error())

	var requestParamErr *RequestParamError
	require.ErrorAs(t, err, &requestParamErr)
	require.Equal(t, 400, requestParamErr.StatusCode)
}

func TestGetVideoGenerationJobRequestUsageRejectsNegativeNSeconds(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"model":"video-model",
		"prompt":"A city street",
		"n_seconds":-1,
		"seconds":60
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/video/generations/jobs",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	_, err := GetVideoGenerationJobRequestUsage(ctx, model.ModelConfig{})
	require.Error(t, err)
	require.Equal(t, "invalid n_seconds: must be non-negative", err.Error())

	var requestParamErr *RequestParamError
	require.ErrorAs(t, err, &requestParamErr)
	require.Equal(t, 400, requestParamErr.StatusCode)
}

func TestGetVideoGenerationJobRequestUsageIgnoresVideosSecondsField(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"model":"video-model",
		"prompt":"A city street",
		"seconds":-1
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/video/generations/jobs",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	usage, err := GetVideoGenerationJobRequestUsage(ctx, model.ModelConfig{})
	require.NoError(t, err)
	require.Zero(t, usage.Usage.OutputTokens)
	require.Zero(t, usage.Usage.TotalTokens)
}

func TestGetVideoGenerationJobRequestPriceSetsPerSecondUnitForConditionalPrices(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	var body bytes.Buffer

	writer := multipart.NewWriter(&body)
	require.NoError(t, writer.WriteField("model", "video-model"))
	require.NoError(t, writer.WriteField("prompt", "A city street"))
	require.NoError(t, writer.WriteField("width", "1280"))
	require.NoError(t, writer.WriteField("height", "720"))
	require.NoError(t, writer.Close())

	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/video/generations/jobs",
		&body,
	)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	price, err := GetVideoGenerationJobRequestPrice(ctx, model.ModelConfig{
		Price: model.Price{
			OutputPrice: 0.2,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{Resolution: []string{"1280x720"}},
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
