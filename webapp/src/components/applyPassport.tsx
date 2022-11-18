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
  var [user_id, setUserId] = useState(0)
  var [user_name, setUserName] = useState<string>("")

  useEffect(() => {
    const queryParams = new URLSearchParams(location.search);

  var id = queryParams.get('user_tg_id');   // get id as string from query
  let int_id : number = +id;                // similar to parseInt()
  var name = queryParams.get('user_tg_name');
  
  

  setUserId(int_id);
  setUserName(name);
  
  }, []);
  

  async function applyPersonalPassport(event:React.FormEvent) {
    event.preventDefault()
    if(!window.ethereum) return    
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const TGPassport:Contract = new ethers.Contract(addressContract, abi, signer)
   // let passport_fee_wei = ethers.utils.formatUnits(1000,"wei");
    //let passport_fee_custom_gwei = ethers.utils.formatUnits(2000000,"gwei"); // 1 gwei = 1'000'000'000 wei, 2m gwei = 0,002 (estimateGas on approval = 0.02, so we need to take that fee for gas)
    //let passport_fee_wei = ethers.utils.formatUnits(passport_fee_custom_gwei,"wei");
    //let passport_fee_wei_hardcode = ethers.utils.formatUnits(2000000000000000,"wei");
    TGPassport.ApplyForPassport(user_id,user_name,{value:ethers.utils.formatUnits(2000000000000000,"wei")})
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
      <Input id="tgid" type="number" required  onChange={(e) => setUserId(parseInt(e.target.value))} value={user_id} my={3}/>
     
      <Input id="tg_name" type="text" required  onChange={(e) => setUserName(e.target.value)} value={user_name} my={3}/>
      <Button type="submit" isDisabled={!currentAccount}>Apply for Passport</Button>
    </FormControl>
    </form>
  )
}