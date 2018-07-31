package log

import (
	"fmt"
	"testing"

	"github.com/wxcsdb88/gin-quick/log/logrus"
	"github.com/wxcsdb88/gin-quick/log/zap"
)

func TestLogger(t *testing.T) {
	// var globalLogger *Logger
	// var globalLogger Logger = &logrus.LogrusLogger{}
	var globalLogger = logrus.GetLogger()

	Register("logrus", globalLogger)
	fmt.Printf("loggers is %#v\n", loggers)

	var zapLogger = zap.GetLogger()
	Register("zap", zapLogger)
	fmt.Printf("loggers is %#v\n", loggers)

}
