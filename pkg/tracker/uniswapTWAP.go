// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"encoding/json"
	"math"
	"math/big"
	"strings"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	tellorCommon "github.com/tellor-io/telliot/pkg/common"
	uniswap "github.com/tellor-io/telliot/pkg/contracts/uniswap"
	"github.com/tellor-io/telliot/pkg/rpc"
)

// Uniswap implements DataSource interface.
type Uniswap struct {
	pair    string
	address string
	client  rpc.ETHClient
}

func (u *Uniswap) String() string {
	return "Uniswap"
}

// NewUniswap creates new Uniswap for provided pair and pair address.
func NewUniswap(pair string, address string) *Uniswap {
	pairAddress := strings.Split(address, ":")[1]
	return &Uniswap{
		pair:    pair,
		address: pairAddress,
	}
}

// Get calculates price for the provided pair.
func (u *Uniswap) Get(ctx context.Context) ([]byte, error) {
	//cast client using type assertion since context holds generic interface{}
	u.client = ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)
	// getting price cumulative last onchain.
	price, err := u.getSpotPrice()
	if err != nil {
		return nil, err
	}
	priceF64, _ := price.Float64()
	return json.Marshal([]float64{priceF64})
}

func (u *Uniswap) getSpotPrice() (*big.Float, error) {
	var pairContract *uniswap.IUniswapV2PairCaller
	var err error
	pairContract, err = uniswap.NewIUniswapV2PairCaller(common.HexToAddress(u.address), u.client)
	if err != nil {
		return nil, errors.Wrap(err, "getting pair contract")
	}
	// ensure that there's liquidity in the pair.
	reserve, err := pairContract.GetReserves(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "getting reserves")
	}
	if reserve.Reserve0.Uint64() == 0 || reserve.Reserve1.Uint64() == 0 {
		return nil, errors.New("there's no liquidity in the pair")
	}
	decimals, err := pairContract.Decimals(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "getting decimals")
	}

	// Get token decimals.
	// Getting tokens addresses.
	token0, err := pairContract.Token0(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "getting token0")
	}
	token1, err := pairContract.Token1(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "getting token0")
	}

	// Call on erc20 contracts.
	var erc20TokenCaller *uniswap.IERC20Caller
	erc20TokenCaller, err = uniswap.NewIERC20Caller(token0, u.client)
	if err != nil {
		return nil, errors.Wrap(err, "getting token0 contract")
	}
	decimals0, err := erc20TokenCaller.Decimals(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "getting token0 decimals")
	}
	erc20TokenCaller, err = uniswap.NewIERC20Caller(token1, u.client)
	if err != nil {
		return nil, errors.Wrap(err, "getting token1 contract")
	}
	decimals1, err := erc20TokenCaller.Decimals(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "getting token1 decimals")
	}

	// Calculating spot price based on reserve values.
	return calculateSpotPrice(reserve.Reserve0, reserve.Reserve1, decimals0, decimals1, decimals), nil
}

// calculateSpotPrice calculates spot price from uniswap pair reserve values.
func calculateSpotPrice(reserve0, reserve1 *big.Int, decimals0, decimals1, decimalsPair uint8) *big.Float {
	price0 := new(big.Float).Quo(big.NewFloat(0).SetInt(reserve0), big.NewFloat(0).SetInt(reserve1))
	price0 = new(big.Float).Quo(price0, new(big.Float).SetFloat64(math.Pow10(int(decimalsPair))))
	price0 = new(big.Float).Quo(price0, new(big.Float).Quo(new(big.Float).SetFloat64(math.Pow10(int(decimals0))),
		new(big.Float).SetFloat64(math.Pow10(int(decimals1)))))
	return price0
}
