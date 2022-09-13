#!/bin/bash


# Allows us to use the view contract without deploying on mainnet

set -u

source .env

PK1=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

cd eth

# --block-time 1
anvil --fork-url $FORK_RPC_URL --no-storage-caching --port $PORT --block-time 15 &
P1=$!

while ! lsof -i :$PORT
do
    echo "...waiting for anvil"
    sleep 1
done
echo http://localhost:${PORT}
forge create src/View.sol:View --private-key $PK1 --rpc-url http://localhost:${PORT} &
P2=$!

#kill both when done
wait $P1 $P2