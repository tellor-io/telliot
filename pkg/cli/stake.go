// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/math"
)

type depositCmd struct {
	cfgGasAddr
}

func (self depositCmd) Run() error {
	logger := logging.NewLogger()
	ctx := context.Background()

	_, err := config.ParseConfig(logger, string(self.Config)) // Load the env file.
	if err != nil {
		return errors.Wrap(err, "creating config")
	}
	client, err := ethereum.NewClient(ctx, logger)
	if err != nil {
		return errors.Wrap(err, "creating ethereum client")
	}

	account, err := ethereum.GetAccountByPubAddess(self.Addr)
	if err != nil {
		return err
	}
	contract, err := contracts.NewITellor(client)
	if err != nil {
		return errors.Wrap(err, "create tellor contract instance")
	}

	balance, err := contract.BalanceOf(&bind.CallOpts{Context: ctx}, account.Address)
	if err != nil {
		return errors.Wrap(err, "get TRB balance")
	}

	status, startTime, err := contract.GetStakerInfo(&bind.CallOpts{Context: ctx}, account.Address)
	if err != nil {
		return errors.Wrap(err, "get stake status")
	}

	if status.Uint64() != 0 && status.Uint64() != 2 {
		printStakeStatus(logger, status, startTime)
		return nil
	}

	stakeAmt, err := contract.GetUintVar(nil, ethereum.Keccak256([]byte("_STAKE_AMOUNT")))
	if err != nil {
		return errors.Wrap(err, "fetching stake amount")
	}

	if balance.Cmp(stakeAmt) < 0 {
		return errors.Errorf("insufficient mining stake TRB balance actual: %v, required:%v",
			math.BigInt18eToFloat(balance),
			math.BigInt18eToFloat(stakeAmt))
	}

	var gasPrice *big.Int
	if self.GasPrice > 0 {
		gasPrice = big.NewInt(int64(self.GasPrice) * params.GWei)
	}

	auth, err := ethereum.PrepareEthTransaction(ctx, client, account, gasPrice)
	if err != nil {
		return errors.Wrap(err, "prepare ethereum transaction")
	}

	tx, err := contract.DepositStake(auth)
	if err != nil {
		return errors.Wrap(err, "contract failed")
	}
	level.Info(logger).Log("msg", "stake depositied", "tx", tx.Hash())
	return nil
}

type withdrawCmd struct {
	cfgGasAddr
}

func (self withdrawCmd) Run() error {
	logger := logging.NewLogger()
	ctx := context.Background()

	_, err := config.ParseConfig(logger, string(self.Config)) // Load the env file.
	if err != nil {
		return errors.Wrap(err, "creating config")
	}

	client, err := ethereum.NewClient(ctx, logger)
	if err != nil {
		return errors.Wrap(err, "creating ethereum client")
	}

	account, err := ethereum.GetAccountByPubAddess(self.Addr)
	if err != nil {
		return err
	}
	contract, err := contracts.NewITellor(client)
	if err != nil {
		return errors.Wrap(err, "create tellor contract instance")
	}

	status, startTime, err := contract.GetStakerInfo(nil, account.Address)
	if err != nil {
		return errors.Wrap(err, "get stake status")
	}
	if status.Uint64() != 2 {
		level.Info(logger).Log("msg", "can't withdraw")
		printStakeStatus(logger, status, startTime)
		return nil
	}

	var gasPrice *big.Int
	if self.GasPrice > 0 {
		gasPrice = big.NewInt(int64(self.GasPrice) * params.GWei)
	}

	auth, err := ethereum.PrepareEthTransaction(ctx, client, account, gasPrice)
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

type requestCmd struct {
	cfgGasAddr
}

func (self requestCmd) Run() error {
	logger := logging.NewLogger()
	ctx := context.Background()

	_, err := config.ParseConfig(logger, string(self.Config)) // Load the env file.
	if err != nil {
		return errors.Wrap(err, "creating config")
	}

	client, err := ethereum.NewClient(ctx, logger)
	if err != nil {
		return errors.Wrap(err, "creating ethereum client")
	}
	account, err := ethereum.GetAccountByPubAddess(self.Addr)
	if err != nil {
		return err
	}
	contract, err := contracts.NewITellor(client)
	if err != nil {
		return errors.Wrap(err, "create tellor contract instance")
	}

	status, startTime, err := contract.GetStakerInfo(nil, account.Address)
	if err != nil {
		return errors.Wrap(err, "get stake status")
	}
	if status.Uint64() != 1 {
		printStakeStatus(logger, status, startTime)
		return nil
	}

	var gasPrice *big.Int
	if self.GasPrice > 0 {
		gasPrice = big.NewInt(int64(self.GasPrice) * params.GWei)
	}

	auth, err := ethereum.PrepareEthTransaction(ctx, client, account, gasPrice)
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

type statusCmd struct {
	cfgAddr
}

func (self statusCmd) Run() error {
	logger := logging.NewLogger()
	ctx := context.Background()

	_, err := config.ParseConfig(logger, string(self.Config)) // Load the env file.
	if err != nil {
		return errors.Wrap(err, "creating config")
	}

	client, err := ethereum.NewClient(ctx, logger)
	if err != nil {
		return errors.Wrap(err, "creating ethereum client")
	}

	contract, err := contracts.NewITellor(client)
	if err != nil {
		return errors.Wrap(err, "create tellor contract instance")
	}

	valid := common.IsHexAddress(self.Addr)
	if !valid {
		return errors.Errorf("invalid etherum address:%v", self.Addr)
	}
	addr := common.HexToAddress(self.Addr)

	status, startTime, err := contract.GetStakerInfo(&bind.CallOpts{Context: ctx}, addr)
	if err != nil {
		return errors.Wrap(err, "get stake status")
	}

	printStakeStatus(logger, status, startTime)
	return nil
}

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
