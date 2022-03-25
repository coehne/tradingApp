import { useForm } from "react-hook-form"
import { FormContainer, InputText } from "../components/FormElements"
import { useAsync } from "../hooks/useAsync"
import { useAuth } from "../context/AuthContext"
import { Alert } from "../components/Alert"

interface FormData {
  email: string
  password: string
}

function Login() {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({ mode: "onSubmit" })

  const { run, error } = useAsync<any>()
  const { login } = useAuth()

  const onSubmit = handleSubmit(({ email, password }) =>
    run(login({ email, password }))
  )

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
            {error && error.statusCode && error.statusCode >= 400 && (
              <Alert msg={"Your password or email is incorrect!"} />
            )}

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
