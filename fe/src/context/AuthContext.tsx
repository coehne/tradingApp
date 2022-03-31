import axios from "../utils/apiClient"
import {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useMemo,
} from "react"
import * as auth from "../AuthProvider"
import { User } from "../models/User"
import FullPageSpinner from "../components/FullPageSpinner"
import FullPageErrorFallback from "../components/FullPageErrorFallback"
import { useAsync } from "../hooks/useAsync"

interface AuthContextType {
  user: User | null
  login: (form: { email: string; password: string }) => Promise<any>
  signup: (form: {
    email: string
    password: string
    firstName: string
  }) => Promise<any>
  logout: () => void
}

// Get user data from me endpoint
const bootstrapAppData = async () => {
  // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
  let user: User = null!

  await axios
    .get<User>("identity/me")
    .then((res) => (user = res.data))
    .catch(() => {})

  return user
}

let AuthContext = createContext<AuthContextType>(null!)

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

  // Get user data when initialize
  useEffect(() => {
    const appDataPromise = bootstrapAppData()
    run(appDataPromise)
  }, [run])

  const login = useCallback(
    (form) => auth.login(form).then((user) => setData(user as any)), //FIXME: Lazy typing.
    [setData]
  )
  const signup = useCallback(
    (form) => auth.signup(form).then((user) => setData(user as any)), //FIXME: Lazy typing.
    [setData]
  )

  const logout = useCallback(() => {
    auth.logout()
    setData(null as any)
  }, [setData])

  const value = useMemo(
    () => ({ user, login, logout, signup }),
    [signup, login, logout, user]
  )
  if (isLoading || isIdle) {
    return <FullPageSpinner />
  }

  if (isError) {
    return <FullPageErrorFallback error={error} />
  }

  if (isSuccess) {
    return <AuthContext.Provider value={value} {...props} />
  }

  throw new Error(`Unhandled status: ${status}`)
}

const useAuth = () => {
  const context = useContext(AuthContext)
  if (context === undefined) {
    throw new Error(`useAuth must be used within a AuthProvider`)
  }
  return context
}

export { AuthProvider, useAuth }
