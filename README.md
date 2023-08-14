# Node entrypoint

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/quantcdn/entrypoint-node)
[![Go Report Card](https://goreportcard.com/badge/github.com/quantcdn/entrypoint-node)](https://goreportcard.com/report/github.com/quantcdn/entrypoint-node)
[![Coverage Status](https://coveralls.io/repos/github/quantcdn/entrypoint-node/badge.svg?branch=main)](https://coveralls.io/github/quantcdn/entrypoint-node?branch=main)
[![Release](https://img.shields.io/github/v/release/quantcdn/entrypoint-node)](https://github.com/quantcdn/entrypoint-node/releases/latest)

## Installation

As this is intended to be a docker entrypoint the preferred way to install is using with a dockerfile.

```Dockerfile
COPY --from=ghcr.io/quantcdn/entrypoint-node:latest /usr/local/bin/entrypoint-node /usr/local/bin/entrypoint-node
```

This can be run directly from the docker image:

```sh
docker run --rm ghcr.io/quantcdn/entrypoint-node:latest entrypoint-node --version
```

## Usage

```
$ entrypoint-node --help                                                                                                                             
usage: entrypoint-node [<flags>] [<commands>...]

Flags:
  --help      Show context-sensitive help (also try --help-long and --help-man).
  --dir=DIR   Directory to execute node commands in.
  --url=URL   The backend url (optional).
  --retry=10  Times to retry the backend connection.
  --delay=5   Delay between backend requests.
  --version   Show application version.

Args:
  [<commands>]  Node JS commands to execute.
```

For example to execute `build` and `start` after a backend connection has been established.

```
entrypoint-node --url http://localhost build start
```

## Local development

### Build
```sh
git clone git@github.com:quantcdn/entrypoint-node.git && cd entrypoint-node
go generate ./...
go build -ldflags="-s -w" -o build/entrypoint-node .
go run . -h
```

### Run tests
```sh
go test -v ./... -coverprofile=build/coverage.out
```

View coverage results:
```sh
go tool cover -html=build/coverage.out
```

### Documentation
```sh
cd docs
npm install
npm run dev
```