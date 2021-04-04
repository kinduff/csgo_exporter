![](extra/logo.png)

# CSGO Exporter

[![.github/workflows/release.yml](https://github.com/kinduff/csgo_exporter/actions/workflows/release.yml/badge.svg)](https://github.com/kinduff/csgo_exporter/actions/workflows/release.yml)
[![Docker Pulls](https://img.shields.io/docker/pulls/kinduff/csgo_exporter.svg?maxAge=604800)][hub]

The CSGO Exporter allows to fetch statistics for one or more players from the [CS:GO](https://store.steampowered.com/app/730/CounterStrike_Global_Offensive/) game by [Valve](https://www.valvesoftware.com/en/).

## Running this software

### From binaries

Download the most suitable binary from [the releases tab](https://github.com/kinduff/csgo_exporter/releases)

Then:

```shell
./csgo_exporter <flags>
```

### Using the docker image

```shell
docker run --rm -d -p 7355:7355 --name csgo_exporter kinduff/blackbox-exporter:latest -steamid=1234567890
```

[hub]: https://hub.docker.com/r/kinduff/csgo_exporter/
