// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package ops

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tellor-io/TellorMiner/pkg/contracts/getter"
	"github.com/tellor-io/TellorMiner/pkg/contracts/tellor"
	"github.com/tellor-io/TellorMiner/pkg/testutil"
)

// TestMain check for goroutine leaks.
// It ensures that at the end of the tests there are no remaining go routines.
func TestMain(m *testing.M) {
	testutil.TolerantVerifyLeakMain(m)
}

func TestProfitCalc(t *testing.T) {
	// Generate a new random account and a funded simulator
	key, err := crypto.GenerateKey()
	testutil.Ok(t, err)
	auth := bind.NewKeyedTransactor(key)

	sim := backends.NewSimulatedBackend(core.GenesisAlloc{auth.From: {Balance: big.NewInt(10000000000)}}, 100000000000000000)

	// Deploy a token contract on the simulated blockchain
	_, tx1, tellor, err := tellor.DeployTellor(auth, sim)
	testutil.Ok(t, err)

	_, tx2, tellorGetter, err := getter.DeployTellorGetters(auth, sim)
	testutil.Ok(t, err)

	sim.Commit()

	name, err := tellor.Name(&bind.CallOpts{Pending: true})
	testutil.Ok(t, err)

	fmt.Println("name", name)

	rcpt, err := bind.WaitMined(context.Background(), sim, tx1)
	testutil.Ok(t, err)
	if rcpt.Status != 1 {
		t.Fatal("deploy transaction failed")
	}
	rcpt, err = bind.WaitMined(context.Background(), sim, tx2)
	testutil.Ok(t, err)
	if rcpt.Status != 1 {
		t.Fatal("deploy transaction failed")
	}

	a, b, c, d, e, f, err := tellorGetter.GetCurrentVariables(nil)
	testutil.Ok(t, err)
	fmt.Println(a, b, c, d, e, f)

}
