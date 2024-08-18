package data

import "errors"

type PriceOverView struct {
	AllPrices []PricingHourResponse
	Cheapest  PricingHourResponse
}

func ListEnergyPrices() (PriceOverView, error) {
	pricesApi := PricesResponse{}
	return fetchEnergy(pricesApi)
}

func fetchEnergy(pricing pricing) (PriceOverView, error) {
	prices, err := pricing.fetchPrices()
	if err != nil {
		return PriceOverView{}, errors.New("error fetching prices")
	}
	return PriceOverView{
		AllPrices: prices,
		Cheapest:  findCheapestPrice(prices),
	}, nil
}

func findCheapestPrice(prices []PricingHourResponse) PricingHourResponse {
	var cheapestPrice = prices[0]
	for _, price := range prices {
		if price.PriceDkk < cheapestPrice.PriceDkk {
			cheapestPrice = price
		}
	}
	return cheapestPrice
}
