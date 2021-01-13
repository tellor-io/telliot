// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	tellorCommon "github.com/tellor-io/telliot/pkg/common"
	uniswap "github.com/tellor-io/telliot/pkg/contracts/uniswap"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
)

// UniswapPriceCumulative will be used to save/load data into/from the DB, so that
// we can minimize onchain contract calls.
type UniswapPriceCumulative struct {
	PriceCumulative string
	Decimals        uint8
	Decimals0       uint8
	Decimals1       uint8
	Timestamp       uint32
	PriceTWAP       float64
}

// Uniswap implements DataSource interface.
type Uniswap struct {
	pair    string
	address string
	client  rpc.ETHClient
}

func (u *Uniswap) String() string {
	return "Uniswap"
}

func (u *Uniswap) dBCumulativePriceKey() string {
	return fmt.Sprintf("uniswap-%s-cumulative", u.pair)
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
	var err error
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)

	// getting price cumulative last onchain.
	price, err := u.getCumulativePrice()
	if err != nil {
		return nil, errors.Wrap(err, "getting price cumulative last")
	}
	// Try to get previous price0Cumulative, timestamp from DB so that we minimize contract calls.
	var priceCache *UniswapPriceCumulative
	priceCache, err = u.pullPriceFromDB(DB)
	if err != nil {
		return nil, errors.Wrap(err, "getting price cumulative from DB")
	}
	// Return the price data based on conditions.
	var priceData []float64
	if priceCache == nil {
		priceData = []float64{0}
	} else if price.Timestamp == priceCache.Timestamp {
		priceData = []float64{priceCache.PriceTWAP}
	} else {
		// Calculate uniswap TWAP using cumulative prices.
		// https://uniswap.org/docs/v2/core-concepts/oracles
		price.PriceTWAP, err = calculateTWAP(priceCache, price)
		if err != nil {
			return nil, errors.Wrap(err, "calculate TWAP")
		}
		priceData = []float64{price.PriceTWAP}
	}

	// Save cumulative price to the DB.
	err = u.savePriceIntoDB(DB, price)
	if err != nil {
		return nil, err
	}
	return json.Marshal(priceData)
}

func (u *Uniswap) pullPriceFromDB(DB db.DB) (*UniswapPriceCumulative, error) {
	_prevPriceData, err := DB.Get(u.dBCumulativePriceKey())
	if err != nil {
		return nil, errors.Wrap(err, "loading cumulative price from DB")
	}

	// Key doesn't exist.
	if _prevPriceData == nil {
		return nil, nil
	}

	prevPrice := &UniswapPriceCumulative{}
	prevPriceJSON := hexutil.MustDecode(string(_prevPriceData))
	err = json.Unmarshal(prevPriceJSON, prevPrice)
	if err != nil {
		return nil, errors.Wrap(err, "parsing previous cumulative price")
	}
	return prevPrice, nil
}

func (u *Uniswap) savePriceIntoDB(DB db.DB, in *UniswapPriceCumulative) error {
	priceJSON, err := json.Marshal(in)
	if err != nil {
		return errors.Wrap(err, "parsing UniswapPriceCumulative to JSON")
	}
	priceHex := hexutil.Encode(priceJSON)
	err = DB.Put(u.dBCumulativePriceKey(), []byte(priceHex))
	if err != nil {
		return errors.Wrap(err, "saving price into DB")
	}
	return nil
}

// calculateTWAP calculates Time Weighted Average Price from pair of UniswapPriceCumulative.
func calculateTWAP(old, last *UniswapPriceCumulative) (price0 float64, err error) {
	priceCumulativeLast, ok := new(big.Float).SetString(last.PriceCumulative)
	if !ok {
		err = errors.Wrap(err, "parsing last.PriceCumulative to big.Float")
		return
	}
	priceCumulativeOld, ok := new(big.Float).SetString(old.PriceCumulative)
	if !ok {
		err = errors.Wrap(err, "parsing old.PriceCumulative to big.Float")
		return
	}

	timeElapsed := new(big.Float).SetUint64(uint64(last.Timestamp - old.Timestamp))
	_price0 := new(big.Float).Quo(new(big.Float).Sub(priceCumulativeLast, priceCumulativeOld), timeElapsed)
	// Calculate float value of the price.
	_price0 = new(big.Float).Quo(_price0, new(big.Float).SetFloat64(math.Pow10(int(last.Decimals))))
	price0, _ = new(big.Float).Quo(_price0, new(big.Float).Quo(new(big.Float).SetFloat64(math.Pow10(int(last.Decimals0))),
		new(big.Float).SetFloat64(math.Pow10(int(last.Decimals1))))).Float64()
	return

}
func (u *Uniswap) getCumulativePrice() (priceLast *UniswapPriceCumulative, err error) {
	var pairContract *uniswap.IUniswapV2PairCaller
	pairContract, err = uniswap.NewIUniswapV2PairCaller(common.HexToAddress(u.address), u.client)
	if err != nil {
		err = errors.Wrap(err, "getting pair contract")
		return
	}
	// ensure that there's liquidity in the pair.
	reserve, err := pairContract.GetReserves(&bind.CallOpts{})
	if err != nil {
		err = errors.Wrap(err, "getting reserves")
		return
	}
	if reserve.Reserve0.Uint64() == 0 || reserve.Reserve1.Uint64() == 0 {
		err = errors.New("there's no liquidity in the pair")
		return
	}
	// Getting price information.
	price0CumulativeLast, err := pairContract.Price0CumulativeLast(&bind.CallOpts{})
	if err != nil {
		err = errors.Wrap(err, "getting price cumulative last")
		return
	}
	decimals, err := pairContract.Decimals(&bind.CallOpts{})
	if err != nil {
		err = errors.Wrap(err, "getting decimals")
		return
	}

	// Get token decimals.

	// Getting tokens addresses.
	token0, err := pairContract.Token0(&bind.CallOpts{})
	if err != nil {
		err = errors.Wrap(err, "getting token0")
		return
	}
	token1, err := pairContract.Token1(&bind.CallOpts{})
	if err != nil {
		err = errors.Wrap(err, "getting token0")
		return
	}

	// Call on erc20 contracts.
	var erc20TokenCaller *uniswap.IERC20Caller
	erc20TokenCaller, err = uniswap.NewIERC20Caller(token0, u.client)
	if err != nil {
		err = errors.Wrap(err, "getting token0 contract")
		return
	}
	decimals0, err := erc20TokenCaller.Decimals(&bind.CallOpts{})
	if err != nil {
		err = errors.Wrap(err, "getting token0 decimals")
		return
	}
	erc20TokenCaller, err = uniswap.NewIERC20Caller(token1, u.client)
	if err != nil {
		err = errors.Wrap(err, "getting token1 contract")
		return
	}
	decimals1, err := erc20TokenCaller.Decimals(&bind.CallOpts{})
	if err != nil {
		err = errors.Wrap(err, "getting token1 decimals")
		return
	}

	priceLast = &UniswapPriceCumulative{
		PriceCumulative: price0CumulativeLast.String(),
		Decimals:        decimals,
		Decimals0:       decimals0,
		Decimals1:       decimals1,
		Timestamp:       reserve.BlockTimestampLast,
	}
	return
}
