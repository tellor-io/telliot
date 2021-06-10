// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tellorAccess

import (
	"math"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/aggregator"
)

const (
	ComponentName      = "psrTellorAccess"
	DefaultGranularity = 1000000
)

func New(logger log.Logger, cfg Config, aggregator *aggregator.Aggregator) *Psr {
	return &Psr{
		logger:     log.With(logger, "component", ComponentName),
		aggregator: aggregator,
		cfg:        cfg,
	}
}

type Config struct {
	MinConfidence float64
}

type Psr struct {
	logger     log.Logger
	aggregator *aggregator.Aggregator
	cfg        Config
}

func (self *Psr) GetValue(reqID int64, ts time.Time) (int64, error) {
	val, err := self.getValue(reqID, ts)
	return int64(math.Round(val * DefaultGranularity)), err
}

func (self *Psr) getValue(reqID int64, ts time.Time) (float64, error) {
	val, err := self.aggregator.ManualValue("tellorAccess", reqID, ts)
	if err != nil {
		level.Error(self.logger).Log("msg", "get manual value", "reqID", reqID, "err", err)
	}
	if val != 0 {
		level.Warn(self.logger).Log("msg", "USING MANUAL VALUE", "reqID", reqID, "val", val)
		return val, nil
	}

	var conf float64
	switch reqID {
	case 1:
		val, conf, err = self.aggregator.MedianAt("ETH/USD", ts)
	case 2:
		val, conf, err = self.aggregator.MedianAt("BTC/USD", ts)
	default:
		return 0, errors.Errorf("undeclared request ID:%v", reqID)
	}

	if err != nil {
		return 0, err
	}

	if conf < self.cfg.MinConfidence {
		return 0, errors.Errorf("not enough confidence - value:%v, conf:%v,confidence threshold:%v", val, conf, self.cfg.MinConfidence)
	}

	return val, err
}
