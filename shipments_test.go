// handlers_test.go
package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "net/http"
  "net/http/httptest"
  "os"
  "testing"
)

func Test_ShipmentsSuccessEmptySlice(t *testing.T) {
  expected, _ := json.Marshal(ResponseGetShipments{Shipments: []ShipmentInfo{}})
  Helper_TestApi(
    t,
    http.MethodGet,
    SHIPMENTS_API_ENDPOINT,
    "",
    nil,
    string(expected),
    http.StatusOK,
    []ShipmentInfo{},
  )
}

func Test_ShipmentSuccessDomesticWithinEU(t *testing.T) {
  requestString, _ := json.Marshal(ShipmentInfo{
    Sender_country:   "SE",
    Receiver_country: "SE",
    Weight:           2.5,
  })
  expected, _ := json.Marshal(ShipmentInfo{
    Sender_country:   "SE",
    Receiver_country: "SE",
    Weight:           2.5,
    Price:            "100.00 " + SWEDISH_KRONA,
  })
  Helper_TestApi(
    t,
    http.MethodPost,
    SHIPMENT_API_ENDPOINT,
    APPLICATION_JSON,
    requestString,
    string(expected),
    http.StatusOK,
    []ShipmentInfo{},
  )
}

func Helper_TestApi(
  t *testing.T,
  apiType string,
  apiName string,
  contentType string,
  requestString []byte,
  expected string,
  statusCode int,
  initialShipments []ShipmentInfo,
) {
  req, err := http.NewRequest(apiType, apiName, bytes.NewBuffer(requestString))
  if err != nil {
    t.Fatal(err)
  }

  req.Header.Set(CONTENT_TYPE, contentType)
  // Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
  rr := httptest.NewRecorder()

  shipment := Shipment{
    all: initialShipments,
  }

  if apiName == SHIPMENT_API_ENDPOINT || apiName == SHIPMENTS_API_ENDPOINT {
    handler := http.Handler(ApiHandler(shipment.handleAPI))
    handler.ServeHTTP(rr, req)
  } else {
    fmt.Println("-- Invalid API Name --")
    return
  }

  // Check the status code is what we expect.
  if status := rr.Code; status != statusCode {
    t.Errorf("handler returned wrong status code: got %v want %v",
      status, statusCode)
  }

  // Check the response body is what we expect.
  if rr.Body.String() != expected {
    t.Errorf("handler returned unexpected body: got %v want %v",
      rr.Body.String(), expected)
  }
}

func TestMain(m *testing.M) {
  fmt.Println("------------------------ BEGIN TestMain -----------------------")
  setLoggerFlag()
  statusCode := m.Run()
  fmt.Println("---------------------------------------------------------------")
  if statusCode != 0 {
    fmt.Println("-------------- A Test Case Failed Ending TestMain -------------")
  } else {
    fmt.Println("----------- All Test Cases Passed Ending TestMain -------------")
  }
  fmt.Println("---------------------------------------------------------------")
  // fmt.Println("------------------------- END TestMain ------------------------")
  os.Exit(statusCode)
}
