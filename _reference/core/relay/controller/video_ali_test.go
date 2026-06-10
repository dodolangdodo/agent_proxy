//nolint:testpackage
package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/model"
	"github.com/stretchr/testify/require"
)

func TestValidateAliVideoRequestRejectsTooLongDuration(t *testing.T) {
	t.Parallel()

	ctx := newAliVideoJSONTestContext(t, `{
		"model":"wan2.5-t2v-preview",
		"input":{"prompt":"A city street"},
		"parameters":{"duration":6,"size":"720P"}
	}`)

	err := ValidateAliVideoRequest(ctx, model.ModelConfig{
		MaxVideoGenerationSeconds: 5,
	})
	require.Error(t, err)
	require.Equal(t, "seconds must be less than or equal to 5", err.Error())

	var requestParamErr *RequestParamError
	require.ErrorAs(t, err, &requestParamErr)
	require.Equal(t, 400, requestParamErr.StatusCode)
}

func TestValidateAliVideoRequestRejectsUnsupportedResolution(t *testing.T) {
	t.Parallel()

	ctx := newAliVideoJSONTestContext(t, `{
		"model":"wan2.5-t2v-preview",
		"input":{"prompt":"A city street"},
		"parameters":{"duration":5,"size":"1080P"}
	}`)

	err := ValidateAliVideoRequest(ctx, model.ModelConfig{
		AllowedResolutions: []string{"720p"},
	})
	require.Error(t, err)
	require.Equal(
		t,
		"unsupported video resolution `1080P`, supported resolutions: 720p",
		err.Error(),
	)
}

func TestValidateAliVideoRequestRejectsNegativeDuration(t *testing.T) {
	t.Parallel()

	ctx := newAliVideoJSONTestContext(t, `{
		"model":"wan2.5-t2v-preview",
		"input":{"prompt":"A city street"},
		"parameters":{"duration":-1,"size":"720P"}
	}`)

	err := ValidateAliVideoRequest(ctx, model.ModelConfig{})
	require.Error(t, err)
	require.Equal(t, "invalid duration: must be non-negative", err.Error())
}

func TestGetAliVideoRequestUsageUsesAliResolution(t *testing.T) {
	t.Parallel()

	ctx := newAliVideoJSONTestContext(t, `{
		"model":"wan2.5-t2v-preview",
		"input":{"prompt":"A city street"},
		"parameters":{"duration":5,"size":"720P"}
	}`)

	usage, err := GetAliVideoRequestUsage(ctx, model.ModelConfig{
		AllowedResolutions: []string{"720p"},
	})
	require.NoError(t, err)
	require.Zero(t, usage.Usage.OutputTokens)
	require.Equal(t, "720P", usage.Context.Resolution)
	require.Equal(t, "720P", usage.Context.NativeResolution)
}

func newAliVideoJSONTestContext(t *testing.T, body string) *gin.Context {
	t.Helper()

	gin.SetMode(gin.TestMode)

	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	return ctx
}
