import React, { useState } from 'react'
import {Button, Input , NumberInput,  NumberInputField,  FormControl,  FormLabel } from '@chakra-ui/react'
import {ethers} from 'ethers'
import {parseEther } from 'ethers/lib/utils'
import {abi} from '../../../artifacts/contracts/Union.sol/Union.json'
import { Contract } from "ethers"
import { TransactionResponse,TransactionReceipt } from "@ethersproject/abstract-provider"
import {useRouter} from "next/router";

interface Props {
    addressContract: string,
    currentAccount: string | undefined
}

declare let window: any;

export default function ApplyDaoTG(props:Props){
  const addressContract = props.addressContract
  const currentAccount = props.currentAccount
  const [user_id, setUserId] = useState<string>("")
  const [chat_id,setChatId] = useState<string>("")
  const [user_name, setUserName] = useState<string>("")

  const { query } = useRouter();

  async function applyDao(event:React.FormEvent) {
    event.preventDefault()
    if(!window.ethereum) return    
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const Union:Contract = new ethers.Contract(addressContract, abi, signer)

    Union.ApplyForUnion(user_id,chat_id,{value:ethers.utils.formatUnits(1000,"wei")})
     .then((tr: TransactionResponse) => {
        console.log(`TransactionResponse TX hash: ${tr.hash}`)
        tr.wait().then((receipt:TransactionReceipt) => {console.log("applying receipt", receipt)})
        })
         .catch((e:Error) => console.log(e))
     }

  
  //const handleChange = (value:string) => setUserId(value)

  return (
    <form onSubmit={applyDao}>
    <FormControl>
      <FormLabel htmlFor='TGID'>Input chat id: </FormLabel>
      <Input id="tgid" type="text" required  onChange={(e) => setUserId(e.target.value)} value={query.user_id} my={3}/>
      <Input id="chatid" type="text" required  onChange={(e) => setChatId(e.target.value)} value={query.chat_id} my={3}/>
      <Button type="submit" isDisabled={!currentAccount}>Apply DAO for Union</Button>
    </FormControl>
    </form>
  )
}