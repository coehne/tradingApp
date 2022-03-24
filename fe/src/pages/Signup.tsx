import { useForm } from "react-hook-form"
import { FormContainer, InputText } from "../components/FormElements"
import axios from "../utils/apiClient"
import { useNavigate } from "react-router-dom"

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
  } = useForm<FormData>({ mode: "onSubmit" })

  const navigate = useNavigate()

  const onSubmit = handleSubmit(({ email, password, firstName }) => {
    axios
      .post("/identity/signup", { email, password, firstName })
      .then(() => navigate("/", { replace: true }))
      .catch((error) => console.log(error))
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
        <FormContainer>
          <form onSubmit={onSubmit} className="space-y-6">
            <InputText
              errors={errors}
              registerHandler={() =>
                register("firstName", {
                  required: "This is a required field",
                })
              }
              name="firstName"
              type="text"
            >
              First Name
            </InputText>
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

export default Signup
