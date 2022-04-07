import { Navigate, Outlet, Route, Routes, useLocation } from "react-router-dom"
import "./App.css"
import FullPageError404 from "./components/FullPageError404"
import FullPageErrorFallback from "./components/FullPageErrorFallback"
import Navbar from "./components/Navbar"
import { useAuth } from "./context/AuthContext"
import Home from "./pages/Home"
import Login from "./pages/Login"
import NewTrade from "./pages/NewTrade"
import NewTsx from "./pages/NewTsx"
import Signup from "./pages/Signup"
import Trade from "./pages/Trade"
import TradeHistory from "./pages/Trades"
import Transactions from "./pages/Transactions"


function App() {
  return (
    <Routes>
      <Route path="/" element={<Layout />}>

        {/* Public Routes */}
        <Route index element={<Home />} />
        <Route path="signup" element={<Signup />} />
        <Route path="login" element={<Login />} />

        {/* Private Routes */}
        <Route path="transactions" element={<RequireAuth><Transactions /></RequireAuth>} />
        <Route
          path="transactions/deposit"
          element={<RequireAuth><NewTsx type={"deposit"} /></RequireAuth>}
        />
        <Route
          path="transactions/withdraw"
          element={<RequireAuth><NewTsx type={"withdraw"} /></RequireAuth>}
        />
        <Route path="trade/:id" element={<RequireAuth><Trade/></RequireAuth>} />
        <Route path="tradehistory" element={<RequireAuth><TradeHistory/></RequireAuth>} />
        <Route path="trades/buy" element={<RequireAuth><NewTrade type="buy" /></RequireAuth>} />
        <Route path="trades/sell" element={<RequireAuth><NewTrade type="sell" /></RequireAuth>} />
        
        {/* Error Routes */}
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

function RequireAuth({ children }: { children: JSX.Element }) {
  const { user } = useAuth()
  const location = useLocation();

  if (!user) {
    // Redirect them to the /login page, but save the current location they were
    // trying to go to when they were redirected. This allows us to send them
    // along to that page after they login, which is a nicer user experience
    // than dropping them off on the home page.
    return <Navigate to="/login" state={{ from: location }} replace />;
  }

  return children;
}

export default App
