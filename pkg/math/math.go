// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package math

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
)

func PercentageDiff(old, new float64) (delta float64) {
	diff := float64(new - old)

	if old > new {
		return (diff / float64(old)) * 100
	}
	return (diff / float64(new)) * 100
}

func FloatToBigInt18e(v float64) (*big.Int, error) {
	v = v * params.Ether
	g := new(big.Int)

	_, ok := g.SetString(strconv.FormatFloat(v, 'f', -1, 64), 10)
	if !ok {
		return nil, errors.Errorf("invalid float:%v", v)
	}
	if len(g.Bytes()) > 32 {
		return nil, errors.Errorf("invalid size larger than 256 bits:%v", v)
	}
	return g, nil
}

func BigInt18eToFloat(input *big.Int) float64 {
	f := 0.0
	if input != nil {
		divisor := big.NewInt(1e18 / 100)
		divisor.Div(input, divisor)
		f = float64(divisor.Uint64()) / 100
	}
	return f
}
