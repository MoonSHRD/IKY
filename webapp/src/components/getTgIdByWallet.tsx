import React, { useState, useEffect } from 'react'
import {Button, Input , NumberInput,  NumberInputField,  FormControl,  FormLabel, Text } from '@chakra-ui/react'
import {ethers} from 'ethers'
import {abi} from '../../../artifacts/contracts/TGPassport.sol/TGPassport.json'
import { Contract } from "ethers"


interface Props {
    addressContract: string,
    currentAccount: string | undefined
}

declare let window: any;


export default function GetWalletByTelegramNickNameTG(props:Props){
  const addressContract = props.addressContract
  const currentAccount = props.currentAccount
 // var [user_name, setUserName] = useState<string>("")
  var [tgid, setTgId] = useState<string>("")
  var [user_wallet, setUserWallet] = useState<string>("")

  useEffect(() => {
  const queryParams = new URLSearchParams(location.search);
  var _wallet = queryParams.get('user_wallet');
  
  
  setUserWallet(_wallet);
  
  }, []);
  


   async function getTgIdByWalletAddress(event:React.FormEvent) {
    event.preventDefault()
    if(!window.ethereum) return    
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const TGPassport:Contract = new ethers.Contract(addressContract, abi, signer)
    TGPassport.GetTgIdByAddress(user_wallet)
     .then((result:string) => {
        console.log(result)
        setTgId(result)
     })    
   }

  return (
    <form onSubmit={getTgIdByWalletAddress}>
    <FormControl>
      <FormLabel htmlFor='TGID'>Input eth wallet to get tgid </FormLabel>
      <Input id="tg_name" type="text" required  onChange={(e) => setUserWallet(e.target.value)} value={user_wallet} my={3}/>
      <Button type="submit" isDisabled={!currentAccount}>Get tgid associated with this wallet</Button>
    </FormControl>
    <div>
        <Text><b>TGID associated with this wallet</b>: {tgid}</Text>
    </div>
    </form>


  )
}