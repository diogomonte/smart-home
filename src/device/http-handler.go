package device

import (
	"database/sql"
	"encoding/json"
	device "github.com/montediogo/home/src/device/registry"
	"github.com/montediogo/home/src/mqtt"
	"log"
	"net/http"
)

type Api struct {
	MqttClient mqtt.Connection
	Db         *sql.DB
}

func (api *Api) InitializeAPI() {
	http.HandleFunc("GET /devices", api.getDevices)
	http.HandleFunc("POST /devices/{deviceId}/action", api.sendMessageToDevice)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal("error initializing http server")
	}
}

func (api *Api) getDevices(w http.ResponseWriter, r *http.Request) {
	allDevices, err := device.ListDevices(api.Db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(allDevices)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *Api) sendMessageToDevice(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	if message == "" {
		http.Error(w, "Missing 'message' query parameter", http.StatusBadRequest)
		return
	}

	deviceId := r.PathValue("deviceId")

	err := api.MqttClient.Publish("home/device/"+deviceId+"/action", message)
	if err != nil {
		http.Error(w, "Error publishing message to device", http.StatusInternalServerError)
		return
	}
}
