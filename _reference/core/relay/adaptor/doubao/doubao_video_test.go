//nolint:testpackage
package doubao

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

func boolValue(value *bool) bool {
	return value != nil && *value
}

func TestConvertNativeVideoRequestPreservesBodyAndRewritesModel(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/api/v3/contents/generations/tasks",
		bytes.NewBufferString(
			`{"model":"doubao-seedance-2-0-260128","content":[{"type":"text","text":"go"},{"type":"video_url","video_url":{"url":"https://example.com/in.mp4"}}],"resolution":"720p","generate_audio":false}`,
		),
	)
	req.Header.Set("Content-Type", "application/json")

	m := meta.NewMeta(nil, mode.DoubaoVideo, "doubao-seedance-2-0-260128", coremodel.ModelConfig{})
	m.ActualModel = "mapped-seedance"

	result, err := ConvertDoubaoNativeVideoRequest(m, req)
	if err != nil {
		t.Fatalf("ConvertDoubaoNativeVideoRequest returned error: %v", err)
	}

	var body map[string]any
	if err := json.NewDecoder(result.Body).Decode(&body); err != nil {
		t.Fatalf("decode converted body: %v", err)
	}

	if body["model"] != "mapped-seedance" {
		t.Fatalf("model was not rewritten: %#v", body["model"])
	}

	if body["resolution"] != "720p" {
		t.Fatalf("resolution was not preserved: %#v", body["resolution"])
	}

	content, ok := body["content"].([]any)
	if !ok || len(content) != 2 {
		t.Fatalf("content was not preserved: %#v", body["content"])
	}

	if usageContext := doubaoNativeVideoRequestUsageContext(m); usageContext.Resolution != "720p" ||
		usageContext.NativeResolution != "720p" ||
		boolValue(usageContext.InputVideo) != true ||
		boolValue(usageContext.OutputAudio) != false {
		t.Fatalf(
			"unexpected native request usage context: %#v input_video=%t output_audio=%t",
			usageContext,
			boolValue(usageContext.InputVideo),
			boolValue(usageContext.OutputAudio),
		)
	}
}

func TestConvertNativeVideoRequestCountsDraftTaskAsInputVideo(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/api/v3/contents/generations/tasks",
		bytes.NewBufferString(
			`{"model":"doubao-seedance-1-5-pro-251215","content":[{"type":"text","text":"go"},{"type":"draft_task","draft_task":{"id":"task-123"}}],"resolution":"720p"}`,
		),
	)
	req.Header.Set("Content-Type", "application/json")

	m := meta.NewMeta(
		nil,
		mode.DoubaoVideo,
		"doubao-seedance-1-5-pro-251215",
		coremodel.ModelConfig{},
	)

	if _, err := ConvertDoubaoNativeVideoRequest(m, req); err != nil {
		t.Fatalf("ConvertDoubaoNativeVideoRequest returned error: %v", err)
	}

	usageContext := doubaoNativeVideoRequestUsageContext(m)
	if usageContext.InputVideo == nil || !*usageContext.InputVideo {
		t.Fatalf("expected draft_task to count as input video, got %#v", usageContext)
	}
}

func TestDoubaoNativeVideoSubmitHandlerPassesThroughAndStoresTask(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/", nil)

	store := &doubaoTestStore{}
	m := meta.NewMeta(nil, mode.DoubaoVideo, "doubao-seedance-2-0-260128", coremodel.ModelConfig{})
	m.Group = coremodel.GroupCache{ID: "group-1"}
	m.Token = coremodel.TokenCache{ID: 7}
	m.Channel = meta.ChannelMeta{ID: 42}
	setDoubaoVideoMetadata(m, doubaoVideoStoreMetadata{
		Resolution:  "1080p",
		Ratio:       "16:9",
		ServiceTier: "priority",
		InputVideo:  new(true),
		OutputAudio: new(false),
	})

	respBody := `{"id":"task-123","model":"doubao-seedance-2-0-260128","status":"queued","resolution":"720p","ratio":"16:9","service_tier":"default"}`
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": {"application/json"}},
		Body:       io.NopCloser(bytes.NewBufferString(respBody)),
	}

	result, relayErr := DoubaoNativeVideoSubmitHandler(m, store, ctx, resp)
	if relayErr != nil {
		t.Fatalf("DoubaoNativeVideoSubmitHandler returned error: %v", relayErr)
	}

	if recorder.Body.String() != respBody {
		t.Fatalf("unexpected passthrough body: %s", recorder.Body.String())
	}

	if result.UpstreamID != "task-123" || !result.AsyncUsage {
		t.Fatalf("unexpected result: %#v", result)
	}

	if result.UsageContext.Resolution != "720p" ||
		result.UsageContext.NativeResolution != "720p" ||
		result.UsageContext.ServiceTier != "default" ||
		result.UsageContext.InputVideo == nil ||
		!*result.UsageContext.InputVideo ||
		result.UsageContext.OutputAudio == nil ||
		*result.UsageContext.OutputAudio {
		t.Fatalf("unexpected native usage context: %#v", result.UsageContext)
	}

	if len(store.saved) != 1 {
		t.Fatalf("expected one store save, got %d", len(store.saved))
	}

	if store.saved[0].ID != coremodel.VideoGenerationStoreID("task-123") ||
		store.saved[0].ChannelID != 42 ||
		store.saved[0].Model != "doubao-seedance-2-0-260128" {
		t.Fatalf("unexpected saved store: %#v", store.saved[0])
	}
}

func TestDoubaoNativeVideoSubmitHandlerUsesNativeRequestResolutionFallback(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/", nil)

	m := meta.NewMeta(nil, mode.DoubaoVideo, "doubao-seedance-2-0-260128", coremodel.ModelConfig{})
	setDoubaoVideoMetadata(m, doubaoVideoStoreMetadata{
		Resolution:  "1080p",
		Ratio:       "16:9",
		ServiceTier: "priority",
		InputVideo:  new(true),
		OutputAudio: new(false),
	})

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": {"application/json"}},
		Body: io.NopCloser(bytes.NewBufferString(
			`{"id":"task-123","model":"doubao-seedance-2-0-260128","status":"queued"}`,
		)),
	}

	result, relayErr := DoubaoNativeVideoSubmitHandler(m, nil, ctx, resp)
	if relayErr != nil {
		t.Fatalf("DoubaoNativeVideoSubmitHandler returned error: %v", relayErr)
	}

	if result.UsageContext.Resolution != "1080p" ||
		result.UsageContext.NativeResolution != "1080p" ||
		result.UsageContext.ServiceTier != "priority" ||
		result.UsageContext.InputVideo == nil ||
		!*result.UsageContext.InputVideo ||
		result.UsageContext.OutputAudio == nil ||
		*result.UsageContext.OutputAudio {
		t.Fatalf("unexpected native usage context fallback: %#v", result.UsageContext)
	}
}

func TestDoubaoNativeVideoSubmitHandlerRequiresID(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/", nil)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": {"application/json"}},
		Body:       io.NopCloser(bytes.NewBufferString(`{"status":"queued"}`)),
	}

	_, relayErr := DoubaoNativeVideoSubmitHandler(&meta.Meta{}, nil, ctx, resp)
	if relayErr == nil {
		t.Fatal("expected missing id error")
	}
}

func TestDoubaoNativeVideoTaskHandlerBackfillsMissingIDFromMeta(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/", nil)

	store := &doubaoTestStore{}
	m := &meta.Meta{
		Mode:        mode.DoubaoVideoTasks,
		VideoID:     "task-123",
		OriginModel: "doubao-seedance-2-0-260128",
		Group:       coremodel.GroupCache{ID: "group-1"},
		Token:       coremodel.TokenCache{ID: 7},
		Channel:     meta.ChannelMeta{ID: 42},
	}
	respBody := `{"model":"doubao-seedance-2-0-260128","status":"succeeded","resolution":"720p","ratio":"16:9","content":{"video_url":"https://example.com/out.mp4"}}`
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": {"application/json"}},
		Body:       io.NopCloser(bytes.NewBufferString(respBody)),
	}

	result, relayErr := DoubaoNativeVideoTaskHandler(m, store, ctx, resp)
	if relayErr != nil {
		t.Fatalf("DoubaoNativeVideoTaskHandler returned error: %v", relayErr)
	}

	if recorder.Body.String() != respBody {
		t.Fatalf("unexpected passthrough body: %s", recorder.Body.String())
	}

	if result.UpstreamID != "task-123" {
		t.Fatalf("expected upstream id from meta, got %#v", result.UpstreamID)
	}

	if result.UsageContext.Resolution != "720p" ||
		result.UsageContext.NativeResolution != "720p" {
		t.Fatalf("unexpected native usage context: %#v", result.UsageContext)
	}

	if len(store.saved) != 1 {
		t.Fatalf("expected one store save, got %d", len(store.saved))
	}

	if store.saved[0].ID != coremodel.VideoGenerationStoreID("task-123") ||
		store.saved[0].ChannelID != 42 ||
		store.saved[0].Model != "doubao-seedance-2-0-260128" {
		t.Fatalf("unexpected saved store: %#v", store.saved[0])
	}
}
