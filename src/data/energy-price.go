package data

import (
	"encoding/json"
	"net/http"
)

type pricing interface {
	fetchPrices() ([]PricingHourResponse, error)
}

type prices struct {
	Prices []pricingHour `json:"records"`
}

type pricingHour struct {
	Time     string  `json:"HourDK"`
	Area     string  `json:"PriceArea"`
	PriceDkk float32 `json:"SpotPriceDKK"`
	PriceEur float32 `json:"SpotPriceEUR"`
}

type PricesResponse struct {
	Prices []PricingHourResponse `json:"records"`
}

type PricingHourResponse struct {
	Time     string  `json:"time"`
	Area     string  `json:"area"`
	PriceDkk float32 `json:"price_dkk"`
	PriceEur float32 `json:"price_eur"`
}

func (p PricesResponse) fetchPrices() ([]PricingHourResponse, error) {
	resp, err := http.Get("https://api.energidataservice.dk/dataset/Elspotprices?limit=24&filter={%22PriceArea%22:[%22DK2%22]}&sort=HourUTC%20DESC&timezone=dk}")
	if err != nil {
		return []PricingHourResponse{}, err
	}

	priceOverview := prices{}
	err = json.NewDecoder(resp.Body).Decode(&priceOverview)
	if err != nil {
		return []PricingHourResponse{}, err
	}

	var pricesArray []PricingHourResponse
	for _, price := range priceOverview.Prices {
		pricesArray = append(pricesArray, PricingHourResponse{
			Time:     price.Time,
			Area:     price.Area,
			PriceEur: price.PriceEur,
			PriceDkk: price.PriceDkk,
		})
	}
	return pricesArray, nil
}
