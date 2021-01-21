// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package ops

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/testutil"
	"github.com/tellor-io/telliot/pkg/util"
)

func TestDataServerOps(t *testing.T) {
	cfg := config.OpenTestConfig(t)
	exitCh := make(chan os.Signal)
	logger := util.SetupLogger("debug")
	DB, cleanup := db.OpenTestDB(t)
	defer t.Cleanup(cleanup)
	client := rpc.NewMockClient()

	proxy, err := db.OpenLocal(cfg, DB)
	testutil.Ok(t, err)

	ctx := context.Background()
	ops, err := CreateDataServerOps(ctx, logger, cfg, proxy, client, nil, nil, exitCh)
	testutil.Ok(t, err)

	testutil.Ok(t, ops.Start(ctx), "starting server")
	time.Sleep(2 * time.Second)
	exitCh <- os.Interrupt
	time.Sleep(1 * time.Second)
	testutil.Assert(t, !ops.Running, "data server is still running after stopping")
}
