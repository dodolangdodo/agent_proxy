package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/adaptor/openai"
	"github.com/labring/aiproxy/core/relay/utils"
)

func GetEmbedRequestUsage(c *gin.Context, _ model.ModelConfig) (RequestUsage, error) {
	textRequest, err := utils.UnmarshalGeneralOpenAIRequest(c.Request)
	if err != nil {
		return RequestUsage{}, err
	}

	return NewRequestUsage(model.Usage{
		InputTokens: model.ZeroNullInt64(openai.CountTokenInput(
			textRequest.Input,
			textRequest.Model,
		)),
	}), nil
}
