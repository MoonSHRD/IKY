//SPDX-License-Identifier: MIT


// We can look at Governor.sol from OZ, but we need to simply implement off-chain voting by tg api, so we should look at snapshot mechanism



pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./TGPassport.sol";

contract UnionOfDAO is Ownable {


    uint private _passportFee;
    address private _owner = owner();

    // Meta information about dao
    struct DAO {
      address chatOwnerAddress;
      string tgId;
      bool valid;
      address multisigAddress;
               }



    // set passport contract address
    constructor(address passportContract_){
        _passportContract = passportContract_;
        tgpassport = TGPassport(passportContract_);
    }

    // TODO: import Multisig contract, make sure we map tgid to multisig contract, not address!
    mapping (string => address) public daoAddresses;

    // mapping from multisig address to attached meta-info
    mapping(address => DAO) public daos;

    address private _passportContract;
    TGPassport public tgpassport;

    

    /**  This function suggest applying for union for any dao
    *   REQUIREMENTS:
    *   1.dao should have it's multisig address
    *   2.owner of multisig must be registred in Passport contract with it's personal tg_id
    *   3.  this tg_id must be equal to tgid of appling chat admin.
    *   Last check can be done only by oracle
    *   @param applyerTg -- tgid of user who sent apply
    *   @param daoTg -- tgid of chat
    *   @param dao_ -- multisig address
    *
    */
    function applyForUnion (string memory applyerTg, string memory daoTg, address dao_) public payable {
      // TODO: add require for check if dao is a gnosis safe multisig! (check support interface?)
      // require(...)
      
      // add passport and owner check
        address daoOwner = tgpassport.getPassportWallet(applyerTg);
        require(daoOwner == msg.sender,"User did not registred");
      
        // TODO: add requirements if Multisig owner is eqal daoOwner



      require(daoAddresses[daoTg] == address(0x0), "this chat tgid already taken");
      daoAddresses[daoTg] = dao_;      //  
      require (msg.value == _passportFee, "Passport fee is not paid");
      daos[dao_] = DAO(msg.sender, applyerTg, false, dao_);
      (bool feePaid,) = _owner.call{value: _passportFee}("");
      require(feePaid, "Unable to transfer fee");
   }


    // This function intended to be used by bot, cause only bot can check if tg id of multisig owner is eqal of tg id of chat admin
    function approveJoin(address dao_address) public onlyOwner {
      DAO memory org = daos[dao_address];
      org.valid = true;
      daos[dao_address] = org;

    }


}