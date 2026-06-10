//nolint:testpackage
package qianfan

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	coremodel "github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/adaptor"
	"github.com/labring/aiproxy/core/relay/meta"
	"github.com/labring/aiproxy/core/relay/mode"
	relaymodel "github.com/labring/aiproxy/core/relay/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var noPayloadKeys = []string{
	"reasoning_effort",
	"enable_thinking",
	"thinking_budget",
	"thinking",
	"reasoning",
}

func TestAdaptorSupportModeGemini(t *testing.T) {
	adaptor := &Adaptor{}

	if !adaptor.SupportMode(&meta.Meta{Mode: mode.Gemini}) {
		t.Fatal("expected Gemini mode to be supported")
	}
}

func TestAdaptorSupportModeResponses(t *testing.T) {
	adaptor := &Adaptor{}

	supportedModes := []mode.Mode{
		mode.Responses,
		mode.ResponsesGet,
		mode.ResponsesDelete,
		mode.ResponsesInputItems,
	}
	for _, m := range supportedModes {
		if !adaptor.SupportMode(&meta.Meta{
			Mode:        m,
			OriginModel: "deepseek-v3.2",
			ActualModel: "mapped-model",
		}) {
			t.Fatalf("expected mode %s to be supported", m)
		}
	}

	if adaptor.SupportMode(&meta.Meta{Mode: mode.ResponsesCancel}) {
		t.Fatal("expected ResponsesCancel to be unsupported")
	}
}

func TestAdaptorSupportModeResponsesModelWhitelist(t *testing.T) {
	adaptor := &Adaptor{}

	if !adaptor.SupportMode(&meta.Meta{
		Mode:        mode.Responses,
		OriginModel: "unsupported-alias",
		ActualModel: "deepseek-v3.1-250821",
	}) {
		t.Fatal("expected Responses to be supported when actual model is whitelisted")
	}

	if adaptor.SupportMode(&meta.Meta{
		Mode:        mode.Responses,
		OriginModel: "ernie-4.5-turbo-128k",
		ActualModel: "ernie-4.5-turbo-128k",
	}) {
		t.Fatal("expected Responses to be unsupported for non-whitelisted model")
	}
}

func TestAdaptorSupportModeResponsesModelConfig(t *testing.T) {
	adaptor := &Adaptor{}

	if !adaptor.SupportMode(&meta.Meta{
		Mode:        mode.Responses,
		OriginModel: "custom-responses-model",
		ActualModel: "upstream-custom-model",
		ChannelConfigs: coremodel.ChannelConfigs{
			"response_models": []string{"custom-responses-model"},
		},
	}) {
		t.Fatal("expected Responses to be supported by channel response_models config")
	}
}

func TestAdaptorGetRequestURLResponses(t *testing.T) {
	adaptor := &Adaptor{}
	channel := &coremodel.Channel{BaseURL: "https://qianfan.baidubce.com/v2"}

	tests := []struct {
		name       string
		mode       mode.Mode
		responseID string
		wantMethod string
		wantURL    string
	}{
		{
			name:       "responses create",
			mode:       mode.Responses,
			wantMethod: http.MethodPost,
			wantURL:    "https://qianfan.baidubce.com/v2/responses",
		},
		{
			name:       "responses get",
			mode:       mode.ResponsesGet,
			responseID: "resp_123",
			wantMethod: http.MethodGet,
			wantURL:    "https://qianfan.baidubce.com/v2/responses/resp_123",
		},
		{
			name:       "responses delete",
			mode:       mode.ResponsesDelete,
			responseID: "resp_123",
			wantMethod: http.MethodDelete,
			wantURL:    "https://qianfan.baidubce.com/v2/responses/resp_123",
		},
		{
			name:       "responses input items",
			mode:       mode.ResponsesInputItems,
			responseID: "resp_123",
			wantMethod: http.MethodGet,
			wantURL:    "https://qianfan.baidubce.com/v2/responses/resp_123/input_items",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := meta.NewMeta(
				channel,
				tt.mode,
				"ernie-4.5-turbo-128k",
				coremodel.ModelConfig{},
				meta.WithResponseID(tt.responseID),
			)

			got, err := adaptor.GetRequestURL(m, nil, nil)
			require.NoError(t, err)
			assert.Equal(t, tt.wantMethod, got.Method)
			assert.Equal(t, tt.wantURL, got.URL)
		})
	}
}

func TestAdaptorGetRequestURLResponsesCancelUnsupported(t *testing.T) {
	adaptor := &Adaptor{}
	m := meta.NewMeta(
		&coremodel.Channel{BaseURL: "https://qianfan.baidubce.com/v2"},
		mode.ResponsesCancel,
		"ernie-4.5-turbo-128k",
		coremodel.ModelConfig{},
		meta.WithResponseID("resp_123"),
	)

	_, err := adaptor.GetRequestURL(m, nil, nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported mode")
}

func TestAdaptorConvertRequestResponses(t *testing.T) {
	adaptor := &Adaptor{}
	m := meta.NewMeta(
		nil,
		mode.Responses,
		"ernie-4.5-turbo-128k",
		coremodel.ModelConfig{},
	)

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"/v1/responses",
		strings.NewReader(`{"model":"ernie-4.5-turbo-128k","input":"hello","stream":true}`),
	)
	require.NoError(t, err)

	result, err := adaptor.ConvertRequest(m, nil, req)
	require.NoError(t, err)

	body, err := io.ReadAll(result.Body)
	require.NoError(t, err)

	var responseReq relaymodel.CreateResponseRequest
	require.NoError(t, json.Unmarshal(body, &responseReq))
	assert.Equal(t, "ernie-4.5-turbo-128k", responseReq.Model)
	assert.True(t, responseReq.Stream)
}

func TestAdaptorConvertRequestReasoning(t *testing.T) {
	adaptor := &Adaptor{}

	t.Run(
		"chat reasoning_effort none maps to disabled thinking for thinking models",
		func(t *testing.T) {
			m := meta.NewMeta(nil, mode.ChatCompletions, "deepseek-v3.2", coremodel.ModelConfig{})
			req, err := http.NewRequestWithContext(
				context.Background(),
				http.MethodPost,
				"/v1/chat/completions",
				strings.NewReader(`{
				"model":"deepseek-v3.2",
				"reasoning_effort":"none",
				"messages":[{"role":"user","content":"hello"}]
			}`),
			)
			require.NoError(t, err)

			result, err := adaptor.ConvertRequest(m, nil, req)
			require.NoError(t, err)

			body, err := io.ReadAll(result.Body)
			require.NoError(t, err)

			var payload map[string]any
			require.NoError(t, json.Unmarshal(body, &payload))
			assert.NotContains(t, payload, "reasoning_effort")

			thinking, ok := payload["thinking"].(map[string]any)
			require.True(t, ok)
			assert.Equal(t, "disabled", thinking["type"])
		},
	)

	t.Run("chat low reasoning_effort maps to enable_thinking with budget", func(t *testing.T) {
		m := meta.NewMeta(nil, mode.ChatCompletions, "qwen3-14b", coremodel.ModelConfig{})
		req, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			"/v1/chat/completions",
			strings.NewReader(`{
				"model":"qwen3-14b",
				"reasoning_effort":"low",
				"messages":[{"role":"user","content":"hello"}]
			}`),
		)
		require.NoError(t, err)

		result, err := adaptor.ConvertRequest(m, nil, req)
		require.NoError(t, err)

		body, err := io.ReadAll(result.Body)
		require.NoError(t, err)

		var payload map[string]any
		require.NoError(t, json.Unmarshal(body, &payload))
		assert.Equal(t, true, payload["enable_thinking"])
		assert.Equal(t, float64(2048), payload["thinking_budget"])
		assert.NotContains(t, payload, "reasoning_effort")
		assert.NotContains(t, payload, "thinking")
	})

	t.Run("chat xhigh reasoning_effort maps to deepseek v4 thinking budget", func(t *testing.T) {
		m := meta.NewMeta(nil, mode.ChatCompletions, "deepseek-v4-pro", coremodel.ModelConfig{})
		req, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			"/v1/chat/completions",
			strings.NewReader(`{
				"model":"deepseek-v4-pro",
				"reasoning_effort":"xhigh",
				"messages":[{"role":"user","content":"hello"}]
			}`),
		)
		require.NoError(t, err)

		result, err := adaptor.ConvertRequest(m, nil, req)
		require.NoError(t, err)

		body, err := io.ReadAll(result.Body)
		require.NoError(t, err)

		var payload map[string]any
		require.NoError(t, json.Unmarshal(body, &payload))
		thinking, ok := payload["thinking"].(map[string]any)
		require.True(t, ok)
		assert.Equal(t, relaymodel.ClaudeThinkingTypeEnabled, thinking["type"])
		assert.Equal(t, float64(16384), payload["thinking_budget"])
		assert.NotContains(t, payload, "reasoning_effort")
	})

	t.Run("native thinking wins over reasoning_effort", func(t *testing.T) {
		m := meta.NewMeta(nil, mode.ChatCompletions, "deepseek-v3.2", coremodel.ModelConfig{})
		req, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			"/v1/chat/completions",
			strings.NewReader(`{
				"model":"deepseek-v3.2",
				"reasoning_effort":"none",
				"thinking":{"type":"enabled","budget_tokens":2048},
				"messages":[{"role":"user","content":"hello"}]
			}`),
		)
		require.NoError(t, err)

		result, err := adaptor.ConvertRequest(m, nil, req)
		require.NoError(t, err)

		body, err := io.ReadAll(result.Body)
		require.NoError(t, err)

		var payload map[string]any
		require.NoError(t, json.Unmarshal(body, &payload))
		assert.NotContains(t, payload, "reasoning_effort")

		thinking, ok := payload["thinking"].(map[string]any)
		require.True(t, ok)
		assert.Equal(t, "enabled", thinking["type"])
		assert.Equal(t, float64(2048), thinking["budget_tokens"])
	})

	t.Run("chat reasoning_effort none maps to enable_thinking false", func(t *testing.T) {
		m := meta.NewMeta(nil, mode.ChatCompletions, "qwen3-14b", coremodel.ModelConfig{})
		req, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			"/v1/chat/completions",
			strings.NewReader(`{
				"model":"qwen3-14b",
				"reasoning_effort":"none",
				"messages":[{"role":"user","content":"hello"}]
			}`),
		)
		require.NoError(t, err)

		result, err := adaptor.ConvertRequest(m, nil, req)
		require.NoError(t, err)

		body, err := io.ReadAll(result.Body)
		require.NoError(t, err)

		var payload map[string]any
		require.NoError(t, json.Unmarshal(body, &payload))
		assert.Equal(t, false, payload["enable_thinking"])
		assert.NotContains(t, payload, "thinking_budget")
		assert.NotContains(t, payload, "thinking")
		assert.NotContains(t, payload, "reasoning_effort")
	})

	t.Run("gemini disabled thinking maps to enable_thinking false", func(t *testing.T) {
		m := meta.NewMeta(nil, mode.Gemini, "qwen3-14b", coremodel.ModelConfig{})
		req, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			"/v1beta/models/qwen3-14b:generateContent",
			strings.NewReader(`{
				"generationConfig": {
					"thinkingConfig": {
						"thinkingBudget": 0
					}
				},
				"contents":[{"role":"user","parts":[{"text":"hello"}]}]
			}`),
		)
		require.NoError(t, err)

		result, err := adaptor.ConvertRequest(m, nil, req)
		require.NoError(t, err)

		body, err := io.ReadAll(result.Body)
		require.NoError(t, err)

		var openAIReq relaymodel.GeneralOpenAIRequest
		require.NoError(t, json.Unmarshal(body, &openAIReq))
		require.NotNil(t, openAIReq.EnableThinking)
		assert.False(t, *openAIReq.EnableThinking)
		assert.Nil(t, openAIReq.Thinking)
		assert.Nil(t, openAIReq.ReasoningEffort)
	})

	t.Run("responses reasoning_effort none maps to disabled thinking", func(t *testing.T) {
		m := meta.NewMeta(nil, mode.Responses, "deepseek-v3.2", coremodel.ModelConfig{})
		req, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			"/v1/responses",
			strings.NewReader(`{
				"model":"deepseek-v3.2",
				"input":"hello",
				"reasoning":{"effort":"none"}
			}`),
		)
		require.NoError(t, err)

		result, err := adaptor.ConvertRequest(m, nil, req)
		require.NoError(t, err)

		body, err := io.ReadAll(result.Body)
		require.NoError(t, err)

		var payload map[string]any
		require.NoError(t, json.Unmarshal(body, &payload))
		assert.NotContains(t, payload, "reasoning")

		thinking, ok := payload["thinking"].(map[string]any)
		require.True(t, ok)
		assert.Equal(t, "disabled", thinking["type"])
	})

	t.Run("responses native thinking wins over reasoning", func(t *testing.T) {
		m := meta.NewMeta(nil, mode.Responses, "deepseek-v3.2", coremodel.ModelConfig{})
		req, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			"/v1/responses",
			strings.NewReader(`{
				"model":"deepseek-v3.2",
				"input":"hello",
				"thinking":{"type":"enabled"},
				"reasoning":{"effort":"none"}
			}`),
		)
		require.NoError(t, err)

		result, err := adaptor.ConvertRequest(m, nil, req)
		require.NoError(t, err)

		body, err := io.ReadAll(result.Body)
		require.NoError(t, err)

		var payload map[string]any
		require.NoError(t, json.Unmarshal(body, &payload))
		assert.NotContains(t, payload, "reasoning")

		thinking, ok := payload["thinking"].(map[string]any)
		require.True(t, ok)
		assert.Equal(t, "enabled", thinking["type"])
	})

	t.Run("actual model fallback selects qianfan parameter family", func(t *testing.T) {
		m := meta.NewMeta(nil, mode.ChatCompletions, "alias-model", coremodel.ModelConfig{})
		m.ActualModel = "qwen3-14b"
		req, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			"/v1/chat/completions",
			strings.NewReader(`{
				"model":"alias-model",
				"reasoning_effort":"medium",
				"messages":[{"role":"user","content":"hello"}]
			}`),
		)
		require.NoError(t, err)

		result, err := adaptor.ConvertRequest(m, nil, req)
		require.NoError(t, err)

		body, err := io.ReadAll(result.Body)
		require.NoError(t, err)

		var payload map[string]any
		require.NoError(t, json.Unmarshal(body, &payload))
		assert.Equal(t, true, payload["enable_thinking"])
		assert.Equal(t, float64(8192), payload["thinking_budget"])
	})

	t.Run("model family fallback handles versioned qwen3 models", func(t *testing.T) {
		m := meta.NewMeta(
			nil,
			mode.ChatCompletions,
			"qwen3-14b-thinking-260101",
			coremodel.ModelConfig{},
		)
		req, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			"/v1/chat/completions",
			strings.NewReader(`{
				"model":"qwen3-14b-thinking-260101",
				"reasoning_effort":"high",
				"messages":[{"role":"user","content":"hello"}]
			}`),
		)
		require.NoError(t, err)

		result, err := adaptor.ConvertRequest(m, nil, req)
		require.NoError(t, err)

		body, err := io.ReadAll(result.Body)
		require.NoError(t, err)

		var payload map[string]any
		require.NoError(t, json.Unmarshal(body, &payload))
		assert.Equal(t, true, payload["enable_thinking"])
		assert.Equal(t, float64(16384), payload["thinking_budget"])
	})

	t.Run("keyword fallback handles vl enable_thinking models", func(t *testing.T) {
		m := meta.NewMeta(nil, mode.ChatCompletions, "ernie-4.5-vl-custom", coremodel.ModelConfig{})
		req, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			"/v1/chat/completions",
			strings.NewReader(`{
				"model":"ernie-4.5-vl-custom",
				"reasoning_effort":"none",
				"messages":[{"role":"user","content":"hello"}]
			}`),
		)
		require.NoError(t, err)

		result, err := adaptor.ConvertRequest(m, nil, req)
		require.NoError(t, err)

		body, err := io.ReadAll(result.Body)
		require.NoError(t, err)

		var payload map[string]any
		require.NoError(t, json.Unmarshal(body, &payload))
		assert.Equal(t, false, payload["enable_thinking"])
		assert.NotContains(t, payload, "thinking")
	})

	t.Run("unsupported model strips normalized reasoning controls", func(t *testing.T) {
		m := meta.NewMeta(
			nil,
			mode.ChatCompletions,
			"ernie-4.5-turbo-128k",
			coremodel.ModelConfig{},
		)
		req, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			"/v1/chat/completions",
			strings.NewReader(`{
				"model":"ernie-4.5-turbo-128k",
				"reasoning_effort":"high",
				"messages":[{"role":"user","content":"hello"}]
			}`),
		)
		require.NoError(t, err)

		result, err := adaptor.ConvertRequest(m, nil, req)
		require.NoError(t, err)

		body, err := io.ReadAll(result.Body)
		require.NoError(t, err)

		var payload map[string]any
		require.NoError(t, json.Unmarshal(body, &payload))
		assert.NotContains(t, payload, "reasoning_effort")
		assert.NotContains(t, payload, "thinking")
		assert.NotContains(t, payload, "enable_thinking")
		assert.NotContains(t, payload, "thinking_budget")
	})
}

func TestQianfanReasoningNodeFieldFamilies(t *testing.T) {
	tests := []struct {
		name      string
		origin    string
		actual    string
		effort    string
		want      map[string]any
		wantThink string
		absent    []string
	}{
		{
			name:   "deepseek v4 thinking model high with budget",
			origin: "deepseek-v4-pro",
			effort: "high",
			want: map[string]any{
				"thinking_budget": float64(16384),
			},
			wantThink: relaymodel.ClaudeThinkingTypeEnabled,
			absent:    []string{"reasoning_effort", "enable_thinking", "reasoning"},
		},
		{
			name:   "deepseek v4 family fallback thinking model xhigh with budget",
			origin: "deepseek-v4-lite-260101",
			effort: "xhigh",
			want: map[string]any{
				"thinking_budget": float64(16384),
			},
			wantThink: relaymodel.ClaudeThinkingTypeEnabled,
			absent:    []string{"reasoning_effort", "enable_thinking", "reasoning"},
		},
		{
			name:      "deepseek v4 thinking model disabled",
			origin:    "deepseek-v4-flash",
			effort:    "none",
			wantThink: relaymodel.ClaudeThinkingTypeDisabled,
			absent: []string{
				"reasoning_effort",
				"enable_thinking",
				"thinking_budget",
				"reasoning",
			},
		},
		{
			name:   "enable thinking model enabled with budget",
			origin: "qwen3-14b",
			effort: "medium",
			want: map[string]any{
				"enable_thinking": true,
				"thinking_budget": float64(8192),
			},
			absent: []string{"reasoning_effort", "thinking", "reasoning"},
		},
		{
			name:   "enable thinking model disabled",
			origin: "qwen3-14b",
			effort: "none",
			want: map[string]any{
				"enable_thinking": false,
			},
			absent: []string{"reasoning_effort", "thinking_budget", "thinking", "reasoning"},
		},
		{
			name:   "vl keyword fallback enables thinking without budget",
			origin: "ernie-4.5-vl-custom",
			effort: "low",
			want: map[string]any{
				"enable_thinking": true,
			},
			absent: []string{"reasoning_effort", "thinking_budget", "thinking", "reasoning"},
		},
		{
			name:      "thinking model enabled",
			origin:    "deepseek-v3.2",
			effort:    "high",
			wantThink: relaymodel.ClaudeThinkingTypeEnabled,
			absent: []string{
				"reasoning_effort",
				"enable_thinking",
				"thinking_budget",
				"reasoning",
			},
		},
		{
			name:      "thinking model disabled",
			origin:    "deepseek-v3.2",
			effort:    "none",
			wantThink: relaymodel.ClaudeThinkingTypeDisabled,
			absent: []string{
				"reasoning_effort",
				"enable_thinking",
				"thinking_budget",
				"reasoning",
			},
		},
		{
			name:   "thinking model with budget enabled",
			origin: "deepseek-v3.1-250821",
			effort: "low",
			want: map[string]any{
				"thinking_budget": float64(2048),
			},
			wantThink: relaymodel.ClaudeThinkingTypeEnabled,
			absent:    []string{"reasoning_effort", "enable_thinking", "reasoning"},
		},
		{
			name:   "deepseek r1 thinking model with budget enabled",
			origin: "deepseek-r1-250528",
			effort: "minimal",
			want: map[string]any{
				"thinking_budget": float64(1024),
			},
			wantThink: relaymodel.ClaudeThinkingTypeEnabled,
			absent:    []string{"reasoning_effort", "enable_thinking", "reasoning"},
		},
		{
			name:      "deepseek r1 thinking model disabled",
			origin:    "deepseek-r1-250528",
			effort:    "none",
			wantThink: relaymodel.ClaudeThinkingTypeDisabled,
			absent: []string{
				"reasoning_effort",
				"enable_thinking",
				"thinking_budget",
				"reasoning",
			},
		},
		{
			name:   "thinking keyword fallback budget only",
			origin: "custom-thinking-model",
			effort: "xhigh",
			want: map[string]any{
				"thinking_budget": float64(16384),
			},
			absent: []string{"reasoning_effort", "enable_thinking", "thinking", "reasoning"},
		},
		{
			name:   "actual fallback selects capability",
			origin: "alias-model",
			actual: "qwen3-30b-a3b",
			effort: "medium",
			want: map[string]any{
				"enable_thinking": true,
				"thinking_budget": float64(8192),
			},
			absent: []string{"reasoning_effort", "thinking", "reasoning"},
		},
		{
			name:   "actual fallback selects thinking capability",
			origin: "alias-model",
			actual: "deepseek-v4-pro",
			effort: "medium",
			want: map[string]any{
				"thinking_budget": float64(8192),
			},
			wantThink: relaymodel.ClaudeThinkingTypeEnabled,
			absent:    []string{"reasoning_effort", "enable_thinking", "reasoning"},
		},
		{
			name:   "unsupported model strips controls",
			origin: "ernie-4.5-turbo-128k",
			effort: "high",
			absent: noPayloadKeys,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload := convertQianfanChatPayload(
				t,
				tt.origin,
				tt.actual,
				`"reasoning_effort":"`+tt.effort+`"`,
			)

			for key, want := range tt.want {
				assert.Equal(t, want, payload[key], key)
			}

			if tt.wantThink != "" {
				thinking, ok := payload["thinking"].(map[string]any)
				require.True(t, ok)
				assert.Equal(t, tt.wantThink, thinking["type"])
			}

			for _, key := range tt.absent {
				assert.NotContains(t, payload, key)
			}
		})
	}
}

func TestQianfanReasoningRequestFieldFamilies(t *testing.T) {
	enabledReasoning := relaymodel.NormalizedReasoning{
		Specified: true,
		Effort:    relaymodel.ReasoningEffortHigh,
	}
	disabledReasoning := relaymodel.NormalizedReasoning{
		Specified: true,
		Disabled:  true,
		Effort:    relaymodel.ReasoningEffortNone,
	}

	tests := []struct {
		name         string
		model        string
		reasoning    relaymodel.NormalizedReasoning
		wantEffort   string
		hasEffort    bool
		wantEnable   bool
		hasEnable    bool
		wantBudget   int
		hasBudget    bool
		wantThinking string
		hasThinking  bool
	}{
		{
			name:         "request deepseek v4 thinking model enabled with budget",
			model:        "deepseek-v4-pro",
			reasoning:    enabledReasoning,
			wantBudget:   16384,
			hasBudget:    true,
			wantThinking: relaymodel.ClaudeThinkingTypeEnabled,
			hasThinking:  true,
		},
		{
			name:         "request deepseek v4 thinking model disabled",
			model:        "deepseek-v4-pro",
			reasoning:    disabledReasoning,
			wantThinking: relaymodel.ClaudeThinkingTypeDisabled,
			hasThinking:  true,
		},
		{
			name:       "request enable_thinking model enabled",
			model:      "qwen3-14b",
			reasoning:  enabledReasoning,
			wantEnable: true,
			hasEnable:  true,
			wantBudget: 16384,
			hasBudget:  true,
		},
		{
			name:       "request enable_thinking model disabled",
			model:      "qwen3-14b",
			reasoning:  disabledReasoning,
			wantEnable: false,
			hasEnable:  true,
		},
		{
			name:         "request thinking model enabled",
			model:        "deepseek-v3.2",
			reasoning:    enabledReasoning,
			wantThinking: relaymodel.ClaudeThinkingTypeEnabled,
			hasThinking:  true,
		},
		{
			name:         "request thinking model disabled",
			model:        "deepseek-v3.2",
			reasoning:    disabledReasoning,
			wantThinking: relaymodel.ClaudeThinkingTypeDisabled,
			hasThinking:  true,
		},
		{
			name:         "request deepseek r1 thinking model with budget enabled",
			model:        "deepseek-r1-250528",
			reasoning:    enabledReasoning,
			wantBudget:   16384,
			hasBudget:    true,
			wantThinking: relaymodel.ClaudeThinkingTypeEnabled,
			hasThinking:  true,
		},
		{
			name:         "request deepseek r1 thinking model disabled",
			model:        "deepseek-r1-250528",
			reasoning:    disabledReasoning,
			wantThinking: relaymodel.ClaudeThinkingTypeDisabled,
			hasThinking:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &relaymodel.GeneralOpenAIRequest{}
			m := meta.NewMeta(nil, mode.Gemini, tt.model, coremodel.ModelConfig{})

			applyQianfanReasoningToRequest(m, req, tt.reasoning)

			if !tt.hasEffort {
				assert.Nil(t, req.ReasoningEffort)
			} else {
				require.NotNil(t, req.ReasoningEffort)
				assert.Equal(t, tt.wantEffort, *req.ReasoningEffort)
			}

			if !tt.hasEnable {
				assert.Nil(t, req.EnableThinking)
			} else {
				require.NotNil(t, req.EnableThinking)
				assert.Equal(t, tt.wantEnable, *req.EnableThinking)
			}

			if !tt.hasBudget {
				assert.Nil(t, req.ThinkingBudget)
			} else {
				require.NotNil(t, req.ThinkingBudget)
				assert.Equal(t, tt.wantBudget, *req.ThinkingBudget)
			}

			if !tt.hasThinking {
				assert.Nil(t, req.Thinking)
			} else {
				require.NotNil(t, req.Thinking)
				assert.Equal(t, tt.wantThinking, req.Thinking.Type)
			}
		})
	}
}

func TestQianfanReasoningEffortAndBudgetMapping(t *testing.T) {
	tests := []struct {
		effort     string
		wantEffort string
		wantBudget float64
	}{
		{effort: "minimal", wantEffort: "high", wantBudget: 1024},
		{effort: "low", wantEffort: "high", wantBudget: 2048},
		{effort: "medium", wantEffort: "high", wantBudget: 8192},
		{effort: "high", wantEffort: "high", wantBudget: 16384},
		{effort: "xhigh", wantEffort: "max", wantBudget: 16384},
	}

	for _, tt := range tests {
		t.Run(tt.effort+" maps deepseek v4 thinking budget", func(t *testing.T) {
			payload := convertQianfanChatPayload(
				t,
				"deepseek-v4-pro",
				"",
				`"reasoning_effort":"`+tt.effort+`"`,
			)

			thinking, ok := payload["thinking"].(map[string]any)
			require.True(t, ok)
			assert.Equal(t, relaymodel.ClaudeThinkingTypeEnabled, thinking["type"])
			assert.NotContains(t, payload, "reasoning_effort")
			assert.Equal(t, tt.wantBudget, payload["thinking_budget"])
		})

		t.Run(tt.effort+" maps budget", func(t *testing.T) {
			payload := convertQianfanChatPayload(
				t,
				"qwen3-14b",
				"",
				`"reasoning_effort":"`+tt.effort+`"`,
			)

			assert.Equal(t, true, payload["enable_thinking"])
			assert.Equal(t, tt.wantBudget, payload["thinking_budget"])
		})
	}
}

func TestQianfanReasoningNativeThinkingCleanup(t *testing.T) {
	payload := convertQianfanChatPayload(
		t,
		"qwen3-14b",
		"",
		`"reasoning_effort":"none",
		"enable_thinking":true,
		"thinking_budget":8192,
		"thinking":{"type":"enabled","budget_tokens":2048}`,
	)

	assert.NotContains(t, payload, "reasoning_effort")
	assert.NotContains(t, payload, "enable_thinking")
	assert.NotContains(t, payload, "thinking_budget")

	thinking, ok := payload["thinking"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "enabled", thinking["type"])
	assert.Equal(t, float64(2048), thinking["budget_tokens"])
}

func TestQianfanReasoningInvalidEffortPassesThrough(t *testing.T) {
	payload := convertQianfanChatPayload(
		t,
		"qwen3-14b",
		"",
		`"reasoning_effort":"invalid"`,
	)

	assert.Equal(t, "invalid", payload["reasoning_effort"])
	assert.NotContains(t, payload, "enable_thinking")
	assert.NotContains(t, payload, "thinking_budget")
	assert.NotContains(t, payload, "thinking")
}

func TestQianfanReasoningCompletionsMode(t *testing.T) {
	adaptor := &Adaptor{}
	m := meta.NewMeta(nil, mode.Completions, "qwen3-14b", coremodel.ModelConfig{})
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"/v1/completions",
		strings.NewReader(`{
			"model":"qwen3-14b",
			"prompt":"hello",
			"reasoning_effort":"high"
		}`),
	)
	require.NoError(t, err)

	result, err := adaptor.ConvertRequest(m, nil, req)
	require.NoError(t, err)

	payload := readJSONMap(t, result.Body)
	assert.Equal(t, true, payload["enable_thinking"])
	assert.Equal(t, float64(16384), payload["thinking_budget"])
	assert.NotContains(t, payload, "reasoning_effort")
}

func TestQianfanReasoningAnthropicMode(t *testing.T) {
	adaptor := &Adaptor{}
	m := meta.NewMeta(nil, mode.Anthropic, "qwen3-14b", coremodel.ModelConfig{})
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"/v1/messages",
		strings.NewReader(`{
			"model":"qwen3-14b",
			"thinking":{"type":"adaptive"},
			"output_config":{"effort":"high"},
			"messages":[{"role":"user","content":"hello"}]
		}`),
	)
	require.NoError(t, err)

	result, err := adaptor.ConvertRequest(m, nil, req)
	require.NoError(t, err)

	var openAIReq relaymodel.GeneralOpenAIRequest
	require.NoError(t, json.NewDecoder(result.Body).Decode(&openAIReq))
	require.NotNil(t, openAIReq.EnableThinking)
	assert.True(t, *openAIReq.EnableThinking)
	require.NotNil(t, openAIReq.ThinkingBudget)
	assert.Equal(t, 16384, *openAIReq.ThinkingBudget)
	assert.Nil(t, openAIReq.Thinking)
	assert.Nil(t, openAIReq.ReasoningEffort)
}

func TestQianfanReasoningGeminiEnabledBudget(t *testing.T) {
	adaptor := &Adaptor{}
	m := meta.NewMeta(nil, mode.Gemini, "qwen3-14b", coremodel.ModelConfig{})
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"/v1beta/models/qwen3-14b:generateContent",
		strings.NewReader(`{
			"generationConfig":{"thinkingConfig":{"thinkingBudget":2048,"includeThoughts":true}},
			"contents":[{"role":"user","parts":[{"text":"hello"}]}]
		}`),
	)
	require.NoError(t, err)

	result, err := adaptor.ConvertRequest(m, nil, req)
	require.NoError(t, err)

	var openAIReq relaymodel.GeneralOpenAIRequest
	require.NoError(t, json.NewDecoder(result.Body).Decode(&openAIReq))
	require.NotNil(t, openAIReq.EnableThinking)
	assert.True(t, *openAIReq.EnableThinking)
	require.NotNil(t, openAIReq.ThinkingBudget)
	assert.Equal(t, 2048, *openAIReq.ThinkingBudget)
	assert.Nil(t, openAIReq.Thinking)
	assert.Nil(t, openAIReq.ReasoningEffort)
}

func TestQianfanReasoningResponsesEnabled(t *testing.T) {
	adaptor := &Adaptor{}
	m := meta.NewMeta(nil, mode.Responses, "deepseek-v4-pro", coremodel.ModelConfig{})
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"/v1/responses",
		strings.NewReader(`{
			"model":"deepseek-v4-pro",
			"input":"hello",
			"reasoning":{"effort":"xhigh"}
		}`),
	)
	require.NoError(t, err)

	result, err := adaptor.ConvertRequest(m, nil, req)
	require.NoError(t, err)

	payload := readJSONMap(t, result.Body)
	thinking, ok := payload["thinking"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, relaymodel.ClaudeThinkingTypeEnabled, thinking["type"])
	assert.Equal(t, float64(16384), payload["thinking_budget"])
	assert.NotContains(t, payload, "reasoning")
	assert.NotContains(t, payload, "reasoning_effort")
}

func TestQianfanReasoningResponsesInvalidEffortPassesThrough(t *testing.T) {
	adaptor := &Adaptor{}
	m := meta.NewMeta(nil, mode.Responses, "deepseek-v4-pro", coremodel.ModelConfig{})
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"/v1/responses",
		strings.NewReader(`{
			"model":"deepseek-v4-pro",
			"input":"hello",
			"reasoning":{"effort":"invalid"}
		}`),
	)
	require.NoError(t, err)

	result, err := adaptor.ConvertRequest(m, nil, req)
	require.NoError(t, err)

	payload := readJSONMap(t, result.Body)
	reasoning, ok := payload["reasoning"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "invalid", reasoning["effort"])
	assert.NotContains(t, payload, "reasoning_effort")
	assert.NotContains(t, payload, "thinking")
}

func convertQianfanChatPayload(
	t *testing.T,
	originModel string,
	actualModel string,
	extraFields string,
) map[string]any {
	t.Helper()

	adaptor := &Adaptor{}

	m := meta.NewMeta(nil, mode.ChatCompletions, originModel, coremodel.ModelConfig{})
	if actualModel != "" {
		m.ActualModel = actualModel
	}

	body := `{
		"model":"` + originModel + `",
		` + extraFields + `,
		"messages":[{"role":"user","content":"hello"}]
	}`
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"/v1/chat/completions",
		strings.NewReader(body),
	)
	require.NoError(t, err)

	result, err := adaptor.ConvertRequest(m, nil, req)
	require.NoError(t, err)

	return readJSONMap(t, result.Body)
}

func readJSONMap(t *testing.T, r io.Reader) map[string]any {
	t.Helper()

	var payload map[string]any
	require.NoError(t, json.NewDecoder(r).Decode(&payload))

	return payload
}

func TestAdaptorDoResponseResponsesDeleteNoContent(t *testing.T) {
	gin.SetMode(gin.TestMode)

	adaptor := &Adaptor{}
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	resp := &http.Response{
		StatusCode: http.StatusNoContent,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader("")),
	}

	_, err := adaptor.DoResponse(
		&meta.Meta{Mode: mode.ResponsesDelete},
		nil,
		ctx,
		resp,
	)

	require.Nil(t, err)
	assert.Equal(t, http.StatusNoContent, ctx.Writer.Status())
}

func TestAdaptorConvertRequestResponsesCancelUnsupported(t *testing.T) {
	adaptor := &Adaptor{}
	m := meta.NewMeta(
		nil,
		mode.ResponsesCancel,
		"ernie-4.5-turbo-128k",
		coremodel.ModelConfig{},
	)

	_, err := adaptor.ConvertRequest(m, nil, httptest.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"/v1/responses/resp_123/cancel",
		nil,
	))

	require.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported mode")
}

var _ adaptor.Adaptor = (*Adaptor)(nil)

func TestAdaptorSetupRequestHeaderWithAppID(t *testing.T) {
	adaptor := &Adaptor{}
	m := meta.NewMeta(
		&coremodel.Channel{
			Key: "test-key",
			Configs: coremodel.ChannelConfigs{
				"appid": " app-test ",
			},
		},
		mode.ChatCompletions,
		"ernie-4.5-turbo-128k",
		coremodel.ModelConfig{},
	)
	req := httptest.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"https://qianfan.baidubce.com/v2/chat/completions",
		nil,
	)

	err := adaptor.SetupRequestHeader(m, nil, nil, req)

	require.NoError(t, err)
	assert.Equal(t, "Bearer test-key", req.Header.Get("Authorization"))
	assert.Equal(t, "app-test", req.Header.Get("Appid"))
}

func TestAdaptorSetupRequestHeaderWithoutAppID(t *testing.T) {
	adaptor := &Adaptor{}
	m := meta.NewMeta(
		&coremodel.Channel{
			Key: "test-key",
		},
		mode.ChatCompletions,
		"ernie-4.5-turbo-128k",
		coremodel.ModelConfig{},
	)
	req := httptest.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"https://qianfan.baidubce.com/v2/chat/completions",
		nil,
	)

	err := adaptor.SetupRequestHeader(m, nil, nil, req)

	require.NoError(t, err)
	assert.Equal(t, "Bearer test-key", req.Header.Get("Authorization"))
	assert.Empty(t, req.Header.Get("Appid"))
}

func TestAdaptorMetadataConfigSchema(t *testing.T) {
	adaptor := &Adaptor{}
	metaInfo := adaptor.Metadata()

	properties, ok := metaInfo.ConfigSchema["properties"].(map[string]any)
	require.True(t, ok)

	field, ok := properties["appid"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "string", field["type"])

	field, ok = properties["response_models"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "array", field["type"])
}

func TestAdaptorConvertRequestGemini(t *testing.T) {
	adaptor := &Adaptor{}
	m := meta.NewMeta(
		nil,
		mode.Gemini,
		"ernie-4.5-turbo-128k",
		coremodel.ModelConfig{},
	)

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"/v1beta/models/ernie-4.5-turbo-128k:streamGenerateContent",
		strings.NewReader(`{"contents":[{"role":"user","parts":[{"text":"hello"}]}]}`),
	)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	result, err := adaptor.ConvertRequest(m, nil, req)
	if err != nil {
		t.Fatalf("ConvertRequest returned error: %v", err)
	}

	body, err := io.ReadAll(result.Body)
	if err != nil {
		t.Fatalf("failed to read converted body: %v", err)
	}

	var openAIReq relaymodel.GeneralOpenAIRequest
	if err := json.Unmarshal(body, &openAIReq); err != nil {
		t.Fatalf("failed to unmarshal converted body: %v", err)
	}

	if openAIReq.Model != "ernie-4.5-turbo-128k" {
		t.Fatalf("expected model ernie-4.5-turbo-128k, got %s", openAIReq.Model)
	}

	if !openAIReq.Stream {
		t.Fatal("expected stream to be enabled")
	}
}
