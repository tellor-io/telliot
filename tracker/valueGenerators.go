package tracker

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tellor-io/TellorMiner/apiOracle"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/db"
	"math"
	"math/big"
	"sort"
	"time"
)

//a function to consolidate the recorded API values to a single value
type IndexProcessor func([]*IndexTracker)(float64, bool)

type ValueGenerator interface {
	Require() map[string]IndexProcessor
	Update(map[string]float64) float64
	Multiplier() float64
}


var sensitivity map[string][]int

var requirements map[int]map[string]IndexProcessor


func InitPSRs() {
	sensitivity = make(map[string][]int)
	requirements = make(map[int]map[string]IndexProcessor)
	for requestID, handler := range PSRs {
		reqs := handler.Require()
		for symbol := range reqs {
			sensitivity[symbol] = append(sensitivity[symbol], requestID)
		}
		requirements[requestID] = reqs
	}
}

func UpdatePSRs(ctx context.Context, updatedSymbols []string) error {
	//generate a set of all affected PSRs
	toUpdate := make(map[int]bool)
	for _,symbol := range updatedSymbols {
		for _,requestID := range sensitivity[symbol] {
			toUpdate[requestID] = true
		}
	}

	//update all affected PSRs
	for requestID := range toUpdate {
		//compute the requirements
		reqs := requirements[requestID]
		allGood := true
		values := make(map[string]float64)
		for symbol, fn := range reqs {
			val, ok := fn(indexes[symbol])
			if !ok {
				allGood = false
				break
			}
			values[symbol] = val
		}
		//we need all the requirements to update a PSR
		if !allGood {
			continue
		}
		amt := PSRs[requestID].Update(values)
		amt *= PSRs[requestID].Multiplier()

		//convert it directly from a float to a bigInt so that we don't risk overflowing a uint64
		bigVal := new(big.Float)
		bigVal.SetFloat64(amt)
		bigInt := new(big.Int)
		bigVal.Int(bigInt)

		//encode it and store to DB
		DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)
		enc := hexutil.EncodeBig(bigInt)
		err := DB.Put(fmt.Sprintf("%s%d", db.QueriedValuePrefix, requestID), []byte(enc))
		if err != nil {
			return err
		}
	}
	return nil
}


