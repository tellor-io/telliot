// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package ops

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kit/kit/log"
	tellorCommon "github.com/tellor-io/telliot/pkg/common"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
)

// TxnSubmitter just concrete type for txn submitter.
type TxnSubmitter struct {
	logger log.Logger
}

// NewSubmitter creates a new TxnSubmitter instance.
func NewSubmitter(logger log.Logger) TxnSubmitter {
	return TxnSubmitter{logger: log.With(logger, "component", "submitter")}
}

// Submit relies on rpc package to prepare and submit transactions.
func (s TxnSubmitter) Submit(ctx context.Context, proxy db.DataServerProxy, ctxName string, callback tellorCommon.TransactionGeneratorFN) (*types.Transaction, error) {
	return rpc.SubmitContractTxn(ctx, s.logger, proxy, ctxName, callback)
}
