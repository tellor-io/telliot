package tracker

import (
	"encoding/json"
	"fmt"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/util"
	"math"
	"math/big"
	"sort"
	"time"
)

type ValueProcessor interface {
	Transform(*PrespecifiedRequest, [][]byte) (uint64, error)
	Value(*PrespecifiedRequest) *big.Int
}

type ValueP struct {
	ValueProcessor
}

func (f *ValueP)UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	switch str {
	case "median": f.ValueProcessor = MedianProc{}
	case "value": f.ValueProcessor = DefaultProcessor{}
	case "dayAvg": f.ValueProcessor = TimeAverage{}
	case "": f.ValueProcessor = DefaultProcessor{}
	default: return fmt.Errorf("unrecognized transformation in PSR: %s", str)
	}
	return nil
}

//defaults for all PSRs
//override these functions to change behaviour
type DefaultProcessor struct {
}

func (p DefaultProcessor)Transform(r *PrespecifiedRequest, payloads [][]byte) (uint64, error) {
	return mean(parsePayloads(r, payloads)), nil
}

func (p DefaultProcessor)Value(r *PrespecifiedRequest) *big.Int {
	ti := GetLatestRequestValue(r.RequestID)
	bigVal := new(big.Int)
	bigVal.SetUint64(ti.Val)
	return bigVal
}



type MedianProc struct {
	DefaultProcessor
}
func (m MedianProc)Transform(r *PrespecifiedRequest, payloads [][]byte) (uint64, error) {
	return median(parsePayloads(r, payloads)), nil
}


//does a time average over 1 day
type TimeAverage struct {
	DefaultProcessor
}
func (t TimeAverage)Value(r *PrespecifiedRequest) *big.Int {
	cfg := config.GetConfig()

	timeInterval := 24 * time.Hour

	//get all the data we have saved locally for the past day
	now := time.Now()
	vals := GetRequestValuesForTime(r.RequestID, now, timeInterval)

	max := timeInterval/cfg.TrackerSleepCycle.Duration

	//require at least 60% of the values from the past day
	ratio := float64(len(vals))/float64(max)
	if ratio < 0.6 {
		estimate := time.Duration(0.6 * float64(timeInterval))
		psrLog.Info("Insufficient data for request ID %d, expected in %s", r.RequestID, estimate.String())
		return nil
	}

	uintVals := make([]uint64, len(vals))
	for i := range vals {
		uintVals[i] = vals[i].Val
	}
	result := new(big.Int)
	result.SetUint64(mean(uintVals))
	return result
}



//common shared functions among different value processors

func parsePayloads(r *PrespecifiedRequest, payloads [][]byte) []uint64 {
	vals := make([]uint64, 0, len(payloads))
	for i, pl := range payloads {
		if pl == nil {
			continue
		}
		v, err := util.ParsePayload(pl, r.Granularity, r.argGroups[i])
		if err != nil {
			continue
		}
		vals = append(vals, uint64(v))
	}
	if len(vals) == 0 {
		psrLog.Error("no sucessful api hits, no value stored for id %d", r.RequestID)
	}
	return vals
}

//an alternative weighting scheme.
//
// new values     1.00
// 6 hours old    0.50
// 24 hours old   0.05
func expTimeWeightedMean(vals []*TimedInt) uint64 {
	now := time.Now()
	sum := 0.0
	for _,v := range vals {
		delta := now.Sub(v.Created).Seconds()
		sum += float64(v.Val) * 1/(math.Exp(delta/(86400/3)))
	}
	return uint64(sum/float64(len(vals)))
}

func mean(vals []uint64) uint64 {
	//compute the mean
	sum := 0.0
	for _,v := range vals {
		sum += float64(v)
	}
	avg:= sum / float64(len(vals))
	return uint64(avg)
}

func median(vals []uint64) uint64 {
	sort.Slice(vals, func (i, j int) bool {
		return vals[i] < vals[j]
	})
	return vals[len(vals)/2]
}