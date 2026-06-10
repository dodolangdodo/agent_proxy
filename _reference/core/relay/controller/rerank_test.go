//nolint:testpackage
package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/model"
)

func TestGetRerankRequestUsageSupportsMultimodalContent(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/rerank",
		bytes.NewBufferString(`{
			"model":"qwen3-vl-rerank",
			"query":{"type":"text","text":"find the same product"},
			"documents":[
				"plain text",
				{"text":"product a"},
				{"type":"image_url","image_url":{"url":"https://example.com/a.png"}},
				{"type":"video_url","video_url":{"url":"https://example.com/b.mp4"}}
			],
			"top_n":1,
			"fps":2.0
		}`),
	)
	c.Request.Header.Set("Content-Type", "application/json")

	usage, err := GetRerankRequestUsage(c, model.ModelConfig{})
	if err != nil {
		t.Fatalf("GetRerankRequestUsage returned error: %v", err)
	}

	if usage.Usage.InputTokens == 0 {
		t.Fatalf("text content should be counted in request usage, got %#v", usage)
	}

	if usage.Usage.ImageInputTokens != 0 || usage.Usage.TotalTokens != 0 {
		t.Fatalf("unexpected non-text usage fields: %#v", usage)
	}
}

func TestGetRerankRequestUsageStillSupportsStringRequest(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/rerank",
		bytes.NewBufferString(`{
			"model":"gte-rerank-v2",
			"query":"what is rerank",
			"documents":["doc one","doc two"]
		}`),
	)
	c.Request.Header.Set("Content-Type", "application/json")

	usage, err := GetRerankRequestUsage(c, model.ModelConfig{})
	if err != nil {
		t.Fatalf("GetRerankRequestUsage returned error: %v", err)
	}

	if usage.Usage.InputTokens == 0 {
		t.Fatalf("string request should be counted in request usage, got %#v", usage)
	}
}

func TestGetRerankRequestUsageRejectsEmptyStringQuery(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/rerank",
		bytes.NewBufferString(`{
			"model":"gte-rerank-v2",
			"query":"",
			"documents":["doc one"]
		}`),
	)
	c.Request.Header.Set("Content-Type", "application/json")

	_, err := GetRerankRequestUsage(c, model.ModelConfig{})
	if err == nil {
		t.Fatal("GetRerankRequestUsage should reject empty string query")
	}

	if !strings.Contains(err.Error(), "query must not be empty") {
		t.Fatalf("unexpected error: %v", err)
	}
}
