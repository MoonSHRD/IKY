import React, { useState, useEffect } from 'react'
import {Button, Input , NumberInput,  NumberInputField,  FormControl,  FormLabel, Text } from '@chakra-ui/react'
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


export default function NOTTrustTGID(props:Props){
  const addressContract = props.addressContract
  const currentAccount = props.currentAccount
  var [user_id, setUserId] = useState(0)
  var [friend_id, setFriendId] = useState(0)
 // var [user_name, setUserName] = useState<string>("")
  var [friend_user_name,setFriendUserName] = useState<string>("")

  useEffect(() => {
    const queryParams = new URLSearchParams(location.search);

  var id = queryParams.get('user_tg_id');   // get id as string from query
  let int_id : number = +id;                // similar to parseInt()
  
  var f_id = queryParams.get('friend_tg_id');
  let f_int_id : number = +f_id; 
  
  var friend_user_name = queryParams.get('friend_user_name');
  

  setUserId(int_id);
  setFriendId(f_int_id);
  setFriendUserName(friend_user_name);
  
  }, []);
  

async function untrustToTgID(event:React.FormEvent) {
    event.preventDefault()
    if(!window.ethereum) return    
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const TGPassport:Contract = new ethers.Contract(addressContract, abi, signer)
        TGPassport.INotTrustID(user_id, friend_id)
         .then((tr: TransactionResponse) => {
            console.log(`TransactionResponse TX hash: ${tr.hash}`)
            tr.wait().then((receipt:TransactionReceipt) => {console.log("trust receipt", receipt)})
            })
             .catch((e:Error) => console.log(e))  
    }
  
  //const handleChange = (value:string) => setUserId(value)
  //http://localhost:3000?user_tg_id=1337&friend_tg_id=1997&friend_user_name=sbekket


    /*
      <div>
       // <Text><b>Ethereum address associated with this telegram nickname</b>: {user_wallet}</Text>
    </div>
    */
  return (
    <form onSubmit={untrustToTgID}>
    <div>
        <Text><b>You about to NOT TRUST this username:</b>: {friend_user_name}</Text>
    </div>
    <FormControl>
      <FormLabel htmlFor='TGID'>Do not change variables </FormLabel>
      <Input id="tgid" type="number" required  onChange={(e) => setUserId(parseInt(e.target.value))} value={user_id} my={3}/>
      <Input id="friend_tgid" type="number" required  onChange={(e) => setFriendId(parseInt(e.target.value))} value={user_id} my={3}/>

      <Button type="submit" isDisabled={!currentAccount}>NOT Trust this ID!</Button>
    </FormControl>
    </form>
  )
}