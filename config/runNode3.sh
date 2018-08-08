#!/bin/sh
NODE_ID=3
BIN_NAME="raft-blockchain"

NAME="node${NODE_ID}"
HOST_PORT="127.0.0.1:808${NODE_ID}"
NETWORK_NODES="config/node${NODE_ID}.json"

cd ..
go build
mv $BIN_NAME $NAME
./${NAME} $HOST_PORT $NETWORK_NODES
