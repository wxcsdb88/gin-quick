package logruslogger

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/wxcsdb88/gin-quick/config"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var logMap map[string]*LogrusLogger
var getLogMutex sync.Mutex

// Options config the logger
type Options struct {
	PrintLog       *bool
	WriteLog       *bool
	LogLevel       *logrus.Level
	Depth          *int
	WithCallerHook *bool
}

// GetLoggerWithOptions with options config
func GetLoggerWithOptions(logName string, options *Options, conf *config.GlobalConfig) *LogrusLogger {
	getLogMutex.Lock()
	defer getLogMutex.Unlock()

	if logMap == nil {
		logMap = make(map[string]*LogrusLogger)
	}
	curLog, ok := logMap[logName]

	if ok {
		return curLog
	}

	log := logrus.New()

	logConf := conf.Log
	defaultLogFilePrex := logConf.LogFilePrefix

	commonConf := conf.Common

	// get logLevel
	logLevel := GetLogLevel(logConf.Level)

	// logger config set
	printLog := !logConf.DisableConsole
	writeLog := logConf.Write
	depth := logConf.Depth
	maxAge := logConf.MaxAge
	rotationTime := logConf.RotationTime

	withCallerHook := logConf.WithCallerHook
	tempFolder := commonConf.TempFolder

	// options set
	cpOptions := *options
	if &cpOptions != nil {
		if cpOptions.PrintLog != nil {
			printLog = *cpOptions.PrintLog
		}
		if cpOptions.WriteLog != nil {
			writeLog = *cpOptions.WriteLog
		}
		if cpOptions.LogLevel != nil {
			logLevel = *cpOptions.LogLevel
		}
		if cpOptions.Depth != nil {
			depth = *cpOptions.Depth
		}
		if cpOptions.WithCallerHook != nil {
			withCallerHook = *cpOptions.WithCallerHook
		}
	}

	log.SetLevel(logLevel)

	if writeLog {
		storeLogDir := filepath.Join(tempFolder, logConf.LogDir)

		err := os.MkdirAll(storeLogDir, os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("creating log file failed: %s", err.Error()))
		}

		path := filepath.Join(storeLogDir, logConf.LogFileName)
		writer, err := rotatelogs.New(
			path+".%Y%m%d%H%M%S",
			rotatelogs.WithClock(rotatelogs.Local),
			rotatelogs.WithMaxAge(time.Duration(maxAge)*time.Hour),
			rotatelogs.WithRotationTime(time.Duration(rotationTime)*time.Hour),
		)
		if err != nil {
			panic(fmt.Sprintf("rotatelogs log failed: %s", err.Error()))
		}

		var formatter logrus.Formatter

		formatter = &logrus.TextFormatter{}
		if conf.Log.Formatter == "json" {
			formatter = &logrus.JSONFormatter{}
		}

		log.AddHook(lfshook.NewHook(
			lfshook.WriterMap{
				logrus.DebugLevel: writer,
				logrus.InfoLevel:  writer,
				logrus.WarnLevel:  writer,
				logrus.ErrorLevel: writer,
				logrus.FatalLevel: writer,
			},
			formatter,
		))

		pathMap := lfshook.PathMap{
			logrus.DebugLevel: fmt.Sprintf("%s/%sdebug.log", storeLogDir, defaultLogFilePrex),
			logrus.InfoLevel:  fmt.Sprintf("%s/%sinfo.log", storeLogDir, defaultLogFilePrex),
			logrus.WarnLevel:  fmt.Sprintf("%s/%swarn.log", storeLogDir, defaultLogFilePrex),
			logrus.ErrorLevel: fmt.Sprintf("%s/%serror.log", storeLogDir, defaultLogFilePrex),
			logrus.FatalLevel: fmt.Sprintf("%s/%sfatal.log", storeLogDir, defaultLogFilePrex),
		}
		log.AddHook(lfshook.NewHook(
			pathMap,
			formatter,
		))
	} else {
		if printLog {
			log.Out = os.Stdout
		}
		fmt.Printf("disableconsole .... %v\n", printLog)
	}

	if withCallerHook {
		log.AddHook(&CallerHook{depth: depth, module: logName}) // add caller hook to print caller's file and line number
	}
	curLog = &LogrusLogger{
		log: log,
	}
	logMap[logName] = curLog
	fmt.Printf("register logger %v, current loggers: %v\n", logName, logMap)
	return curLog
}

// GetLogger gets logrus.Logger object according to logName
// each module can have its own logger
func GetLogger(logName string, printConsole bool, conf *config.GlobalConfig) *LogrusLogger {
	// writeFile flag use global config
	return GetLoggerWithOptions(logName, &Options{PrintLog: &printConsole}, conf)
}

// GetLoggerWithCaller gets logrus.Logger object according to logName
// with paramters printConsole and withCaller
// Used for caller control
func GetLoggerWithCaller(logName string, printConsole bool, withCaller bool, conf *config.GlobalConfig) *LogrusLogger {
	// writeFile flag use global config
	return GetLoggerWithOptions(logName, &Options{PrintLog: &printConsole, WithCallerHook: &withCaller}, conf)
}
