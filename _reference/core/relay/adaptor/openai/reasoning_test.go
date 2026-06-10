//nolint:testpackage
package openai

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/meta"
	"github.com/labring/aiproxy/core/relay/mode"
	relaymodel "github.com/labring/aiproxy/core/relay/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpenAIReasoningEffortForModel(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		originModel string
		actualModel string
		effort      relaymodel.ReasoningEffort
		want        relaymodel.ReasoningEffort
	}{
		{
			name:        "gpt-5.5 minimal migrates to low",
			originModel: "gpt-5.5",
			effort:      relaymodel.ReasoningEffortMinimal,
			want:        relaymodel.ReasoningEffortLow,
		},
		{
			name:        "gpt-5.4 supports xhigh",
			originModel: "gpt-5.4",
			effort:      relaymodel.ReasoningEffortXHigh,
			want:        relaymodel.ReasoningEffortXHigh,
		},
		{
			name:        "gpt-5.2 supports xhigh",
			originModel: "gpt-5.2",
			effort:      relaymodel.ReasoningEffortXHigh,
			want:        relaymodel.ReasoningEffortXHigh,
		},
		{
			name:        "gpt-5.4-pro minimal migrates to medium",
			originModel: "gpt-5.4-pro",
			effort:      relaymodel.ReasoningEffortMinimal,
			want:        relaymodel.ReasoningEffortMedium,
		},
		{
			name:        "gpt-5.2-pro none migrates to medium",
			originModel: "gpt-5.2-pro",
			effort:      relaymodel.ReasoningEffortNone,
			want:        relaymodel.ReasoningEffortMedium,
		},
		{
			name:        "gpt-5 xhigh migrates to high",
			originModel: "gpt-5",
			effort:      relaymodel.ReasoningEffortXHigh,
			want:        relaymodel.ReasoningEffortHigh,
		},
		{
			name:        "gpt-5 none migrates to minimal",
			originModel: "gpt-5",
			effort:      relaymodel.ReasoningEffortNone,
			want:        relaymodel.ReasoningEffortMinimal,
		},
		{
			name:        "gpt-5.1 xhigh migrates to high",
			originModel: "gpt-5.1",
			effort:      relaymodel.ReasoningEffortXHigh,
			want:        relaymodel.ReasoningEffortHigh,
		},
		{
			name:        "unknown codex model preserves xhigh",
			originModel: "gpt-5.1-codex-max",
			effort:      relaymodel.ReasoningEffortXHigh,
			want:        relaymodel.ReasoningEffortXHigh,
		},
		{
			name:        "gpt-5-pro only supports high",
			originModel: "gpt-5-pro",
			effort:      relaymodel.ReasoningEffortLow,
			want:        relaymodel.ReasoningEffortHigh,
		},
		{
			name:        "unknown model preserves effort",
			originModel: "custom-gpt",
			effort:      relaymodel.ReasoningEffortXHigh,
			want:        relaymodel.ReasoningEffortXHigh,
		},
		{
			name:        "unknown gpt-like model preserves none",
			originModel: "my-gpt-5-like-model",
			effort:      relaymodel.ReasoningEffortNone,
			want:        relaymodel.ReasoningEffortNone,
		},
		{
			name:        "unknown future gpt model preserves unsupported effort",
			originModel: "gpt-5.9",
			effort:      relaymodel.ReasoningEffortMinimal,
			want:        relaymodel.ReasoningEffortMinimal,
		},
		{
			name:        "actual model fallback applies when origin does not match",
			originModel: "customer-alias",
			actualModel: "gpt-5.5",
			effort:      relaymodel.ReasoningEffortMinimal,
			want:        relaymodel.ReasoningEffortLow,
		},
		{
			name:        "origin model wins when both match",
			originModel: "gpt-5",
			actualModel: "gpt-5.5",
			effort:      relaymodel.ReasoningEffortNone,
			want:        relaymodel.ReasoningEffortMinimal,
		},
		{
			name:        "series keyword matching handles dated snapshot names",
			originModel: "openai/gpt-5.5-2026-05-14",
			effort:      relaymodel.ReasoningEffortMinimal,
			want:        relaymodel.ReasoningEffortLow,
		},
		{
			name:        "series keyword matching handles mini dated snapshot names",
			originModel: "gpt-5.4-mini-2026-03-17",
			effort:      relaymodel.ReasoningEffortMinimal,
			want:        relaymodel.ReasoningEffortLow,
		},
		{
			name:        "series keyword matching handles nano dated snapshot names",
			originModel: "gpt-5.4-nano-2026-03-17",
			effort:      relaymodel.ReasoningEffortMinimal,
			want:        relaymodel.ReasoningEffortLow,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := openAIReasoningEffortForModel(tt.originModel, tt.actualModel, tt.effort)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestConvertRequest_OpenAIReasoningEffortCompatibility(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		mode        mode.Mode
		originModel string
		actualModel string
		body        string
		wantEffort  string
	}{
		{
			name:        "native chat gpt-5.5 minimal to low",
			mode:        mode.ChatCompletions,
			actualModel: "gpt-5.5",
			body:        `{"model":"alias","messages":[{"role":"user","content":"hi"}],"reasoning_effort":"minimal"}`,
			wantEffort:  "low",
		},
		{
			name:        "native completions gpt-5 xhigh to high",
			mode:        mode.Completions,
			actualModel: "gpt-5",
			body:        `{"model":"alias","prompt":"hi","reasoning_effort":"xhigh"}`,
			wantEffort:  "high",
		},
		{
			name:        "native responses gpt-5.1 xhigh to high",
			mode:        mode.Responses,
			actualModel: "gpt-5.1",
			body:        `{"model":"alias","input":"hi","reasoning":{"effort":"xhigh"}}`,
			wantEffort:  "high",
		},
		{
			name:        "native responses gpt-5.4 mini snapshot minimal to low",
			mode:        mode.Responses,
			actualModel: "gpt-5.4-mini-2026-03-17",
			body:        `{"model":"alias","input":"hi","reasoning":{"effort":"minimal"}}`,
			wantEffort:  "low",
		},
		{
			name:        "native chat origin match beats actual fallback",
			mode:        mode.ChatCompletions,
			originModel: "gpt-5",
			actualModel: "gpt-5.5",
			body:        `{"model":"alias","messages":[{"role":"user","content":"hi"}],"reasoning_effort":"none"}`,
			wantEffort:  "minimal",
		},
		{
			name:        "native completions actual fallback when origin alias does not match",
			mode:        mode.Completions,
			originModel: "customer-alias",
			actualModel: "gpt-5-pro",
			body:        `{"model":"alias","prompt":"hi","reasoning_effort":"none"}`,
			wantEffort:  "high",
		},
		{
			name:        "native chat unknown model preserves effort",
			mode:        mode.ChatCompletions,
			actualModel: "custom-model",
			body:        `{"model":"alias","messages":[{"role":"user","content":"hi"}],"reasoning_effort":"xhigh"}`,
			wantEffort:  "xhigh",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			httpReq := httptest.NewRequestWithContext(t.Context(),
				http.MethodPost,
				"/v1/test",
				strings.NewReader(tt.body),
			)
			m := &meta.Meta{
				Mode:        tt.mode,
				OriginModel: tt.originModel,
				ActualModel: tt.actualModel,
			}

			result, err := ConvertRequest(m, nil, httpReq)
			require.NoError(t, err)
			require.NotNil(t, result.Body)

			bodyBytes, err := io.ReadAll(result.Body)
			require.NoError(t, err)

			var payload map[string]any
			require.NoError(t, json.Unmarshal(bodyBytes, &payload))

			if tt.mode == mode.ChatCompletions &&
				IsResponsesOnlyModelAny(&m.ModelConfig, m.OriginModel, m.ActualModel) {
				reasoning, ok := payload["reasoning"].(map[string]any)
				require.True(t, ok)
				assert.Equal(t, tt.wantEffort, reasoning["effort"])
				return
			}

			if tt.mode == mode.Responses {
				reasoning, ok := payload["reasoning"].(map[string]any)
				require.True(t, ok)
				assert.Equal(t, tt.wantEffort, reasoning["effort"])
				return
			}

			assert.Equal(t, tt.wantEffort, payload["reasoning_effort"])
		})
	}
}

func TestConvertChatCompletionToResponsesRequest_ReasoningEffortCompatibility(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequestWithContext(t.Context(),
		http.MethodPost,
		"/v1/chat/completions",
		strings.NewReader(
			`{"model":"alias","messages":[{"role":"user","content":"hi"}],"reasoning_effort":"minimal"}`,
		),
	)
	m := &meta.Meta{
		OriginModel: "alias",
		ActualModel: "gpt-5.5",
	}

	result, err := ConvertChatCompletionToResponsesRequest(m, req)
	require.NoError(t, err)

	var responsesReq relaymodel.CreateResponseRequest
	require.NoError(t, json.NewDecoder(result.Body).Decode(&responsesReq))
	require.NotNil(t, responsesReq.Reasoning)
	require.NotNil(t, responsesReq.Reasoning.Effort)
	assert.Equal(t, "low", *responsesReq.Reasoning.Effort)
}

func TestDoResponse_MapReasoningToReasoningContent(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/chat/completions",
		nil,
	)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(strings.NewReader(
			`{"id":"chatcmpl-1","object":"chat.completion","created":1,"model":"upstream-model","choices":[{"index":0,"message":{"role":"assistant","reasoning":"internal-thought","content":"final-answer"},"finish_reason":"stop"}],"usage":{"prompt_tokens":1,"completion_tokens":2,"total_tokens":3}}`,
		)),
		Header: make(http.Header),
	}

	m := &meta.Meta{
		Mode:        mode.ChatCompletions,
		OriginModel: "gpt-4o",
		ActualModel: "gpt-4o",
		ChannelConfigs: model.ChannelConfigs{
			"map_reasoning_to_reasoning_content": true,
		},
	}

	result, err := DoResponse(m, nil, c, resp)
	require.Nil(t, err)
	assert.Equal(t, "chatcmpl-1", result.UpstreamID)

	var body map[string]any
	require.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &body))

	choices, ok := body["choices"].([]any)
	require.True(t, ok)
	require.Len(t, choices, 1)

	choice, ok := choices[0].(map[string]any)
	require.True(t, ok)
	message, ok := choice["message"].(map[string]any)
	require.True(t, ok)

	assert.Equal(t, "internal-thought", message["reasoning_content"])
	_, exists := message["reasoning"]
	assert.False(t, exists)
}

func TestDoResponseStream_MapReasoningToReasoningContent(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/chat/completions",
		nil,
	)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewBufferString(
			"data: {\"id\":\"chatcmpl-1\",\"object\":\"chat.completion.chunk\",\"created\":1,\"model\":\"upstream-model\",\"choices\":[{\"index\":0,\"delta\":{\"role\":\"assistant\",\"reasoning\":\"internal-thought\"}}]}\n\n" +
				"data: [DONE]\n\n",
		)),
		Header: http.Header{
			"Content-Type": {"text/event-stream"},
		},
	}

	m := &meta.Meta{
		Mode:        mode.ChatCompletions,
		OriginModel: "gpt-4o",
		ActualModel: "gpt-4o",
		ChannelConfigs: model.ChannelConfigs{
			"map_reasoning_to_reasoning_content": true,
		},
	}

	result, err := DoResponse(m, nil, c, resp)
	require.Nil(t, err)
	assert.Equal(t, "chatcmpl-1", result.UpstreamID)

	body := recorder.Body.String()
	assert.Contains(t, body, `"reasoning_content":"internal-thought"`)
	assert.NotContains(t, body, `"reasoning":"internal-thought"`)
}
