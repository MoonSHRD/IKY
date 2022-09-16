import React, { useState } from 'react'
import {Button, Input , NumberInput,  NumberInputField, Select,  FormControl,  FormLabel, Radio, RadioGroup, Stack, Center } from '@chakra-ui/react'
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

  const [votingType,setVotingType] = useState<string>("")
  const [votingTokenContract,setVotingTokenContract] = useState<string>("")

  const [user_name, setUserName] = useState<string>("")

  const { query } = useRouter();

  async function applyDao(event:React.FormEvent) {
    event.preventDefault()
    if(!window.ethereum) return    
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const Union:Contract = new ethers.Contract(addressContract, abi, signer)

    Union.ApplyForUnion(user_id,chat_id,votingType,votingTokenContract,{value:ethers.utils.formatUnits(1000,"wei")})
     .then((tr: TransactionResponse) => {
        console.log(`TransactionResponse TX hashonChange={(e) => setUserId(e.target.value)} value={query.user_id} my={3} : ${tr.hash}`)
        tr.wait().then((receipt:TransactionReceipt) => {console.log("applying receipt", receipt)})
        })
         .catch((e:Error) => console.log(e))
     }

  
  //const handleChange = (value:string) => setUserId(value)

  return (
    <form onSubmit={applyDao}>
    <FormControl>
      <FormLabel htmlFor='TGID'>Enter data: </FormLabel>
      <Input id="user_id" type="text" placeholder="Your Telegram ID" required  onChange={(e) => setUserId(e.target.value)} value={query.user_id} my={3}/>
      <Input id="chat_id" type="text" placeholder="Chat's Telegram ID"required  onChange={(e) => setChatId(e.target.value)} value={query.chat_id} my={3}/>
      <Select id="votingtype" placeholder="Select voting token's type:" onChange={(e) => setVotingType(e.target.value)} value={query.votingtype} my={3}>
      <option value='0'>ERC20</option>
      <option value='1'>ERC20Snapshot</option>
      <option value='2'>ERC721</option>
      </Select>
      <Input id="votingtokencontract" placeholder="Voting token's contract address" type="text" required  onChange={(e) => setVotingTokenContract(e.target.value)} value={query.votingtokencontract} my={3}/>
      <Button type="submit" isDisabled={!currentAccount}>Apply DAO for Union</Button>
    </FormControl>
    </form>

    
  )
}