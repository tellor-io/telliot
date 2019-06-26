# Tellor Miner

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
