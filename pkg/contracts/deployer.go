package contracts

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/tellor-io/telliot/pkg/contracts/oldTellor"
	"github.com/tellor-io/telliot/pkg/contracts/tellorProxy"
)

func GetTellorCore() error {
	backend, transactors, err := GetbackendBackend()
	if err != nil {
		return err
	}

	initialMiner := common.HexToAddress("e010aC6e0248790e08F42d5F697160DEDf97E024")
	fmt.Println(initialMiner)

	// Deploy oldTellor
	add, _, _, err := deployOldTellor(transactors[0], backend)
	if err != nil {
		return err
	}
	// Deploy proxyMaster
	masterAdd, _, proxy, err := tellorProxy.DeployTellorMaster(transactors[0], backend, add)
	if err != nil {
		return err
	}
	backend.Commit()
	fmt.Println(masterAdd)
	fmt.Println(proxy)
	tellor, err := oldTellor.NewOldTellor(masterAdd, backend)
	if err != nil {
		return err
	}

	bal, err := proxy.BalanceOf(nil, initialMiner)
	if err != nil {
		return err
	}
	fmt.Println(bal)

	bal2, err := tellor.AddTip(transactors[0], big.NewInt(0), big.NewInt(30))
	if err != nil {
		return err
	}
	fmt.Println(bal2)

	// // Add Request Ids
	for i := 0; i < 52; i++ {
		fmt.Println(i)
		x := "USD" + fmt.Sprint(i)
		api := "api"
		_, err := tellor.RequestData(transactors[0], api, x, big.NewInt(1000), big.NewInt(int64(52-i)))
		if err != nil {
			return err
		}
		backend.Commit()
	}

	// //need 5 initial miners to mine 1st block
	// for i := 1; i < 5; i++ {
	// 	_, err := tellor.SubmitMiningSolution(transactors[i], "nonce", big.NewInt(1), big.NewInt(1200))
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	// backend.Commit()

	// // Deploy most recent tellor
	// addv26, tx, _, err := tellorMaster.DeployTellor(transactors[0], backend)
	// if err != nil {
	// 	return err
	// }

	// _, err = proxy.ChangeTellorContract(transactors[0], addv26)

	// backend.Commit()
	// fmt.Println("fafaf")
	return nil
}

func deployOldTellor(transactor *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *bind.BoundContract, error) {
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

func getBinAndAbi(contractName string) (string, string, error) {
	println(os.Getwd())
	b, err := ioutil.ReadFile("../contracts/oldTellor/bin/" + contractName + ".bin")
	if err != nil {
		fmt.Print(err)
		return "", "", nil
	}
	abi, err := ioutil.ReadFile("../contracts/oldTellor/abi/" + contractName + ".abi")
	return string(b), string(abi), nil
}
