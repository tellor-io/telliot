package ops

import (
	"context"
	"fmt"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/pow"
	"math"
	"math/big"
	"time"
)

func RunBenchmark(ctx context.Context) error {
	cfg := config.GetConfig()

	group, err := pow.SetupMiningGroup(cfg)
	if err != nil {
		return fmt.Errorf("failed to setup miners: %s", err.Error())
	}

	input := make(chan *pow.Work)
	output := make(chan *pow.Result)

	go group.Mine(input, output)

	input <- &pow.Work{
		Challenge: &pow.MiningChallenge{
			Challenge: make([]byte, 32, 32),
			Difficulty: big.NewInt(math.MaxInt64),
			RequestID: big.NewInt(0),
		},
		Start:     0,
		N:         math.MaxInt64,
	}

	time.Sleep(1 * time.Second)

	input <- nil
	<-output

	group.PrintHashRateSummary()

	return nil
}
