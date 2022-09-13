// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.6;

import "./external/UFragmentsPolicy.sol";
import "./external/AaveLibs.sol";
import "./external/DSMath.sol";


interface IPriceOracleGetter {
  function getAssetPrice(address asset) external view returns (uint256);
}

interface ILendingPoolAddressesProvider {
  function getPriceOracle() external view returns (address);
}

interface ILendingPool {
    function getUserAccountData(address user) external view returns (
      uint256 totalCollateralETH,
      uint256 totalDebtETH,
      uint256 availableBorrowsETH,
      uint256 currentLiquidationThreshold,
      uint256 ltv,
      uint256 healthFactor
    );
    function getReserveData(address asset) external view returns (DataTypes.ReserveData memory);
    function getAddressesProvider() external view returns (address);
}

interface IERC20 {
    function balanceOf(address account) external view returns (uint256);
    function transfer(address to, uint256 amount) external returns (bool);
    function allowance(address owner, address spender) external view returns (uint256);
    function approve(address spender, uint256 amount) external returns (bool);
    function totalSupply() external view returns (uint256);
    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) external returns (bool);
}

contract View is DSMath {
    using SafeMath for uint256;
    using SafeMath for uint128;
    using SafeMathInt for int256;
    using UInt256Lib for uint256;
    using ReserveConfiguration for DataTypes.ReserveConfigurationMap;

    address public aampl = 0x1E6bb68Acec8fefBD87D192bE09bb274170a0548;
    ILendingPool public lendingPool = ILendingPool(0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9);
    IERC20 public weth = IERC20(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2);
    IERC20 public uFrags = IERC20(0xD46bA6D942050d489DBd938a2C909A5d5039A161); // ampl token
    UFragmentsPolicy public uPolicy = UFragmentsPolicy(0x1B228a749077b8e307C5856cE62Ef35d96Dca2ea);
    
    uint public baseCpi = 109195000000000000000;// private in policy, grabbed from slot 105 bc its immutable
    uint public borrowPercentOfPower = 9e17;// borrow 80% of max
    uint public borrowDurationSecs = 52560;// assume a borrow for 10mins: 6 * 24 * 365 = 52560

    /* EXTERNAL */
    uint256 private constant DECIMALS = 18;
    int256 private constant ONE = int256(10**DECIMALS);
    uint256 private constant MAX_RATE = 10**6 * 10**DECIMALS;
    uint256 private constant MAX_SUPPLY = uint256(type(int256).max) / MAX_RATE;// MAX_SUPPLY = MAX_INT256 / MAX_RATE

    constructor() {}

    // seconds until earliest rebase 
    function secondsUntilRebase() public view returns (uint) {
        uint lowerOffset = uPolicy.rebaseWindowOffsetSec();
        uint upperOffset = lowerOffset + uPolicy.rebaseWindowLengthSec();
        uint minRebaseTimeIntervalSec = uPolicy.minRebaseTimeIntervalSec();
        uint nowTimestamp = block.timestamp;
        uint lastRebaseTimestamp = uPolicy.lastRebaseTimestampSec();
        return secondsUntilRebasePure(lowerOffset, upperOffset, minRebaseTimeIntervalSec, nowTimestamp, lastRebaseTimestamp);
    }

    // pure for testing
    function secondsUntilRebasePure(
        uint lowerOffset, 
        uint upperOffset, 
        uint minRebaseTimeIntervalSec, 
        uint nowTimestamp, 
        uint lastRebaseTimestamp
    ) public pure returns (uint) {
        uint minRebaseTime = lastRebaseTimestamp + minRebaseTimeIntervalSec;
        if (nowTimestamp > minRebaseTime) {
            minRebaseTime =nowTimestamp;
        }

        uint remainder = minRebaseTime % minRebaseTimeIntervalSec;

        if (remainder < lowerOffset) {
            minRebaseTime = minRebaseTime - remainder + lowerOffset;
        } else if (remainder > upperOffset) {
            minRebaseTime = minRebaseTime + minRebaseTimeIntervalSec - remainder + lowerOffset;
        }
        return minRebaseTime - nowTimestamp;
    }

    function userDebt(address user) public view returns (uint) {
        DataTypes.ReserveData memory reserveData = lendingPool.getReserveData(address(uFrags));
        return IERC20(reserveData.variableDebtTokenAddress).balanceOf(user);
    }

    // expected profit, net of interest paid
    function expectedProfit(address user) public returns (int profitETH, uint borrowAmountAMPL) {
        uint availAMPL = uFrags.balanceOf(aampl);
        (,,uint256 availableBorrowsETH,,,) = lendingPool.getUserAccountData(user);

        // don't borrow full borrow power, dont get liquidated
        availableBorrowsETH = wmul(availableBorrowsETH, borrowPercentOfPower);

        DataTypes.ReserveData memory reserveData = lendingPool.getReserveData(address(uFrags));

        address _addressesProvider = lendingPool.getAddressesProvider();
        address oracle = ILendingPoolAddressesProvider(_addressesProvider).getPriceOracle();
        uint amplPrice = IPriceOracleGetter(oracle).getAssetPrice(address(uFrags));

        uint availAmplinETH = amplPrice.mul(availAMPL).div(10**reserveData.configuration.getDecimals());

        uint borrowAmountETH = availAmplinETH > availableBorrowsETH ? availableBorrowsETH : availAmplinETH;// take the min, cant borrow more than capacity
        borrowAmountAMPL = wdiv(borrowAmountETH, amplPrice);

        uint interestPaid = reserveData.currentVariableBorrowRate.mul(borrowAmountETH).div(borrowDurationSecs).div(RAY);
        (int supplyDelta, uint totalSupply) = getRebasePercentage();

        int amplSupplyAfter = toInt(totalSupply).add(supplyDelta);
        require(amplSupplyAfter > 0, "supply delta greater than supply");
        uint ethBalAfter = borrowAmountETH.mul(uint(amplSupplyAfter)).div(totalSupply);
        profitETH = toInt(ethBalAfter.sub(borrowAmountETH)).sub(toInt(interestPaid));
    }

    // return supplyDelta, and total Supply 
    function getRebasePercentage() public returns (int supplyDelta, uint totalSupply) {
        uint256 cpi;
        bool cpiValid;
        (cpi, cpiValid) = uPolicy.cpiOracle().getData();
        require(cpiValid, "rate not valid");

        uint256 targetRate = cpi.mul(10**DECIMALS).div(baseCpi);
        
        uint256 exchangeRate;
        bool rateValid;
        (exchangeRate, rateValid) = uPolicy.marketOracle().getData();
        require(rateValid, "rate not valid");

        if (exchangeRate > MAX_RATE) {
            exchangeRate = MAX_RATE;
        }

        supplyDelta = computeSupplyDelta(exchangeRate, targetRate);

        if (supplyDelta > 0 && uFrags.totalSupply().add(uint256(supplyDelta)) > MAX_SUPPLY) {
            supplyDelta = (MAX_SUPPLY.sub(uFrags.totalSupply())).toInt256Safe();
        }

        return (supplyDelta, uFrags.totalSupply());
    }

    function toInt(uint256 a) internal pure returns (int256) {
        require(a <= ~(uint256(1) << 255));
        return int256(a);
    }

    /// EXTERNAL: copied from uFrags Policy  /// 

    /**
     * @return Computes the total supply adjustment in response to the exchange rate
     *         and the targetRate.
     * copy function here and make public
     */
    function computeSupplyDelta(uint256 rate, uint256 targetRate) public view returns (int256) {
        if (withinDeviationThreshold(rate, targetRate)) {
            return 0;
        }
        int256 targetRateSigned = targetRate.toInt256Safe();
        int256 normalizedRate = rate.toInt256Safe().mul(ONE).div(targetRateSigned);
        int256 rebasePercentage = uPolicy.computeRebasePercentage(
            normalizedRate,
            uPolicy.rebaseFunctionLowerPercentage(),
            uPolicy.rebaseFunctionUpperPercentage(),
            uPolicy.rebaseFunctionGrowth()
        );

        return uFrags.totalSupply().toInt256Safe().mul(rebasePercentage).div(ONE);
    }

    function withinDeviationThreshold(uint256 rate, uint256 targetRate)
        internal
        view
        returns (bool)
    {
        uint256 absoluteDeviationThreshold = targetRate.mul(uPolicy.deviationThreshold()).div(10**DECIMALS);

        return
            (rate >= targetRate && rate.sub(targetRate) < absoluteDeviationThreshold) ||
            (rate < targetRate && targetRate.sub(rate) < absoluteDeviationThreshold);
    }
}
