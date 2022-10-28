#!/usr/bin/env sh

set -e

_registry="$1"
_tag="$2"
_app="$3"

if [ -z "$_registry" ] || [ -z "$_tag" ]; then
  echo "Please specify image repository and tag."
  exit 0;
fi

# prepare dir ./bin
mkdir -p ./bin

# build demo app
GOOS=linux GOARCH=amd64 go build -tags prod -o bin/backend "./cmd/$_app"
docker build -f "build/package/$_app.Dockerfile" -t "$_registry/$_app:$_tag" .

# clean dir ./bin
rm -rf ./bin
