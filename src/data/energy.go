package data

import (
	"encoding/json"
	"net/http"
)

type API struct {
}

func (api *API) InitializeAPI() {
	http.HandleFunc("GET /energy", api.listPrices)
}

func (api *API) listPrices(w http.ResponseWriter, r *http.Request) {
	prices, err := ListEnergyPrices()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(prices)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
