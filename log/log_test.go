package log

import (
	"fmt"
	"testing"

	"github.com/wxcsdb88/gin-quick/log/logrus"
	"github.com/wxcsdb88/gin-quick/log/zap"
)

func TestLogger(t *testing.T) {
	var logrusLogger = logrus.GetLogger()

	Register("logrus", logrusLogger)
	fmt.Printf("loggers is %#v\n", loggers)
	fmt.Printf("GetLevel %#v\n", logrusLogger.GetLevel())

	var zapLogger = zap.GetLogger()
	Register("zap", zapLogger)
	fmt.Printf("loggers is %#v\n", loggers)
	fmt.Printf("GetLevel %#v\n", zapLogger.GetLevel())

}
