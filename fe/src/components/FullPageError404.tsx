import React from "react"

const FullPageErrorFallback = () => {
  return (
    <div className="min-h-screen bg-gray-200 flex flex-col justify-center">
      <div className="max-w-md mx-auto">
        <img
          className="py-8"
          src={require("../images/404Error.jpg")}
          alt="A car flying around the earth in space"
        />
        <p>Uh oh... You took the wrong turn!.</p>
        <div
          className="p-4 mb-4 text-sm text-red-700 bg-red-100 rounded-lg dark:bg-red-200 dark:text-red-800 my-8"
          role="alert"
        >
          <span className="font-medium text-center">Error 404</span>
        </div>
        <pre></pre>
      </div>
    </div>
  )
}

export default FullPageErrorFallback
