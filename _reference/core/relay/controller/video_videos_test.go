//nolint:testpackage
package controller

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/model"
	"github.com/stretchr/testify/require"
)

func TestGetVideosRequestUsageMultipart(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	var body bytes.Buffer

	writer := multipart.NewWriter(&body)
	require.NoError(t, writer.WriteField("model", "sora-2"))
	require.NoError(t, writer.WriteField("prompt", "Animate the reference"))
	require.NoError(t, writer.WriteField("seconds", "6"))
	require.NoError(t, writer.Close())

	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/videos",
		&body,
	)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	usage, err := GetVideosRequestUsage(ctx, model.ModelConfig{})
	require.NoError(t, err)
	require.Zero(t, usage.Usage.OutputTokens)
	require.Zero(t, usage.Usage.TotalTokens)
}

func TestValidateVideosRequestRejectsTooLongSeconds(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"model":"video-model",
		"prompt":"A city street",
		"seconds":6
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/videos",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	err := ValidateVideosRequest(ctx, model.ModelConfig{
		MaxVideoGenerationSeconds: 5,
	})
	require.Error(t, err)
	require.Equal(t, "seconds must be less than or equal to 5", err.Error())

	var requestParamErr *RequestParamError
	require.ErrorAs(t, err, &requestParamErr)
	require.Equal(t, 400, requestParamErr.StatusCode)
}

func TestValidateVideosRequestRejectsUnsupportedMultipartSize(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	var body bytes.Buffer

	writer := multipart.NewWriter(&body)
	require.NoError(t, writer.WriteField("model", "video-model"))
	require.NoError(t, writer.WriteField("prompt", "A city street"))
	require.NoError(t, writer.WriteField("size", "1920x1080"))
	require.NoError(t, writer.WriteField("seconds", "5"))
	require.NoError(t, writer.Close())

	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/videos/video-123/remix",
		&body,
	)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	err := ValidateVideosRequest(ctx, model.ModelConfig{
		AllowedResolutions: []string{"720p"},
	})
	require.Error(t, err)
	require.Equal(
		t,
		"unsupported video resolution `1920x1080`, supported resolutions: 1280x720",
		err.Error(),
	)
}

func TestValidateVideosRequestRejectsInvalidSizeFormat(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"model":"video-model",
		"prompt":"A city street",
		"size":"720P",
		"seconds":5
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/videos",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	err := ValidateVideosRequest(ctx, model.ModelConfig{
		AllowedResolutions: []string{"720p"},
	})
	require.Error(t, err)
	require.Equal(t, "invalid video size `720p`, supported resolutions: 1280x720", err.Error())
}

func TestGetVideosRequestUsageRejectsNonOpenAIDimensionDelimiter(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"model":"video-model",
		"prompt":"A city street",
		"size":"1280*720",
		"seconds":5
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/videos",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	_, err := GetVideosRequestUsage(ctx, model.ModelConfig{})
	require.Error(t, err)
	require.Equal(
		t,
		"invalid video size `1280*720`, supported resolutions: <width>x<height>",
		err.Error(),
	)
}

func TestValidateVideosRequestAllowsAdvertised4KSize(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"model":"video-model",
		"prompt":"A city street",
		"size":"3840x2160",
		"seconds":5
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/videos",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	err := ValidateVideosRequest(ctx, model.ModelConfig{
		AllowedResolutions: []string{"4k"},
	})
	require.NoError(t, err)
}

func TestGetVideosRequestUsageIgnoresJobOnlyFields(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"model":"video-model",
		"prompt":"A city street",
		"seconds":4,
		"n_seconds":60,
		"n_variants":10,
		"width":1920,
		"height":1080
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/videos",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	usage, err := GetVideosRequestUsage(ctx, model.ModelConfig{
		AllowedResolutions:        []string{"720p"},
		MaxVideoGenerationSeconds: 5,
	})
	require.NoError(t, err)
	require.Zero(t, usage.Usage.OutputTokens)
	require.Empty(t, usage.Context.Resolution)
}

func TestGetVideosRequestUsageUsesOfficialVideoFields(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	body := `{
		"model":"video-model",
		"prompt":"A city street",
		"seconds":4,
		"size":"1280x720",
		"n_seconds":60,
		"n_variants":10,
		"width":1920,
		"height":1080
	}`
	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/videos",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	usage, err := GetVideosRequestUsage(ctx, model.ModelConfig{})
	require.NoError(t, err)
	require.Zero(t, usage.Usage.OutputTokens)
	require.Zero(t, usage.Usage.TotalTokens)
	require.Equal(t, "1280x720", usage.Context.Resolution)
}
