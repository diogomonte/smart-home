package data

import (
	"encoding/json"
	"net/http"
)

type pricing interface {
	fetchPrices() (prices, error)
}

type prices struct {
	Prices []PricingHour `json:"records"`
}

type PricingHour struct {
	Time     string  `json:"HourDK"`
	Area     string  `json:"PriceArea"`
	PriceDkk float32 `json:"SpotPriceDKK"`
	PriceEur float32 `json:"SpotPriceEUR"`
}

func (p prices) fetchPrices() (prices, error) {
	resp, err := http.Get("https://api.energidataservice.dk/dataset/Elspotprices?limit=24&filter={%22PriceArea%22:[%22DK2%22]}&sort=HourUTC%20DESC&timezone=dk}")
	if err != nil {
		return prices{}, err
	}

	priceOverview := prices{}
	err = json.NewDecoder(resp.Body).Decode(&priceOverview)
	if err != nil {
		return prices{}, err
	}

	return priceOverview, nil
}
