// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package dataServer

import (
	"context"
	"math/big"

	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/go-kit/kit/log/level"
	_ "github.com/prometheus/client_golang/prometheus"
	_ "github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/rest"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/testutil"
)

func TestDataServer(t *testing.T) {
	cfg := config.OpenTestConfig(t)
	logger := logging.NewLogger()
	DB, cleanup := db.OpenTestDB(t)
	defer t.Cleanup(cleanup)

	done := make(chan bool)
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
	proxy, err := db.OpenLocal(logger, cfg, DB)
	testutil.Ok(t, err)

	accounts, err := rpc.NewAccounts(cfg)
	testutil.Ok(t, err)
	contract, err := contracts.NewITellor(client)
	testutil.Ok(t, err)
	// We need to unregister prometheus counter.
	ds, err := CreateServer(logger, cfg, proxy, client, contract, accounts)
	testutil.Ok(t, err, "creating server in test")
	testutil.Ok(t, ds.Start(context.Background(), done), "starting server")

	srv, err := rest.Create(logger, cfg, context.Background(), proxy, cfg.DataServer.ListenHost, cfg.DataServer.ListenPort)
	testutil.Ok(t, err)
	srv.Start()

	time.Sleep(2 * time.Second)
	resp, err := http.Get("http://" + cfg.DataServer.ListenHost + ":" + strconv.Itoa(int(cfg.DataServer.ListenPort)) + "/balance")
	testutil.Ok(t, err)
	defer resp.Body.Close()
	level.Info(logger).Log("response finished", "resp", resp)
	done <- true
	time.Sleep(1 * time.Second)
	testutil.Assert(t, ds.Stopped, "Did not stop server")
}
