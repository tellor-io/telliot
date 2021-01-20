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

func TestBalancerPrice(t *testing.T) {
	bPoolContract := eth_common.HexToAddress("0x7860E28EBFB8AE052BFE279C07AC5D94C9CD2937")
	token1Address := eth_common.HexToAddress("0xA0B86991C6218B36C1D19D4A2E9EB0CE3606EB48")
	token2Address := eth_common.HexToAddress("0xD46BA6D942050D489DBD938A2C909A5D5039A161")
	spotPrice, _ := new(big.Int).SetString("1113832303486165407237", 10)
	opts := &rpc.MockOptions{
		BPoolContractAddress: bPoolContract,
		BPoolCurrentTokens: []eth_common.Address{
			token1Address,
			token2Address,
		},
		BPoolSpotPrice: spotPrice,
		TokenSymbols: map[string]string{
			token1Address.Hex(): "USDC",
			token2Address.Hex(): "AMPL",
		},
		Decimals: map[string]int{
			bPoolContract.Hex(): 18,
			token1Address.Hex(): 6,
			token2Address.Hex(): 9,
		},
	}
	client := rpc.NewMockClientWithValues(opts)

	tracker := NewBalancer("USDC/AMPL", bPoolContract.Hex(), client)
	priceJSON, err := tracker.Get()
	testutil.Ok(t, err)

	var priceInfo []float64
	err = json.Unmarshal(priceJSON, &priceInfo)
	testutil.Ok(t, err)
	testutil.Equals(t, []float64{1.1138323034861655}, priceInfo)
	t.Logf("AMPL/USD price on Balancer: %v\n", priceInfo)

}
