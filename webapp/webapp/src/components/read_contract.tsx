import React, {useEffect, useState } from 'react';
import {Text} from '@chakra-ui/react'
//import {ERC20ABI as abi} from 'abi/ERC20ABI'
import {abi as abi} from '../../../../artifacts/contracts/TGPassport.sol/TGPassport.json'

import {Contract, ethers} from 'ethers'

//import TGPassport from '../../../../artifacts/contracts/TGPassport.sol/TGPassport.json'




interface Props {
    addressContract: string,
    currentAccount: string | undefined
}

declare let window: any;

export default function ReadERC20(props:Props){
  const addressContract = props.addressContract
  const currentAccount = props.currentAccount
  const [totalSupply,setTotalSupply]=useState<string>()
  const [symbol,setSymbol]= useState<string>("")
  const [owner, setOwner] = useState<string>()


  // called only once, use it as constructor
  useEffect( () => {
    if(!window.ethereum) return

    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const TGPassport = new ethers.Contract(addressContract, abi, provider);
    TGPassport.owner().then((result:string)=>{
        setOwner(result)
    }).catch('error', console.error);

    /*
    erc20.totalSupply().then((result:string)=>{
        setTotalSupply(ethers.utils.formatEther(result))
    }).catch('error', console.error);
    */
    //called only once
  },[])  

  return (
    /**
    <div>
        <Text><b>Telegram Passport Contract</b>: {addressContract}</Text>
        <Text><b>ClassToken totalSupply</b>:{totalSupply} {symbol}</Text>
        <Text my={4}><b>ClassToken in current account</b>:</Text>
    </div>
    */

    <div>
        <Text><b>Telegram Passport Contract</b>: {addressContract}</Text>
        <Text><b>Owner of contract</b>:{owner}</Text>
        <Text my={4}><b>ClassToken in current account</b>:</Text>
    </div>

  )
}