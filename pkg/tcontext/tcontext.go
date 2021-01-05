// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tcontext

import (
	"context"
	"math/big"
	"testing"

	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"

	"github.com/tellor-io/telliot/pkg/common"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts/master"
	"github.com/tellor-io/telliot/pkg/contracts/proxy"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/testutil"
)

func CreateTestContext(t *testing.T) (context.Context, *config.Config, func()) {
	cfg := config.OpenTestConfig(t)
	// Don't need any trackers for the tests.
	cfg.Trackers = make(map[string]bool)

	dbLocal, cleanup := db.OpenTestDB(t)
	ctx := context.WithValue(context.Background(), common.DBContextKey, dbLocal)

	hash := math.PaddedBigBytes(big.NewInt(256), 32)
	var b32 [32]byte
	for i, v := range hash {
		b32[i] = v
	}

	opts := &rpc.MockOptions{
		ETHBalance:    big.NewInt(300000),
		Nonce:         1,
		GasPrice:      big.NewInt(7000000000),
		TokenBalance:  big.NewInt(0),
		Top50Requests: []*big.Int{},
		CurrentChallenge: &rpc.CurrentChallenge{
			ChallengeHash: b32, RequestID: big.NewInt(1),
			Difficulty: big.NewInt(500), QueryString: "json(https://coinbase.com)",
			Granularity: big.NewInt(1000), Tip: big.NewInt(0),
		},
	}
	client := rpc.NewMockClientWithValues(opts)
	ctx = context.WithValue(ctx, common.ClientContextKey, client)

	contractAddress := eth_common.HexToAddress(cfg.ContractAddress)
	m, err := master.NewTellor(contractAddress, client)
	testutil.Ok(t, err)

	ctx = context.WithValue(ctx, common.ContractAddress, contractAddress)
	ctx = context.WithValue(ctx, common.ContractsTellorContextKey, m)

	instanceTellor, err := master.NewTellor(contractAddress, client)
	testutil.Ok(t, err)

	ctx = context.WithValue(ctx, common.ContractsTellorContextKey, instanceTellor)

	instanceGetter, err := proxy.NewTellorGetters(contractAddress, client)
	testutil.Ok(t, err)

	ctx = context.WithValue(ctx, common.ContractsGetterContextKey, instanceGetter)

	proxy, err := db.OpenLocalProxy(dbLocal)
	testutil.Ok(t, err)

	ctx = context.WithValue(ctx, common.DataProxyKey, proxy)

	return ctx, cfg, cleanup
}
