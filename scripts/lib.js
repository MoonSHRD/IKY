const { ethers } = require("hardhat");

function toWei(number) {
  return ethers.utils.parseUnits(number.toString(), 18);
}

function fromWei(number) {
  return ethers.utils.formatUnits(number.toString(), 18);
}

module.exports = {
  toWei,
  fromWei,
};
