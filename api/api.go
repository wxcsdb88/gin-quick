package api

import (
	"golang.org/x/sync/errgroup"

	"github.com/wxcsdb88/gin-quick/config"
	"github.com/wxcsdb88/gin-quick/log"
	"github.com/wxcsdb88/gin-quick/log/logruslogger"
)

var (
	eGroup errgroup.Group
)

// API api server
type API struct {
	config     *config.GlobalConfig
	log        log.Logger
	ErrorGroup *errgroup.Group
}

// New create new API
func New(conf *config.GlobalConfig) (*API, error) {
	confCopy := *conf
	alog := logruslogger.GetLogger("api", true, conf)

	return &API{
		config:     &confCopy,
		log:        alog,
		ErrorGroup: &eGroup,
	}, nil
}

// Start start the api server
func (api *API) Start() error {
	//run server with config and logs
	server := GetServer(api)
	return server.Run()
}
