// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package ops

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/rpc"
)

func PrepareEthTransaction(
	ctx context.Context,
	client contracts.ETHClient,
	account *rpc.Account,
) (*bind.TransactOpts, error) {

	nonce, err := client.PendingNonceAt(ctx, account.Address)
	if err != nil {
		return nil, errors.Wrap(err, "getting pending nonce")
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "getting gas price")
	}

	ethBalance, err := client.BalanceAt(ctx, account.Address, nil)
	if err != nil {
		return nil, errors.Wrap(err, "getting balance")
	}

	cost := new(big.Int)
	cost.Mul(gasPrice, big.NewInt(700000))
	if ethBalance.Cmp(cost) < 0 {
		return nil, errors.Errorf("insufficient ethereum to send a transaction: %v < %v", ethBalance, cost)
	}

	netID, err := client.NetworkID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "getting network id")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(account.PrivateKey, netID)
	if err != nil {
		return nil, errors.Wrap(err, "creating transactor")
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	return auth, nil
}
