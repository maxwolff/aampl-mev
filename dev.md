- AMPL Token Contract,
  - proxy `0xD46bA6D942050d489DBd938a2C909A5d5039A161`
  - impl `0xD0e3F82ab04B983C05263cF3BF52481FbAa435b1`
  - `rebase(uint256 epoch, int256 supplyDelta) onlyMonetaryPolicy`
    - sets new total supply

MonetaryPolicy

- proxy `0x1B228a749077b8e307C5856cE62Ef35d96Dca2ea`
- impl `0xe345465b36cc55275250f30bd32127d3bfe45d53`
- `rebase() external onlyOrchestrator`
  - calls oracle, does some math, calls ampl token

Orchestrator

- `0x6FB00a180781E75F87E2B690Af0196bAa77C7e7C`
- msg.sender == tx.origin
- for tx, call tx
  https://mainnet.infura.io/v3/a92bf60569df4e02b242dbd8357fffac

aAMPL

- `0x1E6bb68Acec8fefBD87D192bE09bb274170a0548`

v2 Lending Pool
`0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9`

## TODO:

- go should know about rebase times

- loop

  - get policy.rebaseState

    - if shortly before rebase time:
      - if profitable
        - borrow .75x max capacity thru fb

  - if rebase log is seen, trigger repay (flashbots bundle) && dump
  - health check

=====================================

## before

- deploy a multicall
- calculate gas used

## off chain

- loop
  - sim rebase tx
  - sim arb tx
  - if success
    - send bundle

## on chain

- ampl.rebase()

- uniswap the excess ampl to eth
- require excess eth > gas cost

filepaths=$(ls external_abis | xargs -I {} echo external_abis/{})
