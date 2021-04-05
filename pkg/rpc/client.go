// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package rpc

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/logging"
)

const ComponentName = "rpc"

// clientInstance is the concrete implementation of the ETHClient.
type clientInstance struct {
	ethClient *ethclient.Client
	timeout   time.Duration
	logger    log.Logger
}

var (
	// retry delays that range from 100ms to 2mins.
	backoff = []uint64{100, 500, 1000, 2000, 5000, 10000, 15000, 30000, 60000, 120000}

	// rate to print errors if continue to occur in retry loop.
	errorPrintTick = time.Duration(5000)
)

// NewClient creates a new client instance.
func NewClient(logger log.Logger, cfg *config.Config, url string) (contracts.ETHClient, error) {
	timeout := time.Duration(cfg.EthClientTimeout) * time.Second
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	logger, err = logging.ApplyFilter(*cfg, ComponentName, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	return &clientInstance{
		ethClient: client,
		timeout:   timeout,
		logger:    log.With(logger, "component", ComponentName),
	}, nil
}

func (c *clientInstance) withTimeout(ctx context.Context, fn func(*context.Context) error) error {
	wTo, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	tryCount := 0
	nextTick := time.Now().Add(errorPrintTick)
	for tryCount < 20 {
		err := fn(&wTo)
		if err == nil {
			return nil
		}
		if strings.Contains(err.Error(), "nonce too low") {
			return err
		}
		if strings.Contains(err.Error(), "replacement transaction underpriced") {
			return err
		}
		level.Debug(c.logger).Log("msg", "calling eth client", "err", err)
		//pause for a bit and try again
		sleepTime := backoff[tryCount%len(backoff)]
		tryCount++
		if time.Now().After(nextTick) {
			level.Error(c.logger).Log("msg", "calling eth client", "err", err)
			nextTick = time.Now().Add(errorPrintTick)
		}

		time.Sleep(time.Duration(sleepTime))
		dl, _ := wTo.Deadline()
		if dl.Before(time.Now()) {
			return err
		}
	}
	err := fn(&wTo)
	return err
}

func (c *clientInstance) Close() {
	level.Info(c.logger).Log("msg", "closing ETHClient")
	c.ethClient.Close()
}

func (c *clientInstance) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		return c.ethClient.SendTransaction(*_ctx, tx)
	})
	return _err
}

func (c *clientInstance) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	_ = c.withTimeout(ctx, func(_ctx *context.Context) error {
		tx, isPending, err = c.ethClient.TransactionByHash(*_ctx, hash)
		return nil
	})
	return
}

func (c *clientInstance) PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error) {
	var res []byte
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.PendingCallContract(*_ctx, call)
		res = r
		return e
	})
	return res, _err
}

func (c *clientInstance) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	var res []byte
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.PendingCodeAt(*_ctx, account)
		res = r
		return e
	})
	return res, _err
}

func (c *clientInstance) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	var res []types.Log
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.FilterLogs(*_ctx, query)
		res = r
		return e
	})
	return res, _err
}

func (c *clientInstance) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	var res ethereum.Subscription
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.SubscribeFilterLogs(*_ctx, query, ch)
		res = r
		return e
	})
	return res, _err
}

func (c *clientInstance) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	return c.ethClient.SubscribeNewHead(ctx, ch)
}

func (c *clientInstance) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	var res []byte
	level.Debug(c.logger).Log("msg", "getting code at address", "contract", contract)
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.CodeAt(*_ctx, contract, blockNumber)
		if e != nil {
			level.Error(c.logger).Log("msg", "getting code at address", "err", e)
		}
		level.Debug(c.logger).Log(
			"msg", "found bytes of code at address",
			"bytes", len(r),
			"address", contract,
		)
		res = r
		return e
	})
	return res, _err
}

func (c *clientInstance) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return c.ethClient.TransactionReceipt(ctx, txHash)
}

func (c *clientInstance) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	var res []byte
	fn := hexutil.Encode(call.Data[0:4])
	level.Debug(c.logger).Log("msg", "calling contract", "fn", fn)
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.CallContract(*_ctx, call, blockNumber)
		if e != nil {
			level.Error(c.logger).Log("msg", "calling", "fn", fn, "err", e)
		}
		for i := 0; i < len(r); i += 32 {
			level.Debug(c.logger).Log("msg", "get slice", "index", i, "slice", hexutil.Encode(r[i:i+32]))
		}
		level.Debug(c.logger).Log(
			"msg", "called fn",
			"fn", fn,
			"result", r,
			"len", len(r),
		)
		res = r
		return e
	})
	return res, _err
}

func (c *clientInstance) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	var res uint64
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.PendingNonceAt(*_ctx, address)
		res = r
		return e
	})
	return res, _err
}

func (c *clientInstance) NonceAt(ctx context.Context, address common.Address) (uint64, error) {
	var res uint64
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.NonceAt(*_ctx, address, nil)
		res = r
		return e
	})
	return res, _err
}

func (c *clientInstance) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	var res *big.Int
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.SuggestGasPrice(*_ctx)
		res = r
		return e
	})
	if _err != nil {
		return nil, _err
	}
	return res, _err
}

func (c *clientInstance) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	var res uint64
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.EstimateGas(*_ctx, call)
		res = r
		return e
	})
	return res, _err
}

func (c *clientInstance) BalanceAt(ctx context.Context, address common.Address, block *big.Int) (*big.Int, error) {
	var res *big.Int
	level.Debug(c.logger).Log("msg", "getting balance of address")
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.BalanceAt(*_ctx, address, block)
		level.Debug(c.logger).Log("msg", "getting balance for", "address", address, "r", r)
		res = r
		return e
	})
	return res, _err
}

func (c *clientInstance) IsSyncing(ctx context.Context) (bool, error) {
	var syncing bool
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.SyncProgress(*_ctx)
		syncing = r != nil
		return e
	})
	return syncing, _err
}

func (c *clientInstance) NetworkID(ctx context.Context) (*big.Int, error) {
	var id *big.Int
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.NetworkID(*_ctx)
		id = r
		return e
	})
	return id, _err
}

func (c *clientInstance) HeaderByNumber(ctx context.Context, num *big.Int) (*types.Header, error) {
	var res *types.Header
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.HeaderByNumber(*_ctx, num)
		res = r
		return e
	})
	return res, _err
}
