package tracker

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/contracts1"
	"github.com/tellor-io/TellorMiner/rpc"
	"github.com/tellor-io/TellorMiner/util"
	"io/ioutil"
	"math"
	"math/big"
	"strings"
	"time"
)

var disputeLogger = util.NewLogger("tracker", "disputeChecker")

type disputeChecker struct {
	lastCheckedBlock uint64
}


func (c *disputeChecker) String() string {
	return "DisputeChecker"
}


type ValueCheckResult struct {
	High, Low float64
	WithinRange bool
	Datapoints []*TimedFloat
}

func CheckValueAtTime(reqId uint64, val *big.Int, at time.Time) *ValueCheckResult {
	cfg := config.GetConfig()

	datapoints := GetRequestValuesForTime(reqId, at, cfg.DisputeTimeDelta.Duration)

	if len(datapoints) == 0 {
		//no data for requestID
		return nil
	}

	min := math.MaxFloat64
	max := 0.0

	for _,dp := range datapoints {
		val := float64(dp.Val)
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}

	min *= 1 - cfg.DisputeThreshold
	max *= 1 + cfg.DisputeThreshold

	floatVal := float64(val.Uint64())


	withinRange := (floatVal > min) && (floatVal < max)

	return &ValueCheckResult{
		Low:         min,
		High:        max,
		WithinRange: withinRange,
		Datapoints:  datapoints,
	}
}

func (c *disputeChecker) Exec(ctx context.Context) error {

	client := ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)


	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to get latest eth block header: %v", err)
	}
	if c.lastCheckedBlock == 0 {
		c.lastCheckedBlock = header.Number.Uint64()
	}

	toCheck := header.Number.Uint64()

	const blockDelay = 100
	if toCheck - c.lastCheckedBlock < blockDelay {
		return nil
	}

	tokenAbi, err := abi.JSON(strings.NewReader(contracts1.TellorLibraryABI))
	if err != nil {
		return fmt.Errorf("failed to parse abi: %v", err)
	}
	contractAddress := ctx.Value(tellorCommon.ContractAddress).(common.Address)

	//just use nil for most of the variables, only using this object to call UnpackLog which only uses the abi
	bar := bind.NewBoundContract(contractAddress, tokenAbi, nil, nil, nil)

	checkUntil := toCheck - blockDelay
	nonceSubmitID := tokenAbi.Events["NonceSubmitted"].ID()
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(c.lastCheckedBlock)),
		ToBlock:   big.NewInt(int64(checkUntil)),
		Addresses: []common.Address{contractAddress},
		Topics: [][]common.Hash{{nonceSubmitID}},
	}
	logs, err := client.FilterLogs(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to filter eth logs: %v", err)
	}
	blockTimes := make(map[uint64]time.Time)
	for _,l := range logs {
		nonceSubmit := contracts1.TellorLibraryNonceSubmitted{}
		err := bar.UnpackLog(&nonceSubmit,"NonceSubmitted", l)
		if err != nil {
			return fmt.Errorf("failed to unpack into object: %v", err)
		}
		blockTime, ok := blockTimes[l.BlockNumber]
		if !ok {
			header, err := client.HeaderByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
			if err != nil {
				return fmt.Errorf("failed to get nonce block header: %v", err)
			}
			blockTime = time.Unix(int64(header.Time), 0)
			blockTimes[l.BlockNumber] = blockTime
		}
		reqID := nonceSubmit.RequestId.Uint64()
		result := CheckValueAtTime(reqID, nonceSubmit.Value, blockTime)
		if result == nil {
			disputeLogger.Warn("no value data for reqID %d at %s", reqID, blockTime)
			continue
		}
		if !result.WithinRange {
			s := fmt.Sprintf("suspected incorrect value for requestID %d at %s:\n", reqID, blockTime)
			s += fmt.Sprintf("nearest values:\n")
			for _,pt := range result.Datapoints {
				s += fmt.Sprintf("\t%d, ", pt.Val)
				delta := blockTime.Sub(pt.Created)
				if delta > 0 {
					s += fmt.Sprintf("%s before\n", delta.String())
				} else {
					s += fmt.Sprintf("%s after\n", (-delta).String())
				}
			}
			s += fmt.Sprintf("value submitted by miner with address %s", nonceSubmit.Miner)
			disputeLogger.Error(s)
			filename := fmt.Sprintf("possible-dispute-%s.txt", blockTime)
			err := ioutil.WriteFile(filename, []byte(s), 0655)
			if err != nil {
				disputeLogger.Error("failed to save dispute data to %s: %v", filename, err)
			}
		} else {
			disputeLogger.Info("value of %s for requestid %d at %s appears to be within expected range", nonceSubmit.Value, reqID, blockTime.String())
		}

	}
	c.lastCheckedBlock = checkUntil
	return nil
}