---
description: Telliot tweaks and settings to keep your rig running smoothly.
---

# Configuration reference

## CLI reference

Telliot commands and config file options are as the following:

#### Required Flags <a id="docs-internal-guid-d1a57725-7fff-a753-9236-759dd3f42eed"></a>

* `--config` \(path to your config file.\)

#### Telliot Commands


* `accounts`   \(Show accounts\)


* `approve` \(Address\)\(Amount\)\(Account\(optional\)\)  \(Approve tokens\)


* `balance` \(Address\(optional\)\)  \(Check the balance of an address\)


* `dataserver`   \(launch only a dataserver instance\)


* `dispute`   \(Perform commands related to disputes\)


* `dispute new` \(requestId: the request id to dispute it\)\(timestamp: the submitted timestamp to dispute\)\(minerIndex: the miner index to dispute\)\(Account\(optional\)\)  \(start a new dispute\)


* `dispute show` \(Account\(optional\)\)  \(show open disputes\)


* `dispute vote` \(disputeId: the dispute id\)\(support: true or false\)\(Account\(optional\)\)  \(vote on a open dispute\)


* `migrate`   \(Migrate funds from the old oracle contract\)


* `mine`   \(mine TRB and submit values\)


* `stake`   \(Perform one of the stake operations\)


* `stake deposit` \(Account\(optional\)\)  \(deposit a stake\)


* `stake request` \(Account\(optional\)\)  \(request to withdraw stake\)


* `stake status` \(Account\(optional\)\)  \(show stake status\)


* `stake withdraw` \(Address\)\(Account\(optional\)\)  \(withdraw stake\)


* `transfer` \(Address\)\(Amount\)\(Account\(optional\)\)  \(Transfer tokens\)


* `version`   \(Show the CLI version information\)

#### .env file options:


* `ETH_PRIVATE_KEYS` \(required\) - list of private keys separated by `,`

* `NODE_WEBSOCKET_URL` \(required\) - node URL \(e.g [wss://mainnet.infura.io/bbbb](wss://mainnet.infura.io/bbbb) or [wss://localhost:8546](ws://localhost:8546) if own node\)


#### Config file options:

* `ApiFile` -  \(default: configs/api.json\) -

* `DBFile` -  \(default: db\) -

* `DataServer.ListenHost` -  \(default: localhost\) -

* `DataServer.ListenPort` -  \(default: 5000\) -

* `EnvFile` -  \(default: configs/.env\) -

* `EthClientTimeout` -  \(default: 3000\) -

* `GasMax` -  \(default: 10\) -

* `GasMultiplier` -  \(default: 1\) -

* `HistoryFile` -  \(default: configs/saved.json\) -

* `Logger` -  \(default: map[apiOracle:info dataServer:info db:info ops:info pow::info rest:info rpc:info tracker:info]\) -

* `ManualDataFile` -  \(default: configs/manualData.json\) -

* `Mine.Heartbeat.Duration` -  \(default: 15s\) -

* `Mine.ListenHost` -  \(default: localhost\) -

* `Mine.ListenPort` -  \(default: 9090\) -

* `Mine.MinSubmitPeriod.Duration` -  \(default: 15m1s\) -

* `Mine.MiningInterruptCheckInterval.Duration` -  \(default: 15s\) -

* `Mine.ProfitThreshold` -

* `Mine.RemoteDBHost` -

* `Mine.RemoteDBPort` -

* `ServerWhitelist` -

* `Trackers.DisputeThreshold` -  \(default: 0.01\) -

* `Trackers.DisputeTimeDelta.Duration` -  \(default: 5m0s\) -

* `Trackers.FetchTimeout.Duration` -  \(default: 30s\) -

* `Trackers.MinConfidence` -  \(default: 0.2\) -

* `Trackers.Names` -  \(default: map[disputeChecker:false gas:true indexers:true]\) -

* `Trackers.SleepCycle.Duration` -  \(default: 30s\) -

### LogConfig file options

The logging.config file consists of two fields: \* component \* level

The component is the package.component combination.

E.G. the Runner component in the tracker package would be: tracker.Runner

To turn on logging, add the component and the according level. Note the default level is "INFO", so to turn down the number of logs, enter "WARN" or "ERROR"

DEBUG - logs everything in INFO and additional developer logs

INFO - logs most information about the mining operation

WARN - logs all warnings and errors

ERROR - logs only serious errors
