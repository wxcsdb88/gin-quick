package logrus

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	logger *logrus.Logger
}

func GetLogger() *LogrusLogger {
	return &LogrusLogger{}
}

func (l *LogrusLogger) GetLevel() {
	fmt.Printf("logrus")
}

// func (l *LogrusLogger) SetLevel(level logrus.Level) {
// 	l.logger.SetLevel(level)
// }

// func (l *LogrusLogger) GetLevel() logrus.Level {
// 	return l.logger.Level
// }
