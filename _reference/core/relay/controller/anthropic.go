package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/adaptor/openai"
	"github.com/labring/aiproxy/core/relay/utils"
)

func GetAnthropicRequestUsage(c *gin.Context, _ model.ModelConfig) (RequestUsage, error) {
	textRequest, err := utils.UnmarshalAnthropicMessageRequest(c.Request)
	if err != nil {
		return RequestUsage{}, err
	}

	return NewRequestUsage(model.Usage{
		InputTokens: model.ZeroNullInt64(openai.CountTokenMessages(
			textRequest.Messages,
			textRequest.Model,
			false,
		)),
	}), nil
}
