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

export default function ApproveDaoTG(props:Props){
  const addressContract = props.addressContract
  const currentAccount = props.currentAccount
  const [approved_dao_wallet, setApprovedDaoWallet] = useState<string>("")

  const { query } = useRouter();


  // function for bot, for approving passport. make it here for tests
  async function approveDao(event:React.FormEvent) { 
    event.preventDefault()
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const Union:Contract = new ethers.Contract(addressContract, abi, signer)

    Union.ApproveJoin(approved_dao_wallet)
     .then((tr: TransactionResponse) => {
        console.log(`TransactionResponse TX hash: ${tr.hash}`)
        tr.wait().then((receipt:TransactionReceipt) => {console.log("approving receipt", receipt)})
        })
       .catch((e:Error) => console.log(e))
  }

  const handleChange = (value:string) => setApprovedDaoWallet(value)

  return (
    <form onSubmit={approveDao}>
    <FormControl>
      <FormLabel htmlFor='WADDR'>Wallet Address: </FormLabel>
      <Input id="approved_user_wallet" type="text" required  onChange={(e) => setApprovedDaoWallet(e.target.value)} value={query.approved_dao_wallet} my={3}/>
      <Button type="submit" isDisabled={!currentAccount}>Approve DAO</Button>
    </FormControl>
    </form>
  )
}