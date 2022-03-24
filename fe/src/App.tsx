import { Outlet, Route, Routes } from "react-router-dom"
import "./App.css"
import Navbar from "./components/molecules/Navbar"
import Home from "./pages/Home"
import Login from "./pages/Login"
import NewTrade from "./pages/NewTrade"
import NewTsx from "./pages/NewTsx"
import Signup from "./pages/Signup"
import Trade from "./pages/Trade"
import TradeHistory from "./pages/Trades"
import Transactions from "./pages/Transactions"

//TODO:
// - AuthProvider abstraction
// - AuthContext
// - Handle loading state
// - Handle error state

function App() {
  return (
    <Routes>
      <Route path="/" element={<Layout />}>
        <Route index element={<Home />} />
        <Route path="signup" element={<Signup />} />
        <Route path="login" element={<Login />} />
        <Route path="transactions" element={<Transactions />} />
        <Route
          path="transactions/deposit"
          element={<NewTsx type={"deposit"} />}
        />
        <Route
          path="transactions/withdraw"
          element={<NewTsx type={"withdraw"} />}
        />
        <Route path="trade/:id" element={<Trade />} />
        <Route path="tradehistory" element={<TradeHistory />} />
        <Route path="trades/buy" element={<NewTrade type="buy" />} />
        <Route path="trades/sell" element={<NewTrade type="sell" />} />
      </Route>
    </Routes>
  )
}
function Layout() {
  return (
    <div>
      <Navbar />

      <Outlet />
    </div>
  )
}

export default App
