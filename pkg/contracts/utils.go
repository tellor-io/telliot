package contracts

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// ABILinkLibrary replaces references to a library
// with the actual addresses to those library contracts
func ABILinkLibrary(bin string, libraryName string, libraryAddress common.Address) string {
	libstr := fmt.Sprintf("_+%v_+", libraryName)
	libraryRexp := regexp.MustCompile(libstr)

	// Remove the 0x prefix from those addresses, just need the actual hex string
	cleanLibraryAddr := strings.Replace(libraryAddress.Hex(), "0x", "", -1)

	modifiedBin := libraryRexp.ReplaceAllString(bin, cleanLibraryAddr)

	return modifiedBin
}

// DeployContractWithLinks patches a contract bin with provided library addresses
func DeployContractWithLinks(
	opts *bind.TransactOpts,
	backend bind.ContractBackend,
	abiString string,
	bin string,
	libraries map[string]common.Address,
	params ...interface{},
) (common.Address, *types.Transaction, *bind.BoundContract, error) {

	for libraryName, libraryAddress := range libraries {
		bin = ABILinkLibrary(bin, libraryName, libraryAddress)
	}

	parsed, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	return bind.DeployContract(opts, parsed, common.FromHex(bin), backend, params...)
}

// DeployContractWithLinks patches a contract bin with provided library addresses
func DeployContractWithLibs(
	opts *bind.TransactOpts,
	backend bind.ContractBackend,
	abiString string,
	bin string,
	libraries map[string]common.Address,
	params ...interface{},
) (common.Address, error) {

	for libraryName, libraryAddress := range libraries {
		bin = ABILinkLibrary(bin, libraryName, libraryAddress)
	}

	parsed, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		return common.Address{}, err
	}

	address, _, _, err := bind.DeployContract(opts, parsed, common.FromHex(bin), backend, params...)
	if err != nil {
		return common.Address{}, err
	}
	return address, err
}

func GetbackendBackend() (*backends.SimulatedBackend, []*bind.TransactOpts, error) {
	initialKeys := []string{
		"3a10b4bc1258e8bfefb95b498fb8c0f0cd6964a811eabca87df5630bcacd7216", "d32132133e03be292495035cf32e0e2ce0227728ff7ec4ef5d47ec95097ceeed", "d13dc98a245bd29193d5b41203a1d3a4ae564257d60e00d6f68d120ef6b796c5", "4beaa6653cdcacc36e3c400ce286f2aefd59e2642c2f7f29804708a434dd7dbe", "78c1c7e40057ea22a36a0185380ce04ba4f333919d1c5e2effaf0ae8d6431f14", "4bdc16637633fa4b4854670fbb83fa254756798009f52a1d3add27fb5f5a8e16", "42ef6879f87950460bc162070839a42690ad76200e2460e30e944f69026a7f0b", "fa991490959b6cf3c31115271c8ee63070dd57b6078582a6b8b5be97ca9a8061", "37ed7b1172f31891e6fb38a361c72768954031f46c3e4018f9f8578ea2b6804c", "8b73fa2c839ccea66e8eddf0aa95f6bc4c6aaa11e2fa126c1d9334985b0e7666",
	}
	adds := map[common.Address]core.GenesisAccount{}
	transactors := []*bind.TransactOpts{}
	for _, sk := range initialKeys {
		s, err := crypto.HexToECDSA(sk)
		if err != nil {
			return nil, nil, err
		}
		address := crypto.PubkeyToAddress(s.PublicKey)
		adds[address] = core.GenesisAccount{Balance: new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(9))}
		transactor, err := bind.NewKeyedTransactorWithChainID(s, big.NewInt(1337))
		if err != nil {
			return nil, nil, err
		}
		transactors = append(transactors, transactor)
	}

	addr := map[common.Address]core.GenesisAccount{
		common.BytesToAddress([]byte{1}): {Balance: big.NewInt(1)}, // ECRecover
		common.BytesToAddress([]byte{2}): {Balance: big.NewInt(1)}, // SHA256
		common.BytesToAddress([]byte{3}): {Balance: big.NewInt(1)}, // RIPEMD
		common.BytesToAddress([]byte{4}): {Balance: big.NewInt(1)}, // Identity
		common.BytesToAddress([]byte{5}): {Balance: big.NewInt(1)}, // ModExp
		common.BytesToAddress([]byte{6}): {Balance: big.NewInt(1)}, // ECAdd
		common.BytesToAddress([]byte{7}): {Balance: big.NewInt(1)}, // ECScalarMul
		common.BytesToAddress([]byte{8}): {Balance: big.NewInt(1)}, // ECPairing
	}
	for k, c := range adds {
		addr[k] = c
	}
	alloc := core.GenesisAlloc(addr)

	return backends.NewSimulatedBackend(alloc, 80000000), transactors, nil
}
