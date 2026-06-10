//nolint:testpackage
package openai

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/meta"
	"github.com/labring/aiproxy/core/relay/mode"
)

func TestConvertRerankRequestPatchesMultimodalContent(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/rerank",
		bytes.NewBufferString(`{
			"model":"jina-reranker-m0",
			"query":"small language model data extraction",
			"documents":[
				"plain text",
				{"type":"text","text":"reader lm"},
				{"type":"image_url","image_url":{"url":"https://example.com/a.png"}},
				{"image_url":"https://example.com/b.png"},
				{"image":"data:image/png;base64,abc"}
			],
			"return_documents":false
		}`),
	)
	req.Header.Set("Content-Type", "application/json")

	m := meta.NewMeta(nil, mode.Rerank, "jina-reranker-m0", model.ModelConfig{})
	m.ActualModel = "mapped-jina-reranker-m0"

	result, err := ConvertRerankRequest(m, req)
	if err != nil {
		t.Fatalf("ConvertRerankRequest returned error: %v", err)
	}

	body := readRerankConvertResultBody(t, result.Body)
	if body["model"] != "mapped-jina-reranker-m0" {
		t.Fatalf("model was not rewritten: %#v", body["model"])
	}

	documents, ok := body["documents"].([]any)
	if !ok || len(documents) != 5 {
		t.Fatalf("unexpected documents: %#v", body["documents"])
	}

	if documents[0] != "plain text" {
		t.Fatalf("string document should be preserved: %#v", documents[0])
	}

	assertRerankObject(t, documents[1], map[string]string{"text": "reader lm"})
	assertRerankObject(t, documents[2], map[string]string{"image": "https://example.com/a.png"})
	assertRerankObject(t, documents[3], map[string]string{"image": "https://example.com/b.png"})
	assertRerankObject(t, documents[4], map[string]string{"image": "data:image/png;base64,abc"})
}

func TestConvertRerankRequestPatchesQueryImageURL(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/rerank",
		bytes.NewBufferString(`{
			"model":"jina-reranker-m0",
			"query":{"type":"image_url","image_url":{"url":"https://example.com/query.png"}},
			"documents":["plain text"]
		}`),
	)
	req.Header.Set("Content-Type", "application/json")

	m := meta.NewMeta(nil, mode.Rerank, "jina-reranker-m0", model.ModelConfig{})
	m.ActualModel = "mapped-jina-reranker-m0"

	result, err := ConvertRerankRequest(m, req)
	if err != nil {
		t.Fatalf("ConvertRerankRequest returned error: %v", err)
	}

	body := readRerankConvertResultBody(t, result.Body)
	assertRerankObject(t, body["query"], map[string]string{
		"image": "https://example.com/query.png",
	})
}

func readRerankConvertResultBody(t *testing.T, body io.Reader) map[string]any {
	t.Helper()

	data, err := io.ReadAll(body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}

	var got map[string]any
	if err := sonic.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal body %s: %v", string(data), err)
	}

	return got
}

func assertRerankObject(t *testing.T, got any, want map[string]string) {
	t.Helper()

	gotMap, ok := got.(map[string]any)
	if !ok {
		t.Fatalf("expected object, got %#v", got)
	}

	if len(gotMap) != len(want) {
		t.Fatalf("unexpected object keys: got %#v want %#v", gotMap, want)
	}

	for key, wantValue := range want {
		if gotMap[key] != wantValue {
			t.Fatalf("unexpected %s: got %#v want %q", key, gotMap[key], wantValue)
		}
	}
}
