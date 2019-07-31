package tracker

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	tellor "github.com/tellor-io/TellorMiner/contracts"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rpc"
	"github.com/tellor-io/TellorMiner/util"
)

var runnerLog = util.NewLogger("tracker", "Runner")

//Runner will execute all configured trackers
type Runner struct {
	client rpc.ETHClient
	db     db.DB
}

//NewRunner will create a new runner instance
func NewRunner(client rpc.ETHClient, db db.DB) (*Runner, error) {
	return &Runner{client: client, db: db}, nil
}

//Start will kick off the runner until the given exit channel selects.
func (r *Runner) Start(ctx context.Context, exitCh chan int) error {
	cfg, err := config.GetConfig()
	if err != nil {
		runnerLog.Error("Problem getting config", err)
		return err
	}
	sleep := cfg.TrackerSleepCycle
	trackerNames := cfg.Trackers
	trackers := make([]Tracker, len(trackerNames))
	for i := 0; i < len(trackers); i++ {
		t, err := createTracker(trackerNames[i])
		if err != nil {
			runnerLog.Error("Problem creating tracker: %s\n", err.Error())
		}
		trackers[i] = t
	}

	masterInstance := ctx.Value(tellorCommon.MasterContractContextKey)
	if masterInstance == nil {
		contractAddress := common.HexToAddress(cfg.ContractAddress)
		masterInstance, err = tellor.NewTellorMaster(contractAddress, r.client)
		if err != nil {
			runnerLog.Error("Problem creating tellor master instance: %v\n", err)
			return err
		}
		ctx = context.WithValue(ctx, tellorCommon.MasterContractContextKey, masterInstance)
	}

	sleepTime := time.Duration(sleep) * time.Second
	runnerLog.Info("Trackers will run every %v\n", sleepTime)
	ticker := time.NewTicker(sleepTime)
	if ctx.Value(tellorCommon.ClientContextKey) == nil {
		ctx = context.WithValue(ctx, tellorCommon.ClientContextKey, r.client)
	}
	if ctx.Value(tellorCommon.DBContextKey) == nil {
		ctx = context.WithValue(ctx, tellorCommon.DBContextKey, r.db)
	}

	go func() {
		r.callTrackers(ctx, &trackers)
		for {
			runnerLog.Info("Waiting for next tracker run cycle...")
			select {
			case _ = <-exitCh:
				{
					runnerLog.Info("Exiting run loop")
					ticker.Stop()
					return
				}
			case _ = <-ticker.C:
				{
					runnerLog.Info("Running trackers...")
					r.callTrackers(ctx, &trackers)
				}
			}
		}
	}()

	return nil

}

func (r *Runner) callTrackers(ctx context.Context, trackers *[]Tracker) error {
	for _, t := range *trackers {
		runnerLog.Info("Calling tracker: %v\n", t)
		err := t.Exec(ctx)
		if err != nil {
			runnerLog.Error("Problem in tracker: %v\n", err)
		}
	}
	return nil
}
