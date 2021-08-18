# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/) and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

NOTE: As semantic versioning states all 0.y.z releases can contain breaking changes in API \(flags, grpc API, any backward compatibility\)

We use _breaking :warning:_ to mark changes that are not backward compatible \(relates only to v0.y.z releases.\)

## [v6.0.0](https://github.com/tellor-io/telliot/releases/tag/v6.0.0) - 2021.08.18

### Changed
* [\#465](https://github.com/tellor-io/telliot/pull/465) Update data ID 10 (AMPL/USD) to use VWAP api endpoints directly and no local aggregation.

### Added

### Fixed
* [\#484](https://github.com/tellor-io/telliot/pull/484) Switch "now" and "end of day" for ID 10 (AMPL/USD) url params.

## [v5.9.0](https://github.com/tellor-io/telliot/releases/tag/v5.9.0) - 2021.08.04

### Changed
* [\#479](https://github.com/tellor-io/telliot/pull/479) Update personal consumption expenditures index (PCE) from the [Bureau of Economic Analysis](https://www.bea.gov/news/2021/personal-income-and-outlays-june-2021-and-annual-update).
* [\#469](https://github.com/tellor-io/telliot/pull/469) Update PCE index.

### Added
* [\#477](https://github.com/tellor-io/telliot/pull/477) Add support for ETH/JPY data. Add bash script for getting latest telliot release.
* [\#474](https://github.com/tellor-io/telliot/pull/474) Add polygon addresses.
* [\#471](https://github.com/tellor-io/telliot/pull/471) Add command to tally votes. Add optional gasPrice argument to all commands.

### Fixed
* [\#482](https://github.com/tellor-io/telliot/pull/482) Fix ID 28.
* [\#470](https://github.com/tellor-io/telliot/pull/470) Fix percentage difference function for negative change.
* [\#466](https://github.com/tellor-io/telliot/pull/466) Fix various access bugs.

## [v5.8.0](https://github.com/tellor-io/telliot/releases/tag/v5.8.0) - 2021.06.15

### Changed
* [\#440](https://github.com/tellor-io/telliot/pull/440) Completely refactored the internal architecture to make it more modular, easier to understand and maintain. In this gigantic PR also added a submitter for the new tellor access oracle. All config files now follow a new structure so see the latest [docs](https://docs.tellor.io/tellor/telliot) for more details.

* [\#461](https://github.com/tellor-io/telliot/pull/461) Data ID 57 - TVL now uses different API endpoints to return consistent values and to include only TVL from ethereum and the aggregation is switched from Mean to Median.

* Data ID 41 is now using the default granularity of 6 digits after the decimal point. For example 113.406333 should be submitted as 113406333.

### Added
* [\#432](https://github.com/tellor-io/telliot/pull/432) At startup it prints the current version and git tag.

* [\#452](https://github.com/tellor-io/telliot/pull/452) At startup it prints a message when the cli has a new release to notify people that they can upgrade.

* [\#434](https://github.com/tellor-io/telliot/pull/434) Added k8s deployments for alertmanager and telegram bot alerting.

* [\#441](https://github.com/tellor-io/telliot/pull/441) Auto generating the docs from the code itself for the CLI args and the configs so now they will be allways up to date. Added a check in the CI to make sure the docs are regenerated on any code changes.

* [\#446](https://github.com/tellor-io/telliot/pull/446) Added a new dispute tracker module which tracks all submitted values in the oracle and exposes metrics to allow comparing the values. Eventually will also add settings and docs how to create alerting so that people submitting values can set an elrt when submitted values looks different than what is expected.

### Fixed
* [\#448](https://github.com/tellor-io/telliot/pull/448) Return an error at startup when an entry in the index.json file contains an env variable, but the variable is not set.

* [\#463](https://github.com/tellor-io/telliot/pull/463) When aggregating a median of odd values count, use use mean for the 2 middle numbers.

## [v5.7.0](https://github.com/tellor-io/telliot/releases/tag/v5.7.0) - 2021.02.23

### Changed
* [\#399](https://github.com/tellor-io/telliot/pull/399) We no longer use the`"publicAddress"` config in the config file, instead we get public addresses from the `"ETH_PRIVATE_KEYS"` environment variable.
* [\#399](https://github.com/tellor-io/telliot/pull/399) We changed `"ETH_PRIVATE_KEY"` environment variable to a list of private keys (separated by `,`) and now it will be defined using `"ETH_PRIVATE_KEYS"` environment variable.
* [\#403](https://github.com/tellor-io/telliot/pull/403) The config file now has fields separated by commands(eg. Dataserve, mine,etc) where all the command specific configuration goes. It now uses strict parsing, meaning that unused fields throw an error, which will likely happen to a lot of users. For a clear view of the conifg, please take a look at the [config file](https://github.com/tellor-io/telliot/blob/master/pkg/config/config.go#L105)

* [\#403](https://github.com/tellor-io/telliot/pull/403) Renamed `indexes.json` to `api.json`.

* [\#403](https://github.com/tellor-io/telliot/pull/403) Removed the configs folder and now user should use  apiFile and manualDataFile.

* [\#410](https://github.com/tellor-io/telliot/pull/410) Public addresses should be prefixed with `0x`.

* [\#410](https://github.com/tellor-io/telliot/pull/410) Integration and testing to use the newer contracts.

### Added
* [\#406](https://github.com/tellor-io/telliot/pull/406) Added new command  `migrate` to migrate old tokens for the new one.
### Fixed
* [\#410](https://github.com/tellor-io/telliot/pull/410) Fixed most submit races, causing fewer submission errors. More effort will be dedicated to completely removing them in the next release.

## [v5.6.0](https://github.com/tellor-io/telliot/releases/tag/v5.6.0) - 2021.02.08

### Changed
* [\#240](https://github.com/tellor-io/telliot/issues/240) Replaced the cli package to allow for command specific flags and configuration. Now all flags should be passed last. Example:
  Instead of: `./telliot --config="config.json" stake deposit`, it becomes:`./telliot stake deposit --config="config.json"`
  Removed the `RemoteMining` `-r` flag. Remote is active when specifying a `RemoteDBHost` for the `Mine` command. See the `configs/config.json` for an example.

* [\#378](https://github.com/tellor-io/telliot/pull/378) Removed GPU mining, as it weren't being used.

* [\#390](https://github.com/tellor-io/telliot/pull/390) Removed the feature to connect using Stratum Pool, as it weren't being used.

* [\#386](https://github.com/tellor-io/telliot/pull/386) Removed the need to provide the tellor contract address in the config file

* [\#386](https://github.com/tellor-io/telliot/pull/386) Removed the logLevel flag. Now all logging setup is in the config file. See the example at `configs/config.json`.

### Added
* [\#347](https://github.com/tellor-io/telliot/pull/347) Added onchain trackers for Uniswap and Balancer.

* [\#385](https://github.com/tellor-io/telliot/pull/385) Added a new psr for Defi Market cap, for id 58. Miners will need to create a free api key in CoinMarketCap pro to be able to read the apis.


### Fixed

## [v5.5.0](https://github.com/tellor-io/telliot/releases/tag/v5.5.0) - 2021.01.18

### Changed

* [\#372](https://github.com/tellor-io/telliot/pull/372) Split the configs of the mine and dataserver command to avoid confusion and be more explicit. This also fixes an issue where can't run a miner and dataserver on the same machine\(now that the miner also runs an HTTP server to expose metrics\). The config format has changed so users need to update their configs. See the `configs/config.json` for an example of the new format.
* [\#374](https://github.com/tellor-io/telliot/pull/374) Changed DEFITVL from median to mean as it has only 2 API endpoints and with mean the calcs return more accurate results.

### Added

### Fixed

## [v5.4.0](https://github.com/tellor-io/telliot/releases/tag/v5.4.0) - 2021.01.13

### Changed

* [\#321](https://github.com/tellor-io/telliot/pull/321) Unified all configuration files. LoggingConfig and LogLevel now reside in the main config file.
* [\#366](https://github.com/tellor-io/telliot/pull/366) Refactored the `index.json` parsing to be more flexible and allow using different parsers. With the notion of `parser` and `param` can allow combining different parsers and parsers parameters. The default is still `jsonPath`, but current users need to rename `jsonPath` to `param` in their `index.json` file.

### Added

* [\#321](https://github.com/tellor-io/telliot/pull/363) `interval` field in the `indexes.json` file. This sets a custom trackerCycle for a specific \(e.g. an `interval: 600` would lead to the API being updated every hour\)
* [\#321](https://github.com/tellor-io/telliot/pull/363) `minSubmitPeriod` field in the `config.json` file. This sets a limit on how often telliot can send submits. The default is 15min which is what the current oracle contract allows.
* [\#339](https://github.com/tellor-io/telliot/pull/339) Initial support for Prometheus metrics.
* [\#340](https://github.com/tellor-io/telliot/pull/340) Manifest files to run in k8s google cloud with Prometheus and Grafana monitoring. The team will run a public telliot miner dashboard at [http://monitor.tellor.io/](http://monitor.tellor.io/)
* [\#334](https://github.com/tellor-io/telliot/pull/334) DEFITVL feed as a new request ID 57. Miners would have to update the binary and `index.json` to be able to submit.

## [v5.3.0](https://github.com/tellor-io/telliot/releases/tag/v5.3.0) - 2020.12.21

### Changed

* [\#317](https://github.com/tellor-io/telliot/pull/317) Removed nodeURL and private key from config file
* [\#318](https://github.com/tellor-io/telliot/pull/318) `indexes.json` file format migrated to JSONPath format.

### Added

* [\#272](https://github.com/tellor-io/telliot/pull/272) Automated Docker images on every push to master and with every tagged release.

## [v5.2.0](https://github.com/tellor-io/telliot/releases/tag/v5.2.0) - 2020.11.12

* [\#254](https://github.com/tellor-io/telliot/pull/254)
  * Added support for expanding variables in the indexer api url.
  * Added config to specify the `.env` file location. The default is now `configs/.env` so users should either specify a custom location in the `config.json` or move it inside the default config folder.

## [v5.0.0](https://github.com/tellor-io/telliot/releases/tag/v5.0.0) - 2020.11.02

### Added

* Profitability calculations which is set through the `ProfitThreshold`\(in percents\) settings in the config,
* Docs how to contribute.
