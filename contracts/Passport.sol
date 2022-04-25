//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract TGPassport is Ownable {
   
   uint passportFee;

   struct passport {
      address userAddress;
      string tgId;
      bool valid;
      address validatorAddress;
   }

passport Passport;

   mapping(string => address) public tgIdToAddress;
   mapping(address => passport) public passports;
 


   function updateAddress(string memory tgId, address userAddress) public {
      require(tgIdToAddress[tgId] == 0x0000000000000000000000000000000000000000, "There's address connected to that TG ID already.");
      tgIdToAddress[tgId] = userAddress;
   }

   function applyForPassport (address applyerAddress, string memory applyerTg) public payable {
      address passportCouncil = owner();
      require (tgIdToAddress[applyerTg] == msg.sender, "Address and Telegram ID do not match");
      require (msg.value == passportFee, "Passport fee is not paid");
      passports[msg.sender] = passport(applyerAddress, applyerTg, false, 0x0000000000000000000000000000000000000000);
      passportCouncil.call{value: passportFee}('');
   }

   function approvePassport (address passportToApprove) public onlyOwner {
        string memory _tgId = passports[passportToApprove].tgId;
        passports[passportToApprove] = passport(passportToApprove, _tgId, true, msg.sender);  
   }
   

    function setPassportFee(uint _passportFee) public onlyOwner {
        passportFee = _passportFee;
    }

    function getPassportFee() public view returns (uint) {
        return passportFee;
    }
}