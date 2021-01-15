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

func TestUniswapPrice(t *testing.T) {
	logSetup := util.SetupLogger()
	logSetup("debug")
	bPairContract := "0x0000000000000000000000000000000000000001"
	opts := &rpc.MockOptions{
		UniPairContractAddress:  eth_common.HexToAddress(bPairContract),
		UniPrice0CumulativeLast: new(big.Int).SetInt64(2000000000),
		UniToken0:               eth_common.HexToAddress("0x0000000000000000000000000000000000000002"),
		UniToken1:               eth_common.HexToAddress("0x0000000000000000000000000000000000000003"),
		UniReserves: &rpc.CurrentReserves{
			Reserve0:           big.NewInt(1328),
			Reserve1:           big.NewInt(10),
			BlockTimestampLast: 200,
		},
		Decimals: map[string]int{
			"0x0000000000000000000000000000000000000001": 3,
			"0x0000000000000000000000000000000000000002": 6,
			"0x0000000000000000000000000000000000000003": 3,
		},
	}
	client := rpc.NewMockClientWithValues(opts)

	tracker := NewUniswap("T1/T2", fmt.Sprintf("ethereum:%s", bPairContract))
	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	priceJSON, err := tracker.Get(ctx)
	testutil.Ok(t, err)

	var priceInfo []float64
	err = json.Unmarshal(priceJSON, &priceInfo)
	testutil.Ok(t, err)
	// Calculated according to this:
	// https://uniswap.org/docs/v2/core-concepts/oracles
	testutil.Equals(t, []float64{132.8}, priceInfo)
	t.Logf("AMPL/ETH price on Uniswap: %f\n", priceInfo[0])
}
