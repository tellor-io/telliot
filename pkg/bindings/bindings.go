// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package bindings

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nanmu42/etherscan-api"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

func DownloadContracts(network etherscan.Network, address string, dstFolder, name string) (map[string]string, error) {
	client := etherscan.New(network, "")
	rep, err := client.ContractSource(address)
	if err != nil {
		return nil, errors.Wrap(err, "get contract source")
	}

	// os.RemoveAll(dstFolder)
	if err := os.MkdirAll(dstFolder, os.ModePerm); err != nil {
		return nil, errors.Wrapf(err, "create download folder:%v", dstFolder)
	}

	var contractFiles = make(map[string]string)

	if codes, ok := isMultiContract(rep[0].SourceCode); ok {
		for filePath := range codes {
			content := codes[filePath].Content
			filePath := filepath.Join(dstFolder, filepath.Base(filePath))
			if err := write(filePath, content); err != nil {
				return nil, err
			}
			contractFiles[filePath] = strings.Split(rep[0].CompilerVersion, "+")[0]
		}
	} else {
		filePath := filepath.Join(dstFolder, name+".sol")
		if err := write(filePath, rep[0].SourceCode); err != nil {
			return nil, err
		}
		contractFiles[filePath] = strings.Split(rep[0].CompilerVersion, "+")[0]
	}

	return contractFiles, nil
}

func write(filePath, content string) (errFinal error) {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			errFinal = multierr.Append(errFinal, err)
		}
	}()
	w := bufio.NewWriter(f)
	defer func() {
		if err := w.Flush(); err != nil {
			errFinal = multierr.Append(errFinal, err)
		}
	}()

	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		// Flatten the import subtree.
		// Rewrite the imports to remove all parent folder.
		if strings.HasPrefix(line, "import") {
			t := strings.SplitN(line, "\"", 3)
			line = t[0] + "\"./" + filepath.Base(t[1]) + "\"" + t[2]
		}
		line += "\n"
		if _, err := w.Write([]byte(line)); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func GetContractObjects(contractFiles map[string]string) (abis []string, types []string, bins []string, sigs []map[string]string, libs map[string]string, err error) {
	libs = make(map[string]string)
	for contractPath, solcVersion := range contractFiles {
		solcPath, err := downloadSolc(solcVersion)
		if err != nil {
			return nil, nil, nil, nil, nil, errors.Wrap(err, "download solc")
		}
		contracts, err := compiler.CompileSolidity(solcPath, contractPath)
		if err != nil {
			return nil, nil, nil, nil, nil, errors.Wrap(err, "build Solidity contract")
		}

		for name, contract := range contracts {
			abi, err := json.Marshal(contract.Info.AbiDefinition)
			if err != nil {
				return nil, nil, nil, nil, nil, errors.Wrap(err, "flatten the compiler parse")
			}
			abis = append(abis, string(abi))
			bins = append(bins, contract.Code)
			sigs = append(sigs, contract.Hashes)
			nameParts := strings.Split(name, ":")
			types = append(types, nameParts[len(nameParts)-1])

			libPattern := crypto.Keccak256Hash([]byte(name)).String()[2:36]
			libs[libPattern] = nameParts[len(nameParts)-1]
		}

	}

	return types, abis, bins, sigs, libs, nil
}

func GenerateABI(folder, filename string, abis []string) error {
	var a []byte
	for _, abi := range abis {
		if len(abi) > 2 {
			a = append(a, abi[1:len(abi)-1]...)
			a = append(a, []byte(",")...)

		}
	}
	a = a[:len(a)-1] // Remove the last comma from the array.
	a = append([]byte(`[`), a...)
	a = append(a, []byte("]")...)

	fpath := filepath.Join(folder, filename+".json")
	if err := ioutil.WriteFile(fpath, a, os.ModePerm); err != nil {
		return errors.Wrapf(err, "write file:%v", fpath)
	}

	return nil
}

func GeneratePackage(pkgFolder, pkgName string, types []string, abis []string, bins []string, sigs []map[string]string, libs map[string]string) error {

	code, err := bind.Bind(types, abis, bins, sigs, pkgName, bind.LangGo, libs, map[string]string{})
	if err != nil {
		return errors.Wrapf(err, "generates the Go wrapper:%v", pkgName)
	}
	pkgFolderName := filepath.Join(pkgFolder, pkgName)

	pkgPath := filepath.Join(pkgFolderName, pkgName+".go")

	os.RemoveAll(pkgFolderName)
	if err := os.MkdirAll(pkgFolderName, os.ModePerm); err != nil {
		return errors.Wrapf(err, "create destination folder:%v", pkgFolderName)
	}

	if err := ioutil.WriteFile(pkgPath, []byte(code), os.ModePerm); err != nil {
		return errors.Wrapf(err, "write package file:%v", pkgPath)
	}
	return nil
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
		log.Println("downloading solc version", solcVersion)
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

type MultiContract struct {
	Language string
	Sources  map[string]Src
}

type Src struct {
	Content string
}

func isMultiContract(s string) (map[string]Src, bool) {
	out := &MultiContract{}

	// Etherscan has inconsistent api responses so need to deal with these here.
	if err := json.Unmarshal([]byte(s), &out.Sources); err == nil {
		return out.Sources, true
	}

	s = strings.ReplaceAll(s, "{{", "{") // Deal with another wierdness of etherscan.
	s = strings.ReplaceAll(s, "}}", "}")

	if err := json.Unmarshal([]byte(s), out); err == nil {
		return out.Sources, true
	}
	return nil, false
}
