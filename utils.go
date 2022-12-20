package main

import (
  "encoding/json"
  "io"
  "log"
  "os"
)

func readJSON(fileName string) ([]byte, error) {
  // Open json file
  jsonFile, err := os.Open(fileName)
  // if os.Open returns an error
  if err != nil {
    return []byte{}, NewHTTPError(nil, 500, "Server error - "+err.Error())
  }
  // defer the closing of json file
  defer jsonFile.Close()

  return io.ReadAll(jsonFile)
}

func contains(container []string, value string) bool {
  for _, v := range container {
    if v == value {
      return true
    }
  }

  return false
}

func getAllEUCountryCodes() ([]string, error) {
  byteValue, err := readJSON(EU_COUNTRIES)
  if err != nil {
    return []string{}, NewHTTPError(nil, 500, "Server error - "+err.Error())
  }

  var euCountries []CountryInfo
  err = json.Unmarshal([]byte(byteValue), &euCountries)
  if err != nil {
    return []string{}, NewHTTPError(nil, 500, "Server error - "+err.Error())
  }

  euCountryCodes := []string{}
  for _, country := range euCountries {
    euCountryCodes = append(euCountryCodes, country.Code)
  }

  return euCountryCodes, nil
}

func retrieveShipments() ([]ShipmentInfo) {
  byteValue, err := readJSON(STORAGE)
  if err != nil {
    return []ShipmentInfo{}
  }

  var shipments []ShipmentInfo
  _ = json.Unmarshal([]byte(byteValue), &shipments)

  return shipments
}

func persitShipments(shipments []ShipmentInfo) {
  jsonFile, err := os.Create(STORAGE)
  // if os.Create returns an error
  if err != nil {
    return
  }

  defer jsonFile.Close()

  encoder := json.NewEncoder(jsonFile)
  _ = encoder.Encode(shipments)
}

// Sets the logger flag to include file name and number in the logs
// to be printed.
func setLoggerFlag() {
  log.SetFlags(log.Lshortfile)
}
