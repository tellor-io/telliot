package rpc

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/util"
)

//ETHClient is the main abstraction interface for client operations
type ETHClient interface {

	//close the client
	Close()

	// CodeAt returns the code of the given account. This is needed to differentiate
	// between contract internal errors and the local chain being out of sync.
	CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error)

	// ContractCall executes an Ethereum contract call with the specified data as the
	// input.
	CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	NonceAt(ctx context.Context, address common.Address) (uint64, error)
	PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error)

	// PendingCodeAt returns the code of the given account in the pending state.
	PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error)

	//PendingNonceAt gets the given address's nonce for submitting transactions
	PendingNonceAt(ctx context.Context, address common.Address) (uint64, error)
	EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error)
	SuggestGasPrice(ctx context.Context) (*big.Int, error)

	FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error)
	SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
	BalanceAt(ctx context.Context, address common.Address, block *big.Int) (*big.Int, error)
	SendTransaction(ctx context.Context, tx *types.Transaction) error
	IsSyncing(ctx context.Context) (bool, error)
	NetworkID(ctx context.Context) (*big.Int, error)
	HeaderByNumber(ctx context.Context, num *big.Int) (*types.Header, error)
}

//clientInstance is the concrete implementation of the ETHClient
type clientInstance struct {
	ethClient *ethclient.Client
	timeout   time.Duration
	log       *util.Logger
}

var (
	//retry delays that range from 100ms to 2mins
	backoff = []uint64{100, 500, 1000, 2000, 5000, 10000, 15000, 30000, 60000, 120000}

	//rate to print errors if continue to occur in retry loop
	errorPrintTick = time.Duration(5000)
)

//NewClient creates a new client instance
func NewClient(url string) (ETHClient, error) {
	cfg := config.GetConfig()
	timeout := time.Duration(cfg.EthClientTimeout) * time.Second
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &clientInstance{ethClient: client, timeout: timeout, log: util.NewLogger("rpc", "client")}, nil
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
		c.log.Debug("Problem in calling eth client: %v", err)
		//pause for a bit and try again
		sleepTime := backoff[tryCount%len(backoff)]
		tryCount++
		if time.Now().After(nextTick) {
			c.log.Error("Problem calling ethClient: %v\n", err)
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
	c.log.Info("Closing ETHClient")
	c.ethClient.Close()
}

func (c *clientInstance) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		//c.log.Info("Sending txn on-chain: %v\n", tx)
		fmt.Println("TX SENT CLIENT", fmt.Sprintf("%x", tx))
		e := c.ethClient.SendTransaction(*_ctx, tx)
		return e
	})
	return _err
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

func (c *clientInstance) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	var res []byte
	c.log.Debug("Getting code at address", contract)
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.CodeAt(*_ctx, contract, blockNumber)
		if e != nil {
			c.log.Error("Problem getting code from eth client: %v", e)
		}
		log.Printf("_normalLog Found %d bytes of code at address: %v", len(r), contract)
		c.log.Debug("Found %d bytes of code at address: %v", len(r), contract)
		res = r
		return e
	})
	return res, _err
}

func (c *clientInstance) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	var res []byte
	fn := hexutil.Encode(call.Data[0:4])
	c.log.Debug("Calling contract fn: %v\n", fn)
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.CallContract(*_ctx, call, blockNumber)
		if e != nil {
			c.log.Error("Problem calling %s: %v", fn, e)
		}
		for i := 0; i < len(r); i += 32 {
			c.log.Debug("Slice %d: %s\n", i, hexutil.Encode(r[i:i+32]))
		}
		c.log.Debug("Called fn: %s with result %v (len: %d)", fn, r, len(r))
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
	c.log.Debug("Getting balance of address")
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.BalanceAt(*_ctx, address, block)
		c.log.Debug("Getting balance for address %v: $v\n", address, r)
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
