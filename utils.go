package main

import (
  "log"
)

func contains(container []string, value string) bool {
  for _, v := range container {
    if v == value {
      return true
    }
  }

  return false
}

// Sets the logger flag to include file name and number in the logs
// to be printed.
func setLoggerFlag() {
  log.SetFlags(log.Lshortfile)
}
