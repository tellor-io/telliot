// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package cli

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/format"
)

type TRBAmount struct {
	*big.Int
}

func (a *TRBAmount) Set(v string) error {
	f, _, err := big.ParseFloat(v, 10, 256, 0)
	if err != nil {
		return err
	}
	scale := big.NewFloat(1e18)
	f.Mul(f, scale)
	a.Int, _ = f.Int(nil)
	return nil
}

func (a *TRBAmount) String() string {
	return format.ERC20Balance(a.Int)
}

func (a *TRBAmount) IsDefault() bool {
	return true
}

type ETHAddress struct {
	addr common.Address
}

func (a *ETHAddress) Set(v string) error {
	valid := common.IsHexAddress(v)
	if !valid {
		return errors.Errorf("invalid etherum address:%v", v)
	}
	a.addr = common.HexToAddress(v)
	return nil
}

func (a *ETHAddress) String() string {
	return a.addr.String()
}

func (a *ETHAddress) IsDefault() bool {
	return true
}

type EthereumInt struct {
	*big.Int
}

func (b *EthereumInt) Set(v string) error {
	g := new(big.Int)
	_, ok := g.SetString(v, 10)
	if !ok {
		return errors.Errorf("invalid integer:%v", v)
	}
	if len(g.Bytes()) > 32 {
		return errors.Errorf("invalid size larger than 256 bits:%v", v)
	}
	b.Int = g
	return nil
}

func (b *EthereumInt) String() string {
	return b.Int.Text(10)
}

func (b *EthereumInt) IsDefault() bool {
	return true
}
