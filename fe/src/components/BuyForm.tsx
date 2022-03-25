import axios from "../utils/apiClient"
import { useForm } from "react-hook-form"
import { useNavigate } from "react-router-dom"
import { ErrorMessage } from "@hookform/error-message"
import { useEffect, useState } from "react"
import { Quote } from "../models/Quote"
import { numberToUSD } from "../utils/formatting"

interface FormData {
  symbol: string
  qty: number
}

export const BuyForm = () => {
  const [quote, setQuote] = useState<Quote>(null!)
  const {
    register,
    handleSubmit,
    formState: { errors },
    watch,
  } = useForm<FormData>({ mode: "onSubmit" })
  const navigate = useNavigate()

  const onSubmit = handleSubmit(({ qty, symbol }) => {
    const formData: FormData = {
      qty: parseInt(qty as unknown as string),
      symbol: symbol,
    }

    axios
      .post("/trade/", formData)
      .then(() => navigate("/"))
      .catch((error) => console.log(error))
  })

  useEffect(() => {
    const delayDebounce = setTimeout(() => {
      axios
        .get(`/quote/${watch("symbol")}`)
        .then((res) => setQuote(res.data))
        .catch((error) => setQuote(null!))
    }, 500)

    return () => clearTimeout(delayDebounce)
  }, [watch("symbol")])
  return (
    <div>
      <div className="max-w-md w-full-md mx-auto border p-8 border-gray-300 mt-4  bg-gray-700 text-gray-200 rounded-md">
        <div className="text-3xl font-bold  mt-2 text-center my-5 uppercase">
          Buy Stonks!
        </div>
        <form onSubmit={onSubmit}>
          <fieldset>
            <label className="text-sm font-bold block">
              Which stonk do you want to buy?
            </label>
            <input
              {...register("symbol", {
                required:
                  "Please provide the symbol of the stock you want to buy",
              })}
              type="text"
              className="w-full p-2 border border-gray-300 rounded mt-1 text-black"
              placeholder="Symbol"
            />
            <ErrorMessage errors={errors} name={"amount"} />
            {quote !== null ? (
              <p className="mt-4">{`Current price for ${
                quote.companyName
              } is ${numberToUSD(quote.latestPrice)}`}</p>
            ) : (
              <p className="mt-4">
                Please provide a valid stock symbol to get a quote!
              </p>
            )}
            <label className="text-sm font-bold block mt-5">
              How many stonks do you want to buy?
            </label>
            <input
              {...register("qty", {
                validate: {
                  positive: (v) => parseInt(v as unknown as string) > 0,
                },
              })}
              type="number"
              className="w-full p-2 border border-gray-300 rounded mt-1 text-black"
              placeholder="Volume"
            />
            <ErrorMessage errors={errors} name={"amount"} />
            {quote !== null && watch("qty") !== 0 ? (
              <p className="mt-4">{`Total order cost is ${numberToUSD(
                quote.latestPrice * watch("qty")
              )}`}</p>
            ) : (
              <></>
            )}
          </fieldset>
          <button className="w-full mt-8 py-2 px-4 bg-primary hover:bg-green-600 rounded uppercase text-black font-bold">
            Submit Order
          </button>
        </form>
      </div>
    </div>
  )
}
