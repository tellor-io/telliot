package ops

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	tellor "github.com/tellor-io/TellorMiner/contracts"
	"github.com/tellor-io/TellorMiner/rpc"
	"math/big"
	"time"
)

type Submission struct {
	Val *big.Int
	MinerAddr common.Address
}

type MinedBlock struct {
	BlockNum *big.Int
	Result *big.Int
	Sources [5]Submission
}

func allActivityLaterThan(instance *tellor.TellorMaster,reqId int, earliest time.Time) ([]*MinedBlock, error) {
	reqIdBig := big.NewInt(int64(reqId))
	newValCount, err := instance.GetNewValueCountbyRequestId(nil, reqIdBig)
	if err != nil {
		return nil, fmt.Errorf("getValueCount failed: %s", err.Error())
	}
	currIndex := newValCount.Int64() - 1
	if currIndex < 0 {
		fmt.Printf("%d has never been mined\n", reqId)
		return nil, nil
	}
	result := make([]*MinedBlock, 0)
	for {
		index := big.NewInt(currIndex)
		timeRaw, err := instance.GetTimestampbyRequestIDandIndex(nil, reqIdBig, index)
		if err != nil {
			return nil, fmt.Errorf("getTimestamp failed: %s", err.Error())
		}

		goTime := time.Unix(timeRaw.Int64(), 0)
		if goTime.Before(earliest) {
			return result, nil
		}
		block := &MinedBlock{}
		block.Result, err = instance.RetrieveData(nil, reqIdBig, timeRaw)
		if err != nil {
			return nil, fmt.Errorf("retrieveData failed: %s", err.Error())
		}
		minerVals, err := instance.GetSubmissionsByTimestamp(nil, reqIdBig, timeRaw)
		if err != nil {
			return nil, fmt.Errorf("getSubmissions failed: %s", err.Error())
		}
		minerAddrs, err := instance.GetMinersByRequestIdAndTimestamp(nil, reqIdBig, timeRaw)
		if err != nil {
			return nil, fmt.Errorf("getMiners failed: %s", err.Error())
		}
		for i,minerVal := range minerVals {
			block.Sources[i] = Submission{
				Val:       minerVal,
				MinerAddr: minerAddrs[i],
			}
		}
		block.BlockNum, err = instance.GetMinedBlockNum(nil, reqIdBig, timeRaw)
		if err != nil {
			return nil, fmt.Errorf("getMinedBlockNum failed: %s", err.Error())
		}
		result = append(result, block)
		currIndex--
	}
}

func Activity(ctx context.Context) error {
	cfg, err := config.GetConfig()
	if err != nil {
		depositLog.Error("Problem getting config: %+v", err)
		return err
	}

	client := ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)

	//privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	//if err != nil {
	//	depositLog.Error("Problem getting private key: %+v", err)
	//	return err
	//}
	contractAddress := common.HexToAddress(cfg.ContractAddress)

	instance, err := tellor.NewTellorMaster(contractAddress, client)
	if err != nil {
		depositLog.Error("Problem creating contract: %+v", err)
		//log.Fatal(err)
		return err
	}

	//earliest time we can file a dispute for
	earliest := time.Now().Add(-24 * time.Hour)
	count := 0
	for i := 1; i <= 50; i++ {
		blocks, err := allActivityLaterThan(instance, i, earliest)
		if err != nil {
			return fmt.Errorf("request id %d failed: %s", i, err.Error())
		}
		count += len(blocks)
		if len(blocks) > 0 {
			fmt.Printf("id %d had %d blocks\n", i, len(blocks))
			n := 10
			if n > len(blocks) {
				n = len(blocks)
			}
			fmt.Printf("newest 10 block IDs:")
			for i := 0; i < n; i++ {
				fmt.Printf("\t%s\n", blocks[len(blocks) - (i + 1)].BlockNum.Text(10))
			}
		}
	}
	fmt.Printf("found %d blocks in the past 24 hours\n", count)

	//dat := crypto.Keccak256([]byte("stakerCount"))
	//var dat32 [32]byte
	//copy(dat32[:], dat)
	//stakerCount, err := instance.GetUintVar(nil, dat32)
	//if err != nil {
	//	return err
	//}
	//fmt.Printf("tellor currently has %s stakers\n", stakerCount.Text(10))

	//newValCount, err := instance.GetNewValueCountbyRequestId(nil, reqID)
	//if err != nil {
	//	return err
	//}
	//fmt.Printf("newValueCount = %d\n", newValCount.Uint64())
	//
	//granularity := 1000000.0
	//for i := newValCount.Uint64()-1; i > newValCount.Uint64()-5; i-- {
	//	index := big.NewInt(int64(i))
	//	timeRaw, err := instance.GetTimestampbyRequestIDandIndex(nil, reqID, index)
	//	if err != nil {
	//		return err
	//	}
	//	val, err := instance.RetrieveData(nil, reqID, timeRaw)
	//	if err != nil {
	//		return err
	//	}
	//	minerVals, err := instance.GetSubmissionsByTimestamp(nil, reqID, timeRaw)
	//	if err != nil {
	//		return err
	//	}
	//	minerAddrs, err := instance.GetMinersByRequestIdAndTimestamp(nil, reqID, timeRaw)
	//	if err != nil {
	//		return err
	//	}
	//	for i,minerVal := range minerVals {
	//		fmt.Printf("[%s %f], ", minerAddrs[i].String()[:8],float64(minerVal.Uint64())/granularity)
	//	}
	//	goTime := time.Unix(timeRaw.Int64(), 0)
	//	fmt.Printf("index %d - %d @ %s\n", i, val.Uint64(), goTime)
	//}

	return nil
	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	depositLog.Error("Problem extract public key")
	//	return fmt.Errorf("error casting public key to ECDSA")
	//}

}

