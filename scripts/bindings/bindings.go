package main

import (
	"encoding/json"
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
	if err := os.RemoveAll(contractsFolder); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(contractsFolder, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	generate(cfg.ContractAddress, filepath.Join(contractsFolder, "proxy.sol"))
	log.Println("Generated proxy contract in:", contractsFolder)
	time.Sleep(5 * time.Second)
	// TODO how to detect that the proxy has changed and this needs updating.
	generate("0x7e05e8a675e649261acc19423db34dd4826f9a98", filepath.Join(contractsFolder, "master.sol"))
	log.Println("Generated master contract in:", contractsFolder)
}

func generate(address string, contractFPath string) {
	filename := filepath.Base(contractFPath)
	pkgName := strings.TrimSuffix(filename, filepath.Ext(filename))
	client := etherscan.New(etherscan.Rinkby, "")
	src, err := client.ContractSource(address)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(contractFPath, []byte(src[0].SourceCode), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll("tmp", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	solcPath := filepath.Join("tmp", "solc")
	if _, err := os.Stat(solcPath); os.IsNotExist(err) {
		log.Println("downloading solc")
		err = downloadFile(solcPath, "https://github.com/ethereum/solidity/releases/download/v0.5.16/solc-static-linux")
		if err != nil {
			log.Fatal(err)
		}
		if err := os.Chmod(solcPath, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	contracts, err := compiler.CompileSolidity(solcPath, contractFPath)
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
