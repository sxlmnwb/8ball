#!/usr/bin/env bash

# set up cosmosvisor directory structure
# put binary in genesis bin    
mkdir cosmovisor
mkdir cosmovisor/genesis
mkdir cosmovisor/genesis/bin
cp 8ball cosmovisor/genesis/bin
mkdir cosmovisor/upgrades

# set inflation to zero instead of defaults
sed -i "s/\"inflation\": \"0.130000000000000000\"/\"inflation\": \"0.000000000000000000\"/" config/genesis.json
sed -i "s/\"inflation_rate_change\": \"0.130000000000000000\"/\"inflation_rate_change\": \"0.000000000000000000\"/" config/genesis.json
sed -i "s/\"inflation_max\": \"0.200000000000000000\"/\"inflation_max\": \"0.000000000000000000\"/" config/genesis.json
sed -i "s/\"inflation_min\": \"0.070000000000000000\"/\"inflation_min\": \"0.000000000000000000\"/" config/genesis.json
sed -i "s/\"goal_bonded\": \"0.670000000000000000\"/\"goal_bonded\": \"0.010000000000000000\"/" config/genesis.json

