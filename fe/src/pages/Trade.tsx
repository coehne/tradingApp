import React, { useEffect, useState } from "react"
import { TradeCard } from "../components/molecules/TradeCard"
import { useParams } from "react-router-dom"
import axios from "../utils/apiClient"
import { Trade as ITrade } from "../models/Trade"

function Trade() {
  const { id } = useParams()

  const [data, setData] = useState<ITrade | undefined>(undefined)
  useEffect(() => {
    axios
      .get(`trade/${id}`)
      .then((res) => setData(res.data))
      .catch((error) => console.log(error))
  }, [id])
  return (
    <div>
      <div className="min-h-screen bg-gray-200 flex flex-col py-8">
        <div className="max-w-md w-full mx-auto"></div>
        <TradeCard trade={data} />
      </div>
    </div>
  )
}

export default Trade
