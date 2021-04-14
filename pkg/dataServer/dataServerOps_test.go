// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package dataServer

// FIXME: commented because of unrelated error: duplicate metrics collector registration attempted!
// import (
// 	"context"
// 	"testing"
// 	"time"

// 	"github.com/tellor-io/telliot/pkg/config"
// 	"github.com/tellor-io/telliot/pkg/contracts"
// 	"github.com/tellor-io/telliot/pkg/db"
// 	"github.com/tellor-io/telliot/pkg/logging"
// 	"github.com/tellor-io/telliot/pkg/rpc"
// 	"github.com/tellor-io/telliot/pkg/testutil"
// )

// func TestDataServerOps(t *testing.T) {
// 	cfg,err := config.OpenTestConfig()
// testutil.Ok(t, err)
// 	logger := logging.NewLogger()
// 	DB, cleanup := db.OpenTestDB(t)
// 	defer t.Cleanup(cleanup)
// 	client := rpc.NewMockClient()

// 	proxy, err := db.OpenLocal(logger, cfg, DB)
// 	testutil.Ok(t, err)

// 	ctx := context.Background()
// 	ops, err := CreateDataServerOps(ctx, logger, cfg, proxy, client, &contracts.ITellor{}, nil)
// 	testutil.Ok(t, err)

// 	testutil.Ok(t, ops.Start(), "starting server")
// 	time.Sleep(2 * time.Second)
// 	ops.Stop()
// 	time.Sleep(1 * time.Second)
// }
