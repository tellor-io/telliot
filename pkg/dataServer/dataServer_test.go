// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package dataServer

import (
	"fmt"

	"net/http"
	"os"
	"strconv"
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
func TestDataServer(t *testing.T) {
	exitCh := make(chan int)
	logger := setupLogger()
	ctx, cfg, cleanup := testutil.CreateContext(t)
	defer t.Cleanup(cleanup)

	ds, err := CreateServer(ctx, logger)
	if err != nil {
		level.Error(logger).Log("msg", "error creating server in test", "err", err)
		os.Exit(1)
	}
	if err := ds.Start(ctx, logger, exitCh); err != nil {
		t.Fatal(err)
	}

	time.Sleep(2 * time.Second)

	resp, err := http.Get("http://" + cfg.ServerHost + ":" + strconv.Itoa(int(cfg.ServerPort)) + "/balance")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Printf("Finished: %+v", resp)
	exitCh <- 1
	time.Sleep(1 * time.Second)
	if !ds.Stopped {
		t.Fatal("Did not stop server")
	}
}
