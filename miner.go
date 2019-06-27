/*
Tellor Miner

This is the workhorse of the Miner system as it takes on solving the PoW challenge.  
As for the feel of this piece, we are going to follow the Geth model to keep a familiar interface for users. 

The main goals for this build are:
	Solve PoW
	Submit Transaction to the Ethereum network

Other features that it should also have to properly do these main tasks:
	Grab challenge details from database
	Grab request data from database
	Optimize transaction details (gas price)
	Pull current gas price from database
	Handle transaction failure when submitting to Ethereum
	Show low gas warning (more ETH needed for gas)
	IPC interface for interacting with Miner (Geth model)

Resources:
https://goethereumbook.org/client-setup/

*/

package main

import (
    "fmt"
    "log"
    "math/rand"
    "time"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    tellor "./contracts" // for demo
)

//Variables - eventually put in config file
const contractAddress bytes = "0x..."
const nodeURL string = "http://localhost:8545" // or "https://mainnet.infura.io"
const networkID uint8 = 1
const publicKey bytes = "0x..."
const privateKey bytes = "0x..."
const databaseURL string = "http://localhost7545"
const abi string = json.abi;


/*Variables to Grab
	Current Challenge
	Current Request ID
	Current Difficulty
	Current Gas Price
	Requested Data
	Current Ethereum balance (check if you have enough)
*/

func getCurrentChallenge(){

	return currentChallenge, difficulty, requestId

}

func solveChallenge(challenge bytes, difficulty uint32) uin32{
	//intermittently check if challenge has changed
	solution := 0
	i := 0 //
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

	for solution = 0 && i < 100000{
		v = Web3.toHex(Web3.sha3(hexstr=_string));
		z= hashlib.new('ripemd160',bytes.fromhex(v[2:])).hexdigest()
		n = "0x" + hashlib.new('sha256',bytes.fromhex(z)).hexdigest()
		hash1 = int(n,16);
		if hash1 % difficulty == 0{
			return solution
		}
		i++
	}

	return solution

}

func getRequestedValues() (bool,uint) {

	return true,value

}

func checkGasPriceAndBalance() (bool,uint){
	myBalance := 0;
	gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
		  log.Fatal(err)
		}
	if(myBalance > gasPrice * transactionCost){
		return true, gasPrice
	}
	else{
		return false,0
	}
}

func submitTransaction(solution uint, requestId uint){
	//pull current gas price
	//check if you have enough gas (warn if low)
	//submit transaction
	goodToSubmit,gasCost := checkGasPriceAndBalance()
	if goodToSubmit {
		goodToSubmit,value = process.argv[4] - 0
		if goodToSubmit{
		    client, err := ethclient.Dial("https://rinkeby.infura.io")
		    if err != nil {
		        log.Fatal(err)
		    }

		    privateKey, err := crypto.HexToECDSA(privateKey)
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

		    auth := bind.NewKeyedTransactor(privateKey)
		    auth.Nonce = big.NewInt(int64(nonce))
		    auth.Value = big.NewInt(0)     // in wei
		    auth.GasLimit = uint64(300000) // in units
		    auth.GasPrice = gasCost

		    address := common.HexToAddress(contractAddress)
		    instance, err := tellor.submitMiningSolution(solution,requestId,value)
		    if err != nil {
		        log.Fatal(err)
		    }

		    key := [32]byte{}
		    value := [32]byte{}
		    copy(key[:], []byte("foo"))
		    copy(value[:], []byte("bar"))

		    tx, err := instance.SetItem(auth, key, value)
		    if err != nil {
		        log.Fatal(err)
		    }

		    fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

		    result, err := instance.Items(nil, key)
		    if err != nil {
		        log.Fatal(err)
		    }

		    fmt.Println(string(result[:])) // "bar"

		}

	}

}

func main() {
	client, err := ethclient.Dial(nodeURL)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("we have a connection")
    _ = client // we'll use this in the upcoming sections

	var nonce uint
	for {
		getCurrentChallenge()
		nonce = solveChallenge()
		if nonce > 0 {
			getRequestedValues()
			submitTransaction()
			nonce = 0
		}
	}
}

package main

import (
    "fmt"
    "log"



)
