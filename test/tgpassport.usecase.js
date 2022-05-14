const {
    expect
} = require("chai");
const {
    ethers, waffle
} = require("hardhat");
const {
    toWei, fromWei
} = require("./lib.js");

describe("TGPassport", function() {
    let TGPassport, tgpassport, provider;
    let owner, user1;
    let tgid1, tgid2; 

    const passportFee = toWei(10);

    before(async () => {
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

    it("should apply for passport", async () => {
        let userBalanceBeforeFee = await provider.getBalance(user1.address);
        const applyForPass = await tgpassport.connect(user1).applyForPassport(tgid1, {value: passportFee});

        const receipt = await applyForPass.wait();
        expect(tgid1).to.equal(receipt.events[0].args[0]);
        expect(receipt.events[0].event).to.equal("passportApplied");

        let userBalanceAfterFee = await provider.getBalance(user1.address);
        let ubbf = fromWei(userBalanceBeforeFee);
        let ubaf = fromWei(userBalanceAfterFee);
        expect(parseInt(ubbf)).to.be.greaterThan(parseInt(ubaf));
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