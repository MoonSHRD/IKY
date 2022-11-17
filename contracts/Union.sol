//SPDX-License-Identifier: MIT


// We can look at Governor.sol from OZ, but we need to simply implement off-chain voting by tg api, so we should look at snapshot mechanism



pragma solidity ^0.8.0;

//import "hardhat/console.sol";

// direct imports -- use it for compile contracts and webapp

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/IERC721Enumerable.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";


// relative imports (for building ABI and go) -- use it for build
/*
import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";
import "../node_modules/@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../node_modules/@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "../node_modules/@openzeppelin/contracts/token/ERC721/extensions/IERC721Enumerable.sol";
import "../node_modules/@openzeppelin/contracts/utils/Counters.sol";
import "../node_modules/@openzeppelin/contracts/access/AccessControl.sol";
*/


import "./TGPassport.sol";

contract Union is Ownable, AccessControl {

    using Counters for Counters.Counter;

    uint private _passportFee;
    address private _owner = owner();
    address private murs = 0x383A9e83E36796106EaC11E8c2Fbe8b92Ff46D3a;

    bytes4 private constant _INTERFACE_ID_IERC721ENUMERABLE = 0x780e9d63;

    bytes32 public constant moderator = keccak256("moderator");
    

    // events
    event ApplicationForJoin(int64 chat_id, int64 applier_id,address multy_wallet_address,VotingType vote_type, address voting_token_address, string group_name);
    event ApplicationForJoinIndexed(int64 indexed chat_id, int64 applier_id,address multy_wallet_address,VotingType vote_type, address voting_token_address, string group_name);
    event ApprovedJoin(int chat_id,address multy_wallet_address,VotingType vote_type, address voting_token_address, string group_name);
    event DeclinedApplication(int chat_id,address multy_wallet_address,VotingType vote_type, address voting_token_address, string group_name);



    //
    enum VotingType {erc20, erc20Votes, erc721 }


    // Meta information about dao
    struct DAO {
      address chatOwnerAddress;
      int64 tgId;
      bool valid;
      address multisigAddress;
      VotingType votingType;
      address votingToken;
      string group_name;
               }



    // set passport contract address
    constructor(address passportContract_){
        _passportContract = passportContract_;
        tgpassport = TGPassport(passportContract_);
       // console.logBytes32(moderator);
        _grantRole(DEFAULT_ADMIN_ROLE,msg.sender);
        _grantRole(moderator,msg.sender);

       // console.logBool(flag);
        _grantRole(moderator,murs);
        
       // console.logBool(flag2);

    }

    // TODO: import Multisig contract, make sure we map tgid to multisig contract, not address!
    mapping (int64 => address) public daoAddresses;

    int64[] public Chat_id_array;

    Counters.Counter dao_count;

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
    *   @param votingType_ -- represents voting token's type: 0=erc20 1=erc20Snapshot 2=erc721
    *   @param dao_name_ -- string name of group chat. can be uses as a link (if link is https://t.me/eth_ru then name is @eth_ru)
    */
    function ApplyForUnion (int64 applyerTg, int64 daoTg, address dao_, VotingType votingType_, address votingTokenContract_, string memory dao_name_) public payable {
      // TODO: add require for check if dao is a gnosis safe multisig! (check support interface?)
      // require(...)
      
      // add passport and owner check
        address daoOwner = tgpassport.GetPassportWalletByID(applyerTg);
        require(daoOwner == msg.sender,"User did not registred in TGP");

      require(daoAddresses[daoTg] == address(0x0), "this chat tgid already taken");
      daoAddresses[daoTg] = dao_;      
      bool checkStandard = _checkStandardVotingToken(votingType_, votingTokenContract_);
      require(checkStandard == true,"Contract does not match with corresponding type");

      _passportFee = tgpassport.GetPassportFee();
      daos[dao_] = DAO(msg.sender, daoTg, false, dao_, votingType_, votingTokenContract_, dao_name_);
      (bool feePaid,) = _owner.call{value: _passportFee}("");  
      require(feePaid, "Unable to transfer fee");
      require (msg.value == _passportFee, "Passport fee is not paid");
      emit ApplicationForJoin(daoTg,applyerTg,dao_,votingType_,votingTokenContract_, dao_name_);
      emit ApplicationForJoinIndexed(daoTg,applyerTg,dao_,votingType_,votingTokenContract_, dao_name_);
   }


    
    /**
     *  @dev This function intended to be used by bot, cause only bot can check if tg id of multisig owner is eqal of tg id of chat admin
     *  @param daoAddress address of multisig wallet
     */
    function ApproveJoin(address daoAddress) public onlyRole(moderator) {
      DAO memory org = daos[daoAddress];
      require(org.valid == false, "already has been approved OR didn't applied at all");
      org.valid = true;
      daos[daoAddress] = org;
      dao_count.increment();
      Chat_id_array.push(org.tgId);
      emit ApprovedJoin(org.tgId,org.multisigAddress,org.votingType,org.votingToken, org.group_name);
    }

    /**
     *  @dev function for decline join (for erase unvalid data as example)
     *  @param daoAddress address of multisig
     */
    function DeclineJoin(address daoAddress) public onlyRole(moderator) {
        DAO memory org = daos[daoAddress];
        require(org.valid == false, "already has been approved OR didn't applied at all");
        delete daos[daoAddress];
        delete daoAddresses[org.tgId];
       // daoAddresses[org.tgId] = address(0x0);
        emit DeclinedApplication(org.tgId,org.multisigAddress,org.votingType,org.votingToken, org.group_name);
    }


    /**
     *  @dev internal function to check interface id of voting token contract
     */
    function  _checkStandardVotingToken(VotingType votingType_, address votingTokenContract_) internal view returns (bool success) {
      if (votingType_ == VotingType.erc721) {
      (success) = IERC721Enumerable(votingTokenContract_).
          supportsInterface(_INTERFACE_ID_IERC721ENUMERABLE);
          return success;
      }
      if (votingType_ == VotingType.erc20) {
        // TODO: check this. decimals of standard token should be equal 18. Probably remove this check
        (success) = IERC20Metadata(votingTokenContract_).decimals() == 18;
      }
      // TODO: add check for snapshot

    }


  /**
   *  @dev get multisig address by tgid of chat
   */
  function getDaoAddressbyChatId(int64 chat_id) public view returns (address) {
        address dao = daoAddresses[chat_id];
        return dao;
    }


  function getDaoCount() public view returns (uint256) {
     return dao_count.current();
  }

  function getModeratorIdentifier() public pure returns (bytes32) {
    return moderator;
  }

}