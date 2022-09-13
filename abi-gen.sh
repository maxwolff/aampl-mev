#!/bin/bash

# make bindings for abis in "eth/out" from foundry
for filename in "View" "ILendingPool" "IPolicy"
do

    lower=$(echo ${filename} | awk '{print tolower($0)}')
    mkdir -p bindings/${lower}
    cat eth/out/${filename}.sol/${filename}.json | jq .abi | abigen --abi - --pkg=${lower} --out=bindings/${lower}/${lower}.go
done
