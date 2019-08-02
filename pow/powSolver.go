package pow

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/apexskier/cryptoPadding"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	tellor1 "github.com/tellor-io/TellorMiner/contracts1"
	"github.com/tellor-io/TellorMiner/rpc"
	"golang.org/x/crypto/ripemd160"
)

//PoWSolver state for mining operation
type PoWSolver struct {
	canMine bool
	mining  bool
}

func randInt() string {
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(max, big.NewInt(1))

	//Generate cryptographically strong pseudo-random between 0 - max
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		//error handling
	}
	//n := big.NewInt(9000)
	//String representation of n in base 16
	nonce := fmt.Sprintf("%x", n.String()) //n.Text(16)
	return nonce
}

func decodeHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return b
}

//CreateMiner creates a new miner instance
func CreateMiner() *PoWSolver {
	return &PoWSolver{canMine: true, mining: false}
}

//SolveChallenge performs PoW
func (p *PoWSolver) SolveChallenge(challenge []byte, _difficulty *big.Int) string {
	var padding cryptoPadding.ZeroPadding
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	if !p.canMine {
		return ""
	}
	p.mining = true
	defer func() {
		p.mining = false
	}()
	fmt.Println("thisChallenge", fmt.Sprintf("%x", challenge))
	fmt.Println("Solving for difficulty: ", _difficulty)
	for i := 0; i < 100000000; i++ {
		if !p.canMine {
			return ""
		}

		nonce := randInt() //do we need to use big number?
		_string := fmt.Sprintf("%x", challenge) + cfg.PublicAddress + nonce
		hash := solsha3.SoliditySHA3(
			solsha3.Bytes32(decodeHex(_string)),
		)
		hasher := ripemd160.New()
		hasher.Write([]byte(hash))
		hash1 := hasher.Sum(nil)
		paddedData, _ := padding.Pad(hash1, 32)
		n := sha256.Sum256(paddedData)
		q := fmt.Sprintf("%x", n)
		p := new(big.Int)
		p, ok := p.SetString(q, 16)
		if !ok {
			fmt.Println("SetString: error")
			return ""
		}
		x := new(big.Int)
		x.Mod(p, _difficulty)
		if x.Cmp(big.NewInt(0)) == 0 {
			fmt.Println("Solution Found", nonce)
			return nonce
		}
	}
	return ""
}

//Stop mining operations
func (p *PoWSolver) Stop() {
	p.canMine = false
}

//IsMining checks whether the miner is currently working on a PoW
func (p *PoWSolver) IsMining() bool {
	return p.mining
}

//SubmitSolution signs transaction and submits on-chain
func SubmitSolution(ctx context.Context, solution string, value, requestId *big.Int) error {
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}
	client := ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)

	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		return err
	}

	cost := new(big.Int)
	cost.Mul(gasPrice, big.NewInt(700000))
	if balance.Cmp(cost) < 0 {
		//FIXME: notify someone that we're out of funds!
		return fmt.Errorf("Insufficient funds to send transaction: %v < %v", balance, cost)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	instance := ctx.Value(tellorCommon.TransactorContractContextKey).(*tellor1.TellorTransactor)

	fmt.Printf("Calling contract with vars: %v, %v, %v, %v\n", auth, solution, requestId, value)

	tx, err := instance.SubmitMiningSolution(auth, solution, requestId, value)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())

	return nil
}

func SubmitTransaction(solution string, value, requestId *big.Int) error {

	//get the single config instance
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
		return err
	}

	client, err := ethclient.Dial(cfg.NodeURL)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	cost := new(big.Int)
	cost.Mul(gasPrice, big.NewInt(700000))
	if balance.Cmp(cost) >= 0 {
		auth := bind.NewKeyedTransactor(privateKey)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)      // in wei
		auth.GasLimit = uint64(3000000) // in units
		auth.GasPrice = gasPrice

		contractAddress := common.HexToAddress(cfg.ContractAddress)
		instance, err := tellor1.NewTellorTransactor(contractAddress, client)
		if err != nil {
			log.Fatal(err)
		}
		solution := string(solution)
		fmt.Println(auth, solution, requestId, value)
		tx, err := instance.SubmitMiningSolution(auth, solution, requestId, value)
		if err != nil {
			log.Fatal(err)
			return err
		}

		fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	}
	return nil
}
