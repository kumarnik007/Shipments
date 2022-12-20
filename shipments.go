// ------------------------------------
// Handler Functions for API
// ------------------------------------

package main

import (
  "encoding/json"
  "io"
  "net/http"
  "strconv"

  // https://pkg.go.dev/github.com/biter777/countries#pkg-overview
  "github.com/biter777/countries"
)

type Shipment struct {
  all []ShipmentInfo
}

func (shipment *Shipment) handleAPI(w http.ResponseWriter, r *http.Request) error {
  if r.URL.Path == SHIPMENTS_API_ENDPOINT {
    if r.Method == http.MethodGet {
      return shipment.getAll(w, r)
    }

    return NewHTTPError(nil, 405, "Method not allowed - Expecting GET /shipments")
  }

  if r.URL.Path == SHIPMENT_API_ENDPOINT && r.Method == http.MethodPost {
    return shipment.post(w, r)
  }

  return NewHTTPError(nil, 405, "Method not allowed - Expecting POST /shipment")
}

func (shipment *Shipment) getAll(w http.ResponseWriter, r *http.Request) error {
  shipments := &ResponseGetShipments{
    Shipments: shipment.all,
  }

  getAllResponse, errorResponse := json.Marshal(shipments)
  if errorResponse != nil {
    // Return 500 Internal Server Error.
    return NewHTTPError(errorResponse, 500, "unable to prepare JSON response")
  }

  w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
  w.Write(getAllResponse)

  return nil
}

func (shipment *Shipment) post(w http.ResponseWriter, r *http.Request) error {
  contentType := r.Header.Get(CONTENT_TYPE)
  if contentType != APPLICATION_JSON {
    // Return 400 Bad Request.
    return NewHTTPError(nil, 400, "Bad Request : " + CONTENT_TYPE + " header is not " + APPLICATION_JSON)
  }

  if r.ContentLength == 0 {
    // Return 400 Bad Request.
    return NewHTTPError(nil, 400, "Bad Request : No body present")
  }

  // Read body
  postBody, errorResponse := io.ReadAll(r.Body)
  if errorResponse != nil {
    // Return 500 Internal Server Error.
    return NewHTTPError(errorResponse, 500, "request body read error")
  }

  // Unmarshal
  var msg ShipmentInfo
  errorResponse = json.Unmarshal(postBody, &msg)
  if errorResponse != nil {
    // Return 400 Bad Request.
    return NewHTTPError(
      errorResponse,
      400,
      "Bad Request : JSON input has invalid value",
    )
  }

  errorResponse = validatePostShipmentInputParams(
    msg.Sender_country,
    msg.Receiver_country,
    msg.Weight,
  )
  if errorResponse != nil {
    return errorResponse
  }

  price, errorResponse := calculatePrice(
    msg.Sender_country,
    msg.Receiver_country,
    msg.Weight,
  )
  if errorResponse != nil {
    return errorResponse
  }

  shipmentInfo := &ShipmentInfo{
    Sender_country:   msg.Sender_country,
    Receiver_country: msg.Receiver_country,
    Weight:           msg.Weight,
    Price:            strconv.FormatFloat(price, 'f', 2, 64) + " " + SWEDISH_KRONA,
  }

  shipment.all = append(shipment.all, *shipmentInfo)
  // Store updated shipments to the storage.
  persitShipments(shipment.all)

  postResponse, errorResponse := json.Marshal(shipmentInfo)
  if errorResponse != nil {
    // Return 500 Internal Server Error.
    return NewHTTPError(errorResponse, 500, "unable to prepare JSON response")
  }

  w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
  w.Write(postResponse)

  return nil
}

func validatePostShipmentInputParams(sender, receiver string, weight float64) error {
  var allCountryCodes []string

  allCountryCodes = make([]string, len(countries.All()))
  for _, country := range countries.AllInfo() {
    allCountryCodes = append(allCountryCodes, country.Alpha2)
  }

  errString := ""

  if !contains(allCountryCodes, sender) {
    errString = "Sender country code is not recognised."
  }

  if !contains(allCountryCodes, receiver) {
    errString += " Receiver country code is not recognised."
  }

  if weight < 0 || weight > 1000 {
    errString += " Shipment weight value outside permissible limits [0 - 1000]."
  }

  if errString == "" {
    return nil
  }

  return NewHTTPError(nil, 400, "Bad Input - "+errString)
}
