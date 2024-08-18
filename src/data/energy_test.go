package data

import "testing"

type MockPrices struct {
}

func (p MockPrices) fetchPrices() ([]PricingHourResponse, error) {
	return []PricingHourResponse{
		{
			Time:     "2024-04-28T23:00:00",
			Area:     "DK2",
			PriceDkk: 370.720001,
			PriceEur: 49.709999,
		},
		{
			Time:     "2024-04-27T22:00:00",
			Area:     "DK2",
			PriceDkk: 306.510010,
			PriceEur: 41.099998,
		},
		{
			Time:     "2024-04-26T22:00:00",
			Area:     "DK2",
			PriceDkk: 304.510010,
			PriceEur: 40.099998,
		},
	}, nil
}

func TestFetchEnergy(t *testing.T) {
	mockedPrices := MockPrices{}
	overview, _ := fetchEnergy(mockedPrices)

	if len(overview.AllPrices) != 3 {
		t.Error("expects 3 elements to be returned")
	}

	if overview.Cheapest.PriceDkk <= 0 {
		t.Error("expects the data to be greater that zero")
	}
}

func TestFindCheapestPrice(t *testing.T) {
	mockedPrices := MockPrices{}
	prices, _ := mockedPrices.fetchPrices()
	cheapest := findCheapestPrice(prices)

	if cheapest.PriceDkk != 304.510010 {
		t.Error("expects the cheapest price form the list")
	}
}
