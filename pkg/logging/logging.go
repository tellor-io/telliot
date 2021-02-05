// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package logging

import (
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
)

// NewLogger create a new logger.
func NewLogger() log.Logger {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	return log.With(logger, "ts", log.TimestampFormat(func() time.Time { return time.Now().UTC() }, "Jan 02 15:04:05.99"), "caller", log.DefaultCaller)
}

// ApplyFilter applies a filter to logger based on component name.
func ApplyFilter(cfg config.Config, componentName string, logger log.Logger) (log.Logger, error) {
	lvl := level.AllowInfo()
	if configLevel, ok := cfg.Logger[componentName]; ok {
		switch configLevel {
		case "error":
			lvl = level.AllowError()
		case "warn":
			lvl = level.AllowWarn()
		case "info":
			lvl = level.AllowInfo()
		case "debug":
			lvl = level.AllowDebug()
		default:
			return nil, errors.Errorf("unexpected log level:%v", configLevel)
		}
	}

	return level.NewFilter(logger, lvl), nil
}
