/* eslint-disable prettier/prettier */
/* eslint-disable prefer-const */
/* eslint-disable node/no-extraneous-require */
/* eslint-disable no-sequences */
/* eslint-disable prettier/prettier */
/* eslint-disable no-unused-vars */
const {inputToConfig} = require("@ethereum-waffle/compiler");
const {
    expect
} = require("chai");
const {
    ethers, waffle
} = require("hardhat");
const {
    toWei, fromWei
} = require("./lib.js");

describe("EPassport unit test", async () => {
    let EPassport, epassport, provider;
    let owner, user1;
    let passHash, link;

    before(async () => {
        [owner, user1] = await ethers.getSigners();
        provider = waffle.provider;
        passHash = "56c607f1dca246cd250ea9d76241e4fadd30936a1b0a199b9de130765c399dd8";
        link = "https://ipfs.io/ipfs/QmcsZbtDmbZ4dqErUpVkUzYNFUpCxSQQ32AzqmLwtCZT36?filename=photo_2021-09-27_02-05-09.jpg";
        EPassport = await ethers.getContractFactory("EPassport");
        epassport = await EPassport.deploy();
    });

    it("should apply epassport ", async () => {
        const apply = await epassport.connect(user1).applyForPassport(passHash, link);
        const receipt = await apply.wait();
        // console.log(receipt.events[0].args);
        expect(passHash).to.equal(receipt.events[0].args[0]);
        expect(user1.address).to.equal(receipt.events[0].args[1]);
        expect(receipt.events[0].event).to.equal("passportApplied");
    });

    it("should NOT approve epassport ", async () => {
        await expect(
            epassport
              .connect(user1)
              .approvePassport(passHash)
          ).to.be.revertedWith("Ownable: caller is not the owner");
    });

    it("should approve epassport by admin", async () => {
        const approve = await epassport.connect(owner).approvePassport(passHash);
        const receipt = await approve.wait();
        expect(passHash).to.equal(receipt.events[0].args.passportIdHash);
        expect(user1.address).to.equal(receipt.events[0].args.wallet_address);
        expect(owner.address).to.equal(receipt.events[0].args.issuer);
        expect(receipt.events[0].event).to.equal("passportApproved");
    });

    it("should get passport by ID hash", async () => {
        const getpassport = await epassport.getPassportByIdHash(passHash);
        expect(getpassport.ID_hash).to.equal(passHash);
        expect(getpassport.IPFS_link).to.equal(link);
        expect(getpassport.user_wallet).to.equal(user1.address);
        expect(getpassport.approved).to.equal(true);
    });

    it("should get passport by EID", async () => {
        const getpassport = await epassport.getPassportByEid(1);
        expect(getpassport.ID_hash).to.equal(passHash);
        expect(getpassport.IPFS_link).to.equal(link);
        expect(getpassport.user_wallet).to.equal(user1.address);
        expect(getpassport.approved).to.equal(true);
    });

    it("should NOT exile epassport ", async () => {
        await expect(
            epassport
              .connect(user1)
              .exile(passHash)
          ).to.be.revertedWith("Ownable: caller is not the owner");
    });

    it("should exile epassport by admin", async () => {
        const exile = await epassport.connect(owner).exile(passHash);
        const receipt = await exile.wait();
        expect(passHash).to.equal(receipt.events[0].args.passportIdHash);
        expect(user1.address).to.equal(receipt.events[0].args.wallet_address);
        expect(owner.address).to.equal(receipt.events[0].args.kyc_center);
        expect(receipt.events[0].event).to.equal("passportBanned");
    });

    it("should NOT be approved because banned", async () => {
        const getPassportById = await epassport.getPassportByEid(1);
        expect(getPassportById.approved).to.equal(false);

        const getPassportByHash = await epassport.getPassportByIdHash(passHash);
        expect(getPassportByHash.approved).to.equal(false);
    });
});