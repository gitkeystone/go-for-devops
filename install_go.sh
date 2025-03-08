#!/usr/bin/env bash

read -p "Please input version [1.24.1]:" VERSION
VERSION=${VERSION:-"1.24.1"}

curl -LOj https://go.dev/dl/go${VERSION}.linux-amd64.tar.gz
rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go${VERSION}.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin && go version

rm go${VERSION}.linux-amd64.tar.gz
