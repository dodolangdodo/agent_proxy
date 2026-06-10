package middleware_test

import (
	"testing"

	"github.com/labring/aiproxy/core/middleware"
	"github.com/labring/aiproxy/core/relay/mode"
	"github.com/stretchr/testify/require"
)

func TestCheckRelayModeGeminiVideo(t *testing.T) {
	require.True(t, middleware.CheckRelayMode(mode.GeminiVideo, mode.GeminiVideo))
	require.True(t, middleware.CheckRelayMode(mode.VideoGenerationsJobs, mode.GeminiVideo))
	require.True(t, middleware.CheckRelayMode(mode.Videos, mode.GeminiVideo))
	require.True(t, middleware.CheckRelayMode(mode.VideoGenerationsContent, mode.GeminiVideo))
	require.True(t, middleware.CheckRelayMode(mode.VideosContent, mode.GeminiVideo))
	require.True(t, middleware.CheckRelayMode(mode.GeminiVideoOperations, mode.GeminiVideo))
	require.True(t, middleware.CheckRelayMode(mode.GeminiFiles, mode.GeminiVideo))

	require.False(t, middleware.CheckRelayMode(mode.VideosDelete, mode.GeminiVideo))
	require.False(t, middleware.CheckRelayMode(mode.ChatCompletions, mode.GeminiVideo))
	require.False(t, middleware.CheckRelayMode(mode.Gemini, mode.GeminiVideo))
	require.False(t, middleware.CheckRelayMode(mode.Responses, mode.GeminiVideo))
	require.False(t, middleware.CheckRelayMode(mode.GeminiVideo, mode.VideoGenerationsJobs))
	require.False(t, middleware.CheckRelayMode(mode.GeminiVideo, mode.Videos))
}

func TestCheckRelayModeGeminiTTS(t *testing.T) {
	require.True(t, middleware.CheckRelayMode(mode.ChatCompletions, mode.GeminiTTS))
	require.True(t, middleware.CheckRelayMode(mode.Gemini, mode.GeminiTTS))
	require.True(t, middleware.CheckRelayMode(mode.AudioSpeech, mode.GeminiTTS))

	require.False(t, middleware.CheckRelayMode(mode.ImagesGenerations, mode.GeminiTTS))
	require.False(t, middleware.CheckRelayMode(mode.GeminiVideo, mode.GeminiTTS))
}

func TestCheckRelayModeGeminiImage(t *testing.T) {
	require.True(t, middleware.CheckRelayMode(mode.ChatCompletions, mode.GeminiImage))
	require.True(t, middleware.CheckRelayMode(mode.Gemini, mode.GeminiImage))
	require.True(t, middleware.CheckRelayMode(mode.ImagesGenerations, mode.GeminiImage))

	require.False(t, middleware.CheckRelayMode(mode.ImagesEdits, mode.GeminiImage))
	require.False(t, middleware.CheckRelayMode(mode.AudioSpeech, mode.GeminiImage))
	require.False(t, middleware.CheckRelayMode(mode.GeminiVideo, mode.GeminiImage))
}

func TestCheckRelayModeResponsesCompatibility(t *testing.T) {
	require.True(t, middleware.CheckRelayMode(mode.ChatCompletions, mode.Responses))
	require.True(t, middleware.CheckRelayMode(mode.Anthropic, mode.Responses))
	require.True(t, middleware.CheckRelayMode(mode.Gemini, mode.Responses))
	require.True(t, middleware.CheckRelayMode(mode.Responses, mode.Responses))
	require.True(t, middleware.CheckRelayMode(mode.ResponsesGet, mode.Responses))
	require.True(t, middleware.CheckRelayMode(mode.ResponsesDelete, mode.Responses))
	require.True(t, middleware.CheckRelayMode(mode.ResponsesCancel, mode.Responses))
	require.True(t, middleware.CheckRelayMode(mode.ResponsesInputItems, mode.Responses))

	require.False(t, middleware.CheckRelayMode(mode.Completions, mode.Responses))
	require.False(t, middleware.CheckRelayMode(mode.ChatCompletions, mode.ResponsesGet))
	require.False(t, middleware.CheckRelayMode(mode.ChatCompletions, mode.ResponsesDelete))
	require.False(t, middleware.CheckRelayMode(mode.ChatCompletions, mode.ResponsesCancel))
	require.False(t, middleware.CheckRelayMode(mode.ChatCompletions, mode.ResponsesInputItems))
	require.False(t, middleware.CheckRelayMode(mode.Anthropic, mode.ResponsesGet))
	require.False(t, middleware.CheckRelayMode(mode.Gemini, mode.ResponsesGet))
}
