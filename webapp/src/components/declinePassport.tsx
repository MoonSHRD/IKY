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

export default function DeclinePassportTG(props:Props){
  const addressContract = props.addressContract
  const currentAccount = props.currentAccount
  const [declined_user_wallet, setDeclinedUserWallet] = useState<string>("")

  const { query } = useRouter();


  // function for bot, for approving passport. make it here for tests
  async function declinePassport(event:React.FormEvent) { 
    event.preventDefault()
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const TGPassport:Contract = new ethers.Contract(addressContract, abi, signer)

    TGPassport.DeclinePassport(declined_user_wallet)
     .then((tr: TransactionResponse) => {
        console.log(`TransactionResponse TX hash: ${tr.hash}`)
        tr.wait().then((receipt:TransactionReceipt) => {console.log("decline receipt", receipt)})
        })
       .catch((e:Error) => console.log(e))
  }

  const handleChange = (value:string) => setDeclinedUserWallet(value)

  return (
    <form onSubmit={declinePassport}>
    <FormControl>
      <FormLabel htmlFor='WADDR'>Wallet Address: </FormLabel>
      <Input id="declined_user_wallet" type="text" required  onChange={(e) => setDeclinedUserWallet(e.target.value)} value={query.user_wallet} my={3}/>
      <Button type="submit" isDisabled={!currentAccount}>Decline Passport</Button>
    </FormControl>
    </form>
  )
}