#!/bin/bash

cd eth
# forge script script/Arb.s.sol > f.txt
RES=""
while read line; do
    if [[ $line == supplyDelta* ]] || [[ $line == totalSupply* ]] ;
    then 
        IFS=' '
        read -a strarr <<< "$line"
        echo "${strarr[2]}"
    fi
done < f.txt

# rm f.txt
