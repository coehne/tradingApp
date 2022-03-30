import { useAuth } from "../context/AuthContext"
import FullPageSpinner from "../components/FullPageSpinner"
import { Suspense } from "react"
import { AuthenticatedApp } from "./AuthenticatedApp"
import UnauthenticatedApp from "./UnauthenticatedApp"

function Home() {
  const { user } = useAuth()
  const baseurl = process.env.REACT_APP_API_BASE_URL
  console.log(baseurl)
  return (
    <Suspense fallback={<FullPageSpinner />}>
      {user ? <AuthenticatedApp /> : <UnauthenticatedApp />}
    </Suspense>
  )
}

export default Home
