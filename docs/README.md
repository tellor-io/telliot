# Introduction

[![tellor.io](../.gitbook/assets/Tellor.png)](https://www.tellor.io/)

[![Twitter WeAreTellor](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/WeAreTellor)

> ### ⚠️ Note!
>
> Telliot docs are synced from the project repository. To edit please visit: [Telliot gitbook space](https://app.gitbook.com/@tellor-2/s/telliot/) or [Telliot github repo](https://github.com/tellor-io/telliot/tree/master/docs)

## Telliot - the tellor.io tasker

This is the main CLI of the project. One of its most useful features is to run in mining mode\(solve a POW challenge\) and submit values to the tellor oracle contract. It's built with Go and utilizes a modular structure so people wanting to use their own logic can use only some modules. See the [internal architecture page](internal-architecture.md) for more details about each module.

The Tellor system is a way to push data on-chain. Note that the data does NOT correspond to a specific API. The tellor mining system is set up to pull API or manually entered data to generate values requested by the on-chain Tellor smart contract and then to submit once a correct nonce is mined. Any specific APIs in the telliot repo are just suggestions and you can use any api you desire to generate the proper corresponding data. The system is not guaranteed to work for everyone. It is up to the consensus of the Tellor token holders to determine what a correct value is. As an example, request ID 4 is BTC/USD. If the APIs listed in the telliot repo all go down, it is the responsibility of the miner to still submit a valid BTC/USD price. If they do not, they risk being disputed and slashed. For these reasons, please contribute openly to the official Telliot system \(or an open-source variant\), as consensus here is key. If your miner gets a different value than the majority of the other miners, you risk being punished.

![MinerSpecs](../.gitbook/assets/minerspecs.png)


## Becoming a Miner

For over a decade now, the Bitcoin network has shown how proof-of-work can incentivize individuals and companies to compete for the honor of finding block rewards and achieving consensus. This phenomenon is global and anonymous. The network is democratized and decentralized because the creators have no direct control over who is providing computing power on their network.

Tellor takes this concept and applies it directly to the delivery of oracle data. Anyone who is able may start up `telliot` and begin competing for blocks. There is no whitelisting. Miners compete very much the same way that Bitcoin miners do, but with a twist. _Tellor Miners must also run a database from which to pull values to submit to the Tellor oracle._ When a "block" is found, the winners submit their data.

Mining is one of the most exciting ways to help Tellor grow and become a leader in the DeFi / Oracle space. Here are a few things to consider before jumping in:

* Mining requires access to an Ethereum node. If you don’t have your own node, you can use an Infura API endpoint.
* Miners must hold a balance of ETH to cover gas fees, which can be significant.
* There is no guarantee of profit from mining and rely heavily on gas cost. There is no promise that Tellor Tributes currently hold or will ever hold any value.

If you are building a competing client, please contact us. A lot of the miner specifications are off-chain and a significant portion of the mining process hinges on the consensus of the Tellor community to determine what proper values are. Competing clients that change different pieces run the risk of being disputed by the community.

As an example, request ID 4 is BTC/USD. If the APIs all go down, it is the responsibility of the miner to still submit a valid BTC/USD price. If they submit incorrect values, they risk being disputed and slashed. For these reasons, please contribute openly to the official telliot cli \(or an open source variant\), as consensus here is key. If your miner gets a different value than the majority of the other miners, you risk being punished!

### Maintainers

This repository is maintained by the [Tellor team](https://github.com/orgs/tellor-io/people)

