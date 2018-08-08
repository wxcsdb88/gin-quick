package api

import (
	"errors"
	"net/http"

	"golang.org/x/sync/errgroup"
)

// Server api server
type Server struct {
	Server *http.Server
	G      *errgroup.Group
	*API
}

// GetServer get the server
func GetServer(api *API) (server *Server) {
	serverConfig := api.config.Server
	runMode := api.config.RunMode

	currentEngineConfig := &EngineConfig{
		middleware:       nil,
		LimitConnections: serverConfig.LimitConnection,
		RunMode:          runMode,
	}

	return &Server{
		G:   api.ErrorGroup,
		API: api,
		Server: &http.Server{
			Handler:        currentEngineConfig.Init(api),
			ReadTimeout:    serverConfig.ReadTimeout,
			WriteTimeout:   serverConfig.WriteTimeout,
			IdleTimeout:    serverConfig.IdleTimeout,
			MaxHeaderBytes: serverConfig.MaxHeaderBytes,
		},
	}
}

// Run start the server
func (server *Server) Run() error {
	// run http server
	server.runServer()

	if server.config.Server.EnableHTTPS {
		if server.config.TLS.CertFile == "" || server.config.TLS.KeyFile == "" {
			return errors.New("use https should config the cert and key files")
		}
		server.runServerTLS()
	}
	if err := server.G.Wait(); err != nil {
		server.API.log.Fatalln(err)
		return err
	}
	return nil
}

// runServer run our server in a goroutine so that it doesn't block.
func (server *Server) runServer() {
	server.G.Go(func() error {
		server.API.log.Info("running server %v", server.config.Server.ListenAddr)
		return http.ListenAndServe(server.config.Server.ListenAddr, server.Server.Handler)
	})
}

// runServerTLS run our server with tls in a goroutine so that it doesn't block.
func (server *Server) runServerTLS() {
	server.G.Go(func() error {
		server.API.log.Info("running server TLS %v", server.config.Server.ListenAddr)
		return http.ListenAndServeTLS(server.config.Server.HTTPSAddr, server.config.TLS.CertFile, server.config.TLS.KeyFile, server.Server.Handler)
	})
}
