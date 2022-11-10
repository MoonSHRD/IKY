// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// in-direct imports
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


// direct imports, use it for go generation
//import "../../node_modules/@openzeppelin/contracts/access/Ownable.sol";
//import "../../node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol";


abstract contract ERC20VotesSample is ERC20Votes, Ownable {
    constructor(uint256 initialSupply)
        ERC20Votes()
        
    {
        mint(msg.sender,initialSupply);
    }

    function mint(address to, uint256 amount) public virtual onlyOwner {
        super._mint(to, amount);
    }
}
