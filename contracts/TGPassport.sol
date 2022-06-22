//SPDX-License-Identifier: MIT
// TODO: documentation
pragma solidity ^0.8.0;

//import "hardhat/console.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract TGPassport is Ownable {
   
   
   // TODO set this price in constructor, not in hard-code
   uint private _passportFee; 
   address private _owner; 

   //address public owner_view = owner();

   struct Passport {
      address userAddress;
      string tgId;
      bool valid;
      address validatorAddress;
   }

   //passport Passport;

   mapping(string => address) public tgIdToAddress;
   mapping(address => Passport) public passports;
 
   // EVENTS

   //
   event passportApplied(string applyerTg, address wallet_address);
   event passportApproved(string applyerTg, address wallet_address, address issuer);


   constructor() Ownable() {
      _passportFee = 100 wei;
      _owner = owner();
   }



   // TODO: add functionality to change address
   function _updateAddress(string memory tgId, address userAddress) internal {
      require(tgIdToAddress[tgId] == address(0x0), "There's address connected to that TG ID already.");  // if cell is not empty revert
      tgIdToAddress[tgId] = userAddress;
   }

   function ApplyForPassport (string memory applyerTg) public payable {
      address applyerAddress = msg.sender;      // ЛИЧНАЯ ПОДАЧА ПАСПОРТА В ТРЕТЬЕ ОКОШКО МФЦ
      _updateAddress(applyerTg,applyerAddress);  
      require (msg.value == _passportFee, "Passport fee is not paid");
      passports[msg.sender] = Passport(applyerAddress, applyerTg, false, address(0x0));
      emit passportApplied(applyerTg, msg.sender);
      (bool feePaid,) = _owner.call{value: _passportFee}("");
      require(feePaid, "Unable to transfer fee");
   }

   
   function ApprovePassport (address passportToApprove) public onlyOwner {
        string memory _tgId = passports[passportToApprove].tgId;
        passports[passportToApprove] = Passport(passportToApprove, _tgId, true, msg.sender);  
        emit passportApproved(_tgId,passportToApprove,msg.sender);
   }
   

    function SetPassportFee(uint passportFee_) public onlyOwner {
        _passportFee = passportFee_;
    }

    function GetPassportFee() public view returns (uint) {
        return _passportFee;
    }

   
   function GetPassportWallet(string memory tgId_) public view returns(address){
      return tgIdToAddress[tgId_];
   }

   function GetOwner() public view returns(address) {
      return _owner;
   }

}