import React, { useState , useEffect} from 'react'
import {Button, Input , NumberInput,  NumberInputField, Select,  FormControl,  FormLabel, Radio, RadioGroup, Stack, Center } from '@chakra-ui/react'
import {ethers} from 'ethers'
import {parseEther } from 'ethers/lib/utils'
import {abi} from '../../../artifacts/contracts/Union.sol/Union.json'
import { Contract } from "ethers"
import { TransactionResponse,TransactionReceipt } from "@ethersproject/abstract-provider"

interface Props {
    addressContract: string,
    currentAccount: string | undefined
}

declare let window: any;

export default function ApplyDaoTG(props:Props){
  const addressContract = props.addressContract
  const currentAccount = props.currentAccount
  var [user_id, setUserId] = useState<number>(0)
  var [chat_id,setChatId] = useState<number>(0)
  var [DAOaddress,setDAOaddress] = useState<string>("")

  var [votingType,setVotingType] = useState<string>("")
  var [votingTokenContract,setVotingTokenContract] = useState<string>("")

  var [DAOname,setDAOname] = useState<string>("")
  const [user_name, setUserName] = useState<string>("")

  //after the page is rendered, we get the data from the URL with this hook
  //as this is done only once, there are no conflicts with editing that data later via the UI
  useEffect(() => {
    const queryParams = new URLSearchParams(location.search);
  var user = queryParams.get('user_id');
  let int_user : number = +user;
  var chat = queryParams.get('chat_id');
  let int_chat : number = +chat;
  
  var address = queryParams.get('address')

  var type = queryParams.get('votingtype');
  var contract = queryParams.get('votingtokencontract');
  var name = queryParams.get('daoname');
 
  setUserId(int_user);
  setChatId(int_chat);
  setDAOaddress(address)
  setVotingType(type);
  setVotingTokenContract(contract);
  setDAOname(name);
  
  }, []);

  async function applyDao(event:React.FormEvent) {
    event.preventDefault()
    if(!window.ethereum) return    
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const Union:Contract = new ethers.Contract(addressContract, abi, signer)

    Union.ApplyForUnion(user_id,chat_id,DAOaddress,votingType,votingTokenContract,DAOname,{value:ethers.utils.formatUnits(2000000000000000,"wei")})
     .then((tr: TransactionResponse) => {
        console.log(`TransactionResponse TX hashonChange={(e) => setUserId(e.target.value)} value={query.user_id} my={3} : ${tr.hash}`)
        tr.wait().then((receipt:TransactionReceipt) => {console.log("applying receipt", receipt)})
        })
         .catch((e:Error) => console.log(e))
     }

  
  //const handleChange = (value:string) => setUserId(value) 
//localhost:3000/dao?user_id=1337&chat_id=1337&address=23746624386&votingtype=1&votingtokencontract=3278465ASDW23&daoname=lol
  return (
    <form onSubmit={applyDao}>
    <FormControl>
      <FormLabel htmlFor='TGID'>Enter data: </FormLabel>
      <Input id="user_id" type="text" placeholder="Your Telegram ID" required  onChange={(e) => setUserId(parseInt(e.target.value))} value = {user_id} my={3}/>
      <Input id="chat_id" type="text" placeholder="Chat's Telegram ID"required  onChange={(e) => setChatId(parseInt(e.target.value))} value = {chat_id} my={3}/>
      <Input id="DAOaddress" type="text" placeholder="Your Multisig address"required  onChange={(e) => setDAOaddress(e.target.value)} value= {DAOaddress}  my={3}/>

      <Select id="votingtype" placeholder="Select voting token's type:" onChange={(e) => setVotingType(e.target.value)} value= {votingType}  my={3}>
      <option value='0'>ERC20</option>
      <option value='1'>ERC20Snapshot</option>
      <option value='2'>ERC721</option>
      </Select>
      <Input id="votingtokencontract" placeholder="Voting token's contract address" type="text" required  onChange={(e) => setVotingTokenContract(e.target.value)} value = {votingTokenContract} my={3}/>
      <Input id="daoname" placeholder="Your DAO name" type="text" required  onChange={(e) => setDAOname(e.target.value)} value = {DAOname} my={3}/>
      <Button type="submit" isDisabled={!currentAccount}>Apply DAO for Union</Button>
    </FormControl>
    </form>

    
  )
}
