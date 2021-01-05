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
	"github.com/tellor-io/telliot/pkg/rpc"
	pool "github.com/tellor-io/telliot/pkg/tracker/balancer/balancerpool"
	balancerToken "github.com/tellor-io/telliot/pkg/tracker/balancer/balancertoken"
)

// BalancerPair to be fetched onchain.
type BalancerPair struct {
	token1Address  common.Address
	token2Address  common.Address
	token1Decimals uint64
	token2Decimals uint64
}

// BalancerGetter implements DataSource interface.
type BalancerGetter struct {
	address string
	token1  string
	token2  string
	client  rpc.ETHClient
}

func (b *BalancerGetter) String() string {
	return "BalancerGetter"
}

func NewBalancerGetter(pair, address string) *BalancerGetter {
	_address := strings.Split(address, ":")
	tokens := strings.Split(pair, "/")
	return &BalancerGetter{
		address: _address[1],
		token1:  tokens[0],
		token2:  tokens[1],
	}
}

func (b *BalancerGetter) Get(ctx context.Context) ([]byte, error) {
	//cast client using type assertion since context holds generic interface{}.
	b.client = ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)

	// Getting current pair info from input pool
	pair, err := b.getPair()
	if err != nil {
		return nil, errors.Wrap(err, "getting pair info from balancer pool")
	}
	// use balancer pool own GetSpotPrice to minimize onchain calls.
	price, err := b.getSpotPrice(pair)
	if err != nil {
		return nil, errors.Wrap(err, "getting price info from balancer pool")
	}
	// Output to index tracker.
	return json.Marshal([]float64{price})
}

func (b *BalancerGetter) getPair() (pair *BalancerPair, err error) {
	var poolCaller *pool.BalancerpoolCaller
	poolCaller, err = pool.NewBalancerpoolCaller(common.HexToAddress(b.address), b.client)
	if err != nil {
		return
	}
	currentTokens, err := poolCaller.GetCurrentTokens(&bind.CallOpts{})
	if err != nil {
		return
	}

	pair = &BalancerPair{}
	var token1Seen, token2Seen bool
	for _, token := range currentTokens {
		var tokenCaller *balancerToken.BalancertokenCaller
		tokenCaller, err = balancerToken.NewBalancertokenCaller(token, b.client)
		if err != nil {
			return
		}
		var symbol string
		var decimals *big.Int
		symbol, err = tokenCaller.Symbol(&bind.CallOpts{})
		if err != nil {
			return
		}
		decimals, err = tokenCaller.Decimals(&bind.CallOpts{})
		if err != nil {
			return
		}
		if symbol == b.token1 {
			pair.token1Address = token
			pair.token1Decimals = decimals.Uint64()
			token1Seen = true
		} else if symbol == b.token2 {
			pair.token2Address = token
			pair.token2Decimals = decimals.Uint64()
			token2Seen = true
		}
	}
	if !token1Seen || !token2Seen {
		err = errors.New("we expected this pool to have the provided tokens")
		return
	}
	return
}

func (b *BalancerGetter) getSpotPrice(pair *BalancerPair) (float64, error) {
	var poolCaller *pool.BalancerpoolCaller
	poolCaller, err := pool.NewBalancerpoolCaller(common.HexToAddress(b.address), b.client)
	if err != nil {
		return 0, err
	}

	spotPrice, err := poolCaller.GetSpotPrice(&bind.CallOpts{}, pair.token1Address, pair.token2Address)
	if err != nil {
		return 0, err
	}
	decimals, err := poolCaller.Decimals(&bind.CallOpts{})
	if err != nil {
		return 0, err
	}

	_spotPrice := new(big.Float).Quo(big.NewFloat(0).SetInt(spotPrice), new(big.Float).SetFloat64(math.Pow10(int(decimals))))
	price, _ := new(big.Float).Quo(_spotPrice, new(big.Float).Quo(new(big.Float).SetFloat64(math.Pow10(int(pair.token1Decimals))),
		new(big.Float).SetFloat64(math.Pow10(int(pair.token2Decimals))))).Float64()
	return price, nil
}
