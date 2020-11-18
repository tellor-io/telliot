// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"github.com/tellor-io/TellorMiner/pkg/util"
)

var serverLog = util.NewLogger("rest", "Server")

// Server wraps http server with pre-configured paths.
type Server struct {
	server    *http.Server
	dataProxy db.DataServerProxy
}

// Create a new server instance for the given host/port.
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

// Start the server listening for incoming requests.
func (s *Server) Start() error {
	go func() {
		serverLog.Info("Starting server on %+v\n", s.server.Addr)
		return s.server.ListenAndServe()
	}()
}

// Stop stops the server listening.
func (s *Server) Stop() error {
	serverLog.Info("Stopping server")
	return s.server.Close()
}
