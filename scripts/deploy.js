// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const { ethers } = require("hardhat");
const { toWei, fromWei } = require("./lib.js");
const hre = require("hardhat");

async function main() {
  // Hardhat always runs the compile task when running scripts with its command
  // line interface.
  //
  // If this script is run directly using `node` you may want to call compile
  // manually to make sure everything is compiled
  // await hre.run('compile');

  console.log(hre.network.name);
  let murs_account = ethers.utils.getAddress("0x383A9e83E36796106EaC11E8c2Fbe8b92Ff46D3a");
  let account_owner = await hre.ethers.getSigner();
  const balance = await ethers.provider.getBalance(account_owner.address);

  console.log(ethers.utils.formatEther(balance), "ETH");


  let owner;
  owner = await hre.ethers.getSigner();
  console.log("owner address:", owner.address);


  // We get the contract to deploy
  const TGPassport = await hre.ethers.getContractFactory("TGPassport");
  const tgpassport = await TGPassport.deploy();

  await tgpassport.deployed();
  console.log("TGPassport deployed to:", tgpassport.address);

  const Union = await hre.ethers.getContractFactory("Union");
  const union = await Union.deploy(tgpassport.address);

  await union.deployed();
  console.log("Union deployed to:", union.address);

  // Checking bytes32 moderator
  const moderator_identifier = await union.connect(owner)
  .getModeratorIdentifier();
    console.log("bytes32 moderator:", moderator_identifier);






  // Checking roles set
  //const flag1 = await union.connect(owner)
  //.hasRole(moderator_identifier,owner.address);
  //console.log("owner account have moderator role:", flag1);
  const flag2 = await union.connect(owner)
  .hasRole(moderator_identifier,murs_account);
  console.log("Murs account have moderator role:", flag2);
  


  //const ERC20Sample = await ethers.getContractFactory("ERC20Sample");
  //let initialSupply = toWei(100);
  //const erc20sample = await ERC20Sample.deploy("Token", "TKN", initialSupply);
  //await erc20sample.deployed();
  //console.log("Sample erc20 deployed to: ", erc20sample.address);

  let erc20SampleDeployed = ethers.utils.getAddress("0xFbC45497848cc7438c528015271d73B9d8712385");
  console.log("Sample ERC20 deployed to: ", erc20SampleDeployed )

  //const ERC20Votes = await ethers.getContractFactory("ERC20VotesSample");
  //const erc20VotesSample = await ERC20Votes.deploy(initialSupply);
  //await erc20VotesSample.deployed();
  //console.log("Sample Votes ERC20 deployed to:", erc20VotesSample.address);

  const dao_test_address = ethers.utils.getAddress("0x18f060e4E6A7ff6f432d45629085AeF5E6Cc5081");
  console.log("test dao address:",dao_test_address);
  //const example_address = erc20sample.address;




  // retriving passport fee:
  const passportFee = await tgpassport.connect(owner)
  .GetPassportFee();
  console.log("passport fee:", passportFee);


  // for test data:
  // passports:
  const applyForPass = await tgpassport.connect(owner)
  .ApplyForPassport(1234,"test_username", { value : passportFee});
  const receipt = await applyForPass.wait();
  //console.log("receipt for applying:", receipt);
  const getPassportWallet = await tgpassport.GetWalletByNickName("test_username");
  console.log("applierd user wallet :", getPassportWallet);

  const approveForPass = await tgpassport.connect(owner)
  .ApprovePassport(getPassportWallet);
  const receipt2 = await approveForPass.wait();
  //console.log("receipt for approving:", receipt2);
  const getPassportWallet2 = await tgpassport.GetWalletByNickName("test_username");
  console.log("approved user wallet :", getPassportWallet2);

 




  // daos:
  const applyForUn = await union.connect(owner)
  .ApplyForUnion(1234,12345,dao_test_address,0,erc20SampleDeployed,"test_dao_username", { value : passportFee});
  const receipt3 = await applyForUn.wait();
 // console.log("receipt for applying dao:", receipt3);
  const getDaoAddress = await union.getDaoAddressbyChatId(12345);
  console.log("apllied dao address:", getDaoAddress);

  const approveForUn = await union.connect(owner)
  .ApproveJoin(getDaoAddress);
  const receipt4 = await approveForUn.wait();
 // console.log("receipt for approving join union: ", receipt4);
  const getDaoAddress2 = await union.getDaoAddressbyChatId(12345);
  console.log("approved dao address:", getDaoAddress2);


}


// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
