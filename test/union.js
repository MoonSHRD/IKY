const { expect } = require("chai");
const { ethers } = require("hardhat");
const { toWei, fromWei } = require("./lib.js");

describe("Union", function () {
  let Union,
    union,
    TGPassport,
    tgpassport,
    ERC20Sample,
    erc20sample,
    applyerTg,
    daoTg;
  let owner, user1, dao;

  let tgUserName1, tgUserName2;

  const passportFee = toWei(1);

  before(async () => {
    applyerTg = "@voter1";
    daoTg = "@gov";
    tgUserName1 = "username1";
    tgUserName2 = "username2";
    [owner, user1, dao] = await ethers.getSigners();
    ERC20Sample = await ethers.getContractFactory("ERC20Sample");
    erc20sample = await ERC20Sample.deploy("Token", "TKN");
    TGPassport = await ethers.getContractFactory("TGPassport");
    tgpassport = await TGPassport.deploy();
    Union = await ethers.getContractFactory("Union");
    union = await Union.deploy(tgpassport.address);
  });

  it("should mint some tokens", async () => {
    await erc20sample.connect(owner).mint(owner.address, toWei(100));
    const ownerBalance = await erc20sample.balanceOf(owner.address);
    expect(ownerBalance).to.equal(toWei(100));
  });

  it("should set passport fee and get it", async () => {
    const tgPassAddress = await union.tgpassport();
    await tgpassport.attach(tgPassAddress);
    await tgpassport.connect(owner).SetPassportFee(passportFee);
    const getPassportFee = await tgpassport.GetPassportFee();
    expect(getPassportFee).to.equal(passportFee);
  });

  it("should apply for passport and approve it", async () => {
    await tgpassport
      .connect(owner)
      .ApplyForPassport(applyerTg, tgUserName1, { value: passportFee });
    await tgpassport.connect(owner).ApprovePassport(user1.address);
  });

  it("should apply for union with erc20", async () => {
    const apply = await union
      .connect(owner)
      .ApplyForUnion(applyerTg, daoTg, dao.address, 0, erc20sample.address, {
        value: passportFee,
      });
    const receipt = await apply.wait();
    expect(receipt.events[0].event).to.equal("ApplicationForJoin");
    expect(daoTg).to.equal(receipt.events[0].args.chat_id);
    expect(applyerTg).to.equal(receipt.events[0].args.applier_id);
    expect(dao.address).to.equal(receipt.events[0].args.multy_wallet_address);
    expect(receipt.events[0].args.vote_type).to.equal(0);
    expect(erc20sample.address).to.equal(
      receipt.events[0].args.voting_token_address
    );
  });

  it("should NOT approve join", async () => {
    await expect(
      union.connect(user1).ApproveJoin(dao.address)
    ).to.be.revertedWith("Ownable: caller is not the owner");
  });

  it("should approve join", async () => {
    const approve = await union.connect(owner).ApproveJoin(dao.address);
    const receipt = await approve.wait();
    expect(receipt.events[0].event).to.equal("ApprovedJoin");
    expect(daoTg).to.equal(receipt.events[0].args.chat_id);
    expect(dao.address).to.equal(receipt.events[0].args.multy_wallet_address);
    expect(receipt.events[0].args.vote_type).to.equal(0);
    expect(erc20sample.address).to.equal(
      receipt.events[0].args.voting_token_address
    );
    const status = await union.daos(dao.address);
    expect(status.valid).to.equal(true);
  });

  it("should NOT approve join because already approved", async () => {
    await expect(
      union.connect(owner).ApproveJoin(dao.address)
    ).to.be.revertedWith("already has been approved OR didn't applied at all");
  });
});
