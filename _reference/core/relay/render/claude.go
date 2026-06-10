package render

import (
	"errors"
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
)

func ClaudeData(c *gin.Context, data []byte) {
	if len(c.Errors) > 0 {
		return
	}

	if c.IsAborted() {
		return
	}

	c.Render(-1, &Anthropic{Data: data})
	c.Writer.Flush()
}

func ClaudeEventData(c *gin.Context, event string, data []byte) {
	if len(c.Errors) > 0 {
		return
	}

	if c.IsAborted() {
		return
	}

	c.Render(-1, &Anthropic{Event: event, Data: data})
	c.Writer.Flush()
}

func ClaudeObjectData(c *gin.Context, object any) error {
	if len(c.Errors) > 0 {
		return c.Errors.Last()
	}

	if c.IsAborted() {
		return errors.New("context aborted")
	}

	jsonData, err := sonic.Marshal(object)
	if err != nil {
		return fmt.Errorf("error marshalling object: %w", err)
	}

	c.Render(-1, &Anthropic{Data: jsonData})
	c.Writer.Flush()

	return nil
}

func ClaudeEventObjectData(c *gin.Context, event string, object any) error {
	if len(c.Errors) > 0 {
		return c.Errors.Last()
	}

	if c.IsAborted() {
		return errors.New("context aborted")
	}

	jsonData, err := sonic.Marshal(object)
	if err != nil {
		return fmt.Errorf("error marshalling object: %w", err)
	}

	c.Render(-1, &Anthropic{Event: event, Data: jsonData})
	c.Writer.Flush()

	return nil
}
