// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package rest

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/db"
)

func TestServer(t *testing.T) {
	DB, err := db.Open(filepath.Join(os.TempDir(), "test_server"))
	if err != nil {
		t.Fatal(err)
	}
	balInt := big.NewInt(350000)
	err = DB.Put(db.BalanceKey, []byte(hexutil.EncodeBig(balInt)))
	if err != nil {
		t.Fatal(err)
	}
	proxy, err := db.OpenRemoteDB(DB)
	if err != nil {
		t.Fatal(err)
	}
	cfg := config.GetConfig()
	ctx := context.WithValue(context.Background(), common.DBContextKey, DB)
	ctx = context.WithValue(ctx, common.DataProxyKey, proxy)
	srv, err := Create(ctx, cfg.ServerHost, cfg.ServerPort)
	if err != nil {
		t.Fatal(err)
	}

	srv.Start()
	defer func() {
		if err := srv.Stop(); err != nil {
			fmt.Println("error stoping the server", err)
		}
	}()

	data, err := proxy.Get(db.BalanceKey)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatal("Expected data to be returned")
	}
	asInt, err := hexutil.DecodeBig(string(data))
	if err != nil {
		t.Fatal(err)
	}
	if asInt.Cmp(balInt) != 0 {
		t.Fatalf("Expected %v but received %v as balance", balInt, asInt)
	}

	t.Logf("Retrieved balance from server: %+v\n", asInt)

	/***
	ctx := context.WithValue(context.Background(), common.DBContextKey, DB)
	srv, err := Create(ctx, "localhost", 5000)
	if err != nil {
		t.Fatal(err)
	}
	srv.Start()
	defer srv.Stop()

	resp, err := http.Get("http:// ocalhost:5000/balance")
	if err != nil {
		t.Fatal(err)
	}
	var bal BalTest
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&bal)
	if !strings.Contains(bal.Balance, "0x") {
		t.Fatal("Missing balance in response")
	} else {
		t.Logf("Retrieved balance from server: %+v\n", bal)
	}
	**/

}
