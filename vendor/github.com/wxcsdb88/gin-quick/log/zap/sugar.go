package zap

import (
	"fmt"

	"github.com/wxcsdb88/gin-quick/log/types"
	zap "go.uber.org/zap"
)

// SugaredLogger zap Sugared logger
type SugaredLogger struct {
	logger *zap.SugaredLogger
}

// GetSugaredLogger get sugar logger
func GetSugaredLogger() *SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("GetSugaredLogger error for zap %v\n", err)
	}
	return &SugaredLogger{
		logger: logger.Sugar(),
	}
}

// GetLevel get level
func (l *SugaredLogger) GetLevel() types.Level {
	fmt.Println("zap sugar GetLevel")
	return 0
}

// SetLevel set level
func (l *SugaredLogger) SetLevel(level types.Level) {
	fmt.Println("zap sugar SetLevel")
}
