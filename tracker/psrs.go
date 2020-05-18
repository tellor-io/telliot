package tracker

import (
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/tellor-io/TellorMiner/apiOracle"
	"github.com/tellor-io/TellorMiner/config"
)

var switchTime, _ = time.Parse(time.RFC3339, "2010-06-01T00:00:00+00:00")

var PSRs = map[int]ValueGenerator{
	// 1: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "ETH/USD~api.pro.coinbase.com", granularity: 1000, transform: MeanAt},
	// 	after:  &SingleSymbol{symbol: "ETH/USD", granularity: 1000000, transform: MedianAt},
	// 	at:     switchTime,
	// },
	//2: &TimedSwitch{
	//	before: &SingleSymbol{symbol: "BTC/USD~api.binance.com", granularity: 1000, transform: MeanAt},
	//	after:  &SingleSymbol{symbol: "BTC/USD", granularity: 1000000, transform: MedianAt},
	//	at:     switchTime,
	//},
	// 3: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "BNB/USD~dex.binance.org", granularity: 1000, transform: MeanAt},
	// 	after:  &SingleSymbol{symbol: "BNB/USD", granularity: 1000000, transform: MedianAt},
	// 	at:     switchTime,
	// },
	// 4: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "BTC/USD", granularity: 100000, transform: MedianAt},
	// 	after:  &SingleSymbol{symbol: "BTC/USD", granularity: 1000000, transform: TimeWeightedAvg(24*time.Hour, ExpDecay)},
	// 	at:     switchTime,
	// },
	// 5: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "ETH/BTC~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 	after:  &SingleSymbol{symbol: "ETH/BTC", granularity: 1000000, transform: MedianAt},
	// 	at:     switchTime,
	// },
	// 6: &SingleSymbol{symbol: "BNB/BTC~api.binance.com", granularity: 1000000, transform: MeanAt},

	// 7: &SingleSymbol{symbol: "BNB/ETH~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 8: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "ETH/USD~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 	after:  &SingleSymbol{symbol: "ETH/USD", granularity: 1000000, transform: TimeWeightedAvg(24*time.Hour, ExpDecay)},
	// 	at:     switchTime,
	// },
	// 9: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "LINK/USDT~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 	after:  &SingleSymbol{symbol: "ETH/USD", granularity: 1000000, transform: EODMedian},
	// 	at:     switchTime,
	// },
	10: &TimedSwitch{
		before: &SingleSymbol{symbol: "ETC/ETH~api.binance.com", granularity: 1000000, transform: MeanAt},
		after:  &Ampl{granularity: 1000000},
		at:     switchTime,
	},
	// 11: &SingleSymbol{symbol: "ZEC/ETH~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 12: &SingleSymbol{symbol: "TRX/ETH~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 13: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "XRP/BTC~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 	after:  &SingleSymbol{symbol: "XRP/USD", granularity: 1000000, transform: MedianAt},
	// 	at:     switchTime,
	// },
	// 14: &SingleSymbol{symbol: "XMR/ETH~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 15: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "XLM/BTC~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 	after:  &SingleSymbol{symbol: "ATOM/USD", granularity: 1000000, transform: MedianAt},
	// 	at:     switchTime,
	// },
	// 16: &SingleSymbol{symbol: "LTC/USDT~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 17: &SingleSymbol{symbol: "WAVES/BTC~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 18: &SingleSymbol{symbol: "REP/BTC~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 19: &SingleSymbol{symbol: "TUSD/ETH~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 20: &SingleSymbol{symbol: "EOS/USDT~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 21: &SingleSymbol{symbol: "IOTA/USDT~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 22: &SingleSymbol{symbol: "ETC/USDT~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 23: &SingleSymbol{symbol: "ETH/PAX~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 24: &SingleSymbol{symbol: "ETH/USDC~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 25: &SingleSymbol{symbol: "USDC/USDT~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 26: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "RCN/BTC~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 	after:  &SingleSymbol{symbol: "DEFISCORE", granularity: 1000000, transform: MeanAt},
	// 	at:     switchTime,
	// },
	// 27: &SingleSymbol{symbol: "LINK/USDC~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 28: &SingleSymbol{symbol: "ZRX/BNB~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 29: &SingleSymbol{symbol: "ZEC/USDC~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 30: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "DASH/BNB~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 	after:  &SingleSymbol{symbol: "XAU/USD", granularity: 1000000, transform: MedianAt},
	// 	at:     switchTime,
	// },
	// 31: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "MATIC/USD~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 	after:  &SingleSymbol{symbol: "MATIC/USD", granularity: 1000000, transform: MedianAt},
	// 	at:     switchTime,
	// },
	// 32: &SingleSymbol{symbol: "BAT/USDC~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 33: &SingleSymbol{symbol: "ALGO/USDT~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 34: &SingleSymbol{symbol: "ZRX/USDT~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 35: &SingleSymbol{symbol: "COS/USDT~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 36: &SingleSymbol{symbol: "BCH/USD~api.kraken.com", granularity: 1000000, transform: MeanAt},
	// 37: &SingleSymbol{symbol: "REP/USD~api.coingecko.com", granularity: 1000000, transform: MeanAt},
	// 38: &SingleSymbol{symbol: "GNO/USD~api.kraken.com", granularity: 1000000, transform: MeanAt},
	// 39: &SingleSymbol{symbol: "DAI/USD~api.kraken.com", granularity: 1000000, transform: MeanAt},
	// 40: &SingleSymbol{symbol: "STEEM/BTC~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 41: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "LINK/USDT~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 	after:  &SingleSymbol{symbol: "USPCE", granularity: 1000, transform: MedianAt},
	// 	at:     switchTime,
	// },
	// 42: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "WAN/BTC~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 	after:  &SingleSymbol{symbol: "BTC/USD", granularity: 1000000, transform: EODMedian},
	// 	at:     switchTime,
	// },
	// 43: &SingleSymbol{symbol: "GNT/ETH~api.binance.com", granularity: 1000000, transform: MeanAt},
	// 44: &SingleSymbol{symbol: "BTC/USD~api-pub.bitfinex.com", granularity: 1000000, transform: MeanAt},
	// 45: &SingleSymbol{symbol: "BTC/USD~api.coingecko.com", granularity: 1000000, transform: MeanAt},
	// 46: &SingleSymbol{symbol: "ETH/USD~api.coingecko.com", granularity: 1000000, transform: MeanAt},
	// 47: &SingleSymbol{symbol: "LTC/USD~api.coingecko.com", granularity: 1000000, transform: MeanAt},
	// 48: &SingleSymbol{symbol: "MAKER/USD~api.coingecko.com", granularity: 1000000, transform: MeanAt},
	// 49: &SingleSymbol{symbol: "EOS/USD~api.coingecko.com", granularity: 1000000, transform: MeanAt},
	// 50: &SingleSymbol{symbol: "TRB/USD~api.coingecko.com", granularity: 1000000, transform: MeanAt},
}

//these weight functions map values of x between 0 (brand new) and 1 (old) to weights between 0 and 1
//also returns the integral of the weight over the range [0,1]
//weights the oldest data (1) as being 1/3 as important (1/e)
func ExpDecay(x float64) (float64, float64) {
	return 1 / math.Exp(x), 0.63212
}

//weights the oldest data at 0
func LinearDecay(x float64) (float64, float64) {
	return 1 - x, 0.5
}

//weights all data in the time interval evenly
func NoDecay(x float64) (float64, float64) {
	return 1, 1
}

func TimeWeightedAvg(interval time.Duration, weightFn func(float64) (float64, float64)) IndexProcessor {
	return func(apis []*IndexTracker, at time.Time) (apiOracle.PriceInfo, float64) {
		cfg := config.GetConfig()
		sum := 0.0
		weightSum := 0.0
		numVals := 0
		minTime := at
		maxVolume := 0.0
		for _, api := range apis {
			values := apiOracle.GetRequestValuesForTime(api.Identifier, at, interval)
			for _, v := range values {
				normDelta := at.Sub(v.Created).Seconds() / interval.Seconds()
				weight, _ := weightFn(normDelta)
				sum += v.Price * weight
				weightSum += weight
				numVals += 1
				if minTime.Sub(v.Created).Seconds() > 0 {
					minTime = v.Created
				}
				if v.Volume > maxVolume {
					maxVolume = v.Volume
				}
			}
		}
		// number of APIs * rate * interval
		maxWeight := float64(len(apis)) * (1 / cfg.TrackerSleepCycle.Duration.Seconds()) * interval.Seconds()
		fmt.Println("Number of Values in Time Weighted : ", numVals)
		fmt.Println("Time of last value : ", minTime)
		//average weight is the integral of the weight fn over [0,1]
		_, avgWeight := weightFn(0)
		targetWeight := maxWeight * avgWeight

		var result apiOracle.PriceInfo
		result.Price = sum / weightSum

		//use the highest volume seen over all values. works well when the time averaging window is equal to the interval of volume reporting
		// ie, 24 hour average on an api that returns 24hr volume
		result.Volume = maxVolume
		return result, math.Min(weightSum/targetWeight, 1.0)
	}
}

func VolumeWeightedAPIs(processor IndexProcessor) IndexProcessor {
	return func(apis []*IndexTracker, at time.Time) (apiOracle.PriceInfo, float64) {
		var results []apiOracle.PriceInfo
		totalConfidence := 0.0
		for _,api := range apis {
			value, confidence := processor([]*IndexTracker{api}, at)
			fmt.Printf("AMPL api %s: %f (%f)\n", api.Name, value, confidence)
			if confidence > 0.3 {
				results = append(results, value)
				totalConfidence += confidence
			}
		}
		return VolumeWeightedAvg(results), totalConfidence/float64(len(results))
	}
}

func getLatest(apis []*IndexTracker, at time.Time) ([]apiOracle.PriceInfo, float64) {
	var values []apiOracle.PriceInfo
	totalConf := 0.0
	for _, api := range apis {
		b, _ := apiOracle.GetNearestTwoRequestValue(api.Identifier, at)
		if b != nil {
			//penalize values more than 5 minutes old
			totalConf += math.Min(5/at.Sub(b.Created).Minutes(), 1.0)
			values = append(values, b.PriceInfo)
		}
	}
	return values, totalConf / float64(len(apis))
}

func MedianAt(apis []*IndexTracker, at time.Time) (apiOracle.PriceInfo, float64) {
	values, confidence := getLatest(apis, at)
	if confidence == 0 {
		return apiOracle.PriceInfo{}, 0
	}
	return Median(values), confidence
}

func MedianAtEOD(apis []*IndexTracker, at time.Time) (apiOracle.PriceInfo, float64) {
	now := time.Now().UTC()
	d := 24 * time.Hour
	eod := now.Truncate(d)
	return MedianAt(apis, eod)
	//interval := 2 * time.Minute
	//var s []*apiOracle.PriceStamp
	//for _, api := range apis {
	//	values := apiOracle.GetRequestValuesForTime(api.Identifier, eod, interval)
	//	for _, v := range values {
	//		s = append(s, v)
	//	}
	//}
	//if len(s) > 0 {
	//	sort.Slice(s, func(i, j int) bool {
	//		return s[i].Price < s[j].Price
	//	})
	//	return s[len(s)/2].Price, float64(len(s) / len(apis))
	//}
	//return 0, 0
}

func Median(values []apiOracle.PriceInfo) apiOracle.PriceInfo {
	var result apiOracle.PriceInfo
	sort.Slice(values, func(i, j int) bool {
		return values[i].Price < values[j].Price
	})
	fmt.Println("median of ", len(values))
	result.Price = values[len(values)/2].Price
	for _, val := range values {
		result.Volume += val.Volume
	}
	return result
}

func MeanAt(apis []*IndexTracker, at time.Time) (apiOracle.PriceInfo, float64) {
	values, confidence := getLatest(apis, at)
	if confidence == 0 {
		return apiOracle.PriceInfo{}, 0
	}
	return Mean(values), confidence
}

func Mean(vals []apiOracle.PriceInfo) apiOracle.PriceInfo {
	var result apiOracle.PriceInfo
	priceSum := 0.0
	volSum := 0.0
	for _, val := range vals {
		priceSum += val.Price
		volSum += val.Volume
	}
	result.Price = priceSum / float64(len(vals))
	result.Volume = volSum
	return result
}


func VolumeWeightedAvg(vals []apiOracle.PriceInfo) apiOracle.PriceInfo {
	priceSum := 0.0
	volSum := 0.0
	for _, val := range vals {
		priceSum += val.Price * val.Volume
		volSum += val.Volume
	}
	if volSum > 0 {
		return apiOracle.PriceInfo{Price: priceSum / volSum, Volume: volSum}
	}
	//if there was no volume data, just do a normal average instead
	priceSum = 0
	for _, val := range vals {
		priceSum += val.Price
	}
	return apiOracle.PriceInfo{Price: priceSum / float64(len(vals)), Volume: 0}
}

