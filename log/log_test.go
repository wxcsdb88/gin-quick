package log

import (
	"fmt"
	"testing"

	logrus "github.com/wxcsdb88/gin-quick/log/logrus"
	"github.com/wxcsdb88/gin-quick/log/zap"
)

func TestLogger(t *testing.T) {
	// var globalLogger *Logger
	var globalLogger Logger = &logrus.LogrusLogger{}
	// globalLogger = logrus.GetLogger()

	Register("logrus", globalLogger)
	fmt.Printf("loggers is %#v\n", loggers)

	var zapLogger Logger = &zap.ZapLogger{}
	Register("zap", zapLogger)
	fmt.Printf("loggers is %#v\n", loggers)

	loggers["zap"].GetLevel()
	loggers["logrus"].GetLevel()
	loggers["zap"].GetLevel()

}
