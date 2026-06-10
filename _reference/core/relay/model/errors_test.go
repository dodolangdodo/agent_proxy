//nolint:testpackage
package model

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/labring/aiproxy/core/relay/mode"
	"github.com/stretchr/testify/require"
)

func TestWrapperErrorWithMessageUsesGeminiErrorForNativeVideo(t *testing.T) {
	t.Parallel()

	err := WrapperErrorWithMessage(mode.GeminiVideo, http.StatusBadRequest, "bad resolution")

	var body GeminiErrorResponse
	require.NoError(t, json.Unmarshal(mustMarshalError(t, err), &body))
	require.Equal(t, "bad resolution", body.Error.Message)
	require.Equal(t, ErrorTypeAIPROXY, body.Error.Status)
	require.Equal(t, http.StatusBadRequest, body.Error.Code)
}

func TestWrapperErrorWithMessageKeepsOpenAIVideoError(t *testing.T) {
	t.Parallel()

	err := WrapperErrorWithMessage(mode.Videos, http.StatusBadRequest, "bad size")

	var body OpenAIVideoError
	require.NoError(t, json.Unmarshal(mustMarshalError(t, err), &body))
	require.Equal(t, "bad size", body.Detail)
}

func mustMarshalError(t *testing.T, err interface {
	MarshalJSON() ([]byte, error)
},
) []byte {
	t.Helper()

	data, marshalErr := err.MarshalJSON()
	require.NoError(t, marshalErr)

	return data
}
