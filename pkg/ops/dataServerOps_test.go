// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package ops

import (
	"os"
	"testing"
	"time"

	"github.com/tellor-io/telliot/pkg/tcontext"
	"github.com/tellor-io/telliot/pkg/testutil"
	"github.com/tellor-io/telliot/pkg/util"
)

func TestDataServerOps(t *testing.T) {

	exitCh := make(chan os.Signal)
	logSetup := util.SetupLogger()
	logger := logSetup("debug")
	ctx, _, cleanup := tcontext.CreateTestContext(t)
	defer t.Cleanup(cleanup)

	ops, err := CreateDataServerOps(ctx, logger, exitCh)
	testutil.Ok(t, err)

	testutil.Ok(t, ops.Start(ctx), "starting server")
	time.Sleep(2 * time.Second)
	exitCh <- os.Interrupt
	time.Sleep(1 * time.Second)
	testutil.Assert(t, !ops.Running, "data server is still running after stopping")
}
