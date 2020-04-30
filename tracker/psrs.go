package tracker

import (
	"math"
	"sort"
	"time"

	"github.com/tellor-io/TellorMiner/apiOracle"
	"github.com/tellor-io/TellorMiner/config"
)

var switchTime, _ = time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")

var PSRs = map[int]ValueGenerator{
	// 1: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "ETH/USD~api.pro.coinbase.com", granularity: 1000, transform: CurrentMean},
	// 	after:  &SingleSymbol{symbol: "ETH/USD", granularity: 1000000, transform: CurrentMedian},
	// 	at:     switchTime,
	// },
	// 2: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "BTC/USD~api.binance.com", granularity: 1000, transform: CurrentMean},
	// 	after:  &SingleSymbol{symbol: "BTC/USD", granularity: 1000000, transform: CurrentMedian},
	// 	at:     switchTime,
	// },
	// 3: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "BNB/USD~dex.binance.org", granularity: 1000, transform: CurrentMean},
	// 	after:  &SingleSymbol{symbol: "BNB/USD", granularity: 1000000, transform: CurrentMedian},
	// 	at:     switchTime,
	// },
	// 4: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "BTC/USD", granularity: 100000, transform: CurrentMedian},
	// 	after:  &SingleSymbol{symbol: "BTC/USD", granularity: 1000000, transform: TimeWeightedAvg(24*time.Hour, ExpDecay)},
	// 	at:     switchTime,
	// },
	// 5: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "ETH/BTC~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 	after:  &SingleSymbol{symbol: "ETH/BTC", granularity: 1000000, transform: CurrentMedian},
	// 	at:     switchTime,
	// },
	// 6: &SingleSymbol{symbol: "BNB/BTC~api.binance.com", granularity: 1000000, transform: CurrentMean},

	// 7: &SingleSymbol{symbol: "BNB/ETH~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 8: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "ETH/USDT~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 	after:  &SingleSymbol{symbol: "ETH/USD", granularity: 1000000, transform: TimeWeightedAvg(24*time.Hour, ExpDecay)},
	// 	at:     switchTime,
	// },
	// 9: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "LINK/USDT~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 	after:  &SingleSymbol{symbol: "ETH/USD", granularity: 1000000, transform: EOD},
	// 	at:     switchTime,
	// },
	// 10: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "ETC/ETH~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 	after:  &ChainedPrice{chain: []string{"AMPL/USD", "AMPL/BTC", "BTC/USD"}, granularity: 1000000, transform: TimeWeightedAvg(10*time.Minute, NoDecay)},
	// 	at:     switchTime,
	// },
	// 11: &SingleSymbol{symbol: "ZEC/ETH~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 12: &SingleSymbol{symbol: "TRX/ETH~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 13: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "XRP/BTC~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 	after:  &SingleSymbol{symbol: "XRP/USD", granularity: 1000000, transform: CurrentMedian},
	// 	at:     switchTime,
	// },
	// 14: &SingleSymbol{symbol: "XMR/ETH~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 15: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "XLM/BTC~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 	after:  &SingleSymbol{symbol: "ATOM/USD", granularity: 1000000, transform: CurrentMedian},
	// 	at:     switchTime,
	// },
	// 16: &SingleSymbol{symbol: "LTC/USDT~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 17: &SingleSymbol{symbol: "WAVES/BTC~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 18: &SingleSymbol{symbol: "REP/BTC~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 19: &SingleSymbol{symbol: "TUSD/ETH~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 20: &SingleSymbol{symbol: "EOS/USDT~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 21: &SingleSymbol{symbol: "IOTA/USDT~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 22: &SingleSymbol{symbol: "ETC/USDT~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 23: &SingleSymbol{symbol: "ETH/PAX~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 24: &SingleSymbol{symbol: "ETH/USDC~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 25: &SingleSymbol{symbol: "USDC/USDT~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 26: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "RCN/BTC~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 	after:  &SingleSymbol{symbol: "DEFISCORE", granularity: 1000000, transform: CurrentMean},
	// 	at:     switchTime,
	// },
	// 27: &SingleSymbol{symbol: "LINK/USDC~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 28: &SingleSymbol{symbol: "ZRX/BNB~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 29: &SingleSymbol{symbol: "ZEC/USDC~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 30: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "DASH/BNB~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 	after:  &SingleSymbol{symbol: "XAU/USD", granularity: 1000000, transform: CurrentMedian},
	// 	at:     switchTime,
	// },
	// 31: &TimedSwitch{
	// 	before: &SingleSymbol{symbol: "MATIC/USDT~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 	after:  &SingleSymbol{symbol: "MATIC/USD", granularity: 1000000, transform: CurrentMedian},
	// 	at:     switchTime,
	// },
	// 32: &SingleSymbol{symbol: "BAT/USDC~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 33: &SingleSymbol{symbol: "ALGO/USDT~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 34: &SingleSymbol{symbol: "ZRX/USDT~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 35: &SingleSymbol{symbol: "COS/USDT~api.binance.com", granularity: 1000000, transform: CurrentMean},
	// 36: &SingleSymbol{symbol: "BCH/USD~api.kraken.com", granularity: 1000000, transform: CurrentMean},
	// 37: &SingleSymbol{symbol: "REP/USD~api.coingecko.com", granularity: 1000000, transform: CurrentMean},
	// 38: &SingleSymbol{symbol: "GNO/USD~api.kraken.com", granularity: 1000000, transform: CurrentMean},
	// 39: &SingleSymbol{symbol: "DAI/USD~api.kraken.com", granularity: 1000000, transform: CurrentMean},
	// 40: &SingleSymbol{symbol: "STEEM/BTC~api.binance.com", granularity: 1000000, transform: CurrentMean},
	41: &TimedSwitch{
		before: &SingleSymbol{symbol: "LINK/USDT~api.binance.com", granularity: 1000000, transform: CurrentMean},
		after:  &SingleSymbol{symbol: "USPCE", granularity: 1000, transform: CurrentMedian},
		at:     switchTime,
	},
	42: &TimedSwitch{
		before: &SingleSymbol{symbol: "WAN/BTC~api.binance.com", granularity: 1000000, transform: CurrentMean},
		after:  &SingleSymbol{symbol: "BTC/USD", granularity: 1000000, transform: EOD},
		at:     switchTime,
	},
	43: &SingleSymbol{symbol: "GNT/ETH~api.binance.com", granularity: 1000000, transform: CurrentMean},
	44: &SingleSymbol{symbol: "BTC/USD~api-pub.bitfinex.com", granularity: 1000000, transform: CurrentMean},
	45: &SingleSymbol{symbol: "BTC/USD~api.coingecko.com", granularity: 1000000, transform: CurrentMean},
	46: &SingleSymbol{symbol: "ETH/USD~api.coingecko.com", granularity: 1000000, transform: CurrentMean},
	47: &SingleSymbol{symbol: "LTC/USD~api.coingecko.com", granularity: 1000000, transform: CurrentMean},
	48: &SingleSymbol{symbol: "MAKER/USD~api.coingecko.com", granularity: 1000000, transform: CurrentMean},
	49: &SingleSymbol{symbol: "EOS/USD~api.coingecko.com", granularity: 1000000, transform: CurrentMean},
	50: &SingleSymbol{symbol: "TRB/USD~api.coingecko.com", granularity: 1000000, transform: CurrentMean},
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
	return func(apis []*IndexTracker, at time.Time) (float64, float64) {
		cfg := config.GetConfig()
		sum := 0.0
		weightSum := 0.0
		for _, api := range apis {
			values := apiOracle.GetRequestValuesForTime(api.Identifier, at, interval)
			for _, v := range values {
				normDelta := at.Sub(v.Created).Seconds() / interval.Seconds()
				weight, _ := weightFn(normDelta)
				sum += v.Price * weight
				weightSum += weight
			}
		}
		// number of APIs * rate * interval
		maxWeight := float64(len(apis)) * (1 / cfg.TrackerSleepCycle.Duration.Seconds()) * interval.Seconds()

		//average weight is the integral of the weight fn over [0,1]
		_, avgWeight := weightFn(0)
		targetWeight := maxWeight * avgWeight
		return sum / weightSum, math.Min(weightSum/targetWeight, 1.0)
	}
}

func getLatest(apis []*IndexTracker, at time.Time) ([]*apiOracle.PriceStamp, float64) {
	var values []*apiOracle.PriceStamp
	totalConf := 0.0
	for _, api := range apis {
		b, _ := apiOracle.GetNearestTwoRequestValue(api.Identifier, at)
		if b != nil {
			//penalize values more than 5 minutes old
			totalConf += math.Min(5/at.Sub(b.Created).Minutes(), 1.0)
			values = append(values, b)
		}
	}
	return values, totalConf / float64(len(apis))
}

func CurrentMedian(apis []*IndexTracker, at time.Time) (float64, float64) {
	values, confidence := getLatest(apis, at)
	if confidence == 0 {
		return 0, 0
	}
	sort.Slice(values, func(i, j int) bool {
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
	for _, val := range values {
		sum += val.Price
	}
	return sum / float64(len(values)), confidence
}

func EOD(apis []*IndexTracker, at time.Time) (float64, float64) {
	if at.Hour() == 0 && at.Minute() == 0 {
		return CurrentMedian(apis, at)
	} else {
		return 0, 0
	}
}
