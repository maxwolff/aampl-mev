## Ampleforth Arbitrage MEV Bot

If the upcoming ampleforth rebase would be profitable, borrow from aave right before the rebase, wait until it does happen, and then repay the loan.
Uses a View contract inside `/eth` as a utility to figure out when the rebase is happening, and if the trade would be profitable.

Uses a local anvil rpc forked from mainnet to deploy the view contract and read values from chain. Uses a ws endpoint to read events. Sends via flashbots private transaction.

## Usage

to install:

- install [forge](https://book.getfoundry.sh/getting-started/installation)
- cd eth
- forge install

in root

- in new tab, `./fork.sh`
- go build
- go run .

### TODO:

- simulate transactions before sent
