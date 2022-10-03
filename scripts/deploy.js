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


  // We get the contract to deploy
  const TGPassport = await hre.ethers.getContractFactory("TGPassport");
  const tgpassport = await TGPassport.deploy();

  await tgpassport.deployed();
  console.log("TGPassport deployed to:", tgpassport.address);

  const Union = await hre.ethers.getContractFactory("Union");
  const union = await Union.deploy(tgpassport.address);

  await union.deployed();
  console.log("Union deployed to:", union.address);

  const ERC20Sample = await ethers.getContractFactory("ERC20Sample");
  const erc20sample = await ERC20Sample.deploy("Token", "TKN");
  await erc20sample.deployed();
  console.log("Sample erc20 deployed to: ", erc20sample.address);

  const example_address = erc20sample.address;


  let owner;
  owner = await hre.ethers.getSigner();
  console.log("owner address:", owner.address);

  // retriving passport fee:
  const passportFee = await tgpassport.connect(owner)
  .GetPassportFee();
  console.log("passport fee:", passportFee);


  // for test data:
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

  const applyForUn = await union.connect(owner)
  .ApplyForUnion(1234,12345,example_address,0,example_address,"test_dao_username", { value : passportFee});
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
