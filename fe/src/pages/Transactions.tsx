import { useEffect, useState } from "react"
import { Link } from "react-router-dom"
import { Transaction } from "../models/Transaction"
import axios from "../utils/apiClient"
import { numberToUSD, stringToDate } from "../utils/formatting"

function Transactions() {
  const [data, setData] = useState<Transaction[] | undefined>()
  useEffect(() => {
    axios
      .get("transaction")
      .then((res) => setData(res.data))
      .catch((error) => console.log(error))
  }, [])

  console.log(data)
  return (
    <div className="min-h-screen bg-gray-200 flex flex-col ">
      <div className="max-w-5xl w-full mx-auto my-10">
        <Link
          to="/transactions/deposit"
          className="py-2 px-5 bg-primary text-black font-bold rounded hover:bg-green-600 transition duration-300"
        >
          Deposit
        </Link>
        <Link
          to="/transactions/withdraw"
          className="mx-4 py-2 px-5 bg-primary text-black font-bold rounded hover:bg-green-600 transition duration-300"
        >
          Withdraw
        </Link>
      </div>
      <div className="max-w-5xl w-full mx-auto">
        <div className="relative overflow-x-auto shadow-md sm:rounded-lg">
          <table className="w-full text-sm text-left text-gray-500 dark:text-gray-400">
            <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
              <tr>
                <th scope="col" className="px-6 py-3">
                  ID
                </th>
                <th scope="col" className="px-6 py-3">
                  Type
                </th>
                <th scope="col" className="px-6 py-3">
                  Amount
                </th>
                <th scope="col" className="px-6 py-3">
                  Time
                </th>
                <th scope="col" className="px-6 py-3">
                  <span className="sr-only">Edit</span>
                </th>
              </tr>
            </thead>
            <tbody>
              {data?.map((tx) => {
                return (
                  <tr
                    className="bg-white border-b dark:bg-gray-800 dark:border-gray-700"
                    key={tx.id}
                  >
                    <th
                      scope="row"
                      className="px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap"
                    >
                      {tx.id}
                    </th>
                    <td className="px-6 py-4">
                      {tx.tradeId === 0 && tx.amount < 0
                        ? "Cash withdraw"
                        : tx.tradeId === 0 && tx.amount > 0
                        ? "Cash deposit"
                        : `Trade: ${tx.tradeId}`}
                    </td>
                    <td className="px-6 py-4">{numberToUSD(tx.amount)}</td>
                    <td className="px-6 py-4">{stringToDate(tx.createdAt)}</td>
                    <td className="px-6 py-4 text-right">
                      {tx.tradeId === 0 ? (
                        ""
                      ) : (
                        <a
                          href={`trade/${tx.tradeId}`}
                          className="font-medium text-blue-600 dark:text-blue-500 hover:underline"
                        >
                          Trade
                        </a>
                      )}
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

export default Transactions
