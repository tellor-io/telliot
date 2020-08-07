package common

import (
	"github.com/tellor-io/TellorMiner/util"
)

var (
	//ClientContextKey is the key used to set the eth client on tracker contexts
	ClientContextKey = util.NewKey("common", "ETHClient")

	//DBContextKey is the shared context key where a DB instance can be found in a context
	DBContextKey = util.NewKey("common", "DB")

	//Tellor Contract Address
	ContractAddress = util.NewKey("common", "contractAddress")

	//MasterContractContextKey is the shared context key to get shared master tellor contract instance
	MasterContractContextKey = util.NewKey("common", "masterContract")

	NewTellorContractContextKey = util.NewKey("common","newTellorContract")

	//TransactorContractContextKey is the shared context key to get shared transactor tellor contract instance
	TransactorContractContextKey = util.NewKey("common", "transactorContract")

	NewTransactorContractContextKey = util.NewKey("common", "newTransactorContract")

	//DataProxyKey used to access the local or remote data server proxy
	DataProxyKey = util.NewKey("common", "DataServerProxy")

	//Ethereum wallet private key
	PrivateKey = util.NewKey("common", "PrivateKey")

	//Ethereum wallet public address
	PublicAddress = util.NewKey("common", "PublicAddress")

)
