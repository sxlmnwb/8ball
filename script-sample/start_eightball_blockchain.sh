#!/usr/bin/env bash

# Install cosmovisor and put it in bin/
FILE=bin/cosmovisor
if [ ! -f "$FILE" ]; then
  ./scripts/get_cosmovisor.sh
fi

FILE=$HOME/.8ball/cosmovisor/genesis/bin/8ball
if [ ! -f "$FILE" ]; then
    git checkout genesis
    make build-bin
    mkdir -p $HOME/.8ball/cosmovisor/genesis/bin
    cp 8ball $HOME/.8ball/cosmovisor/genesis/bin/8ball
fi

# Every git TAG in the repository that starts with v, for example "v2.4.5" represents an upgrade version

for TAG in `git tag -l 'v*'`
do
    FILE=$HOME/.8ball/cosmovisor/upgrades/$TAG/bin/8ball
    if [ ! -f "$FILE" ]; then
        git checkout $TAG
        make build-bin
        mkdir -p $HOME/.8ball/cosmovisor/upgrades/$TAG/bin
        cp 8ball $HOME/.8ball/cosmovisor/upgrades/$TAG/bin/8ball
    fi
done

# remember to checkout main after building
git checkout main


# Start node with cosmovisor
DAEMON_HOME=~/.8ball DAEMON_NAME=8ball DAEMON_ALLOW_DOWNLOAD_BINARIES=false DAEMON_RESTART_AFTER_UPGRADE=true DAEMON_RESTART_DELAY=1s DAEMON_POLL_INTERVAL=1s ./bin/cosmovisor run --log_level info start
