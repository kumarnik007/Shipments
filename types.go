package main

type ShipmentInfo struct {
  Sender_country   string  `json:"sender_country"`
  Receiver_country string  `json:"receiver_country"`
  Weight           float64 `json:"weight"`
  Price            string  `json:"price"`
}

type ResponseGetShipments struct {
  Shipments []ShipmentInfo `json:"shipments"`
}

// Country Name and Code as per the
// https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#:~:text=ISO%203166%2D1%20alpha%2D2%20codes%20are%20two%2Dletter,special%20areas%20of%20geographical%20interest.
type CountryInfo struct {
  Name string `json:"name"`
  Code string `json:"code"`
}

type WeightInfo struct {
  Begin float64 `json:"begin"`
  End   float64 `json:"end"`
  Price float64 `json:"price"`
}

type WeightRule struct {
  Small  WeightInfo `json:"small"`
  Medium WeightInfo `json:"medium"`
  Large  WeightInfo `json:"large"`
  Huge   WeightInfo `json:"huge"`
}

type RegionRule struct {
  Domestic      float64 `json:"domestic"`
  Eu            float64 `json:"eu"`
  International float64 `json:"international"`
}

type Pricing struct {
  WeightClass WeightRule `json:"weight"`
  Multiplier  RegionRule `json:"region"`
  Currency    string     `json:"currency"`
  IsReady bool
}

const (
  CONTENT_TYPE     = "Content-type"
  APPLICATION_JSON = "application/json"

  SHIPMENT_API_ENDPOINT  = "/shipment"
  SHIPMENTS_API_ENDPOINT = "/shipments"

  API_CORE      = "[API_CORE]"
  SWEDISH_KRONA = "SEK"

  EU_COUNTRIES = "config/eu.json"
  STORAGE      = "storage.json"
  PRICING      = "config/pricing.json"
)
