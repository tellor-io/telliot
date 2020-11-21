// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package ops

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	tellorCommon "github.com/tellor-io/telliot/pkg/common"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
)

// TxnSubmitter just concrete type for txn submitter.
type TxnSubmitter struct {
}

// NewSubmitter creates a new TxnSubmitter instance.
func NewSubmitter() TxnSubmitter {
	return TxnSubmitter{}
}

// Submit relies on rpc package to prepare and submit transactions.
func (s TxnSubmitter) Submit(ctx context.Context, proxy db.DataServerProxy, ctxName string, callback tellorCommon.TransactionGeneratorFN) (*types.Transaction, error) {
	return rpc.SubmitContractTxn(ctx, proxy, ctxName, callback)
}
