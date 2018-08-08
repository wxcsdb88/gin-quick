/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package handlers

import (
	"github.com/wxcsdb88/gin-quick/config"
	hlog "github.com/wxcsdb88/gin-quick/log"
	"github.com/wxcsdb88/gin-quick/log/logruslogger"
)

var (
	log hlog.Logger
)

// SetAPIHandlerLog set api handler logger
func SetAPIHandlerLog(name string, printLog bool, conf *config.GlobalConfig) {
	// log = dlog.GetLogger(name, common.PrintLog)
	log = logruslogger.GetLogger("api-handlers", true, conf)
}
