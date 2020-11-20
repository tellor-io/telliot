// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package dataServer

import (
	"fmt"

	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/tellor-io/TellorMiner/pkg/tcontext"
	"github.com/tellor-io/TellorMiner/pkg/testutil"
	"github.com/tellor-io/TellorMiner/pkg/util"
)

func TestDataServer(t *testing.T) {
	exitCh := make(chan int)
	logSetup := util.SetupLogger()
	logger := logSetup("debug")
	ctx, cfg, cleanup := tcontext.CreateTestContext(t)
	defer t.Cleanup(cleanup)

	ds, err := CreateServer(ctx, logger)
	testutil.Ok(t, err, "creating server in test")
	testutil.Ok(t, ds.Start(ctx, exitCh), "starting server")

	time.Sleep(2 * time.Second)
	resp, err := http.Get("http://" + cfg.ServerHost + ":" + strconv.Itoa(int(cfg.ServerPort)) + "/balance")
	testutil.Ok(t, err)
	defer resp.Body.Close()
	fmt.Printf("Finished: %+v", resp)
	exitCh <- 1
	time.Sleep(1 * time.Second)
	testutil.Assert(t, ds.Stopped, "Did not stop server")
}
