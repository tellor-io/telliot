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
	lastChallenge *miningCycle
	miner         *pow.PoWSolver
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
					//FIXME: this needs to be refactored/designed because it assumes that the
					//mining cycle will complete before the next time cycle. Even though the
					//mining is synchronous, many time entries will be pushed into the ticker
					//channel and a bunch of extraneous requests will happen after a full mine.
					cycle, err := ops.buildNextCycle(ctx)
					if err == nil {
						if cycle != nil && !ops.miner.IsMining() {
							ops.log.Info("Requesting mining cycle with vars: %+v\n", cycle)
							go ops.mine(ctx, cycle)
						}
					}
				}
			}
		}
	}()
}

func (ops *MinerOps) buildNextCycle(ctx context.Context) (*miningCycle, error) {

	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	DB := ctx.Value(common.DBContextKey).(db.DB)
	currentChallenge, err := DB.Get(db.CurrentChallengeKey)
	if err != nil {
		ops.log.Error("Problem reading challenge in miner run loop: %v\n", err)
		return nil, err
	}
	if ops.lastChallenge != nil && bytes.Compare(currentChallenge, ops.lastChallenge.challenge) == 0 {
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
	fmt.Println("mining status")
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
			pow.RequestData(ctx)
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
	if lastCycle == nil || bytes.Compare(lastCycle.challenge, cycle.challenge) != 0 {
		ops.log.Info("Mining for PoW nonce...")
		//FIXME: need to make sure that if the machine is stopped that any ongoing PoW computation will end
		nonce := ops.miner.SolveChallenge(cycle.challenge, cycle.difficulty)
		ops.log.Info("Mined nonce", nonce)
		if nonce != "" {
			val, err := DB.Get(fmt.Sprintf("%s%d", db.QueriedValuePrefix, cycle.requestID.Uint64()))
			var priceValue *big.Int
			if err != nil {
				ops.log.Error("Problem reading price data from DB: %v. Using last known value: %v\n", err, cycle.value)
				priceValue = cycle.value
			} else {
				priceValue, err = hexutil.DecodeBig(string(val))
				if err != nil {
					ops.log.Error("Problem decoding price value from DB data: %v\n", err)
					priceValue = nil
				}
			}
			if priceValue != nil {
				ops.lastChallenge = cycle
				ops.log.Info("Submitting solution: %v, %v, %v", nonce, priceValue, cycle.requestID)
				//pow.SubmitTransaction(nonce, priceValue, cycle.requestID)
				pow.SubmitSolution(ctx, nonce, priceValue, cycle.requestID)
				//pow.SubmitSolution(ctx, nonce, big.NewInt(221000), big.NewInt(1))
			}
		}

	}

}
