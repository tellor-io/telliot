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

```
Usage: telliot accounts

Show accounts

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `approve`

```
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

```
Usage: telliot balance [<address>]

Check the balance of an address

Arguments:
  [<address>]

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `dataserver`

```
Usage: telliot dataserver

launch only a dataserver instance

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `dispute`

```
Usage: telliot dispute <command>

Perform commands related to disputes

Flags:
  -h, --help    Show context-sensitive help.

Commands:
  dispute new [<account>]
    start a new dispute

  dispute vote [<account>]
    vote on a open dispute

  dispute list [<account>]
    list open disputes

```

* `dispute list`

```
Usage: telliot dispute list [<account>]

list open disputes

Arguments:
  [<account>]

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `dispute new`

```
Usage: telliot dispute new [<account>]

start a new dispute

Arguments:
  [<account>]

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `dispute vote`

```
Usage: telliot dispute vote [<account>]

vote on a open dispute

Arguments:
  [<account>]

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `mine`

```
Usage: telliot mine

Submit data to oracle contracts

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `stake`

```
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

```
Usage: telliot stake deposit [<account>]

deposit a stake

Arguments:
  [<account>]

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `stake request`

```
Usage: telliot stake request [<account>]

request to withdraw stake

Arguments:
  [<account>]

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `stake status`

```
Usage: telliot stake status [<account>]

show stake status

Arguments:
  [<account>]

Flags:
  -h, --help                  Show context-sensitive help.

      --config=CONFIG-PATH    path to config file

```

* `stake withdraw`

```
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

```
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

```
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
		"LogLevel": "(Required: false)  - Default: info",
		"ManualDataFile": "(Required: false)  - Default: configs/manualData.json"
	},
	"Db": {
		"LogLevel": "(Required: false)  - Default: info",
		"Path": "(Required: false)  - Default: db",
		"RemoteHost": "(Required: false)  - Default: ",
		"RemotePort": "(Required: false)  - Default: 0",
		"RemoteTimeout": {
			"Duration": "(Required: false)  - Default: 5s"
		}
	},
	"DisputeTracker": {
		"LogLevel": "(Required: false)  - Default: info"
	},
	"IndexTracker": {
		"IndexFile": "(Required: false)  - Default: configs/index.json",
		"Interval": {
			"Duration": "(Required: false)  - Default: 30s"
		},
		"LogLevel": "(Required: false)  - Default: info"
	},
	"Mining": {
		"Heartbeat": "(Required: false)  - Default: 1m0s",
		"LogLevel": "(Required: false)  - Default: info"
	},
	"ProfitTracker": {
		"LogLevel": "(Required: false)  - Default: info"
	},
	"PsrTellor": {
		"MinConfidence": "(Required: false)  - Default: 70"
	},
	"PsrTellorAccess": {
		"MinConfidence": "(Required: false)  - Default: 0"
	},
	"SubmitterTellor": {
		"Enabled": "(Required: false)  - Default: true",
		"LogLevel": "(Required: false)  - Default: info",
		"MinSubmitPeriod": {
			"Duration": "(Required: false)  - Default: 15m1s"
		},
		"ProfitThreshold": "(Required: false)  - Default: 0"
	},
	"SubmitterTellorAccess": {
		"Enabled": "(Required: false)  - Default: false",
		"LogLevel": "(Required: false)  - Default: info"
	},
	"Tasker": {
		"LogLevel": "(Required: false)  - Default: info"
	},
	"Transactor": {
		"GasMax": "(Required: false)  - Default: 10",
		"GasMultiplier": "(Required: false)  - Default: 1",
		"LogLevel": "(Required: false)  - Default: info"
	},
	"Web": {
		"ListenHost": "(Required: false)  - Default: ",
		"ListenPort": "(Required: false)  - Default: 9090",
		"LogLevel": "(Required: false)  - Default: info",
		"ReadTimeout": {
			"Duration": "(Required: false)  - Default: 0s"
		}
	},
	"envFile": "(Required: false)  - Default: configs/.env"
}
```
Here are the config defaults in json format:
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
		"RemoteTimeout": "5s"
	},
	"DisputeTracker": {
		"LogLevel": "info"
	},
	"IndexTracker": {
		"IndexFile": "configs/index.json",
		"Interval": "30s",
		"LogLevel": "info"
	},
	"Mining": {
		"Heartbeat": 60000000000,
		"LogLevel": "info"
	},
	"ProfitTracker": {
		"LogLevel": "info"
	},
	"PsrTellor": {
		"MinConfidence": 70
	},
	"PsrTellorAccess": {
		"MinConfidence": 0
	},
	"SubmitterTellor": {
		"Enabled": true,
		"LogLevel": "info",
		"MinSubmitPeriod": "15m1s",
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
		"ReadTimeout": "0s"
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
