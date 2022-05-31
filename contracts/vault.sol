// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract VaultContract {
    mapping(address => uint256) public balances;
    uint256 totalSupply;

    constructor() {
        totalSupply = 0;
    }

    event Transfer(address indexed from, address indexed to, uint256 amount);
    event Deposit(address indexed account, uint256 amount);
    event Withdraw(address indexed account, uint256 amount);
    event CreateAccount(address indexed account, uint256 amount);

    function createAccount(address _addr, uint256 _amount) public {
        balances[_addr] = _amount;

        totalSupply += _amount;

        emit CreateAccount(_addr, _amount);
    }

    function deposit(address _addr, uint256 _amount) public payable {
        balances[_addr] += _amount;
        totalSupply += _amount;

        emit Deposit(_addr, _amount);
    }

    function withdraw(address _addr, uint256 _amount) public {
        require(balances[_addr] >= _amount);
        balances[_addr] -= _amount;
        totalSupply -= _amount;

        emit Withdraw(_addr, _amount);
    }

    function transfer(
        address _from,
        address _to,
        uint256 _amount
    ) public {
        require(balances[_from] >= _amount);
        balances[_from] -= _amount;
        balances[_to] += _amount;

        emit Transfer(_from, _to, _amount);
    }

    function balanceOf(address _addr) public view returns (uint256) {
        return balances[_addr];
    }

    function getTotalSupply() public view returns (uint256) {
        return totalSupply;
    }
}
