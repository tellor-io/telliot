// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nanmu42/etherscan-api"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile | log.Lmsgprefix)

	f, err := ioutil.ReadFile(filepath.Join("configs", "config.json"))
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.Config{}

	if err := json.Unmarshal(f, &cfg); err != nil {
		log.Fatal(err)
	}

	contractsFolder := filepath.Join("scripts", "bindings", "contracts")
	uinswapContractFolder := filepath.Join(contractsFolder, "uniswap")
	balancerContractFolder := filepath.Join(contractsFolder, "balancer")
	if err := os.RemoveAll(contractsFolder); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(contractsFolder, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := os.Mkdir(balancerContractFolder, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := os.Mkdir(uinswapContractFolder, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	generate(contracts.TellorRinkebyAddress, filepath.Join(contractsFolder, "proxy.sol"), "v0.5.16")
	log.Println("Generated proxy contract in:", contractsFolder)
	time.Sleep(5 * time.Second)
	// TODO how to detect that the proxy has changed and this needs updating.
	generate("0x7e05e8a675e649261acc19423db34dd4826f9a98", filepath.Join(contractsFolder, "master.sol"), "v0.5.16")
	log.Println("Generated master contract in:", contractsFolder)
	time.Sleep(5 * time.Second)
	// Generating Balancer core pool factory contract.
	generate("0x9C84391B443ea3a48788079a5f98e2EaD55c9309", balancerContractFolder, "")
	log.Println("Generated balancer contract in:", balancerContractFolder)
	time.Sleep(5 * time.Second)
	// Generating Uniswap factory contract.
	generate("0x03E6c12eF405AC3F642B9184eDed8E1322de1a9e", filepath.Join(uinswapContractFolder, "uniswap.sol"), "")
	log.Println("Generated balancer contract in:", uinswapContractFolder)

}

func generate(address string, contractFPath string, solcVersion string) {
	filename := filepath.Base(contractFPath)
	pkgName := strings.TrimSuffix(filename, filepath.Ext(filename))
	client := etherscan.New(etherscan.Rinkby, "")
	src, err := client.ContractSource(address)
	if err != nil {
		log.Fatal(err)
	}

	var contractFiles = []string{contractFPath}
	var codes map[string]Code
	var ok bool

	if codes, ok = isJSONString(src[0].SourceCode); ok {
		contractFiles = []string{}
		for fileName := range codes {
			filePath := filepath.Join(contractFPath, fileName)
			contractFiles = append(contractFiles, filePath)
			if err := ioutil.WriteFile(filePath, []byte(codes[fileName].Content), os.ModePerm); err != nil {
				log.Fatal(err)
			}
		}

	} else if err := ioutil.WriteFile(contractFPath, []byte(src[0].SourceCode), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// Get solc version from etherscan if empty.
	if solcVersion == "" {
		solcVersion = strings.Split(src[0].CompilerVersion, "+")[0]
	}
	solcPath, err := downloadSolc(solcVersion)
	if err != nil {
		log.Fatal(err)
	}

	contracts, err := compiler.CompileSolidity(solcPath, contractFiles...)
	if err != nil {
		utils.Fatalf("Failed to build Solidity contract: %v", err)
	}
	var (
		abis  []string
		types []string
		bins  []string
		sigs  []map[string]string
		libs  = make(map[string]string)
	)
	// Gather all non-excluded contract for binding
	for name, contract := range contracts {
		abi, err := json.Marshal(contract.Info.AbiDefinition) // Flatten the compiler parse
		if err != nil {
			log.Fatal(err)
		}
		abis = append(abis, string(abi))
		bins = append(bins, contract.Code)
		sigs = append(sigs, contract.Hashes)
		nameParts := strings.Split(name, ":")
		types = append(types, nameParts[len(nameParts)-1])

		libPattern := crypto.Keccak256Hash([]byte(name)).String()[2:36]
		libs[libPattern] = nameParts[len(nameParts)-1]
	}

	code, err := bind.Bind(types, abis, bins, sigs, pkgName, bind.LangGo, libs, map[string]string{})
	if err != nil {
		log.Fatal(err)
	}

	pkgFolder := filepath.Join("pkg", "contracts", pkgName)
	pkgPath := filepath.Join(pkgFolder, pkgName+".go")

	os.RemoveAll(pkgFolder)
	if err := os.MkdirAll(pkgFolder, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(pkgPath, []byte(code), os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

// downloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return errors.Wrap(err, "gettings the file")
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return errors.Wrap(err, "creating destination file")
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return errors.Wrap(err, "writing the file")
}

// downloadSolc will download @solcVersion of the Solc compiler to tmp/solc directory.
func downloadSolc(solcVersion string) (string, error) {
	solcDir := filepath.Join("tmp", "solc")
	if err := os.MkdirAll(solcDir, os.ModePerm); err != nil {
		return "", err
	}
	solcPath := filepath.Join(solcDir, solcVersion)
	if _, err := os.Stat(solcPath); os.IsNotExist(err) {
		log.Println("downloading solc")
		err = downloadFile(solcPath, fmt.Sprintf("https://github.com/ethereum/solidity/releases/download/%s/solc-static-linux", solcVersion))
		if err != nil {
			return "", err
		}
		if err := os.Chmod(solcPath, os.ModePerm); err != nil {
			return "", err
		}
	}
	return solcPath, nil
}

type Code struct {
	Content string
}

func isJSONString(s string) (map[string]Code, bool) {
	var out map[string]Code
	ok := json.Unmarshal([]byte(s), &out) == nil
	return out, ok
}
