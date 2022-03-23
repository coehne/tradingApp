import { BrowserRouter, Route, Routes } from "react-router-dom"
import "./App.css"
import Navbar from "./components/molecules/Navbar"
import Home from "./pages/Home"
import Login from "./pages/Login"
import NewTransaction from "./pages/NewTransaction"
import Signup from "./pages/Signup"
import Trade from "./pages/Trade"
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
        <Route
          path="/transactions/deposit"
          element={<NewTransaction type={"deposit"} />}
        />
        <Route
          path="/transactions/withdraw"
          element={<NewTransaction type={"withdraw"} />}
        />
        <Route path="/trades" element={<Trades />} />
        <Route path="/trade/:id" element={<Trade />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
