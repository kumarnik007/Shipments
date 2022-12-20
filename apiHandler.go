package main

import (
    "log"
    "net/http"
)

type ApiHandler func(http.ResponseWriter, *http.Request) error

// ApiHandler implements http.Handler interface.
func (handlerCb ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  errorResponse := handlerCb(w, r) // Call handler function
  if errorResponse == nil {
      return
  }
  
  // Error handling logic starts.
  log.Printf(API_CORE + " An error occured: %v", errorResponse)

  // Check if it is a ClientError.
  clientError, ok := errorResponse.(ClientError)
  if !ok {
      // If the error is not ClientError, assume that it is ServerError.
      // return 500 Internal Server Error.
      w.WriteHeader(500)
      return
  }

  // Try to get response body of ClientError.
  body, errorResponse := clientError.ResponseBody()
  if errorResponse != nil {
      log.Printf(API_CORE + " An error accured: %v", errorResponse)
      w.WriteHeader(500)
      return
  }
  // Get http status code and headers.
  status, headers := clientError.ResponseHeaders()
  for k, v := range headers {
      w.Header().Set(k, v)
  }
  w.WriteHeader(status)
  w.Write(body)
}
