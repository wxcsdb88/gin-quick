package zap

import (
	"fmt"
)

type ZapLogger struct {
}

func GetLogger() *ZapLogger {
	return &ZapLogger{}
}

func (l *ZapLogger) GetLevel() {
	fmt.Printf("zap")
}
