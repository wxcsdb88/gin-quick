package logruslogger

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/wxcsdb88/gin-quick/log"
)

// LogrusLogger logrus logger
type LogrusLogger struct {
	log *logrus.Logger
}

// Debug wrapper Debug logger
func (m *LogrusLogger) Debug(f interface{}, args ...interface{}) {
	m.log.Debug(log.FormatLog(f, args...))
}

// Info wrapper Info logger
func (m *LogrusLogger) Info(f interface{}, args ...interface{}) {
	m.log.Info(log.FormatLog(f, args...))
}

// Warn wrapper Warn logger
func (m *LogrusLogger) Warn(f interface{}, args ...interface{}) {
	m.log.Warn(log.FormatLog(f, args...))
}

// Printf wrapper Printf logger
func (m *LogrusLogger) Printf(f interface{}, args ...interface{}) {
	m.log.Print(log.FormatLog(f, args...))
}

// Panic wrapper Panic logger
func (m *LogrusLogger) Panic(f interface{}, args ...interface{}) {
	m.log.Panic(log.FormatLog(f, args...))
}

// Fatal wrapper Fatal logger
func (m *LogrusLogger) Fatal(f interface{}, args ...interface{}) {
	m.log.Fatal(log.FormatLog(f, args...))
}

// Error wrapper Error logger
func (m *LogrusLogger) Error(f interface{}, args ...interface{}) {
	m.log.Error(log.FormatLog(f, args...))
}

// Debugln wrapper Debugln logger
func (m *LogrusLogger) Debugln(v ...interface{}) {
	m.log.Debug(fmt.Sprintln(v...))
}

// Infoln wrapper Infoln logger
func (m *LogrusLogger) Infoln(args ...interface{}) {
	m.log.Info(fmt.Sprintln(args...))
}

// Warnln wrapper Warnln logger
func (m *LogrusLogger) Warnln(args ...interface{}) {
	m.log.Warn(fmt.Sprintln(args...))
}

// Printfln wrapper Printfln logger
func (m *LogrusLogger) Printfln(args ...interface{}) {
	m.log.Print(fmt.Sprintln(args...))
}

// Panicln wrapper Panicln logger
func (m *LogrusLogger) Panicln(args ...interface{}) {
	m.log.Panic(fmt.Sprintln(args...))
}

// Fatalln wrapper Fatalln logger
func (m *LogrusLogger) Fatalln(args ...interface{}) {
	m.log.Fatal(fmt.Sprintln(args...))
}

// Errorln wrapper Errorln logger
func (m *LogrusLogger) Errorln(args ...interface{}) {
	m.log.Error(fmt.Sprintln(args...))
}
