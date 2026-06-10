package baidu

import (
	"github.com/labring/aiproxy/core/relay/model"
)

type Error struct {
	ErrorMsg  string `json:"error_msg"`
	ErrorCode int    `json:"error_code"`
}

type ErrorResponse struct {
	*Error `json:"error"`
	ID     string `json:"id"`
}

type ChatResponse struct {
	Usage            *model.ChatUsage `json:"usage"`
	*Error           `json:"error"`
	ID               string `json:"id"`
	Object           string `json:"object"`
	Result           string `json:"result"`
	Created          int64  `json:"created"`
	IsTruncated      bool   `json:"is_truncated"`
	NeedClearHistory bool   `json:"need_clear_history"`
}

type ChatStreamResponse struct {
	ChatResponse
	SentenceID int  `json:"sentence_id"`
	IsEnd      bool `json:"is_end"`
}
