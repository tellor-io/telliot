---
description: Here are the nuts and bolts for mining TRB.
---

# The Guide

## Download the Latest TellorMiner

This is the workhorse of the Miner system.

[https://github.com/tellor-io/TellorMiner/releases](https://github.com/tellor-io/TellorMiner/releases)

```text
wget https://github.com/tellor-io/TellorMiner/releases/[release-num]/download/tellor
```

Depending on your miner setup you may need to give TellorMiner permission to run. If so, this will need to be done after updates as well.

```text
chmod +x tellor
```

## Download and Edit config.json

config.json is where you will enter your wallet address and configure TellorMiner for your machine.

```text
wget https://raw.githubusercontent.com/tellor-io/TellorMiner/master/configs/config.json
```

Open config.json and update the following values:

* Set `"nodeURL"` to an Ethereum node endpoint. \(e.g. Infura API endpoint\)
* Set `"publicAddress"` to the public key for the Ethereum wallet you plan to use for mining. Remove the 0x prefix at the beginning of the address.

## Create .env file

Create a file named `configs/.env` \(Note: This step can be skipped in you plan to mine on a pool.\)

Copy and paste the following into your `.env` file, and edit this to match your mining address private key.

```text
ETH_PRIVATE_KEY="3a10b4bc1258e8bfefb95b498fb8c0f0cd6964a811eabca87df56xxxxxxxxxxxx"
```

## Download the API Index and Logging Config Files

Run the following commands:

```text
wget https://raw.githubusercontent.com/tellor-io/TellorMiner/master/configs/indexes.json

wget https://raw.githubusercontent.com/tellor-io/TellorMiner/master/configs/loggingConfig.json
```

## Download and Edit the Manual Data Entry File

Tellor currently has one data point which must be manually created. The rolling 3 month average of the US PCE . It is updated monthly. _Make sure to keep this file up to date._

Run the following command:

```text
wget https://raw.githubusercontent.com/tellor-io/TellorMiner/master/configs/manualData.json
```

For testing purposes, or if you want to hardcode in a specific value to enter, you can use the manualdata.json file to add manual data for a given requestID. Similar to the manual data structure, you add the request ID, a given value \(with granularity\), and a date which the manual data is valid until.

The following example shows request ID 4, inputting a value of 9000 with a 1,000,000 granularity. Note the date is a unix timestamp.

```text
"4":{
    "VALUE":9000000000,
    "DATE":1596153600
}
```

## Deposit your Initial Stake

You will need 500 TRB to run your own server for mining. Your stake is locked for a minimum of 7 days after you run the command to request withdrawal.

{% hint style="info" %}
You do not need to stake 500 TRB if you plan to mine on a pool.
{% endhint %}

Run the following command to deposit your stake:

```text
tellor --config=./configs/config.json stake deposit
```

If needed, change the name of config.json to match the name of your config file.

## Run the miner

Start the dataServer:

```text
tellor --config=./configs/config.json dataServer
```

Start the miner by running this command in another terminal or process:

```text
tellor --config=./configs/config.json mine
```

After starting the miner, observe the logs it outputs to confirm it's working correctly. In the next section we will look at some configuration options that can help improve performance.

## Unstaking / Ending Mining Operations

To unstake your tokens, you need to request a withdraw:

```text
tellor --config=./configs/config.json stake request
```

One week after the request, the tokens are free to move at your discretion after running the command:

```text
tellor --config=./configs/config.json stake withdraw
```

## Running the Disputer

Tellor as a system only functions properly if parties actively monitor the tellor network and dispute bad values. Expecting parties to manually look at every value submitted is obviously burdensome. The Tellor disputer automates this fact checking of values.

The way that it work is that your dataServer will store historical values \(e.g. the last 10 minutes\) and then compare any submitted values to the min/max of the historical values. If the value submitted is outside a certain threshold \(e.g. 10% of the min/max\), then the party will be notified and they can choose if they wish to dispute the bad value.

To start the disputer, add the following line to your config file IN THE TRACKERS ARRAY:

```text
"disputeChecker"
```

Now when running the dataServer, you will store historical values and check for whether the submitted values were within min/max of the range of historical values given a threshold \(e.g. 1% outside\). The variables for configuring the time range of the historical values and the threshold are as follows:

```text
  disputeTimeDelta: 5,
  disputeThreshold: 0.01,
```

Where 5 and .01 are the defaults, the variables are the amount of time in minutes to store historical values for comparison and the the threshold outside the min/max of the values \(e.g. 0.01 = 1%\);

If the disputer is successful and finds a submitted outside of your acceptable range, a text file containing pertinent information will be created in your working directory \(the one you're running the miner out of\) in the format: `"possible-dispute-(blocktime).txt"`

