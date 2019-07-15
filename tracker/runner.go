package tracker

import (
	"context"
	"fmt"
	"time"

	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rpc"
)

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
		return err
	}
	sleep := cfg.TrackerSleepCycle
	trackerNames := cfg.Trackers
	trackers := make([]Tracker, len(trackerNames))
	for i := 0; i < len(trackers); i++ {
		t, err := createTracker(trackerNames[i])
		if err != nil {
			fmt.Printf("Problem creating tracker: %s\n", err.Error())
		}
		trackers[i] = t
	}

	ticker := time.NewTicker(time.Duration(sleep) * time.Second)
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

					c := context.WithValue(ctx, ClientContextKey, r.client)
					c = context.WithValue(c, DBContextKey, r.db)
					for _, t := range trackers {
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
