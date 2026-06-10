//nolint:testpackage
package cohere

import (
	"io"
	"net/http"
	"strings"
	"testing"

	relaymodel "github.com/labring/aiproxy/core/relay/model"
)

func TestErrorHandlerParsesCohereError(t *testing.T) {
	t.Parallel()

	resp := &http.Response{
		StatusCode: http.StatusTooManyRequests,
		Body: io.NopCloser(strings.NewReader(
			`{"message":"too many requests"}`,
		)),
	}

	err := ErrorHandler(resp)
	if err.StatusCode() != http.StatusTooManyRequests {
		t.Fatalf("expected status %d, got %d", http.StatusTooManyRequests, err.StatusCode())
	}

	body, marshalErr := err.MarshalJSON()
	if marshalErr != nil {
		t.Fatalf("marshal error: %v", marshalErr)
	}

	if !strings.Contains(string(body), `"message":"too many requests"`) {
		t.Fatalf("expected cohere message, got %s", body)
	}

	if !strings.Contains(string(body), `"type":"`+relaymodel.ErrorTypeUpstream+`"`) {
		t.Fatalf("expected upstream type, got %s", body)
	}
}
