interface FullPageErrorFallbackProps {
  error:
    | (Error & {
        statusCode?: number | undefined
      })
    | null
}

const FullPageErrorFallback = ({ error }: FullPageErrorFallbackProps) => {
  return (
    <div className="min-h-screen bg-gray-200 flex flex-col justify-center">
      <div className="max-w-md mx-auto">
        <img
          className="py-8"
          src={require("../images/errorMeme.jpg")}
          alt="An exploding rocket and Elon Musk with one eye brow raised"
        />
        <p>
          Uh oh... You broke the app in exciting new ways which I did not
          predict. Try refreshing the app.
        </p>
        <div
          className="p-4 mb-4 text-sm text-red-700 bg-red-100 rounded-lg dark:bg-red-200 dark:text-red-800 my-8"
          role="alert"
        >
          <span className="font-medium">Error Message:</span> {error?.message}
        </div>
        <pre></pre>
      </div>
    </div>
  )
}

export default FullPageErrorFallback
