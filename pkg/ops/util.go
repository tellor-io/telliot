// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package ops

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/rpc"
)

func PrepareEthTransaction(ctx context.Context) (*bind.TransactOpts, error) {

	client := ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)

	publicAddress := ctx.Value(tellorCommon.PublicAddress).(common.Address)

	nonce, err := client.PendingNonceAt(ctx, publicAddress)
	if err != nil {
		return nil, fmt.Errorf("problem getting pending nonce: %+v", err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("problem getting gas price: %+v", err)
	}

	ethBalance, err := client.BalanceAt(ctx, publicAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("problem getting balance: %+v", err)
	}

	cost := new(big.Int)
	cost.Mul(gasPrice, big.NewInt(700000))
	if ethBalance.Cmp(cost) < 0 {
		return nil, fmt.Errorf("insufficient ethereum to send a transaction: %v < %v", ethBalance, cost)
	}

	privateKey := ctx.Value(tellorCommon.PrivateKey).(*ecdsa.PrivateKey)
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	return auth, nil
}
