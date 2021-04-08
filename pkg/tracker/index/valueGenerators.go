// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package index

import (
	"math"
	"time"

	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/apiOracle"
)

// IndexProcessor consolidates the recorded API values to a single value.
type IndexProcessor func([]*IndexTracker, time.Time, float64) (apiOracle.PriceInfo, float64, error)

type ValueGenerator interface {
	// Require reports what a PSR requires to produce a value.
	Require() map[string]IndexProcessor

	// ValueAt returns the best estimate of a value at a given time, and the confidence
	// if confidence == 0, the value has no meaning
	ValueAt(map[string]apiOracle.PriceInfo) float64

	// Granularity returns the currency granularity.
	Granularity() int64

	// Symbol returns the tracker Symbol.
	Symbol() string
}

func InitPSRs() error {
	//check that we have all the symbols asked for
	for requestID, handler := range PSRs {
		reqs := handler.Require()
		for symbol := range reqs {
			_, ok := indexes[symbol]
			if !ok {
				return errors.Errorf("requires non-existent symbol: %s on PSR: %d", symbol, requestID)
			}
		}
	}
	return nil
}

func PSRValueForTime(requestID int, at time.Time, trackersInterval float64) (float64, float64, error) {
	// Get the requirements.
	reqs := PSRs[requestID].Require()
	values := make(map[string]apiOracle.PriceInfo)
	minConfidence := math.MaxFloat64

	for symbol, fn := range reqs {
		val, confidence, err := fn(indexes[symbol], at, trackersInterval)
		if err != nil {
			return 0, 0, err
		}
		if confidence == 0 {
			return 0, 0, nil
		}
		if confidence < minConfidence {
			minConfidence = confidence
		}
		values[symbol] = val
	}

	return PSRs[requestID].ValueAt(values), minConfidence, nil
}
