// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package ops

import (
	"os"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/tellor-io/TellorMiner/pkg/tcontext"
	"github.com/tellor-io/TellorMiner/pkg/testutil"
	"github.com/tellor-io/TellorMiner/pkg/util"
)

func TestDataServerOps(t *testing.T) {

	exitCh := make(chan os.Signal)
	logSetup := util.SetupLogger()
	logger := logSetup("debug")
	ctx, _, cleanup := tcontext.CreateTestContext(t)
	defer t.Cleanup(cleanup)

	ops, err := CreateDataServerOps(ctx, logger, exitCh)
	if err != nil {
		testutil.Ok(t, err)
	}
	if err := ops.Start(ctx); err != nil {
		testutil.Ok(t, errors.Wrap(err, "error starting the data server"))
	}
	time.Sleep(2 * time.Second)
	exitCh <- os.Interrupt
	time.Sleep(1 * time.Second)
	if ops.Running {
		testutil.Ok(t, errors.New("data server is still running after stopping"))
	}
}
