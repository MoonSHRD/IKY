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

export default function DeletePassportTG(props:Props){
  const addressContract = props.addressContract
  const currentAccount = props.currentAccount
  const [delete_user_wallet, setDeleteUserWallet] = useState<string>("")

  const { query } = useRouter();


  async function deletePassport(event:React.FormEvent) { 
    event.preventDefault()
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const TGPassport:Contract = new ethers.Contract(addressContract, abi, signer)

    TGPassport.DeletePassport(delete_user_wallet)
     .then((tr: TransactionResponse) => {
        console.log(`TransactionResponse TX hash: ${tr.hash}`)
        tr.wait().then((receipt:TransactionReceipt) => {console.log("DELETE receipt", receipt)})
        })
       .catch((e:Error) => console.log(e))
  }

  const handleChange = (value:string) => setDeleteUserWallet(value)

  return (
    <form onSubmit={deletePassport}>
    <FormControl>
      <FormLabel htmlFor='WADDR'>Wallet Address: </FormLabel>
      <Input id="delete_user_wallet" type="text" required  onChange={(e) => setDeleteUserWallet(e.target.value)} value={query.user_wallet} my={3}/>
      <Button type="submit" isDisabled={!currentAccount}>Delete Passport</Button>
    </FormControl>
    </form>
  )
}