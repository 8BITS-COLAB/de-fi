// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract VaultContract {
    mapping(address => uint256) public balances;
    uint256 totalSupply;

    constructor() {
        totalSupply = 0;
    }

    function createAccount(address _owner) public {
        balances[_owner] = 0;
    }

    function deposit(address _addr, uint256 _amount) public payable {
        balances[_addr] += _amount;
        totalSupply += _amount;
    }

    function withdraw(address _addr, uint256 _amount) public {
        require(balances[_addr] >= _amount);
        balances[_addr] -= _amount;
        totalSupply -= _amount;
    }

    function balanceOf(address _addr) public view returns (uint256) {
        return balances[_addr];
    }

    function getTotalSupply() public view returns (uint256) {
        return totalSupply;
    }
}
