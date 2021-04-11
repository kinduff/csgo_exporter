![](extra/logo.png)

# CSGO Exporter

[![GoDoc](https://godoc.org/github.com/kinduff/csgo_exporter?status.svg)][godoc]
[![Go Report Card](https://goreportcard.com/badge/github.com/kinduff/csgo_exporter)][goreport]
[![Test / Build](https://github.com/kinduff/csgo_exporter/actions/workflows/ci.yml/badge.svg?branch=master)][workflow-c]
[![Release](https://github.com/kinduff/csgo_exporter/actions/workflows/release.yml/badge.svg)][workflow-r]
[![Docker Pulls](https://img.shields.io/docker/pulls/kinduff/csgo_exporter.svg?maxAge=604800)][dockerhub]
![GitHub all releases](https://img.shields.io/github/downloads/kinduff/csgo_exporter/total)

The CSGO Exporter allows to fetch statistics for one player from the [CS:GO][csgo] game.

## Prerequisites

One of the two, depending on your running method.

- [Go][go]
- [Docker][docker]

## Running this exporter

See [Configuration][configuration] in order to set the necessary params to run the exporter.

### Using a binary

You can download the latest version of the binary built for your architecture [here][releases].

### Using Docker

The exporter is also available as a Docker image in [DockerHub][dockerhub] and [Github CR][ghcr]. You can run it using the following example and pass the configuration as environment variables:

```shell
$ docker run \
  --name csgo_exporter \
  -p 7355:7355 \
  -e STEAM_API_KEY=<your-api-key> \
  -e STEAM_NAME=<your-user-name> \
  kinduff/csgo_exporter:latest
```

Alternative, you can use `ghcr.io/kinduff/csgo_exporter` if you want to use the Github Container Registry.

### Using the source

Optionally, you can download and build it from the sources. You have to retrieve the project sources by using one of the following way:

```shell
$ go get -u github.com/kinduff/csgo_exporter
# or
$ git clone https://github.com/kinduff/csgo_exporter.git
```

Install the needed vendors:

```shell
$ GO111MODULE=on go mod vendor
```

Then, build the binary:

```shell
$ go build -o csgo_exporter ./cmd/csgo_exporter
```

## Configuration

This exporter uses environment variables, there are no CLI support for now. The reason behind this is that it's easier to setup and treat it as a running server, instead of a CLI tool.

| Environment variable                    | Description                                                                                | Default                         | Required |
| --------------------------------------- | ------------------------------------------------------------------------------------------ | ------------------------------- | -------- |
| `HTTP_PORT`                             | The port the exporter will be running the HTTP server                                      | 7355<sup id="a1">[1](#f1)</sup> |          |
| `STEAM_API_KEY`                         | Your personal API key from Steam, get one using [this link][steam-api]                     |                                 | ✅       |
| `STEAM_ID`<sup id="a2">[2](#f2)</sup>                              | The Steam ID you want to fetch the data from for the player statistics                     |                                 | ✅       |
| `STEAM_NAME` | If you don't want to provide a `STEAM_ID` you can provide your username, see the footnotes |                                 |          |

## Available Prometheus metrics

| Metric name                | Description                                                                                            |
| -------------------------- | ------------------------------------------------------------------------------------------------------ |
| `csgo_stats_metric`        | All the stats from the player, it includes last_match data, totals per weapon, among other cool things |
| `csgo_achievements_metric` | All achievements done by the player, with value `1` or `0` for achieved or not                         |
| `csgo_news_metric`         | The latest news from the CS:GO community, can be used in a table. Value is an epoch                    |

## Footnotes

- <b id="f1">[1]</b>: This port is being assigned for fun, since the bomb code from Counter Strike is `7355608`. [↩](#a1)
- <b id="f2">[2]</b>: Please note that the `STEAM_ID` environment variable is not required if you provide a `STEAM_NAME`, but this will add 1 HTTP call in order to fetch the SteamID. [↩](#a2)

[configuration]: #configuration
[csgo]: https://store.steampowered.com/app/730/CounterStrike_Global_Offensive
[docker]: https://docs.docker.com
[dockerhub]: https://hub.docker.com/r/kinduff/csgo_exporter
[ghcr]: #ghcr
[go]: https://golang.org
[godoc]: https://godoc.org/github.com/kinduff/csgo_exporter
[goreport]: https://goreportcard.com/report/github.com/gustavo-iniguez-goya/opensnitch
[releases]: https://github.com/kinduff/csgo_exporter/releases
[steam-api]: https://steamcommunity.com/dev/apikey
[workflow-c]: https://github.com/kinduff/csgo_exporter/actions/workflows/ci.yml
[workflow-r]: https://github.com/kinduff/csgo_exporter/actions/workflows/release.yml
