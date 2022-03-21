import { ErrorMessage } from "@hookform/error-message"
import React from "react"
import { useForm } from "react-hook-form"

interface FormData {
  firstName: string
  email: string
  password: string
}

function Signup() {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({ mode: "onTouched" })

  const onSubmit = handleSubmit((data) => {
    console.log(data)
  })

  return (
    <div className="min-h-screen bg-gray-200 flex flex-col justify-center">
      <div className="max-w-md w-full mx-auto">
        <div className="text-3xl font-bold text-gray-900 mt-2 text-center">
          Signup
        </div>
        <div className="text-center font-medium text-xl mt-2">
          Signup now for free and start trading!
        </div>
        <div className="max-w-md w-full-md mx-auto bg-white border p-8 border-gray-300 mt-4">
          <form onSubmit={onSubmit} className="space-y-6">
            <div>
              <label
                htmlFor=""
                className="text-sm font-bold text-gray-600 block"
              >
                First Name
              </label>
              <input
                {...register("firstName", {
                  required: "Please fill in your first name",
                })}
                type="text"
                className="w-full p-2 border border-gray-300 rounded mt-1"
              />
              <ErrorMessage errors={errors} name="firstName" />
            </div>
            <div>
              <label
                htmlFor=""
                className="text-sm font-bold text-gray-600 block"
              >
                Email
              </label>
              <input
                {...register("email", {
                  required: "Please fill in your email",
                })}
                type="text"
                className="w-full p-2 border border-gray-300 rounded mt-1"
              />
              <ErrorMessage errors={errors} name="email" />
            </div>
            <div>
              <label
                htmlFor=""
                className="text-sm font-bold text-gray-600 block"
              >
                Password
              </label>
              <input
                {...register("password", {
                  required: "Please fill in your password",
                })}
                type="password"
                className="w-full p-2 border border-gray-300 rounded mt-1"
              />
              <ErrorMessage errors={errors} name="password" />
            </div>
            <div className="flex items-center justify-end">
              <a href="" className="font-medium text-sm text-green-600">
                Forgot Password?
              </a>
            </div>
            <div>
              <button className="w-full py-2 px-4 bg-primary hover:bg-green-600 rounded text-black font-bold">
                Submit
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  )
}

export default Signup
