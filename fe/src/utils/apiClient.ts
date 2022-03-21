import * as auth from "../AuthProvider"

const apiURL = "http://localhost:8000/api"

interface Config {
  data?: Record<string, any> | [] | FormData
  headers?: HeadersInit
  method?: string
  fileUpload?: boolean
  fileDownload?: boolean
}

export const client = async <ApiResponse,>(
  endpoint: string,
  {
    data,
    headers: customHeaders,
    method = "GET",
    fileUpload = false,
    fileDownload = false,
    ...customConfig
  }: Config
): Promise<ApiResponse> => {
  let config: RequestInit | undefined | FormData | Record<string, any>
  if (fileUpload === false) {
    config = {
      fileDownload: fileDownload,
      method: method,
      body: data ? JSON.stringify(data) : undefined,
      headers: {
        "Content-Type": data ? "application/json" : "",
        ...customHeaders,
      },
      credentials: "include",
      ...customConfig,
    }
  } else {
    config = {
      fileDownload: fileDownload,
      method: method,
      body: data,
      credentials: "include"
    }
  }

  return fetchRetry<ApiResponse>(`${apiURL}/${endpoint}`, config, 2)
}

const fetchRetry = async <ApiResponse,>(
  url: string,
  config: Config,
  retries: number = 2
): Promise<ApiResponse> => {
  return fetch(url, config).then(async (response) => {
    if (response.status === 401) {
      if (retries === 0) {
        await auth.logout()
        // refresh the page for them
        window.location.assign(window.location.href)
        return Promise.reject({
          message: "Please re-authenticate.",
          statusCode: response.status,
        })
      }
    }

    if (response.status === 204) {
      return Promise.resolve({}) as Promise<ApiResponse>
    }

    let data

    if (config.fileDownload) {
      const blobData = await response.blob()
      data = window.URL.createObjectURL(new Blob([blobData]))
    } else {
      data = await response.json()
    }

    if (response.ok) {
      return data as Promise<ApiResponse>
    } else {
      return Promise.reject({
        message: data.message,
        statusCode: response.status,
      })
    }
  })
}
