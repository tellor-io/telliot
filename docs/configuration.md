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

* `NODE_URL` \(required\) - websocket node URL \(e.g [wss://mainnet.infura.io/bbbb](wss://mainnet.infura.io/bbbb) or [wss://localhost:8546](ws://localhost:8546) if own node\)


#### Config file options:
```json
{
	"Aggregator": {
		"LogLevel": "info",
		"ManualDataFile": "configs/manualData.json"
	},
	"Db": {
		"LogLevel": "info",
		"Path": "db",
		"RemoteHost": "",
		"RemotePort": 0,
		"RemoteTimeout": {
			"Duration": 5000000000
		}
	},
	"Disputer": {
		"DisputeThreshold": 0,
		"DisputeTimeDelta": {
			"Duration": 0
		},
		"LogLevel": ""
	},
	"EnvFile": "(Required: false)  - Default: configs/.env",
	"Ethereum": {
		"LogLevel": "info",
		"Timeout": 3000
	},
	"IndexTracker": {
		"IndexFile": "configs/index.json",
		"Interval": {
			"Duration": 30000000000
		},
		"LogLevel": "info"
	},
	"Mining": {
		"Address": "0x0000000000000000000000000000000000000000",
		"Heartbeat": 60000000000,
		"LogLevel": "info"
	},
	"ProfitTracker": {
		"LogLevel": "info"
	},
	"PsrTellor": {
		"MinConfidence": 0.7
	},
	"PsrTellorAccess": {
		"MinConfidence": 0
	},
	"SubmitterTellor": {
		"Enabled": true,
		"LogLevel": "info",
		"MinSubmitPeriod": {
			"Duration": 901000000000
		},
		"ProfitThreshold": 0
	},
	"SubmitterTellorAccess": {
		"Enabled": false,
		"LogLevel": "info"
	},
	"Tasker": {
		"LogLevel": "info"
	},
	"Transactor": {
		"GasMax": 10,
		"GasMultiplier": 1,
		"LogLevel": "info"
	},
	"Web": {
		"ListenHost": "",
		"ListenPort": 9090,
		"LogLevel": "info",
		"ReadTimeout": {
			"Duration": 0
		}
	}
}
```
Here are the config defaults in json format:
```json
{
	"Web": {
		"LogLevel": "info",
		"ListenHost": "",
		"ListenPort": 9090,
		"ReadTimeout": "0s"
	},
	"Mining": {
		"LogLevel": "info",
		"Address": "0x0000000000000000000000000000000000000000",
		"Heartbeat": 60000000000
	},
	"SubmitterTellor": {
		"Enabled": true,
		"LogLevel": "info",
		"ProfitThreshold": 0,
		"MinSubmitPeriod": "15m1s"
	},
	"SubmitterTellorAccess": {
		"Enabled": false,
		"LogLevel": "info"
	},
	"ProfitTracker": {
		"LogLevel": "info"
	},
	"Tasker": {
		"LogLevel": "info"
	},
	"Transactor": {
		"LogLevel": "info",
		"GasMax": 10,
		"GasMultiplier": 1
	},
	"IndexTracker": {
		"LogLevel": "info",
		"Interval": "30s",
		"IndexFile": "configs/index.json"
	},
	"Disputer": {
		"LogLevel": "",
		"DisputeTimeDelta": "0s",
		"DisputeThreshold": 0
	},
	"Ethereum": {
		"LogLevel": "info",
		"Timeout": 3000
	},
	"Aggregator": {
		"LogLevel": "info",
		"ManualDataFile": "configs/manualData.json"
	},
	"PsrTellor": {
		"MinConfidence": 0.7
	},
	"PsrTellorAccess": {
		"MinConfidence": 0
	},
	"Db": {
		"LogLevel": "info",
		"Path": "db",
		"RemoteHost": "",
		"RemotePort": 0,
		"RemoteTimeout": "5s"
	},
	"envFile": "configs/.env"
}
```
### Log levels
Note the default level is "INFO", so to turn down the number of logs, enter "WARN" or "ERROR".

DEBUG - logs everything in INFO and additional developer logs

INFO - logs most information about the mining operation

WARN - logs all warnings and errors

ERROR - logs only serious errors
