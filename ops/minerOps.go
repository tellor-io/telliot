package ops

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/pow"
	"github.com/tellor-io/TellorMiner/util"
)

//miningCycle holds all details for the current mining challenge and fields needed to submit a result
type miningCycle struct {
	challenge  []byte
	difficulty *big.Int
	nonce      string
	requestID  *big.Int
	value      *big.Int
}

//MinerOps holds items for mining ops
type MinerOps struct {
	exitCh        chan os.Signal
	log           *util.Logger
	Running       bool
	lastChallenge []byte
	miner         *pow.PoWSolver
	Requesting	bool
}

//CreateMinerOps creates a new miner operation ready to start run loop
func CreateMinerOps(ctx context.Context, exitCh chan os.Signal) (*MinerOps, error) {
	miner := pow.CreateMiner()
	return &MinerOps{exitCh: exitCh, log: util.NewLogger("ops", "MinerOps"), Running: false, miner: miner}, nil
}

//Start will start the mining run loop
func (ops *MinerOps) Start(ctx context.Context) {
	ops.Running = true
	ops.log.Info("Starting miner")
	go func() {
		ticker := time.NewTicker(35 * time.Second)
		for {
			select {
			case _ = <-ops.exitCh:
				{
					ops.log.Info("Stopping miner")
					ops.miner.Stop()
					ops.Running = false
					return
				}
			case _ = <-ticker.C:
				{
					if !ops.Requesting{
						fmt.Println("Building new Cycle...")
						cycle, err := ops.buildNextCycle(ctx)
						if err == nil && cycle != nil {
							fmt.Println("Checking Cycle")
							if (ops.lastChallenge == nil || bytes.Compare(ops.lastChallenge,cycle.challenge) != 0)  && !ops.miner.IsMining() {
								ops.lastChallenge = cycle.challenge
								ops.log.Info("Requesting mining cycle with vars: %+v\n", cycle)
								go ops.mine(ctx, cycle)
							}else{
								fmt.Println("Miner is Mining : ",ops.miner.IsMining())
							}
						}else{
							fmt.Println("Error Building Cycle",err)
						}
					}else{
						fmt.Println("Miner is requesting Data")
					}
				}
			}
		}
	}()
}

func (ops *MinerOps) buildNextCycle(ctx context.Context) (*miningCycle, error) {

	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Println("Couldn't get config")
		return nil, err
	}

	DB := ctx.Value(common.DBContextKey).(db.DB)
	currentChallenge, err := DB.Get(db.CurrentChallengeKey)
	if err != nil {
		ops.log.Error("Problem reading challenge in miner run loop: %v\n", err)
		return nil, err
	}
	if ops.lastChallenge != nil && bytes.Compare(currentChallenge, ops.lastChallenge) == 0 {
		fmt.Println("Challeng being grabbed")
		return nil, nil
	}

	diff, err := DB.Get(db.DifficultyKey)
	if err != nil {
		ops.log.Error("Problem reading difficult from DB: %v\n", err)
		return nil, err
	}
	miningStatus, err := DB.Get(db.MiningStatusKey)
	if err != nil {
		ops.log.Error("Problem reading miningStatus from DB: %v\n", err)
		return nil, err
	}
	if bytes.Compare(miningStatus, []byte{1}) == 0 {
		fmt.Println("Already Mined")
		return nil, nil
	}
	difficulty, err := hexutil.DecodeBig(string(diff))
	if err != nil {
		ops.log.Error("Problem decoding difficulty: %v\n", err)
		return nil, err
	}

	requestID, err := DB.Get(db.RequestIdKey)
	if err != nil {
		ops.log.Error("Problem reading request id from DB: %v\n", err)
		return nil, err
	}
	asInt, err := hexutil.DecodeBig(string(requestID))
	if err != nil {
		ops.log.Error("Problem decoding request id as big int: %v\n", err)
		return nil, err
	}
	if asInt.Cmp(big.NewInt(0)) == 0 {
		fmt.Println("RequestID is zero")
		if cfg.RequestData > 0 {
			fmt.Println("Requesting Data")
			ops.Requesting = true
			err= pow.RequestData(ctx)
			fmt.Println("Done Requesting", err)
			ops.Requesting = false
		}
		return nil, nil
	}
	val, err := DB.Get(fmt.Sprintf("%s%d", db.QueriedValuePrefix, asInt.Uint64()))
	if err != nil {
		ops.log.Error("Problem reading price data from DB: %v\n", err)
		return nil, err
	}
	if len(val) > 0 {
		value, err := hexutil.DecodeBig(string(val))
		if err != nil {
			ops.log.Error("Problem decoding price value: %v\n", err)
			return nil, err
		}
		return &miningCycle{challenge: currentChallenge, difficulty: difficulty, nonce: "", requestID: asInt, value: value}, nil
	}
	ops.log.Warn("No price data found for request id: %d\n", asInt.Uint64())
	return nil, nil
}

func (ops *MinerOps) mine(ctx context.Context, cycle *miningCycle) {
	lastCycle := ops.lastChallenge
	DB := ctx.Value(common.DBContextKey).(db.DB)
	if !ops.Running {
		return
	}
	if lastCycle == nil || bytes.Compare(lastCycle, cycle.challenge) != 0 {
		ops.log.Info("Mining for PoW nonce...")
		nonce := ops.miner.SolveChallenge(cycle.challenge, cycle.difficulty)
		ops.log.Info("Mined nonce", nonce)
		if nonce != "" {
			val, err := DB.Get(fmt.Sprintf("%s%d", db.QueriedValuePrefix, cycle.requestID.Uint64()))
			var priceValue *big.Int
			if err != nil {
				ops.log.Error("Problem reading price data from DB: %v. Using last known value: %v\n", err, cycle.value)
				priceValue = cycle.value
			} else if val != nil{
				priceValue, err = hexutil.DecodeBig(string(val))
				if err != nil {
					ops.log.Error("Problem decoding price value from DB data: %v\n", err)
					priceValue = nil
				}
			}else{
				ops.log.Info("Price is nil, check API and/or PSR value")
				ops.lastChallenge=nil
				return
			}
			if priceValue != nil {
				ops.log.Info("Submitting solution: %v, %v, %v", nonce, priceValue, cycle.requestID)
				pow.SubmitSolution(ctx, cycle.challenge, nonce, priceValue, cycle.requestID)
			}else{
				ops.log.Info("Price is nil, check API and/or PSR value")
				ops.lastChallenge=nil
				return
			}
		}else{
			ops.log.Info("Nonce is nil")
			ops.lastChallenge=nil
			return
		}

	}

}
