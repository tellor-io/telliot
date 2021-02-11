package contracts

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tellor-io/telliot/pkg/contracts/ITellor"
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
	oldTellorAddress, _, proxyAddress, _, err := deployOldTellorAndMaster(transactors[0], backend)
	if err != nil {
		return err
	}
	backend.Commit()
	fmt.Println(oldTellorAddress)
	fmt.Println(proxyAddress)

	tellor, err := ITellor.NewITellor(proxyAddress, backend)
	if err != nil {
		return err
	}
	backend.Commit()

	proxy, err := tellorProxy.NewTellorMaster(proxyAddress, backend)
	if err != nil {
		return err
	}

	bal, err := proxy.TotalSupply(nil)
	if err != nil {
		return err
	}
	fmt.Println("bal")
	fmt.Println(bal)
	backend.Commit()
	bal2, err := tellor.Decimals(nil)
	if err != nil {
		return err
	}
	fmt.Println(bal2)
	fmt.Println("bal2")
	// // Add Request Ids
	// for i := 0; i < 52; i++ {
	// 	fmt.Println(i)
	// 	x := "USD" + fmt.Sprint(i)
	// 	api := "api"
	// 	_, err := tellor.RequestData(transactors[0], api, x, big.NewInt(1000), big.NewInt(int64(52-i)))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	backend.Commit()
	// }

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

func deployOldTellorAndMaster(transactor *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *bind.BoundContract, common.Address, *bind.BoundContract, error) {
	// Deploy Tellor Transfer

	fmt.Println("Transfer")
	oldTellorTransferBin, oldTellorTransferABI, err := getBinAndAbi("OldTellorTransfer")
	oldTellorTransfer, err := DeployContractWithLibs(transactor, backend, oldTellorTransferABI, oldTellorTransferBin, map[string]common.Address{})
	if err != nil {
		return common.Address{}, nil, common.Address{}, nil, err
	}
	fmt.Println(oldTellorTransfer)
	fmt.Println("Dispute")
	// Deploy Tellor Dispute
	oldTellorDisputeBin, oldTellorDisputeABI, err := getBinAndAbi("OldTellorDispute")
	oldTellorDispute, err := DeployContractWithLibs(transactor, backend, oldTellorDisputeABI, oldTellorDisputeBin, map[string]common.Address{
		"OldTellorTransfer": oldTellorTransfer,
	})
	if err != nil {
		return common.Address{}, nil, common.Address{}, nil, err
	}
	fmt.Println(oldTellorDispute)

	fmt.Println("Stake")
	// Deploy Tellor Stake
	oldTellorBinStake, oldTellorABIStake, err := getBinAndAbi("OldTellorStake")
	oldTellorStake, err := DeployContractWithLibs(transactor, backend, oldTellorABIStake, oldTellorBinStake, map[string]common.Address{
		"OldTellorTransfer": oldTellorTransfer,
		"OldTellorDispute":  oldTellorDispute,
	})
	if err != nil {
		return common.Address{}, nil, common.Address{}, nil, err
	}

	fmt.Println("Library")
	// Deploy Tellor Library
	oldTellorBinLibrary, oldTellorABILibrary, err := getBinAndAbi("OldTellorLibrary")
	oldTellorLibrary, err := DeployContractWithLibs(transactor, backend, oldTellorABILibrary, oldTellorBinLibrary, map[string]common.Address{
		"OldTellorTransfer": oldTellorTransfer,
		"OldTellorDispute":  oldTellorDispute,
		"OldTellorStake":    oldTellorStake,
	})
	if err != nil {
		return common.Address{}, nil, common.Address{}, nil, err
	}

	fmt.Println("Getter lib")
	// Deploy OldTellorGetterLibraries
	oldTellorGettersLibraryBin, oldTellorGettersLibraryABI, err := getBinAndAbi("OldTellorGettersLibrary")
	oldTellorGettersLibrary, err := DeployContractWithLibs(transactor, backend, oldTellorGettersLibraryABI, oldTellorGettersLibraryBin, nil)
	if err != nil {
		return common.Address{}, nil, common.Address{}, nil, err
	}

	fmt.Println("Old Tellor")
	// Deploy Old Tellor
	oldTellorBin, oldTellorABI, err := getBinAndAbi("OldTellor")
	oldTellorAddress, _, oldTellor, err := DeployContractWithLinks(transactor, backend, oldTellorABI, oldTellorBin, map[string]common.Address{
		"OldTellorTransfer": oldTellorTransfer,
		"OldTellorDispute":  oldTellorDispute,
		"OldTellorStake":    oldTellorStake,
		"OldTellorLibrary":  oldTellorLibrary,
	})
	if err != nil {
		return common.Address{}, nil, common.Address{}, nil, err
	}

	fmt.Println("Master")
	// Deploy Old Tellor Master
	oldTellorMasterBin, oldTellorMasterABI, err := getBinAndAbi("OldTellorMaster")
	masterAddress, _, oldTellorMaster, err := DeployContractWithLinks(transactor, backend, oldTellorMasterABI, oldTellorMasterBin, map[string]common.Address{
		"OldTellorTransfer":       oldTellorTransfer,
		"OldTellorStake":          oldTellorStake,
		"OldTellorGettersLibrary": oldTellorGettersLibrary,
	}, oldTellorAddress)
	if err != nil {
		return common.Address{}, nil, common.Address{}, nil, err
	}

	return oldTellorAddress, oldTellor, masterAddress, oldTellorMaster, nil
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
