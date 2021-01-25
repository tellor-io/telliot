// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"encoding/json"
	"math"
	"math/big"
	"strings"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/tellor-io/telliot/pkg/contracts"
	uniswap "github.com/tellor-io/telliot/pkg/contracts/uniswap"
)

// Uniswap implements DataSource interface.
type Uniswap struct {
	symbol0 string
	symbol1 string
	address string
	client  contracts.ETHClient
}

func (u *Uniswap) String() string {
	return "Uniswap"
}

// NewUniswap creates new Uniswap for provided pair and pair address.
func NewUniswap(pair string, address string, client contracts.ETHClient) *Uniswap {
	symbols := strings.Split(pair, "/")
	return &Uniswap{
		symbol0: symbols[0],
		symbol1: symbols[1],
		address: address,
		client:  client,
	}
}

// Get calculates price for the provided pair.
func (u *Uniswap) Get() ([]byte, error) {
	// Getting price on-chain.
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

	// Ensure that there's liquidity in the pair.
	reserve, err := pairContract.GetReserves(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "getting reserves")
	}
	if reserve.Reserve0.Uint64() == 0 || reserve.Reserve1.Uint64() == 0 {
		return nil, errors.New("there's no liquidity in the pair")
	}

	// Getting tokens addresses.
	token0, err := pairContract.Token0(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "getting token0")
	}
	token1, err := pairContract.Token1(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "getting token1")
	}

	// Getting token decimals
	decimals0, err := u.getTokenDecimals(token0)
	if err != nil {
		return nil, err
	}
	decimals1, err := u.getTokenDecimals(token1)
	if err != nil {
		return nil, err
	}

	// Getting the price side for our calculations.
	side, err := u.getSide(token0, token1)
	if err != nil {
		return nil, err
	}

	// Calculating spot price based on reserve values.
	if side == 0 {
		return calculateSpotPrice(reserve.Reserve0, reserve.Reserve1, decimals0, decimals1), nil
	}
	return calculateSpotPrice(reserve.Reserve1, reserve.Reserve0, decimals1, decimals0), nil
}

func (u *Uniswap) getTokenDecimals(token common.Address) (uint8, error) {
	// Get token decimals.
	// Call on erc20 contracts.
	var erc20TokenCaller *uniswap.IERC20Caller
	erc20TokenCaller, err := uniswap.NewIERC20Caller(token, u.client)
	if err != nil {
		return 0, errors.Wrapf(err, "getting token(%s) contract", token.Hex())
	}
	decimals, err := erc20TokenCaller.Decimals(&bind.CallOpts{})
	if err != nil {
		return 0, errors.Wrapf(err, "getting token(%s) decimals", token.Hex())
	}
	return decimals, nil
}

func (u *Uniswap) getSide(token0, token1 common.Address) (int, error) {
	// Get price side.
	symbol0, err := u.getTokenSymbol(token0)
	if err != nil {
		return -1, err
	}
	symbol1, err := u.getTokenSymbol(token1)
	if err != nil {
		return -1, err
	}
	if symbol0 == u.symbol0 {
		// Maybe this check isn't necessary!
		if symbol1 == u.symbol1 {
			return 0, nil
		}

	} else if symbol0 == u.symbol1 { // Maybe this check isn't necessary too and could be ignored to reduce on-chain calls!

		if symbol1 == u.symbol0 {
			return 1, nil
		}
	}
	return -1, errors.New("wrong pair of input symbols were provided")
}

func (u *Uniswap) getTokenSymbol(token common.Address) (string, error) {
	// Get token symbol.
	// Call on erc20 contracts.
	var erc20TokenCaller *uniswap.IERC20Caller
	erc20TokenCaller, err := uniswap.NewIERC20Caller(token, u.client)
	if err != nil {
		return "", errors.Wrapf(err, "getting token(%s) contract", token.Hex())
	}
	symbol, err := erc20TokenCaller.Symbol(&bind.CallOpts{})
	if err != nil {
		return "", errors.Wrapf(err, "getting token(%s) symbol", token.Hex())
	}
	if symbol == "WETH" {
		symbol = "ETH"
	}
	return symbol, nil
}

// calculateSpotPrice calculates spot price from uniswap pair reserve values.
func calculateSpotPrice(reserve0, reserve1 *big.Int, decimals0, decimals1 uint8) *big.Float {
	price0 := new(big.Float).Quo(big.NewFloat(0).SetInt(reserve1), big.NewFloat(0).SetInt(reserve0))
	return new(big.Float).Quo(price0, new(big.Float).Quo(new(big.Float).SetFloat64(math.Pow10(int(decimals1))),
		new(big.Float).SetFloat64(math.Pow10(int(decimals0)))))
}
