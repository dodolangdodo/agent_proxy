package vertexai

import (
	"bytes"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/relay/adaptor"
	"github.com/labring/aiproxy/core/relay/adaptor/gemini"
	"github.com/labring/aiproxy/core/relay/meta"
	"github.com/labring/aiproxy/core/relay/mode"
	"github.com/labring/aiproxy/core/relay/utils"
)

type Adaptor struct{}

func (a *Adaptor) ConvertRequest(
	meta *meta.Meta,
	_ adaptor.Store,
	request *http.Request,
) (adaptor.ConvertResult, error) {
	switch meta.Mode {
	case mode.Anthropic:
		return gemini.ConvertClaudeRequest(meta, request)
	case mode.Gemini:
		return gemini.NativeConvertRequest(meta, request, gemini.CleanFunctionResponseID)
	case mode.AudioSpeech:
		return gemini.ConvertTTSRequest(meta, request)
	case mode.ImagesGenerations:
		return gemini.ConvertImageRequest(meta, request)
	case mode.GeminiVideo:
		return convertNativeVideoRequest(meta, request)
	case mode.GeminiVideoOperations:
		return gemini.ConvertVideoNoBodyRequest(meta, request)
	case mode.GeminiFiles:
		return gemini.ConvertVideoNoBodyRequest(meta, request)
	case mode.VideoGenerationsJobs:
		return convertOpenAIVideoRequest(meta, request, gemini.ConvertVideoGenerationJobRequest)
	case mode.Videos:
		return convertOpenAIVideoRequest(meta, request, gemini.ConvertVideosRequest)
	case mode.VideoGenerationsGetJobs:
		return gemini.ConvertVideoGenerationsGetJobsRequest(meta, request)
	case mode.VideoGenerationsContent:
		return gemini.ConvertVideoGenerationsContentRequest(meta, request)
	case mode.VideosGet:
		return gemini.ConvertVideosGetRequest(meta, request)
	case mode.VideosContent:
		return gemini.ConvertVideosContentRequest(meta, request)
	case mode.VideosDelete:
		return gemini.ConvertVideoNoBodyRequest(meta, request)
	default:
		return gemini.ConvertRequest(meta, request)
	}
}

func convertOpenAIVideoRequest(
	meta *meta.Meta,
	request *http.Request,
	convert func(*meta.Meta, *http.Request) (adaptor.ConvertResult, error),
) (adaptor.ConvertResult, error) {
	result, err := convert(meta, request)
	if err != nil {
		return adaptor.ConvertResult{}, err
	}

	body, err := io.ReadAll(result.Body)
	if err != nil {
		return adaptor.ConvertResult{}, err
	}

	body, err = gemini.ConvertVideoRequestParametersToVertex(body)
	if err != nil {
		return adaptor.ConvertResult{}, err
	}

	result.Body = bytes.NewReader(body)
	result.Header.Set("Content-Length", strconv.Itoa(len(body)))

	return result, nil
}

func convertNativeVideoRequest(
	meta *meta.Meta,
	request *http.Request,
) (adaptor.ConvertResult, error) {
	result, err := gemini.NativeVideoConvertRequest(meta, request)
	if err != nil {
		return adaptor.ConvertResult{}, err
	}

	body, err := io.ReadAll(result.Body)
	if err != nil {
		return adaptor.ConvertResult{}, err
	}

	body, err = geminiVideoRequestBodyToVertex(body)
	if err != nil {
		return adaptor.ConvertResult{}, err
	}

	result.Body = bytes.NewReader(body)
	result.Header.Set("Content-Length", strconv.Itoa(len(body)))

	return result, nil
}

func geminiVideoRequestBodyToVertex(body []byte) ([]byte, error) {
	return gemini.ConvertVideoRequestParametersToVertex(body)
}

func (a *Adaptor) SetupRequestHeader(
	_ *meta.Meta,
	_ adaptor.Store,
	_ *gin.Context,
	_ *http.Request,
) error {
	return nil
}

func (a *Adaptor) DoResponse(
	meta *meta.Meta,
	store adaptor.Store,
	c *gin.Context,
	resp *http.Response,
) (adaptor.DoResponseResult, adaptor.Error) {
	switch meta.Mode {
	case mode.Anthropic:
		if utils.IsStreamResponse(resp) {
			return gemini.ClaudeStreamHandler(meta, c, resp)
		}
		return gemini.ClaudeHandler(meta, c, resp)
	case mode.Gemini:
		// For Gemini mode (native format), pass through the response as-is
		if utils.IsStreamResponse(resp) {
			return gemini.NativeStreamHandler(meta, c, resp)
		}
		return gemini.NativeHandler(meta, c, resp)
	case mode.AudioSpeech:
		return gemini.TTSHandler(meta, c, resp)
	case mode.ImagesGenerations:
		return gemini.ImageHandler(meta, c, resp)
	case mode.GeminiVideo:
		return gemini.NativeVideoHandler(meta, store, c, resp)
	case mode.GeminiVideoOperations:
		return gemini.NativeVideoOperationHandler(meta, store, c, resp)
	case mode.GeminiFiles:
		return gemini.GeminiFileHandler(meta, c, resp)
	case mode.VideoGenerationsJobs:
		return gemini.VideoGenerationJobSubmitHandler(meta, store, c, resp)
	case mode.Videos:
		return gemini.VideosSubmitHandler(meta, store, c, resp)
	case mode.VideoGenerationsGetJobs:
		return gemini.VideoGenerationJobStatusHandler(meta, store, c, resp)
	case mode.VideoGenerationsContent:
		return gemini.VideoGenerationJobContentHandler(meta, c, resp)
	case mode.VideosGet:
		return gemini.VideosStatusHandler(meta, store, c, resp)
	case mode.VideosContent:
		return gemini.VideosContentHandler(meta, c, resp)
	default:
		if utils.IsStreamResponse(resp) {
			return gemini.StreamHandler(meta, c, resp)
		}
		return gemini.Handler(meta, c, resp)
	}
}
