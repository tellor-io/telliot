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
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
	uniswapcontract "github.com/tellor-io/telliot/pkg/tracker/uniswap"
)

// UniswapPriceCumulativeLast will be used to save/load data into/from the DB, so that
// we can minimize onchain contract calls.
type UniswapPriceCumulativeLast struct {
	Price0CumulativeLast string
	Decimals             uint8
	Decimals0            uint8
	Decimals1            uint8
	TimestampLast        uint32
	CalculatedPrice0Last float64
}

// UniswapGetter implements DataSource interface.
type UniswapGetter struct {
	pair    string
	address string
	client  rpc.ETHClient
}

func (u *UniswapGetter) String() string {
	return "UniswapGetter"
}

func (u *UniswapGetter) dBCumulativePriceKey() string {
	return fmt.Sprintf("uniswap-%s-cumulative", u.pair)
}

// NewUniswapGetter creates new UniswapGetter for provided pair and pair address.
func NewUniswapGetter(pair string, address string) *UniswapGetter {
	pairAddress := strings.Split(address, ":")[1]
	return &UniswapGetter{
		pair:    pair,
		address: pairAddress,
	}
}

// Get calculates price for the provided pair.
func (u *UniswapGetter) Get(ctx context.Context) ([]byte, error) {
	//cast client using type assertion since context holds generic interface{}
	u.client = ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)
	var err error
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)

	// getting price cumulative last onchain.
	priceCumulativeLast, err := u.getCumulativePriceLast()
	if err != nil {
		return nil, errors.Wrap(err, "getting price cumulative last")
	}
	// Try to get previous price0Cumulative, timestamp from DB so that we minimize contract calls.
	var priceCumulative *UniswapPriceCumulativeLast
	priceCumulative, err = u.pullPriceFromDB(DB)
	if err != nil {
		return nil, errors.Wrap(err, "getting price cumulative from DB")
	}
	// Return the price data based on conditions.
	var priceData []float64
	if priceCumulative == nil {
		priceData = []float64{0}

	} else if priceCumulative.Price0CumulativeLast == priceCumulativeLast.Price0CumulativeLast {
		priceData = []float64{priceCumulative.CalculatedPrice0Last}
	} else {
		// Calculate uniswap TWAP using cumulative prices. see https://uniswap.org/docs/v2/core-concepts/oracles.
		priceCumulativeLast.CalculatedPrice0Last, err = calculateTWAP(priceCumulative, priceCumulativeLast)
		if err != nil {
			return nil, errors.Wrap(err, "calculate TWAP")
		}
		priceData = []float64{priceCumulativeLast.CalculatedPrice0Last}
	}

	// Save cumulative price to the DB.
	err = u.savePriceIntoDB(DB, priceCumulativeLast)
	if err != nil {
		return nil, err
	}
	return json.Marshal(priceData)
}

func (u *UniswapGetter) pullPriceFromDB(DB db.DB) (*UniswapPriceCumulativeLast, error) {
	_prevPriceData, err := DB.Get(u.dBCumulativePriceKey())
	if err != nil {
		return nil, errors.Wrap(err, "loading cumulative price from DB")
	}

	// Key doesn't exist.
	if _prevPriceData == nil {
		return nil, nil
	}

	prevPrice := &UniswapPriceCumulativeLast{}
	prevPriceJSON := hexutil.MustDecode(string(_prevPriceData))
	err = json.Unmarshal(prevPriceJSON, prevPrice)
	if err != nil {
		return nil, errors.Wrap(err, "parsing previous cumulative price")
	}
	return prevPrice, nil
}

func (u *UniswapGetter) savePriceIntoDB(DB db.DB, in *UniswapPriceCumulativeLast) error {
	priceJSON, err := json.Marshal(in)
	if err != nil {
		return errors.Wrap(err, "parsing UniswapPriceCumulativeLast to JSON")
	}
	priceHex := hexutil.Encode(priceJSON)
	err = DB.Put(u.dBCumulativePriceKey(), []byte(priceHex))
	if err != nil {
		return errors.Wrap(err, "saving price into DB")
	}
	return nil
}

// calculate Time Weighted Average Price from pair of UniswapPriceCumulativeLast.
func calculateTWAP(old, last *UniswapPriceCumulativeLast) (price0 float64, err error) {
	price0CumulativeLast, ok := new(big.Float).SetString(last.Price0CumulativeLast)
	if !ok {
		err = errors.Wrap(err, "parsing last.Price0CumulativeLast to big.Float")
		return
	}
	price0CumulativeOld, ok := new(big.Float).SetString(old.Price0CumulativeLast)
	if !ok {
		err = errors.Wrap(err, "parsing old.Price0CumulativeLast to big.Float")
		return
	}

	timeElapsed := new(big.Float).SetUint64(uint64(last.TimestampLast - old.TimestampLast))
	_price0 := new(big.Float).Quo(new(big.Float).Sub(price0CumulativeLast, price0CumulativeOld), timeElapsed)
	// Calculate float value of the price.
	_price0 = new(big.Float).Quo(_price0, new(big.Float).SetFloat64(math.Pow10(int(last.Decimals))))
	price0, _ = new(big.Float).Quo(_price0, new(big.Float).Quo(new(big.Float).SetFloat64(math.Pow10(int(last.Decimals0))),
		new(big.Float).SetFloat64(math.Pow10(int(last.Decimals1))))).Float64()
	return

}
func (u *UniswapGetter) getCumulativePriceLast() (priceLast *UniswapPriceCumulativeLast, err error) {
	var pairContract *uniswapcontract.IUniswapV2PairCaller
	pairContract, err = uniswapcontract.NewIUniswapV2PairCaller(common.HexToAddress(u.address), u.client)
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
	var erc20TokenCaller *uniswapcontract.IERC20Caller
	erc20TokenCaller, err = uniswapcontract.NewIERC20Caller(token0, u.client)
	if err != nil {
		err = errors.Wrap(err, "getting token0 contract")
		return
	}
	decimals0, err := erc20TokenCaller.Decimals(&bind.CallOpts{})
	if err != nil {
		err = errors.Wrap(err, "getting token0 decimals")
		return
	}
	erc20TokenCaller, err = uniswapcontract.NewIERC20Caller(token1, u.client)
	if err != nil {
		err = errors.Wrap(err, "getting token1 contract")
		return
	}
	decimals1, err := erc20TokenCaller.Decimals(&bind.CallOpts{})
	if err != nil {
		err = errors.Wrap(err, "getting token1 decimals")
		return
	}

	priceLast = &UniswapPriceCumulativeLast{
		Price0CumulativeLast: price0CumulativeLast.String(),
		Decimals:             decimals,
		Decimals0:            decimals0,
		Decimals1:            decimals1,
		TimestampLast:        reserve.BlockTimestampLast,
	}
	return
}
