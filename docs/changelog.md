# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/) and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

NOTE: As semantic versioning states all 0.y.z releases can contain breaking changes in API \(flags, grpc API, any backward compatibility\)

We use _breaking :warning:_ to mark changes that are not backward compatible \(relates only to v0.y.z releases.\)

## Unreleased

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
* Added onchain trackers for Uniswap and Balancer.[\#347](https://github.com/tellor-io/telliot/pull/347)

* Added a new psr for Defi Market cap, for id 58. Miners will need to create a free api key in CoinMarketCap pro to be able to read the apis.[\#385](https://github.com/tellor-io/telliot/pull/385)


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
