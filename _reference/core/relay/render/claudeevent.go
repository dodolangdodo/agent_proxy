package render

import (
	"net/http"

	"github.com/labring/aiproxy/core/common"
	"github.com/labring/aiproxy/core/common/conv"
)

type Anthropic struct {
	Event string
	Data  []byte
}

func (r *Anthropic) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)

	event := r.Event

	if event == "" {
		eventNode, err := common.GetJSONNodeNoCopy(r.Data, "type")
		if err != nil {
			return err
		}

		event, err = eventNode.String()
		if err != nil {
			return err
		}
	}

	for _, bytes := range [][]byte{
		eventBytes,
		conv.StringToBytes(event),
		nBytes,
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

func (r *Anthropic) WriteContentType(w http.ResponseWriter) {
	WriteSSEContentType(w)
}
