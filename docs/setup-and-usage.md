---
description: Here are the nuts and bolts for using the CLI
---

# Setup and usage

## Get the CLI

The CLI support only linux and is provided as a pre-built binary with every release and also as a docker image.

[Github releases](https://github.com/tellor-io/telliot/releases)

[https://hub.docker.com/u/tellor](https://hub.docker.com/u/tellor)

## Config files.
 - `.env` - keeps private information(private keys, api keys etc.). Most commands require some secrets and these are kept in this file as a precaution against accidental exposure. For a working setup it is required to at least add one private key in your `"ETH_PRIVATE_KEYS"` environment variable. Multiple private keys are supported separated by `,`.
 - `index.json` - all api endpoint for data providers. The cli uses these provider endpoints to gather data which is then used to submit to the onchain oracle.
 - `manualdata.json` - for providing data manually. There is currently one data point which must be manually created. The rolling 3 month average of the US PCE . It is updated monthly. _Make sure to keep this file up to date._
 For testing purposes, or if you want to hardcode in a specific value, you can use the file to add manual data for a given requestID. Add the request ID, a given value \(with granularity\), and a date on which the manual data expires.
The following example shows request ID 4, inputting a value of 9000 with 6 digits granularity. Note the date is a unix timestamp.
```bash
"4":{
    "VALUE":9000.123456,
    "DATE":1596153600
}
```
 - `config.json` - optional config file to override any of the defaults. See [configuration page](configuration.md) for full reference.


> by default the cli looks for these in the `./configs` folder relative to the cli folder.

### Here is a quick reference how to run the cli with the default configs.

```
mkdir ./configs
cd ./configs
wget https://raw.githubusercontent.com/tellor-io/telliot/master/configs/index.json
wget https://raw.githubusercontent.com/tellor-io/telliot/master/configs/manualData.json
wget https://raw.githubusercontent.com/tellor-io/telliot/master/configs/env.example
mv env.example .env
cd ../
wget https://github.com/tellor-io/telliot/releases/latest/download/telliot
chmod +x telliot
```

## Deposit or withdraw a stake

As of now, mining requires you to deposit 500 TRB to be allowed to submit values to the oracle and earn rewards. This is a security deposit. If you are a malicious actor \(aka submit a bad value\), the community can vote to slash your 500 tokens.
Your stake is locked for a minimum of 7 days after you run the command to request withdrawal.

Run the following command to deposit your stake:

```bash
./telliot stake deposit
```

To unstake your tokens, you need to request a withdraw:

```bash
./telliot stake request
```

One week after the request, the tokens are free to move at your discretion after running the command:

```bash
./telliot stake withdraw
```

## Start mining.
{% hint style="info" %}
The same instance can be used with multiple private keys in the `.env` file separated by a comma.
{% endhint %}

```bash
./telliot mine
```

## DataServer - a shared data API feeds.

{% hint style="info" %}
Advanced usage! If you are setting up a Tellor miner for the first time, it might be a good idea to skip this section and come back after you're up and running with one miner. See the [configuration page](configuration.md) for the required configs.
{% endhint %}

Some oracle feeds require 24h avarages and for these enough historical data is needed. Running a dataserver is the solution to always have enough historical data to generate these averages.

The network topology of this setup looks like the diagram below.
One ore more miners are connected to the same data server for fetching current or historical data to submit to the oracle.
The data server pulls data from the API providers, the 5 staked miners pull data from the data server and submit on-chain to the Tellor Core smart contracts.

```bash
                            /(0xE037)\
                Miner      | (0xcdd8) |
Tellor     <-> (multiple   | (0xb9dD) | <-> Data Server <-> Data APIs
(on chain)      keys)      | (0x2305) |
                            \(0x3233)/
```


## Run with Docker - [https://hub.docker.com/u/tellor](https://hub.docker.com/u/tellor)

```bash
cp configs/.env.example configs/.env # Edit the file after the copy.
docker run -v $(pwd)/configs:/configs tellor/telliot:master mine
```

## Run cli in mining mode with k8s

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
export INSTANCE_NAME=lat # Use max 3 characters due to k8s limitation for port names.
export CFG_FOLDER=.local/configs/$INSTANCE_NAME # Configs will be copied to this folder.
export DEPL_NAME=telliot-m # This is the name of the deployment file.
export DEPL_INSTANCE_NAME=$DEPL_NAME-$INSTANCE_NAME
mkdir -p $CFG_FOLDER

# Create the secret file.
cp configs/.env.example $CFG_FOLDER/.env # Edit the file after the copy.

touch $CFG_FOLDER/config.json # Create an empty file and if needed overwrite the defaults.

# Copy the manual data file.
cp configs/manualData.json $CFG_FOLDER/manualData.json

# Apply the configs.
kubectl create secret generic $DEPL_INSTANCE_NAME --from-env-file=$CFG_FOLDER/.env
kubectl create configmap $DEPL_INSTANCE_NAME \
  --from-file=configs/index.json \
  --from-file=$CFG_FOLDER/config.json \
  --from-file=$CFG_FOLDER/manualData.json \
  -o yaml --dry-run=client | kubectl apply -f -

# Copy the manifest and run it.
cp configs/manifests/$DEPL_NAME.yml $CFG_FOLDER/$DEPL_NAME.yml
sed -i "s/$DEPL_NAME/$DEPL_INSTANCE_NAME/g" $CFG_FOLDER/$DEPL_NAME.yml
kubectl apply -f $CFG_FOLDER/$DEPL_NAME.yml
```

### Run the cli in dataserver mode.

```bash
export INSTANCE_NAME=lat # Use max 3 characters due to k8s limitation for port names.
export CFG_FOLDER=.local/configs/db
export DEPL_NAME=telliot-db
mkdir -p $CFG_FOLDER

# Run the same commands as the mining deployment.

See [configuration page](configuration.md) on how to setup other instances to connect to this remote dataserver

### To run another instance.

```bash
export NAME= # Put an instance name here. Something short as some properties are limited by length(e.g `export NAME=PR1`).
# Run all the other commands from initial k8s setup.
```

### To delete an instance.

```bash
kubectl delete statefulsets.apps $DEPL_INSTANCE_NAME
kubectl delete service $DEPL_INSTANCE_NAME
kubectl delete configmap $DEPL_INSTANCE_NAME
kubectl delete secret $DEPL_INSTANCE_NAME
kubectl delete persistentvolumeclaims $DEPL_INSTANCE_NAME
```

### To run a custom docker image.

```bash
export REPO= # Your docker repository name.
docker build . -t $REPO/telliot:custom
docker push $REPO/telliot:latest

sed -i "s/tellor\/telliot:latest/$REPO\/telliot:custom/g" $CFG_FOLDER/telliot-m.yml
kubectl apply -f $CFG_FOLDER/telliot-m.yml
```

### Optionally deploy the monitoring stack with Prometheus and Grafana.

```bash
kubectl apply -f configs/manifests/monitoring-persist.yml
kubectl apply -f configs/manifests/monitoring.yml
```

###  Optionally deploy the alerting manager and get alerts on your Telegram bot.

This uses the alertmanager bot. see [here](https://github.com/metalmatze/alertmanager-bot) for more info and available commands.

```bash
# Create a secret for the telegram authentication.
kubectl create secret generic alertmanager-bot \
  --from-literal=admin='<telegram admin>' \
  --from-literal=token='<telegram token>'
kubectl apply -f configs/manifests/alerting-persist.yml
kubectl apply -f configs/manifests/alerting.yml
```

## Becoming a Miner

For over a decade now, the Bitcoin network has shown how proof-of-work can incentivize individuals and companies to compete for the honor of finding block rewards and achieving consensus. This phenomenon is global and anonymous. The network is democratized and decentralized because the creators have no direct control over who is providing computing power on their network.

Tellor takes this concept and applies it directly to the delivery of oracle data. Anyone who is able may start up `telliot` and begin competing for blocks. There is no whitelisting. Miners compete very much the same way that Bitcoin miners do, but with a twist. _Tellor Miners must also run a database from which to pull values to submit to the Tellor oracle._ When a "block" is found, the winners submit their data.

Mining is one of the most exciting ways to help Tellor grow and become a leader in the DeFi / Oracle space. Here are a few things to consider before jumping in:

* Mining requires access to an Ethereum node. If you don’t have your own node, you can use an Infura API endpoint.
* Miners must hold a balance of ETH to cover gas fees, which can be significant.
* There is no guarantee of profit from mining and rely heavily on gas cost. There is no promise that Tellor Tributes currently hold or will ever hold any value.

If you are building a competing client, please contact us. A lot of the miner specifications are off-chain and a significant portion of the mining process hinges on the consensus of the Tellor community to determine what proper values are. Competing clients that change different pieces run the risk of being disputed by the community.

As an example, request ID 4 is BTC/USD. If the APIs all go down, it is the responsibility of the miner to still submit a valid BTC/USD price. If they submit incorrect values, they risk being disputed and slashed. For these reasons, please contribute openly to the official telliot cli \(or an open source variant\), as consensus here is key. If your miner gets a different value than the majority of the other miners, you risk being punished!

