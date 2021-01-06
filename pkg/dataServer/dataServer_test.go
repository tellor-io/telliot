// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package dataServer

import (
	"context"
	"fmt"
	"math/big"

	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rest"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/testutil"
	"github.com/tellor-io/telliot/pkg/util"
)

func TestDataServer(t *testing.T) {
	cfg := config.OpenTestConfig(t)
	exitCh := make(chan int)
	logSetup := util.SetupLogger()
	logger := logSetup("debug")
	DB, cleanup := db.OpenTestDB(t)
	defer t.Cleanup(cleanup)

	startBal := big.NewInt(356000)
	opts := &rpc.MockOptions{
		ETHBalance:    startBal,
		Nonce:         1,
		GasPrice:      big.NewInt(700000000),
		TokenBalance:  big.NewInt(0),
		Top50Requests: []*big.Int{},
		DisputeStatus: big.NewInt(1),
	}
	client := rpc.NewMockClientWithValues(opts)
	proxy, err := db.OpenLocalProxy(DB)
	testutil.Ok(t, err)

	ctx := context.Background()
	account, err := rpc.NewAccount(cfg)
	testutil.Ok(t, err)
	contract, err := contracts.NewTellor(cfg, client)
	testutil.Ok(t, err)
	ds, err := CreateServer(ctx, logger, cfg, DB, client, &contract, &account)
	testutil.Ok(t, err, "creating server in test")
	testutil.Ok(t, ds.Start(ctx, exitCh), "starting server")

	srv, err := rest.Create(ctx, proxy, cfg.ServerHost, cfg.ServerPort)
	testutil.Ok(t, err)
	srv.Start()

	time.Sleep(2 * time.Second)
	resp, err := http.Get("http://" + cfg.ServerHost + ":" + strconv.Itoa(int(cfg.ServerPort)) + "/balance")
	testutil.Ok(t, err)
	defer resp.Body.Close()
	fmt.Printf("Finished: %+v", resp)
	exitCh <- 1
	time.Sleep(1 * time.Second)
	testutil.Assert(t, ds.Stopped, "Did not stop server")
}
