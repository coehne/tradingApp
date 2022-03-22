import { BrowserRouter, Route, Routes } from "react-router-dom"
import { AuthProvider } from "./context/AuthContext"
import "./App.css"
import Navbar from "./components/molecules/Navbar"
import Home from "./pages/Home"
import Login from "./pages/Login"
import Signup from "./pages/Signup"
import Transactions from "./pages/Transactions"

function App() {
  return (
    <BrowserRouter>
      <AuthProvider>
        <Navbar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/signup" element={<Signup />} />
          <Route path="/login" element={<Login />} />
          <Route path="/transactions" element={<Transactions />} />
        </Routes>
      </AuthProvider>
    </BrowserRouter>
  )
}

export default App
