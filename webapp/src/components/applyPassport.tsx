import React, { useState, useEffect } from 'react'
import {Button, Input , NumberInput,  NumberInputField,  FormControl,  FormLabel } from '@chakra-ui/react'
import {ethers} from 'ethers'
import {parseEther } from 'ethers/lib/utils'
import {abi} from '../../../artifacts/contracts/TGPassport.sol/TGPassport.json'
import { Contract } from "ethers"
import { TransactionResponse,TransactionReceipt } from "@ethersproject/abstract-provider"


interface Props {
    addressContract: string,
    currentAccount: string | undefined
}

declare let window: any;


export default function ApplyPassportTG(props:Props){
  const addressContract = props.addressContract
  const currentAccount = props.currentAccount
  const [user_id, setUserId] = useState<string>("")
  const [user_name, setUserName] = useState<string>("")

  useEffect(() => {
    const queryParams = new URLSearchParams(location.search);

  var id = queryParams.get('user_tg_id');
  var name = queryParams.get('user_tg_name');
  
  setUserId(id);
  setUserName(name);
  
  }, []);
  

  async function applyPersonalPassport(event:React.FormEvent) {
    event.preventDefault()
    if(!window.ethereum) return    
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const TGPassport:Contract = new ethers.Contract(addressContract, abi, signer)
    TGPassport.ApplyForPassport(user_id,user_name,{value:ethers.utils.formatUnits(1000,"wei")})
     .then((tr: TransactionResponse) => {
        console.log(`TransactionResponse TX hash: ${tr.hash}`)
        tr.wait().then((receipt:TransactionReceipt) => {console.log("applying receipt", receipt)})
        })
         .catch((e:Error) => console.log(e))
     }


  
  //const handleChange = (value:string) => setUserId(value)
  //http://localhost:3000?user_tg_id=1337&user_tg_name=Alice
  return (
    <form onSubmit={applyPersonalPassport}>
    <FormControl>
      <FormLabel htmlFor='TGID'>User Telegram Id (not nickname!): </FormLabel>
      <Input id="tgid" type="text" required  onChange={(e) => setUserId(e.target.value)} value={user_id} my={3}/>
     
      <Input id="tg_name" type="text" required  onChange={(e) => setUserName(e.target.value)} value={user_name} my={3}/>
      <Button type="submit" isDisabled={!currentAccount}>Apply for Passport</Button>
    </FormControl>
    </form>
  )
}