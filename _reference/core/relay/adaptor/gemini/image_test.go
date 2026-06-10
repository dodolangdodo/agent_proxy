package gemini_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/model"
	adaptorapi "github.com/labring/aiproxy/core/relay/adaptor"
	"github.com/labring/aiproxy/core/relay/adaptor/gemini"
	"github.com/labring/aiproxy/core/relay/adaptor/openai"
	"github.com/labring/aiproxy/core/relay/meta"
	"github.com/labring/aiproxy/core/relay/mode"
	relaymodel "github.com/labring/aiproxy/core/relay/model"
	"github.com/stretchr/testify/assert"
)

func TestConvertImageRequestMapsOpenAIImageToGemini(t *testing.T) {
	t.Parallel()

	channel := &model.Channel{Type: model.ChannelTypeGoogleGemini}
	meta := meta.NewMeta(
		channel,
		mode.ImagesGenerations,
		"gemini-3-pro-image-preview",
		model.ModelConfig{Type: mode.GeminiImage},
	)

	req, err := http.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"http://localhost/v1/images/generations",
		bytes.NewBufferString(
			`{"model":"gemini-3-pro-image-preview","prompt":"Draw a cat.","size":"1536x1024","response_format":"url"}`,
		),
	)
	assert.NoError(t, err)

	result, err := gemini.ConvertImageRequest(meta, req)
	assert.NoError(t, err)
	assert.Equal(t, "url", meta.GetString(openai.MetaResponseFormat))

	bodyBytes, err := io.ReadAll(result.Body)
	assert.NoError(t, err)

	var geminiReq relaymodel.GeminiChatRequest

	err = json.Unmarshal(bodyBytes, &geminiReq)
	assert.NoError(t, err)
	assert.Len(t, geminiReq.Contents, 1)
	assert.Equal(t, "Draw a cat.", geminiReq.Contents[0].Parts[0].Text)
	assert.Equal(
		t,
		[]string{relaymodel.GeminiModalityImage},
		geminiReq.GenerationConfig.ResponseModalities,
	)
	assert.Equal(t, "3:2", geminiReq.GenerationConfig.ImageConfig.AspectRatio)
	assert.Equal(t, "2K", geminiReq.GenerationConfig.ImageConfig.ImageSize)
}

func TestConvertImageRequestMapsSmallDimensionsToImageSize(t *testing.T) {
	t.Parallel()

	meta := meta.NewMeta(
		&model.Channel{Type: model.ChannelTypeGoogleGemini},
		mode.ImagesGenerations,
		"gemini-3-pro-image-preview",
		model.ModelConfig{Type: mode.GeminiImage},
	)

	req, err := http.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"http://localhost/v1/images/generations",
		bytes.NewBufferString(
			`{"model":"gemini-3-pro-image-preview","prompt":"Draw a cat.","size":"512x512"}`,
		),
	)
	assert.NoError(t, err)

	result, err := gemini.ConvertImageRequest(meta, req)
	assert.NoError(t, err)

	bodyBytes, err := io.ReadAll(result.Body)
	assert.NoError(t, err)

	var geminiReq relaymodel.GeminiChatRequest
	assert.NoError(t, json.Unmarshal(bodyBytes, &geminiReq))
	assert.Equal(t, "1:1", geminiReq.GenerationConfig.ImageConfig.AspectRatio)
	assert.Equal(t, "512", geminiReq.GenerationConfig.ImageConfig.ImageSize)
}

func TestConvertImageRequestNormalizesDimensionDelimiters(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		size            string
		wantAspectRatio string
		wantImageSize   string
	}{
		{
			name:            "asterisk",
			size:            "1024*1024",
			wantAspectRatio: "1:1",
			wantImageSize:   "1K",
		},
		{
			name:            "multiplication sign",
			size:            "1536×1024",
			wantAspectRatio: "3:2",
			wantImageSize:   "2K",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			meta := meta.NewMeta(
				&model.Channel{Type: model.ChannelTypeGoogleGemini},
				mode.ImagesGenerations,
				"gemini-3-pro-image-preview",
				model.ModelConfig{Type: mode.GeminiImage},
			)

			req, err := http.NewRequestWithContext(
				t.Context(),
				http.MethodPost,
				"http://localhost/v1/images/generations",
				bytes.NewBufferString(
					`{"model":"gemini-3-pro-image-preview","prompt":"Draw a cat.","size":"`+tt.size+`"}`,
				),
			)
			assert.NoError(t, err)

			result, err := gemini.ConvertImageRequest(meta, req)
			assert.NoError(t, err)

			bodyBytes, err := io.ReadAll(result.Body)
			assert.NoError(t, err)

			var geminiReq relaymodel.GeminiChatRequest
			assert.NoError(t, json.Unmarshal(bodyBytes, &geminiReq))
			assert.Equal(t, tt.wantAspectRatio, geminiReq.GenerationConfig.ImageConfig.AspectRatio)
			assert.Equal(t, tt.wantImageSize, geminiReq.GenerationConfig.ImageConfig.ImageSize)
		})
	}
}

func TestConvertImageRequestMapsSquareDimensionsToGeminiImageSize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		size          string
		wantImageSize string
	}{
		{
			size:          "1024x1024",
			wantImageSize: "1K",
		},
		{
			size:          "2048x2048",
			wantImageSize: "2K",
		},
		{
			size:          "4096x4096",
			wantImageSize: "4K",
		},
	}

	for _, tt := range tests {
		t.Run(tt.size, func(t *testing.T) {
			t.Parallel()

			meta := meta.NewMeta(
				&model.Channel{Type: model.ChannelTypeGoogleGemini},
				mode.ImagesGenerations,
				"gemini-3-pro-image-preview",
				model.ModelConfig{Type: mode.GeminiImage},
			)

			req, err := http.NewRequestWithContext(
				t.Context(),
				http.MethodPost,
				"http://localhost/v1/images/generations",
				bytes.NewBufferString(
					`{"model":"gemini-3-pro-image-preview","prompt":"Draw a cat.","size":"`+tt.size+`"}`,
				),
			)
			assert.NoError(t, err)

			result, err := gemini.ConvertImageRequest(meta, req)
			assert.NoError(t, err)

			bodyBytes, err := io.ReadAll(result.Body)
			assert.NoError(t, err)

			var geminiReq relaymodel.GeminiChatRequest
			assert.NoError(t, json.Unmarshal(bodyBytes, &geminiReq))
			assert.Equal(t, "1:1", geminiReq.GenerationConfig.ImageConfig.AspectRatio)
			assert.Equal(t, tt.wantImageSize, geminiReq.GenerationConfig.ImageConfig.ImageSize)
		})
	}
}

func TestConvertImageEditRequestMapsMultipartToGemini(t *testing.T) {
	t.Parallel()

	meta := meta.NewMeta(
		&model.Channel{Type: model.ChannelTypeGoogleGemini},
		mode.ImagesEdits,
		"gemini-3-pro-image-preview",
		model.ModelConfig{Type: mode.GeminiImage},
	)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	assert.NoError(t, writer.WriteField("model", "gemini-3-pro-image-preview"))
	assert.NoError(t, writer.WriteField("prompt", "Add a hat."))
	assert.NoError(t, writer.WriteField("size", "1024x1536"))
	assert.NoError(t, writer.WriteField("response_format", "b64_json"))

	part, err := writer.CreateFormFile("image", "input.png")
	assert.NoError(t, err)
	_, err = part.Write([]byte{
		0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a,
		0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
		0x08, 0x02, 0x00, 0x00, 0x00, 0x90, 0x77, 0x53, 0xde,
		0x00, 0x00, 0x00, 0x0c, 0x49, 0x44, 0x41, 0x54,
		0x08, 0xd7, 0x63, 0xf8, 0xcf, 0xc0, 0x00, 0x00,
		0x03, 0x01, 0x01, 0x00, 0x18, 0xdd, 0x8d, 0xb0,
		0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44,
		0xae, 0x42, 0x60, 0x82,
	})
	assert.NoError(t, err)
	assert.NoError(t, writer.Close())

	req, err := http.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"http://localhost/v1/images/edits",
		body,
	)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	result, err := gemini.ConvertImageEditRequest(meta, req)
	assert.NoError(t, err)
	assert.Equal(t, "b64_json", meta.GetString(openai.MetaResponseFormat))

	bodyBytes, err := io.ReadAll(result.Body)
	assert.NoError(t, err)

	var geminiReq relaymodel.GeminiChatRequest
	assert.NoError(t, json.Unmarshal(bodyBytes, &geminiReq))
	assert.Len(t, geminiReq.Contents, 1)
	assert.Len(t, geminiReq.Contents[0].Parts, 2)
	assert.NotNil(t, geminiReq.Contents[0].Parts[0].InlineData)
	assert.Equal(t, "image/png", geminiReq.Contents[0].Parts[0].InlineData.MimeType)
	assert.Equal(t, "Add a hat.", geminiReq.Contents[0].Parts[1].Text)
	assert.Equal(t, "2:3", geminiReq.GenerationConfig.ImageConfig.AspectRatio)
	assert.Equal(t, "2K", geminiReq.GenerationConfig.ImageConfig.ImageSize)
}

func TestConvertImageEditRequestAcceptsImageArrayFiles(t *testing.T) {
	t.Parallel()

	meta := meta.NewMeta(
		&model.Channel{Type: model.ChannelTypeGoogleGemini},
		mode.ImagesEdits,
		"gemini-3-pro-image-preview",
		model.ModelConfig{Type: mode.GeminiImage},
	)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	assert.NoError(t, writer.WriteField("model", "gemini-3-pro-image-preview"))
	assert.NoError(t, writer.WriteField("prompt", "Blend them."))

	for _, name := range []string{"input-1.png", "input-2.png"} {
		part, err := writer.CreateFormFile("image[]", name)
		assert.NoError(t, err)
		_, err = part.Write([]byte{
			0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a,
			0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52,
			0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
			0x08, 0x02, 0x00, 0x00, 0x00, 0x90, 0x77, 0x53, 0xde,
		})
		assert.NoError(t, err)
	}

	assert.NoError(t, writer.Close())

	req, err := http.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"http://localhost/v1/images/edits",
		body,
	)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	result, err := gemini.ConvertImageEditRequest(meta, req)
	assert.NoError(t, err)

	bodyBytes, err := io.ReadAll(result.Body)
	assert.NoError(t, err)

	var geminiReq relaymodel.GeminiChatRequest
	assert.NoError(t, json.Unmarshal(bodyBytes, &geminiReq))
	assert.Len(t, geminiReq.Contents, 1)
	assert.Len(t, geminiReq.Contents[0].Parts, 3)
	assert.NotNil(t, geminiReq.Contents[0].Parts[0].InlineData)
	assert.NotNil(t, geminiReq.Contents[0].Parts[1].InlineData)
	assert.Equal(t, "Blend them.", geminiReq.Contents[0].Parts[2].Text)
}

func TestConvertImageEditRequestUsesRequestContextForRemoteImages(t *testing.T) {
	t.Parallel()

	remoteImageStarted := make(chan struct{})
	remoteImageRelease := make(chan struct{})

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		close(remoteImageStarted)
		<-remoteImageRelease
		w.Header().Set("Content-Type", "image/png")
		_, _ = w.Write([]byte("not reached"))
	}))
	defer server.Close()
	defer close(remoteImageRelease)

	meta := meta.NewMeta(
		&model.Channel{Type: model.ChannelTypeGoogleGemini},
		mode.ImagesEdits,
		"gemini-3-pro-image-preview",
		model.ModelConfig{Type: mode.GeminiImage},
	)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	assert.NoError(t, writer.WriteField("model", "gemini-3-pro-image-preview"))
	assert.NoError(t, writer.WriteField("prompt", "Use the remote image."))
	assert.NoError(t, writer.WriteField("image_url", server.URL+"/image.png"))
	assert.NoError(t, writer.Close())

	ctx, cancel := context.WithCancel(t.Context())
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"http://localhost/v1/images/edits",
		body,
	)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	errCh := make(chan error, 1)
	go func() {
		_, err := gemini.ConvertImageEditRequest(meta, req)
		errCh <- err
	}()

	<-remoteImageStarted
	cancel()

	err = <-errCh
	assert.ErrorIs(t, err, context.Canceled)
}

func TestConvertImageEditRequestKeepsRemoteImagesWhenAutoDownloadDisabled(t *testing.T) {
	t.Parallel()

	meta := meta.NewMeta(
		&model.Channel{
			Type: model.ChannelTypeGoogleGemini,
			Configs: model.ChannelConfigs{
				"disable_auto_image_url_to_base64": true,
			},
		},
		mode.ImagesEdits,
		"gemini-3-pro-image-preview",
		model.ModelConfig{Type: mode.GeminiImage},
	)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	assert.NoError(t, writer.WriteField("model", "gemini-3-pro-image-preview"))
	assert.NoError(t, writer.WriteField("prompt", "Use the remote image."))
	assert.NoError(t, writer.WriteField("image_url", "https://example.com/image.png"))
	assert.NoError(t, writer.Close())

	req, err := http.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"http://localhost/v1/images/edits",
		body,
	)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	result, err := gemini.ConvertImageEditRequest(meta, req)
	assert.NoError(t, err)

	bodyBytes, err := io.ReadAll(result.Body)
	assert.NoError(t, err)

	var geminiReq relaymodel.GeminiChatRequest
	assert.NoError(t, json.Unmarshal(bodyBytes, &geminiReq))
	assert.Len(t, geminiReq.Contents, 1)
	assert.Len(t, geminiReq.Contents[0].Parts, 2)
	assert.Nil(t, geminiReq.Contents[0].Parts[0].InlineData)
	assert.NotNil(t, geminiReq.Contents[0].Parts[0].FileData)
	assert.Equal(
		t,
		"https://example.com/image.png",
		geminiReq.Contents[0].Parts[0].FileData.FileURI,
	)
	assert.Equal(t, "Use the remote image.", geminiReq.Contents[0].Parts[1].Text)
}

func TestConvertImageRequestMapsImageSizePreset(t *testing.T) {
	t.Parallel()

	meta := meta.NewMeta(
		&model.Channel{Type: model.ChannelTypeGoogleGemini},
		mode.ImagesGenerations,
		"gemini-3-pro-image-preview",
		model.ModelConfig{Type: mode.GeminiImage},
	)

	req, err := http.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"http://localhost/v1/images/generations",
		bytes.NewBufferString(
			`{"model":"gemini-3-pro-image-preview","prompt":"Draw a cat.","size":"2K"}`,
		),
	)
	assert.NoError(t, err)

	result, err := gemini.ConvertImageRequest(meta, req)
	assert.NoError(t, err)

	bodyBytes, err := io.ReadAll(result.Body)
	assert.NoError(t, err)

	var geminiReq relaymodel.GeminiChatRequest

	err = json.Unmarshal(bodyBytes, &geminiReq)
	assert.NoError(t, err)
	assert.Equal(t, "2K", geminiReq.GenerationConfig.ImageConfig.ImageSize)
	assert.Empty(t, geminiReq.GenerationConfig.ImageConfig.AspectRatio)
}

func TestConvertImageRequestPreservesStream(t *testing.T) {
	t.Parallel()

	meta := meta.NewMeta(
		&model.Channel{Type: model.ChannelTypeGoogleGemini},
		mode.ImagesGenerations,
		"gemini-3-pro-image-preview",
		model.ModelConfig{Type: mode.GeminiImage},
	)

	req, err := http.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"http://localhost/v1/images/generations",
		bytes.NewBufferString(
			`{"model":"gemini-3-pro-image-preview","prompt":"Draw a cat.","stream":true}`,
		),
	)
	assert.NoError(t, err)

	_, err = gemini.ConvertImageRequest(meta, req)
	assert.NoError(t, err)
	assert.True(t, meta.GetBool("stream"))

	adaptor := &gemini.Adaptor{}
	requestURL, err := adaptor.GetRequestURL(meta, nil, nil)
	assert.NoError(t, err)
	assert.Contains(t, requestURL.URL, ":streamGenerateContent?alt=sse")
}

func TestConvertImageRequestMissingPromptReturnsRelayError(t *testing.T) {
	t.Parallel()

	meta := meta.NewMeta(
		&model.Channel{Type: model.ChannelTypeGoogleGemini},
		mode.ImagesGenerations,
		"gemini-3-pro-image-preview",
		model.ModelConfig{Type: mode.GeminiImage},
	)

	req := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/images/generations",
		bytes.NewBufferString(`{"model":"gemini-3-pro-image-preview"}`),
	)

	_, err := gemini.ConvertImageRequest(meta, req)
	assert.Error(t, err)

	var relayErr adaptorapi.Error
	assert.ErrorAs(t, err, &relayErr)
	assert.Equal(t, http.StatusBadRequest, relayErr.StatusCode())
}

func TestConvertImageRequestAllowsMultipleRequestedImages(t *testing.T) {
	t.Parallel()

	meta := meta.NewMeta(
		&model.Channel{Type: model.ChannelTypeGoogleGemini},
		mode.ImagesGenerations,
		"gemini-3-pro-image-preview",
		model.ModelConfig{Type: mode.GeminiImage},
	)

	req, err := http.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"http://localhost/v1/images/generations",
		bytes.NewBufferString(
			`{"model":"gemini-3-pro-image-preview","prompt":"Draw a cat.","n":2}`,
		),
	)
	assert.NoError(t, err)

	result, err := gemini.ConvertImageRequest(meta, req)
	assert.NoError(t, err)

	bodyBytes, err := io.ReadAll(result.Body)
	assert.NoError(t, err)

	var geminiReq relaymodel.GeminiChatRequest

	err = json.Unmarshal(bodyBytes, &geminiReq)
	assert.NoError(t, err)
	assert.Zero(t, geminiReq.GenerationConfig.CandidateCount)
}

func TestGeminiImageAspectRatioFromSize(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "1:1", gemini.GeminiImageAspectRatioFromSizeForTest("1024x1024"))
	assert.Equal(t, "3:2", gemini.GeminiImageAspectRatioFromSizeForTest("1536x1024"))
	assert.Equal(t, "2:3", gemini.GeminiImageAspectRatioFromSizeForTest("1024x1536"))
}

func TestImageHandlerConvertsGeminiImageResponseToOpenAI(t *testing.T) {
	t.Parallel()

	meta := meta.NewMeta(
		&model.Channel{Type: model.ChannelTypeGoogleGemini},
		mode.ImagesGenerations,
		"gemini-3-pro-image-preview",
		model.ModelConfig{Type: mode.GeminiImage},
		meta.WithRequestUsage(model.Usage{
			InputTokens:  model.ZeroNullInt64(1),
			OutputTokens: model.ZeroNullInt64(1),
			TotalTokens:  model.ZeroNullInt64(2),
		}),
	)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": {"application/json"}},
		Body: io.NopCloser(bytes.NewBufferString(`{
			"candidates":[{
				"content":{"parts":[
					{"text":"done"},
					{"inlineData":{"mimeType":"image/png","data":"aW1hZ2U="}}
				]}
			}],
			"usageMetadata":{
				"promptTokenCount":10,
				"candidatesTokenCount":20,
				"totalTokenCount":30,
				"promptTokensDetails":[{"modality":"IMAGE","tokenCount":3}],
				"candidatesTokensDetails":[{"modality":"IMAGE","tokenCount":20}]
			}
		}`)),
	}

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/images/generations",
		nil,
	)

	result, err := gemini.ImageHandler(meta, c, resp)
	assert.Nil(t, err)
	assert.Equal(t, int64(10), int64(result.Usage.InputTokens))
	assert.Equal(t, int64(20), int64(result.Usage.OutputTokens))
	assert.Equal(t, int64(20), int64(result.Usage.ImageOutputTokens))
	assert.Equal(t, int64(30), int64(result.Usage.TotalTokens))

	var imageResp relaymodel.ImageResponse

	assert.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &imageResp))
	assert.Len(t, imageResp.Data, 1)
	assert.Equal(t, "aW1hZ2U=", imageResp.Data[0].B64Json)
	assert.NotNil(t, imageResp.Usage)
	assert.Equal(t, int64(20), imageResp.Usage.OutputTokens)
	assert.Equal(t, int64(20), imageResp.Usage.OutputTokensDetails.ImageTokens)

	var raw map[string]any
	assert.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &raw))
	assert.NotContains(t, raw, "candidates")
	assert.NotContains(t, raw, "usageMetadata")
}

func TestImageHandlerHonorsURLResponseFormat(t *testing.T) {
	t.Parallel()

	meta := meta.NewMeta(
		&model.Channel{Type: model.ChannelTypeGoogleGemini},
		mode.ImagesGenerations,
		"gemini-3-pro-image-preview",
		model.ModelConfig{Type: mode.GeminiImage},
	)
	meta.Set(openai.MetaResponseFormat, "url")

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": {"application/json"}},
		Body: io.NopCloser(bytes.NewBufferString(`{
			"candidates":[{
				"content":{"parts":[
					{"inlineData":{"mimeType":"image/png","data":"aW1hZ2U="}}
				]}
			}]
		}`)),
	}

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/images/generations",
		nil,
	)

	_, err := gemini.ImageHandler(meta, c, resp)
	assert.Nil(t, err)

	var imageResp relaymodel.ImageResponse
	assert.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &imageResp))
	assert.Len(t, imageResp.Data, 1)
	assert.Empty(t, imageResp.Data[0].B64Json)
	assert.Equal(t, "data:image/png;base64,aW1hZ2U=", imageResp.Data[0].URL)
}

func TestImageHandlerEmptyImageErrorIncludesGeminiText(t *testing.T) {
	t.Parallel()

	meta := meta.NewMeta(
		&model.Channel{Type: model.ChannelTypeGoogleGemini},
		mode.ImagesGenerations,
		"gemini-3.1-flash-image-preview",
		model.ModelConfig{Type: mode.GeminiImage},
	)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": {"application/json"}},
		Body: io.NopCloser(bytes.NewBufferString(`{
			"candidates":[{
				"finishReason":"STOP",
				"content":{"parts":[{"text":"I need a more specific image prompt."}]}
			}]
		}`)),
	}

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/images/generations",
		nil,
	)

	_, err := gemini.ImageHandler(meta, c, resp)
	assert.NotNil(t, err)

	body, marshalErr := err.MarshalJSON()
	assert.NoError(t, marshalErr)
	assert.Contains(t, string(body), "gemini image response image is empty")
	assert.Contains(t, string(body), "I need a more specific image prompt.")
	assert.Contains(t, string(body), "finish_reason=STOP")
}

func TestImageHandlerChargesActualGeminiImageCount(t *testing.T) {
	t.Parallel()

	meta := meta.NewMeta(
		&model.Channel{Type: model.ChannelTypeGoogleGemini},
		mode.ImagesGenerations,
		"gemini-3-pro-image-preview",
		model.ModelConfig{Type: mode.GeminiImage},
		meta.WithRequestUsageContext(model.UsageContext{
			Resolution: "1024x1024",
		}),
		meta.WithRequestUsage(model.Usage{
			InputTokens:  model.ZeroNullInt64(5),
			OutputTokens: model.ZeroNullInt64(2),
			TotalTokens:  model.ZeroNullInt64(7),
		}),
	)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": {"application/json"}},
		Body: io.NopCloser(bytes.NewBufferString(`{
			"candidates":[{
				"content":{"parts":[
					{"inlineData":{"mimeType":"image/png","data":"aW1hZ2U="}}
				]}
			}]
		}`)),
	}

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/images/generations",
		nil,
	)

	result, err := gemini.ImageHandler(meta, c, resp)
	assert.Nil(t, err)
	assert.Equal(t, int64(1120), int64(result.Usage.OutputTokens))
	assert.Equal(t, int64(1120), int64(result.Usage.ImageOutputTokens))
	assert.Equal(t, int64(1120), int64(result.Usage.TotalTokens))

	var imageResp relaymodel.ImageResponse

	assert.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &imageResp))
	assert.NotNil(t, imageResp.Usage)
	assert.Equal(t, int64(1120), imageResp.Usage.OutputTokens)
	assert.Equal(t, int64(1120), imageResp.Usage.OutputTokensDetails.ImageTokens)
}

func TestImageStreamHandlerConvertsGeminiStreamToOpenAIImageEvents(t *testing.T) {
	t.Parallel()

	meta := meta.NewMeta(
		&model.Channel{Type: model.ChannelTypeGoogleGemini},
		mode.ImagesGenerations,
		"gemini-3-pro-image-preview",
		model.ModelConfig{Type: mode.GeminiImage},
	)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": {"text/event-stream"}},
		Body: io.NopCloser(bytes.NewBufferString(
			"data: {\"candidates\":[{\"content\":{\"parts\":[{\"inlineData\":{\"mimeType\":\"image/png\",\"data\":\"cGFydGlhbA==\"}}]}}]}\n\n" +
				"data: {\"usageMetadata\":{\"promptTokenCount\":4,\"candidatesTokenCount\":8,\"totalTokenCount\":12,\"promptTokensDetails\":[],\"candidatesTokensDetails\":[{\"modality\":\"IMAGE\",\"tokenCount\":8}]}}\n\n",
		)),
	}

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/v1/images/generations",
		nil,
	)

	result, err := gemini.ImageHandler(meta, c, resp)
	assert.Nil(t, err)
	assert.Equal(t, int64(8), int64(result.Usage.ImageOutputTokens))

	body := recorder.Body.String()
	assert.Contains(t, body, "event: "+relaymodel.ImageStreamEventPartialImage+"\n")
	assert.Contains(t, body, `"type":"`+relaymodel.ImageStreamEventPartialImage+`"`)
	assert.Contains(t, body, `"partial_image_index":0`)
	assert.Contains(t, body, `"b64_json":"cGFydGlhbA=="`)
	assert.Contains(t, body, "event: "+relaymodel.ImageStreamEventCompleted+"\n")
	assert.Contains(t, body, `"type":"`+relaymodel.ImageStreamEventCompleted+`"`)
	assert.Contains(t, body, `"b64_json":"cGFydGlhbA=="`)
	assert.NotContains(t, body, `[DONE]`)
}
