solc --abi Store.sol -o build

abigen --abi=./build/Store.abi --pkg=store --out=Store.go
