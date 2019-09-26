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

	miners []*WorkerWrapper

	//miner's exit channel
	//minerExit chan os.Signal

	dataRequester *DataRequester
	//data requester's exit channel
	requesterExit chan os.Signal
}

//WorkerWrapper allows us to stand up multiple mining workers, one per cpu
type WorkerWrapper struct {
	miner  *pow.Worker
	exitCh chan os.Signal
}

//CreateMiningManager creates a new manager that mananges mining and data requests
func CreateMiningManager(ctx context.Context, exitCh chan os.Signal, submitter tellorCommon.TransactionSubmitter) (*MiningMgr, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	var miners []*WorkerWrapper
	miners = make([]*WorkerWrapper, cfg.NumProcessors, cfg.NumProcessors)
	for i := 0; i < cfg.NumProcessors; i++ {
		mExit := make(chan os.Signal)
		miner := pow.CreateWorker(mExit, i, submitter, cfg.MiningInterruptCheckInterval)
		miners[i] = &WorkerWrapper{miner: miner, exitCh: mExit}
	}

	rExit := make(chan os.Signal)

	dataRequester := CreateDataRequester(rExit, submitter, cfg.RequestDataInterval)
	log := util.NewLogger("ops", "MiningMgr")

	return &MiningMgr{exitCh: exitCh, log: log, Running: false, miners: miners, dataRequester: dataRequester,
		requesterExit: rExit}, nil
}

//Start will start the mining run loop
func (mgr *MiningMgr) Start(ctx context.Context) {
	mgr.log.Info("Starting mining operations")
	for _, m := range mgr.miners {
		m.miner.Start(ctx)
	}
	//mgr.miningWorker.Start(ctx)

	mgr.dataRequester.Start(ctx)
	go func() {
		<-mgr.exitCh
		mgr.log.Info("Stopping mining worker and data requester")
		//mgr.minerExit <- os.Interrupt
		running := false
		for _, m := range mgr.miners {
			m.exitCh <- os.Interrupt
		}
		mgr.requesterExit <- os.Interrupt
		time.Sleep(1 * time.Second)
		for _, m := range mgr.miners {
			if m.miner.CanMine() {
				//at least one can mine
				running = true
			}
		}

		mgr.Running = running || mgr.dataRequester.IsRunning()
	}()
}
