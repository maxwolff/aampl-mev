// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.6;

interface IUniswapV2 {
    function swapExactTokensForETH(uint amountIn, uint amountOutMin, address[] calldata path, address to, uint deadline) external returns (uint[] memory amounts);
}    