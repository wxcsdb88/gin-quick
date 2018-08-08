package api

import (
	"net/http"
	"time"

	"github.com/wxcsdb88/gin-quick/log/logruslogger"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"github.com/wxcsdb88/gin-quick/api/routers"
	"github.com/wxcsdb88/gin-quick/log"
)

// EngineConfig engine config
type EngineConfig struct {
	middleware       []func(*gin.Context)
	log              *log.Logger
	LimitConnections int
	RunMode          string // runMode, ex: debug,release,test
	RootRouterPrefix string // root router, default ""
}

// initEngineConfig init engine config
func (config *EngineConfig) initEngineConfig(api *API) *gin.Engine {
	if config == nil {
		panic("engine config should not be nil")
	}

	conf := api.config
	// set gin mode release(hide handlers info)
	gin.SetMode(conf.RunMode)

	e := gin.New()

	e.Use(gzip.Gzip(gzip.DefaultCompression))

	confCP := *conf
	confCP.Log.Depth = 6 // may be changed

	logger := logruslogger.GetLogger("request", true, &confCP) // re create logger
	// gin api handlers, used for api log info
	e.Use(logruslogger.Logger(logger.GetLogger()))

	// use recovery middleware
	e.Use(gin.Recovery())

	corsConfig := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	corsConfig.AllowAllOrigins = true
	e.Use(cors.New(corsConfig))

	// By default, http.ListenAndServe (which gin.Run wraps) will serve an unbounded number of requests.
	// Limiting the number of simultaneous connections can sometimes greatly speed things up under load
	if config.LimitConnections > 0 {
		e.Use(limit.MaxAllowed(config.LimitConnections))
	}

	return e
}

// Init engine init
func (config *EngineConfig) Init(api *API) http.Handler {
	e := config.initEngineConfig(api)

	// init the db
	// db.Init()
	// here init the routers, need refactor
	routers.InitRouters(e, api.config)
	return e
}
