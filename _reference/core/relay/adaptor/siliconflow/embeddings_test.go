package siliconflow_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	coremodel "github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/adaptor/siliconflow"
	"github.com/labring/aiproxy/core/relay/meta"
	"github.com/labring/aiproxy/core/relay/mode"
)

func TestConvertRequestVLEmbeddingsPatchesOpenAIImageInput(t *testing.T) {
	adaptor := &siliconflow.Adaptor{}
	m := meta.NewMeta(
		nil,
		mode.Embeddings,
		"Qwen/Qwen3-VL-Embedding-8B",
		coremodel.ModelConfig{},
	)

	req := newEmbeddingsRequest(t, []byte(`{
		"model": "Qwen/Qwen3-VL-Embedding-8B",
		"input": [
			{
				"type": "image_url",
				"image_url": {
					"url": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUg"
				}
			},
			{
				"type": "text",
				"text": "describe this image"
			},
			"plain text",
			{
				"image": "https://example.com/image.jpg"
			}
		],
		"encoding_format": "float",
		"dimensions": 768,
		"user": "user_123",
		"truncate": "right"
	}`))

	result, err := adaptor.ConvertRequest(m, nil, req)
	if err != nil {
		t.Fatalf("ConvertRequest returned error: %v", err)
	}

	got := readConvertResultBody(t, result.Body)

	if got["model"] != "Qwen/Qwen3-VL-Embedding-8B" {
		t.Fatalf("expected model to be set, got %#v", got["model"])
	}

	if got["encoding_format"] != "float" {
		t.Fatalf("expected encoding_format to be preserved, got %#v", got["encoding_format"])
	}

	if dimensions, ok := got["dimensions"].(float64); !ok || int(dimensions) != 768 {
		t.Fatalf("expected dimensions=768, got %#v", got["dimensions"])
	}

	if got["user"] != "user_123" {
		t.Fatalf("expected user to be preserved, got %#v", got["user"])
	}

	if got["truncate"] != "right" {
		t.Fatalf("expected truncate to be preserved, got %#v", got["truncate"])
	}

	input, ok := got["input"].([]any)
	if !ok {
		t.Fatalf("expected input array, got %#v", got["input"])
	}

	if len(input) != 4 {
		t.Fatalf("expected 4 input items, got %d", len(input))
	}

	assertContentString(t, input[0], "image", "data:image/png;base64,iVBORw0KGgoAAAANSUhEUg")
	assertContentString(t, input[1], "text", "describe this image")

	if input[2] != "plain text" {
		t.Fatalf("expected plain text string, got %#v", input[2])
	}

	assertContentString(t, input[3], "image", "https://example.com/image.jpg")
}

func TestConvertRequestVLEmbeddingsPatchesSingleObjectInput(t *testing.T) {
	adaptor := &siliconflow.Adaptor{}
	m := meta.NewMeta(
		nil,
		mode.Embeddings,
		"alias-model",
		coremodel.ModelConfig{
			Config: coremodel.NewModelConfig(coremodel.WithModelConfigVision(true)),
		},
	)
	m.ActualModel = "Qwen/Qwen3-VL-Embedding-8B"

	req := newEmbeddingsRequest(t, []byte(`{
		"model": "alias-model",
		"input": {
			"type": "image_url",
			"image_url": {
				"url": "https://example.com/image.jpg"
			}
		}
	}`))

	result, err := adaptor.ConvertRequest(m, nil, req)
	if err != nil {
		t.Fatalf("ConvertRequest returned error: %v", err)
	}

	got := readConvertResultBody(t, result.Body)

	if got["model"] != "Qwen/Qwen3-VL-Embedding-8B" {
		t.Fatalf("expected actual model to be set, got %#v", got["model"])
	}

	assertContentString(t, got["input"], "image", "https://example.com/image.jpg")
}

func TestConvertRequestClassicEmbeddingsUsesDefaultConversion(t *testing.T) {
	adaptor := &siliconflow.Adaptor{}
	m := meta.NewMeta(nil, mode.Embeddings, "BAAI/bge-large-zh-v1.5", coremodel.ModelConfig{})

	req := newEmbeddingsRequest(t, []byte(`{
		"model": "BAAI/bge-large-zh-v1.5",
		"input": "hello"
	}`))

	result, err := adaptor.ConvertRequest(m, nil, req)
	if err != nil {
		t.Fatalf("ConvertRequest returned error: %v", err)
	}

	got := readConvertResultBody(t, result.Body)

	if got["input"] != "hello" {
		t.Fatalf("expected classic input string to be preserved, got %#v", got["input"])
	}
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
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("failed to unmarshal body %s: %v", string(data), err)
	}

	return got
}

func assertContentString(t *testing.T, got any, key, want string) {
	t.Helper()

	gotMap, ok := got.(map[string]any)
	if !ok {
		t.Fatalf("expected content object, got %T", got)
	}

	if len(gotMap) != 1 {
		t.Fatalf("expected content to have 1 key, got %#v", gotMap)
	}

	if gotMap[key] != want {
		t.Fatalf("expected %s=%q, got %#v", key, want, gotMap[key])
	}
}
