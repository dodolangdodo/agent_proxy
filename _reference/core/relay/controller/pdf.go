package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/model"
)

func GetPdfRequestUsage(_ *gin.Context, _ model.ModelConfig) (RequestUsage, error) {
	return NewRequestUsage(model.Usage{}), nil
}
