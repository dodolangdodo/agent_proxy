//nolint:testpackage
package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/meta"
	"github.com/labring/aiproxy/core/relay/mode"
	relaymodel "github.com/labring/aiproxy/core/relay/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConvertImagesRequest_DefaultIncludesModel(t *testing.T) {
	meta := meta.NewMeta(nil, mode.ImagesGenerations, "gpt-image-1", model.ModelConfig{})

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"http://example.com/v1/images/generations",
		strings.NewReader(`{"prompt":"test","response_format":"b64_json"}`),
	)
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	result, err := ConvertImagesRequest(meta, req)
	require.NoError(t, err)

	body, err := io.ReadAll(result.Body)
	require.NoError(t, err)

	var payload map[string]any

	err = json.Unmarshal(body, &payload)
	require.NoError(t, err)

	assert.Equal(t, "gpt-image-1", payload["model"])
	assert.Equal(t, "b64_json", meta.GetString(MetaResponseFormat))
}

func TestConvertImagesRequest_CanRemoveModelDynamically(t *testing.T) {
	meta := meta.NewMeta(nil, mode.ImagesGenerations, "gpt-image-1", model.ModelConfig{})

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"http://example.com/v1/images/generations",
		strings.NewReader(`{"model":"ignored","prompt":"test","response_format":"b64_json"}`),
	)
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	result, err := ConvertImagesRequest(meta, req, ImagesRequestRemoveModel)
	require.NoError(t, err)

	body, err := io.ReadAll(result.Body)
	require.NoError(t, err)

	var payload map[string]any

	err = json.Unmarshal(body, &payload)
	require.NoError(t, err)

	_, ok := payload["model"]
	assert.False(t, ok)
	assert.Equal(t, "b64_json", meta.GetString(MetaResponseFormat))
}

func TestImagesHandlerClearsURLWhenB64JSONRequested(t *testing.T) {
	gin.SetMode(gin.TestMode)

	imageServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "image/png")
			_, _ = w.Write([]byte("png-bytes"))
		}),
	)
	defer imageServer.Close()

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"/v1/images/generations",
		nil,
	)

	m := meta.NewMeta(nil, mode.ImagesGenerations, "gpt-image-1", model.ModelConfig{})
	m.Set(MetaResponseFormat, "b64_json")

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Header:     make(http.Header),
		Body: io.NopCloser(strings.NewReader(`{
			"created": 1710000000,
			"data": [{"url": "` + imageServer.URL + `/one.png"}]
		}`)),
	}

	_, adaptorErr := ImagesHandler(m, ctx, resp)
	require.Nil(t, adaptorErr)

	var imageResponse relaymodel.ImageResponse
	require.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &imageResponse))
	require.Len(t, imageResponse.Data, 1)
	assert.Empty(t, imageResponse.Data[0].URL)
	assert.NotEmpty(t, imageResponse.Data[0].B64Json)
}

func TestConvertImagesEditsRequest_DefaultIncludesModel(t *testing.T) {
	meta := meta.NewMeta(nil, mode.ImagesEdits, "gpt-image-1", model.ModelConfig{})

	var body bytes.Buffer

	writer := multipart.NewWriter(&body)
	require.NoError(t, writer.WriteField("model", "ignored"))
	require.NoError(t, writer.WriteField("prompt", "edit prompt"))
	part, err := writer.CreateFormFile("image", "test.png")
	require.NoError(t, err)
	_, err = part.Write([]byte("png-bytes"))
	require.NoError(t, err)
	require.NoError(t, writer.Close())

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"http://example.com/v1/images/edits",
		bytes.NewReader(body.Bytes()),
	)
	require.NoError(t, err)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.ContentLength = int64(body.Len())

	result, err := ConvertImagesEditsRequest(meta, req, true)
	require.NoError(t, err)

	convertedBody, err := io.ReadAll(result.Body)
	require.NoError(t, err)

	convertedReq, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"http://example.com",
		bytes.NewReader(convertedBody),
	)
	require.NoError(t, err)
	convertedReq.Header.Set("Content-Type", result.Header.Get("Content-Type"))
	convertedReq.ContentLength = int64(len(convertedBody))

	err = convertedReq.ParseMultipartForm(1024 * 1024 * 4)
	require.NoError(t, err)

	assert.Equal(t, "gpt-image-1", convertedReq.MultipartForm.Value["model"][0])
	assert.Equal(t, "edit prompt", convertedReq.MultipartForm.Value["prompt"][0])
}

func TestConvertImagesEditsRequest_CanExcludeModel(t *testing.T) {
	meta := meta.NewMeta(nil, mode.ImagesEdits, "gpt-image-1", model.ModelConfig{})

	var body bytes.Buffer

	writer := multipart.NewWriter(&body)
	require.NoError(t, writer.WriteField("model", "ignored"))
	require.NoError(t, writer.WriteField("prompt", "edit prompt"))
	part, err := writer.CreateFormFile("image", "test.png")
	require.NoError(t, err)
	_, err = part.Write([]byte("png-bytes"))
	require.NoError(t, err)
	require.NoError(t, writer.Close())

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"http://example.com/v1/images/edits",
		bytes.NewReader(body.Bytes()),
	)
	require.NoError(t, err)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.ContentLength = int64(body.Len())

	result, err := ConvertImagesEditsRequest(meta, req, false)
	require.NoError(t, err)

	convertedBody, err := io.ReadAll(result.Body)
	require.NoError(t, err)

	convertedReq, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"http://example.com",
		bytes.NewReader(convertedBody),
	)
	require.NoError(t, err)
	convertedReq.Header.Set("Content-Type", result.Header.Get("Content-Type"))
	convertedReq.ContentLength = int64(len(convertedBody))

	err = convertedReq.ParseMultipartForm(1024 * 1024 * 4)
	require.NoError(t, err)

	assert.Nil(t, convertedReq.MultipartForm.Value["model"])
	assert.Equal(t, "edit prompt", convertedReq.MultipartForm.Value["prompt"][0])
}

func TestImagesStreamHandlerPassesEventsAndExtractsUsage(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"/v1/images/generations",
		nil,
	)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Header: http.Header{
			"Content-Type": {"text/event-stream"},
		},
		Body: io.NopCloser(strings.NewReader(strings.Join([]string{
			"event: " + relaymodel.ImageStreamEventPartialImage,
			`data: {"type":"` + relaymodel.ImageStreamEventPartialImage + `","partial_image_index":0,"b64_json":"partial"}`,
			"",
			"event: " + relaymodel.ImageStreamEventCompleted,
			`data: {"type":"` + relaymodel.ImageStreamEventCompleted + `","b64_json":"final","usage":{"input_tokens":10,"output_tokens":20,"total_tokens":30,"input_tokens_details":{"text_tokens":4,"image_tokens":6},"output_tokens_details":{"image_tokens":20}}}`,
			"",
		}, "\n"))),
	}

	result, err := ImagesStreamHandler(
		meta.NewMeta(nil, mode.ImagesGenerations, "gpt-image-1", model.ModelConfig{}),
		c,
		resp,
	)

	require.Nil(t, err)
	assert.Equal(t, model.ZeroNullInt64(10), result.Usage.InputTokens)
	assert.Equal(t, model.ZeroNullInt64(6), result.Usage.ImageInputTokens)
	assert.Equal(t, model.ZeroNullInt64(20), result.Usage.OutputTokens)
	assert.Equal(t, model.ZeroNullInt64(20), result.Usage.ImageOutputTokens)
	assert.Equal(t, model.ZeroNullInt64(30), result.Usage.TotalTokens)

	body := recorder.Body.String()
	assert.Contains(t, body, "event: "+relaymodel.ImageStreamEventPartialImage+"\n")
	assert.Contains(t, body, `data: {"type":"`+relaymodel.ImageStreamEventPartialImage+`"`)
	assert.Contains(t, body, "event: "+relaymodel.ImageStreamEventCompleted+"\n")
	assert.Contains(t, body, `data: {"type":"`+relaymodel.ImageStreamEventCompleted+`"`)
	assert.NotContains(t, body, `[DONE]`)
	assert.Equal(t, "text/event-stream", recorder.Header().Get("Content-Type"))
}

func TestImagesHandlerCountsReturnedImagesWhenUsageMissing(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"/v1/images/generations",
		nil,
	)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Header:     make(http.Header),
		Body: io.NopCloser(strings.NewReader(`{
			"created": 1713833628,
			"data": [
				{"url":"https://example.com/one.png"},
				{"b64_json":"two"}
			]
		}`)),
	}

	result, err := ImagesHandler(
		meta.NewMeta(nil, mode.ImagesGenerations, "dall-e-3", model.ModelConfig{}),
		c,
		resp,
	)

	require.Nil(t, err)
	assert.Equal(t, model.ZeroNullInt64(2), result.Usage.OutputTokens)
	assert.Equal(t, model.ZeroNullInt64(2), result.Usage.ImageOutputTokens)
	assert.Equal(t, model.ZeroNullInt64(2), result.Usage.TotalTokens)
}

func TestConvertVideosRequestMultipartRewritesModel(t *testing.T) {
	meta := meta.NewMeta(nil, mode.Videos, "sora-2", model.ModelConfig{})

	var body bytes.Buffer

	writer := multipart.NewWriter(&body)
	require.NoError(t, writer.WriteField("model", "ignored"))
	require.NoError(t, writer.WriteField("prompt", "Animate the reference"))
	part, err := writer.CreateFormFile("input_reference", "reference.png")
	require.NoError(t, err)
	_, err = part.Write([]byte("png-bytes"))
	require.NoError(t, err)
	require.NoError(t, writer.Close())

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"http://example.com/v1/videos",
		bytes.NewReader(body.Bytes()),
	)
	require.NoError(t, err)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.ContentLength = int64(body.Len())

	result, err := ConvertVideosRequest(meta, req)
	require.NoError(t, err)

	convertedBody, err := io.ReadAll(result.Body)
	require.NoError(t, err)

	convertedReq, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"http://example.com",
		bytes.NewReader(convertedBody),
	)
	require.NoError(t, err)
	convertedReq.Header.Set("Content-Type", result.Header.Get("Content-Type"))
	convertedReq.ContentLength = int64(len(convertedBody))

	err = convertedReq.ParseMultipartForm(1024 * 1024 * 4)
	require.NoError(t, err)

	assert.Equal(t, "sora-2", convertedReq.MultipartForm.Value["model"][0])
	assert.Equal(t, "Animate the reference", convertedReq.MultipartForm.Value["prompt"][0])
	require.Len(t, convertedReq.MultipartForm.File["input_reference"], 1)
}
