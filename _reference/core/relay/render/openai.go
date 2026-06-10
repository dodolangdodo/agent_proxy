package render

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/common/conv"
	"github.com/labring/aiproxy/core/relay/model"
)

type OpenaiSSE struct {
	Data []byte
}

func (r *OpenaiSSE) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)

	for _, bytes := range [][]byte{
		dataBytes,
		r.Data,
		nnBytes,
	} {
		// nosemgrep:
		// go.lang.security.audit.xss.no-direct-write-to-responsewriter.no-direct-write-to-responsewriter
		if _, err := w.Write(bytes); err != nil {
			return err
		}
	}

	return nil
}

func (r *OpenaiSSE) WriteContentType(w http.ResponseWriter) {
	WriteSSEContentType(w)
}

func OpenaiStringData(c *gin.Context, str string) {
	OpenaiBytesData(c, conv.StringToBytes(str))
}

func OpenaiBytesData(c *gin.Context, data []byte) {
	if len(c.Errors) > 0 {
		return
	}

	if c.IsAborted() {
		return
	}

	c.Render(-1, &OpenaiSSE{Data: data})
	c.Writer.Flush()
}

func OpenaiObjectData(c *gin.Context, object any) error {
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

	c.Render(-1, &OpenaiSSE{Data: jsonData})
	c.Writer.Flush()

	return nil
}

func OpenaiDone(c *gin.Context) {
	OpenaiStringData(c, DONE)
}

type OpenaiTtsSSE struct {
	Audio string // base64 encode audio data
	Usage *model.TextToSpeechUsage
}

func (r *OpenaiTtsSSE) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)

	payload := model.TextToSpeechSSEResponse{
		Audio: r.Audio,
		Usage: r.Usage,
	}
	if r.Usage != nil {
		payload.Type = model.TextToSpeechSSEResponseTypeDone
	} else {
		payload.Type = model.TextToSpeechSSEResponseTypeDelta
	}

	jsonData, err := sonic.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshalling object: %w", err)
	}

	for _, bytes := range [][]byte{
		dataBytes,
		jsonData,
		nnBytes,
	} {
		// nosemgrep:
		// go.lang.security.audit.xss.no-direct-write-to-responsewriter.no-direct-write-to-responsewriter
		if _, err := w.Write(bytes); err != nil {
			return err
		}
	}

	return nil
}

func (r *OpenaiTtsSSE) WriteContentType(w http.ResponseWriter) {
	WriteSSEContentType(w)
}

func OpenaiAudioData(c *gin.Context, audio string) {
	if len(c.Errors) > 0 {
		return
	}

	if c.IsAborted() {
		return
	}

	c.Render(-1, &OpenaiTtsSSE{Audio: audio})
	c.Writer.Flush()
}

type OpenaiAudioDataWriter struct {
	c *gin.Context
}

func NewOpenaiAudioDataWriter(c *gin.Context) *OpenaiAudioDataWriter {
	return &OpenaiAudioDataWriter{c: c}
}

func (w *OpenaiAudioDataWriter) Write(p []byte) (n int, err error) {
	OpenaiAudioData(w.c, base64.StdEncoding.EncodeToString(p))
	return len(p), nil
}

func OpenaiAudioDone(c *gin.Context, usage model.TextToSpeechUsage) {
	if len(c.Errors) > 0 {
		return
	}

	if c.IsAborted() {
		return
	}

	c.Render(-1, &OpenaiTtsSSE{Usage: &usage})
	c.Writer.Flush()
}
