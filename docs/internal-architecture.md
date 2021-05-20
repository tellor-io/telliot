---
description: Internal architecture of the project.
---

# Internal architecture

## Tasker

Monitors the oracle contract for new data requests.
As soon as a new challenge is emitted it is sent to the miner module to find a solution.

## Miner

It simply solves the PoW challenge which is required by the oracle contract wen submitting values.

## Submitter

Submits the currently requested data IDs to the oracle contract.
It makes all the necessary checks to prepare the data accordingly to avoid failed transactions.
The data is taken from a local or remote DB.
The data is collected by the Index Tracker module and aggregated by the Aggregator module.

## Trackers

A tracker is module that gets and parses data from an HTTP API or a Blockchain smart contract for later usage or aggregation.
There are different types of trackers the main one being the [Index Tracker](index-tracker.md)

## Aggregator

Aggregates data to expose median, mean, TWAP, VWAP etc.
It also exposes all required data IDs required to submit data to the Tellor oracle.
These are specified in the module itself.

## API

The cli exposes an api to query all collected data from the trackers.
The api is an exact copy of the [Prometheus API](https://prometheus.io/docs/prometheus/latest/querying/api/) which uses the [promql query language](https://prometheus.io/docs/prometheus/latest/querying/basics).


