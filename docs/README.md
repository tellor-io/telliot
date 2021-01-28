# Introduction

[![tellor.io](../.gitbook/assets/Tellor.png)](https://www.tellor.io/)

[![Twitter WeAreTellor](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/WeAreTellor)

> ### ⚠️ Note!
>
> Telliot docs are synced from the project repository. To edit please visit: [Telliot gitbook space](https://app.gitbook.com/@tellor-2/s/telliot/) or [Telliot github repo](https://github.com/tellor-io/telliot/tree/master/docs)

## Telliot - the tellor.io tasker

This is the main CLI of the project. One of its most useful features is to run in mining mode\(solve a POW challenge\) and submit values to the tellor oracle contract. It's built with Go and utilizes a split structure. The database piece is a LevelDB that keeps track of all variables \(challenges, difficulty, values to submit, etc.\) and the miner simply solves the PoW challenge. This enables parties to split the pieces for optimization.

The Tellor system is a way to push data on-chain. Note that the data does NOT correspond to a specific API. The tellor mining system is set up to pull API or manually entered data to generate values requested by the on-chain Tellor smart contract and then to submit once a correct nonce is mined. Any specific APIs in the telliot repo are just suggestions and you can use any api you desire to generate the proper corresponding data. The system is not guaranteed to work for everyone. It is up to the consensus of the Tellor token holders to determine what a correct value is. As an example, request ID 4 is BTC/USD. If the APIs listed in the telliot repo all go down, it is the responsibility of the miner to still submit a valid BTC/USD price. If they do not, they risk being disputed and slashed. For these reasons, please contribute openly to the official Telliot system \(or an open-source variant\), as consensus here is key. If your miner gets a different value than the majority of the other miners, you risk being punished.

A list of all PSR's\(pre-specified requests\) and the expected data can be found [here](https://github.com/tellor-io/telliot/blob/master/pkg/tracker/psrs.go).

![MinerSpecs](../.gitbook/assets/minerspecs.png)

### Maintainers

This repository is maintained by the [Tellor team](https://github.com/orgs/tellor-io/people)

