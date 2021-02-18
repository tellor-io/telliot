// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package submitter

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kit/kit/log"
	tellorCommon "github.com/tellor-io/telliot/pkg/common"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
)

// TxnSubmitter just concrete type for txn submitter.
type TxnSubmitter struct {
	cfg      *config.Config
	client   contracts.ETHClient
	contract *contracts.ITellor
	account  *rpc.Account
	logger   log.Logger
}

// NewSubmitter creates a new TxnSubmitter instance.
func NewSubmitter(
	logger log.Logger,
	cfg *config.Config,
	client contracts.ETHClient,
	tellor *contracts.ITellor,
	account *rpc.Account) TxnSubmitter {
	return TxnSubmitter{
		cfg:      cfg,
		client:   client,
		contract: tellor,
		account:  account,
		logger:   log.With(logger, "component", ComponentName),
	}
}

// Submit relies on rpc package to prepare and submit transactions.
func (s TxnSubmitter) Submit(ctx context.Context, proxy db.DataServerProxy, ctxName string, callback tellorCommon.TransactionGeneratorFN) (*types.Transaction, error) {
	return rpc.SubmitContractTxn(ctx, s.logger, s.cfg, proxy, s.client, s.contract, s.account, ctxName, callback)
}

// Address returns the public key address of the account.
func (s TxnSubmitter) Address() common.Address {
	return s.account.Address
}
