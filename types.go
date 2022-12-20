package main

type ShipmentInfo struct {
  Sender_country   string  `json:"sender_country"`
  Receiver_country string  `json:"receiver_country"`
  Weight           float64 `json:"weight"`
  Price            string  `json:"price"`
}

const (
  CONTENT_TYPE     = "Content-type"
  APPLICATION_JSON = "application/json"

  SHIPMENT_API_ENDPOINT  = "/shipment"
  SHIPMENTS_API_ENDPOINT = "/shipments"

  API_CORE = "[API_CORE]"
)
