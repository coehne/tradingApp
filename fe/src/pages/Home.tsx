import React, { useEffect } from "react"
import Depot from "./Depot"
import { useAuth } from "../context/AuthContext"

function Home() {
  const { user } = useAuth()
  useEffect(() => {
    console.log(user)
  }, [user])

  return (
    <div>
      <p>{`Hi ${user?.firstName}`}</p>
      <Depot />
    </div>
  )
}

export default Home
