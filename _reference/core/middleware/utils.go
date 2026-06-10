package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/common"
	"github.com/labring/aiproxy/core/relay/mode"
	relaymodel "github.com/labring/aiproxy/core/relay/model"
)

func AbortLogWithMessageWithMode(
	m mode.Mode,
	c *gin.Context,
	statusCode int,
	message string,
	opts ...relaymodel.WrapperErrorOptionFunc,
) {
	common.GetLogger(c).Error(message)
	AbortWithMessageWithMode(m, c, statusCode, message, opts...)
}

func AbortWithMessageWithMode(
	m mode.Mode,
	c *gin.Context,
	statusCode int,
	message string,
	opts ...relaymodel.WrapperErrorOptionFunc,
) {
	c.JSON(statusCode,
		relaymodel.WrapperErrorWithMessage(m, statusCode, message, opts...),
	)
	c.Abort()
}

func AbortLogWithMessage(
	c *gin.Context,
	statusCode int,
	message string,
	opts ...relaymodel.WrapperErrorOptionFunc,
) {
	common.GetLogger(c).Error(message)
	AbortWithMessage(c, statusCode, message, opts...)
}

func AbortWithMessage(
	c *gin.Context,
	statusCode int,
	message string,
	opts ...relaymodel.WrapperErrorOptionFunc,
) {
	c.JSON(statusCode,
		relaymodel.WrapperErrorWithMessage(GetMode(c), statusCode, message, opts...),
	)
	c.Abort()
}

func GetMode(c *gin.Context) mode.Mode {
	m, exists := c.Get(Mode)
	if !exists {
		return mode.Unknown
	}

	v, ok := m.(mode.Mode)
	if !ok {
		panic(fmt.Sprintf("mode type error: %T, %v", v, v))
	}

	return v
}
