#!/usr/bin/env bash -e

# Set up go path
export GOPATH=$(pwd)
export GOBIN=$GOPATH/bin

# DIRS=$(ls src/**/*.go)
DIRS=$(find $(PWD) -type f -name "*.go")

for f in $DIRS; do
    # echo building $f...
    go install $f
done
