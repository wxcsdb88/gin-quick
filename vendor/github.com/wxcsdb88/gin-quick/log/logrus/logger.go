package logrus

import (
	"fmt"

	logrus "github.com/sirupsen/logrus"
	"github.com/wxcsdb88/gin-quick/log/types"
)

// Logger logrus logger
type Logger struct {
	logger *logrus.Logger
}

// GetLogger get logger
func GetLogger() *Logger {
	return &Logger{
		logger: logrus.New(),
	}
}

// GetLevel get level
func (l *Logger) GetLevel() types.Level {
	fmt.Println("logrus GetLevel")
	return 0
}

// SetLevel set level
func (l *Logger) SetLevel(level types.Level) {
	fmt.Println("logrus SetLevel")
}

// func (l *LogrusLogger) SetLevel(level logrus.Level) {
// 	l.logger.SetLevel(level)
// }

// func (l *LogrusLogger) GetLevel() logrus.Level {
// 	return l.logger.Level
// }
