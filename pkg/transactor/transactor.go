// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package transactor

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/tracker/gasPrice"
)

const ComponentName = "transactor"

type Config struct {
	LogLevel      string
	GasMax        uint
	GasMultiplier int
}

// Transactor takes care of sending transactions over the blockchain network.
type Transactor interface {
	Transact(context.Context, func(*bind.TransactOpts) (*types.Transaction, error)) (*types.Transaction, *types.Receipt, error)
}

// TransactorDefault implements the Transactor interface.
type TransactorDefault struct {
	cfg             Config
	logger          log.Logger
	gasPriceTracker *gasPrice.GasTracker
	client          contracts.ETHClient
	account         *ethereum.Account
}

func New(
	logger log.Logger,
	cfg Config,
	gasPriceTracker *gasPrice.GasTracker,
	client contracts.ETHClient,
	account *ethereum.Account,
) (*TransactorDefault, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}

	return &TransactorDefault{
		cfg:             cfg,
		logger:          log.With(logger, "component", ComponentName, "addr", account.Address.String()[:6]),
		gasPriceTracker: gasPriceTracker,
		client:          client,
		account:         account,
	}, nil
}

func (self *TransactorDefault) Transact(ctx context.Context, contractCall func(*bind.TransactOpts) (*types.Transaction, error)) (*types.Transaction, *types.Receipt, error) {
	nonce, err := self.client.NonceAt(ctx, self.account.Address)
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting nonce for miner address")
	}

	// Use the same nonce in case there is a stuck transaction so thaself iself submits with the currenself nonce buself higher gas price.
	IntNonce := int64(nonce)

	_gasPrice, err := self.gasPriceTracker.Query(ctx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting data from the db")
	}
	gasPrice := big.NewInt(int64(_gasPrice))

	mul := self.cfg.GasMultiplier
	if mul > 0 {
		level.Info(self.logger).Log("msg", "settings gas price multiplier", "value", mul)
		gasPrice = gasPrice.Mul(gasPrice, big.NewInt(int64(mul)))
	}

	var finalError error
	for i := 0; i <= 5; i++ {
		balance, err := self.client.BalanceAt(ctx, self.account.Address, nil)
		if err != nil {
			finalError = err
			continue
		}

		cost := big.NewInt(1)
		cost = cost.Mul(gasPrice, big.NewInt(200000))
		if balance.Cmp(cost) < 0 {
			finalError = errors.Errorf("insufficient funds to send transaction: %v < %v", balance, cost)
			continue
		}

		netID, err := self.client.NetworkID(ctx)
		if err != nil {
			return nil, nil, errors.Wrap(err, "getting network id")
		}
		auth, err := bind.NewKeyedTransactorWithChainID(self.account.PrivateKey, netID)
		if err != nil {
			return nil, nil, errors.Wrap(err, "creating transactor")
		}
		auth.Nonce = big.NewInt(IntNonce)
		auth.Value = big.NewInt(0)      // in weiF
		auth.GasLimit = uint64(3000000) // in units
		if gasPrice.Cmp(big.NewInt(0)) == 0 {
			gasPrice = big.NewInt(100)
		}
		if i > 1 {
			gasPrice1 := new(big.Int).Set(gasPrice)
			gasPrice1.Mul(gasPrice1, big.NewInt(int64(i*11))).Div(gasPrice1, big.NewInt(int64(100)))
			auth.GasPrice = gasPrice1.Add(gasPrice, gasPrice1)
		} else {
			// First time, try base gas price.
			auth.GasPrice = gasPrice
		}
		max := self.cfg.GasMax
		var maxGasPrice *big.Int
		gasPrice1 := big.NewInt(params.GWei)
		if max > 0 {
			maxGasPrice = gasPrice1.Mul(gasPrice1, big.NewInt(int64(max)))
		} else {
			maxGasPrice = gasPrice1.Mul(gasPrice1, big.NewInt(int64(100)))
		}

		if auth.GasPrice.Cmp(maxGasPrice) > 0 {
			level.Info(self.logger).Log("msg", "gas price too high, will default to the max price", "current", auth.GasPrice, "defaultMax", maxGasPrice)
			auth.GasPrice = maxGasPrice
		}

		tx, err := contractCall(auth)
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "nonce too low") { // Can't use error type matching because of the way the eth client is implemented.
				IntNonce = IntNonce + 1
				level.Warn(self.logger).Log("msg", "last transaction has been confirmed so will increase the nonce and resend the transaction.")

			} else if strings.Contains(strings.ToLower(err.Error()), "replacement transaction underpriced") { // Can't use error type matching because of the way the eth client is implemented.
				level.Warn(self.logger).Log("msg", "last transaction is stuck so will increase the gas price and try to resend")
				finalError = err
			} else {
				finalError = errors.Wrap(err, "contract call SubmitMiningSolution")
			}

			delay := 15 * time.Second
			level.Info(self.logger).Log("msg", "will retry a send", "retryDelay", delay)
			select {
			case <-ctx.Done():
				return nil, nil, errors.New("the submit context was canceled")
			case <-time.After(delay):
				continue
			}
		}

		receipt, err := bind.WaitMined(ctx, self.client, tx)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "transaction result tx:%v", tx.Hash())
		}
		return tx, receipt, nil
	}
	return nil, nil, errors.Wrapf(finalError, "submit tx after 5 attempts")
}
