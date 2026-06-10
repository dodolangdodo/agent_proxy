//nolint:testpackage
package coze

import (
	"io"
	"net/http"
	"strings"
	"testing"

	relaymodel "github.com/labring/aiproxy/core/relay/model"
)

func TestErrorHandlerParsesCozeError(t *testing.T) {
	t.Parallel()

	resp := &http.Response{
		StatusCode: http.StatusBadRequest,
		Body: io.NopCloser(strings.NewReader(
			`{"code":4001,"msg":"invalid bot id"}`,
		)),
	}

	err := ErrorHandler(resp)
	if err.StatusCode() != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, err.StatusCode())
	}

	body, marshalErr := err.MarshalJSON()
	if marshalErr != nil {
		t.Fatalf("marshal error: %v", marshalErr)
	}

	if !strings.Contains(string(body), `"message":"invalid bot id"`) {
		t.Fatalf("expected coze message, got %s", body)
	}

	if !strings.Contains(string(body), `"code":4001`) {
		t.Fatalf("expected coze code, got %s", body)
	}

	if !strings.Contains(string(body), `"type":"`+relaymodel.ErrorTypeUpstream+`"`) {
		t.Fatalf("expected upstream type, got %s", body)
	}
}
