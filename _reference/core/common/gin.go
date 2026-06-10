package common

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var fieldsPool = sync.Pool{
	New: func() any {
		return make(logrus.Fields, 6)
	},
}

func GetLogFields() logrus.Fields {
	fields, ok := fieldsPool.Get().(logrus.Fields)
	if !ok {
		panic(fmt.Sprintf("fields pool type error: %T, %v", fields, fields))
	}

	return fields
}

func PutLogFields(fields logrus.Fields) {
	clear(fields)
	fieldsPool.Put(fields)
}

func GetLogger(c *gin.Context) *logrus.Entry {
	return GetLoggerFromReq(c.Request)
}

type ginLoggerKey struct{}

func GetLoggerFromReq(req *http.Request) *logrus.Entry {
	ctx := req.Context()
	if log := ctx.Value(ginLoggerKey{}); log != nil {
		v, ok := log.(*logrus.Entry)
		if !ok {
			panic(fmt.Sprintf("log type error: %T, %v", v, v))
		}

		return v
	}

	entry := NewLogger()
	SetLogger(req, entry)

	return entry
}

func SetLogger(req *http.Request, entry *logrus.Entry) {
	newCtx := context.WithValue(req.Context(), ginLoggerKey{}, entry)
	*req = *req.WithContext(newCtx)
}

func NewLogger() *logrus.Entry {
	return &logrus.Entry{
		Logger: logrus.StandardLogger(),
		Data:   GetLogFields(),
	}
}

func TruncateDuration(d time.Duration) time.Duration {
	if d > time.Hour {
		return d.Truncate(time.Minute)
	}

	if d > time.Minute {
		return d.Truncate(time.Second)
	}

	if d > time.Second {
		return d.Truncate(time.Millisecond)
	}

	if d > time.Millisecond {
		return d.Truncate(time.Microsecond)
	}

	return d
}
