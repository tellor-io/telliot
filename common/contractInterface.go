package common

import (
	"context"
	"math/big"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/ethereum/go-ethereum/core/types"
)

//TransactionGeneratorFN is a callback function that a TransactionSubmitter uses to actually invoke
//a contract and generate a transaction
type TransactionGeneratorFN func(ctx context.Context, contract ContractInterface) (*types.Transaction, error)

//ContractInterface represents an abstraction of function definitions that can be
//called on the smart contract. This is mostly so that we can do unit tests without
//needing to call the actual contract
type ContractInterface interface {
	AddTip(requestID *big.Int, amount *big.Int) (*types.Transaction, error)
	SubmitSolution(solution string, requestID *big.Int, value *big.Int) (*types.Transaction, error)
	NewSubmitSolution(solution string, requestID [5]*big.Int, value [5]*big.Int) (*types.Transaction, error)
	DidMine(challenge [32]byte) (bool, error)
}
//TransactionSubmitter is an abstraction for something that will callback a generator fn
//with a contract able to submit and generate transactions. This, like the ContractInterface,
//is an abstraction mainly so we can test isolated functionality.
type TransactionSubmitter interface {

	//PrepareTransaction creates a ContractInterface and sends it to the generatorFN. The ctxName is
	//primarily for logging under which context the transaction is being prepared.
	PrepareTransaction(ctx context.Context,proxy db.DataServerProxy, ctxName string, factoryFn TransactionGeneratorFN) error
}
