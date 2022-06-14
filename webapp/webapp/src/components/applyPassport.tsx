import React, { useState } from 'react'
import {Button, Input , NumberInput,  NumberInputField,  FormControl,  FormLabel } from '@chakra-ui/react'
import {ethers} from 'ethers'
import {parseEther } from 'ethers/lib/utils'
import {abi} from '../../../../artifacts/contracts/TGPassport.sol/TGPassport.json'
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
 // const [amount,setAmount]=useState<string>('100')
 // const [toAddress, setToAddress]=useState<string>("")
  const [user_id, setUserId] = useState<string>("")


  async function transfer(event:React.FormEvent) {
    event.preventDefault()
    if(!window.ethereum) return    
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const TGPassport:Contract = new ethers.Contract(addressContract, abi, signer)

    TGPassport.applyForPassport(user_id)
     .then((tr: TransactionResponse) => {
        console.log(`TransactionResponse TX hash: ${tr.hash}`)
        tr.wait().then((receipt:TransactionReceipt) => {console.log("transfer receipt", receipt)})
        })
         .catch((e:Error) => console.log(e))
     }

     


    /*
    erc20.transfer(toAddress,parseEther(amount))
      .then((tr: TransactionResponse) => {
        console.log(`TransactionResponse TX hash: ${tr.hash}`)
        tr.wait().then((receipt:TransactionReceipt)=>{console.log("transfer receipt",receipt)})
      })
      .catch((e:Error)=>console.log(e))
      */
  

  const handleChange = (value:string) => setUserId(value)

  return (
    <form onSubmit={transfer}>
    <FormControl>
      <FormLabel htmlFor='TGID'>User Telegram Id (not nickname!): </FormLabel>
      <Input id="tgid" type="text" required  onChange={(e) => setUserId(e.target.value)} my={3}/>
      <Button type="submit" isDisabled={!currentAccount}>Apply for Passport</Button>
    </FormControl>
    </form>
  )
}