# Golang Project Template

## Introduction

This is the starting point for a Golang project with Gometalinter, Govendor, TDD, Precommit git hooks, security best practices, etc.

## Contents

- [Install](#install)
- [Build](#build)
- [Run and Test](#run-and-test)
- [cURL](#curl)
- [TODO](#todo)

## Install

```bash
# Govendor
go get -u github.com/kardianos/govendor
govendor sync
```

## Build

```bash
./build.sh
```

## Run and Test

### Mock

```bash
# Run
docker-compose -f mock.docker-compose.yml up -d

# Test
curl -d '{"key":"name","val":"ryan"}' localhost:7100/set
curl -d '{"key":"name"}'              localhost:7100/get

# Stop
docker-compose -f mock.docker-compose.yml down
```

### Redis

```bash
# Run
docker-compose -f redis.docker-compose.yml up -d

# Test
curl -d '{"key":"name","val":"ryan"}' localhost:7100/set
curl -d '{"key":"name"}'              localhost:7100/get

# Stop
docker-compose -f redis.docker-compose.yml down
```
