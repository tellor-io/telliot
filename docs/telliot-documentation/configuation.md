---
description: Telliot tweaks and settings to keep your rig running smoothly.
---

# Configuation

## Cli reference

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
* `stake deposit` \(indicates to deposit 1000 tokens in the contract\)
* `stake request` \(indicates you wish to withdraw your stake\)
* `stake withdraw` \(withdraws your stake, run 1 week after request\)
* `stake status` \(shows your staking balance\)
* `balance` \(shows your balance\)

#### Config file options:

* `contractAddress` \(required\) - address of TellorContract
* `nodeURL` \(required\) - node URL \(e.g [https://mainnet.infura.io/bbbb](https://mainnet.infura.io/bbbb) or [https://localhost:8545](https://localhost:8545) if own node\)
* `privateKey` \(required\) - privateKey for your address
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
* `numProcessors` - an integer number of CPU cores/threads to use for mining. \(cpu mining is disabled if there is a suitable GPU is found
* `disputeTimeDelta` - how far back to store values for min/max range - default 5 \(in minutes\)
* `disputeThreshold` - percentage of acceptable range outside min/max for dispute checking - default
* `psrFolder` - folder location holding your psr.json file, default working directory

## Configuation

### useGPU

If you have one or more GPUs, they will be used for mining by default. Currently only Nvidia cards are supported, and the default behavior will work well for miners.

You can configure your GPUs as you wish by adding / editing the following lines in your config.json file:

```text
"useGPU":true,
        "gpuConfig":{
        "default":{
            "groupSize":256,
            "groups":4096,
            "count":16
        }
    },
```

If you have multiple GPUs connected, you can specify different config settings for each card in the following format:

```text
"gpuConfig":{
        "<GPUName1>":{
            "groupSize":256,
            "groups":4096,
            "count":16
        },
        "<GPUName2>":{
            "groupSize":256,
            "groups":4096,
            "count":16
        },
        "<GPUName3>":{
            "disabled":true
        },

    }
```

Replace "GPUName1" with system name for each card. On most systems you can get a list of GPUs with `nvidia-smi -l`

You may edit the variables to try to improve hashrate for your GPU model:

* disabled: boolean on whether or not to disable a specific GPU
* groupSize - number of groups to split work into
* groups - number of groups of work submitted to the gpu at once
* count: number of hashes each thread executes in one pass

{% hint style="info" %}
Edit these variables at your own risk! Reach out to the community if you're unsure.
{% endhint %}

### Connecting to a Pool

There are mining pools available for mining TRB without staking 500 tokens. The pool server operator stakes 500 tokens for you, and you receive rewards roughly proportional to your hashrate as a fraction of the pool's hashrate.

{% hint style="info" %}
Each pool as different fees and instructions for hooking up. Be sure to read your pools documentation. Feel free to reach out to the community if you need help with mining pools.
{% endhint %}

Add the following lines to your config file:

```text
"enablePoolWorker": true,
"poolURL": "<poolURL>",
```

Where the poolURL is the link to your pool. \(e.g. [http://tellorpool.org](http://tellorpool.org) \)

You can change the job duration if needed. This is the time in seconds to grab information from the pool. The default time is 15 seconds.

```text
"poolJobDuration":10
```

### Running a Remote Data Server for Multiple Miners

{% hint style="info" %}
If you are setting up a Tellor miner for the first time, it might be a good idea to skip this section and come back after you're up and running with one miner.
{% endhint %}

If you are running multiple miners, there is no reason to run multiple databases \(the values you will submit should be identical\). In addition, querying the same API from multiple processes can lead to rate limits on the public API's. To get around this, you can utilize a system where you run one.

`tellor dataServer`

You can then start multiple miners that all read off the one database.

`tellor mine -r`

Create a config file for the dataServer and add the public addresses of all remote miners that want to read the database: \(0x prefix is retained here\)

```text
    “serverWhitelist”: [
        "0xyour-miner-address1",
        "0xyour-miner-address2"
    ]
```

Create a separate copy of config.json for each miner and edit these to match your miner's characteristics.

In both your miner and dataServer config files, you must change the IP addresses to match the server and port you are hosting/reading from.

```text
   "serverHost": "1.2.3.4",
   "serverPort": 5000,
```

Note that your dataServer and miners must be started with separate commands.

For more detailed instructions:[ https://docs.google.com/document/d/1k8ELb1cXkEpztHkHUt8QTL4JCcnHw5\_yQjTKIHCaSCE](https://docs.google.com/document/d/1k8ELb1cXkEpztHkHUt8QTL4JCcnHw5_yQjTKIHCaSCE)

### LogConfig file options

The logging.config file consists of two fields: \* component \* level

The component is the package.component combination.

E.G. the Runner component in the tracker package would be: tracker.Runner

To turn on logging, add the component and the according level. Note the default level is "INFO", so to turn down the number of logs, enter "WARN" or "ERROR"

DEBUG - logs everything in INFO and additional developer logs

INFO - logs most information about the mining operation

WARN - logs all warnings and errors

ERROR - logs only serious errors

