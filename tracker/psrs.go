package tracker

import (
	"github.com/tellor-io/TellorMiner/apiOracle"
	"github.com/tellor-io/TellorMiner/config"
	"math"
	"sort"
	"time"
)

var switchTime, _ = time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")

var PSRs = map[int]ValueGenerator {
	1: &SingleSymbol{symbol: "BTC/USD", granularity:1000, transform:CurrentMedian},
	2: &SingleSymbol{symbol: "BTC/USD", granularity:1000, transform:TimeWeightedAvg(24*time.Hour, ExpDecay)},
	3: &SingleSymbol{symbol: "BTC/USD~api.binance.com", granularity:1000, transform:CurrentMean},
	4: &TimedSwitch{
		before:SingleSymbol{symbol:"ETH/USD", granularity: 10000, transform:CurrentMedian},
		after:SingleSymbol{symbol:"ETH/USD", granularity: 10000, transform:TimeWeightedAvg(10*time.Minute, NoDecay)},
		at: switchTime,
	},
	//computes the TRB/USD price indirectly
	5: &ChainedPrice{chain: []string{"TRB/ETH", "ETH/BTC", "BTC/USD"}},
}

//these weight functions map values of x between 0 (brand new) and 1 (old) to weights between 0 and 1
//also returns the integral of the weight over the range [0,1]
//weights the oldest data (1) as being 1/3 as important (1/e)
func ExpDecay(x float64) (float64, float64) {
	return 1/math.Exp(x), 0.63212
}
//weights the oldest data at 0
func LinearDecay(x float64) (float64,float64) {
	return 1 - x, 0.5
}
//weights all data in the time interval evenly
func NoDecay(x float64) (float64, float64) {
	return 1, 1
}

func TimeWeightedAvg(interval time.Duration, weightFn func(float64)(float64, float64)) IndexProcessor {
	return func(apis []*IndexTracker, at time.Time) (float64, float64) {
		cfg := config.GetConfig()
		sum := 0.0
		weightSum := 0.0
		for _, api := range apis {
			values := apiOracle.GetRequestValuesForTime(api.Identifier, at, interval)
			for _, v := range values {
				normDelta := at.Sub(v.Created).Seconds()/interval.Seconds()
				weight,_ := weightFn(normDelta)
				sum += v.Price * weight
				weightSum += weight
			}
		}
		// number of APIs * rate * interval
		maxWeight := float64(len(apis)) * (1/cfg.TrackerSleepCycle.Duration.Seconds()) * interval.Seconds()

		//average weight is the integral of the weight fn over [0,1]
		_,avgWeight :=  weightFn(0)
		targetWeight := maxWeight * avgWeight
		return sum / weightSum, math.Min(weightSum/targetWeight,1.0)
	}
}

func getLatest(apis []*IndexTracker, at time.Time) ([]*apiOracle.PriceStamp, float64) {
	var values []*apiOracle.PriceStamp
	totalConf := 0.0
	for _,api := range apis {
		b,_ := apiOracle.GetNearestTwoRequestValue(api.Identifier, at)
		if b != nil {
			//penalize values more than 5 minutes old
			totalConf += math.Min(5/at.Sub(b.Created).Minutes(), 1.0)
			values = append(values, b)
		}
	}
	return values, totalConf/float64(len(apis))
}


func CurrentMedian(apis []*IndexTracker, at time.Time) (float64, float64) {
	values, confidence := getLatest(apis, at)
	if confidence == 0 {
		return 0, 0
	}
	sort.Slice(values, func (i, j int) bool {
		return values[i].Price < values[j].Price
	})
	return values[len(values)/2].Price, confidence
}

func CurrentMean(apis []*IndexTracker, at time.Time) (float64, float64) {
	values, confidence := getLatest(apis, at)
	if confidence == 0 {
		return 0, 0
	}
	sum := 0.0
	for _,val := range values {
		sum += val.Price
	}
	return sum/float64(len(values)), confidence
}
