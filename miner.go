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

*/

package main


/*Hardcoded Variables (in config file)
	Contract Address to submit to
	Node URL (your own node or infura)
	Network ID (1 if mainnet)
	public key
	private key
	database url (our public one or localhost)
*/

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

func solveChallenge(){
	//intermittently check if challenge has changed

	return solution

}

func getRequestedValues(){

}

func submitTransaction(){
	//pull current gas price
	//check if you have enough gas (warn if low)
	//submit transaction
}

func main() {
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