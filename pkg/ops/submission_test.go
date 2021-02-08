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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/contracts/oldTellor"
	"github.com/tellor-io/telliot/pkg/contracts/proxy"
	"github.com/tellor-io/telliot/pkg/testutil"
)

// TestMain check for goroutine leaks.
// It ensures that at the end of the tests there are no remaining go routines.
func TestMain(m *testing.M) {
	testutil.TolerantVerifyLeakMain(m)
}

func TestSubmission(t *testing.T) {

	initialMiner := common.HexToAddress("e010aC6e0248790e08F42d5F697160DEDf97E024")
	backend, transactor, _ := getbackendBackend(t)
	// Deploy a token contract on the simulated blockchain
	_, tx1, _, err := contracts.DeployOldTellor(transactor, backend)
	testutil.Ok(t, err)
	backend.Commit()

	addr, err := bind.WaitDeployed(context.Background(), backend, tx1)
	testutil.Ok(t, err)
	fmt.Println(addr)

	masterAdd, tx2, master, err := proxy.DeployTellorMaster(transactor, backend, addr)
	testutil.Ok(t, err)
	backend.Commit()

	_, err = bind.WaitDeployed(context.Background(), backend, tx2)
	testutil.Ok(t, err)
	backend.Commit()

	tellor, err := oldTellor.NewOldTellor(masterAdd, backend)
	testutil.Ok(t, err)
	backend.Commit()

	// _, err = tellor.RequestStakingWithdraw(transactor)
	// testutil.Ok(t, err)

	tx, err := tellor.RequestData(transactor, "data", "sym", big.NewInt(10), big.NewInt(0))
	testutil.Ok(t, err)
	println(tx)
	// a, b, c, d, e, f, err := master.GetCurrentVariables(nil)
	// testutil.Ok(t, err)
	// fmt.Println(a, b, c, d, e, f)
	t.Fatal()
	fmt.Println(master)
	fmt.Println(tellor)
	fmt.Println(initialMiner)
}

func getbackendBackend(t *testing.T) (*backends.SimulatedBackend, *bind.TransactOpts, common.Address) {
	sk, err := crypto.HexToECDSA("3a10b4bc1258e8bfefb95b498fb8c0f0cd6964a811eabca87df5630bcacd7216")
	testutil.Ok(t, err)
	faucetAddr := crypto.PubkeyToAddress(sk.PublicKey)
	addr := map[common.Address]core.GenesisAccount{
		common.BytesToAddress([]byte{1}): {Balance: big.NewInt(1)}, // ECRecover
		common.BytesToAddress([]byte{2}): {Balance: big.NewInt(1)}, // SHA256
		common.BytesToAddress([]byte{3}): {Balance: big.NewInt(1)}, // RIPEMD
		common.BytesToAddress([]byte{4}): {Balance: big.NewInt(1)}, // Identity
		common.BytesToAddress([]byte{5}): {Balance: big.NewInt(1)}, // ModExp
		common.BytesToAddress([]byte{6}): {Balance: big.NewInt(1)}, // ECAdd
		common.BytesToAddress([]byte{7}): {Balance: big.NewInt(1)}, // ECScalarMul
		common.BytesToAddress([]byte{8}): {Balance: big.NewInt(1)}, // ECPairing
		faucetAddr:                       {Balance: new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(9))},
	}
	alloc := core.GenesisAlloc(addr)
	transactor, err := bind.NewKeyedTransactorWithChainID(sk, big.NewInt(1337))
	testutil.Ok(t, err)
	return backends.NewSimulatedBackend(alloc, 80000000), transactor, faucetAddr
}
