import { ErrorMessage } from "@hookform/error-message"
import axios from "../utils/apiClient"
import { useForm } from "react-hook-form"
import { useNavigate } from "react-router-dom"

type TransactionType = "withdraw" | "deposit"

interface FormData {
  amount: number
}

const NewTsx: React.FC<{ type: TransactionType }> = ({ type }) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({ mode: "onSubmit" })
  const navigate = useNavigate()

  const onSubmit = handleSubmit(({ amount }) => {
    const formData: FormData = {
      amount: parseInt(amount as unknown as string),
    }
    if (type === "withdraw") {
      formData.amount = formData.amount * -1
    }
    axios
      .post("/transaction/", formData)
      .then(() => navigate("/transactions"))
      .catch((error) => console.log(error))
  })

  return (
    <div>
      <div className="max-w-md w-full-md mx-auto border p-8 border-gray-300 mt-4  bg-gray-700 text-gray-200 rounded-md">
        <div className="text-3xl font-bold  mt-2 text-center my-5 uppercase">
          {type}
        </div>
        <form onSubmit={onSubmit}>
          <label className="text-sm font-bold block">{`How much do you want to ${type}`}</label>
          <input
            {...register("amount", {
              validate: {
                positive: (v) => parseInt(v as unknown as string) > 0,
              },
            })}
            type="number"
            className="w-full p-2 border border-gray-300 rounded mt-1 text-black"
          />
          <ErrorMessage errors={errors} name={"amount"} />
          <button className="w-full my-8 py-2 px-4 bg-primary hover:bg-green-600 rounded uppercase text-black font-bold">
            {`${type} now!`}
          </button>
        </form>
      </div>
    </div>
  )
}

export default NewTsx
