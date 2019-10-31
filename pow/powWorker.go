package pow

import (
	"context"
	"os"
	"sync"
	"time"

	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/db"
)

//Worker state for mining operation
type Worker struct {
	running bool
	loop    *miningLoop
	tasker  *miningTasker
	handler *solutionHandler
}

var (
	//global lock for submitting solution
	submitMutex sync.Mutex
)

//CreateWorker creates a new worker instance
func CreateWorker(
	id int,
	submitter tellorCommon.TransactionSubmitter,
	checkIntervalSeconds time.Duration,
	proxy db.DataServerProxy) (*Worker, error) {
	if checkIntervalSeconds == 0 {
		checkIntervalSeconds = 15
	}
	loop, err := createMiningLoop(id)
	if err != nil {
		return nil, err
	}
	tasker, err := createTasker(id, loop.taskCh, loop.cancelCh, checkIntervalSeconds, proxy)
	if err != nil {
		return nil, err
	}
	handler, err := createSolutionHandler(id, loop.solutionCh, submitter, proxy)
	if err != nil {
		return nil, err
	}

	return &Worker{

		loop:    loop,
		tasker:  tasker,
		handler: handler,
	}, nil
}

//Start kicks of go routines to check for challenge changes, mine, and submit solutions
func (w *Worker) Start(ctx context.Context) {
	if w.running {
		return
	}
	w.handler.Start(ctx)
	w.tasker.Start(ctx)
	w.loop.Start(ctx)
}

//Stop stops worker and its resources
func (w *Worker) Stop(ctx context.Context) {
	w.loop.exitCh <- os.Interrupt
	w.tasker.exitCh <- os.Interrupt
	w.handler.exitCh <- os.Interrupt
	w.running = false
}

//CanMine checks whether this mining worker is allowed to mine right now
func (w *Worker) CanMine() bool {
	return w.loop.canMine
}
