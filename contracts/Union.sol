//SPDX-License-Identifier: MIT


// We can look at Governor.sol from OZ, but we need to simply implement off-chain voting by tg api, so we should look at snapshot mechanism



pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/IERC721Enumerable.sol";
import "./TGPassport.sol";

contract Union is Ownable {


    uint private _passportFee;
    address private _owner = owner();

    bytes4 private constant _INTERFACE_ID_IERC721ENUMERABLE = 0x780e9d63;

    //
    enum VotingType {erc20, erc721 }


    // Meta information about dao
    struct DAO {
      address chatOwnerAddress;
      string tgId;
      bool valid;
      address multisigAddress;
      VotingType votingType;
      address votingToken;
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
    function applyForUnion (string memory applyerTg, string memory daoTg, address dao_, VotingType votingType_, address votingTokenContract_) public payable {
      // TODO: add require for check if dao is a gnosis safe multisig! (check support interface?)
      // require(...)
      
      // add passport and owner check
        address daoOwner = tgpassport.getPassportWallet(applyerTg);
        require(daoOwner == msg.sender,"User did not registred");
      
        // TODO: add requirements if Multisig owner is eqal daoOwner



      require(daoAddresses[daoTg] == address(0x0), "this chat tgid already taken");
      daoAddresses[daoTg] = dao_;      //  
      bool checkStandard = _checkStandardVotingToken(votingType_, votingTokenContract_);
      require(checkStandard == true,"Contract does not match with corresponding type");


      daos[dao_] = DAO(msg.sender, applyerTg, false, dao_, votingType_, votingTokenContract_);

      (bool feePaid,) = _owner.call{value: _passportFee}("");
      require(feePaid, "Unable to transfer fee");
      require (msg.value == _passportFee, "Passport fee is not paid");
   }


    // This function intended to be used by bot, cause only bot can check if tg id of multisig owner is eqal of tg id of chat admin
    function approveJoin(address daoAddress) public onlyOwner {
      DAO memory org = daos[daoAddress];
      org.valid = true;
      daos[daoAddress] = org;

    }

    function  _checkStandardVotingToken(VotingType votingType_, address votingTokenContract_) internal view returns (bool success) {
      if (votingType_ == VotingType.erc721) {
      (success) = IERC721Enumerable(votingTokenContract_).
          supportsInterface(_INTERFACE_ID_IERC721ENUMERABLE);
          return success;
      }
      if (votingType_ == VotingType.erc20) {
        // TODO: check this. decomals of standard token should be equal 18. Probably remove this check
        (success) = IERC20Metadata(votingTokenContract_).decimals() == 1e18;
      }
    }


}