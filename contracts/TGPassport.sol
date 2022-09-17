//SPDX-License-Identifier: MIT
// TODO: documentation
pragma solidity ^0.8.0;

//import "hardhat/console.sol";
import "@openzeppelin/contracts/access/Ownable.sol";                  // @WARN: it's direct import change to ../node_modules/ for ABI
//import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";    // @WARN: it's an absolute path witch is required for build abi, binaries and go files

contract TGPassport is Ownable {
   
   
   uint private _passportFee; 
   address private _owner; 

   struct Passport {
      address userAddress;
      string tgId;
      bool valid;
      address validatorAddress;
      string userName; // can be changed, do not trust it
   }

   //mappings
   mapping(string => address) public tgIdToAddress;
   mapping(address => Passport) public passports;
   mapping(string => address) public username_wallets;  // usernames can be changed, do not trust it, use as utility
 
   // EVENTS

   //
   event passportApplied(string applyerTg, address wallet_address);
   event passportApproved(string applyerTg, address wallet_address, address issuer);
   event passportDenied(string applyerTg, address wallet);


   constructor() Ownable() {
      _passportFee = 1000 wei; // TODO: calculate gas costs      
      _owner = owner();
   }


   function _updateAddress(string memory tgId, address userAddress, string memory user_name_) internal {
      require(tgIdToAddress[tgId] == address(0x0), "There's address connected to that TG ID already.");  // if cell is not empty revert
      tgIdToAddress[tgId] = userAddress;
      username_wallets[user_name_] = userAddress;
   }

   // This function update user nicname if user change it
   function UpdateUserName(string memory new_user_name_) public {
     Passport memory p = GetPassportByAddress(msg.sender);
     require(p.userAddress == msg.sender, "you don't now own this username");
     p.userName = new_user_name_;
     passports[msg.sender] = p;
   }

   // This function for USER who try to obtain some tg_id
   function ApplyForPassport (string memory applyerTg, string memory user_name_) public payable {
      address applyerAddress = msg.sender;      // ЛИЧНАЯ ПОДАЧА ПАСПОРТА В ТРЕТЬЕ ОКОШКО МФЦ
      
      // TODO: Discuss and remove function from smart contract because passport can be changed or be broken picture
      //       In this case we should reapply and reapprove new picture. Address updating should be in ApprovePassport and
      //       should be applyed for approved passports only
      _updateAddress(applyerTg,applyerAddress,user_name_);  
      
      require (msg.value == _passportFee, "Passport fee is not paid");
      passports[msg.sender] = Passport(applyerAddress, applyerTg, false, address(0x0),user_name_);
      emit passportApplied(applyerTg, msg.sender);
      (bool feePaid,) = _owner.call{value: _passportFee}("");
      require(feePaid, "Unable to transfer fee");
   }

   // This function approving passport (use for bot) which approve that user owns it's tg_id and nicname he want to attach with
   function ApprovePassport (address passportToApprove) public onlyOwner {
        string memory _tgId = passports[passportToApprove].tgId;
        string memory user_name_ = passports[passportToApprove].userName;
        require(passports[passportToApprove].valid == false, "already approved OR do not exists yet");
        passports[passportToApprove] = Passport(passportToApprove, _tgId, true, msg.sender, user_name_);  
        emit passportApproved(_tgId,passportToApprove,msg.sender);
   }

   // This function decline application end erase junk data
   function DeclinePassport (address passportToDecline) public onlyOwner {
      string memory _tgId = passports[passportToDecline].tgId;
      string memory user_name_ = passports[passportToDecline].userName;
      require(passports[passportToDecline].valid == false, "already approved OR do not exists yet"); // it also means that record exists
      delete passports[passportToDecline];
      delete tgIdToAddress[_tgId];
      delete username_wallets[user_name_];
      emit passportDenied(_tgId,passportToDecline);
   }
   

    function SetPassportFee(uint passportFee_) public onlyOwner {
        _passportFee = passportFee_;
    }

    function GetPassportFee() public view returns (uint) {
        return _passportFee;
    }

   
   function GetPassportWalletByID(string memory tgId_) public view returns(address){
      return tgIdToAddress[tgId_];
   }

   function GetPassportByAddress(address user_wallet) public view returns(Passport memory) {
      Passport memory p = passports[user_wallet];
      return p;
   }

   function GetWalletByNickName(string memory user_name_) public view returns (address) {
      return username_wallets[user_name_];
   }

   function GetPassportByNickName(string memory user_name_) public view returns (Passport memory) {
      address wallet_ = GetWalletByNickName(user_name_);
      Passport memory p = passports[wallet_];
      return p;
   }

   function GetOwner() public view returns(address) {
      return _owner;
   }

}