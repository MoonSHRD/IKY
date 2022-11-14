// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// in-direct imports

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


// direct imports, use it for go generation
/*
import "../../node_modules/@openzeppelin/contracts/access/Ownable.sol";
import "../../node_modules/@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "../../node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "../../node_modules/@openzeppelin/contracts/token/ERC20/IERC20.sol";
*/

contract ERC20Sample is ERC20, Ownable {
    constructor(string memory name_, string memory symbol_, uint256 initialSupply)
        ERC20(name_, symbol_)
        
    {
        mint(msg.sender,initialSupply);
    }

    function mint(address to, uint256 amount) public virtual onlyOwner {
        super._mint(to, amount);
    }
}
