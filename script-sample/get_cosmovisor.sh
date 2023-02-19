#!/usr/bin/env bash

# cosmovisor automatically handles blockchain software updates
wget --auth-no-challenge --no-check-certificate https://github.com/cosmos/cosmos-sdk/releases/download/cosmovisor%2Fv1.3.0/cosmovisor-v1.3.0-linux-amd64.tar.gz
tar xfvz cosmovisor-v1.3.0-linux-amd64.tar.gz cosmovisor
mv cosmovisor bin
rm cosmovisor-v1.3.0-linux-amd64.tar.gz
