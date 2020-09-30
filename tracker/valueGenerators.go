package tracker

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tellor-io/TellorMiner/apiOracle"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/db"
)

//a function to consolidate the recorded API values to a single value
type IndexProcessor func([]*IndexTracker, time.Time) (apiOracle.PriceInfo, float64)

type ValueGenerator interface {
	//PSRs report what they require to produce a value with this
	Require(time.Time) map[string]IndexProcessor

	//return the best estimate of a value at a given time, and the confidence
	// if confidence == 0, the value has no meaning
	ValueAt(map[string]apiOracle.PriceInfo, time.Time) float64
}

func InitPSRs() error {
	//check that we have all the symbols asked for
	now := clck.Now()
	for requestID, handler := range PSRs {
		reqs := handler.Require(now)
		for symbol := range reqs {
			_, ok := indexes[symbol]
			if !ok {
				return fmt.Errorf("PSR %d requires non-existent symbol %s", requestID, symbol)
			}
		}
	}
	return nil
}

func PSRValueForTime(requestID int, at time.Time) (float64, float64) {
	//get the requirements
	reqs := PSRs[requestID].Require(at)
	values := make(map[string]apiOracle.PriceInfo)
	minConfidence := math.MaxFloat64
	for symbol, fn := range reqs {

		val, confidence := fn(indexes[symbol], at)
		if confidence == 0 {
			return 0, 0
		}
		if confidence < minConfidence {
			minConfidence = confidence
		}
		values[symbol] = val
		//fmt.Println("Value Updated", symbol, " : ", requestID, ": ", val)
	}
	//fmt.Println("values", values)
	return PSRs[requestID].ValueAt(values, at), minConfidence
}

func UpdatePSRs(ctx context.Context, updatedSymbols []string) error {
	now := clck.Now()
	//generate a set of all affected PSRs
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

	//update all affected PSRs
	for _, requestID := range toUpdate {
		amt, conf := PSRValueForTime(requestID, now)
		// if requestID == 10{
			// 	fmt.Println("ID : ",requestID," Confidence: ",conf," || Value: ",amt)
			// }
			cfg := config.GetConfig()
			if conf < cfg.MinConfidence || math.IsNaN(amt) {
				//fmt.Println("ID : ",requestID," Confidence too low: ",conf," || Min required: ",cfg.MinConfidence)
				//confidence in this signal is too low to use
				continue
			}
			//fmt.Print("\ntester ", requestID)

		//convert it directly from a float to a bigInt so that we don't risk overflowing a uint64
		bigVal := new(big.Float)
		bigVal.SetFloat64(amt)
		bigInt := new(big.Int)
		bigVal.Int(bigInt)

		//encode it and store to DB
		DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)
		enc := hexutil.EncodeBig(bigInt)
		//fmt.Print("\nrequestId: ", requestID, " ", db.QueriedValuePrefix, requestID)
		err := DB.Put(fmt.Sprintf("%s%d", db.QueriedValuePrefix, requestID), []byte(enc))
		if err != nil {
			return err
		}
	}
	//fmt.Print("exited with requestIds: ", toUpdate)
	return nil
}
