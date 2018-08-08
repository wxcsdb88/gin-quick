package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wxcsdb88/gin-quick/api/common"
)

// Ping return without content
func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Debug("ping %v", time.Now().Local())
		log.Info("ping %v", time.Now().Local())
		log.Warn("ping %v", time.Now().Local())
		responseData := common.NewResponseData(200, "server is ok", nil, c.Request.RequestURI)
		ResponseJSON(c, responseData)
	}
}
