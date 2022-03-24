import React from "react"
import { BuyForm } from "../components/BuyForm"
import { SellForm } from "../components/SellForm"

type TradeType = "buy" | "sell"

const NewTrade: React.FC<{ type: TradeType }> = ({ type }) => {
  if (type === "buy") {
    return <BuyForm />
  }
  return <SellForm />
}

export default NewTrade
