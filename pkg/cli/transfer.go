// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/format"
)

func prepareTransfer(
	ctx context.Context,
	logger log.Logger,
	client *ethclient.Client,
	tellor *contracts.ITellor,
	account *ethereum.Account,
	amt *big.Int,
	_gasPrice int,
) (*bind.TransactOpts, error) {
	balance, err := tellor.BalanceOf(nil, account.Address)
	if err != nil {
		return nil, errors.Wrap(err, "get balance")
	}
	level.Info(logger).Log("msg", "check my balance", format.ERC20Balance(balance))
	if balance.Cmp(amt) < 0 {
		return nil, errors.Errorf("insufficient balance TRB actual: %v, requested: %v",
			format.ERC20Balance(balance),
			format.ERC20Balance(amt))
	}

	var gasPrice *big.Int
	if _gasPrice > 0 {
		gasPrice = big.NewInt(int64(_gasPrice) * params.GWei)
	}

	auth, err := ethereum.PrepareEthTransaction(ctx, client, account, gasPrice)
	if err != nil {
		return nil, errors.Wrap(err, "preparing ethereum transaction")
	}
	return auth, nil
}

func Transfer(
	ctx context.Context,
	logger log.Logger,
	client *ethclient.Client,
	tellor *contracts.ITellor,
	account *ethereum.Account,
	toAddress common.Address,
	amt *big.Int,
	gasPrice int,
) error {
	auth, err := prepareTransfer(ctx, logger, client, tellor, account, amt, gasPrice)
	if err != nil {
		return errors.Wrap(err, "preparing transfer")
	}

	tx, err := tellor.Transfer(auth, toAddress, amt)
	if err != nil {
		return errors.Wrap(err, "calling transfer")
	}
	level.Info(logger).Log(
		"msg", "transferred",
		"amount", format.ERC20Balance(amt),
		"to", toAddress.String()[:12],
		"tx Hash", tx.Hash().Hex(),
	)
	return nil
}

func Approve(
	ctx context.Context,
	logger log.Logger,
	client *ethclient.Client,
	tellor *contracts.ITellor,
	account *ethereum.Account,
	spender common.Address,
	amt *big.Int,
	gasPrice int,
) error {
	auth, err := prepareTransfer(ctx, logger, client, tellor, account, amt, gasPrice)
	if err != nil {
		return errors.Wrap(err, "preparing transfer")
	}

	tx, err := tellor.Approve(auth, spender, amt)
	if err != nil {
		return errors.Wrap(err, "calling approve")
	}
	level.Info(logger).Log("msg", "approved", "amount", format.ERC20Balance(amt), "spender", spender.String()[:12], "tx Hash", tx.Hash().Hex())
	return nil
}

func Balance(ctx context.Context, logger log.Logger, client *ethclient.Client, tellor *contracts.ITellor,
	addr common.Address) error {
	ethBalance, err := client.BalanceAt(ctx, addr, nil)
	if err != nil {
		return errors.Wrap(err, "get eth balance")
	}
	trbBalance, err := tellor.BalanceOf(nil, addr)
	if err != nil {
		return errors.Wrapf(err, "getting trb balance")
	}
	level.Info(logger).Log(
		"msg", "balance check",
		"address", addr.String(),
		"ETH", fmt.Sprintf("%10s", format.ERC20Balance(ethBalance)),
		"TRB", fmt.Sprintf("%10s", format.ERC20Balance(trbBalance)),
	)
	return nil
}
