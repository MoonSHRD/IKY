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

export default function DeclineDAOTG(props:Props){
  const addressContract = props.addressContract
  const currentAccount = props.currentAccount
  const [declined_dao_wallet, setDeclinedDaoWallet] = useState<string>("")

  const { query } = useRouter();


  // function for bot, for approving passport. make it here for tests
  async function declineDao(event:React.FormEvent) { 
    event.preventDefault()
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const Union:Contract = new ethers.Contract(addressContract, abi, signer)

    Union.DeclineJoin(declined_dao_wallet)
     .then((tr: TransactionResponse) => {
        console.log(`TransactionResponse TX hash: ${tr.hash}`)
        tr.wait().then((receipt:TransactionReceipt) => {console.log("decline receipt", receipt)})
        })
       .catch((e:Error) => console.log(e))
  }

  const handleChange = (value:string) => setDeclinedDaoWallet(value)

  return (
    <form onSubmit={declineDao}>
    <FormControl>
      <FormLabel htmlFor='WADDR'>Wallet Address: </FormLabel>
      <Input id="declined_user_wallet" type="text" required  onChange={(e) => setDeclinedDaoWallet(e.target.value)} value={query.dao_wallet} my={3}/>
      <Button type="submit" isDisabled={!currentAccount}>Decline DAO</Button>
    </FormControl>
    </form>
  )
}