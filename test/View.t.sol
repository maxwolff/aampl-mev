// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "../src/View.sol";

contract ViewTest is Test {
    uint256 mainnetFork;
    View public arbView;

    // test user with assets on aave
    address user = 0xb394e0A990A577F277dDB65547b6399898D78051;
    
    function setUp() public {
        string memory rpcUrl = vm.envString("FORK_RPC_URL");
        mainnetFork = vm.createFork(rpcUrl);
        vm.selectFork(mainnetFork);
        arbView = new View();
    }

    // non deterministic test, just to eyeball correctness
    function testEx() public {
        (int delta, uint totalSupply) = arbView.getRebasePercentage();
        emit log_uint(totalSupply);
        emit log_int(delta);
        
        (int profit, uint amt) = arbView.expectedProfit(user);
        emit log_int(profit);
        emit log_uint(amt);
    }

    function testSecondsTilRebasePure() public {
        // waits for cool off period, rebase at t_15
        uint secondsUntilRebase = arbView.secondsUntilRebasePure(3, 7, 10, 9, 5);
        assertEq(secondsUntilRebase, 6);

        // above window, rebase at t_33 
        secondsUntilRebase = arbView.secondsUntilRebasePure(3, 7, 10, 28, 1);
        assertEq(secondsUntilRebase, 5);

        // above window, rebase at t_23
        secondsUntilRebase = arbView.secondsUntilRebasePure(3, 7, 10, 21, 1);
        assertEq(secondsUntilRebase, 2);

        // in window, rebase at t_25
        secondsUntilRebase = arbView.secondsUntilRebasePure(3, 7, 10, 25, 1);
        assertEq(secondsUntilRebase, 0);

    }


    function testSecondsTilRebase() public {
        uint secondsUntilRebase = arbView.secondsUntilRebase();
        emit log_uint(secondsUntilRebase);
    }

}
