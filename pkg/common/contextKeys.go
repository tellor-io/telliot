// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package common

import (
	"github.com/tellor-io/TellorMiner/pkg/util"
)

var (
	// ClientContextKey is the key used to set the eth client on tracker contexts.
	ClientContextKey = util.NewKey("common", "ETHClient")

	// DBContextKey is the shared context key where a DB instance can be found in a context.
	DBContextKey = util.NewKey("common", "DB")

	// Tellor Contract Address.
	ContractAddress = util.NewKey("common", "contractAddress")

	// ContractsTellorContextKey is the shared context key to get the tellor contract instance.
	ContractsTellorContextKey = util.NewKey("common", "ContractsTellorInstance")

	// ContractsGetterContextKey is the shared context key to get the getter contract instance.
	ContractsGetterContextKey = util.NewKey("common", "ContractsGetterInstance")

	// DataProxyKey used to access the local or remote data server proxy.
	DataProxyKey = util.NewKey("common", "DataServerProxy")

	// Ethereum wallet private key.
	PrivateKey = util.NewKey("common", "PrivateKey")

	// Ethereum wallet public address.
	PublicAddress = util.NewKey("common", "PublicAddress")

	// PriceTXs is the key used to save transactions cost
	// These are used to calculate the profitability when submitting a solution.
	PriceTXs = "PriceTXSlot"
)
