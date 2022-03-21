import React, { useState } from "react"

//TODO: Refactor into components

function Navbar() {
  const [isOpen, setIsOpen] = useState(false)

  const MobileNavigationMenuItem: React.FC<{
    to: string
    toggleMenu?: () => void
  }> = ({ children, to, toggleMenu }) => (
    <a
      href={to}
      onClick={toggleMenu}
      className="block py-2 px-4 text-sm hover:bg-gray-700 text-gray-200"
    >
      {children}
    </a>
  )

  const MobileNavigationHamburgerButton: React.FC<{
    toggleMenu: () => void
  }> = ({ toggleMenu }) => (
    <button onClick={toggleMenu}>
      <svg
        xmlns="http://www.w3.org/2000/svg"
        className="h-6 w-6 text-white"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
        stroke-width="2"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M4 6h16M4 12h16M4 18h16"
        />
      </svg>
    </button>
  )

  const NavigationMenuItem: React.FC<{
    to: string
  }> = ({ children, to }) => (
    <a href={to} className="py-5 px-3 text-gray-200 hover:text-green-600">
      {children}
    </a>
  )

  const Logo = () => (
    <div>
      <a href="#" className="flex items-center py-5 px-2 text-gray-400">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          className="h-6 w-6 mr-1 text-green-600"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
          stroke-width="2"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
          />
        </svg>
        <span className="font-bold text-green-600">Trading App</span>
      </a>
    </div>
  )

  return (
    <>
      <nav className="bg-black">
        <div className="max-w-5xl mx-auto px-8">
          <div className="flex justify-between">
            <div className="flex space-x-4">
              {/* Logo */}
              <Logo />
              {/* primary nav */}
              <div className="hidden md:flex items-center space-x-1 ">
                <NavigationMenuItem to="#">Transactions</NavigationMenuItem>
                <NavigationMenuItem to="#">Trades</NavigationMenuItem>
              </div>
            </div>
            {/* secondary nav */}
            <div className="hidden md:flex items-center space-x-1">
              <NavigationMenuItem to="#">Login</NavigationMenuItem>

              {/* TODO: Abstract into button component */}
              <a
                href="#"
                className="py-2 px-3 bg-primary text-black font-bold rounded hover:bg-green-600 transition duration-300"
              >
                Signup
              </a>
            </div>
            {/* Mobile button */}
            <div className="md:hidden flex items-center">
              <MobileNavigationHamburgerButton
                toggleMenu={() => setIsOpen(!isOpen)}
              />
            </div>
          </div>
        </div>
        {/* Mobile menu */}
        <div className={isOpen == false ? "hidden" : "md:hidden text-center"}>
          <MobileNavigationMenuItem to="#" toggleMenu={() => setIsOpen(false)}>
            Transactions
          </MobileNavigationMenuItem>
          <MobileNavigationMenuItem to="#" toggleMenu={() => setIsOpen(false)}>
            Trades
          </MobileNavigationMenuItem>
        </div>
      </nav>
    </>
  )
}

export default Navbar
