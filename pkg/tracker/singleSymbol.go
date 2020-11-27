// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"time"

	"github.com/tellor-io/telliot/pkg/apiOracle"
)

type SingleSymbol struct {
	symbol      string
	granularity float64
	transform   IndexProcessor
}

func (s SingleSymbol) Require(at time.Time) map[string]IndexProcessor {
	r := make(map[string]IndexProcessor)
	r[s.symbol] = s.transform
	return r
}

func (s SingleSymbol) ValueAt(vals map[string]apiOracle.PriceInfo, at time.Time) float64 {
	return vals[s.symbol].Price * s.granularity
}

func (s SingleSymbol) Granularity() int64 {
	return int64(s.granularity)
}
