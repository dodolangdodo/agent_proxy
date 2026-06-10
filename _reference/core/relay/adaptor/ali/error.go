package ali

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/labring/aiproxy/core/common"
	"github.com/labring/aiproxy/core/relay/adaptor"
	"github.com/labring/aiproxy/core/relay/meta"
	relaymodel "github.com/labring/aiproxy/core/relay/model"
)

// https://help.aliyun.com/zh/model-studio/error-code?userCode=okjhlpr5

func ErrorHanlder(resp *http.Response) adaptor.Error {
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

	statusCode, openAIError := getAliErrorWithBody(resp.StatusCode, respBody)
	statusCode, openAIError = normalizeAliError(statusCode, openAIError)

	return relaymodel.NewOpenAIError(statusCode, openAIError)
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

func OpenAIVideoErrorHandlerWithBody(statusCode int, respBody []byte) adaptor.Error {
	statusCode, openAIError := getAliErrorWithBody(statusCode, respBody)
	statusCode, openAIError = normalizeAliError(statusCode, openAIError)

	return relaymodel.NewOpenAIVideoError(statusCode, relaymodel.OpenAIVideoError{
		Detail: openAIError.Message,
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

func normalizeAliError(
	statusCode int,
	openAIError relaymodel.OpenAIError,
) (int, relaymodel.OpenAIError) {
	// {"error":{"code":"ServiceUnavailable","message":"<503> InternalError.Algo: An error occurred in model serving, error message is: [Too many requests. Your requests are being throttled due to system capacity limits. Please try again later.]","type":"ServiceUnavailable"}}
	switch openAIError.Type {
	case "ServiceUnavailable":
		statusCode = http.StatusServiceUnavailable
		openAIError.Type = relaymodel.ErrorTypeUpstream
	case "RequestTimeOut":
		statusCode = http.StatusRequestTimeout
		openAIError.Type = relaymodel.ErrorTypeUpstream
	}

	if strings.Contains(openAIError.Message, "object is not iterable") {
		statusCode = http.StatusBadRequest
	}

	return statusCode, openAIError
}

func getAliErrorWithBody(statusCode int, respBody []byte) (int, relaymodel.OpenAIError) {
	openAIError := relaymodel.OpenAIError{
		Type:  relaymodel.ErrorTypeUpstream,
		Code:  relaymodel.ErrorCodeBadResponse,
		Param: strconv.Itoa(statusCode),
	}

	var errResponse struct {
		Error     relaymodel.OpenAIError `json:"error"`
		Code      string                 `json:"code"`
		Message   string                 `json:"message"`
		RequestID string                 `json:"request_id"`
	}

	if err := sonic.Unmarshal(respBody, &errResponse); err != nil {
		openAIError.Message = string(respBody)
		return statusCode, openAIError
	}

	if errResponse.Error.Message != "" {
		openAIError = errResponse.Error
	} else if errResponse.Message != "" || errResponse.Code != "" {
		openAIError.Message = firstNonEmpty(errResponse.Message, errResponse.Code)
		openAIError.Code = firstNonEmpty(errResponse.Code, relaymodel.ErrorCodeBadResponse)
	}

	if openAIError.Message == "" {
		openAIError.Message = string(respBody)
	}

	return statusCode, openAIError
}
