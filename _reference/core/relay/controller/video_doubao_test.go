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

func TestValidateDoubaoVideoRequestRejectsTooLongDuration(t *testing.T) {
	t.Parallel()

	ctx := newDoubaoVideoJSONTestContext(t, `{
		"model":"doubao-seedance-2-0-260128",
		"content":[{"type":"text","text":"A city street"}],
		"duration":6,
		"resolution":"720p"
	}`)

	err := ValidateDoubaoVideoRequest(ctx, model.ModelConfig{
		MaxVideoGenerationSeconds: 5,
	})
	require.Error(t, err)
	require.Equal(t, "seconds must be less than or equal to 5", err.Error())
}

func TestValidateDoubaoVideoRequestRejectsUnsupportedResolution(t *testing.T) {
	t.Parallel()

	ctx := newDoubaoVideoJSONTestContext(t, `{
		"model":"doubao-seedance-2-0-260128",
		"content":[{"type":"text","text":"A city street"}],
		"duration":5,
		"resolution":"1080p"
	}`)

	err := ValidateDoubaoVideoRequest(ctx, model.ModelConfig{
		AllowedResolutions: []string{"720p"},
	})
	require.Error(t, err)
	require.Equal(
		t,
		"unsupported video resolution `1080p`, supported resolutions: 720p",
		err.Error(),
	)
}

func TestValidateDoubaoVideoRequestRejectsNegativeDuration(t *testing.T) {
	t.Parallel()

	ctx := newDoubaoVideoJSONTestContext(t, `{
		"model":"doubao-seedance-2-0-260128",
		"content":[{"type":"text","text":"A city street"}],
		"duration":-1,
		"resolution":"720p"
	}`)

	err := ValidateDoubaoVideoRequest(ctx, model.ModelConfig{})
	require.Error(t, err)
	require.Equal(t, "invalid duration: must be non-negative", err.Error())
}

func TestGetDoubaoVideoRequestUsageUsesDoubaoResolution(t *testing.T) {
	t.Parallel()

	ctx := newDoubaoVideoJSONTestContext(t, `{
		"model":"doubao-seedance-2-0-260128",
		"content":[{"type":"text","text":"A city street"}],
		"duration":5,
		"resolution":"720p"
	}`)

	usage, err := GetDoubaoVideoRequestUsage(ctx, model.ModelConfig{
		AllowedResolutions: []string{"720p"},
	})
	require.NoError(t, err)
	require.Zero(t, usage.Usage.OutputTokens)
	require.Equal(t, "720p", usage.Context.Resolution)
	require.Equal(t, "720p", usage.Context.NativeResolution)
}

func newDoubaoVideoJSONTestContext(t *testing.T, body string) *gin.Context {
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
