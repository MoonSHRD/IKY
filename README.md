# Know Your Customer / I Know You 
This is KYC/IKY solution to tether ethereum wallet of user to it's accounts at third platforms

Basically we want to tether telegram_id to ethereum wallet, which allow us to do gassless transactions at messaging platforms and protokols.

Contracts

TGPassport -- basically mapping contract which tether tg_id and wallet
Union -- extension for registring group chats tg_id to multisig wallet address 

# Deploy
1. create .env file in the root of project (.env.example as reference). Add your RPC end-point and your private key with 0x prefix
2. `npx hardhat run --network rinkeby scripts/deploy.js ` for Rinkeby testnet
3. don't forget to add deployed addresses to adresses.txt file and make change in `webapp/src/pages/index.tsx` if you need to run webapp

## Local deploy
To make deploy to local node run `npx hardhat node` in one terminal and run `npx hardhat run --network localhost scripts/deploy.js` in *other* terminal window
it will create test enviroment like an early bundle truffle+ganache


# Frontend is in webapp directory

# Golang artifacts for telegram bot

0. You should change imports in contracts from `@openzeppeline/contracts/...` to `../node_modules/@openzeppeline/contracts/...` it's required for proper generation of ABI


1. You can generate abi as is:
```
solc --abi --bin ./contracts/TGPassport.sol -o build ..=.. --overwrite --allow-paths *,/node_modules/,

solc --abi --bin ./contracts/Union.sol -o build ..=.. --overwrite --allow-paths *,/node_modules/,
solc --abi --bin ./contracts/erc20/ERC20Sample.sol -o build ..=.. --overwrite --allow-paths *,/node_modules/,
solc --abi --bin ./contracts/erc721/ERC721Sample.sol -o build ..=.. --overwrite --allow-paths *,/node_modules/,
solc --abi --bin ./contracts/erc20Votes/ERC20VotesSample.sol -o build ..=.. --overwrite --allow-paths *,/node_modules/, 


```

2. You can generate Golang artifacts as is:
```
abigen --abi="build/TGPassport.abi" --pkg=Passport --out="./go/TGPassport.go"
abigen --abi="build/Union.abi" --pkg=Union --out="./go/Union.go"
abigen --abi="build/ERC20Sample.abi" --pkg=TokenERC20 --out="./go/ERC20.go"
abigen --abi="build/ERC721Sample.abi" --pkg=TokenERC721 --out="./go/ERC721.go"
abigen --abi="build/ERC20VotesSample.abi" --pkg=TokenERC20Votes --out="./go/ERC20Votes.go"


```

# Advanced Sample Hardhat Project

This project demonstrates an advanced Hardhat use case, integrating other tools commonly used alongside Hardhat in the ecosystem.

The project comes with a sample contract, a test for that contract, a sample script that deploys that contract, and an example of a task implementation, which simply lists the available accounts. It also comes with a variety of other tools, preconfigured to work with the project code.

Try running some of the following tasks:

```shell
npx hardhat accounts
npx hardhat compile
npx hardhat clean
npx hardhat test
npx hardhat node
npx hardhat help
REPORT_GAS=true npx hardhat test
npx hardhat coverage
npx hardhat run scripts/deploy.js
node scripts/deploy.js
npx eslint '**/*.js'
npx eslint '**/*.js' --fix
npx prettier '**/*.{json,sol,md}' --check
npx prettier '**/*.{json,sol,md}' --write
npx solhint 'contracts/**/*.sol'
npx solhint 'contracts/**/*.sol' --fix
```


