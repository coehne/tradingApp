import { useEffect } from "react"
import { useForm } from "react-hook-form"
import { useNavigate } from "react-router-dom"
import { FormContainer, InputText } from "../components/atoms/FormElements"
import { useAuth } from "../context/AuthContext"
import { useAsync } from "../hooks/useAsync"

interface FormData {
  email: string
  password: string
}

function Login() {
  const navigate = useNavigate()
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({ mode: "onSubmit" })
  const { login } = useAuth()
  const { run, isSuccess } = useAsync<any>()
  const onSubmit = handleSubmit(({ email, password }) => {
    run(login({ email, password }))
  })
  useEffect(() => {
    isSuccess && navigate("/", { replace: true })
  }, [isSuccess, navigate])

  return (
    <div className="min-h-screen bg-gray-200 flex flex-col justify-center">
      <div className="max-w-md w-full mx-auto">
        <div className="text-3xl font-bold text-gray-900 mt-2 text-center">
          Login
        </div>
        <FormContainer>
          <form onSubmit={onSubmit} className="space-y-6">
            <InputText
              errors={errors}
              registerHandler={() =>
                register("email", {
                  required: "This is a required field",
                })
              }
              name="email"
              type="text"
            >
              Email
            </InputText>
            <InputText
              errors={errors}
              registerHandler={() =>
                register("password", {
                  required: "This is a required field",
                })
              }
              name="password"
              type="password"
            >
              Password
            </InputText>

            <div>
              <button className="w-full py-2 px-4 bg-primary hover:bg-green-600 rounded text-black font-bold">
                Submit
              </button>
            </div>
          </form>
        </FormContainer>
      </div>
    </div>
  )
}

export default Login