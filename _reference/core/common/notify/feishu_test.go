package notify_test

import (
	"os"
	"testing"

	"github.com/labring/aiproxy/core/common/notify"
)

func TestPostToFeiShuv2(t *testing.T) {
	fshook := os.Getenv("FEISHU_WEBHOOK")
	if fshook == "" {
		return
	}

	err := notify.PostToFeiShuv2(
		t.Context(),
		notify.FeishuColorRed,
		"Error",
		"Error Message",
		os.Getenv("FEISHU_WEBHOOK"))
	if err != nil {
		t.Error(err)
	}
}
