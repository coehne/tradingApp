import React from "react"

interface FullPageErrorFallbackProps {
  error:
    | (Error & {
        statusCode?: number | undefined
      })
    | null
}

const FullPageErrorFallback = ({ error }: FullPageErrorFallbackProps) => {
  return (
    <div>
      <p>Uh oh... There's a problem. Try refreshing the app.</p>
      <pre>{error?.message}</pre>
    </div>
  )
}

export default FullPageErrorFallback
