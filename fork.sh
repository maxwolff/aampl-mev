#!/bin/bash

# sudo kill -9 `sudo lsof -t -i:8546`

set -u

source .env

PORT=8546
PK1=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

cd eth

anvil --fork-url $FORK_RPC_URL --no-storage-caching --port $PORT &
P1=$!

while ! lsof -i :$PORT
do
    echo "...waiting for anvil"
    sleep 1
done
echo http://localhost:${PORT}
forge create src/Arb.sol:Arb --private-key $PK1 --rpc-url http://localhost:${PORT} &
P2=$!

#kill both when done
wait $P1 $P2