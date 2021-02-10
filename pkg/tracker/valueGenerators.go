// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/apiOracle"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/db"
)

// IndexProcessor consolidates the recorded API values to a single value.
type IndexProcessor func([]*IndexTracker, time.Time) (apiOracle.PriceInfo, float64, error)

type ValueGenerator interface {
	// Require reports what a PSR requires to produce a value.
	Require(time.Time) map[string]IndexProcessor

	// ValueAt returns the best estimate of a value at a given time, and the confidence
	// if confidence == 0, the value has no meaning
	ValueAt(map[string]apiOracle.PriceInfo, time.Time) float64

	// Granularity returns the currency granularity.
	Granularity() int64
}

func InitPSRs() error {
	//check that we have all the symbols asked for
	now := clck.Now()
	for requestID, handler := range PSRs {
		reqs := handler.Require(now)
		for symbol := range reqs {
			_, ok := indexes[symbol]
			if !ok {
				return errors.Errorf("requires non-existent symbol: %s on PSR: %d", symbol, requestID)
			}
		}
	}
	return nil
}

func PSRValueForTime(requestID int, at time.Time) (float64, float64, error) {
	// Get the requirements.
	reqs := PSRs[requestID].Require(at)
	values := make(map[string]apiOracle.PriceInfo)
	minConfidence := math.MaxFloat64

	for symbol, fn := range reqs {
		val, confidence, err := fn(indexes[symbol], at)
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

	return PSRs[requestID].ValueAt(values, at), minConfidence, nil
}

func UpdatePSRs(ctx context.Context, DB db.DataServerProxy, updatedSymbols []string) error {
	now := clck.Now()
	// Generate a set of all affected PSRs.
	var toUpdate []int
	for requestID, psr := range PSRs {
		reqs := psr.Require(now)
		for _, symbol := range updatedSymbols {
			_, ok := reqs[symbol]
			if ok {
				toUpdate = append(toUpdate, requestID)
				break
			}
		}
	}

	// Update all affected PSRs.
	for _, requestID := range toUpdate {
		amt, conf, err := PSRValueForTime(requestID, now)
		if err != nil {
			return err
		}

		cfg, err := config.ParseConfig("")
		if err != nil {
			return errors.Wrapf(err, "parsing config")
		}
		if conf < cfg.Trackers.MinConfidence || math.IsNaN(amt) {
			// Confidence in this signal is too low to use.
			continue
		}

		// Convert it directly from a float to a bigInt so that there is no risk of overflowing a uint64.
		bigVal := new(big.Float)
		bigVal.SetFloat64(amt)
		bigInt := new(big.Int)
		bigVal.Int(bigInt)
		// Encode it and store to DB.
		enc := hexutil.EncodeBig(bigInt)
		err = DB.Put(fmt.Sprintf("%s%d", db.QueriedValuePrefix, requestID), []byte(enc))
		if err != nil {
			return err
		}
	}
	return nil
}
