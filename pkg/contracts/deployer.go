package contracts

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/tellor-io/telliot/pkg/contracts/oldTellor"
)

func DeployOldTellor(transactor *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *bind.BoundContract, error) {
	// Deploy Tellor Transfer
	oldTellorTransferBin, oldTellorTransferABI, err := getBinAndAbi("OldTellorTransfer")
	oldTellorTransfer, err := DeployContractWithLibs(transactor, backend, oldTellorTransferABI, oldTellorTransferBin, map[string]common.Address{})
	if err != nil {

	}

	// Deploy Tellor Dispute
	oldTellorBinDispute, oldTellorABIDispute, err := getBinAndAbi("OldTellorDispute")
	oldTellorDispute, err := DeployContractWithLibs(transactor, backend, oldTellorABIDispute, oldTellorBinDispute, map[string]common.Address{
		"OldTellorTransfer": oldTellorTransfer,
	})

	// Deploy Tellor Stake
	oldTellorBinStake, oldTellorABIStake, err := getBinAndAbi("OldTellorStake")
	oldTellorStake, err := DeployContractWithLibs(transactor, backend, oldTellorABIStake, oldTellorBinStake, map[string]common.Address{
		"OldTellorTransfer": oldTellorTransfer,
		"OldTellorDispute":  oldTellorDispute,
	})

	// Deploy Tellor Library
	oldTellorBinLibrary, oldTellorABILibrary, err := getBinAndAbi("OldTellorLibrary")
	oldTellorLibrary, err := DeployContractWithLibs(transactor, backend, oldTellorABILibrary, oldTellorBinLibrary, map[string]common.Address{
		"OldTellorTransfer": oldTellorTransfer,
		"OldTellorDispute":  oldTellorDispute,
		"OldTellorStake":    oldTellorStake,
	})

	// Deploy Old Tellor
	oldTellorBin, oldTellorABI, err := getBinAndAbi("OldTellor")
	return DeployContractWithLinks(transactor, backend, oldTellorABI, oldTellorBin, map[string]common.Address{
		"OldTellorTransfer": oldTellorTransfer,
		"OldTellorDispute":  oldTellorDispute,
		"OldTellorStake":    oldTellorStake,
		"OldTellorLibrary":  oldTellorLibrary,
	})
}

func addRequestIds(tellor *oldTellor.OldTellor) error {

}

func getBinAndAbi(contractName string) (string, string, error) {
	println(os.Getwd())
	b, err := ioutil.ReadFile("../contracts/abigenBindings/bin/" + contractName + ".bin")
	if err != nil {
		fmt.Print(err)
		return "", "", nil
	}
	abi, err := ioutil.ReadFile("../contracts/abigenBindings/abi/" + contractName + ".abi")
	return string(b), string(abi), nil
}
