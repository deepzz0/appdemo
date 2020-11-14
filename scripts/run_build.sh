#!/usr/bin/env sh

set -e

# prepare dir ./bin
rm -rf ./bin
mkdir ./bin

# build demo app
GOOS=linux GOARCH=amd64 go build -o bin/ ./cmd/demo
docker build -f build/package/demo.Dockerfile -t appdemo:demo .

# clean dir ./bin
rm -rf ./bin
