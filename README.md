# Tellor Miner

This is the workhorse of the Miner system as it takes on solving the PoW challenge.  
It's built on Go and utilizes a split structure.  The database piece is a LevelDB that keeps track of all variables (challenges, difficulty, values to submit, etc.) and the miner simply solves the PoW challenge.  This enables parties to split the pieces for optimization.

<p align="center">
    <img src= './public/minerspecs.png' width="250" alt='MinerSpecs' />
</p>


### Tellor Deployed Addresses

Mainnet - [0x0ba45a8b5d5575935b8158a88c631e9f9c95a2e5](https://etherscan.io/address/0x0ba45a8b5d5575935b8158a88c631e9f9c95a2e5)

Rinkeby - [0x3f1571e4dfc9f3a016d90e0c9824c56fd8107a3e](https://rinkeby.etherscan.io/address/0x3f1571e4dfc9f3a016d90e0c9824c56fd8107a3e)



### Instructions for deployment


[Tellor Deployment Instructions](https://github.com/tellor-io/TellorMiner/wiki/Launching-the-Miner---Technical)


[Tellor Deployment Instructions -- For the Non-Technical Miner](https://github.com/tellor-io/TellorMiner/wiki/Launching-the-Miner---Non-Technical)


[Tellor Deployment Instructions -- From Source](https://github.com/tellor-io/TellorMiner/wiki/Launching-the-Miner---From-Source)


#### How to Contribute<a name="how2contribute"> </a>  
Join our Discord or Telegram:
[<img src="./public/telegram.png" width="24" height="24">](https://t.me/tellor)
[<img src="./public/discord.png" width="24" height="24">](https://discord.gg/zFcM3G)


#### Contributors<a name="contributors"> </a>

This repository is maintained by the Tellor team - [www.tellor.io](https://www.tellor.io)


#### Copyright

Tellor Inc. 2019
