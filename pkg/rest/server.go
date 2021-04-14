// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/logging"
)

const ComponentName = "rest"

// Server wraps http server with pre-configured paths.
type Server struct {
	server    *http.Server
	dataProxy db.DataServerProxy
	logger    log.Logger
}

// Create a new server instance for the given host/port.
func Create(logger log.Logger, cfg *config.Config, ctx context.Context, proxy db.DataServerProxy, host string, port uint) (*Server, error) {
	srv := &http.Server{Addr: fmt.Sprintf("%s:%d", host, port)}

	remoteHandler, err := CreateRemoteProxy(logger, cfg, ctx, proxy)
	if err != nil {
		return nil, err
	}
	if remoteHandler == nil {
		return nil, errors.Errorf("create a remote proxy")
	}

	logger, err = logging.ApplyFilter(*cfg, ComponentName, logger)
	if err != nil {
		return nil, errors.Wrap(err, "applying filter to looger")
	}

	http.Handle("/", remoteHandler)
	return &Server{
		server:    srv,
		dataProxy: proxy,
		logger:    log.With(logger, "component", ComponentName),
	}, nil
}

// Start the server listening for incoming requests.
func (s *Server) Start() error {
	level.Info(s.logger).Log("msg", "starting server", "addr", s.server.Addr)
	// returns ErrServerClosed on graceful close
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		// NOTE: there is a chance that next line won't have time to run,
		// as main() doesn't wait for this goroutine to stop. don't use
		// code with race conditions like these for production. see post
		// comments below on more discussion on how to handle this.
		// TODO remove this log and return error instead.
		level.Error(s.logger).Log("msg", "ListenAndServe()", "err", err)
		return err
	}
	return nil
}

// Stop stops the server listening.
func (s *Server) Stop() error {
	level.Info(s.logger).Log("msg", "stopping server")
	return s.server.Close()
}
