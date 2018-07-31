package zap

import (
	"fmt"

	"github.com/wxcsdb88/gin-quick/log/types"
	zap "go.uber.org/zap"
)

// Logger zap logger
type Logger struct {
	logger *zap.Logger
}

// GetLogger get logger
func GetLogger() *Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("GetLogger error for zap %v\n", err)
	}
	return &Logger{
		logger: logger,
	}
}

// GetLevel get level
func (l *Logger) GetLevel() types.Level {
	fmt.Println("zap GetLevel")
	return 0
}

// SetLevel set level
func (l *Logger) SetLevel(level types.Level) {
	fmt.Println("zap SetLevel")
}
