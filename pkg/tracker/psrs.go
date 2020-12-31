// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/tellor-io/telliot/pkg/apiOracle"
	"github.com/tellor-io/telliot/pkg/config"
)

const RequestID_TRB_ETH int = 43

var PSRs = map[int]ValueGenerator{
	1: &SingleSymbol{symbol: "ETH/USD", granularity: 1000000, transform: MedianAt},
	2: &SingleSymbol{symbol: "BTC/USD", granularity: 1000000, transform: MedianAt},
	3: &SingleSymbol{symbol: "BNB/USD", granularity: 1000000, transform: MedianAt},
	4: &SingleSymbol{symbol: "BTC/USD", granularity: 1000000, transform: TimeWeightedAvg(24*time.Hour, ExpDecay)},
	5: &SingleSymbol{symbol: "ETH/BTC", granularity: 1000000, transform: MedianAt},
	6: &SingleSymbol{symbol: "BNB/BTC", granularity: 1000000, transform: MedianAt},
	7: &SingleSymbol{symbol: "BNB/ETH", granularity: 1000000, transform: MedianAt},
	8: &SingleSymbol{symbol: "ETH/USD", granularity: 1000000, transform: TimeWeightedAvg(24*time.Hour, ExpDecay)},
	9: &SingleSymbol{symbol: "ETH/USD", granularity: 1000000, transform: MedianAtEOD},
	// For more details see https://docs.google.com/document/d/1RFCApk1PznMhSRVhiyFl_vBDPA4mP2n1dTmfqjvuTNw/edit
	10: &Ampl{granularity: 1000000},
	11: &SingleSymbol{symbol: "ZEC/ETH", granularity: 1000000, transform: MedianAt},
	12: &SingleSymbol{symbol: "TRX/ETH", granularity: 1000000, transform: MedianAt},
	13: &SingleSymbol{symbol: "XRP/USD", granularity: 1000000, transform: MedianAt},
	14: &SingleSymbol{symbol: "XMR/ETH", granularity: 1000000, transform: MedianAt},
	15: &SingleSymbol{symbol: "ATOM/USD", granularity: 1000000, transform: MedianAt},
	16: &SingleSymbol{symbol: "LTC/USD", granularity: 1000000, transform: MedianAt},
	17: &SingleSymbol{symbol: "WAVES/BTC", granularity: 1000000, transform: MedianAt},
	18: &SingleSymbol{symbol: "REP/BTC", granularity: 1000000, transform: MedianAt},
	19: &SingleSymbol{symbol: "TUSD/ETH", granularity: 1000000, transform: MedianAt},
	20: &SingleSymbol{symbol: "EOS/USD", granularity: 1000000, transform: MedianAt},
	21: &SingleSymbol{symbol: "IOTA/USD", granularity: 1000000, transform: MedianAt},
	22: &SingleSymbol{symbol: "ETC/USD", granularity: 1000000, transform: MedianAt},
	23: &SingleSymbol{symbol: "ETH/PAX", granularity: 1000000, transform: MedianAt},
	24: &SingleSymbol{symbol: "ETH/BTC", granularity: 1000000, transform: TimeWeightedAvg(1*time.Hour, NoDecay)},
	25: &SingleSymbol{symbol: "USDC/USDT", granularity: 1000000, transform: MedianAt},
	26: &SingleSymbol{symbol: "XTZ/USD", granularity: 1000000, transform: MedianAt},
	27: &SingleSymbol{symbol: "LINK/USD", granularity: 1000000, transform: MedianAt},
	28: &SingleSymbol{symbol: "ZRX/BNB", granularity: 1000000, transform: MedianAt},
	29: &SingleSymbol{symbol: "ZEC/USD", granularity: 1000000, transform: MedianAt},
	30: &SingleSymbol{symbol: "XAU/USD", granularity: 1000000, transform: MedianAt},
	31: &SingleSymbol{symbol: "MATIC/USD", granularity: 1000000, transform: MedianAt},
	32: &SingleSymbol{symbol: "BAT/USD", granularity: 1000000, transform: MedianAt},
	33: &SingleSymbol{symbol: "ALGO/USD", granularity: 1000000, transform: MedianAt},
	34: &SingleSymbol{symbol: "ZRX/USD", granularity: 1000000, transform: MedianAt},
	35: &SingleSymbol{symbol: "COS/USD", granularity: 1000000, transform: MedianAt},
	36: &SingleSymbol{symbol: "BCH/USD", granularity: 1000000, transform: MedianAt},
	37: &SingleSymbol{symbol: "REP/USD", granularity: 1000000, transform: MedianAt},
	38: &SingleSymbol{symbol: "GNO/USD", granularity: 1000000, transform: MedianAt},
	39: &SingleSymbol{symbol: "DAI/USD", granularity: 1000000, transform: MedianAt},
	40: &SingleSymbol{symbol: "STEEM/BTC", granularity: 1000000, transform: MedianAt},
	// It is three month average for US PCE (monthly levels): https://www.bea.gov/data/personal-consumption-expenditures-price-index-excluding-food-and-energy
	41:                &SingleSymbol{symbol: "USPCE", granularity: 1000, transform: ManualEntry},
	42:                &SingleSymbol{symbol: "BTC/USD", granularity: 1000000, transform: MedianAtEOD},
	RequestID_TRB_ETH: &SingleSymbol{symbol: "TRB/ETH", granularity: 1000000, transform: MedianAt},
	44:                &SingleSymbol{symbol: "BTC/USD", granularity: 1000000, transform: TimeWeightedAvg(1*time.Hour, NoDecay)},
	45:                &SingleSymbol{symbol: "TRB/USD", granularity: 1000000, transform: MedianAtEOD},
	46:                &SingleSymbol{symbol: "ETH/USD", granularity: 1000000, transform: TimeWeightedAvg(1*time.Hour, NoDecay)},
	47:                &SingleSymbol{symbol: "BSV/USD", granularity: 1000000, transform: MedianAt},
	48:                &SingleSymbol{symbol: "MAKER/USD", granularity: 1000000, transform: MedianAt},
	49:                &SingleSymbol{symbol: "BCH/USD", granularity: 1000000, transform: TimeWeightedAvg(24*time.Hour, NoDecay)},
	50:                &SingleSymbol{symbol: "TRB/USD", granularity: 1000000, transform: MedianAt},
	51:                &SingleSymbol{symbol: "XMR/USD", granularity: 1000000, transform: MedianAt},
	52:                &SingleSymbol{symbol: "XFT/USD", granularity: 1000000, transform: MedianAt},
	53:                &SingleSymbol{symbol: "BTCDOMINANCE", granularity: 1000000, transform: MedianAt},
	54:                &SingleSymbol{symbol: "WAVES/USD", granularity: 1000000, transform: MedianAt},
	55:                &SingleSymbol{symbol: "OGN/USD", granularity: 1000000, transform: MedianAt},
	56:                &SingleSymbol{symbol: "VIXEOD", granularity: 1000000, transform: MedianAt},
	57:                &SingleSymbol{symbol: "DEFITVL", granularity: 1000000, transform: MedianAt},
}

// ExpDecay maps values of x between 0 (brand new) and 1 (old) to weights between 0 and 1
// also returns the integral of the weight over the range [0,1]
// weights the oldest data (1) as being 1/3 as important (1/e).
func ExpDecay(x float64) (float64, float64) {
	return 1 / math.Exp(x), 0.63212
}

// LinearDecay weights the oldest data at 0.
func LinearDecay(x float64) (float64, float64) {
	return 1 - x, 0.5
}

// NoDecay weights all data in the time interval evenly.
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
		// Number of APIs * rate * interval.
		maxWeight := float64(len(apis)) * (1 / cfg.TrackerSleepCycle.Duration.Seconds()) * interval.Seconds()
		// Average weight is the integral of the weight fn over [0,1].
		_, avgWeight := weightFn(0)
		targetWeight := maxWeight * avgWeight

		var result apiOracle.PriceInfo
		result.Price = sum / weightSum

		// Use the highest volume seen over all values.
		// Works well when the time averaging window is equal to the interval of volume reporting
		// ie, 24 hour average on an api that returns 24hr volume
		result.Volume = maxVolume
		// if math.Min(weightSum/targetWeight, 1.0) < .5{
		// 	values,_ := apiOracle.GetNearestTwoRequestValue(apis[0].Identifier, at)
		// 	fmt.Println("not enough data for time series, series starts : ", values.Created)
		// }
		//fmt.Println("Time Weighted: ", result)
		return result, math.Min(weightSum/targetWeight, 1.0)
	}
}

func VolumeWeightedAPIs(processor IndexProcessor) IndexProcessor {
	return func(apis []*IndexTracker, at time.Time) (apiOracle.PriceInfo, float64) {
		var results []apiOracle.PriceInfo
		totalConfidence := 0.0
		for _, api := range apis {
			value, confidence := processor([]*IndexTracker{api}, at)
			if confidence > 0 {
				results = append(results, value)
				totalConfidence += confidence
			}
		}
		return VolumeWeightedAvg(results), totalConfidence / float64(len(results))
	}
}

func getLatest(apis []*IndexTracker, at time.Time) ([]apiOracle.PriceInfo, float64) {
	var values []apiOracle.PriceInfo
	totalConf := 0.0
	for _, api := range apis {
		b, _ := apiOracle.GetNearestTwoRequestValue(api.Identifier, at)
		if b != nil {
			// Penalize values more than 5 minutes old.
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

func ManualEntry(apis []*IndexTracker, at time.Time) (apiOracle.PriceInfo, float64) {
	vals, confidence := getLatest(apis, at)
	if confidence == 0 {
		return apiOracle.PriceInfo{}, 0
	}
	for _, val := range vals {
		if int64(val.Volume) < clck.Now().Unix() {
			fmt.Println("Pulled Timestamp: ", val.Volume)
			fmt.Println("Warning: Manual Data Entry is expired, please update")
			return apiOracle.PriceInfo{}, 0
		}
	}
	return Median(vals), confidence
}

func MaxPSRID() uint64 {
	var maxID int
	for id := range PSRs {
		if id > maxID {
			maxID = id
		}
	}
	return uint64(maxID)
}

func MedianAtEOD(apis []*IndexTracker, at time.Time) (apiOracle.PriceInfo, float64) {
	now := clck.Now().UTC()
	d := 24 * time.Hour
	eod := now.Truncate(d)
	return MedianAt(apis, eod)
}

func Median(values []apiOracle.PriceInfo) apiOracle.PriceInfo {
	var result apiOracle.PriceInfo
	sort.Slice(values, func(i, j int) bool {
		return values[i].Price < values[j].Price
	})
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
