TO DO
check rebase % formula

- @notice Initiates a new rebase operation, provided the minimum time period has elapsed.
-
- @dev The supply adjustment equals (\_totalSupply \* DeviationFromTargetRate) / rebaseLag
-      Where DeviationFromTargetRate is (MarketOracleRate - targetRate) / targetRate
-      and targetRate is CpiOracleRate / baseCpi
  \*/

![](https://i.imgur.com/pPOMqxK.png)
rebase % = (((oracle rate - price target) / price target) _ 100 ) / 10
ex
(((1504470070677221900 - 104406333333333340000 )/ 104406333333333340000 _ 100 ) / 10)
1504470070677221900
104406333333333340000
check what the oracle actually and what sources it gets

setup blocknative to watch mempool
watch for add liquidity OR return borrowed ampl

fire off borrow tx to flashbots if expected value is positive

repay loan instantly after rebase (ideally bundle) and insta dump the profits on best AMM (or just hold the free ampl idk doesnt matter)

IMPORTANT
although risk is v.low, add a liquidation/health factor check every so often just to be giga safe

tx conditions
value of amount of AMPL available to borrow > expected % increase of rebase - gas cost of borrow tx - gas cost of repay tx - interest we’ll pay holding the ampl until rebase

notes
https://docs.aave.com/developers/the-core-protocol/lendingpool
https://app.aave.com/reserve-overview/AMPL?pool=AaveV2

deposit is ppl adding the real token to the pool
assetaddy, amont, toaddy, 0

withdraw is ppl removing the real token from the pool
asset (real asset addy), amount, toaddy

borrow is ppl borrowing
assetaddy, amount, interestratemode, 0, debtaddress
cba learning about interest rate mode, fine to set variable and assume perma 48% for calc

debtaddress is either msg.sender if self funding, or address delegated if delegating

repay
assetaddy, amount, ratemode, debtaddress

// view
getReserveData
configuration
liquidityIndex
variableBorrowIndex
currentLiquidityRate
currentVariableBorrowRate
currentStableBorrowRate
lastUpdateTimestamp
aTokenAddress
stableDebtTokenAddress
variableDebtTokenAddress
interestRateStrategyAddress
id

addresses
rebase contract
https://etherscan.io/address/0x6fb00a180781e75f87e2b690af0196baa77c7e7c#code
can be called at 2AM my time, think 3AM CET is what it’s set to

aave lending pool (proxy)
https://etherscan.io/address/0x7d2768de32b0b80b7a3454c06bdac94a69ddc7a9

AMPL token
https://etherscan.io/address/0xd46ba6d942050d489dbd938a2c909a5d5039a161

https://etherscan.io/address/0x1b228a749077b8e307c5856ce62ef35d96dca2ea#readProxyContract
proxy for smth related, gives addess of oracle (marketOracle)
check what rebaseLag is #todo

https://etherscan.io/address/0x99c9775e076fdf99388c029550155032ba2d8914#readContract
market oracle, has 2 providers

providers
https://etherscan.io/address/0x8844dfd05AC492D1924Ad27ddD5e690B8E72D694

https://etherscan.io/address/0xfc4b1Ce32ed7310028DCC0d94C7B3D96dCd880e0#code
currentAnswer
updated ~1.5 hrs before rebase available

https://www.ampleforth.org/dashboard/
