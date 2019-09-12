package ops

import (
	"context"
	"os"
	"time"

	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/pow"
	"github.com/tellor-io/TellorMiner/util"
)

//MiningMgr holds items for mining and requesting data
type MiningMgr struct {
	//primary exit channel
	exitCh  chan os.Signal
	log     *util.Logger
	Running bool

	miningWorker *pow.Worker
	//miner's exit channel
	minerExit chan os.Signal

	dataRequester *DataRequester
	//data requester's exit channel
	requesterExit chan os.Signal
}

//CreateMiningManager creates a new manager that mananges mining and data requests
func CreateMiningManager(ctx context.Context, exitCh chan os.Signal, submitter tellorCommon.TransactionSubmitter) (*MiningMgr, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	mExit := make(chan os.Signal)
	rExit := make(chan os.Signal)
	miner := pow.CreateWorker(mExit, submitter, cfg.MiningInterruptCheckInterval)
	dataRequester := CreateDataRequester(rExit, submitter, cfg.RequestDataInterval)
	log := util.NewLogger("ops", "MiningMgr")

	return &MiningMgr{exitCh: exitCh, log: log, Running: false, miningWorker: miner, minerExit: mExit, dataRequester: dataRequester, requesterExit: rExit}, nil
}

//Start will start the mining run loop
func (mgr *MiningMgr) Start(ctx context.Context) {
	mgr.log.Info("Starting mining operations")
	mgr.miningWorker.Start(ctx)
	mgr.dataRequester.Start(ctx)
	go func() {
		<-mgr.exitCh
		mgr.log.Info("Stopping mining worker and data requester")
		mgr.minerExit <- os.Interrupt
		mgr.requesterExit <- os.Interrupt
		time.Sleep(1 * time.Second)
		mgr.Running = mgr.miningWorker.CanMine() || mgr.dataRequester.IsRunning()
	}()
}
