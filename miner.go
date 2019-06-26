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
)

//Variables - eventually put in config file
const contractAddress bytes = "0x..."
const nodeURL string = "http://localhost:8545" // or "https://mainnet.infura.io"
const networkID uint8 = 1
const publicKey bytes = "0x..."
const privateKey bytes = "0x..."
const databaseURL string = "http://localhost7545"


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

func getRequestedValues(){

}

func submitTransaction(solution uint, requestId uint){
	//pull current gas price
	//check if you have enough gas (warn if low)
	//submit transaction
	value = process.argv[4] - 0


	var address = process.argv[5];
	var abi = json.abi;
	var account = process.argv[6];
	var privateKey = new Buffer(process.argv[7], 'hex');

	let myContract = new web3.eth.Contract(abi,address);
	let data = myContract.methods.submitMiningSolution(solution,requestId,value).encodeABI();

	//web3.eth.sendTransaction({to: oracle.address,from:accounts[0],gas:7000000,data:oracle2.methods.requestData(api,0,0).encodeABI()});

	web3.eth.getTransactionCount(account, function (err, nonce) {
	     var tx = new Tx({
	      nonce: nonce,
	      gasPrice: web3.utils.toHex(web3.utils.toWei('20', 'gwei')),
	      gasLimit: 2000000,
	      to: address,
	      value: 0,
	      data: data,
	    });
	    tx.sign(privateKey);

	    var raw = '0x' + tx.serialize().toString('hex');
	    web3.eth.sendSignedTransaction(raw).on('transactionHash', function (txHash) {
	      }).on('receipt', function (receipt) {
	          //console.log("receipt:" + receipt);
	      }).on('confirmation', function (confirmationNumber, receipt) {
	          //console.log("confirmationNumber:" + confirmationNumber + " receipt:" + receipt);
	      }).on('error', function (error) {
	    });
	  });
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