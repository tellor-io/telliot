package tracker

import (
	"context"
	"testing"
	"time"

	"github.com/tellor-io/TellorMiner/rpc"
)

func TestRunner(t *testing.T) {
	exitCh := make(chan int)
	client := rpc.NewMockClient()
	runner, _ := NewRunner(client)

	ctx := context.Background()
	runner.Start(ctx, exitCh)
	time.Sleep(2 * time.Second)
	exitCh <- 1
	time.Sleep(1 * time.Second)
}
