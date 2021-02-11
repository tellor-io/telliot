// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package main

import (
	"log"
	"path/filepath"
	"time"

	"github.com/nanmu42/etherscan-api"
	"github.com/tellor-io/telliot/pkg/bindings"
	"github.com/tellor-io/telliot/pkg/config"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile | log.Lmsgprefix)

	downlContractsFolder := filepath.Join("scripts", "bindings", "contracts")
	pkgFolder := filepath.Join("pkg", "contracts")

	// Bindings for the oracle proxy.
	downloadAndGenerate(config.TellorRinkebyAddress, downlContractsFolder, pkgFolder, "tellorProxy")
	time.Sleep(5 * time.Second)

	// Bindings for the oracle master.
	downloadAndGenerate("0x7e05e8a675e649261acc19423db34dd4826f9a98", downlContractsFolder, pkgFolder, "tellorMaster")
	time.Sleep(5 * time.Second)

	// Bindings for balancer.
	downloadAndGenerate("0x9C84391B443ea3a48788079a5f98e2EaD55c9309", downlContractsFolder, pkgFolder, "balancer")
	time.Sleep(5 * time.Second)

	// Bindings for uniswap.
	downloadAndGenerate("0x03E6c12eF405AC3F642B9184eDed8E1322de1a9e", downlContractsFolder, pkgFolder, "uniswap")
	time.Sleep(5 * time.Second)

	// Bindings for uniswap.
	downloadAndGenerate("0xAf96A11a622f78399b5a12503D429750525273Bd", downlContractsFolder, pkgFolder, "oldTellor")
	time.Sleep(5 * time.Second)
}

func downloadAndGenerate(addr, downlContractsFolder, pkgFolder, name string) {
	downloadFolder := filepath.Join(downlContractsFolder, name)

	filePaths, err := bindings.DownloadContracts(etherscan.Rinkby, addr, downloadFolder, name)
	ExitOnErr(err, "download contracts")

	log.Printf("Downloaded contract:%+v", filePaths)
	types, abis, bins, sigs, libs, err := bindings.GetContractObjects(filePaths)
	ExitOnErr(err, "get contracts object")

	err = bindings.GenerateABI(downloadFolder, name, abis)
	ExitOnErr(err, "generate ABI")
	log.Println("Generated ABI:", filepath.Join(downloadFolder, name))

	err = bindings.GeneratePackage(pkgFolder, name, types, abis, bins, sigs, libs)
	ExitOnErr(err, "generate GO binding")

	log.Println("generated GO binding:", filepath.Join(pkgFolder, name))
}

func ExitOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("execution error:%+v msg:%+v", err, msg)
	}
}
