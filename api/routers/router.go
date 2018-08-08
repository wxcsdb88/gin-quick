package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wxcsdb88/gin-quick/api/handlers"
	"github.com/wxcsdb88/gin-quick/config"
)

// InitRouters init routers
func InitRouters(e *gin.Engine, conf *config.GlobalConfig) {
	// set api handlers logger
	handlers.SetAPIHandlerLog("api-handlers", conf.Log.DisableConsole, conf)

	// routerGroup API
	rootRouterPrefix := conf.Server.RootRouterPrefix
	if rootRouterPrefix == "" {
		rootRouterPrefix = "/api"
	}
	routerGroupAPI := e.Group(rootRouterPrefix)

	demoAPIGroup := routerGroupAPI.Group("/demo")
	demoAPIGroup.GET("/ping", handlers.Ping())

}
