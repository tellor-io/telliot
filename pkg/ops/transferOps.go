// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package ops

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/contracts"
	proxy "github.com/tellor-io/telliot/pkg/contracts/tellorProxy"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/util"
)

/**
 * This is the operational transfer component. Its purpose is to transfer tellor tokens
 */

func prepareTransfer(
	ctx context.Context,
	client contracts.ETHClient,
	instance *proxy.TellorGetters,
	account *rpc.Account,
	amt *big.Int,
) (*bind.TransactOpts, error) {
	balance, err := instance.BalanceOf(nil, account.Address)
	if err != nil {
		return nil, errors.Wrap(err, "get balance")
	}
	fmt.Println("My balance", util.FormatERC20Balance(balance))
	if balance.Cmp(amt) < 0 {
		return nil, errors.Errorf("insufficient balance TRB actual: %v, requested: %v",
			util.FormatERC20Balance(balance),
			util.FormatERC20Balance(amt))
	}
	auth, err := PrepareEthTransaction(ctx, client, account)
	if err != nil {
		return nil, errors.Wrap(err, "preparing ethereum transaction")
	}
	return auth, nil
}

func Transfer(
	ctx context.Context,
	logger log.Logger,
	client contracts.ETHClient,
	contract *contracts.Tellor,
	account *rpc.Account,
	toAddress common.Address,
	amt *big.Int,
) error {
	auth, err := prepareTransfer(ctx, client, contract.Getter, account, amt)
	if err != nil {
		return errors.Wrap(err, "preparing transfer")
	}

	tx, err := contract.Caller.Transfer(auth, toAddress, amt)
	if err != nil {
		return errors.Wrap(err, "calling transfer")
	}
	level.Info(logger).Log("msg", "transferred", "amount", util.FormatERC20Balance(amt), "to", toAddress.String()[:12], "tx Hash", tx.Hash().Hex())
	return nil
}

func Approve(
	ctx context.Context,
	logger log.Logger,
	client contracts.ETHClient,
	contract *contracts.Tellor,
	account *rpc.Account,
	spender common.Address,
	amt *big.Int,
) error {
	auth, err := prepareTransfer(ctx, client, contract.Getter, account, amt)
	if err != nil {
		return errors.Wrap(err, "preparing transfer")
	}

	tx, err := contract.Caller.Approve(auth, spender, amt)
	if err != nil {
		return errors.Wrap(err, "calling approve")
	}
	level.Info(logger).Log("msg", "approved", "amount", util.FormatERC20Balance(amt), "spender", spender.String()[:12], "tx Hash", tx.Hash().Hex())
	return nil
}

func Balance(ctx context.Context, client contracts.ETHClient, getterInstance *proxy.TellorGetters, addr common.Address) error {
	ethBalance, err := client.BalanceAt(ctx, addr, nil)
	if err != nil {
		return errors.Wrap(err, "get eth balance")
	}
	trbBalance, err := getterInstance.BalanceOf(nil, addr)
	if err != nil {
		return errors.Wrapf(err, "getting trb balance")
	}
	fmt.Printf("%s\n", addr.String())
	fmt.Printf("%10s ETH\n", util.FormatERC20Balance(ethBalance))
	fmt.Printf("%10s TRB\n", util.FormatERC20Balance(trbBalance))
	return nil
}
