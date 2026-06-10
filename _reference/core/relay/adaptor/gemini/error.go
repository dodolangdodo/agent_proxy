package gemini

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/labring/aiproxy/core/common"
	"github.com/labring/aiproxy/core/relay/adaptor"
	"github.com/labring/aiproxy/core/relay/meta"
	relaymodel "github.com/labring/aiproxy/core/relay/model"
)

func ErrorHandler(resp *http.Response) adaptor.Error {
	defer resp.Body.Close()

	respBody, err := common.GetResponseBody(resp)
	if err != nil {
		return relaymodel.NewGeminiError(resp.StatusCode, relaymodel.GeminiError{
			Message: err.Error(),
			Status:  relaymodel.ErrorTypeUpstream,
			Code:    resp.StatusCode,
		})
	}

	return ErrorHandlerWithBody(resp.StatusCode, respBody)
}

func OpenAIVideoErrorHandler(resp *http.Response) adaptor.Error {
	defer resp.Body.Close()

	respBody, err := common.GetResponseBody(resp)
	if err != nil {
		return relaymodel.NewOpenAIVideoError(resp.StatusCode, relaymodel.OpenAIVideoError{
			Detail: err.Error(),
		})
	}

	return OpenAIVideoErrorHandlerWithBody(resp.StatusCode, respBody)
}

func ErrorHandlerWithBody(statusCode int, respBody []byte) adaptor.Error {
	geminiError := parseGeminiError(statusCode, respBody)
	return relaymodel.NewGeminiError(statusCode, geminiError)
}

func OpenAIVideoErrorHandlerWithBody(statusCode int, respBody []byte) adaptor.Error {
	geminiError := parseGeminiError(statusCode, respBody)
	return relaymodel.NewOpenAIVideoError(statusCode, relaymodel.OpenAIVideoError{
		Detail: geminiError.Message,
	})
}

func convertRequestError(meta *meta.Meta, message string) adaptor.Error {
	if meta == nil {
		return relaymodel.WrapperOpenAIErrorWithMessage(
			message,
			"invalid_request_error",
			http.StatusBadRequest,
		)
	}

	return relaymodel.WrapperErrorWithMessage(
		meta.Mode,
		http.StatusBadRequest,
		message,
		relaymodel.WithCode("invalid_request_error"),
	)
}

func parseGeminiError(statusCode int, respBody []byte) relaymodel.GeminiError {
	var errResponse relaymodel.GeminiErrorResponse

	err := sonic.Unmarshal(respBody, &errResponse)
	if err != nil {
		// Maybe it's not a JSON response or different format
		return relaymodel.GeminiError{
			Message: string(respBody),
			Status:  relaymodel.ErrorTypeUpstream,
			Code:    statusCode,
		}
	}

	if errResponse.Error.Message == "" {
		errResponse.Error.Message = string(respBody)
	}

	if errResponse.Error.Code == 0 {
		errResponse.Error.Code = statusCode
	}

	return errResponse.Error
}
