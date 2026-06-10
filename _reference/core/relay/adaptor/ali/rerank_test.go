//nolint:testpackage
package ali

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/labring/aiproxy/core/relay/adaptor"
	"github.com/labring/aiproxy/core/relay/meta"
)

func TestConvertRerankRequestUsesSonicAST(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/rerank",
		bytes.NewBufferString(`{
			"model":"gte-rerank",
			"query":"q",
			"documents":["a","b"],
			"top_n":1,
			"return_documents":true,
			"input":{"preserve":"as_parameter"}
		}`),
	)
	req.Header.Set("Content-Type", "application/json")

	result, err := ConvertRerankRequest(&meta.Meta{ActualModel: "mapped-rerank"}, req)
	if err != nil {
		t.Fatalf("ConvertRerankRequest returned error: %v", err)
	}

	var body map[string]any
	if err := sonic.ConfigStd.NewDecoder(result.Body).Decode(&body); err != nil {
		t.Fatalf("decode converted body: %v", err)
	}

	if body["model"] != "mapped-rerank" {
		t.Fatalf("model was not rewritten: %#v", body["model"])
	}

	if _, ok := body["query"]; ok {
		t.Fatalf("query should be moved to input: %#v", body)
	}

	if _, ok := body["documents"]; ok {
		t.Fatalf("documents should be moved to input: %#v", body)
	}

	input, ok := body["input"].(map[string]any)
	if !ok {
		t.Fatalf("input was not created: %#v", body["input"])
	}

	if input["query"] != "q" {
		t.Fatalf("unexpected input.query: %#v", input["query"])
	}

	documents, ok := input["documents"].([]any)
	if !ok || len(documents) != 2 || documents[0] != "a" || documents[1] != "b" {
		t.Fatalf("unexpected input.documents: %#v", input["documents"])
	}

	parameters, ok := body["parameters"].(map[string]any)
	if !ok {
		t.Fatalf("parameters was not created: %#v", body["parameters"])
	}

	if parameters["top_n"] != float64(1) || parameters["return_documents"] != true {
		t.Fatalf("unexpected parameters: %#v", parameters)
	}

	if _, ok := parameters["input"]; ok {
		t.Fatalf("existing input should be replaced by generated input: %#v", parameters["input"])
	}
}

func TestConvertRerankRequestSupportsMultimodalContent(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/rerank",
		bytes.NewBufferString(`{
			"model":"qwen3-vl-rerank",
			"query":{"type":"image_url","image_url":{"url":"https://example.com/query.png"}},
			"documents":[
				"plain text",
				{"text":"product a"},
				{"type":"image","image_url":{"url":"https://example.com/a.png"}},
				{"image_url":"https://example.com/b.png"},
				{"type":"video_url","video_url":{"url":"https://example.com/c.mp4"}},
				{"video":"https://example.com/d.mp4"}
			],
			"top_n":1,
			"return_documents":true,
			"fps":2.0
		}`),
	)
	req.Header.Set("Content-Type", "application/json")

	result, err := ConvertRerankRequest(&meta.Meta{ActualModel: "mapped-qwen3-vl-rerank"}, req)
	if err != nil {
		t.Fatalf("ConvertRerankRequest returned error: %v", err)
	}

	var body map[string]any
	if err := sonic.ConfigStd.NewDecoder(result.Body).Decode(&body); err != nil {
		t.Fatalf("decode converted body: %v", err)
	}

	if body["model"] != "mapped-qwen3-vl-rerank" {
		t.Fatalf("model was not rewritten: %#v", body["model"])
	}

	input, ok := body["input"].(map[string]any)
	if !ok {
		t.Fatalf("input was not created: %#v", body["input"])
	}

	query, ok := input["query"].(map[string]any)
	if !ok || query["image"] != "https://example.com/query.png" {
		t.Fatalf("unexpected input.query: %#v", input["query"])
	}

	documents, ok := input["documents"].([]any)
	if !ok || len(documents) != 6 {
		t.Fatalf("unexpected input.documents: %#v", input["documents"])
	}

	if documents[0] != "plain text" {
		t.Fatalf("unexpected string document: %#v", documents[0])
	}

	assertRerankContent(t, documents[1], "text", "product a")
	assertRerankContent(t, documents[2], "image", "https://example.com/a.png")
	assertRerankContent(t, documents[3], "image", "https://example.com/b.png")
	assertRerankContent(t, documents[4], "video", "https://example.com/c.mp4")
	assertRerankContent(t, documents[5], "video", "https://example.com/d.mp4")

	for i, document := range documents {
		content, ok := document.(map[string]any)
		if !ok {
			continue
		}

		if _, ok := content["type"]; ok {
			t.Fatalf("document %d should not send type upstream: %#v", i, document)
		}
	}

	parameters, ok := body["parameters"].(map[string]any)
	if !ok {
		t.Fatalf("parameters were not created: %#v", body["parameters"])
	}

	if parameters["fps"] != float64(2) ||
		parameters["top_n"] != float64(1) ||
		parameters["return_documents"] != true {
		t.Fatalf("unexpected parameters: %#v", body["parameters"])
	}
}

func TestConvertRerankRequestAddsAliImageDataURLPrefix(t *testing.T) {
	t.Parallel()

	const pngBase64 = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z8BQDwAEhQGAhKmMIQAAAABJRU5ErkJggg=="

	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/rerank",
		bytes.NewBufferString(`{
			"model":"qwen3-vl-rerank",
			"query":"find image",
			"documents":[{"image":"`+pngBase64+`"}]
		}`),
	)
	req.Header.Set("Content-Type", "application/json")

	result, err := ConvertRerankRequest(&meta.Meta{ActualModel: "qwen3-vl-rerank"}, req)
	if err != nil {
		t.Fatalf("ConvertRerankRequest returned error: %v", err)
	}

	var body map[string]any
	if err := sonic.ConfigStd.NewDecoder(result.Body).Decode(&body); err != nil {
		t.Fatalf("decode converted body: %v", err)
	}

	input, ok := body["input"].(map[string]any)
	if !ok {
		t.Fatalf("input was not created: %#v", body["input"])
	}

	documents, ok := input["documents"].([]any)
	if !ok || len(documents) != 1 {
		t.Fatalf("unexpected input.documents: %#v", input["documents"])
	}

	assertRerankContent(t, documents[0], "image", "data:image/png;base64,"+pngBase64)
}

func TestConvertRerankRequestRejectsInvalidAliImageBase64(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/rerank",
		bytes.NewBufferString(`{
			"model":"qwen3-vl-rerank",
			"query":"find image",
			"documents":[{"image":"not-base64"}]
		}`),
	)
	req.Header.Set("Content-Type", "application/json")

	_, err := ConvertRerankRequest(&meta.Meta{ActualModel: "qwen3-vl-rerank"}, req)
	if err == nil {
		t.Fatal("expected invalid image error")
	}

	var relayErr adaptor.Error

	ok := errors.As(err, &relayErr)
	if !ok {
		t.Fatalf("expected adaptor.Error, got %T: %v", err, err)
	}

	if relayErr.StatusCode() != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", relayErr.StatusCode())
	}
}

func assertRerankContent(t *testing.T, item any, key, value string) {
	t.Helper()

	content, ok := item.(map[string]any)
	if !ok {
		t.Fatalf("content is not an object: %#v", item)
	}

	if content[key] != value {
		t.Fatalf("unexpected %s content: %#v", key, item)
	}
}
