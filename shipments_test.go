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
    Pricing{},
  )
}

func Test_ShipmentsSuccessInitialShipments(t *testing.T) {
  shipement1 := ShipmentInfo{Sender_country: "SE", Receiver_country: "SE", Weight: 1, Price: "100 SEK"}
  shipement2 := ShipmentInfo{Sender_country: "SE", Receiver_country: "DK", Weight: 1, Price: "150 SEK"}
  expected, _ := json.Marshal(ResponseGetShipments{Shipments: []ShipmentInfo{shipement1, shipement2}})
  Helper_TestApi(
    t,
    http.MethodGet,
    SHIPMENTS_API_ENDPOINT,
    "",
    nil,
    string(expected),
    http.StatusOK,
    []ShipmentInfo{shipement1, shipement2},
    Pricing{},
  )
}

func Test_Shipments405MethodNotAllowed(t *testing.T) {
  expected, _ := json.Marshal(HTTPError{Detail: "Method not allowed - Expecting GET /shipments"})
  Helper_TestApi(
    t,
    http.MethodPost,
    SHIPMENTS_API_ENDPOINT,
    "",
    nil,
    string(expected),
    http.StatusMethodNotAllowed,
    []ShipmentInfo{},
    Pricing{},
  )
}

func Test_Shipment405MethodNotAllowed(t *testing.T) {
  expected, _ := json.Marshal(HTTPError{Detail: "Method not allowed - Expecting POST /shipment"})
  Helper_TestApi(
    t,
    http.MethodGet,
    SHIPMENT_API_ENDPOINT,
    "",
    nil,
    string(expected),
    http.StatusMethodNotAllowed,
    []ShipmentInfo{},
    Pricing{},
  )
}

func Test_Shipment400NoBody(t *testing.T) {
  expected, _ := json.Marshal(HTTPError{Detail: "Bad Request : No body present"})
  Helper_TestApi(
    t,
    http.MethodPost,
    SHIPMENT_API_ENDPOINT,
    APPLICATION_JSON,
    nil,
    string(expected),
    http.StatusBadRequest,
    []ShipmentInfo{},
    Pricing{},
  )
}

func Test_Shipment400UnexpectedContentType(t *testing.T) {
  expected, _ := json.Marshal(HTTPError{Detail: "Bad Request : Content-type header is not application/json"})
  Helper_TestApi(
    t,
    http.MethodPost,
    SHIPMENT_API_ENDPOINT,
    "text/plain",
    nil,
    string(expected),
    http.StatusBadRequest,
    []ShipmentInfo{},
    Pricing{},
  )
}

func Test_Shipment400InvalidWeight(t *testing.T) {
  requestString, _ := json.Marshal(ShipmentInfo{
    Sender_country:   "SE",
    Receiver_country: "SE",
    Weight:           1001,
  })
  expected, _ := json.Marshal(
    HTTPError{
      Detail: "Bad Input -  Shipment weight value outside permissible limits [0.00 - 1000.00].",
    },
  )
  Helper_TestApi(
    t,
    http.MethodPost,
    SHIPMENT_API_ENDPOINT,
    APPLICATION_JSON,
    requestString,
    string(expected),
    http.StatusBadRequest,
    []ShipmentInfo{},
    getPricingPlan(),
  )
}

func Test_Shipment400InvalidSenderCountry(t *testing.T) {
  requestString, _ := json.Marshal(ShipmentInfo{
    Sender_country:   "SWEDEN",
    Receiver_country: "SE",
    Weight:           10,
  })
  expected, _ := json.Marshal(
    HTTPError{
      Detail: "Bad Input - Sender country code is not recognised.",
    },
  )
  Helper_TestApi(
    t,
    http.MethodPost,
    SHIPMENT_API_ENDPOINT,
    APPLICATION_JSON,
    requestString,
    string(expected),
    http.StatusBadRequest,
    []ShipmentInfo{},
    getPricingPlan(),
  )
}

func Test_Shipment400InvalidReceiverCountry(t *testing.T) {
  requestString, _ := json.Marshal(ShipmentInfo{
    Sender_country:   "SE",
    Receiver_country: "ABCDEF",
    Weight:           10,
  })
  expected, _ := json.Marshal(
    HTTPError{
      Detail: "Bad Input -  Receiver country code is not recognised.",
    },
  )
  Helper_TestApi(
    t,
    http.MethodPost,
    SHIPMENT_API_ENDPOINT,
    APPLICATION_JSON,
    requestString,
    string(expected),
    http.StatusBadRequest,
    []ShipmentInfo{},
    getPricingPlan(),
  )
}

func Test_Shipment500IssueReadingPricing(t *testing.T) {
  requestString, _ := json.Marshal(ShipmentInfo{
    Sender_country:   "SE",
    Receiver_country: "SE",
    Weight:           10,
  })
  expected, _ := json.Marshal(
    HTTPError{
      Detail: "Server error - Problem with reading pricing details.",
    },
  )
  Helper_TestApi(
    t,
    http.MethodPost,
    SHIPMENT_API_ENDPOINT,
    APPLICATION_JSON,
    requestString,
    string(expected),
    http.StatusInternalServerError,
    []ShipmentInfo{},
    Pricing{},
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
    getPricingPlan(),
  )
}

func Test_ShipmentSuccessWithinEU(t *testing.T) {
  requestString, _ := json.Marshal(ShipmentInfo{
    Sender_country:   "SE",
    Receiver_country: "DK",
    Weight:           2.5,
  })
  expected, _ := json.Marshal(ShipmentInfo{
    Sender_country:   "SE",
    Receiver_country: "DK",
    Weight:           2.5,
    Price:            "150.00 " + SWEDISH_KRONA,
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
    getPricingPlan(),
  )
}

func Test_ShipmentSuccessInternational(t *testing.T) {
  requestString, _ := json.Marshal(ShipmentInfo{
    Sender_country:   "SE",
    Receiver_country: "IN",
    Weight:           2.5,
  })
  expected, _ := json.Marshal(ShipmentInfo{
    Sender_country:   "SE",
    Receiver_country: "IN",
    Weight:           2.5,
    Price:            "250.00 " + SWEDISH_KRONA,
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
    getPricingPlan(),
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
  pricing Pricing,
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
    pricing: pricing,
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
