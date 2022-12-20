// ------------------------------------
// Handler Functions for API
// ------------------------------------

package main

import (
  "net/http"
)

type Shipment struct {}

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
  return nil
}

func (shipment *Shipment) post(w http.ResponseWriter, r *http.Request) error {
  return nil
}
