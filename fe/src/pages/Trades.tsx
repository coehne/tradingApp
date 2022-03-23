import React, { useEffect, useState } from "react"
import { Trade } from "../models/Trade"
import axios from "../utils/apiClient"
import { numberToUSD } from "../utils/formatting"

function Trades() {
  const [data, setData] = useState<Trade[] | undefined>(undefined)
  useEffect(() => {
    axios
      .get("trade")
      .then((res) => setData(res.data))
      .catch((error) => console.log(error))
  }, [])

  return (
    <div className="min-h-screen bg-gray-200 flex flex-col ">
      <div className="max-w-5xl w-full mx-auto my-10">
        <a
          href="/trades/buy"
          className="py-2 px-5 bg-primary text-black font-bold rounded hover:bg-green-600 transition duration-300"
        >
          Buy
        </a>
        <a
          href="/trades/sell"
          className="mx-4 py-2 px-5 bg-primary text-black font-bold rounded hover:bg-green-600 transition duration-300"
        >
          Sell
        </a>
      </div>
      <div className="max-w-5xl w-full mx-auto">
        <div className="relative overflow-x-auto shadow-md sm:rounded-lg">
          <table className="w-full text-sm text-left text-gray-500 dark:text-gray-400">
            <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400 text-center">
              <tr>
                <th scope="col" className="px-6 py-3">
                  ID
                </th>
                <th scope="col" className="px-6 py-3">
                  Company Name
                </th>
                <th scope="col" className="px-6 py-3">
                  Symbol
                </th>
                <th scope="col" className="px-6 py-3">
                  Quantity
                </th>
                <th scope="col" className="px-6 py-3">
                  Price
                </th>
                <th scope="col" className="px-6 py-3">
                  Total
                </th>
              </tr>
            </thead>
            <tbody>
              {data?.map((trade) => {
                return (
                  <tr
                    className="bg-white border-b dark:bg-gray-800 dark:border-gray-700 text-center"
                    key={trade.id}
                  >
                    <th
                      scope="row"
                      className="px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap"
                    >
                      {trade.id}
                    </th>
                    <td className="px-6 py-4">{trade.companyName}</td>
                    <td className="px-6 py-4">{trade.symbol}</td>
                    <td className="px-6 py-4">{trade.qty}</td>
                    <td className="px-6 py-4">{numberToUSD(trade.price)}</td>
                    <td className="px-6 py-4 ">
                      {numberToUSD(trade.qty * trade.price)}
                    </td>
                  </tr>
                )
              })}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  )
}

export default Trades
