package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/util"
)

var serverLog = util.NewLogger("rest", "Server")

//Server wraps http server with pre-configured paths
type Server struct {
	server    *http.Server
	dataProxy db.DataServerProxy
}

//Create a new server instance for the given host/port
func Create(ctx context.Context, host string, port uint) (*Server, error) {
	proxy := ctx.Value(common.DataProxyKey).(db.DataServerProxy)
	srv := &http.Server{Addr: fmt.Sprintf("%s:%d", host, port)}

	remoteHandler, err := CreateRemoteProxy(ctx)
	if err != nil {
		return nil, err
	}
	if remoteHandler == nil {
		return nil, fmt.Errorf("Could not create a remote proxy")
	}

	http.Handle("/", remoteHandler)
	return &Server{server: srv, dataProxy: proxy}, nil
}

//Start the server listening for incoming requests
func (s *Server) Start() {
	go func() {
		serverLog.Info("Starting server on %+v\n", s.server.Addr)
		// returns ErrServerClosed on graceful close
		if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
			// NOTE: there is a chance that next line won't have time to run,
			// as main() doesn't wait for this goroutine to stop. don't use
			// code with race conditions like these for production. see post
			// comments below on more discussion on how to handle this.
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()
}

//Stop stops the server listening
func (s *Server) Stop() error {
	serverLog.Info("Stopping server")
	return s.server.Close()
}
