package contracts

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/contracts/balancer"
	"github.com/tellor-io/telliot/pkg/contracts/lens"
	"github.com/tellor-io/telliot/pkg/contracts/tellor"
	"github.com/tellor-io/telliot/pkg/contracts/tellorAccess"
	"github.com/tellor-io/telliot/pkg/contracts/uniswap"
)

const (
	TellorAddress                      = "0x88dF592F8eb5D7Bd38bFeF7dEb0fBc02cf3778a0"
	TellorAddressGoerli                = "0x90bbE2155deb3d454696Ce659B52B831e754431C"
	TellorAccessAddressRinkeby         = "0x5a991dd4f646ed7efdd090b1ba5b68d222273f7e"
	TellorAccessAddressArbitrumTestnet = "0x7A1e398A228271D1B8b1fb1ede678A3e4c79f50A"
	TellorAccessAddress                = "0x5a991dd4f646ed7efdd090b1ba5b68d222273f7e"
	LensAddressMainnet                 = "0x577417CFaF319a1fAD90aA135E3848D2C00e68CF"
	LensAddressRinkeby                 = "0xebEF7ceB7C43850898e258be0a1ea5ffcdBc3205"
)

type (
	ITellorNewDispute    = tellor.ITellorNewDispute
	TellorNonceSubmitted = tellor.TellorNonceSubmitted
)

const (
	BPoolABI          = balancer.BPoolABI
	BTokenABI         = balancer.BTokenABI
	IERC20ABI         = uniswap.IERC20ABI
	IUniswapV2PairABI = uniswap.IUniswapV2PairABI
	ITellorABI        = tellor.ITellorABI
)

type ITellorAccess struct {
	Address common.Address
	*tellorAccess.TellorAccess
}

type ITellor struct {
	*tellor.ITellor
	*tellor.ITellorNewDispute
	*lens.Main
	Address common.Address
}

func NewITellor(client *ethclient.Client) (*ITellor, error) {
	conractAddr, err := GetTellorAddress(client)
	if err != nil {
		return nil, errors.Wrap(err, "getting contract address")
	}
	tellorInstance, err := tellor.NewITellor(conractAddr, client)
	if err != nil {
		return nil, errors.Wrap(err, "creating contract interface")
	}
	contractAddr, err := GetLensAddress(client)
	if err != nil {
		return nil, errors.Wrap(err, "getting contract address")
	}

	lensInstance, err := lens.NewMain(contractAddr, client)
	if err != nil {
		return nil, errors.Wrap(err, "creating telllor interface")
	}

	return &ITellor{Address: common.HexToAddress(TellorAddress), ITellor: tellorInstance, Main: lensInstance}, nil
}

func NewITellorAccess(client *ethclient.Client) (*ITellorAccess, error) {
	conractAddr, err := GetTellorAccessAddress(client)
	if err != nil {
		return nil, errors.Wrap(err, "getting contract address")
	}
	tellorInstance, err := tellorAccess.NewTellorAccess(conractAddr, client)
	if err != nil {
		return nil, errors.Wrap(err, "creating telllor interface")
	}

	return &ITellorAccess{Address: common.HexToAddress(TellorAccessAddress), TellorAccess: tellorInstance}, nil
}

func GetTellorAccessAddress(client *ethclient.Client) (common.Address, error) {
	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		return common.Address{}, err
	}
	switch netID := networkID.Int64(); netID {
	case 421611:
		return common.HexToAddress(TellorAccessAddressArbitrumTestnet), nil
	case 4:
		return common.HexToAddress(TellorAccessAddressRinkeby), nil
	default:
		return common.Address{}, errors.Errorf("contract address for current network id not found:%v", netID)
	}
}

func GetTellorAddress(client *ethclient.Client) (common.Address, error) {
	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		return common.Address{}, err
	}

	switch netID := networkID.Int64(); netID {
	case 1:
		return common.HexToAddress(TellorAddress), nil
	case 4:
		return common.HexToAddress(TellorAddress), nil
	case 5:
		return common.HexToAddress(TellorAddressGoerli), nil
	default:
		return common.Address{}, errors.Errorf("network id not supported id:%v", netID)
	}
}

func GetLensAddress(client *ethclient.Client) (common.Address, error) {
	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		return common.Address{}, err
	}
	switch netID := networkID.Int64(); netID {
	case 1:
		return common.HexToAddress(LensAddressMainnet), nil
	case 4:
		return common.HexToAddress(LensAddressRinkeby), nil
	default:
		return common.Address{}, errors.Errorf("contract address for current network id not found:%v", netID)
	}
}
