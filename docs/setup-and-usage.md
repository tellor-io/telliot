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
cp configs/api.json .local/configs/$NAME/api.json
cp configs/manualData.json .local/configs/$NAME/manualData.json
# Add the configs.
kubectl create configmap telliot-$NAME \
  --from-file=.local/configs/$NAME/config.json \
  --from-file=.local/configs/$NAME/api.json \
  --from-file=.local/configs/$NAME/manualData.json \
  -o yaml --dry-run=client | kubectl apply -f -

# Copy the StatefulSet and run it.
cp configs/manifests/telliot.yml .local/configs/$NAME/telliot.yml
sed -i "s/telliot-main/telliot-$NAME/g" .local/configs/$NAME/telliot.yml
kubectl apply -f .local/configs/$NAME/telliot.yml
```

#### To run another instance.

```bash
export NAME= # Put an instance name here. Something short as some properties are limited by length(e.g `export NAME=PR320`).
# Run all the other commands from initial k8s setup.
```

#### To delete an instance.

```bash
export NAME=
kubectl delete statefulsets.apps $NAME
kubectl delete service $NAME
kubectl delete configmap $NAME
kubectl delete secret $NAME
kubectl delete persistentvolumeclaims $NAME
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
For a working setup it is required to at least add one private key in your `"ETH_PRIVATE_KEYS"` environment variable. all of public addresses can be determined from your private keys.

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
wget https://raw.githubusercontent.com/tellor-io/telliot/master/configs/api.json
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
telliot mine --config=./configs/config.json
```

## deposit - Deposit or withdraw a stake

{% hint style="info" %}
You do not need to stake 500 TRB if you plan to mine on a pool.
{% endhint %}

You will need 500 TRB to run your own server for mining. Your stake is locked for a minimum of 7 days after you run the command to request withdrawal.

Run the following command to deposit your stake:

```bash
tellor stake deposit --config=./configs/config.json
```

To unstake your tokens, you need to request a withdraw:

```bash
telliot stake request --config=./configs/config.json
```

One week after the request, the tokens are free to move at your discretion after running the command:

```bash
telliot stake withdraw --config=./configs/config.json
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

## dataServer - have a dataserver running all the time.

{% hint style="info" %}
Advanced usage! If you are setting up a Tellor miner for the first time, it might be a good idea to skip this section and come back after you're up and running with one miner.
{% endhint %}

It is recommended to have a dataserver running all the time so data could be saved and also it is needed as some prices need 24h averages.

In this example will a miner connected to a data server. This miner will start the mining process using multiple keys and the 1 data server will fetch required data from the internet. The network topology of this setup is as follow:

```bash
                            /(0xE037)\
                Miner      | (0xcdd8) |
Tellor     <-> (multiple   | (0xb9dD) | <-> Data Server <-> Internet
(on chain)      keys)      | (0x2305) |
                            \(0x3233)/
```

The data server pulls data from the internet, the 5 staked miners pull data from the data server and submit on-chain to the Tellor Core smart contracts. The following instructions cover setting this up locally.

```bash
wget https://raw.githubusercontent.com/tellor-io/telliot/master/configs/config.json
cp config.json config1.json
telliot dataServer --config=config1.json
```

Edit `config1.json` to include the following:

```bash
{
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

After saving this `config1.json` file. Create a copy of this file and edit the `envFile` location to include the 5 staked miner addresses \(the command below do this for you with `cp` and `sed`\):

```bash
cp config1.json config2.json
sed -i -e 's/.env1/.env2/' config2.json
sed -i -e 's/tellorDB/tellorDB2/' config2.json
```

Create `.env` file with the private keys for the miner (if there are more than one private keys, must be seperated by `,`).

```bash
echo "ETH_PRIVATE_KEYS=4bdc16637633fa4b4854670fbb83fa254756798009f52a1d3add27fb5f5a8e16,d32132133e03be292495035cf32e0e2ce0227728ff7ec4ef5d47ec95097ceeed" > .env1

echo "NODE_WEBSOCKET_URL=wss://mainnet.infura.io/v3/ws/xxxxxxxxxxxxx" >> .env1
```

Finaly, make 1 more copy of the config for the data server and update the `serverHost` address to `0.0.0.0`:

```bash
cp config1.json config-dataserver.json
sed -i -e 's/\"serverHost\": \"localhost\"/\"serverHost\": \"0.0.0.0\"/' config-dataserver.json
```

The stakes have already been deposited for these Addresses so you can now move on to starting up the miner using multiple staked keys.

#### Starting the Miners and Data Server

You can do this in 6 separate terminals locally. Run each of the command in each of the terminals and confirm they start up correctly.

| Terminal \# | Command | Description |
| :--- | :--- | :--- |
| 1 | ./telliot dataserver --config=config-dataserver.json | Data Server |
| 2 | ./telliot mine -r --config=config1.json | Miner 1 (Multiple staked keys) |

#### Conclusion

At this point, you will have 3 terminals running: 2 terminals for the `telliot` and 1 terminal for running Ganache. You should see your miner are submitting transactions using multiple staked keys and if you want to check that the network difficulty is rising, you can use Truffle's console again and run the following commands:

```bash
let difficulty = await oracle.getUintVar("0xb12aff7664b16cb99339be399b863feecd64d14817be7e1f042f97e3f358e64e")
difficulty.toNumber()
```

