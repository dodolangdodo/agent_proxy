package qianfan

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/bytedance/sonic/ast"
	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/adaptor"
	"github.com/labring/aiproxy/core/relay/adaptor/openai"
	"github.com/labring/aiproxy/core/relay/adaptor/registry"
	"github.com/labring/aiproxy/core/relay/meta"
	"github.com/labring/aiproxy/core/relay/mode"
	relaymodel "github.com/labring/aiproxy/core/relay/model"
	"github.com/labring/aiproxy/core/relay/utils"
)

type Adaptor struct {
	openai.Adaptor
	configCache utils.ChannelConfigCache[Config]
}

func init() {
	registry.Register(model.ChannelTypeQianfan, &Adaptor{})
}

const baseURL = "https://qianfan.baidubce.com/v2"

func (a *Adaptor) DefaultBaseURL() string {
	return baseURL
}

func (a *Adaptor) SupportMode(mt *meta.Meta) bool {
	m := adaptor.ModeFromMeta(mt)

	switch m {
	case mode.ChatCompletions,
		mode.Completions,
		mode.Anthropic,
		mode.Gemini,
		mode.Embeddings:
		return true
	case mode.Responses,
		mode.ResponsesGet,
		mode.ResponsesDelete,
		mode.ResponsesInputItems:
		return a.supportResponsesModel(mt)
	default:
		return false
	}
}

var builtinResponsesModels = map[string]struct{}{
	"deepseek-v3.2":                  {},
	"deepseek-v3.2-think":            {},
	"deepseek-v3.1-250821":           {},
	"deepseek-v3.1-think-250821":     {},
	"deepseek-v3":                    {},
	"deepseek-r1":                    {},
	"deepseek-r1-250528":             {},
	"kimi-k2-instruct":               {},
	"qwen3-coder-480b-a35b-instruct": {},
	"qwen3-coder-30b-a3b-instruct":   {},
	"qwen3-235b-a22b":                {},
	"qwen3-235b-a22b-thinking-2507":  {},
	"qwen3-235b-a22b-instruct-2507":  {},
	"qwen3-30b-a3b":                  {},
	"qwen3-30b-a3b-instruct-2507":    {},
	"qwen3-30b-a3b-thinking-2507":    {},
	"qwen3-32b":                      {},
	"qwen3-14b":                      {},
	"qwen3-8b":                       {},
	"qwen3-4b":                       {},
	"qwen3-1.7b":                     {},
	"qwen3-0.6b":                     {},
}

func (a *Adaptor) supportResponsesModel(mt *meta.Meta) bool {
	if qianfanModelMatches(mt, isBuiltinResponsesModel) {
		return true
	}

	if mt == nil {
		return false
	}

	cfg, err := a.loadConfig(mt)
	if err != nil {
		return false
	}

	return qianfanModelMatches(mt, func(modelName string) bool {
		return containsModel(cfg.ResponseModels, modelName)
	})
}

func qianfanModelMatches(mt *meta.Meta, match func(string) bool) bool {
	if mt == nil {
		return false
	}

	return utils.FirstMatchingModelName(mt.OriginModel, mt.ActualModel, match) != ""
}

func isBuiltinResponsesModel(modelName string) bool {
	_, ok := builtinResponsesModels[normalizeModelName(modelName)]
	return ok
}

func containsModel(models []string, modelName string) bool {
	target := normalizeModelName(modelName)
	if target == "" {
		return false
	}

	for _, m := range models {
		if normalizeModelName(m) == target {
			return true
		}
	}

	return false
}

func normalizeModelName(modelName string) string {
	return strings.ToLower(strings.TrimSpace(modelName))
}

func (a *Adaptor) GetRequestURL(
	meta *meta.Meta,
	store adaptor.Store,
	c *gin.Context,
) (adaptor.RequestURL, error) {
	if meta.Mode == mode.ResponsesCancel {
		return adaptor.RequestURL{}, fmt.Errorf("unsupported mode: %s", meta.Mode)
	}

	return a.Adaptor.GetRequestURL(meta, store, c)
}

func (a *Adaptor) ConvertRequest(
	meta *meta.Meta,
	store adaptor.Store,
	req *http.Request,
) (adaptor.ConvertResult, error) {
	switch meta.Mode {
	case mode.ResponsesCancel:
		return adaptor.ConvertResult{}, fmt.Errorf("unsupported mode: %s", meta.Mode)
	case mode.ChatCompletions:
		return openai.ConvertChatCompletionsRequest(
			meta,
			req,
			false,
			func(node *ast.Node) error {
				return patchReasoningFromNode(meta, node)
			},
		)
	case mode.Completions:
		return openai.ConvertCompletionsRequest(meta, req, func(node *ast.Node) error {
			return patchReasoningFromNode(meta, node)
		})
	case mode.Anthropic:
		return openai.ConvertClaudeRequest(
			meta,
			req,
			func(openAIReq *relaymodel.GeneralOpenAIRequest) error {
				return patchReasoningRequest(meta, openAIReq)
			},
		)
	case mode.Gemini:
		return openai.ConvertGeminiRequest(
			meta,
			req,
			func(openAIReq *relaymodel.GeneralOpenAIRequest) error {
				return patchReasoningRequest(meta, openAIReq)
			},
		)
	case mode.Responses:
		return openai.ConvertResponseRequest(meta, req, func(node *ast.Node) error {
			return patchResponsesReasoningFromNode(meta, node)
		})
	default:
		return a.Adaptor.ConvertRequest(meta, store, req)
	}
}

func patchReasoningFromNode(meta *meta.Meta, node *ast.Node) error {
	if node == nil {
		return nil
	}

	if node.Get("thinking").Exists() {
		_, _ = node.Unset("reasoning_effort")
		_, _ = node.Unset("enable_thinking")
		_, _ = node.Unset("thinking_budget")
		return nil
	}

	reasoning, err := utils.ParseOpenAIReasoningFromNode(node)
	if err != nil {
		return err
	}

	if !reasoning.Specified {
		return nil
	}

	return applyQianfanReasoningToNode(meta, node, reasoning)
}

func patchResponsesReasoningFromNode(meta *meta.Meta, node *ast.Node) error {
	if node == nil {
		return nil
	}

	if node.Get("thinking").Exists() {
		_, _ = node.Unset("reasoning")
		return nil
	}

	reasoningNode := node.Get("reasoning")
	if !reasoningNode.Exists() || reasoningNode.TypeSafe() == ast.V_NULL {
		return nil
	}

	effortNode := reasoningNode.Get("effort")
	if !effortNode.Exists() {
		return nil
	}

	effort, err := effortNode.String()
	if err != nil {
		return err
	}

	reasoning := parseReasoningEffort(effort)
	if !reasoning.Specified {
		return nil
	}

	return applyQianfanReasoningToNode(meta, node, reasoning)
}

func patchReasoningRequest(meta *meta.Meta, openAIReq *relaymodel.GeneralOpenAIRequest) error {
	if openAIReq == nil {
		return nil
	}

	if openAIReq.Thinking != nil {
		openAIReq.ReasoningEffort = nil
		openAIReq.EnableThinking = nil
		openAIReq.ThinkingBudget = nil
		return nil
	}

	reasoning := utils.ParseOpenAIReasoning(openAIReq)
	if !reasoning.Specified {
		return nil
	}

	openAIReq.ReasoningEffort = nil
	openAIReq.EnableThinking = nil
	openAIReq.ThinkingBudget = nil

	applyQianfanReasoningToRequest(meta, openAIReq, reasoning)

	return nil
}

func qianfanReasoningEffort(reasoning relaymodel.NormalizedReasoning) string {
	switch utils.ReasoningToOpenAIEffort(reasoning) {
	case relaymodel.ReasoningEffortXHigh:
		return "max"
	default:
		return "high"
	}
}

func applyQianfanReasoningToNode(
	meta *meta.Meta,
	node *ast.Node,
	reasoning relaymodel.NormalizedReasoning,
) error {
	_, _ = node.Unset("reasoning_effort")
	_, _ = node.Unset("enable_thinking")
	_, _ = node.Unset("thinking_budget")
	_, _ = node.Unset("thinking")
	_, _ = node.Unset("reasoning")

	disabled := reasoning.Disabled ||
		utils.ReasoningToOpenAIEffort(reasoning) == relaymodel.ReasoningEffortNone

	if qianfanModelMatches(meta, qianfanSupportsReasoningEffort) {
		if disabled {
			return nil
		}

		if _, err := node.Set(
			"reasoning_effort",
			ast.NewString(qianfanReasoningEffort(reasoning)),
		); err != nil {
			return err
		}

		if !qianfanModelMatches(meta, qianfanSupportsThinkingBudget) {
			return nil
		}

		return setQianfanThinkingBudget(meta, node, reasoning)
	}

	if qianfanModelMatches(meta, qianfanSupportsEnableThinking) {
		if _, err := node.Set("enable_thinking", ast.NewBool(!disabled)); err != nil {
			return err
		}

		if disabled || !qianfanModelMatches(meta, qianfanSupportsThinkingBudget) {
			return nil
		}

		return setQianfanThinkingBudget(meta, node, reasoning)
	}

	if qianfanModelMatches(meta, qianfanSupportsThinking) {
		thinkingType := relaymodel.ClaudeThinkingTypeEnabled
		if disabled {
			thinkingType = relaymodel.ClaudeThinkingTypeDisabled
		}

		if _, err := node.SetAny(
			"thinking",
			relaymodel.ClaudeThinking{Type: thinkingType},
		); err != nil {
			return err
		}

		if disabled || !qianfanModelMatches(meta, qianfanSupportsThinkingBudget) {
			return nil
		}

		return setQianfanThinkingBudget(meta, node, reasoning)
	}

	if disabled || !qianfanModelMatches(meta, qianfanSupportsThinkingBudget) {
		return nil
	}

	return setQianfanThinkingBudget(meta, node, reasoning)
}

func applyQianfanReasoningToRequest(
	meta *meta.Meta,
	req *relaymodel.GeneralOpenAIRequest,
	reasoning relaymodel.NormalizedReasoning,
) {
	req.ReasoningEffort = nil
	req.EnableThinking = nil
	req.ThinkingBudget = nil
	req.Thinking = nil

	disabled := reasoning.Disabled ||
		utils.ReasoningToOpenAIEffort(reasoning) == relaymodel.ReasoningEffortNone

	if qianfanModelMatches(meta, qianfanSupportsReasoningEffort) {
		if disabled {
			return
		}

		effort := qianfanReasoningEffort(reasoning)
		req.ReasoningEffort = &effort

		if qianfanModelMatches(meta, qianfanSupportsThinkingBudget) {
			budget := qianfanThinkingBudget(meta, reasoning)
			req.ThinkingBudget = &budget
		}

		return
	}

	if qianfanModelMatches(meta, qianfanSupportsEnableThinking) {
		enableThinking := !disabled
		req.EnableThinking = &enableThinking

		if !disabled && qianfanModelMatches(meta, qianfanSupportsThinkingBudget) {
			budget := qianfanThinkingBudget(meta, reasoning)
			req.ThinkingBudget = &budget
		}

		return
	}

	if qianfanModelMatches(meta, qianfanSupportsThinking) {
		thinkingType := relaymodel.ClaudeThinkingTypeEnabled
		if disabled {
			thinkingType = relaymodel.ClaudeThinkingTypeDisabled
		}

		req.Thinking = &relaymodel.ClaudeThinking{Type: thinkingType}

		if !disabled && qianfanModelMatches(meta, qianfanSupportsThinkingBudget) {
			budget := qianfanThinkingBudget(meta, reasoning)
			req.ThinkingBudget = &budget
		}

		return
	}

	if !disabled && qianfanModelMatches(meta, qianfanSupportsThinkingBudget) {
		budget := qianfanThinkingBudget(meta, reasoning)
		req.ThinkingBudget = &budget
	}
}

func setQianfanThinkingBudget(
	meta *meta.Meta,
	node *ast.Node,
	reasoning relaymodel.NormalizedReasoning,
) error {
	budget := qianfanThinkingBudget(meta, reasoning)
	_, err := node.Set("thinking_budget", ast.NewNumber(strconv.Itoa(budget)))
	return err
}

func qianfanThinkingBudget(
	_ *meta.Meta,
	reasoning relaymodel.NormalizedReasoning,
) int {
	budget := min(max(utils.ReasoningToBudget(reasoning), 100), 16384)

	return budget
}

func qianfanSupportsThinking(modelName string) bool {
	modelName = normalizeModelName(modelName)

	return strings.HasPrefix(modelName, "deepseek") ||
		strings.HasPrefix(modelName, "kimi-k2.5") ||
		strings.HasPrefix(modelName, "glm-5") ||
		strings.HasPrefix(modelName, "glm-4.7")
}

func qianfanSupportsEnableThinking(modelName string) bool {
	modelName = normalizeModelName(modelName)
	if strings.HasPrefix(modelName, "qwen3-") ||
		strings.HasPrefix(modelName, "ernie-5.0-thinking") ||
		strings.Contains(modelName, "vl") {
		return true
	}

	switch modelName {
	case "qwen3-235b-a22b",
		"qwen3-30b-a3b",
		"qwen3-32b",
		"qwen3-14b",
		"qwen3-8b",
		"qwen3-4b",
		"qwen3-1.7b",
		"qwen3-0.6b",
		"ernie-4.5-turbo-vl-preview",
		"ernie-4.5-turbo-vl-32k-preview",
		"ernie-4.5-vl-28b-a3b",
		"ernie-5.0-thinking-preview":
		return true
	default:
		return false
	}
}

func qianfanSupportsThinkingBudget(modelName string) bool {
	modelName = normalizeModelName(modelName)
	if strings.HasPrefix(modelName, "qwen3-") ||
		strings.HasPrefix(modelName, "ernie-5.0-thinking") ||
		strings.HasPrefix(modelName, "deepseek-v4-") ||
		strings.HasPrefix(modelName, "deepseek-r1") ||
		strings.Contains(modelName, "think") ||
		strings.Contains(modelName, "thinking") {
		return true
	}

	switch modelName {
	case "ernie-5.0-thinking-preview",
		"deepseek-v3.2-think",
		"deepseek-v3.1-250821",
		"deepseek-r1-250528",
		"qwen3-235b-a22b-thinking-2507",
		"qwen3-30b-a3b-thinking-2507",
		"qwen3-235b-a22b",
		"qwen3-30b-a3b",
		"qwen3-32b",
		"qwen3-14b",
		"qwen3-8b",
		"qwen3-4b",
		"qwen3-1.7b",
		"qwen3-0.6b":
		return true
	default:
		return false
	}
}

func qianfanSupportsReasoningEffort(modelName string) bool {
	return false
}

func parseReasoningEffort(effort string) relaymodel.NormalizedReasoning {
	normalizedEffort := relaymodel.NormalizeReasoningEffort(effort)
	if normalizedEffort == "" {
		return relaymodel.NormalizedReasoning{}
	}

	return relaymodel.NormalizedReasoning{
		Specified: true,
		Disabled:  normalizedEffort == relaymodel.ReasoningEffortNone,
		Effort:    normalizedEffort,
	}
}

func (a *Adaptor) SetupRequestHeader(
	meta *meta.Meta,
	store adaptor.Store,
	c *gin.Context,
	req *http.Request,
) error {
	if err := a.Adaptor.SetupRequestHeader(meta, store, c, req); err != nil {
		return err
	}

	cfg, err := a.loadConfig(meta)
	if err != nil {
		return err
	}

	if appID := strings.TrimSpace(cfg.AppID); appID != "" {
		req.Header.Set("Appid", appID)
	}

	return nil
}

func (a *Adaptor) DoResponse(
	meta *meta.Meta,
	store adaptor.Store,
	c *gin.Context,
	resp *http.Response,
) (adaptor.DoResponseResult, adaptor.Error) {
	if !adaptor.IsSuccessfulResponseStatus(meta.Mode, resp.StatusCode) {
		return adaptor.DoResponseResult{}, ErrorHandler(resp)
	}

	return a.Adaptor.DoResponse(meta, store, c, resp)
}

func (a *Adaptor) Metadata() adaptor.Metadata {
	return adaptor.Metadata{
		Readme:       "Baidu Qianfan OpenAI-compatible endpoint\nSupports chat, completions, embeddings, Responses API, Anthropic-compatible request conversion, and Gemini-compatible request conversion\nThinking controls: native `thinking` is preserved; normalized reasoning is written as `thinking`, `enable_thinking`, `thinking_budget`, or Qianfan `reasoning_effort` according to model capability.\nKey format example: `bce-v3/aaa/bbb`\nChannel config `appid` sets the upstream `appid` request header.",
		KeyHelp:      "bce-v3/aaa/bbb",
		ConfigSchema: configSchema(),
	}
}
