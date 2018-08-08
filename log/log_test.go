package log_test

import (
	"testing"
	"time"

	"github.com/wxcsdb88/gin-quick/log"

	"github.com/wxcsdb88/gin-quick/config"
	"github.com/wxcsdb88/gin-quick/log/logruslogger"
)

func Test_Log(t *testing.T) {
	type temps struct {
		log log.Logger
	}
	globalConfig := &config.GlobalConfig{
		Log: config.LogOptions{
			WithCallerHook: true,
			Depth:          8,
			Level:          "debug",
			DisableConsole: true,
		},
	}

	a := &temps{
		log: logruslogger.GetLogger("test-logrus", true, globalConfig),
	}

	b := &temps{
		log: logruslogger.GetLogger("test-logrus2", true, globalConfig),
	}

	a.log.Debug("test %v", time.Now().UnixNano())
	a.log.Warn("test %v", time.Now().UnixNano())
	a.log.Info("test %v", time.Now().UnixNano())
	a.log.Printf("test %v", time.Now().UnixNano())
	a.log.Printf("test", time.Now().UnixNano())
	a.log.Error("test %v", time.Now().UnixNano())

	a.log.Debugln("test", time.Now().UnixNano())
	a.log.Warnln("test", time.Now().UnixNano())
	a.log.Infoln("test", time.Now().UnixNano())
	a.log.Printfln("test", time.Now().UnixNano())
	a.log.Printfln("test ", time.Now().UnixNano())
	a.log.Errorln("test", time.Now().UnixNano())

	b.log.Debug("test %v", time.Now().UnixNano())
	b.log.Warn("test %v", time.Now().UnixNano())
	b.log.Info("test %v", time.Now().UnixNano())
	b.log.Printf("test %v", time.Now().UnixNano())
	b.log.Error("test %v", time.Now().UnixNano())

	b.log.Debugln("test", time.Now().UnixNano())
	b.log.Warnln("test", time.Now().UnixNano())
	b.log.Infoln("test", time.Now().UnixNano())
	b.log.Printfln("test", time.Now().UnixNano())
	b.log.Errorln("test", time.Now().UnixNano())
}
