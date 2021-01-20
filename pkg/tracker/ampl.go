// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"time"

	"github.com/tellor-io/telliot/pkg/apiOracle"
)

type Ampl struct {
	granularity float64
}

func (a Ampl) Require(at time.Time) map[string]IndexProcessor {
	return map[string]IndexProcessor{
		"AMPL/USD":  VolumeWeightedAPIs(TimeWeightedAvg(24*time.Hour, NoDecay)),
		"AMPL/BTC":  AmpleChained("BTC/USD"),
		"AMPL/ETH":  TimeWeightedAvg(24*time.Hour, NoDecay),
		"AMPL/USDC": TimeWeightedAvg(24*time.Hour, NoDecay),
	}
}

func (a Ampl) ValueAt(vals map[string]apiOracle.PriceInfo, at time.Time) float64 {
	valSlice := make([]apiOracle.PriceInfo, 0, len(vals))
	for _, v := range vals {
		valSlice = append(valSlice, v)
	}
	return VolumeWeightedAvg(valSlice).Price * a.granularity
}

func (s *Ampl) Granularity() int64 {
	return int64(s.granularity)
}

// compute the average ampl price over a 24 hour period using a chained price feed.
func AmpleChained(chainedPair string) IndexProcessor {
	return func(apis []*IndexTracker, at time.Time) (apiOracle.PriceInfo, float64) {

		eod := clck.Now().UTC()
		d := 24 * time.Hour
		eod = eod.Truncate(d)
		eod = eod.Add(2 * time.Hour)

		//Get the value always at 2am UTC
		//time weight individual 10 minute buckets
		//VWAP based on time at 2am
		numVals := 0
		sum := 0.0
		maxVolume := 0.0

		interval := 10 * time.Minute

		//function to collect API values over an interval
		apiFn := VolumeWeightedAPIs(TimeWeightedAvg(interval, NoDecay))

		for i := 0; i < 144; i++ {
			thisTime := eod.Add(time.Duration(-i) * interval)
			chainedPrice, confidence := MedianAt(indexes[chainedPair], thisTime)
			if confidence < 0.01 {
				//we don't have an accurate estimate of the intermediary price, so we can't convert the AMPL price to USD
				continue
			}
			avg, confidence := apiFn(apis, thisTime)
			if confidence < 0.01 {
				//our estimate of AMPL/intermediary is not good enough right now
				continue
			}
			sum += avg.Price * chainedPrice.Price
			if avg.Volume > maxVolume {
				maxVolume = avg.Volume
			}
			numVals++
		}
		var result apiOracle.PriceInfo
		result.Price = sum / float64(numVals)
		result.Volume = maxVolume
		if sum > 0 {
			return result, float64(numVals) / 144.0

		} else {
			return result, 0
		}
	}
}
