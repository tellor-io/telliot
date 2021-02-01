---
description: Internal architecture of the project.
---

# Internal architecture

### Trackers

A tracker is usually in the type of an index tracker for a coin or token on a blockchain network.

Index trackers take placed in the `pkg/tracker` package of the Telliot repository. the index tracker struct is as follow:

```go
type IndexTracker struct {
	DB               db.DataServerProxy
	Name             string
	Identifier       string
	Symbols          []string
	Source           DataSource
	Interval         time.Duration
	Param            string
	Type             IndexType
	lastRunTimestamp time.Time
}
```

an IndexTracker struct is related to an index object from the `indexes.json` file. based on `Source` there are two main types of trackers: HTTP trackers, on-chain trackers.

#### HTTP trackers

An index tracker that fetches and parses tracker info using an HTTP call. the `indexes.json` entry for such trackers could be like this:

```javascript
    "AMPL/BTC": [
        {
            "URL": "https://api-pub.bitfinex.com/v2/tickers?symbols=tAMPBTC",
            "param": "$[0][7,8]"
        }
    ]
```

#### On-chain trackers

On-chain index trackers fetch data using smart contract calls on the Ethereum blockchain. currently supported on-chain trackers are Uniswap and Balancer trackers. `indexes.json` entry for such trackers could be like this:

```javascript
    "AMPL/ETH": [
        {
          "URL": "Mainnet:0xc5be99a02c6857f9eac67bbce58df5572498f40c,Rinkeby:0x7e62a502232f1feB77Adf8b8ca023cc9fB133418",
          "type": "ethereum",
          "parser": "Uniswap"
        }
    ]
```



Uniswap and Balancer each have their `DataSource` that fetches tracker info using on-chain calls.

`BalancerGetter` is a `DataSource` that fetches tracker info from a [Balancer pool](https://docs.balancer.finance/getting-started/faq#balancer-pools). Balancer pools are liquidity pools for pair of ERC20 tokens and they could resist on both Ethereum mainnet and its testnets.

Sometimes there is no deployed testnet Balancer pool to use in the On-chain index tracker. [here](https://docs.balancer.finance/guides/testing-on-kovan) are some steps to deploy a Balancer pool on the Koven testnet and add some liquidity to it. for a Rinkeby deployment, the step could be a bit different and required to deploy required ERC20 tokens beforehand. [here](https://github.com/hhio618/simple-erc20-example) is a simple javascript project that could be used to deploy an ERC20 token on the Rinkeby testnet. also, [here](https://docs.balancer.finance/smart-contracts/addresses) are addresses of Balancer smart contracts on different Ethereum networks. \(Note: please make sure that use enough gas when using the Seth client in the above-mentioned Balancer guide!\)

`UniswapGetter` is a `DataSource` that fetches tracker info from a [UniswapV2 pair](https://uniswap.org/docs/v2/smart-contracts/pair/). the easiest way to add UniswapV2 testnet pair is to call the [addLiquidity](https://uniswap.org/docs/v2/smart-contracts/router02/#addliquidity) contract method on Uniswap RouterV2 smart contract on different Ethereum networks. see [addresses](https://uniswap.org/docs/v2/smart-contracts/router02/#addresshttps://uniswap.org/docs/v2/smart-contracts/router02/#address). the method is as follow:

```javascript
function addLiquidity(
  address tokenA,
  address tokenB,
  uint amountADesired,
  uint amountBDesired,
  uint amountAMin,
  uint amountBMin,
  address to,
  uint deadline
) external returns (uint amountA, uint amountB, uint liquidity);

```

This required to deploy some ERC20 token beforehand and will create a Uniswap V2 pair if already not exists for the provided pair. there is a factory method that could be used to get the pair address [here](https://uniswap.org/docs/v2/smart-contracts/factory/#getpair).
