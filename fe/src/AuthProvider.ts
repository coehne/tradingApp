import axios from "./utils/apiClient"
import { User } from "./models/User"

const handleUserResponse = ({id, firstName, email}: User) => {
  
  return {id, firstName, email} 
}

const login = ({
  email,
  password,
}: {
  email: string
  password: string
}) => {
  return axios.post("identity/login/", { email, password }).then(() =>
    handleUserResponse
  )
}
const signup = ({
  firstName,
  email,
  password,
}: {
  firstName: string
  email: string
  password: string
}) => {
  return axios.post("identity/signup/", { email, password, firstName }).then(() =>
    handleUserResponse
  )
}

const logout = async () => {
    axios.post("identity/logout")
 
}


export {
  login,
  logout,
  signup
}
