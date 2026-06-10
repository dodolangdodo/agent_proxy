package doubao

import (
	"net/http"
	"strconv"

	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/labring/aiproxy/core/common"
	"github.com/labring/aiproxy/core/relay/adaptor"
	"github.com/labring/aiproxy/core/relay/meta"
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

	statusCode, openAIError := getDoubaoErrorWithBody(resp.StatusCode, respBody)

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
	_, openAIError := getDoubaoErrorWithBody(statusCode, respBody)

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

func getDoubaoErrorWithBody(statusCode int, respBody []byte) (int, relaymodel.OpenAIError) {
	openAIError := relaymodel.OpenAIError{
		Type:  relaymodel.ErrorTypeUpstream,
		Code:  relaymodel.ErrorCodeBadResponse,
		Param: strconv.Itoa(statusCode),
	}

	root, err := sonic.Get(respBody)
	if err != nil {
		openAIError.Message = string(respBody)
		return statusCode, openAIError
	}

	if errNode := root.Get(
		"error",
	); errNode != nil && errNode.Exists() &&
		errNode.TypeSafe() == ast.V_OBJECT {
		var errResponse relaymodel.OpenAIErrorResponse
		if err := sonic.Unmarshal(
			respBody,
			&errResponse,
		); err == nil &&
			errResponse.Error.Message != "" {
			openAIError = errResponse.Error
		} else {
			openAIError.Message = stringFromNode(errNode.Get("message"))
			openAIError.Type = firstNonEmptyString(
				stringFromNode(errNode.Get("type")),
				openAIError.Type,
			)

			openAIError.Param = firstNonEmptyString(
				stringFromNode(errNode.Get("param")),
				openAIError.Param,
			)
			if codeNode := errNode.Get("code"); codeNode != nil && codeNode.Exists() {
				openAIError.Code = anyFromNode(codeNode)
			}
		}
	} else if responseMetadataErr := root.GetByPath(
		"ResponseMetadata",
		"Error",
	); responseMetadataErr != nil &&
		responseMetadataErr.Exists() &&
		responseMetadataErr.TypeSafe() == ast.V_OBJECT {
		openAIError.Message = stringFromNode(responseMetadataErr.Get("Message"))
		if code := stringFromNode(responseMetadataErr.Get("Code")); code != "" {
			openAIError.Code = code
		}
	}

	if openAIError.Message == "" {
		openAIError.Message = string(respBody)
	}

	return statusCode, openAIError
}

func stringFromNode(node *ast.Node) string {
	if node == nil || !node.Exists() || node.TypeSafe() == ast.V_NULL {
		return ""
	}

	value, err := node.String()
	if err != nil {
		return ""
	}

	return value
}

func anyFromNode(node *ast.Node) any {
	if node == nil || !node.Exists() || node.TypeSafe() == ast.V_NULL {
		return nil
	}

	switch node.TypeSafe() {
	case ast.V_STRING:
		return stringFromNode(node)
	case ast.V_NUMBER:
		if value, err := node.Int64(); err == nil {
			return value
		}

		if value, err := node.Float64(); err == nil {
			return value
		}
	case ast.V_TRUE, ast.V_FALSE:
		if value, err := node.Bool(); err == nil {
			return value
		}
	}

	return relaymodel.ErrorCodeBadResponse
}
