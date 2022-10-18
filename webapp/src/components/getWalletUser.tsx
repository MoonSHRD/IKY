import React, { useState, useEffect } from 'react'
import {Button, Input , NumberInput,  NumberInputField,  FormControl,  FormLabel, Text } from '@chakra-ui/react'
import {ethers} from 'ethers'
//import {parseEther } from 'ethers/lib/utils'
import {abi} from '../../../artifacts/contracts/TGPassport.sol/TGPassport.json'
import { Contract } from "ethers"
//import { TransactionResponse,TransactionReceipt } from "@ethersproject/abstract-provider"


interface Props {
    addressContract: string,
    currentAccount: string | undefined
}

declare let window: any;


export default function GetWalletByTelegramNickNameTG(props:Props){
  const addressContract = props.addressContract
  const currentAccount = props.currentAccount
  var [user_name, setUserName] = useState<string>("")
  var [user_wallet, setUserWallet] = useState<string>("")

  useEffect(() => {
  const queryParams = new URLSearchParams(location.search);
  var name = queryParams.get('user_tg_name');
  
  
  setUserName(name);
  
  }, []);
  



   async function getWalletByUsername(event:React.FormEvent) {
    event.preventDefault()
    if(!window.ethereum) return    
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const TGPassport:Contract = new ethers.Contract(addressContract, abi, signer)
    TGPassport.GetWalletByNickName(user_name)
     .then((result:string) => {
        console.log(result)
        setUserWallet(result)
     })
   }

  //const handleChange = (value:string) => setUserId(value)
  //http://localhost:3000?user_tg_id=1337&user_tg_name=Alice
  return (
    <form onSubmit={getWalletByUsername}>
    <FormControl>
      <FormLabel htmlFor='TGID'>Input telegram nickname to get it's eth wallet </FormLabel>
      <Input id="tg_name" type="text" required  onChange={(e) => setUserName(e.target.value)} value={user_name} my={3}/>
      <Button type="submit" isDisabled={!currentAccount}>Get wallet address</Button>
    </FormControl>
    <div>
        <Text><b>Ethereum address associated with this telegram nickname</b>: {user_wallet}</Text>
    </div>
    </form>


  )
}