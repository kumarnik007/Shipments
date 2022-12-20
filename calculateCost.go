package main

// Apply weight Class Rules and get the base price
func getBasePrice(weight float64, rule WeightRule) float64 {
  // Weight Class : Small
  if weight < rule.Small.End {
    return rule.Small.Price
  }

  // Weight Class : Medium
  if weight < rule.Medium.End {
    return rule.Medium.Price
  }

  // Weight Class : Large
  if weight < rule.Large.End {
    return rule.Large.Price
  }

  // Weight Class : Huge
  return rule.Huge.Price
}

// Apply Region Rules and get price multiplier
func getPriceMultiplier(sender, receiver string, multiplier RegionRule) (float64, error) {
  euCountries, err := getAllEUCountryCodes()
  if err != nil {
    return 0, err
  }

  // Shipment is within EU and domestic
  if contains(euCountries, sender) && contains(euCountries, receiver) && sender == receiver {
    return multiplier.Domestic, nil
  }

  // Shipment is within the EU
  if contains(euCountries, sender) && contains(euCountries, receiver) {
    return multiplier.Eu, nil
  }

  // Shipment has either sender or/and receiver country as non-EU member
  return multiplier.International, nil
}

func calculatePrice(
  sender, receiver string,
  weight float64,
  pricing Pricing,
) (float64, error) {
  basePrice := getBasePrice(weight, pricing.WeightClass)
  priceMultiplier, errorResponse := getPriceMultiplier(
    sender,
    receiver,
    pricing.Multiplier,
  )

  return (priceMultiplier * basePrice), errorResponse
}
