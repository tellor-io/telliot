// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

// Copyright (c) The Tellor Authors.
// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"encoding/json"
	"math/big"
	"testing"

	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/testutil"
)

func TestUniswapPrice(t *testing.T) {
	bPairContract := eth_common.HexToAddress("0xc5be99a02c6857f9eac67bbce58df5572498f40c")
	token1Address := eth_common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	token2Address := eth_common.HexToAddress("0xd46ba6d942050d489dbd938a2c909a5d5039a161")
	reserve0, _ := big.NewInt(0).SetString("11073781494155978314322", 10)
	reserve1, _ := big.NewInt(0).SetString("14899909395042275", 10)
	opts := &rpc.MockOptions{
		UniPairContractAddress: bPairContract,
		UniToken0:              token1Address,
		UniToken1:              token2Address,
		UniReserves: &rpc.CurrentReserves{
			Reserve0:           reserve0,
			Reserve1:           reserve1,
			BlockTimestampLast: 200,
		},
		TokenSymbols: map[string]string{
			token1Address.Hex(): "ETH",
			token2Address.Hex(): "AMPL",
		},
		Decimals: map[string]int{
			bPairContract.Hex(): 18,
			token1Address.Hex(): 18,
			token2Address.Hex(): 9,
		},
	}
	client := rpc.NewMockClientWithValues(opts)

	tracker := NewUniswap("ETH/AMPL", bPairContract.Hex(), client)
	priceJSON, err := tracker.Get()
	testutil.Ok(t, err)

	var priceInfo []float64
	err = json.Unmarshal(priceJSON, &priceInfo)
	testutil.Ok(t, err)
	testutil.Equals(t, []float64{1345.5123168996497}, priceInfo)
	t.Logf("AMPL/ETH price on Uniswap: %.5f\n", priceInfo[0])
}
