package tracker

import (
	"context"
	"fmt"
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

	contractAddress := common.HexToAddress(cfg.ContractAddress)
	masterInstance, err := tellor.NewTellorMaster(contractAddress, r.client)
	if err != nil {
		runnerLog.Error("Problem creating tellor master instance: %v\n", err)
		return err
	}

	sleepTime := time.Duration(sleep) * time.Second
	fmt.Printf("Trackers will run every %v\n", sleepTime)
	ticker := time.NewTicker(sleepTime)
	go func() {
		defer r.client.Close()
		defer r.db.Close()

		for {
			select {
			case _ = <-exitCh:
				{
					fmt.Println("Exiting run loop")
					ticker.Stop()
					return
				}
			case _ = <-ticker.C:
				{
					fmt.Println("Running trackers...")
					c := context.WithValue(ctx, tellorCommon.ClientContextKey, r.client)
					c = context.WithValue(c, tellorCommon.DBContextKey, r.db)
					c = context.WithValue(c, tellorCommon.MasterContractContextKey, masterInstance)
					for _, t := range trackers {
						fmt.Printf("Calling tracker: %v\n", t)
						err := t.Exec(c)
						if err != nil {
							fmt.Println("Problem in tracker", err)
						}
					}

				}
			}
		}
	}()

	return nil

}
