// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const hre = require("hardhat");

async function main() {
  // Hardhat always runs the compile task when running scripts with its command
  // line interface.
  //
  // If this script is run directly using `node` you may want to call compile
  // manually to make sure everything is compiled
  // await hre.run('compile');

  // We get the contract to deploy
  const TGPassport = await hre.ethers.getContractFactory("TGPassport");
  const tgpassport = await TGPassport.deploy();

  await tgpassport.deployed();

  console.log("TGPassport deployed to:", tgpassport.address);

  const Union = await hre.ethers.getContractFactory("Union");
  const union = await Union.deploy(tgpassport.address);

  await union.deployed();

  console.log("Union deployed to:", union.address);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
