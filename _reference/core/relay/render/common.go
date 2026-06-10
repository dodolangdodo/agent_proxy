package render

import (
	"bytes"
	"net/http"
	"slices"

	"github.com/labring/aiproxy/core/common/conv"
)

const (
	n     = "\n"
	nn    = "\n\n"
	event = "event: "
	data  = "data: "
)

var (
	nBytes     = conv.StringToBytes(n)
	nnBytes    = conv.StringToBytes(nn)
	eventBytes = conv.StringToBytes(event)
	dataBytes  = conv.StringToBytes(data)
)

const (
	DONE             = "[DONE]"
	DataPrefix       = "data:"
	DataPrefixLength = len(DataPrefix)
)

var (
	DataPrefixBytes = conv.StringToBytes(DataPrefix)
	DoneBytes       = conv.StringToBytes(DONE)
)

// IsValidSSEData checks if data is valid SSE format
func IsValidSSEData(data []byte) bool {
	return len(data) >= DataPrefixLength &&
		slices.Equal(data[:DataPrefixLength], DataPrefixBytes)
}

// ExtractSSEData extracts data from SSE format
func ExtractSSEData(data []byte) []byte {
	return bytes.TrimSpace(data[DataPrefixLength:])
}

// IsSSEDone checks if SSE data indicates completion
func IsSSEDone(data []byte) bool {
	return slices.Equal(data, DoneBytes)
}

func WriteSSEContentType(w http.ResponseWriter) {
	header := w.Header()
	if header.Get("Content-Type") == "text/event-stream" {
		return
	}

	header.Set("Content-Type", "text/event-stream")
	header.Set("Cache-Control", "no-cache")
	header.Set("Connection", "keep-alive")
	header.Set("Transfer-Encoding", "chunked")
	header.Set("X-Accel-Buffering", "no")
}
