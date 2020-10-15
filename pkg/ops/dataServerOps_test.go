// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package ops

import (
	"os"
	"testing"
	"time"

	"github.com/tellor-io/TellorMiner/pkg/testutil"
)

func TestDataServerOps(t *testing.T) {

	exitCh := make(chan os.Signal)

	ctx, _, cleanup := testutil.CreateContext(t)
	defer t.Cleanup(cleanup)

	ops, err := CreateDataServerOps(ctx, exitCh)
	if err != nil {
		t.Fatal(err)
	}
	if err := ops.Start(ctx); err != nil {
		t.Fatal("error starting the data server", err)
	}
	time.Sleep(2 * time.Second)
	exitCh <- os.Interrupt
	time.Sleep(1 * time.Second)
	if ops.Running {
		t.Fatal("data server is still running after stopping")
	}
}
