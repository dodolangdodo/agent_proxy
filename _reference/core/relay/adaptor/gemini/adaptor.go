package gemini

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/adaptor"
	"github.com/labring/aiproxy/core/relay/adaptor/registry"
	"github.com/labring/aiproxy/core/relay/meta"
	"github.com/labring/aiproxy/core/relay/mode"
	relaymodel "github.com/labring/aiproxy/core/relay/model"
	"github.com/labring/aiproxy/core/relay/utils"
)

type Adaptor struct {
	configCache utils.ChannelConfigCache[Config]
}

func init() {
	registry.Register(model.ChannelTypeGoogleGemini, &Adaptor{})
}

const baseURL = "https://generativelanguage.googleapis.com"

func (a *Adaptor) DefaultBaseURL() string {
	return baseURL
}

func (a *Adaptor) SupportMode(mt *meta.Meta) bool {
	m := adaptor.ModeFromMeta(mt)

	if m == mode.AudioSpeech {
		return isGeminiTTSMeta(mt)
	}

	if m == mode.ImagesGenerations || m == mode.ImagesEdits {
		return isGeminiImageMeta(mt)
	}

	return m == mode.ChatCompletions ||
		m == mode.Anthropic ||
		m == mode.Embeddings ||
		m == mode.Gemini ||
		m == mode.GeminiFiles ||
		m == mode.GeminiVideo ||
		m == mode.GeminiVideoOperations ||
		m == mode.GeminiTTS ||
		m == mode.GeminiImage ||
		m == mode.VideoGenerationsJobs ||
		m == mode.VideoGenerationsGetJobs ||
		m == mode.VideoGenerationsContent ||
		m == mode.Videos ||
		m == mode.VideosEdits ||
		m == mode.VideosExtensions ||
		m == mode.VideosGet ||
		m == mode.VideosContent
}

var v1ModelMap = map[string]struct{}{}

func requestVersionModel(meta *meta.Meta) string {
	if meta == nil {
		return ""
	}

	if modelName := utils.FirstMatchingModelName(
		meta.OriginModel,
		meta.ActualModel,
		func(modelName string) bool {
			_, ok := v1ModelMap[modelName]
			return ok
		},
	); modelName != "" {
		return modelName
	}

	return meta.ActualModel
}

func getRequestURL(meta *meta.Meta, action string) adaptor.RequestURL {
	u := meta.Channel.BaseURL
	if u == "" {
		u = baseURL
	}

	version := "v1beta"
	if _, ok := v1ModelMap[requestVersionModel(meta)]; ok {
		version = "v1"
	}

	return adaptor.RequestURL{
		Method: http.MethodPost,
		URL:    fmt.Sprintf("%s/%s/models/%s:%s", u, version, meta.ActualModel, action),
	}
}

func getOperationRequestURL(meta *meta.Meta, operationName string) (adaptor.RequestURL, error) {
	if operationName == "" {
		return adaptor.RequestURL{}, errors.New("operation name is empty")
	}

	u := meta.Channel.BaseURL
	if u == "" {
		u = baseURL
	}

	version := "v1beta"
	if _, ok := v1ModelMap[requestVersionModel(meta)]; ok {
		version = "v1"
	}

	requestURL, err := url.JoinPath(u, version, operationName)
	if err != nil {
		return adaptor.RequestURL{}, err
	}

	return adaptor.RequestURL{
		Method: http.MethodGet,
		URL:    requestURL,
	}, nil
}

func (a *Adaptor) GetRequestURL(
	meta *meta.Meta,
	store adaptor.Store,
	c *gin.Context,
) (adaptor.RequestURL, error) {
	var action string
	switch meta.Mode {
	case mode.Embeddings:
		action = "batchEmbedContents"
	case mode.GeminiVideo,
		mode.VideoGenerationsJobs,
		mode.Videos,
		mode.VideosEdits,
		mode.VideosExtensions:
		action = "predictLongRunning"
	case mode.GeminiVideoOperations:
		return getNativeVideoOperationRequestURL(meta, store)
	case mode.GeminiFiles:
		return getGeminiFileRequestURL(meta, store)
	case mode.VideoGenerationsGetJobs:
		operationID, err := ResolveVideoJobOperationID(meta, store, meta.JobID)
		if err != nil {
			return adaptor.RequestURL{}, err
		}

		return getOperationRequestURL(meta, operationID)
	case mode.VideoGenerationsContent:
		operationID, err := ResolveVideoGenerationOperationID(meta, store, meta.GenerationID)
		if err != nil {
			return adaptor.RequestURL{}, err
		}

		return getOperationRequestURL(meta, operationID)
	case mode.VideosGet:
		operationID, err := ResolveVideoGenerationOperationID(meta, store, meta.VideoID)
		if err != nil {
			return adaptor.RequestURL{}, err
		}

		return getOperationRequestURL(meta, operationID)
	case mode.VideosContent:
		operationID, err := ResolveVideoGenerationOperationID(meta, store, meta.VideoID)
		if err != nil {
			return adaptor.RequestURL{}, err
		}

		return getOperationRequestURL(meta, operationID)
	default:
		action = "generateContent"
	}

	if meta.GetBool("stream") ||
		(meta.Mode == mode.Gemini && c != nil && utils.IsGeminiStreamRequest(c.Request.URL.Path)) {
		action = "streamGenerateContent?alt=sse"
	}

	return getRequestURL(meta, action), nil
}

func (a *Adaptor) SetupRequestHeader(
	meta *meta.Meta,
	_ adaptor.Store,
	_ *gin.Context,
	req *http.Request,
) error {
	req.Header.Set("X-Goog-Api-Key", meta.Channel.Key)
	return nil
}

func (a *Adaptor) ConvertRequest(
	meta *meta.Meta,
	store adaptor.Store,
	req *http.Request,
) (adaptor.ConvertResult, error) {
	switch meta.Mode {
	case mode.Embeddings:
		return ConvertEmbeddingRequest(meta, req)
	case mode.ChatCompletions:
		return a.convertRequest(meta, req)
	case mode.Anthropic:
		return a.convertClaudeRequest(meta, req)
	case mode.Gemini:
		return NativeConvertRequest(meta, req)
	case mode.AudioSpeech:
		return ConvertTTSRequest(meta, req)
	case mode.ImagesGenerations:
		return ConvertImageRequest(meta, req)
	case mode.ImagesEdits:
		return ConvertImageEditRequest(meta, req)
	case mode.GeminiVideo:
		return NativeVideoConvertRequest(meta, req)
	case mode.GeminiVideoOperations:
		return ConvertVideoNoBodyRequest(meta, req)
	case mode.GeminiFiles:
		return ConvertVideoNoBodyRequest(meta, req)
	case mode.VideoGenerationsJobs:
		return ConvertVideoGenerationJobRequest(meta, req)
	case mode.Videos:
		return ConvertVideosRequest(meta, req)
	case mode.VideosEdits:
		return ConvertVideosEditRequestWithStore(meta, store, req)
	case mode.VideosExtensions:
		return ConvertVideosExtensionRequestWithStore(meta, store, req)
	case mode.VideoGenerationsGetJobs:
		return ConvertVideoGenerationsGetJobsRequest(meta, req)
	case mode.VideoGenerationsContent:
		return ConvertVideoGenerationsContentRequest(meta, req)
	case mode.VideosGet:
		return ConvertVideosGetRequest(meta, req)
	case mode.VideosContent:
		return ConvertVideosContentRequest(meta, req)
	case mode.VideosDelete:
		return ConvertVideoNoBodyRequest(meta, req)
	default:
		return adaptor.ConvertResult{}, fmt.Errorf("unsupported mode: %s", meta.Mode)
	}
}

func (a *Adaptor) DoRequest(
	meta *meta.Meta,
	_ adaptor.Store,
	_ *gin.Context,
	req *http.Request,
) (*http.Response, error) {
	return utils.DoRequestWithMeta(req, meta)
}

func (a *Adaptor) DoResponse(
	meta *meta.Meta,
	store adaptor.Store,
	c *gin.Context,
	resp *http.Response,
) (adaptor.DoResponseResult, adaptor.Error) {
	switch meta.Mode {
	case mode.Embeddings:
		return EmbeddingHandler(meta, c, resp)
	case mode.ChatCompletions:
		if utils.IsStreamResponse(resp) {
			return StreamHandler(meta, c, resp)
		}
		return Handler(meta, c, resp)
	case mode.Anthropic:
		if utils.IsStreamResponse(resp) {
			return ClaudeStreamHandler(meta, c, resp)
		}
		return ClaudeHandler(meta, c, resp)
	case mode.Gemini:
		// For Gemini mode (native format), pass through the response as-is
		if utils.IsStreamResponse(resp) {
			return NativeStreamHandler(meta, c, resp)
		}
		return NativeHandler(meta, c, resp)
	case mode.AudioSpeech:
		return TTSHandler(meta, c, resp)
	case mode.ImagesGenerations, mode.ImagesEdits:
		return ImageHandler(meta, c, resp)
	case mode.GeminiVideo:
		return NativeVideoHandler(meta, store, c, resp)
	case mode.GeminiVideoOperations:
		return NativeVideoOperationHandler(meta, store, c, resp)
	case mode.GeminiFiles:
		return GeminiFileHandler(meta, c, resp)
	case mode.VideoGenerationsJobs:
		return VideoGenerationJobSubmitHandler(meta, store, c, resp)
	case mode.Videos, mode.VideosEdits, mode.VideosExtensions:
		return VideosSubmitHandler(meta, store, c, resp)
	case mode.VideoGenerationsGetJobs:
		return VideoGenerationJobStatusHandler(meta, store, c, resp)
	case mode.VideoGenerationsContent:
		return VideoGenerationJobContentHandler(meta, c, resp)
	case mode.VideosGet:
		return VideosStatusHandler(meta, store, c, resp)
	case mode.VideosContent:
		return VideosContentHandler(meta, c, resp)
	default:
		return adaptor.DoResponseResult{}, relaymodel.WrapperOpenAIErrorWithMessage(
			fmt.Sprintf("unsupported mode: %s", meta.Mode),
			"unsupported_mode",
			http.StatusBadRequest,
		)
	}
}

func (a *Adaptor) Metadata() adaptor.Metadata {
	return adaptor.Metadata{
		Readme: "https://ai.google.dev\nGoogle Gemini native API\nSupports chat, embeddings, native Gemini requests, and image generation",
		Models: ModelList,
		ConfigSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"safety": map[string]any{
					"type":        "string",
					"title":       "Safety Threshold",
					"description": "Safety blocking threshold applied to all Gemini safety categories.",
					"enum": []string{
						relaymodel.GeminiSafetyThresholdBlockNone,
						relaymodel.GeminiSafetyThresholdBlockLowAndAbove,
						relaymodel.GeminiSafetyThresholdBlockMediumAndAbove,
						relaymodel.GeminiSafetyThresholdBlockOnlyHigh,
					},
				},
				"disable_auto_image_url_to_base64": map[string]any{
					"type":        "boolean",
					"title":       "Disable Auto Image URL To Base64",
					"description": "Keep image URLs unchanged instead of downloading and converting them to base64.",
				},
				"disable_auto_audio_url_to_base64": map[string]any{
					"type":        "boolean",
					"title":       "Disable Auto Audio URL To Base64",
					"description": "Keep audio URLs unchanged instead of downloading and converting them to base64 inline data.",
				},
				"disable_auto_video_url_to_base64": map[string]any{
					"type":        "boolean",
					"title":       "Disable Auto Video URL To Base64",
					"description": "Keep video URLs unchanged instead of downloading and converting them to base64 inline data.",
				},
				"enable_person_generation_allow_all": map[string]any{
					"type":        "boolean",
					"title":       "Enable Person Generation Allow All",
					"description": "When personGeneration is absent, set it to allow_all for Gemini image/video generation requests that support the field.",
				},
			},
		},
	}
}
