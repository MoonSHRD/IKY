/* eslint-disable prettier/prettier */
/* eslint-disable prefer-const */
/* eslint-disable node/no-extraneous-require */
/* eslint-disable no-sequences */
/* eslint-disable prettier/prettier */
/* eslint-disable no-unused-vars */
const { expect } = require("chai");
const { ethers, waffle} = require("hardhat");
const { toWei, fromWei} = require("./lib.js");

describe("TGPassport", function() {
    let TGPassport, tgpassport, provider;
    let owner, user1;
    let tgid1, tgid2; 

    const passportFee = toWei(10);

    beforeEach(async () => {
        [owner, user1] = await ethers.getSigners();
        tgid1 = "@vouter1";
        tgid2 = "@vouter2";
        provider = waffle.provider;
        TGPassport = await ethers.getContractFactory("TGPassport");
        tgpassport = await TGPassport.deploy();
    });

    it("should NOT set passport fee", async () => {
        await expect(
            tgpassport
              .connect(user1)
              .setPassportFee(passportFee)
          ).to.be.revertedWith("Ownable: caller is not the owner");
    });

    it("should set passport fee and get it", async () => {
        await tgpassport.connect(owner).setPassportFee(passportFee);
        const getPassportFee = await tgpassport.getPassportFee();
        expect(getPassportFee).to.equal(passportFee);
    });

    it("should NOT apply for passport with zero value", async () => {
        await expect(
            tgpassport
              .connect(user1)
              .applyForPassport(tgid1)
          ).to.be.revertedWith("Passport fee is not paid");
    });

    describe("ApplyForPassport", function() {

        it("Should be reverted with empty Passport Fee", async () => {
            await expect (tgpassport.connect(user1).ApplyForPassport(tgid1, "username", {value: 0}))
                .to.be.reverted;
        });

        it("Should be passed with correct Passport Fee", async () => {
            await expect (tgpassport.connect(user1).ApplyForPassport(tgid1, "username", {value: await tgpassport.GetPassportFee()}))
                .to.be.ok;
        });

        it("Event passportApplied should be emitted", async () => {
            await expect (await tgpassport.connect(user1).ApplyForPassport(tgid1, "username", {value: await tgpassport.GetPassportFee()}))
                    .to.emit(tgpassport, "passportApplied")
                    .withArgs(tgid1, user1.address);
        });

        it("Should be increased owner balance for Passport Fee value", async () => {
    
            const ownerBalanceBefore = await provider.getBalance(owner.address);

            await tgpassport.connect(user1).ApplyForPassport(tgid1, "username", {value: await tgpassport.GetPassportFee()});

            const ownerBalanceAfter = await provider.getBalance(owner.address);
        
            await expect(ownerBalanceAfter).to.equal(ownerBalanceBefore.add(ethers.BigNumber.from(await tgpassport.GetPassportFee())));
        });
    });
        
    describe("ApprovePassport", function() {

        it("Should be reverted for non-owner caller", async () => {

            await tgpassport.connect(user1).ApplyForPassport(tgid1, "username", {value: await tgpassport.GetPassportFee()});

            await expect (tgpassport.connect(user1).ApprovePassport(user1.address))
                .to.be.reverted;
        });

        it("Should be passed for owner caller", async () => {

            await tgpassport.connect(user1).ApplyForPassport(tgid1, "username", {value: await tgpassport.GetPassportFee()});

            await expect (tgpassport.connect(owner).ApprovePassport(user1.address))
                .to.emit(tgpassport, "passportApproved")
                .withArgs(tgid1, user1.address, owner.address);
        });
    });

    describe("DeclinePassport", function() {

        it("Should be reverted for non-owner caller", async () => {

            await tgpassport.connect(user1).ApplyForPassport(tgid1, "username", {value: await tgpassport.GetPassportFee()});

            await expect (tgpassport.connect(user1).DeclinePassport(user1.address))
                .to.be.reverted;
        });

        it("Should be passed for owner caller", async () => {

            await tgpassport.connect(user1).ApplyForPassport(tgid1, "username", {value: await tgpassport.GetPassportFee()});

            await expect (tgpassport.connect(owner).DeclinePassport(user1.address))
                .to.emit(tgpassport, "passportDenied")
                .withArgs(tgid1, user1.address);
        });
    });

    it("should get wallet by tg id", async () => {
        const getPassportWallet = await tgpassport.getPassportWallet(tgid1);
        expect(user1.address).to.equal(getPassportWallet);
    });

    it("should NOT approve passport", async () => {
        await expect(
            tgpassport
              .connect(user1)
              .approvePassport(user1.address)
          ).to.be.revertedWith("Ownable: caller is not the owner");
    });

    it("should approve passport by admin", async () => {
        const approvePassport = await tgpassport.connect(owner).approvePassport(user1.address);
        let receipt = await approvePassport.wait();
        expect(tgid1).to.equal(receipt.events[0].args.applyerTg);
        expect(user1.address).to.equal(receipt.events[0].args.wallet_address);
        expect(owner.address).to.equal(receipt.events[0].args.issuer);
        expect(receipt.events[0].event).to.equal("passportApproved");
    });
});