import React from "react"
import { BrowserRouter, Route, Routes } from "react-router-dom"

import "./App.css"
import Navbar from "./components/molecules/Navbar"
import Home from "./pages/Home"
import Signup from "./pages/Signup"

function App() {
  return (
    <BrowserRouter>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/signup" element={<Signup />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
