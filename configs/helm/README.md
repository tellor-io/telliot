# telliot

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 5.8.0](https://img.shields.io/badge/AppVersion-5.8.0-informational?style=flat-square)

A Helm chart for Kubernetes

## Values

| Key                              | Type   | Default                               | Description                                                             |
| -------------------------------- | ------ | ------------------------------------- | ----------------------------------------------------------------------- |
| alertmanager.bot.container.image | string | `"metalmatze/alertmanager-bot:0.4.3"` | Docker image for alertmanager                                           |
| alertmanager.bot.container.port  | int    | `8080`                                |                                                                         |
| alertmanager.bot.enabled         | bool   | `false`                               | Whether to enable alertmanager                                          |
| alertmanager.bot.service.port    | int    | `8080`                                |                                                                         |
| alertmanager.bot.storage         | string | `"1Gi"`                               |                                                                         |
| alertmanager.bot.telegram.admin  | string | `nil`                                 | Telegram admin username                                                 |
| alertmanager.bot.telegram.token  | string | `nil`                                 | Telegram token                                                          |
| alertmanager.container.image     | string | `"prom/alertmanager:v0.19.0"`         | Docker image for alertmanager                                           |
| alertmanager.container.port      | int    | `9093`                                |                                                                         |
| alertmanager.enabled             | bool   | `false`                               |                                                                         |
| alertmanager.service.port        | int    | `9093`                                |                                                                         |
| grafana.container.image          | string | `"grafana/grafana:7.3.6"`             | Docker image for grafana                                                |
| grafana.container.port           | int    | `3000`                                |                                                                         |
| grafana.enabled                  | bool   | `false`                               |                                                                         |
| grafana.ingress.class            | string | `"nginx"`                             | Ingress class to use for grafana                                        |
| grafana.ingress.hostname         | string | `"monitor.tellor.io"`                 | Hostname to use for accessing grafana                                   |
| grafana.ingress.path             | string | `"/"`                                 | Subpath to access grafana                                               |
| grafana.ingress.tls.enabled      | bool   | `false`                               | Enable/Disable TLS for grafana                                          |
| grafana.ingress.tls.secret       | string | `"grafana-tls-secret"`                | Name of TLS secret to use for grafana                                   |
| grafana.persist                  | bool   | `true`                                | Enable persistance for grafana configuration                            |
| grafana.serveFromSubPath         | bool   | `false`                               |                                                                         |
| grafana.service.port             | int    | `80`                                  |                                                                         |
| grafana.storage                  | string | `"5Gi"`                               | Grafana persistent storage size                                         |
| prometheus.container.image       | string | `"prom/prometheus:v2.24.0"`           | Docker image for prometheus                                             |
| prometheus.container.port        | int    | `9090`                                |                                                                         |
| prometheus.enabled               | bool   | `false`                               |                                                                         |
| prometheus.persist               | bool   | `true`                                | Enable persistance for prometheus data                                  |
| prometheus.service.port          | int    | `9090`                                |                                                                         |
| prometheus.storage               | int    | `50`                                  | Prometheus storage size in GB                                           |
| telliot.container.image          | string | `"tellor/telliot:master"`             | Docker image for telliot                                                |
| telliot.container.port           | int    | `9090`                                |                                                                         |
| telliot.modes                    | list   | `["mine"]`                            | Array of commands to spawn separate instances of telliot instances with |
| telliot.service.port             | int    | `9090`                                |                                                                         |
| telliot.storage                  | string | `"2Gi"`                               | telliot persistent storage size                                         |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.5.0](https://github.com/norwoodj/helm-docs/releases/v1.5.0)
