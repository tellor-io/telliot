package tracker

import (
	"github.com/tellor-io/TellorMiner/apiOracle"
	"sort"
	"time"
)

var PSRs = map[int]ValueGenerator {
	1: SingleSymbol{symbol: "BTC/USD", multiplier:1000, transform:CurrentMedian},
	2: SingleSymbol{symbol: "BTC/USD", multiplier:1000, transform:TimeWeightedAvg(24 * time.Hour, ExpDecay)},
	3: SingleSymbol{symbol: "BTC/USD~binance.com", multiplier:1000, transform:CurrentMedian},
}

//these weight functions map values of x between 0 (brand new) and 1 (old) to weights between 0 and 1
func ExpDecay(x float64) float64 {
	return 1/math.Exp(x)
}
func LinearDecay(x float64) float64 {
	return 1 - x
}
func NoDecay(x float64) float64 {
	return 1
}

func TimeWeightedAvg(interval time.Duration, weightFn func(float64)float64) IndexProcessor {
	return func(apis []*IndexTracker) (float64, bool) {
		now := time.Now()
		sum := 0.0
		weightSum := 0.0
		totalV := 0
		for _, api := range apis {
			values := apiOracle.GetRequestValuesForTime(api.Name, now, interval)
			for _, v := range values {
				normDelta := now.Sub(v.Created).Seconds()/interval.Seconds()
				weight := weightFn(normDelta)
				sum += v.Price * weight
				weightSum += weight
			}
			totalV += len(values)
		}
		//expectedNum := len(apis)

		return sum / weightSum, false
	}
}

func CurrentMedian(apis []*IndexTracker) (float64, bool) {
	var values []*apiOracle.PriceStamp
	for _,api := range apis {
		v := apiOracle.GetLatestRequestValue(api.Name)
		if v != nil {
			values = append(values, v)
		}
	}
	if len(values) == 0 {
		return 0, false
	}
	sort.Slice(values, func (i, j int) bool {
		return values[i].Price < values[j].Price
	})
	return values[len(values)/2].Price, true
}

