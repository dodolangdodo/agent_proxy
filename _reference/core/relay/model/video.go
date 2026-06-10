package model

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/labring/aiproxy/core/relay/adaptor"
)

type VideoGenerationJobRequest struct {
	Prompt    string `json:"prompt"`
	Model     string `json:"model"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	NVariants int    `json:"n_variants"`
	NSeconds  int    `json:"n_seconds"`
	Seconds   int    `json:"seconds,omitempty"`
}

type VideosRequest struct {
	Prompt         string `json:"prompt,omitempty"`
	Model          string `json:"model,omitempty"`
	Seconds        any    `json:"seconds,omitempty"`
	Size           string `json:"size,omitempty"`
	InputReference string `json:"input_reference,omitempty"`
}

type VideosRemixRequest struct {
	VideosRequest
}

type VideosEditRequest struct {
	VideosRequest
	Video any `json:"video,omitempty"`
}

type VideosExtensionRequest struct {
	VideosRequest
	Video any `json:"video,omitempty"`
}

type VideoGenerationJobStatus = string

const (
	VideoGenerationJobStatusQueued     VideoGenerationJobStatus = "queued"
	VideoGenerationJobStatusProcessing VideoGenerationJobStatus = "processing"
	VideoGenerationJobStatusRunning    VideoGenerationJobStatus = "running"
	VideoGenerationJobStatusSucceeded  VideoGenerationJobStatus = "succeeded"
)

type VideoGenerationJob struct {
	Object       string                   `json:"object"`
	ID           string                   `json:"id"`
	Status       VideoGenerationJobStatus `json:"status"`
	CreatedAt    int64                    `json:"created_at"`
	FinishedAt   *int64                   `json:"finished_at"`
	ExpiresAt    *int64                   `json:"expires_at"`
	Generations  []VideoGenerations       `json:"generations"`
	Prompt       string                   `json:"prompt"`
	Model        string                   `json:"model"`
	NVariants    int                      `json:"n_variants"`
	NSeconds     int                      `json:"n_seconds"`
	Width        int                      `json:"width"`
	Height       int                      `json:"height"`
	FinishReason *string                  `json:"finish_reason"`
}

type VideoGenerations struct {
	Object    string `json:"object"`
	ID        string `json:"id"`
	JobID     string `json:"job_id"`
	CreatedAt int64  `json:"created_at"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Prompt    string `json:"prompt"`
	NSeconds  int    `json:"n_seconds"`
}

type VideoStatus = string

const (
	VideoStatusQueued     VideoStatus = "queued"
	VideoStatusInProgress VideoStatus = "in_progress"
	VideoStatusCompleted  VideoStatus = "completed"
	VideoStatusSucceeded  VideoStatus = "succeeded"
	VideoStatusFailed     VideoStatus = "failed"
	VideoStatusCancelled  VideoStatus = "cancelled"
)

type Video struct {
	ID        string         `json:"id"`
	Object    string         `json:"object"`
	CreatedAt int64          `json:"created_at,omitempty"`
	Status    VideoStatus    `json:"status,omitempty"`
	Progress  int            `json:"progress,omitempty"`
	Model     string         `json:"model,omitempty"`
	Prompt    string         `json:"prompt,omitempty"`
	Seconds   int            `json:"seconds,omitempty"`
	Size      string         `json:"size,omitempty"`
	Error     map[string]any `json:"error,omitempty"`
}

type OpenAIVideoError struct {
	Detail string `json:"detail"`
}

func NewOpenAIVideoError(statusCode int, err OpenAIVideoError) adaptor.Error {
	return adaptor.NewError(statusCode, err)
}

func WrapperOpenAIVideoError(err error, statusCode int) adaptor.Error {
	return WrapperOpenAIVideoErrorWithMessage(err.Error(), statusCode)
}

func WrapperOpenAIVideoErrorWithMessage(message string, statusCode int) adaptor.Error {
	return NewOpenAIVideoError(statusCode, OpenAIVideoError{
		Detail: message,
	})
}

func VideoResolutionFromDimensions(width, height int) string {
	if width <= 0 || height <= 0 {
		return ""
	}

	shortSide := min(width, height)
	if resolution := VideoResolutionFromHeight(shortSide); resolution != "" {
		return resolution
	}

	return fmt.Sprintf("%dx%d", width, height)
}

func VideoAspectRatioFromSize(size string) string {
	size = strings.ToLower(strings.TrimSpace(size))
	switch size {
	case "16:9", "9:16", "1:1", "4:3", "3:4":
		return size
	}

	width, height, ok := ParseVideoDimensions(size)
	if !ok || width <= 0 || height <= 0 {
		return ""
	}

	return ClosestVideoAspectRatio(width, height)
}

func VideoResolutionFromHeight(height int) string {
	switch {
	case height >= 2000:
		return "4k"
	case height >= 1000:
		return "1080p"
	case height >= 700:
		return "720p"
	case height >= 400:
		return "480p"
	default:
		return ""
	}
}

func ParseVideoDimensions(size string) (int, int, bool) {
	size = strings.ReplaceAll(strings.ToLower(strings.TrimSpace(size)), "×", "x")
	parts := strings.Split(size, "x")

	if len(parts) != 2 {
		return 0, 0, false
	}

	width, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, false
	}

	height, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, 0, false
	}

	return width, height, true
}

func ClosestVideoAspectRatio(width, height int) string {
	type candidate struct {
		label string
		ratio float64
	}

	ratio := float64(width) / float64(height)
	candidates := []candidate{
		{"16:9", 16.0 / 9.0},
		{"9:16", 9.0 / 16.0},
		{"1:1", 1},
		{"4:3", 4.0 / 3.0},
		{"3:4", 3.0 / 4.0},
	}

	best := candidates[0]
	bestDelta := absFloat64(ratio - best.ratio)

	for _, item := range candidates[1:] {
		delta := absFloat64(ratio - item.ratio)
		if delta < bestDelta {
			best = item
			bestDelta = delta
		}
	}

	return best.label
}

func absFloat64(value float64) float64 {
	if value < 0 {
		return -value
	}

	return value
}
