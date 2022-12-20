// ------------------------------------
// Handler Functions for API
// ------------------------------------

package main

import (
  "encoding/json"
  "net/http"
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
  return nil
}
