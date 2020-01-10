package ops

import (
	"context"
	"fmt"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/pow"
	"github.com/tellor-io/TellorMiner/util"
	"log"
	"math"
	"math/rand"
	"os"
)

//MiningMgr holds items for mining and requesting data
type MiningMgr struct {
	//primary exit channel
	exitCh  chan os.Signal
	log     *util.Logger
	Running bool

	group  *pow.MiningGroup
	tasker *pow.MiningTasker
	solHandler *pow.SolutionHandler

	proxy db.DataServerProxy

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

	group, err := pow.SetupMiningGroup(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to setup miners: %s", err.Error())
	}

	proxy := ctx.Value(tellorCommon.DataProxyKey).(db.DataServerProxy)

	tasker := pow.CreateTasker(cfg, proxy)
	solHandler := pow.CreateSolutionHandler(cfg, submitter, proxy)

	rExit := make(chan os.Signal)

	dataRequester := CreateDataRequester(rExit, submitter, cfg.RequestDataInterval.Duration, proxy)
	log := util.NewLogger("ops", "MiningMgr")

	return &MiningMgr{
		exitCh:  exitCh,
		log:     log,
		Running: false,
		group:   group,
		proxy:   proxy,
		tasker:  tasker,
		solHandler: solHandler,
		dataRequester: dataRequester,
		requesterExit: rExit}, nil
}

//Start will start the mining run loop
func (mgr *MiningMgr) Start(ctx context.Context) {
	go func(ctx context.Context) {
		cfg, err := config.GetConfig()
		if err != nil {
			log.Fatal(err)
		}

		start := uint64(rand.Int63())
		for {
			select {
				case <-mgr.exitCh:
					//exit
					return
				default:
					//continue
			}
			challenge := mgr.tasker.PullUpdates()
			hashSettings := pow.NewHashSettings(challenge, cfg.PublicAddress)

			nonce, err := mgr.group.Mine(hashSettings, start, math.MaxInt64, cfg.MiningInterruptCheckInterval.Duration)
			if err != nil {
				log.Fatal(err)
			}
			if nonce != "" {
				mgr.solHandler.HandleSolution(ctx, challenge, nonce)
			}
		}
	}(ctx)
}
