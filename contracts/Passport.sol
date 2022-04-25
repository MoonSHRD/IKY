//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "hardhat/console.sol";

contract Passport {
   mapping(string => address) public tgIdToAddress;
   mapping(address => string) public addressToTgId;
 
   function updateAddress(string memory tgId, address userAddress) public {
      require(tgIdToAddress[tgId] == 0x0000000000000000000000000000000000000000, "There's address connected to that TG ID already.");
      tgIdToAddress[tgId] = userAddress;
   }
}