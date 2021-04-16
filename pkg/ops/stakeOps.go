// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package ops

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/util"
)

/**
 * This is the operational deposit component. Its purpose is to deposit Tellor Tokens so you can mine
 */

func printStakeStatus(logger log.Logger, bigStatus *big.Int, started *big.Int) {
	// 0-not Staked, 1=Staked, 2=LockedForWithdraw 3= OnDispute
	status := bigStatus.Uint64()
	stakeTime := time.Unix(started.Int64(), 0)
	switch status {
	case 0:
		level.Info(logger).Log("msg", "not currently staked")
	case 1:
		level.Info(logger).Log("msg", "staked in good standing since", "UTC", stakeTime.UTC())
	case 2:
		startedRound := started.Int64()
		startedRound = ((startedRound + 86399) / 86400) * 86400
		target := time.Unix(startedRound, 0)
		timePassed := time.Since(target)
		delta := timePassed - (time.Hour * 24 * 7)
		if delta > 0 {
			level.Info(logger).Log("msg", "stake has been eligbile to withdraw for", "delta", delta)
		} else {
			level.Info(logger).Log("msg", "stake will be eligible to withdraw in", "delta", -delta)
		}
	case 3:
		level.Info(logger).Log("msg", "stake is currently under dispute")
	}
}

func Deposit(
	ctx context.Context,
	logger log.Logger,
	client contracts.ETHClient,
	contract *contracts.ITellor,
	account *ethereum.Account,
) error {

	balance, err := contract.BalanceOf(nil, account.Address)
	if err != nil {
		return errors.Wrap(err, "get TRB balance")
	}

	status, startTime, err := contract.GetStakerInfo(nil, account.Address)
	if err != nil {
		return errors.Wrap(err, "get stake status")
	}

	if status.Uint64() != 0 && status.Uint64() != 2 {
		printStakeStatus(logger, status, startTime)
		return nil
	}

	dat := crypto.Keccak256([]byte("_STAKE_AMOUNT"))
	var dat32 [32]byte
	copy(dat32[:], dat)
	stakeAmt, err := contract.GetUintVar(nil, dat32)
	if err != nil {
		return errors.Wrap(err, "fetching stake amount")
	}

	if balance.Cmp(stakeAmt) < 0 {
		return errors.Errorf("insufficient mining stake TRB balance actual: %v, required:%v",
			util.FormatERC20Balance(balance),
			util.FormatERC20Balance(stakeAmt))
	}

	auth, err := ethereum.PrepareEthTransaction(ctx, client, account)
	if err != nil {
		return errors.Wrap(err, "prepare ethereum transaction")
	}

	tx, err := contract.DepositStake(auth)
	if err != nil {
		return errors.Wrap(err, "contract failed")
	}
	level.Info(logger).Log("msg", "stake depositied", "txHash", tx.Hash().Hex())
	return nil
}

func ShowStatus(
	ctx context.Context,
	logger log.Logger,
	client contracts.ETHClient,
	contract *contracts.ITellor,
	account *ethereum.Account,
) error {
	status, startTime, err := contract.GetStakerInfo(nil, account.Address)
	if err != nil {
		return errors.Wrap(err, "get stake status")
	}

	printStakeStatus(logger, status, startTime)
	return nil
}

func RequestStakingWithdraw(
	ctx context.Context,
	logger log.Logger,
	client contracts.ETHClient,
	contract *contracts.ITellor,
	account *ethereum.Account,
) error {

	status, startTime, err := contract.GetStakerInfo(nil, account.Address)
	if err != nil {
		return errors.Wrap(err, "get stake status")
	}
	if status.Uint64() != 1 {
		printStakeStatus(logger, status, startTime)
		return nil
	}

	auth, err := ethereum.PrepareEthTransaction(ctx, client, account)
	if err != nil {
		return errors.Wrap(err, "prepare ethereum transaction")
	}

	tx, err := contract.RequestStakingWithdraw(auth)
	if err != nil {
		return errors.Wrap(err, "contract")
	}

	level.Info(logger).Log("msg", "withdrawal request sent", "txHash", tx.Hash().Hex())
	return nil
}

func WithdrawStake(
	ctx context.Context,
	logger log.Logger,
	client contracts.ETHClient,
	contract *contracts.ITellor,
	account *ethereum.Account,
) error {
	status, startTime, err := contract.GetStakerInfo(nil, account.Address)
	if err != nil {
		return errors.Wrap(err, "get stake status")
	}
	if status.Uint64() != 2 {
		level.Info(logger).Log("msg", "can't withdraw")
		printStakeStatus(logger, status, startTime)
		return nil
	}

	auth, err := ethereum.PrepareEthTransaction(ctx, client, account)
	if err != nil {
		return errors.Wrap(err, "prepare ethereum transaction")
	}

	tx, err := contract.WithdrawStake(auth)
	if err != nil {
		return errors.Wrap(err, "contract")
	}
	level.Info(logger).Log("msg", "withdrew stake", "txHash", tx.Hash().Hex())
	return nil
}
