#!/bin/bash

# make bindings for abis in "external_abis"
for filename in $(ls -d external_abis/*)
do
    base=$(basename $filename .abi)
    mkdir -p bindings/${base}
    abigen --abi $filename --pkg=${base} --out=bindings/${base}/${base}.go
done

# make bindings for abis in "eth/out" from foundry
for filename in "Arb" 
do

    lower=$(echo ${filename} | awk '{print tolower($0)}')
    mkdir -p bindings/${lower}
    cat eth/out/${filename}.sol/${filename}.json | jq .abi | abigen --abi - --pkg=${lower} --out=bindings/${lower}/${lower}.go
done

# idk why this doesnt work
# abigen --combined-json=eth/out/${filename}.sol/${filename}.json --pkg=dummy --out=Dummy.go