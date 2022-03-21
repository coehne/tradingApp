import { User } from "./models/User"

const handleUserResponse = ({
  id,
  firstName,
  email,
}: User) => {
 
   return {
    id,
    firstName,
    email,
  }
}

const login = ({
  email,
  password,
}: {
  email: string
  password: string
}) => {
  return client("identity/login/", { email, password }).then(
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
  return client("identity/signup/", { email, password, firstName }).then(
    handleUserResponse
  )
}

const logout = async () => {
    //TODO:
 
}

const API_AUTH_URL = "http://localhost:8000/api" // Define API_BASE_URL from .env

const client = <ClientData,>(
  endpoint: string,
  data?: ClientData,
  method?: string,
  token?: string
) => {
  const config = {
    method: method ? method : "POST",
    body: data ? JSON.stringify(data) : undefined,
    headers: {
      "Content-Type": "application/json",
    credentials: "include",
    },
  }

  return window
    .fetch(`${API_AUTH_URL}/${endpoint}`, config)
    .then(async (response) => {
      if (response.status === 204) {
        return Promise.resolve()
      }
      const data = await response.json()

      if (response.ok) {
        return data
      } else {
        return Promise.reject({
          message: data.message,
          statusCode: response.status,
        })
      }
    })
}

export {
  login,
  logout,
  signup
}
