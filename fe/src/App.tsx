import { Outlet, Route, Routes } from "react-router-dom"
import "./App.css"
import FullPageError404 from "./components/FullPageError404"
import FullPageErrorFallback from "./components/FullPageErrorFallback"
import Navbar from "./components/Navbar"
import Home from "./pages/Home"
import Login from "./pages/Login"
import NewTrade from "./pages/NewTrade"
import NewTsx from "./pages/NewTsx"
import Signup from "./pages/Signup"
import Trade from "./pages/Trade"
import TradeHistory from "./pages/Trades"
import Transactions from "./pages/Transactions"

//TODO:
// - Handle the request for unauthorized routes

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
        <Route path="*" element={<FullPageError404 />} />
        <Route
          path="error"
          element={
            <FullPageErrorFallback error={Error("Test error message")} />
          }
        />
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
