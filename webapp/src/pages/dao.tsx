import type { NextPage } from 'next'
import Head from 'next/head'
import NextLink from "next/link"
import { VStack, Heading, Box, LinkOverlay, LinkBox} from "@chakra-ui/layout"
import { Text, Button } from '@chakra-ui/react'
import { useState, useEffect} from 'react'
import {ethers} from "ethers"
//import ReadERC20 from 'components/read_contract'
import ReadPassportContract from 'components/read_contract'
//import ApplyPassportTG from 'components/applyPassport'
import ApplyDaoTG from 'components/applyDao'
//import ApprovePassportTG from 'components/approvePassport'
import ApproveDaoTG from 'components/approveDao'
//import DeclinePassportTG from 'components/declinePassport'
import DeclineDAOTG from 'components/declineDao'

declare let window:any

const Home: NextPage = () => {
  const [balance, setBalance] = useState<string | undefined>()
  const [currentAccount, setCurrentAccount] = useState<string | undefined>()
  const [chainId, setChainId] = useState<number | undefined>()
  const [chainname, setChainName] = useState<string | undefined>()

  useEffect(() => {
    if(!currentAccount || !ethers.utils.isAddress(currentAccount)) return
    //client side code
    if(!window.ethereum) return
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    provider.getBalance(currentAccount).then((result)=>{
      setBalance(ethers.utils.formatEther(result))
    })
    provider.getNetwork().then((result)=>{
      setChainId(result.chainId)
      setChainName(result.name)
    })

  },[currentAccount])

  const onClickConnect = () => {
    //client side code
    if(!window.ethereum) {
      console.log("please install MetaMask")
      return
    }
    /*
    //change from window.ethereum.enable() which is deprecated
    //see docs: https://docs.metamask.io/guide/ethereum-provider.html#legacy-methods
    window.ethereum.request({ method: 'eth_requestAccounts' })
    .then((accounts:any)=>{
      if(accounts.length>0) setCurrentAccount(accounts[0])
    })
    .catch('error',console.error)
    */

    //we can do it using ethers.js
    const provider = new ethers.providers.Web3Provider(window.ethereum)

    // MetaMask requires requesting permission to connect users accounts
    provider.send("eth_requestAccounts", [])
    .then((accounts)=>{
      if(accounts.length>0) setCurrentAccount(accounts[0])
    })
    .catch((e)=>console.log(e))
  }

  const onClickDisconnect = () => {
    console.log("onClickDisConnect")
    setBalance(undefined)
    setCurrentAccount(undefined)
  }

  return (
    <>
      <Head>
        <title>My DAPP</title>
      </Head>

      <Heading as="h3"  my={4}>Explore Web3</Heading>          
      <VStack>
        <Box w='100%' my={4}>
        {currentAccount  
          ? <Button type="button" w='100%' onClick={onClickDisconnect}>
                Account:{currentAccount}
            </Button>
          : <Button type="button" w='100%' onClick={onClickConnect}>
                  Connect MetaMask
              </Button>
        }
        </Box>
        {currentAccount  
          ?<Box  mb={0} p={4} w='100%' borderWidth="1px" borderRadius="lg">
          <Heading my={4}  fontSize='xl'>Account info</Heading>
          <Text>ETH Balance of current account: {balance}</Text>
          <Text>Chain Info: ChainId {chainId} name {chainname}</Text>
        </Box>
        :<></>
        }

        <Box  mb={0} p={4} w='100%' borderWidth="1px" borderRadius="lg">
          <Heading my={4}  fontSize='xl'>Attach CHAT from telegram to collective wallet</Heading>
          <ApplyDaoTG 
            addressContract='0x9E296573FD351881631EB7cdDa709AE9cd27D5CD'
            currentAccount={currentAccount}
          />
        </Box>
        <Box  mb={0} p={4} w='100%' borderWidth="1px" borderRadius="lg">
          <Heading my={4}  fontSize='xl'>Approve user intent for registration personal wallet address</Heading>
          <ApproveDaoTG 
            addressContract='0x9E296573FD351881631EB7cdDa709AE9cd27D5CD'
            currentAccount={currentAccount}
          />
        </Box>
        <Box  mb={0} p={4} w='100%' borderWidth="1px" borderRadius="lg">
          <Heading my={4}  fontSize='xl'>Decline user intent for registration personal wallet address</Heading>
          <DeclineDAOTG 
            addressContract='0x9E296573FD351881631EB7cdDa709AE9cd27D5CD'
            currentAccount={currentAccount}
          />
        </Box>
        

...
      </VStack>
    </>
  )
}

export default Home