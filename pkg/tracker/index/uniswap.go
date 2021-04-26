// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package index

import (
	"context"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/tellor-io/telliot/pkg/contracts"
	uniswap "github.com/tellor-io/telliot/pkg/contracts/uniswap"
)

// Uniswap implements DataSource interface.
type Uniswap struct {
	symbol0  string
	symbol1  string
	address  string
	client   contracts.ETHClient
	interval time.Duration
}

// NewUniswap creates new Uniswap for provided pair and pair address.
func NewUniswap(pair string, address string, interval time.Duration, client contracts.ETHClient) *Uniswap {
	symbols := strings.Split(pair, "/")
	return &Uniswap{
		interval: interval,
		symbol0:  symbols[0],
		symbol1:  symbols[1],
		address:  address,
		client:   client,
	}
}

// Get calculates price for the provided pair.
func (self *Uniswap) Get(ctx context.Context) (float64, time.Time, error) {
	// Getting price on-chain.
	price, err := self.getSpotPrice(ctx)
	if err != nil {
		return 0, time.Time{}, err
	}
	priceF64, _ := price.Float64()
	return priceF64, time.Time{}, nil
}

func (self *Uniswap) Interval() time.Duration {
	return self.interval
}

func (self *Uniswap) Source() string {
	return self.address
}

func (self *Uniswap) getSpotPrice(ctx context.Context) (*big.Float, error) {
	var pairContract *uniswap.IUniswapV2PairCaller
	var err error
	pairContract, err = uniswap.NewIUniswapV2PairCaller(common.HexToAddress(self.address), self.client)
	if err != nil {
		return nil, errors.Wrap(err, "getting pair contract")
	}

	// Ensure that there's liquidity in the pair.
	reserve, err := pairContract.GetReserves(&bind.CallOpts{Context: ctx})
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
	decimals0, err := self.getTokenDecimals(token0)
	if err != nil {
		return nil, err
	}
	decimals1, err := self.getTokenDecimals(token1)
	if err != nil {
		return nil, err
	}

	// Getting the price side for our calculations.
	side, err := self.getSide(token0, token1)
	if err != nil {
		return nil, err
	}

	// Calculating spot price based on reserve values.
	if side == 0 {
		return calculateSpotPrice(reserve.Reserve0, reserve.Reserve1, decimals0, decimals1), nil
	}
	return calculateSpotPrice(reserve.Reserve1, reserve.Reserve0, decimals1, decimals0), nil
}

func (self *Uniswap) getTokenDecimals(token common.Address) (uint8, error) {
	// Get token decimals.
	// Call on erc20 contracts.
	var erc20TokenCaller *uniswap.IERC20Caller
	erc20TokenCaller, err := uniswap.NewIERC20Caller(token, self.client)
	if err != nil {
		return 0, errors.Wrapf(err, "getting token(%s) contract", token.Hex())
	}
	decimals, err := erc20TokenCaller.Decimals(&bind.CallOpts{})
	if err != nil {
		return 0, errors.Wrapf(err, "getting token(%s) decimals", token.Hex())
	}
	return decimals, nil
}

func (self *Uniswap) getSide(token0, token1 common.Address) (int, error) {
	// Get price side.
	symbol0, err := self.getTokenSymbol(token0)
	if err != nil {
		return -1, err
	}
	symbol1, err := self.getTokenSymbol(token1)
	if err != nil {
		return -1, err
	}
	if symbol0 == self.symbol0 {
		// Maybe this check isn't necessary!
		if symbol1 == self.symbol1 {
			return 0, nil
		}

	} else if symbol0 == self.symbol1 { // Maybe this check isn't necessary too and could be ignored to reduce on-chain calls!

		if symbol1 == self.symbol0 {
			return 1, nil
		}
	}
	return -1, errors.New("wrong pair of input symbols were provided")
}

func (self *Uniswap) getTokenSymbol(token common.Address) (string, error) {
	// Get token symbol.
	// Call on erc20 contracts.
	var erc20TokenCaller *uniswap.IERC20Caller
	erc20TokenCaller, err := uniswap.NewIERC20Caller(token, self.client)
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
