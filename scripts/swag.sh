#!/usr/bin/env sh

app="$1"
cd cmd/$app && swag init --parseDependency --parseDepth 4;

