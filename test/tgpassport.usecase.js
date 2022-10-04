const { expect } = require("chai");
const { ethers, waffle } = require("hardhat");
const { toWei, fromWei } = require("./lib.js");

describe("TGPassport", function () {
  let TGPassport, tgpassport, provider;
  let owner, user1, user2;
  let tgid1;
  let tgUserName1, tgUserName2;

  const passportFee = toWei(10);

  before(async () => {
    [owner, user1, user2] = await ethers.getSigners();
    tgid1 = 666;
    tgUserName1 = "username1";
    tgUserName2 = "username2";
    provider = waffle.provider;
    TGPassport = await ethers.getContractFactory("TGPassport");
    tgpassport = await TGPassport.deploy();
  });

  it("should NOT set passport fee", async () => {
    await expect(
      tgpassport.connect(user1).SetPassportFee(passportFee)
    ).to.be.revertedWith("Ownable: caller is not the owner");
  });

  it("should set passport fee and get it", async () => {
    await tgpassport.connect(owner).SetPassportFee(passportFee);
    const getPassportFee = await tgpassport.GetPassportFee();
    expect(getPassportFee).to.equal(passportFee);
  });

  it("should NOT apply for passport with zero value", async () => {
    await expect(
      tgpassport.connect(user1).ApplyForPassport(tgid1, tgUserName1)
    ).to.be.revertedWith("Passport fee is not paid");
  });

  it("should apply for passport", async () => {
    const userBalanceBeforeFee = await provider.getBalance(user1.address);
    const applyForPass = await tgpassport
      .connect(user1)
      .ApplyForPassport(tgid1, tgUserName1, { value: passportFee });

    const receipt = await applyForPass.wait();
    expect(tgid1).to.equal(receipt.events[0].args[0]);
    expect(receipt.events[0].event).to.equal("passportApplied");

    const userBalanceAfterFee = await provider.getBalance(user1.address);
    const ubbf = fromWei(userBalanceBeforeFee);
    const ubaf = fromWei(userBalanceAfterFee);
    expect(parseInt(ubbf)).to.be.greaterThan(parseInt(ubaf));
  });

  it("should get wallet by tg id", async () => {
    const getPassportWallet = await tgpassport.GetPassportWalletByID(tgid1);
    expect(user1.address).to.equal(getPassportWallet);
  });

  it("should get wallet by tg nickname", async () => {
    const getPassportWallet = await tgpassport.GetWalletByNickName(tgUserName1);
    expect(user1.address).to.equal(getPassportWallet);
  });

  it("should get passport by address", async () => {
    const getPassportWallet = await tgpassport.GetPassportByAddress(
      user1.address
    );
    expect(getPassportWallet.userAddress).to.equal(user1.address);
    expect(getPassportWallet.tgId).to.equal(tgid1);
    expect(getPassportWallet.userName).to.equal(tgUserName1);
  });

  it("should get passport by tg nickname", async () => {
    const getPassportWallet = await tgpassport.GetPassportByNickName(
      tgUserName1
    );
    expect(getPassportWallet.userAddress).to.equal(user1.address);
    expect(getPassportWallet.tgId).to.equal(tgid1);
    expect(getPassportWallet.userName).to.equal(tgUserName1);
  });

  it("should NOT approve passport", async () => {
    await expect(
      tgpassport.connect(user1).ApprovePassport(user1.address)
    ).to.be.revertedWith("Ownable: caller is not the owner");
  });

  it("should approve passport by admin", async () => {
    const approvePassport = await tgpassport
      .connect(owner)
      .ApprovePassport(user1.address);
    const receipt = await approvePassport.wait();
    expect(tgid1).to.equal(receipt.events[0].args.applyerTg);
    expect(user1.address).to.equal(receipt.events[0].args.wallet_address);
    expect(owner.address).to.equal(receipt.events[0].args.issuer);
    expect(receipt.events[0].event).to.equal("passportApproved");
  });

  it("should NOT update user name", async () => {
    await expect(
      tgpassport.connect(user2).UpdateUserName(tgUserName2)
    ).to.be.revertedWith("you don't now own this username");
  });

  it("should update user name", async () => {
    await tgpassport.connect(user1).UpdateUserName(tgUserName2);
    const updatedPassport = await tgpassport.passports(user1.address);
    expect(updatedPassport.userName).to.equal(tgUserName2);
  });
});
