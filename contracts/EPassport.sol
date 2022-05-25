//SPDX-License-Identifier: MIT
// TODO: documentation

// This contract implements (https://www.facebook.com/100000943480007/posts/pfbid0fZFB2g6wHMfZ2Bu432hZjZJEpmMU1gaYyJAhs7HGfadyhDYPAfNN4KWoPv45pdQSl/?d=n)
// e-Passport is a contract which allow to create electronic passport of citizens with external KYC check of documents

// algorythm of e-passport:
// 1. user applies for passport, submitting (encrypted)archive with documents for KYC check to IPFS. Archive is encrypted by public key of KYC-center 
// 2. contract creates record with passport struct, saving ipfs link, this struct mapped to real passport ID hash.
// 3. KYC center recive event, get archive from IPFS, decrypt it with it's key, operator check documents, set approval or decline
// 4. KYC center can revoke approval






pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

contract EPassport is Ownable {

   using Counters for Counters.Counter;

   uint private _passportFee;
   address private _owner = owner();
   Counters.Counter passport_eid_count;

   struct CitizenPassport {
       // unic order number
     //  uint256 u_counter;
       string ID_hash;      // hash of real passport ID number
       string IPFS_link;    // link to archive with documents
       address user_wallet;     // user wallet address
       bool approved;
       address KYC_operator;


   }

    // from passportID_Hash string to eid
   mapping(string => uint256) hashToeid;
    // from eid to Passport struct
   mapping(uint256 => CitizenPassport) passports;
   // from wallet address to counter
   mapping(address => uint256) eids;

    event passportApplied(string passportIdHash, address wallet_address);
    event passportApproved(string passportIdHash, address wallet_address, address issuer);
    event passportBanned(string passportIdHash, address wallet_address, address kyc_center);

    function applyForPassport (string memory passportIdHash, string memory ipfs_link) public payable {
   //   require (msg.value == _passportFee, "Passport fee is not paid");
      require(hashToeid[passportIdHash]==0,"this passportID number already applied");
      passport_eid_count.increment();
      eids[msg.sender] = passport_eid_count.current();
      passports[passport_eid_count.current()] = CitizenPassport(passportIdHash, ipfs_link,msg.sender,false,address(0x0));
      emit passportApplied(passportIdHash, msg.sender);
     // (bool feePaid,) = _owner.call{value: _passportFee}("");
     // require(feePaid, "Unable to transfer fee");
   }


   function approvePassport (string memory passportIdHash) public onlyOwner {
        uint256 eid = hashToeid[passportIdHash];
        CitizenPassport memory id_card;
        id_card = passports[eid];
        id_card.approved = true;
        id_card.KYC_operator = msg.sender;
        passports[eid] = id_card;
        emit passportApproved(passportIdHash,id_card.user_wallet,msg.sender);
   }

    function getPassportByIdHash(string memory passportIdHash) public view returns (CitizenPassport memory)  {
        CitizenPassport memory id_card;
        uint256 eid = hashToeid[passportIdHash];
        id_card = passports[eid];
        return id_card;
    }

    function getPassportByEid(uint256 eid_) public view returns (CitizenPassport memory) {
        CitizenPassport memory id_card;
        id_card = passports[eid_];
        return id_card;
    }

    function exile (string memory passportIdHash) public onlyOwner {
        uint256 eid = hashToeid[passportIdHash];
        CitizenPassport memory id_card;
        id_card = passports[eid];
        id_card.approved = false;
        id_card.KYC_operator = msg.sender;
        passports[eid] = id_card;
        emit passportBanned(passportIdHash, id_card.user_wallet, msg.sender);
    }

}