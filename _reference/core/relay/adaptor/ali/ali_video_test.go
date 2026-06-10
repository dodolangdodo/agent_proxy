//nolint:testpackage
package ali

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	coremodel "github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/meta"
	"github.com/labring/aiproxy/core/relay/mode"
)

func TestConvertAliNativeVideoRequestPreservesBodyAndRewritesModel(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/api/v1/services/aigc/video-generation/video-synthesis",
		bytes.NewBufferString(
			`{"model":"wan2.5-t2v-preview","input":{"prompt":"go"},"parameters":{"duration":5,"size":"720P"}}`,
		),
	)
	req.Header.Set("Content-Type", "application/json")

	m := meta.NewMeta(nil, mode.AliVideo, "wan2.5-t2v-preview", coremodel.ModelConfig{})
	m.ActualModel = "mapped-wan"

	result, err := ConvertAliNativeVideoRequest(m, req)
	if err != nil {
		t.Fatalf("ConvertAliNativeVideoRequest returned error: %v", err)
	}

	if got := result.Header.Get("X-Dashscope-Async"); got != "enable" {
		t.Fatalf("expected async header, got %q", got)
	}

	var body map[string]any
	if err := json.NewDecoder(result.Body).Decode(&body); err != nil {
		t.Fatalf("decode converted body: %v", err)
	}

	if body["model"] != "mapped-wan" {
		t.Fatalf("model was not rewritten: %#v", body["model"])
	}

	input, ok := body["input"].(map[string]any)
	if !ok || input["prompt"] != "go" {
		t.Fatalf("input was not preserved: %#v", body["input"])
	}

	parameters, ok := body["parameters"].(map[string]any)
	if !ok || parameters["size"] != "720P" {
		t.Fatalf("parameters were not preserved: %#v", body["parameters"])
	}

	if got := aliNativeVideoRequestUsageContext(m).Resolution; got != "720P" {
		t.Fatalf("expected native request usage context resolution, got %q", got)
	}

	if got := m.GetInt(metaAliVideoSeconds); got != 5 {
		t.Fatalf("expected native request duration metadata, got %d", got)
	}
}

func TestAliNativeVideoHandlerPassesThroughAndStoresTask(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/", nil)

	store := &aliTestStore{}
	m := meta.NewMeta(nil, mode.AliVideo, "wan2.5-t2v-preview", coremodel.ModelConfig{})
	m.Group = coremodel.GroupCache{ID: "group-1"}
	m.Token = coremodel.TokenCache{ID: 7}
	m.Channel = meta.ChannelMeta{ID: 42}
	m.Set(metaAliVideoSize, "1080P")
	m.Set(metaAliVideoSeconds, 6)

	respBody := `{"request_id":"req-1","output":{"task_id":"task-123","task_status":"PENDING"},"usage":{"SR":720,"ratio":"16:9"}}`
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": {"application/json"}},
		Body:       io.NopCloser(bytes.NewBufferString(respBody)),
	}

	result, relayErr := AliNativeVideoHandler(m, store, ctx, resp)
	if relayErr != nil {
		t.Fatalf("AliNativeVideoHandler returned error: %v", relayErr)
	}

	if recorder.Body.String() != respBody {
		t.Fatalf("unexpected passthrough body: %s", recorder.Body.String())
	}

	if result.UpstreamID != "task-123" || !result.AsyncUsage {
		t.Fatalf("unexpected result: %#v", result)
	}

	if result.UsageContext.Resolution != "720P" ||
		result.UsageContext.NativeResolution != "720P" {
		t.Fatalf("unexpected native usage context: %#v", result.UsageContext)
	}

	if len(store.saved) != 1 {
		t.Fatalf("expected one store save, got %d", len(store.saved))
	}

	if store.saved[0].ID != coremodel.VideoGenerationStoreID("task-123") ||
		store.saved[0].ChannelID != 42 ||
		store.saved[0].Model != "wan2.5-t2v-preview" {
		t.Fatalf("unexpected saved store: %#v", store.saved[0])
	}

	metadata, err := parseAliVideoStoreMetadata(store.saved[0].Metadata)
	if err != nil {
		t.Fatalf("parse saved metadata: %v", err)
	}

	if metadata.Size != "1080P" || metadata.Seconds != 6 || metadata.UpstreamID != "task-123" {
		t.Fatalf("unexpected saved metadata: %#v", metadata)
	}
}

func TestAliNativeVideoHandlerUsesNativeRequestResolutionFallback(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/", nil)

	m := meta.NewMeta(nil, mode.AliVideo, "wan2.5-t2v-preview", coremodel.ModelConfig{})
	m.Set(metaAliVideoSize, "1080P")

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": {"application/json"}},
		Body: io.NopCloser(bytes.NewBufferString(
			`{"request_id":"req-1","output":{"task_id":"task-123","task_status":"PENDING"}}`,
		)),
	}

	result, relayErr := AliNativeVideoHandler(m, nil, ctx, resp)
	if relayErr != nil {
		t.Fatalf("AliNativeVideoHandler returned error: %v", relayErr)
	}

	if result.UsageContext.Resolution != "1080P" ||
		result.UsageContext.NativeResolution != "1080P" {
		t.Fatalf("unexpected native usage context fallback: %#v", result.UsageContext)
	}
}
