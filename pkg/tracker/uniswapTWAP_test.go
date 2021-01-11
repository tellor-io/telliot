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
	"os"
	"path/filepath"
	"testing"

	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/tellor-io/telliot/pkg/common"
	"github.com/tellor-io/telliot/pkg/db"
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
			Reserve0:           big.NewInt(1),
			Reserve1:           big.NewInt(1),
			BlockTimestampLast: 200,
		},
		Decimals: map[string]int{
			"0x0000000000000000000000000000000000000001": 3,
			"0x0000000000000000000000000000000000000002": 6,
			"0x0000000000000000000000000000000000000003": 3,
		},
	}
	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_uniswapTWAP"))
	testutil.Ok(t, err)

	tracker := NewUniswap("T1/T2", fmt.Sprintf("ethereum:%s", bPairContract))
	// Add a test record to db.
	prevPrice := &UniswapPriceCumulative{
		PriceCumulative: "1000000000",
		Timestamp:       100,
		PriceTWAP:       0.0,
	}
	err = tracker.savePriceIntoDB(DB, prevPrice)
	testutil.Ok(t, err)
	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, DB)
	priceJSON, err := tracker.Get(ctx)
	testutil.Ok(t, err)

	var priceInfo []float64
	err = json.Unmarshal(priceJSON, &priceInfo)
	testutil.Ok(t, err)
	// Calculated according to this:
	// https://uniswap.org/docs/v2/core-concepts/oracles
	testutil.Equals(t, []float64{10.0}, priceInfo)
	t.Logf("AMPL/ETH price on Uniswap: %f\n", priceInfo[0])
	// Check saved db values.
	prevPrice, err = tracker.pullPriceFromDB(DB)
	testutil.Ok(t, err)
	testutil.Equals(t, prevPrice.PriceTWAP, 10.0)
	testutil.Equals(t, prevPrice.Timestamp, uint32(200))
	testutil.Equals(t, prevPrice.Decimals, uint8(3))
	testutil.Equals(t, prevPrice.Decimals0, uint8(6))
	testutil.Equals(t, prevPrice.Decimals1, uint8(3))
	testutil.Equals(t, prevPrice.PriceCumulative, "2000000000")
}
