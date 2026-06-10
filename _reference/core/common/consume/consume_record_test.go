package consume_test

import (
	"net/http"
	"testing"

	"github.com/labring/aiproxy/core/common/consume"
	"github.com/labring/aiproxy/core/relay/meta"
	"github.com/labring/aiproxy/core/relay/mode"
)

func TestNeedRecordConsumeSkipsSuccessfulStoredVideoReads(t *testing.T) {
	tests := []mode.Mode{
		mode.VideoGenerationsGetJobs,
		mode.VideoGenerationsContent,
		mode.VideosGet,
		mode.VideosContent,
		mode.VideosDelete,
		mode.GeminiFiles,
		mode.GeminiVideoOperations,
		mode.AliVideoTasks,
		mode.DoubaoVideoTasks,
		mode.DoubaoVideoTasksDelete,
	}

	for _, relayMode := range tests {
		t.Run(relayMode.String(), func(t *testing.T) {
			if consume.NeedRecordConsumeForTest(
				http.StatusOK,
				&meta.Meta{Mode: relayMode},
			) {
				t.Fatalf("expected successful %s request not to record consume", relayMode)
			}

			if !consume.NeedRecordConsumeForTest(
				http.StatusInternalServerError,
				&meta.Meta{Mode: relayMode},
			) {
				t.Fatalf("expected failed %s request to record consume", relayMode)
			}
		})
	}
}

func TestNeedRecordConsumeSkipsSuccessfulDoubaoNativeDeleteNoContent(t *testing.T) {
	if consume.NeedRecordConsumeForTest(
		http.StatusNoContent,
		&meta.Meta{Mode: mode.DoubaoVideoTasksDelete},
	) {
		t.Fatal("expected successful doubao native delete 204 request not to record consume")
	}
}

func TestNeedRecordConsumeRecordsVideoCreateAndRemix(t *testing.T) {
	tests := []mode.Mode{
		mode.Videos,
		mode.VideosRemix,
	}

	for _, relayMode := range tests {
		t.Run(relayMode.String(), func(t *testing.T) {
			if !consume.NeedRecordConsumeForTest(http.StatusOK, &meta.Meta{Mode: relayMode}) {
				t.Fatalf("expected successful %s request to record consume", relayMode)
			}
		})
	}
}
