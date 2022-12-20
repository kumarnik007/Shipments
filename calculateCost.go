package main

// Apply weight Class Rules and get the base price
func getBasePrice(weight float64) float64 {
  // Weight Class : Small
  if weight < 10 {
    return 100
  }

  // Weight Class : Medium
  if weight < 25 {
    return 300
  }

  // Weight Class : Large
  if weight < 50 {
    return 500
  }

  // Weight Class : Huge
  return 2000
}

// Apply Region Rules and get price multiplier
func getPriceMultiplier(sender, receiver string) (float64, error) {
  euCountries, err := getAllEUCountryCodes()
  if err != nil {
    return 0, err
  }

  // Shipment is within EU and domestic
  if contains(euCountries, sender) && contains(euCountries, receiver) && sender == receiver {
    return 1, nil
  }

  // Shipment is within the EU
  if contains(euCountries, sender) && contains(euCountries, receiver) {
    return 1.5, nil
  }

  // Shipment has either sender or/and receiver country as non-EU member
  return 2.5, nil
}

func calculatePrice(sender, receiver string, weight float64) (float64, error) {
  basePrice := getBasePrice(weight)
  priceMultiplier, errorResponse := getPriceMultiplier(sender, receiver)

  return (priceMultiplier * basePrice), errorResponse
}
