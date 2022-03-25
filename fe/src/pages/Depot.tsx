import React, { useEffect, useState } from "react"
import { Link, Navigate } from "react-router-dom"
import FullPageSpinner from "../components/FullPageSpinner"
import { Trade } from "../models/Trade"
import { Transaction } from "../models/Transaction"
import axios from "../utils/apiClient"
import { numberToUSD } from "../utils/formatting"

function Depot() {
  const [data, setData] = useState<Trade[] | undefined>(undefined)
  const [txData, setTxData] = useState<Transaction[] | undefined>(undefined)
  const [tradeIsLoading, setTradeIsLoading] = useState(true)
  const [transactionIsLoading, setTransactionIsLoading] = useState(true)

  useEffect(() => {
    setTradeIsLoading(true)
    setTransactionIsLoading(true)
    axios
      .get("trades")
      .then((res) => setData(res.data))
      .then(() => setTradeIsLoading(false))
      .catch((error) => console.log(error))

    axios
      .get("transaction")
      .then((res) => setTxData(res.data))
      .then(() => setTransactionIsLoading(false))
      .catch((error) => console.log(error))
  }, [])

  let sumStocks = 0

  // Calculate the total cash of the user
  let sumCash = 0
  txData?.map((tx) => (sumCash += tx.amount))

  if (tradeIsLoading === false && transactionIsLoading === false) {
    return (
      <div className="min-h-screen bg-gray-200 flex flex-col ">
        {txData?.length === 0 ? (
          <Navigate to={"/transactions"} />
        ) : (
          <>
            <div className="max-w-5xl w-full mx-auto my-10">
              <Link
                to="/trades/buy"
                className="py-2 px-5 bg-primary text-black font-bold rounded hover:bg-green-600 transition duration-300"
              >
                Buy
              </Link>
              <Link
                to="/trades/sell"
                className="mx-4 py-2 px-5 bg-primary text-black font-bold rounded hover:bg-green-600 transition duration-300"
              >
                Sell
              </Link>
            </div>
            <div className="max-w-5xl w-full mx-auto">
              <div className="relative overflow-x-auto shadow-md sm:rounded-lg">
                <table className="w-full text-sm text-left text-gray-500 dark:text-gray-400">
                  <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400 text-center">
                    <tr>
                      <th scope="col" className="px-6 py-3">
                        Position
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
                    {data?.map((trade, i) => {
                      sumStocks += trade.qty * trade.price
                      return (
                        <tr
                          className="bg-white border-b dark:bg-gray-800 dark:border-gray-700 text-center"
                          key={trade.symbol}
                        >
                          <th
                            scope="row"
                            className="px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap"
                          >
                            {i + 1}
                          </th>
                          <td className="px-6 py-4">{trade.companyName}</td>
                          <td className="px-6 py-4">{trade.symbol}</td>
                          <td className="px-6 py-4">{trade.qty}</td>
                          <td className="px-6 py-4">
                            {numberToUSD(trade.price)}
                          </td>
                          <td className="px-6 py-4 ">
                            {numberToUSD(trade.qty * trade.price)}
                          </td>
                        </tr>
                      )
                    })}
                    <tr className="bg-white border-b dark:bg-gray-800 dark:border-gray-700 text-center">
                      <td></td>
                      <td></td>
                      <td></td>
                      <td></td>
                      <td className="px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap">
                        Sum stocks:
                      </td>
                      <td>{numberToUSD(sumStocks)}</td>
                    </tr>
                    <tr className="bg-white border-b dark:bg-gray-800 dark:border-gray-700 text-center">
                      <td></td>
                      <td></td>
                      <td></td>
                      <td></td>
                      <td className="px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap">
                        Cash:
                      </td>
                      <td>{numberToUSD(sumCash)}</td>
                    </tr>
                    <tr className="bg-white border-b dark:bg-gray-800 dark:border-gray-700 text-center">
                      <td></td>
                      <td></td>
                      <td></td>
                      <td></td>
                      <td className="px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap">
                        Total:
                      </td>
                      <td className="font-medium text-gray-900 dark:text-white whitespace-nowrap underline">
                        {numberToUSD(sumCash + sumStocks)}
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </>
        )}
      </div>
    )
  }

  return <FullPageSpinner />
}

export default Depot
