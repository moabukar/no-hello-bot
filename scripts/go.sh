#!/bin/bash

set -euo pipefail

BIN="discord-bot-go"

# remove old bin directory and recreate it
rm -f bin/*
mkdir -p bin/

XC_OS="$(go env GOOS)"
XC_ARCH="$(go env GOARCH)"

GIT_COMMIT_SHA=$(git rev-parse HEAD)
BUILD_DATE=$(date -u '+%Y-%m-%d %H:%M:%S')

LD_FLAGS="-X 'github.com/moabukar/no-hello-bot/pkg/version.BuildDate=${BUILD_DATE} UTC' -X 'github.com/moabukar/no-hello-bot/pkg/version.CommitSHA=${GIT_COMMIT_SHA}'"

if [ "$#" -ne 1 ];
then
    echo "[-] Error: No positional arguments provided. Specify 'build' or 'run'"
    exit 1

elif [ "$1" == "build" ];
then
    echo "[*] Building go binary"
    GOOS=${XC_OS} GOARCH=${XC_ARCH} go build -o "bin/${BIN}" cmd/discord-bot-go.go

elif [ "$1" == "run" ];
then
    echo "[*] Running go binary"
    GOOS=${XC_OS} GOARCH=${XC_ARCH} go run cmd/discord-bot-go.go -c docs/config.json --log.level=debug

fi