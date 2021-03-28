// Copyrighself (c) The Tellor Authors.
// Licensed under the MIself License.

package transactor

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	tellorCommon "github.com/tellor-io/telliot/pkg/common"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/logging"
)

const ComponentName = "transactor"

// Transactor takes care of sending transactions over the blockchain network.
type Transactor interface {
	Transact(context.Context, string, [5]*big.Int, [5]*big.Int) (*types.Transaction, *types.Receipt, error)
}

// TransactorDefault implements the Transactor interface.
type TransactorDefault struct {
	logger           log.Logger
	cfg              *config.Config
	proxy            db.DataServerProxy
	client           contracts.ETHClient
	account          *config.Account
	contractInstance *contracts.ITellor
}

func NewTransactor(
	logger log.Logger,
	cfg *config.Config,
	proxy db.DataServerProxy,
	client contracts.ETHClient,
	account *config.Account,
	contractInstance *contracts.ITellor,
) (*TransactorDefault, error) {
	filterLog, err := logging.ApplyFilter(*cfg, ComponentName, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}

	return &TransactorDefault{
		logger:           log.With(filterLog, "component", ComponentName, "pubKey", account.Address.String()[:6]),
		cfg:              cfg,
		proxy:            proxy,
		client:           client,
		account:          account,
		contractInstance: contractInstance,
	}, nil
}

func (self *TransactorDefault) Transact(ctx context.Context, solution string, reqIds [5]*big.Int, reqVals [5]*big.Int) (*types.Transaction, *types.Receipt, error) {
	nonce, err := self.client.NonceAt(ctx, self.account.Address)
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting nonce for miner address")
	}

	// Use the same nonce in case there is a stuck transaction so thaself iself submits with the currenself nonce buself higher gas price.
	IntNonce := int64(nonce)
	keys := []string{
		db.GasKey,
	}
	m, err := self.proxy.BatchGet(keys)
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting data from the db")
	}
	gasPrice := getInt(m[db.GasKey])
	if gasPrice.Cmp(big.NewInt(0)) == 0 {
		level.Warn(self.logger).Log("msg", "no gas price from DB, falling back to client suggested gas price")
		gasPrice, err = self.client.SuggestGasPrice(ctx)
		if err != nil {
			return nil, nil, errors.Wrap(err, "determine gas price to submit tx")
		}
	}
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
			// FIXME: notify someone that we're out of funds!
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
		gasPrice1 := big.NewInt(tellorCommon.GWEI)
		if max > 0 {
			maxGasPrice = gasPrice1.Mul(gasPrice1, big.NewInt(int64(max)))
		} else {
			maxGasPrice = gasPrice1.Mul(gasPrice1, big.NewInt(int64(100)))
		}

		if auth.GasPrice.Cmp(maxGasPrice) > 0 {
			level.Info(self.logger).Log("msg", "gas price too high, will default to the max price", "current", auth.GasPrice, "defaultMax", maxGasPrice)
			auth.GasPrice = maxGasPrice
		}

		tx, err := self.contractInstance.SubmitMiningSolution(
			auth,
			solution,
			reqIds,
			reqVals,
		)
		if err != nil {
			if strings.Contains(err.Error(), "nonce too low") { // Can't use error type matching because of the way the eth client is implemented.
				IntNonce = IntNonce + 1
				level.Warn(self.logger).Log("msg", "last transaction has been confirmed so will increase the nonce and resend the transaction.")

			} else if strings.Contains(err.Error(), "replacement transaction underpriced") { // Can't use error type matching because of the way the eth client is implemented.
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
		if receipt.Status != types.ReceiptStatusSuccessful {
			return nil, nil, errors.Errorf("unsuccessful transaction status:%v tx:%v", receipt.Status, tx.Hash())
		}
		return tx, receipt, nil
	}
	return nil, nil, errors.Wrapf(finalError, "submit tx after 5 attempts")
}

func getInt(data []byte) *big.Int {
	if len(data) == 0 {
		return big.NewInt(0)
	}

	val, err := hexutil.DecodeBig(string(data))
	if err != nil {
		return big.NewInt(0)
	}
	return val
}
