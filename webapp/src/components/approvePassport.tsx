import React, { useState } from 'react'
import {Button, Input , NumberInput,  NumberInputField,  FormControl,  FormLabel } from '@chakra-ui/react'
import {ethers} from 'ethers'
import {parseEther } from 'ethers/lib/utils'
import {abi} from '../../../artifacts/contracts/TGPassport.sol/TGPassport.json'
import { Contract } from "ethers"
import { TransactionResponse,TransactionReceipt } from "@ethersproject/abstract-provider"
import {useRouter} from "next/router";

interface Props {
    addressContract: string,
    currentAccount: string | undefined
}

declare let window: any;

export default function ApplyPassportTG(props:Props){
  const addressContract = props.addressContract
  const currentAccount = props.currentAccount
  const [approved_user_wallet, setApprovedUserWallet] = useState<string>("")

  const { query } = useRouter();


  // function for bot, for approving passport. make it here for tests
  async function approvePassport(event:React.FormEvent) { 
    event.preventDefault()
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const TGPassport:Contract = new ethers.Contract(addressContract, abi, signer)

    TGPassport.ApprovePassport(approved_user_wallet)
     .then((tr: TransactionResponse) => {
        console.log(`TransactionResponse TX hash: ${tr.hash}`)
        tr.wait().then((receipt:TransactionReceipt) => {console.log("approving receipt", receipt)})
        })
       .catch((e:Error) => console.log(e))
  }

  const handleChange = (value:string) => setApprovedUserWallet(value)

  return (
    <form onSubmit={approvePassport}>
    <FormControl>
      <FormLabel htmlFor='WADDR'>Wallet Address: </FormLabel>
      <Input id="approved_user_wallet" type="text" required  onChange={(e) => setApprovedUserWallet(e.target.value)} value={query.approved_user_wallet} my={3}/>
      <Button type="submit" isDisabled={!currentAccount}>Approve Passport</Button>
    </FormControl>
    </form>
  )
}