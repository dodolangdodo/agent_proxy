package doubao

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

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

func init() {
	registry.Register(model.ChannelTypeDoubao, &Adaptor{})
}

func featureModel(meta *meta.Meta) string {
	if meta == nil {
		return ""
	}

	if modelName := utils.FirstMatchingModelName(
		meta.OriginModel,
		meta.ActualModel,
		func(modelName string) bool {
			modelName = strings.ToLower(modelName)
			return strings.HasPrefix(modelName, "bot-") || strings.Contains(modelName, "vision")
		},
	); modelName != "" {
		return modelName
	}

	return utils.PreferredModelName(meta.OriginModel, meta.ActualModel)
}

func GetRequestURL(meta *meta.Meta) (adaptor.RequestURL, error) {
	u := meta.Channel.BaseURL

	modelName := strings.ToLower(featureModel(meta))
	switch meta.Mode {
	case mode.ChatCompletions, mode.Anthropic, mode.Gemini:
		if strings.HasPrefix(modelName, "bot-") {
			url, err := url.JoinPath(u, "/api/v3/bots/chat/completions")
			if err != nil {
				return adaptor.RequestURL{}, err
			}

			return adaptor.RequestURL{
				Method: http.MethodPost,
				URL:    url,
			}, nil
		}

		url, err := url.JoinPath(u, "/api/v3/chat/completions")
		if err != nil {
			return adaptor.RequestURL{}, err
		}

		return adaptor.RequestURL{
			Method: http.MethodPost,
			URL:    url,
		}, nil
	case mode.Embeddings:
		if strings.Contains(modelName, "vision") {
			url, err := url.JoinPath(u, "/api/v3/embeddings/multimodal")
			if err != nil {
				return adaptor.RequestURL{}, err
			}

			return adaptor.RequestURL{
				Method: http.MethodPost,
				URL:    url,
			}, nil
		}

		url, err := url.JoinPath(u, "/api/v3/embeddings")
		if err != nil {
			return adaptor.RequestURL{}, err
		}

		return adaptor.RequestURL{
			Method: http.MethodPost,
			URL:    url,
		}, nil
	case mode.ImagesGenerations:
		url, err := url.JoinPath(u, "/api/v3/images/generations")
		if err != nil {
			return adaptor.RequestURL{}, err
		}

		return adaptor.RequestURL{
			Method: http.MethodPost,
			URL:    url,
		}, nil
	case mode.DoubaoVideo,
		mode.DoubaoVideoTasks,
		mode.DoubaoVideoTasksDelete,
		mode.VideoGenerationsJobs,
		mode.Videos,
		mode.VideosEdits,
		mode.VideosExtensions,
		mode.VideoGenerationsGetJobs,
		mode.VideoGenerationsContent,
		mode.VideosGet,
		mode.VideosContent,
		mode.VideosDelete:
		return doubaoVideoTaskRequestURL(u, meta)
	case mode.Responses:
		url, err := url.JoinPath(u, "/api/v3/responses")
		if err != nil {
			return adaptor.RequestURL{}, err
		}

		return adaptor.RequestURL{
			Method: http.MethodPost,
			URL:    url,
		}, nil
	case mode.ResponsesGet:
		url, err := url.JoinPath(u, "/api/v3/responses", meta.ResponseID)
		if err != nil {
			return adaptor.RequestURL{}, err
		}

		return adaptor.RequestURL{
			Method: http.MethodGet,
			URL:    url,
		}, nil
	case mode.ResponsesDelete:
		url, err := url.JoinPath(u, "/api/v3/responses", meta.ResponseID)
		if err != nil {
			return adaptor.RequestURL{}, err
		}

		return adaptor.RequestURL{
			Method: http.MethodDelete,
			URL:    url,
		}, nil
	case mode.ResponsesCancel:
		url, err := url.JoinPath(u, "/api/v3/responses", meta.ResponseID, "cancel")
		if err != nil {
			return adaptor.RequestURL{}, err
		}

		return adaptor.RequestURL{
			Method: http.MethodPost,
			URL:    url,
		}, nil
	case mode.ResponsesInputItems:
		url, err := url.JoinPath(u, "/api/v3/responses", meta.ResponseID, "input_items")
		if err != nil {
			return adaptor.RequestURL{}, err
		}

		return adaptor.RequestURL{
			Method: http.MethodGet,
			URL:    url,
		}, nil
	default:
		return adaptor.RequestURL{}, fmt.Errorf("unsupported relay mode %d for doubao", meta.Mode)
	}
}

func doubaoVideoTaskRequestURL(baseURL string, meta *meta.Meta) (adaptor.RequestURL, error) {
	switch meta.Mode {
	case mode.DoubaoVideo,
		mode.VideoGenerationsJobs,
		mode.Videos,
		mode.VideosEdits,
		mode.VideosExtensions:
		return doubaoTaskRequestURL(baseURL, http.MethodPost, "")
	case mode.DoubaoVideoTasks, mode.VideosGet, mode.VideosContent:
		return doubaoTaskRequestURL(baseURL, http.MethodGet, meta.VideoID)
	case mode.DoubaoVideoTasksDelete, mode.VideosDelete:
		return doubaoTaskRequestURL(baseURL, http.MethodDelete, meta.VideoID)
	case mode.VideoGenerationsGetJobs:
		return doubaoTaskRequestURL(baseURL, http.MethodGet, meta.JobID)
	case mode.VideoGenerationsContent:
		return doubaoTaskRequestURL(baseURL, http.MethodGet, meta.GenerationID)
	default:
		return adaptor.RequestURL{}, fmt.Errorf(
			"unsupported relay mode %d for doubao video",
			meta.Mode,
		)
	}
}

func doubaoTaskRequestURL(baseURL, method, taskID string) (adaptor.RequestURL, error) {
	parts := []string{"/api/v3/contents/generations/tasks"}
	if taskID != "" {
		parts = append(parts, taskID)
	}

	url, err := url.JoinPath(baseURL, parts...)
	if err != nil {
		return adaptor.RequestURL{}, err
	}

	return adaptor.RequestURL{
		Method: method,
		URL:    url,
	}, nil
}

type Adaptor struct {
	openai.Adaptor
}

const baseURL = "https://ark.cn-beijing.volces.com"

func (a *Adaptor) DefaultBaseURL() string {
	return baseURL
}

func (a *Adaptor) SupportMode(mt *meta.Meta) bool {
	m := adaptor.ModeFromMeta(mt)

	return m == mode.ChatCompletions ||
		m == mode.Anthropic ||
		m == mode.Gemini ||
		m == mode.Embeddings ||
		m == mode.ImagesGenerations ||
		m == mode.DoubaoVideo ||
		m == mode.DoubaoVideoTasks ||
		m == mode.DoubaoVideoTasksDelete ||
		m == mode.VideoGenerationsJobs ||
		m == mode.VideoGenerationsGetJobs ||
		m == mode.VideoGenerationsContent ||
		m == mode.Videos ||
		m == mode.VideosEdits ||
		m == mode.VideosExtensions ||
		m == mode.VideosGet ||
		m == mode.VideosContent ||
		m == mode.VideosDelete ||
		m == mode.Responses ||
		m == mode.ResponsesGet ||
		m == mode.ResponsesDelete ||
		m == mode.ResponsesCancel ||
		m == mode.ResponsesInputItems
}

func (a *Adaptor) Metadata() adaptor.Metadata {
	return adaptor.Metadata{
		Readme: "Doubao / Volcano Engine endpoint\nSupports bot-style models, native Responses API, Gemini-compatible request conversion, and network search metering fields",
		Models: ModelList,
	}
}

func (a *Adaptor) GetRequestURL(
	meta *meta.Meta,
	_ adaptor.Store,
	_ *gin.Context,
) (adaptor.RequestURL, error) {
	return GetRequestURL(meta)
}

func (a *Adaptor) ConvertRequest(
	meta *meta.Meta,
	store adaptor.Store,
	req *http.Request,
) (adaptor.ConvertResult, error) {
	reasoningHook := func(openAIReq *relaymodel.GeneralOpenAIRequest) error {
		reasoning := utils.ParseOpenAIReasoning(openAIReq)
		utils.ApplyReasoningToDoubaoRequest(openAIReq, reasoning)
		return nil
	}

	switch meta.Mode {
	case mode.ImagesGenerations:
		return ConvertImageRequest(meta, req)
	case mode.DoubaoVideo:
		return ConvertDoubaoNativeVideoRequest(meta, req)
	case mode.DoubaoVideoTasks, mode.DoubaoVideoTasksDelete:
		return adaptor.ConvertResult{}, nil
	case mode.VideoGenerationsJobs:
		return ConvertVideoGenerationJobRequest(meta, req)
	case mode.Videos:
		return ConvertVideosRequest(meta, req)
	case mode.VideosEdits:
		return ConvertVideosEditRequest(meta, req)
	case mode.VideosExtensions:
		return ConvertVideosExtensionRequest(meta, req)
	case mode.VideoGenerationsGetJobs:
		return ConvertVideoGenerationsGetJobsRequest(meta, req)
	case mode.VideoGenerationsContent:
		return ConvertVideoGenerationsContentRequest(meta, req)
	case mode.VideosGet:
		return ConvertVideosGetRequest(meta, req)
	case mode.VideosContent:
		return ConvertVideosContentRequest(meta, req)
	case mode.VideosDelete:
		return adaptor.ConvertResult{}, nil
	case mode.Embeddings:
		if strings.Contains(strings.ToLower(featureModel(meta)), "vision") {
			return openai.ConvertEmbeddingsRequest(meta, req, false, patchEmbeddingsVisionInput)
		}
		return openai.ConvertEmbeddingsRequest(meta, req, true)
	case mode.ChatCompletions:
		return ConvertChatCompletionsRequest(meta, req)
	case mode.Anthropic:
		return openai.ConvertClaudeRequest(meta, req, reasoningHook)
	case mode.Gemini:
		return openai.ConvertGeminiRequest(meta, req, reasoningHook)
	default:
		return openai.ConvertRequest(meta, store, req)
	}
}

func (a *Adaptor) DoResponse(
	meta *meta.Meta,
	store adaptor.Store,
	c *gin.Context,
	resp *http.Response,
) (adaptor.DoResponseResult, adaptor.Error) {
	switch meta.Mode {
	case mode.ImagesGenerations:
		if utils.IsStreamResponse(resp) {
			return ImageStreamHandler(meta, c, resp)
		}

		return ImageHandler(meta, c, resp)
	case mode.DoubaoVideo:
		return DoubaoNativeVideoSubmitHandler(meta, store, c, resp)
	case mode.DoubaoVideoTasks:
		return DoubaoNativeVideoTaskHandler(meta, store, c, resp)
	case mode.DoubaoVideoTasksDelete:
		return DoubaoNativeVideoTaskDeleteHandler(meta, c, resp)
	case mode.VideoGenerationsJobs:
		return VideoGenerationJobSubmitHandler(meta, store, c, resp)
	case mode.Videos:
		return VideosSubmitHandler(meta, store, c, resp)
	case mode.VideosEdits, mode.VideosExtensions:
		return VideosSubmitHandler(meta, store, c, resp)
	case mode.VideoGenerationsGetJobs:
		return VideoGenerationJobStatusHandler(meta, store, c, resp)
	case mode.VideoGenerationsContent:
		return VideoGenerationJobContentHandler(meta, c, resp)
	case mode.VideosGet:
		return VideosStatusHandler(meta, store, c, resp)
	case mode.VideosContent:
		return VideosContentHandler(meta, c, resp)
	case mode.VideosDelete:
		return openai.VideoDeleteHandler(meta, c, resp)
	case mode.ChatCompletions:
		websearchCount := int64(0)

		var (
			result adaptor.DoResponseResult
			err    adaptor.Error
		)

		if utils.IsStreamResponse(resp) {
			result, err = openai.StreamHandler(meta, c, resp, newHandlerPreHandler(&websearchCount))
		} else {
			result, err = openai.Handler(meta, c, resp, newHandlerPreHandler(&websearchCount))
		}

		result.Usage.WebSearchCount += model.ZeroNullInt64(websearchCount)

		return result, err
	case mode.Embeddings:
		return openai.EmbeddingsHandler(
			meta,
			c,
			resp,
			embeddingPreHandler,
		)
	case mode.Gemini:
		if utils.IsStreamResponse(resp) {
			return openai.GeminiStreamHandler(meta, c, resp)
		}
		return openai.GeminiHandler(meta, c, resp)
	default:
		return openai.DoResponse(meta, store, c, resp)
	}
}

func (a *Adaptor) GetBalance(_ *model.Channel) (float64, error) {
	return 0, adaptor.ErrGetBalanceNotImplemented
}
