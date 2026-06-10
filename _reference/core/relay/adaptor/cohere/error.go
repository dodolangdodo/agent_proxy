package cohere

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/labring/aiproxy/core/common"
	"github.com/labring/aiproxy/core/relay/adaptor"
	relaymodel "github.com/labring/aiproxy/core/relay/model"
)

func ErrorHandler(resp *http.Response) adaptor.Error {
	defer resp.Body.Close()

	respBody, err := common.GetResponseBody(resp)
	if err != nil {
		return relaymodel.WrapperOpenAIErrorWithMessage(
			err.Error(),
			relaymodel.ErrorCodeBadResponse,
			resp.StatusCode,
			relaymodel.ErrorTypeUpstream,
		)
	}

	return ErrorHandlerWithBody(resp.StatusCode, respBody)
}

func ErrorHandlerWithBody(statusCode int, respBody []byte) adaptor.Error {
	openAIError := relaymodel.OpenAIError{
		Type: relaymodel.ErrorTypeUpstream,
	}

	var errResponse Response
	if err := sonic.Unmarshal(respBody, &errResponse); err != nil {
		openAIError.Message = string(respBody)
		openAIError.Code = relaymodel.ErrorCodeBadResponse
		return relaymodel.NewOpenAIError(statusCode, openAIError)
	}

	openAIError.Message = errResponse.Message
	openAIError.Code = statusCode

	if openAIError.Message == "" {
		openAIError.Message = string(respBody)
	}

	return relaymodel.NewOpenAIError(statusCode, openAIError)
}
