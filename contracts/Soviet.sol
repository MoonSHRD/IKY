//SPDX-License-Identifier: MIT


// We can look at Governor.sol from OZ, but we need to simply implement off-chain voting by tg api, so we should look at snapshot mechanism



pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./TGPassport.sol";

contract Soviet is Ownable {


    uint private _passportFee;
    address private _owner = owner();

    constructor(address passportContract_){
        _passportContract = passportContract_;
        tgpassport = TGPassport(passportContract_);
    }

    // TODO: import Multisig contract, make sure we map tgid to multisig contract, not address!
    mapping (string => address) public daoAddresses;

    address private _passportContract;
    TGPassport public tgpassport;

    


    function applyForSoviet (string memory applyerTg, string memory daoTg, address dao_) public payable {
      // TODO: add require for check if dao is a gnosis safe multisig! (check support interface?)
      // require(...)
      
      // add passport and owner check
        address daoOwner = tgpassport.getPassportWallet(applyerTg);
        require(daoOwner == msg.sender,"User did not registred");
      
        // TODO: add requirements if Multisig owner is eqal daoOwner



      require(daoAddresses[daoTg] == address(0x0), "this chat tgid already taken");
      daoAddresses[daoTg] = dao_;      //  
      require (msg.value == _passportFee, "Passport fee is not paid");
    //  passports[msg.sender] = Passport(applyerAddress, applyerTg, false, address(0x0));
      (bool feePaid,) = _owner.call{value: _passportFee}("");
      require(feePaid, "Unable to transfer fee");
   }



}