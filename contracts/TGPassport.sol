//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract TGPassport is Ownable {
   
   uint passportFee;

   struct Passport {
      address userAddress;
      string tgId;
      bool valid;
      address validatorAddress;
   }

   //passport Passport;

   mapping(string => address) public tgIdToAddress;
   mapping(address => Passport) public passports;
 


   function updateAddress(string memory tgId, address userAddress) internal {
      require(tgIdToAddress[tgId] == address(0x0), "There's address connected to that TG ID already.");  // if cell is not empty revert
      tgIdToAddress[tgId] = userAddress;
   }

   function applyForPassport (string memory applyerTg) public payable {
    //  address passportCouncil = owner();
      address applyerAddress = msg.sender;      // ЛИЧНАЯ ПОДАЧА ПАСПОРТА В ТРЕТЬЕ ОКОШКО МФЦ
      updateAddress(applyerTg,applyerAddress);  
    //  require (tgIdToAddress[applyerTg] == msg.sender, "Address and Telegram ID do not match");
      require (msg.value == passportFee, "Passport fee is not paid");
      passports[msg.sender] = Passport(applyerAddress, applyerTg, false, address(0x0));
      address _owner = owner();
      _owner.call{value: passportFee}('');
   }

   function approvePassport (address passportToApprove) public onlyOwner {
        string memory _tgId = passports[passportToApprove].tgId;
        passports[passportToApprove] = Passport(passportToApprove, _tgId, true, msg.sender);  
   }
   

    function setPassportFee(uint _passportFee) public onlyOwner {
        passportFee = _passportFee;
    }

    function getPassportFee() public view returns (uint) {
        return passportFee;
    }
}