import axios from "./utils/apiClient"
import { User } from "./models/User"
import { AxiosError } from "axios"

export interface SignupForm {
  firstName: string
  email: string
  password: string
}

export interface LoginForm {
  email: string
  password: string
}

const handleUserResponse = ({ id, firstName }: User) => {
  return { id, firstName }
}

const login = ({
  email,
  password,
}: LoginForm) => {
  return axios.post("identity/login/", { email, password }).then((res) =>
    handleUserResponse(res.data)
  ).catch((error: AxiosError) => Promise.reject({
    message: error.message,
    statusCode: error.response?.status
  }
  ))
}
const signup = ({
  firstName,
  email,
  password,
}: SignupForm) => {
  return axios.post("identity/signup/", { email, password, firstName }).then((res) =>
    handleUserResponse(res.data)
  ).catch((error: AxiosError) => Promise.reject({
    message: error.message,
    statusCode: error.response?.status
  }
  ))
}

const logout = async () => {
  axios.post("identity/logout")

}

export {
  login,
  logout,
  signup
}
