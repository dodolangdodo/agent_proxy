package common

import (
	"fmt"
	stdlog "log"
	"os"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var logCallerIgnoreFuncs = map[string]struct{}{
	"github.com/labring/aiproxy/core/middleware.logColor": {},
}

func InitLog(l *log.Logger, debug bool) {
	gin.ForceConsoleColor()

	if debug {
		l.SetLevel(log.DebugLevel)
		l.SetReportCaller(true)
		gin.SetMode(gin.DebugMode)
	} else {
		l.SetLevel(log.InfoLevel)
		l.SetReportCaller(false)
		gin.SetMode(gin.ReleaseMode)
	}

	l.SetOutput(os.Stdout)
	stdlog.SetOutput(l.Writer())

	l.SetFormatter(&log.TextFormatter{
		ForceColors:      true,
		DisableColors:    false,
		ForceQuote:       debug,
		DisableQuote:     !debug,
		DisableSorting:   false,
		FullTimestamp:    true,
		TimestampFormat:  time.DateTime,
		QuoteEmptyFields: true,
		CallerPrettyfier: func(f *runtime.Frame) (function, file string) {
			if _, ok := logCallerIgnoreFuncs[f.Function]; ok {
				return "", ""
			}
			return f.Function, fmt.Sprintf("%s:%d", f.File, f.Line)
		},
	})

	if NeedColor() {
		gin.ForceConsoleColor()
	}
}
