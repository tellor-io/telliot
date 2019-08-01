package common

import (
	"github.com/tellor-io/TellorMiner/util"
)

var (
	//ClientContextKey is the key used to set the eth client on tracker contexts
	ClientContextKey = util.NewKey("common", "ETHClient")

	//DBContextKey is the shared context key where a DB instance can be found in a context
	DBContextKey = util.NewKey("common", "DB")

	//MasterContractContextKey is the shared context key to get shared master tellor contract instance
	MasterContractContextKey = util.NewKey("common", "masterContract")

	//TransactorContractContextKey is the shared context key to get shared transactor tellor contract instance
	TransactorContractContextKey = util.NewKey("common", "transactorContract")
)
