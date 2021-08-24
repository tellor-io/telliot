package contracts

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/contracts/balancer"
	"github.com/tellor-io/telliot/pkg/contracts/lens"
	"github.com/tellor-io/telliot/pkg/contracts/tellor"
	"github.com/tellor-io/telliot/pkg/contracts/tellorMesosphere"
	"github.com/tellor-io/telliot/pkg/contracts/uniswap"
)

const (
	TellorAddress                          = "0x88dF592F8eb5D7Bd38bFeF7dEb0fBc02cf3778a0"
	TellorAddressGoerli                    = "0xA0238859b8626cCd6981438200Eded52F05dB37A"
	TellorAddressHardhat                   = "0x8920050E1126125a27A4EaC5122AD3586c056E51"
	TellorMesosphereAddressRinkeby         = "0xB2a25FD022526c64823FF1bF03bf348Fd0787f2a"
	TellorMesosphereAddressArbitrumTestnet = "0x7A1e398A228271D1B8b1fb1ede678A3e4c79f50A"
	TellorMesosphereAddressPolygonTestnet  = "0x32704dCEb8dA339516f4AE561Cd40a6cBE6d98c9"
	TellorMesosphereAddressPolygonMainnet  = "0xACC2d27400029904919ea54fFc0b18Bf07C57875"
	TellorMesosphereAddress                = "0x5a991dd4f646ed7efdd090b1ba5b68d222273f7e"
	LensAddressMainnet                     = "0x577417CFaF319a1fAD90aA135E3848D2C00e68CF"
	LensAddressRinkeby                     = "0xebEF7ceB7C43850898e258be0a1ea5ffcdBc3205"
	LensAddressHardhat                     = "0x577417CFaF319a1fAD90aA135E3848D2C00e68CF"
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

type ITellorMesosphere struct {
	Address common.Address
	*tellorMesosphere.TellorMesosphere
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

func NewITellorMesosphere(client *ethclient.Client) (*ITellorMesosphere, error) {
	conractAddr, err := GetTellorMesosphereAddress(client)
	if err != nil {
		return nil, errors.Wrap(err, "getting contract address")
	}
	tellorInstance, err := tellorMesosphere.NewTellorMesosphere(conractAddr, client)
	if err != nil {
		return nil, errors.Wrap(err, "creating telllor interface")
	}

	return &ITellorMesosphere{Address: common.HexToAddress(TellorMesosphereAddress), TellorMesosphere: tellorInstance}, nil
}

func GetTellorMesosphereAddress(client *ethclient.Client) (common.Address, error) {
	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		return common.Address{}, err
	}
	switch netID := networkID.Int64(); netID {
	case 421611:
		return common.HexToAddress(TellorMesosphereAddressArbitrumTestnet), nil
	case 4:
		return common.HexToAddress(TellorMesosphereAddressRinkeby), nil
	case 80001:
		return common.HexToAddress(TellorMesosphereAddressPolygonTestnet), nil
	case 137:
		return common.HexToAddress(TellorMesosphereAddressPolygonMainnet), nil
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
	case 31337:
		return common.HexToAddress(TellorAddressHardhat), nil
	case 4: // Rinkeby has the same address as mainnet.
		fallthrough
	case 1:
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
	case 31337:
		return common.HexToAddress(LensAddressHardhat), nil
	case 1:
		return common.HexToAddress(LensAddressMainnet), nil
	case 4:
		return common.HexToAddress(LensAddressRinkeby), nil
	default:
		return common.Address{}, errors.Errorf("contract address for current network id not found:%v", netID)
	}
}
