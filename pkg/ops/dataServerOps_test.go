// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package ops

import (
	"os"
	"testing"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/tellor-io/TellorMiner/pkg/testutil"
)

func setupLogger() log.Logger {
	lvl := level.AllowInfo()

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, lvl)

	return log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
}
func TestDataServerOps(t *testing.T) {

	exitCh := make(chan os.Signal)
	logger := setupLogger()
	ctx, _, cleanup := testutil.CreateContext(t)
	defer t.Cleanup(cleanup)

	ops, err := CreateDataServerOps(ctx, logger, exitCh)
	if err != nil {
		t.Fatal(err)
	}
	if err := ops.Start(ctx, logger); err != nil {
		t.Fatal("error starting the data server", err)
	}
	time.Sleep(2 * time.Second)
	exitCh <- os.Interrupt
	time.Sleep(1 * time.Second)
	if ops.Running {
		t.Fatal("data server is still running after stopping")
	}
}
