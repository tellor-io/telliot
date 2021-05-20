---
description: Internal architecture of the project.
---

# Internal architecture

## Trackers

A tracker is a process that gets data from an HTTP API or a Blockchain smart contract.
All data is then stored in a database for later aggregation.

There are different types of index trackers.
The trackers module uses an `index.json` file.


## Aggregator

The aggregator package aggregated data to expose median, mean, TWAP, VWAP etc.
It also exposes all required data IDs required to submit data to the Tellor oracle. These are specified in the package source file.

## Miner

It simply solves the PoW challenge which is required by the oracle contract wen submitting values.

## API

The cli exposes an api to query all collected data from the trackers.
The api is an exact copy of the [Prometheus API](https://prometheus.io/docs/prometheus/latest/querying/api/) which uses the [promql query language](https://prometheus.io/docs/prometheus/latest/querying/basics).


