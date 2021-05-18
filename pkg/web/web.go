package web

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/route"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/storage"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/util"
	"github.com/tellor-io/telliot/pkg/web/api"
)

const ComponentName = "web"

type Config struct {
	LogLevel    string
	ListenHost  string
	ListenPort  uint
	ReadTimeout util.Duration
}

type Web struct {
	logger log.Logger
	cfg    Config
	ctx    context.Context
	stop   context.CancelFunc
	srv    *http.Server
}

func New(logger log.Logger, ctx context.Context, tsDB storage.SampleAndChunkQueryable, cfg Config) (*Web, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	router := route.New()
	router.Get("/metrics", promhttp.Handler().ServeHTTP)

	opts := promql.EngineOpts{
		Logger:               logger,
		Reg:                  nil,
		MaxSamples:           100000,
		Timeout:              10 * time.Second,
		EnableAtModifier:     true,
		EnableNegativeOffset: true,
	}
	engine := promql.NewEngine(opts)

	api := api.New(logger, ctx, engine, tsDB)
	api.Register(router.WithPrefix("/api/v1"))

	mux := http.NewServeMux()
	mux.Handle("/", router)

	srv := &http.Server{
		Handler:     mux,
		ReadTimeout: cfg.ReadTimeout.Duration,
		Addr:        fmt.Sprintf("%s:%d", cfg.ListenHost, cfg.ListenPort),
	}

	ctx, stop := context.WithCancel(ctx)

	return &Web{
		logger: log.With(logger, "component", ComponentName),
		cfg:    cfg,
		ctx:    ctx,
		stop:   stop,
		srv:    srv,
	}, nil

}

func (self *Web) Start() error {
	level.Info(self.logger).Log("msg", "starting", "addr", self.srv.Addr)
	if err := self.srv.ListenAndServe(); err != http.ErrServerClosed {
		return errors.Wrapf(err, "ListenAndServe")
	}
	return nil
}

func (self *Web) Stop() {
	self.stop()
	if err := self.srv.Close(); err != nil {
		level.Error(self.logger).Log("msg", "closing srv", "err", err)
	}
}
