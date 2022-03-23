import React, { useEffect, useState } from "react"
import { Trade } from "../models/Trade"
import axios from "../utils/apiClient"

function Trades() {
  const [data, setData] = useState<Trade[] | undefined>(undefined)
  useEffect(() => {
    axios
      .get("trade")
      .then((res) => setData(res.data))
      .catch((error) => console.log(error))
  }, [])

  console.log(data)
  return (
    <div>
      {data?.map((t) => (
        <p>{t.symbol}</p>
      ))}
    </div>
  )
}

export default Trades
