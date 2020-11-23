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
	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/contracts/getter"
	"github.com/tellor-io/TellorMiner/pkg/rpc"
	"github.com/tellor-io/TellorMiner/pkg/util"
)

/**
 * This is the operational transfer component. Its purpose is to transfer tellor tokens
 */

func prepareTransfer(ctx context.Context, client rpc.ETHClient, instance *getter.TellorGetters, account tellorCommon.Account, amt *big.Int) (*bind.TransactOpts, error) {
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

func Transfer(ctx context.Context, logger log.Logger, client rpc.ETHClient, contract tellorCommon.Contract, account tellorCommon.Account, toAddress common.Address, amt *big.Int) error {
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

func Approve(ctx context.Context, logger log.Logger, client rpc.ETHClient, contract tellorCommon.Contract, account tellorCommon.Account, _spender common.Address, amt *big.Int) error {
	auth, err := prepareTransfer(ctx, client, contract.Getter, account, amt)
	if err != nil {
		return errors.Wrap(err, "preparing transfer")
	}

	tx, err := contract.Caller.Approve(auth, _spender, amt)
	if err != nil {
		return errors.Wrap(err, "calling approve")
	}
	level.Info(logger).Log("msg", "approved", "amount", util.FormatERC20Balance(amt), "spender", _spender.String()[:12], "tx Hash", tx.Hash().Hex())
	return nil
}

func Balance(ctx context.Context, client rpc.ETHClient, getterInstance *getter.TellorGetters, addr common.Address) error {
	ethBalance, err := client.BalanceAt(context.Background(), addr, nil)
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
