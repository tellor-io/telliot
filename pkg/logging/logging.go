// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package logging

import (
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
)

// NewLogger create a new logger.
func NewLogger() log.Logger {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	return log.With(logger, "ts", log.TimestampFormat(func() time.Time { return time.Now().UTC() }, "jan 02 15:04:05.00"), "caller", log.Caller(5))
}

// ApplyFilter applies a filter to logger based on component name.
func ApplyFilter(configLevel string, logger log.Logger) (log.Logger, error) {
	var lvl level.Option
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

	return level.NewFilter(logger, lvl), nil
}
