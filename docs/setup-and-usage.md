---
description: Here are the nuts and bolts for usinng the CLI
---

# Setup and usage

## Get the CLI

The CLI is provided as a pre-built binary with every release and also as a docker image.

### Run manually
 Download and run the [latest release](https://github.com/tellor-io/telliot/releases/latest)

```bash
wget https://github.com/tellor-io/telliot/releases/latest/download/telliot
chmod +x telliot
```

### Run with Docker - [https://hub.docker.com/u/tellor](https://hub.docker.com/u/tellor)

```bash
docker run -v $(pwd)/local:/configs tellor/telliot:master mine
```

### Run with k8s

{% hint style="info" %}
tested with [google cloud](https://cloud.google.com), but should work with any k8s cluster.
{% endhint %}

 * Install [`gcloud`](https://cloud.google.com/sdk/docs/install)
 * Install [`kubectl`](https://kubernetes.io/docs/tasks/tools/install-kubectl)
 * Create a k8s cluster with a single node
 * Login to the cluster

```bash
gcloud auth login --project projectName
gcloud container clusters get-credentials main --zone europe-west2-a --project projectName
```

 * Deploy the `cli` \(by default deployed to run as a miner\)

```bash
git clone https://github.com/tellor-io/telliot
cd telliot
export NAME=main
mkdir -p .local/configs/$NAME

# Create the secret file.
cp configs/.env.example .local/configs/$NAME/.env # Edit the file after the copy.
kubectl create secret generic telliot-$NAME --from-env-file=.local/configs/$NAME/.env

cp configs/config.json .local/configs/$NAME/config.json # Edit the file after the copy.

# Copy the index, manual. These can be used as it without editing.
cp configs/indexes.json .local/configs/$NAME/indexes.json
cp configs/manualData.json .local/configs/$NAME/manualData.json
# Add the configs.
kubectl create configmap telliot-$NAME \
  --from-file=.local/configs/$NAME/config.json \
  --from-file=.local/configs/$NAME/indexes.json \
  --from-file=.local/configs/$NAME/manualData.json \
  -o yaml --dry-run=client | kubectl apply -f -

# Copy the deployment and run it.
cp configs/manifests/telliot.yml .local/configs/$NAME/telliot.yml
sed -i "s/telliot-main/telliot-$NAME/g" .local/configs/$NAME/telliot.yml
kubectl apply -f .local/configs/$NAME/telliot.yml
```

#### To run another instance.

```bash
export NAME= # Put an instance name here. Something short as some properties are limited by length(e.g `export NAME=PR320`).
# Run all the other commands from initial k8s setup.
```

#### To run a custom docker image.

```bash
export REPO= # Your docker repository name.
docker build . -t $REPO/telliot:latest
docker push $REPO/telliot:latest

sed -i "s/tellor\/telliot:master/$REPO\/telliot:latest/g" .local/configs/$NAME/telliot.yml
kubectl apply -f .local/configs/$NAME/telliot.yml
```

 * Optionally deploy the monitoring stack with Prometheus and Grafana.

```bash
kubectl apply -f configs/manifests/monitoring-persist.yml
kubectl apply -f configs/manifests/monitoring.yml
```

### Download and Edit config.json

`config.json` is where you will enter your wallet address and configure the CLI for your machine.

```bash
wget https://raw.githubusercontent.com/tellor-io/telliot/master/configs/config.json
```

Open config.json and update the following values:

* Set `"publicAddress"` to the public key for the Ethereum wallet you plan to use for mining. Remove the 0x prefix at the beginning of the address.

### Create .env file

Most commands require some secrets and these are kept in a separate `configs/.env`. This is a precaution so that are not accidentally exposed as part of the main config. Make a copy of the `env.example` and edit with your secrets.

## mine - Become a Miner

{% hint style="warning" %}
#### DISCLAIMER

Mine at your own risk.

Mining requires you to deposit 500 Tellor Tributes use as a security deposit. If you are a malicious actor \(aka submit a bad value\), the community can vote to slash your tokens.

Mining also requires submitting on-chain transactions on Ethereum. These transactions cost gas \(ETH\) and can sometimes be significant if the cost of gas on EThereum is high \(i.e. the network is clogged\). Please reach out to the community to find the best tips for keeping gas costs under control or at least being aware of the costs.

If you are building a competing client, please contact us. The miner specifications are off-chain and the validity of the mining process hinges on the consensus of the community to determine what proper values are. Competing clients that change different pieces run the risk of being disputed by the community.

There is no guarantee of profit from mining. There is no promise that Tellor Tributes currently hold or will ever hold any value.
{% endhint %}

{% hint style="info" %}
#### DISCLAIMER

If you are building a competing client, please contact us. A lot of the miner specifications are off-chain and a significant portion of the mining process hinges on the consensus of the Tellor community to determine what proper values are. Competing clients that change different pieces run the risk of being disputed by the community.

As an example, request ID 4 is BTC/USD. If the APIs all go down, it is the responsibility of the miner to still submit a valid BTC/USD price. If they do not, they risk being disputed and slashed. For these reasons, please contribute openly to the official telliot miner \(or an open source variant\), as consensus here is key. If your miner gets a different value than the majority of the other miners, you risk being punished!
{% endhint %}

For over a decade now, the Bitcoin network has shown how proof-of-work can incentivize individuals and companies to compete for the honor of finding block rewards and achieving consensus. This phenomenon is global and anonymous. The network is democratized and decentralized because the creators have no direct control over who is providing computing power on their network.

Tellor takes this concept and applies it directly to the delivery of oracle data. Anyone who is able may start up `telliot` and begin competing for blocks. There is no whitelisting. Miners compete very much the same way that Bitcoin miners do, but with a twist. _Tellor Miners must also run a database from which to pull values to submit to the Tellor oracle._ When a "block" is found, the winners submit their data.

Mining is one of the most exciting ways to help Tellor grow and become a leader in the DeFi / Oracle space. Here are a few things to consider before jumping in:

As of now, mining requires you to deposit 500 Tellor Tributes. This is a security deposit. If you are a malicious actor \(aka submit a bad value\), the community can vote to slash your 500 tokens.

* Mining requires access to an Ethereum node. If you don’t have your own node, you can use an Infura API endpoint.
* Miners must hold a balance of ETH to cover gas fees, which can be significant. Please reach out to the community to find the best tips for keeping gas costs under control.

The guide that follows assumes that you have access to a suitable machine running linux to use for mining. For information about what constitutes a "suitable machine", we recommend reaching out to the community.

### Download the API Index and Logging Config Files

Run the following commands:

```bash
wget https://raw.githubusercontent.com/tellor-io/telliot/master/configs/indexes.json
```

### Download and Edit the Manual Data Entry File

Tellor currently has one data point which must be manually created. The rolling 3 month average of the US PCE . It is updated monthly. _Make sure to keep this file up to date._

Run the following command:

```bash
wget https://raw.githubusercontent.com/tellor-io/telliot/master/configs/manualData.json
```

For testing purposes, or if you want to hardcode in a specific value to enter, you can use the manualdata.json file to add manual data for a given requestID. Similar to the manual data structure, you add the request ID, a given value \(with granularity\), and a date on which the manual data expires.

The following example shows request ID 4, inputting a value of 9000 with a 1,000,000 granularity. Note the date is a unix timestamp.

```bash
"4":{
    "VALUE":9000000000,
    "DATE":1596153600
}
```

### Start mining.

```bash
telliot --config=./configs/config.json mine
```

### Bonus section - connecting to a Pool

There are mining pools available for mining TRB without staking any tokens. The pool server operator stakes the tokens for you, and you receive rewards roughly proportional to your hashrate as a fraction of the pool's hashrate.

{% hint style="info" %}
Each pool has different fees and instructions for hooking up. Be sure to read your pool's documentation. Feel free to reach out to the community if you need help with mining pools.
{% endhint %}

Add the following lines to your config file:

```bash
"enablePoolWorker": true,
"poolURL": "<poolURL>",
```

Where the poolURL is the link to your pool. \(e.g. [http://tellorpool.org](http://tellorpool.org) \)

You can change the job duration if needed. This is the time in seconds to grab information from the pool. The default time is 15 seconds.

```bash
"poolJobDuration":10
```

## deposit - Deposit or withdraw a stake

{% hint style="info" %}
You do not need to stake 500 TRB if you plan to mine on a pool.
{% endhint %}

You will need 500 TRB to run your own server for mining. Your stake is locked for a minimum of 7 days after you run the command to request withdrawal.

Run the following command to deposit your stake:

```bash
tellor --config=./configs/config.json stake deposit
```

To unstake your tokens, you need to request a withdraw:

```bash
telliot --config=./configs/config.json stake request
```

One week after the request, the tokens are free to move at your discretion after running the command:

```bash
telliot --config=./configs/config.json stake withdraw
```

## dispute - monitor submitted values

Tellor as a system only functions properly if parties actively monitor the tellor network and dispute bad values. Expecting parties to manually look at every value submitted is obviously burdensome. The Tellor disputer automates this fact checking of values.

The way that it works is that the dataServer component will store historical values \(e.g. the last 10 minutes\) and then compare any submitted values to the min/max of the historical values. If the value submitted is outside a certain threshold \(e.g. 10% of the min/max\), then the party will be notified and they can choose if they wish to dispute the bad value.

To start the disputer, add the following line to your config file IN THE TRACKERS ARRAY:

```bash
"disputeChecker"
```

Now when running the dataServer, you will store historical values and check for whether the submitted values were within min/max of the range of historical values given a threshold \(e.g. 1% outside\). The variables for configuring the time range of the historical values and the threshold are as follows:

```bash
  disputeTimeDelta: 5,
  disputeThreshold: 0.01,
```

Where 5 and .01 are the defaults, the variables are the amount of time in minutes to store historical values for comparison and the threshold outside the min/max of the values \(e.g. 0.01 = 1%\);

If the disputer is successful and finds a submitted outside of your acceptable range, a text file containing pertinent information will be created in your working directory \(the one you're running the miner out of\) in the format: `"possible-dispute-(blocktime).txt"`

## dataServer - connect more than one miner to work together.

{% hint style="info" %}
Advanced usage! If you are setting up a Tellor miner for the first time, it might be a good idea to skip this section and come back after you're up and running with one miner.
{% endhint %}

If you are running multiple miners, there is no reason to run multiple databases \(the values you will submit should be identical\). In addition, querying the same API from multiple processes can lead to rate limits on the public APIs. To get around this, you can utilize a system where you run one.

In this example will 5 miners connected to a single data server. These 5 miners will start the mining process and the 1 data server will be how each of the 5 miners fetch data from the internet. The network topology of this setup is as follow:

```bash
           <-> Miner (0xE037) <->
           <-> Miner (0xcdd8) <->
Tellor     <-> Miner (0xb9dD) <-> Data Server <-> Internet
(on chain) <-> Miner (0x2305) <->
           <-> Miner (0x3233) <->
```

The data server pulls data from the internet, the 5 staked miners pull data from the data server and submit on-chain to the Tellor Core smart contracts. The following instructions cover setting this up locally.

```bash
wget https://raw.githubusercontent.com/tellor-io/telliot/master/configs/config.json
cp config.json config1.json
telliot --config=config1.json dataServer
```

Edit `config1.json` to include the following:

```bash
{
    "publicAddress": "0xE037EC8EC9ec423826750853899394dE7F024fee",
    "databaseURL":"http://localhost7545",
    "serverWhitelist": [
                "0xE037EC8EC9ec423826750853899394dE7F024fee",
                "0xcdd8FA31AF8475574B8909F135d510579a8087d3",
                "0xb9dD5AfD86547Df817DA2d0Fb89334A6F8eDd891",
                "0x230570cD052f40E14C14a81038c6f3aa685d712B",
                "0x3233afA02644CCd048587F8ba6e99b3C00A34DcC"
    ],
    "serverHost": "localhost",
    "serverPort": 5000,
    "ethClientTimeout": 3000,
    "trackerCycle": 10,
    "requestData":1,
    "gasMultiplier": 1,
    "gasMax":10,
    "trackers": [
          "balance",
          "currentVariables",
          "disputeStatus",
          "gas",
          "top50",
          "tributeBalance",
          "indexers"
    ],
    "dbFile": "./tellorDB"
    "envFile": ".env1"
}
```

After saving this `config1.json` file. Create 4 copies of this file and edit the `dbFile`, `publicAddress`, `envFile` location for each of the files to include the other 5 staked miner addresses \(the command below do this for you with `cp` and `sed`\):

```bash
cp config1.json config2.json
cp config1.json config3.json
cp config1.json config4.json
cp config1.json config5.json

sed -i -e 's/.env1/.env2/' config2.json
sed -i -e 's/.env1/.env3/' config3.json
sed -i -e 's/.env1/.env4/' config4.json
sed -i -e 's/.env1/.env5/' config5.json

sed -i -e 's/tellorDB/tellorDB2/' config2.json
sed -i -e 's/tellorDB/tellorDB3/' config3.json
sed -i -e 's/tellorDB/tellorDB4/' config4.json
sed -i -e 's/tellorDB/tellorDB5/' config5.json

sed -i -e '1,/0xE037EC8EC9ec423826750853899394dE7F024fee/ s/0xE037EC8EC9ec423826750853899394dE7F024fee/0xcdd8FA31AF8475574B8909F135d510579a8087d3/' config2.json
sed -i -e '1,/0xE037EC8EC9ec423826750853899394dE7F024fee/ s/0xE037EC8EC9ec423826750853899394dE7F024fee/0xb9dD5AfD86547Df817DA2d0Fb89334A6F8eDd891/' config3.json
sed -i -e '1,/0xE037EC8EC9ec423826750853899394dE7F024fee/ s/0xE037EC8EC9ec423826750853899394dE7F024fee/0x230570cD052f40E14C14a81038c6f3aa685d712B/' config4.json
sed -i -e '1,/0xE037EC8EC9ec423826750853899394dE7F024fee/ s/0xE037EC8EC9ec423826750853899394dE7F024fee/0x3233afA02644CCd048587F8ba6e99b3C00A34DcC/' config5.json
```

Create `.env` file with the private key for each miner.

```bash
echo "ETH_PRIVATE_KEY=4bdc16637633fa4b4854670fbb83fa254756798009f52a1d3add27fb5f5a8e16" > .env1
echo "ETH_PRIVATE_KEY=d32132133e03be292495035cf32e0e2ce0227728ff7ec4ef5d47ec95097ceeed" > .env2
echo "ETH_PRIVATE_KEY=d13dc98a245bd29193d5b41203a1d3a4ae564257d60e00d6f68d120ef6b796c5" > .env3
echo "ETH_PRIVATE_KEY=4beaa6653cdcacc36e3c400ce286f2aefd59e2642c2f7f29804708a434dd7dbe" > .env4
echo "ETH_PRIVATE_KEY=78c1c7e40057ea22a36a0185380ce04ba4f333919d1c5e2effaf0ae8d6431f14" > .env5


echo "NODE_URL=https://mainnet.infura.io/v3/xxxxxxxxxxxxx" >> .env1
echo "NODE_URL=https://mainnet.infura.io/v3/xxxxxxxxxxxxx" >> .env2
echo "NODE_URL=https://mainnet.infura.io/v3/xxxxxxxxxxxxx" >> .env3
echo "NODE_URL=https://mainnet.infura.io/v3/xxxxxxxxxxxxx" >> .env4
echo "NODE_URL=https://mainnet.infura.io/v3/xxxxxxxxxxxxx" >> .env5
```

Finaly, make 1 more copy of the config for the data server and update the `serverHost` address to `0.0.0.0`:

```bash
cp config1.json config-dataserver.json
sed -i -e 's/\"serverHost\": \"localhost\"/\"serverHost\": \"0.0.0.0\"/' config-dataserver.json
```

The stakes have already been deposited for these Addresses so you can now move on to starting up each of the miners.

#### Starting the Miners and Data Server

You can do this in 6 separate terminals locally. Run each of the command in each of the terminals and confirm they start up correctly.

| Terminal \# | Command | Description |
| :--- | :--- | :--- |
| 1 | ./telliot --config=config-dataserver.json dataserver | Data Server |
| 2 | ./telliot --config=config1.json mine -r | Staked Miner 1 |
| 3 | ./telliot --config=config2.json mine -r | Staked Miner 2 |
| 4 | ./telliot --config=config3.json mine -r | Staked Miner 3 |
| 5 | ./telliot --config=config4.json mine -r | Staked Miner 4 |
| 6 | ./telliot --config=config5.json mine -r | Staked Miner 5 |

#### Conclusion

At this point, you will have 7 terminals running: 6 terminals for the `telliot` and 1 terminal for running Ganache. You should see your miners are submitting transactions and if you want to check that the network difficulty is rising, you can use Truffle's console again and run the following commands:

```bash
let difficulty = await oracle.getUintVar("0xb12aff7664b16cb99339be399b863feecd64d14817be7e1f042f97e3f358e64e")
difficulty.toNumber()
```

