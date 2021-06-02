---
description: Telliot tweaks and settings to keep your rig running smoothly.
---

# Configuration reference

## CLI reference

Telliot commands and config file options are as the following:

#### Required Flags <a id="docs-internal-guid-d1a57725-7fff-a753-9236-759dd3f42eed"></a>

* `--config` \(path to your config file.\)

#### Telliot Commands

* `accounts`
```go
Usage: telliot accounts

Show accounts

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `approve`
```go
Usage: telliot approve <address> <amount> [<account>]

Approve tokens

Arguments:
  <address>
  <amount>
  [<account>]

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `balance`
```go
Usage: telliot balance [<address>]

Check the balance of an address

Arguments:
  [<address>]

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `dataserver`
```go
Usage: telliot dataserver

launch only a dataserver instance

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `dispute`
```go
Usage: telliot dispute <command>

Perform commands related to disputes

Flags:
  -h, --help    Show context-sensitive help.

Commands:
  dispute new [<account>]
    start a new dispute

  dispute vote [<account>]
    vote on a open dispute

  dispute show [<account>]
    show open disputes

```

* `dispute new`
```go
Usage: telliot dispute new [<account>]

start a new dispute

Arguments:
  [<account>]

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `dispute show`
```go
Usage: telliot dispute show [<account>]

show open disputes

Arguments:
  [<account>]

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `dispute vote`
```go
Usage: telliot dispute vote [<account>]

vote on a open dispute

Arguments:
  [<account>]

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `migrate`
```go
Usage: telliot migrate

Migrate funds from the old oracle contract

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `mine`
```go
Usage: telliot mine

mine TRB and submit values

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `stake`
```go
Usage: telliot stake <command>

Perform one of the stake operations

Flags:
  -h, --help    Show context-sensitive help.

Commands:
  stake deposit [<account>]
    deposit a stake

  stake request [<account>]
    request to withdraw stake

  stake withdraw <address> [<account>]
    withdraw stake

  stake status [<account>]
    show stake status

```

* `stake deposit`
```go
Usage: telliot stake deposit [<account>]

deposit a stake

Arguments:
  [<account>]

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `stake request`
```go
Usage: telliot stake request [<account>]

request to withdraw stake

Arguments:
  [<account>]

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `stake status`
```go
Usage: telliot stake status [<account>]

show stake status

Arguments:
  [<account>]

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `stake withdraw`
```go
Usage: telliot stake withdraw <address> [<account>]

withdraw stake

Arguments:
  <address>
  [<account>]

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `transfer`
```go
Usage: telliot transfer <address> <amount> [<account>]

Transfer tokens

Arguments:
  <address>
  <amount>
  [<account>]

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `version`
```go
Usage: telliot version

Show the CLI version information

Flags:
  -h, --help    Show context-sensitive help.

```

#### .env file options:


* `ETH_PRIVATE_KEYS` \(required\) - list of private keys separated by `,`

* `NODE_WEBSOCKET_URL` \(required\) - node URL \(e.g [wss://mainnet.infura.io/bbbb](wss://mainnet.infura.io/bbbb) or [wss://localhost:8546](ws://localhost:8546) if own node\)


#### Config file options:
```json
{
	"ApiFile": "(Required: false)  - Default: configs/api.json",
	"DBFile": "(Required: false)  - Default: db",
	"DataServer": {
		"ListenHost": "localhost",
		"ListenPort": 5000
	},
	"EnvFile": "(Required: false)  - Default: configs/.env",
	"EthClientTimeout": "(Required: false)  - Default: 3000",
	"GasMax": "(Required: false)  - Default: 10",
	"GasMultiplier": "(Required: false)  - Default: 1",
	"HistoryFile": "(Required: false)  - Default: configs/saved.json",
	"Logger": "(Required: false)  - Default: map[apiOracle:info dataServer:info db:info ops:info pow::info rest:info rpc:info tracker:info]",
	"ManualDataFile": "(Required: false)  - Default: configs/manualData.json",
	"Mine": {
		"Heartbeat": {
			"Duration": 15000000000
		},
		"ListenHost": "localhost",
		"ListenPort": 9090,
		"MinSubmitPeriod": {
			"Duration": 901000000000
		},
		"MiningInterruptCheckInterval": {
			"Duration": 15000000000
		},
		"ProfitThreshold": 0,
		"RemoteDBHost": "",
		"RemoteDBPort": 0
	},
	"ServerWhitelist": "(Required: false)  - Default: []",
	"Trackers": {
		"DisputeThreshold": 0.01,
		"DisputeTimeDelta": {
			"Duration": 300000000000
		},
		"FetchTimeout": {
			"Duration": 30000000000
		},
		"MinConfidence": 0.2,
		"Names": {
			"disputeChecker": false,
			"gas": true,
			"indexers": true
		},
		"SleepCycle": {
			"Duration": 30000000000
		}
	}
}
```
So the default config is as follows:
```json
{
	"ApiFile": "configs/api.json",
	"DBFile": "db",
	"EnvFile": "configs/.env",
	"EthClientTimeout": 3000,
	"GasMax": 10,
	"GasMultiplier": 1,
	"HistoryFile": "configs/saved.json",
	"Logger": {
		"apiOracle": "info",
		"dataServer": "info",
		"db": "info",
		"ops": "info",
		"pow:": "info",
		"rest": "info",
		"rpc": "info",
		"tracker": "info"
	},
	"ManualDataFile": "configs/manualData.json",
	"ServerWhitelist": null
}
```
### LogConfig file options

The logging.config file consists of two fields: \* component \* level

The component is the package.component combination.

E.G. the Runner component in the tracker package would be: tracker.Runner

To turn on logging, add the component and the according level. Note the default level is "INFO", so to turn down the number of logs, enter "WARN" or "ERROR"

DEBUG - logs everything in INFO and additional developer logs

INFO - logs most information about the mining operation

WARN - logs all warnings and errors

ERROR - logs only serious errors
