package ops

import (
	"context"
	"fmt"
	"os"
	"time"

	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/pow"
	"github.com/tellor-io/TellorMiner/util"
)

type WorkSource interface {
	GetWork(input chan *pow.Work) (*pow.Work,bool)
}

type SolutionSink interface {
	Submit(context.Context, *pow.Result) bool
}

//MiningMgr holds items for mining and requesting data
type MiningMgr struct {
	//primary exit channel
	exitCh  chan os.Signal
	log     *util.Logger
	Running bool

	group      *pow.MiningGroup
	tasker     WorkSource
	solHandler SolutionSink
	solution   *pow.Result

	dataRequester *DataRequester
	//data requester's exit channel
	requesterExit chan os.Signal
}

//CreateMiningManager creates a new manager that mananges mining and data requests
func CreateMiningManager(ctx context.Context, exitCh chan os.Signal, submitter tellorCommon.TransactionSubmitter) (*MiningMgr, error) {
	cfg := config.GetConfig()

	group, err := pow.SetupMiningGroup(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to setup miners: %s", err.Error())
	}

	mng := &MiningMgr{
		exitCh:     exitCh,
		log:        util.NewLogger("ops", "MiningMgr"),
		Running:    false,
		group:      group,
		tasker:     nil,
		solution: 	nil,
		solHandler: nil,
	}

	if cfg.EnablePoolWorker {
		pool := pow.CreatePool(cfg, group)
		mng.tasker = pool
		mng.solHandler = pool
	} else {
		proxy := ctx.Value(tellorCommon.DataProxyKey).(db.DataServerProxy)
		mng.tasker = pow.CreateTasker(cfg, proxy)
		mng.solHandler = pow.CreateSolutionHandler(cfg, submitter, proxy)
		if cfg.RequestData > 0 {
			fmt.Println("dataRequester created")
			mng.dataRequester = CreateDataRequester(exitCh, submitter, cfg.RequestDataInterval.Duration, proxy)
		}
	}
	return mng, nil
}

//Start will start the mining run loop
func (mgr *MiningMgr) Start(ctx context.Context) {
	mgr.Running = true
	go func(ctx context.Context) {
		cfg := config.GetConfig()

		ticker := time.NewTicker(cfg.MiningInterruptCheckInterval.Duration)

		//if you make these buffered, think about the effects on synchronization!
		input := make(chan *pow.Work)
		output := make(chan *pow.Result)
		if cfg.RequestData > 0 {
			mgr.dataRequester.Start(ctx)
		}

		//start the mining group
		go mgr.group.Mine(input, output)

		// sends work to the mining group
		sendWork := func() {
			if cfg.EnablePoolWorker {
				mgr.tasker.GetWork(input)
			} else {
				work,instantSubmit := mgr.tasker.GetWork(input)
				if instantSubmit{
					if mgr.solution == nil {
						fmt.Println("Instant Submit Called! ")
						mgr.solution = &pow.Result{Work:work, Nonce:"1"}
					} else{
						fmt.Println("Trying Resubmit...")
					}
				}else if work != nil {
					mgr.solution = nil
					input <- work
				}
				if mgr.solution != nil{
					goodSubmit := mgr.solHandler.Submit(ctx,mgr.solution)
					if goodSubmit {
						mgr.solution = nil
					}
				} 
			}
		}
		//send the initial challenge
		sendWork()	
		for {
			select {
			//boss wants us to quit for the day
			case <-mgr.exitCh:
				//exit
				input <- nil

			//found a solution
			case result := <-output:
				if result == nil {
					mgr.Running = false
					return
				}
				mgr.solution = result
				goodSubmit := mgr.solHandler.Submit(ctx,mgr.solution)
				if goodSubmit {
					mgr.solution = nil
				}
				sendWork()

			//time to check for a new challenge
			case _ = <-ticker.C:
				sendWork()
			}
		}
	}(ctx)
}
