package log

import (
	"sync"

	"github.com/wxcsdb88/gin-quick/log/types"
)

const (
	LoggerLogrus = "logrus"
)

var (
	defaultLogger = LoggerLogrus

	// loggerMap map[string]*GlobalLogger
	mutex sync.Mutex
)

// loggerMap map[string]*GlobalLogger
var loggerMap = make(map[string]*interface{})

func init() {
	// loggerMap[LoggerLogrus] = logrus.GetLogger()

}

var (
	loggersMu  sync.RWMutex
	loggers    = make(map[string]Logger)
	loggersMap = make(map[string]Logger) // module:Logger
)

func Register(name string, logger Logger) {
	loggersMu.Lock()
	defer loggersMu.Unlock()
	if logger == nil {
		panic("logger: Register logger is nil")
	}
	if _, dup := loggers[name]; dup {
		panic("sql: Register called twice for logger " + name)
	}
	loggers[name] = logger
}

// Logger global logger
type Logger interface {
	GetLevel() types.Level
	SetLevel(level types.Level)

	// Debugf(format string, args ...interface{})
	// Infof(format string, args ...interface{})
	// Printf(format string, args ...interface{})
	// Warnf(format string, args ...interface{})
	// Warningf(format string, args ...interface{})
	// Errorf(format string, args ...interface{})
	// Fatalf(format string, args ...interface{})
	// Panicf(format string, args ...interface{})

	// Debug(args ...interface{})
	// Info(args ...interface{})
	// Print(args ...interface{})
	// Warn(args ...interface{})
	// Warning(args ...interface{})
	// Error(args ...interface{})
	// Fatal(args ...interface{})
	// Panic(args ...interface{})

	// Debugln(args ...interface{})
	// Infoln(args ...interface{})
	// Println(args ...interface{})
	// Warnln(args ...interface{})
	// Warningln(args ...interface{})
	// Errorln(args ...interface{})
	// Fatalln(args ...interface{})
	// Panicln(args ...interface{})
}
