//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

//import "hardhat/console.sol";
import "@openzeppelin/contracts/access/Ownable.sol";  
import "@openzeppelin/contracts/access/AccessControl.sol";


//import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";    
//import "../node_modules/@openzeppelin/contracts/access/AccessControl.sol";

contract TGPassport is Ownable, AccessControl {
   
   
   uint private _passportFee; 
   address private _owner; 
   bytes32 public constant moderator = keccak256("moderator");
   address private murs = 0x383A9e83E36796106EaC11E8c2Fbe8b92Ff46D3a;
   
   address private bot = 0x0E5279edeD9Fe8281eB0f7277e51068c6DA2fa31;

   struct Passport {
      address userAddress;
      int64 tgId;      // unic Id for telegram (number)
      bool valid;
      address validatorAddress;
      string userName; // can be changed, do not trust it
   }


   //mappings
   mapping(int64 => address) public tgIdToAddress;
   mapping(address => Passport) public passports;
   mapping(string => address) public username_wallets;  // usernames can be changed, do not trust it, use as utility
   
   mapping (int64 => mapping(int64 => bool)) public trust_global; // user id => [] user ids => trust
  
  /**
   *   
   *  1. by defult user  TRUST N0 0NE.
   *  2. we can get int64[] opinion_changed, so we get array of user who express trust/untrust to specific user
   *  3. then we can call GetTrust(user_from,user_to) for each of result from opinion_changed, and get arrays of trusters/untrusters attached to specific user
   *  ..... possibly there is a better way to do it
   */
   mapping (int64 => int64[]) public opinion_changed;
 
   // EVENTS
   //
   event passportApplied(int64 applyerTg, address wallet_address);
   event passportAppliedIndexed(int64 indexed applyerTg, address wallet_address);
   event passportApproved(int64 applyerTg, address wallet_address, address issuer);
   event passportDenied(int64 applyerTg, address wallet);
   event TrustChanged(int64 from, int64 indexed to, bool trust);


   constructor() Ownable() {
      _passportFee = 2000000000000000 wei; 
      _owner = owner();
        _grantRole(DEFAULT_ADMIN_ROLE,msg.sender);
        _grantRole(moderator,msg.sender);
        _grantRole(moderator,bot);
        _grantRole(moderator,murs);
   }


   function _updateAddress(int64 tgId, address userAddress, string memory user_name_) internal {
      require(tgIdToAddress[tgId] == address(0x0), "There's address connected to that TG ID already.");  // if cell is not empty revert
      tgIdToAddress[tgId] = userAddress;
      username_wallets[user_name_] = userAddress;
   }

   /**
   *  @dev This function update user nicname if user change it
   */
   function UpdateUserName(string memory new_user_name_) public {
     Passport memory p = GetPassportByAddress(msg.sender);
     require(p.userAddress == msg.sender, "you don't now own this username");
     p.userName = new_user_name_;
     passports[msg.sender] = p;
   }

   /**
   *   @notice This function for USER who try to obtain some tg_id
   *   @param applyerTg unic id for telegram user, in telegram it's int64 (number)
   *   @param user_name_ is username (like @username)
   **/
   function ApplyForPassport (int64 applyerTg, string memory user_name_) public payable {
      address applyerAddress = msg.sender;      // ЛИЧНАЯ ПОДАЧА ПАСПОРТА В ТРЕТЬЕ ОКОШКО МФЦ
      _updateAddress(applyerTg,applyerAddress,user_name_);  
      require (msg.value == _passportFee, "Passport fee is not paid");

      passports[msg.sender] = Passport(applyerAddress, applyerTg, false, address(0x0),user_name_);
      emit passportApplied(applyerTg, msg.sender);
      emit passportAppliedIndexed(applyerTg, msg.sender);
      (bool feePaid,) = bot.call{value: _passportFee}("");
      require(feePaid, "Unable to transfer fee");
   }

   /** 
   *    @notice  This function approving passport (use for bot) which approve that user owns it's tg_id and nicname he want to attach with
   *    @param passportToApprove address of user wallet which attached to him
   */
   function ApprovePassport (address passportToApprove) public onlyRole(moderator) {
        int64 _tgId = passports[passportToApprove].tgId;
        string memory user_name_ = passports[passportToApprove].userName;
        require(passports[passportToApprove].valid == false, "already approved OR do not exists yet");
        trust_global[_tgId][_tgId] == true;
        passports[passportToApprove] = Passport(passportToApprove, _tgId, true, msg.sender, user_name_);  
        emit passportApproved(_tgId,passportToApprove,msg.sender);
   }

   /**
   *     @notice This function decline application end erase junk data
   *     @param passportToDecline address of user wallet
   */
   function DeclinePassport (address passportToDecline) public onlyRole(moderator) {
      int64 _tgId = passports[passportToDecline].tgId;
      string memory user_name_ = passports[passportToDecline].userName;
      require(passports[passportToDecline].valid == false, "already approved OR do not exists yet"); // it also means that record exists
      delete passports[passportToDecline];
      delete tgIdToAddress[_tgId];
      delete username_wallets[user_name_];
      emit passportDenied(_tgId,passportToDecline);
   }

   /**
    *  @dev This function is a service function which allow Owner to erase already approved passport
    *  and make clean state contract. NOT FOR USE IN PRODUCTION
    */
    function DeletePassport (address passportToDecline) public onlyOwner {
      int64 _tgId = passports[passportToDecline].tgId;
      string memory user_name_ = passports[passportToDecline].userName;
      uint chainID = block.chainid;
      require(chainID == uint(5), "this function work's only for testnet");  
     // require(passports[passportToDecline].valid == false, "already approved OR do not exists yet"); // it also means that record exists
      delete passports[passportToDecline];
      delete tgIdToAddress[_tgId];
      delete username_wallets[user_name_];
      emit passportDenied(_tgId,passportToDecline);
   }  


      /**
       * 
       *  @dev this INTERNAL function is to show trust to other user
       *  @param from tgid user who trust
       *  @param to tgid user who trusted by
       */
      function _iTrustTo(int64 from, int64 to)  internal {
         trust_global[from][to] = true;
      }

      /**
       *   @dev this INTERNAL function is to DISRESPECT youser
       *   by DEFAULT you are TRUST NO 0NE!
       *   @param from tgid user who DONT TRUST
       *   @param to tgid user who looks suspiciouse
       */
      function _iNotTrust(int64 from, int64 to) internal {
         trust_global[from][to] = false;
      }

      /*
      function SetTrustToAddress(address to, bool trust) public {
      
      }
      */

      /**
       *  @dev set trust 'from' tgid 'to' tgid
       */
      function SetTrustToID(int64 from, int64 to, bool trust) public {
         Passport memory from_p = GetPassportByTgId(from);
         address from_address = from_p.userAddress;
         require(from_address == msg.sender, "Your current address mismatch with your registred wallet address"); // we check only registration 'from', 'to' may not be registred
         if (trust == true) {
            _iTrustTo(from,to);
            opinion_changed[from].push(to);     // consider remove
            emit TrustChanged(from,to,trust);
         } else {
            _iNotTrust(from,to);
            opinion_changed[from].push(to);     // consider remove
            emit TrustChanged(from,to,trust);
         }

      }


      /**
       *  @notice get to know if tgid from trust tgid to
       */
      function GetTrust(int64 from, int64 to) public view returns (bool) {
         return trust_global[from][to];
      }

    /**
     *  @dev setting fee for applying for passport
     */
    function SetPassportFee(uint passportFee_) public onlyOwner {
        _passportFee = passportFee_;
    }

    function SetBotAddress(address bot_) public onlyOwner {
      bot = bot_;
    }

    /**
     *  @dev getter to obtain how much user will pay for apply
     */
    function GetPassportFee() public view returns (uint) {
        return _passportFee;
    }


   function GetPassportWalletByID(int64 tgId_) public view returns(address){
      return tgIdToAddress[tgId_];
   }

   function GetTgIdByAddress(address user_wallet) public view returns(int64 tgid) {
      Passport memory p = GetPassportByAddress(user_wallet);
      tgid = p.tgId;
      return tgid;
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

   function GetPassportByTgId(int64 tgId_) public view returns (Passport memory) {
      address wallet = GetPassportWalletByID(tgId_);
      Passport memory p = passports[wallet];
      return p;
   }

   function GetOwner() public view returns(address) {
      return _owner;
   }

   function getModeratorIdentifier() public pure returns (bytes32) {
       return moderator;
  }

}