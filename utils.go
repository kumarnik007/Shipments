package main

import (
  "log"
)

// Sets the logger flag to include file name and number in the logs
// to be printed.
func setLoggerFlag() {
  log.SetFlags(log.Lshortfile)
}
