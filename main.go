package main

import (
  "log"
  "net/http"
  "strconv"
)

const PORT_NUMBER = 8081

func main() {
  setLoggerFlag()
  addr := ":" + strconv.Itoa(PORT_NUMBER)
  shipment := Shipment{
    // Retrieve all shipments from storage.
    all: retrieveShipments(),
  }

  http.Handle(SHIPMENTS_API_ENDPOINT, ApiHandler(shipment.handleAPI))
  http.Handle(SHIPMENT_API_ENDPOINT, ApiHandler(shipment.handleAPI))

  log.Println("== Welcome to my Web Server ==")
  log.Println("== Server Is Listening On Port", addr, " ==")
  log.Fatal(http.ListenAndServe(addr, nil))
}
