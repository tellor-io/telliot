// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

// Copyright (c) The Tellor Authors.
// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/tellor-io/telliot/pkg/common"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/testutil"
	"github.com/tellor-io/telliot/pkg/util"
)

func TestBalancerPrice(t *testing.T) {
	logSetup := util.SetupLogger()
	logSetup("debug")
	bPoolContract := "0x0000000000000000000000000000000000000001"
	opts := &rpc.MockOptions{
		BPoolContractAddress: eth_common.HexToAddress(bPoolContract),
		BPoolCurrentTokens: []eth_common.Address{
			eth_common.HexToAddress("0x0000000000000000000000000000000000000002"),
			eth_common.HexToAddress("0x0000000000000000000000000000000000000003"),
		},
		BPoolSpotPrice: new(big.Int).SetInt64(2000000000),
		BTokenSymbols: map[string]string{
			"0x0000000000000000000000000000000000000002": "T1",
			"0x0000000000000000000000000000000000000003": "T2",
		},
		Decimals: map[string]int{
			"0x0000000000000000000000000000000000000001": 3,
			"0x0000000000000000000000000000000000000002": 6,
			"0x0000000000000000000000000000000000000003": 3,
		},
	}
	client := rpc.NewMockClientWithValues(opts)

	tracker := NewBalancerGetter("T1/T2", fmt.Sprintf("ethereum:%s", bPoolContract))
	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	priceJSON, err := tracker.Get(ctx)
	testutil.Ok(t, err)

	var priceInfo []float64
	err = json.Unmarshal(priceJSON, &priceInfo)
	testutil.Ok(t, err)
	testutil.Equals(t, []float64{2000.0}, priceInfo)
	t.Logf("AMPL/USD price on Balancer: %v\n", priceInfo)

}
