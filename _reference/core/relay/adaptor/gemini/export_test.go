package gemini

import (
	"context"
	"net/http"

	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/adaptor"
	"github.com/labring/aiproxy/core/relay/meta"
	"github.com/labring/aiproxy/core/relay/mode"
	relaymodel "github.com/labring/aiproxy/core/relay/model"
)

func BuildMessagePartForTest(
	message relaymodel.MessageContent,
) *relaymodel.GeminiPart {
	return buildMessageParts(message)
}

func ProcessImageTasksForTest(ctx context.Context, imageTasks []*relaymodel.GeminiPart) error {
	return processImageTasks(ctx, imageTasks)
}

func ProcessMediaTasksForTest(
	ctx context.Context,
	mediaType string,
	mediaTasks []*relaymodel.GeminiPart,
) {
	processMediaTasks(ctx, mediaType, mediaTasks)
}

func ResponseChat2OpenAIForTest(
	meta *meta.Meta,
	response *relaymodel.GeminiChatResponse,
) *relaymodel.TextResponse {
	return responseChat2OpenAI(meta, response)
}

func StreamResponseChat2OpenAIForTest(
	meta *meta.Meta,
	response *relaymodel.GeminiChatResponse,
) *relaymodel.ChatCompletionsStreamResponse {
	return streamResponseChat2OpenAI(meta, response)
}

func ConvertVideoRequestForTest(
	meta *meta.Meta,
	req *http.Request,
) (adaptor.ConvertResult, error) {
	switch meta.Mode {
	case mode.VideoGenerationsJobs:
		return ConvertVideoGenerationJobRequest(meta, req)
	case mode.Videos:
		return ConvertVideosRequest(meta, req)
	case mode.VideosEdits:
		return ConvertVideosEditRequest(meta, req)
	case mode.VideosExtensions:
		return ConvertVideosExtensionRequest(meta, req)
	default:
		return ConvertVideoGenerationJobRequest(meta, req)
	}
}

func ConvertRequestWithConfigForTest(
	meta *meta.Meta,
	req *http.Request,
	cfg Config,
) (adaptor.ConvertResult, error) {
	return convertRequest(meta, req, cfg)
}

func ConvertVideoRequestWithConfigForTest(
	meta *meta.Meta,
	req *http.Request,
	cfg Config,
) (adaptor.ConvertResult, error) {
	return convertOpenAIVideoRequestWithConfig(meta, req, cfg)
}

func GeminiImageAspectRatioFromSizeForTest(size string) string {
	return geminiImageAspectRatioFromSize(size)
}

func GeminiVideoURLByIDForTest(operation *relaymodel.GeminiVideoOperation, id string) string {
	return geminiVideoURLByID(operation, id)
}

func GeminiVideoLocalIDForTest(operationName string) string {
	return geminiVideoLocalID(operationName)
}

func GeminiVideoRequestUsageForTest(meta *meta.Meta) model.Usage {
	return geminiVideoRequestUsage(meta)
}

func GeminiVideoUsageContextForTest(meta *meta.Meta) model.UsageContext {
	return geminiVideoUsageContext(meta, nil)
}

func GeminiVideoStoreMetadataForTest(value string) string {
	return parseGeminiVideoStoreMetadata(value).OperationName
}
