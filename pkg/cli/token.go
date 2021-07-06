// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/math"
)

type tokenCmd struct {
	cfgGas
	From   string  `required:""`
	To     string  `required:""`
	Amount float64 `arg:""`
}

type transferCmd tokenCmd

func (self *transferCmd) Run() error {
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

	valid := common.IsHexAddress(self.From)
	if !valid {
		return errors.Errorf("invalid etherum address:%v", self.From)
	}
	from := common.HexToAddress(self.From)

	contract, err := contracts.NewITellor(client)
	if err != nil {
		return errors.Wrap(err, "create tellor contract instance")
	}

	balance, err := contract.BalanceOf(&bind.CallOpts{Context: ctx}, from)
	if err != nil {
		return errors.Wrap(err, "get balance")
	}
	level.Info(logger).Log("msg", "current balance", math.BigInt18eToFloat(balance))

	amount, err := math.FloatToBigInt18e(self.Amount)
	if err != nil {
		return errors.Wrap(err, "invalid input amount")
	}
	if balance.Cmp(amount) < 0 {
		return errors.Errorf("insufficient balance TRB actual: %v, requested: %v",
			math.BigInt18eToFloat(balance),
			math.BigInt18eToFloat(amount))
	}

	var gasPrice *big.Int
	if self.GasPrice > 0 {
		gasPrice = big.NewInt(int64(self.GasPrice) * params.GWei)
	}

	acc, err := ethereum.GetAccountByPubAddess(self.From)
	if err != nil {
		return errors.Wrap(err, "getting auth account")
	}
	fromAuth, err := ethereum.PrepareEthTransaction(ctx, client, acc, gasPrice)
	if err != nil {
		return errors.Wrap(err, "preparing ethereum transaction")
	}

	valid = common.IsHexAddress(self.To)
	if !valid {
		return errors.Errorf("invalid etherum address:%v", self.From)
	}
	to := common.HexToAddress(self.To)

	tx, err := contract.Transfer(fromAuth, to, amount)
	if err != nil {
		return errors.Wrap(err, "calling transfer")
	}
	level.Info(logger).Log(
		"msg", "transferred",
		"amount", math.BigInt18eToFloat(amount),
		"to", to.String()[:12],
		"tx", tx.Hash(),
	)
	return nil
}

type approveCmd tokenCmd

func (self *approveCmd) Run() error {
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

	valid := common.IsHexAddress(self.From)
	if !valid {
		return errors.Errorf("invalid etherum address:%v", self.From)
	}
	from := common.HexToAddress(self.From)

	balance, err := contract.BalanceOf(&bind.CallOpts{Context: ctx}, from)
	if err != nil {
		return errors.Wrap(err, "get balance")
	}

	amount, err := math.FloatToBigInt18e(self.Amount)
	if err != nil {
		return errors.Wrap(err, "invalid input amount")
	}
	if balance.Cmp(amount) < 0 {
		return errors.Errorf("insufficient balance TRB actual: %v, requested: %v",
			math.BigInt18eToFloat(balance),
			math.BigInt18eToFloat(amount))
	}

	var gasPrice *big.Int
	if self.GasPrice > 0 {
		gasPrice = big.NewInt(int64(self.GasPrice) * params.GWei)
	}

	acc, err := ethereum.GetAccountByPubAddess(self.From)
	if err != nil {
		return errors.Wrap(err, "getting auth account")
	}

	fromAuth, err := ethereum.PrepareEthTransaction(ctx, client, acc, gasPrice)
	if err != nil {
		return errors.Wrap(err, "preparing ethereum transaction")
	}

	valid = common.IsHexAddress(self.To)
	if !valid {
		return errors.Errorf("invalid etherum address:%v", self.To)
	}
	spender := common.HexToAddress(self.To)

	tx, err := contract.Approve(fromAuth, spender, amount)
	if err != nil {
		return errors.Wrap(err, "calling approve")
	}
	level.Info(logger).Log("msg", "approved", "amount", math.BigInt18eToFloat(amount), "spender", spender.String()[:12], "tx", tx.Hash())
	return nil

}

type balanceCmd struct {
	Config  configPath `type:"existingfile" help:"path to config file"`
	Address string     `arg:"" optional:""`
}

func (self *balanceCmd) Run() error {
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

	valid := common.IsHexAddress(self.Address)
	if !valid {
		return errors.Errorf("invalid etherum address:%v", self.Address)
	}
	addr := common.HexToAddress(self.Address)

	ethBalance, err := client.BalanceAt(ctx, addr, nil)
	if err != nil {
		return errors.Wrap(err, "get eth balance")
	}
	trbBalance, err := contract.BalanceOf(&bind.CallOpts{Context: ctx}, addr)
	if err != nil {
		return errors.Wrapf(err, "getting trb balance")
	}

	level.Info(logger).Log(
		"msg", "balance check",
		"address", addr.String(),
		"ETH", math.BigInt18eToFloat(ethBalance),
		"TRB", math.BigInt18eToFloat(trbBalance),
	)
	return nil
}
