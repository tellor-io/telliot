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
	tellor "github.com/tellor-io/TellorMiner/abi/contracts"
	tellor1 "github.com/tellor-io/TellorMiner/abi/contracts1"
	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/rpc"
	"github.com/tellor-io/TellorMiner/pkg/util"
)

/**
 * This is the operational transfer component. Its purpose is to transfer tellor tokens
 */

func prepareTransfer(amt *big.Int, ctx context.Context) (*bind.TransactOpts, error) {
	instance := ctx.Value(tellorCommon.MasterContractContextKey).(*tellor.TellorMaster)
	senderPubAddr := ctx.Value(tellorCommon.PublicAddress).(common.Address)

	balance, err := instance.BalanceOf(nil, senderPubAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %v", err)
	}
	fmt.Println("My balance", util.FormatERC20Balance(balance))
	if balance.Cmp(amt) < 0 {
		return nil, fmt.Errorf("insufficient balance (%s TRB), requested %s TRB",
			util.FormatERC20Balance(balance),
			util.FormatERC20Balance(amt))
	}
	auth, err := PrepareEthTransaction(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare ethereum transaction: %s", err.Error())
	}
	return auth, nil
}

func Transfer(ctx context.Context, logger log.Logger, toAddress common.Address, amt *big.Int) error {
	auth, err := prepareTransfer(amt, ctx)
	if err != nil {
		return err
	}

	instance2 := ctx.Value(tellorCommon.TransactorContractContextKey).(*tellor1.TellorTransactor)
	tx, err := instance2.Transfer(auth, toAddress, amt)
	if err != nil {
		return fmt.Errorf("contract failed: %s", err.Error())
	}
	level.Info(logger).Log("msg", "Transferred", "amount", util.FormatERC20Balance(amt), "to", toAddress.String()[:12], "tx Hash", tx.Hash().Hex())
	return nil
}

func Approve(ctx context.Context, logger log.Logger, _spender common.Address, amt *big.Int) error {
	auth, err := prepareTransfer(amt, ctx)
	if err != nil {
		return err
	}

	instance2 := ctx.Value(tellorCommon.TransactorContractContextKey).(*tellor1.TellorTransactor)
	tx, err := instance2.Approve(auth, _spender, amt)
	if err != nil {
		return err
	}
	level.Info(logger).Log("msg", "Approved", "amount", util.FormatERC20Balance(amt), "spender", _spender.String()[:12], "tx Hash", tx.Hash().Hex())
	return nil
}

func Balance(ctx context.Context, addr common.Address) error {
	client := ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)

	ethBalance, err := client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return fmt.Errorf("problem getting balance: %+v", err)
	}

	instance := ctx.Value(tellorCommon.MasterContractContextKey).(*tellor.TellorMaster)
	trbBalance, err := instance.BalanceOf(nil, addr)
	if err != nil {
		return fmt.Errorf("problem getting balance: %+v", err)
	}
	fmt.Printf("%s\n", addr.String())
	fmt.Printf("%10s ETH\n", util.FormatERC20Balance(ethBalance))
	fmt.Printf("%10s TRB\n", util.FormatERC20Balance(trbBalance))
	return nil
}
