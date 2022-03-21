import * as React from "react"
import * as auth from "../AuthProvider"
import { client } from "../utils/apiClient"
import { User } from "../models/User"
import { useAsync } from "../hooks/useAsync"

interface AuthContextType {
  user: User | null
  login: (form: { email: string; password: string }) => Promise<any>
  logout: () => void
  signup: (form: {
    email: string
    password: string
    firstName: string
  }) => Promise<any>
}

const bootstrapAppData = async () => {
  // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
  let user: User = null!

  const data = await client<User>("identity/me", {})
  user = data

  return user
}

const AuthContext = React.createContext<AuthContextType>(null!)
AuthContext.displayName = "AuthContext"

const AuthProvider = (props: { children: React.ReactNode }) => {
  const {
    state: { data: user },
    status,
    error,
    isLoading,
    isIdle,
    isError,
    isSuccess,
    run,
    setData,
  } = useAsync<User>()

  React.useEffect(() => {
    const appDataPromise = bootstrapAppData()
    run(appDataPromise)
  }, [run])

  const login = React.useCallback(
    (form) => auth.login(form).then((user) => setData(user)),
    [setData]
  )
  const signup = React.useCallback(
    (form) => auth.signup(form).then((user) => setData(user)),
    [setData]
  )

  const logout = React.useCallback(() => {
    auth.logout()
    setData(null as any)
  }, [setData])

  const value = React.useMemo(
    () => ({ user, login, logout, signup }),
    [login, logout, user, signup]
  )

  if (isLoading || isIdle) {
    return <div> Todo: FullPageSpinner </div>
  }

  if (isError) {
    return <div>Todo: FullPage Error </div>
  }

  if (isSuccess) {
    return <AuthContext.Provider value={value} {...props} />
  }

  throw new Error(`Unhandled status: ${status}`)
}

const useAuth = () => {
  const context = React.useContext(AuthContext)
  if (context === undefined) {
    throw new Error(`useAuth must be used within a AuthProvider`)
  }
  return context
}

const useClient = () => {
  return React.useCallback(
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    <Data,>(endpoint: string, config?: Record<string, any>) =>
      client<Data>(endpoint, {
        ...config,
      }),
    []
  )
}
export { AuthProvider, useAuth, useClient }
