# Holodule Bot

This is an application that reads the holodule distribution schedule, gets the broadcast frame information by scraping, and notifies the user with a Discord bot.

## Install

```shell
go mod download
```

## Build

```shell
make build
```

## Debug

```shell
make debug-daily
# or
make debug-coming-soon
```

## Run

```shell
./bin/notify-daily
# or
./bin/notify-coming-soon
```

## Migration

```shell
./scripts/goose.sh up
# or
./scripts/goose.sh down
```
