import { BrowserRouter, Route, Routes } from "react-router-dom"
import "./App.css"
import Navbar from "./components/molecules/Navbar"
import Home from "./pages/Home"
import Login from "./pages/Login"
import Signup from "./pages/Signup"
import Trades from "./pages/Trades"
import Transactions from "./pages/Transactions"

//TODO:
// - AuthProvider abstraction
// - AuthContext

function App() {
  return (
    <BrowserRouter>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/signup" element={<Signup />} />
        <Route path="/login" element={<Login />} />
        <Route path="/transactions" element={<Transactions />} />
        <Route path="/trades" element={<Trades />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
