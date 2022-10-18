//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

//import "hardhat/console.sol";
import "@openzeppelin/contracts/access/Ownable.sol";                  // @WARN: it's direct import change to ../node_modules/ for ABI
//import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";    // @WARN: it's an absolute path witch is required for build abi, binaries and go files

contract TGPassport is Ownable {
   
   
   uint private _passportFee; 
   address private _owner; 

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
   event trustedIndexed(int64 from, int64 indexed to);
   event untrustedIndexed(int64 from, int64 indexed to);


   constructor() Ownable() {
      _passportFee = 1000 wei; // TODO: calculate gas costs
      _owner = owner();
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
      (bool feePaid,) = _owner.call{value: _passportFee}("");
      require(feePaid, "Unable to transfer fee");
   }

   /** 
   *    @notice  This function approving passport (use for bot) which approve that user owns it's tg_id and nicname he want to attach with
   *    @param passportToApprove address of user wallet which attached to him
   */
   function ApprovePassport (address passportToApprove) public onlyOwner {
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
   function DeclinePassport (address passportToDecline) public onlyOwner {
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



      /**
       *  @dev this is PUBLIC function to TRUST username
       * by default you TRUST NO 0NE
       * @param my_username string nickname... can be changed, UNEXPECTING BEHAVIOUR when nickname taken away
       * @param friend_name string nickname of a friend.. can be changed
       */
      function ITrustUsername(string memory my_username, string memory friend_name) public {

         Passport memory my_p = GetPassportByNickName(my_username);
         address my_address = my_p.userAddress;
         require(my_address == msg.sender, "Your current address mismatch with your tgid");
         int64 my_id = my_p.tgId;
         Passport memory to_p = GetPassportByNickName(friend_name);
         address friend_address = to_p.userAddress;
         int64 friend_id = to_p.tgId;
         require(tgIdToAddress[friend_id] == friend_address, "friend username/id missmatch with correspodning address, possible hack");
         int64 to_id = friend_id;
         _iTrustTo(my_id,to_id);
         opinion_changed[my_id].push(to_id);
         emit trustedIndexed(my_id,to_id);
      }


      /**
       *  @dev this is PUBLIC function to TRUST username
       * by default you TRUST NO 0NE
       * @param my_username string nickname... can be changed, UNEXPECTING BEHAVIOUR when nickname taken away
       * @param enemy_name string nickname of a enemy.. can be changed .. 
       */
      function INOTTrustUsername(string memory my_username, string memory enemy_name)  public {

         Passport memory my_p = GetPassportByNickName(my_username);
         address my_address = my_p.userAddress;
         require(my_address == msg.sender, "Your current address mismatch with your tgid");
         int64 my_id = my_p.tgId;
         Passport memory to_p = GetPassportByNickName(enemy_name);
         address enemy_address = to_p.userAddress;
         int64 enemy_id = to_p.tgId;
         require(tgIdToAddress[enemy_id] == enemy_address, "enemy_name username/id missmatch with correspodning address, possible hack");
         int64 to_id = enemy_id;
         _iNotTrust(my_id,to_id);
         opinion_changed[my_id].push(to_id);         
         emit untrustedIndexed(my_id,to_id);
      }


      /**
       *  @notice this is public function to trust some address
       *  this function is more secured than trust some username
       *  @param friend_address is wallet address of a friend
       */
      function ITrustAddress(address friend_address) public {
         Passport memory my_p = GetPassportByAddress(msg.sender);
         address my_address = my_p.userAddress;
         require(my_address == msg.sender, "Your current address mismatch with your registred wallet address");
         int64 my_id = my_p.tgId;
         Passport memory to_p = GetPassportByAddress(friend_address);
         int64 to_id = to_p.tgId;
         _iTrustTo(my_id,to_id);
         opinion_changed[my_id].push(to_id);
         emit trustedIndexed(my_id,to_id);
      }


      /**
       *  @notice this is a public function to NOT to trust some address
       *  @param enemy_address this is a function to not to trust suspiciouse address 
       *  by default you TRUST N0 0NE
       */
      function INOTTrustAddress(address enemy_address) public  {
         Passport memory my_p = GetPassportByAddress(msg.sender);
         address my_address = my_p.userAddress;
         require(my_address == msg.sender, "Your current address mismatch with your registred wallet address");
         int64 my_id = my_p.tgId;
         Passport memory to_p = GetPassportByAddress(enemy_address);
         int64 to_id = to_p.tgId;
         _iNotTrust(my_id,to_id);
         opinion_changed[my_id].push(to_id);
         emit untrustedIndexed(my_id,to_id);
      }

      /**
       *  @dev since telegram bot can obtain tgid directly from dialog we can use this function to set trust using only tgid's
       *  
       */
      function ITrustID(int64 from, int64 to) public {
         Passport memory my_p = GetPassportByTgId(from);
         address my_address = my_p.userAddress;
         require(my_address == msg.sender, "Your current address mismatch with your registred wallet address");
         _iTrustTo(from,to);
         opinion_changed[from].push(to);
         emit trustedIndexed(from,to);
      }

      /**
       *  @dev since telegram bot can obtain tgid directly from dialog we can use this function to set trust using only tgid's
       */
      function INotTrustID(int64 from, int64 to) public {
         Passport memory my_p = GetPassportByTgId(from);
         address my_address = my_p.userAddress;
         require(my_address == msg.sender, "Your current address mismatch with your registred wallet address");
         _iNotTrust(from,to);
         opinion_changed[from].push(to);
         emit untrustedIndexed(from,to);
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

}