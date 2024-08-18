package data

import "errors"

type PriceOverView struct {
	AllPrices []PricingHour
	Cheapest  PricingHour
}

func ListEnergyPrices() (PriceOverView, error) {
	pricesApi := prices{}
	return fetchEnergy(pricesApi)
}

func fetchEnergy(pricing pricing) (PriceOverView, error) {
	prices, err := pricing.fetchPrices()
	if err != nil {
		return PriceOverView{}, errors.New("error fetching prices")
	}
	return PriceOverView{
		AllPrices: prices.Prices,
		Cheapest:  findCheapestPrice(prices),
	}, nil
}

func findCheapestPrice(prices prices) PricingHour {
	var cheapestPrice = prices.Prices[0]
	for _, price := range prices.Prices {
		if price.PriceDkk < cheapestPrice.PriceDkk {
			cheapestPrice = price
		}
	}
	return cheapestPrice
}
