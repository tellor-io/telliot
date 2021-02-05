// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package rest

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/logging"
)

// RemoteProxyRouter handles incoming http requests.
type RemoteProxyRouter struct {
	dataProxy db.DataServerProxy
	logger    log.Logger
}

// CreateRemoteProxy creates a remote proxy instance.
func CreateRemoteProxy(logger log.Logger, cfg *config.Config, ctx context.Context, proxy db.DataServerProxy) (*RemoteProxyRouter, error) {
	logger, err := logging.ApplyFilter(*cfg, ComponentName, logger)
	if err != nil {
		return nil, err
	}
	return &RemoteProxyRouter{
		dataProxy: proxy,
		logger:    log.With(logger, "component", ComponentName),
	}, nil
}

// Default http handler callback which will route to appropriate handler internally.
func (r *RemoteProxyRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// The "/" pattern matches everything, so this is to disable any other requests.
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	w.Header().Add("Content-Type", "application/octet-stream")

	if e := recover(); e != nil {
		level.Info(r.logger).Log("msg", "proxied data request", "err", e)
		fmt.Fprintf(w, "Cannot serve request")
		return
	}
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		level.Error(r.logger).Log("msg", "reading request", "err", err)
		fmt.Fprint(w, "read request data")
		return
	}
	level.Info(r.logger).Log("msg", "getting request", "bytes", len(data))
	outData, err := r.dataProxy.IncomingRequest(data)

	if err != nil {
		level.Error(r.logger).Log("msg", "handling incoming request", "err", err)
		fmt.Fprint(w, "handle request")
		return
	}
	level.Info(r.logger).Log("msg", "produced result", "bytes", len(outData))
	w.WriteHeader(200)
	_, err = w.Write(outData)
	if err != nil {
		level.Error(r.logger).Log("msg", "write response", "err", err)
		http.Error(w, "write the response", 500)
		return
	}
}
