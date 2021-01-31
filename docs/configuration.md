---
description: Telliot tweaks and settings to keep your rig running smoothly.
---

# Configuration reference

## CLI reference

Telliot commands and config file options are as the following:

#### Required Flags <a id="docs-internal-guid-d1a57725-7fff-a753-9236-759dd3f42eed"></a>

* `--config` \(path to your config file.\)

#### Telliot Commands

* `--logConfig` \(location of logging config file; default path is current directory\)
* `mine` \(indicates to run the miner\)
* `mine -r` \(indicates to mine utilizing a remote server\)
* `dataserver` \(indicates to run the dataServer \(no mining\)\)
* `transfer` \(AMOUNT\) \(TOADDRESS\) \(indicates transfer, toAddress is Ethereum address and amount is number of Tributes \(eg. transfer 10 0xea... \(this transfers 10 tokens\)\)\)
* `approve` \(AMOUNT\) \(TOADDRESS\) \(ammount to approve the toaddress to send this amount of tokens
* `stake deposit` \(indicates to deposit tokens in the contract\)
* `stake request` \(indicates you wish to withdraw your stake\)
* `stake withdraw` \(withdraws your stake, run 1 week after request\)
* `stake status` \(shows your staking balance\)
* `balance` \(shows your balance\)

#### .env file options:

* `NODE_URL` \(required\) - node URL \(e.g [https://mainnet.infura.io/bbbb](https://mainnet.infura.io/bbbb) or [https://localhost:8545](https://localhost:8545) if own node\)
* `ETH_PRIVATE_KEY` \(required\) - privateKey for your address
* `$PSR$_KEY` - API key for getting a specific indexes.json api \(required if you use authenticated API's\)

#### Config file options:

* `databaseURL` \(required\) - where you are reading from for the server database \(if hosted\)
* `publicAddress` \(required\) - public address for your miner \(note, no 0x\)
* `ethClientTimeout` \(required\) - timeout for making requests from your node
* `trackerCycle` \(required\) - how often your database updates \(in seconds\)
* `trackers` \(required\) - which pieces of the database you update
* `dbFile` \(required\) - where you want to store your local database \(if self-hosting\)
* `serverHost` \(required\) - location to host server
* `serverWhitelist` \(required\) - whitelists which publicAddress can access the data server
* `fetchTimeout` - timeout for requesting data from an API
* `requestData` - sets wether your miner request data if challenge is 0.  If yes, then you will addTip\(\) to this number.  Enter a uint number representing request id to be requested \(e.g. 2\)
* `requestDataInterval` - min frequency at which to request data at \(in seconds, default 30\)
* `gasMultiplier` - Multiplies the submitted gasPrice \(e.g. 2 will double gas costs\)
* `gasMax` - a max for the gas price in gwei \(note: this max comes BEFORE the gas multiplier.  So a max gas cost of 10 gwei, can have gas prices up to 20 if gasMultiplier is 2\)
* `heartbeat` - an integer that controls how frequently the miner process should report the hashrate \(larger is less frequent, try 1000000 to start\)
* `numProcessors` - an integer number of CPU cores/threads to use for mining.
* `disputeTimeDelta` - how far back to store values for min/max range - default 5 \(in minutes\)
* `disputeThreshold` - percentage of acceptable range outside min/max for dispute checking - default
* `psrFolder` - folder location holding your psr.json file, default working directory


### LogConfig file options

The logging.config file consists of two fields: \* component \* level

The component is the package.component combination.

E.G. the Runner component in the tracker package would be: tracker.Runner

To turn on logging, add the component and the according level. Note the default level is "INFO", so to turn down the number of logs, enter "WARN" or "ERROR"

DEBUG - logs everything in INFO and additional developer logs

INFO - logs most information about the mining operation

WARN - logs all warnings and errors

ERROR - logs only serious errors

