package jina_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/bytedance/sonic"
	coremodel "github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/adaptor/jina"
	"github.com/labring/aiproxy/core/relay/meta"
	"github.com/labring/aiproxy/core/relay/mode"
)

func TestConvertEmbeddingsRequestPatchesOpenAIMultimodalInput(t *testing.T) {
	body := []byte(`{
		"model": "jina-embeddings-v4",
		"input": [
			{
				"type": "image_url",
				"image_url": {
					"url": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUg"
				}
			},
			{
				"type": "text",
				"text": "请描述这张图片"
			},
			"plain text"
		],
		"encoding_format": "float"
	}`)
	req := newEmbeddingsRequest(t, body)
	m := meta.NewMeta(nil, mode.Embeddings, "jina-embeddings-v4", coremodel.ModelConfig{})

	result, err := jina.ConvertEmbeddingsRequest(m, req)
	if err != nil {
		t.Fatalf("ConvertEmbeddingsRequest returned error: %v", err)
	}

	got := readConvertResultBody(t, result.Body)

	if _, ok := got["encoding_format"]; ok {
		t.Fatal("expected encoding_format to be removed")
	}

	input, ok := got["input"].([]any)
	if !ok {
		t.Fatalf("expected input array, got %T", got["input"])
	}

	if len(input) != 3 {
		t.Fatalf("expected 3 input items, got %d", len(input))
	}

	assertMapString(t, input[0], map[string]string{
		"image": "iVBORw0KGgoAAAANSUhEUg",
	})
	assertMapString(t, input[1], map[string]string{
		"text": "请描述这张图片",
	})
	assertMapString(t, input[2], map[string]string{
		"text": "plain text",
	})
}

func TestConvertEmbeddingsRequestPreservesJinaNativeInput(t *testing.T) {
	body := []byte(`{
		"model": "jina-embeddings-v4",
		"task": "text-matching",
		"input": [
			{"text": "A beautiful sunset over the beach"},
			{"image": "https://example.com/beach.jpg"},
			{"image": "iVBORw0KGgoAAAANSUhEUg"}
		]
	}`)
	req := newEmbeddingsRequest(t, body)
	m := meta.NewMeta(nil, mode.Embeddings, "jina-embeddings-v4", coremodel.ModelConfig{})

	result, err := jina.ConvertEmbeddingsRequest(m, req)
	if err != nil {
		t.Fatalf("ConvertEmbeddingsRequest returned error: %v", err)
	}

	got := readConvertResultBody(t, result.Body)

	input, ok := got["input"].([]any)
	if !ok {
		t.Fatalf("expected input array, got %T", got["input"])
	}

	assertMapString(t, input[0], map[string]string{
		"text": "A beautiful sunset over the beach",
	})
	assertMapString(t, input[1], map[string]string{
		"image": "https://example.com/beach.jpg",
	})
	assertMapString(t, input[2], map[string]string{
		"image": "iVBORw0KGgoAAAANSUhEUg",
	})
}

func TestConvertEmbeddingsRequestPatchesStringInput(t *testing.T) {
	body := []byte(`{
		"model": "jina-embeddings-v4",
		"input": "hello"
	}`)
	req := newEmbeddingsRequest(t, body)
	m := meta.NewMeta(nil, mode.Embeddings, "jina-embeddings-v4", coremodel.ModelConfig{})

	result, err := jina.ConvertEmbeddingsRequest(m, req)
	if err != nil {
		t.Fatalf("ConvertEmbeddingsRequest returned error: %v", err)
	}

	got := readConvertResultBody(t, result.Body)

	input, ok := got["input"].([]any)
	if !ok {
		t.Fatalf("expected input array, got %T", got["input"])
	}

	if len(input) != 1 {
		t.Fatalf("expected 1 input item, got %d", len(input))
	}

	assertMapString(t, input[0], map[string]string{
		"text": "hello",
	})
}

func newEmbeddingsRequest(t *testing.T, body []byte) *http.Request {
	t.Helper()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"/v1/embeddings",
		bytes.NewReader(body),
	)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	return req
}

func readConvertResultBody(t *testing.T, body io.Reader) map[string]any {
	t.Helper()

	data, err := io.ReadAll(body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	var got map[string]any
	if err := sonic.Unmarshal(data, &got); err != nil {
		t.Fatalf("failed to unmarshal body %s: %v", string(data), err)
	}

	return got
}

func assertMapString(t *testing.T, got any, want map[string]string) {
	t.Helper()

	gotMap, ok := got.(map[string]any)
	if !ok {
		t.Fatalf("expected map item, got %T", got)
	}

	if len(gotMap) != len(want) {
		t.Fatalf("expected %d keys, got %d: %#v", len(want), len(gotMap), gotMap)
	}

	for key, wantValue := range want {
		gotValue, ok := gotMap[key]
		if !ok {
			t.Fatalf("expected key %q in %#v", key, gotMap)
		}

		if gotValue != wantValue {
			t.Fatalf("expected %s=%q, got %#v", key, wantValue, gotValue)
		}
	}
}
