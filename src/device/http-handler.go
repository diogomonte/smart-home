package device

import (
	"database/sql"
	"encoding/json"
	"fmt"
	device "github.com/montediogo/home/src/device/registry"
	"github.com/montediogo/home/src/mqtt"
	"net/http"
)

type Api struct {
	MqttClient mqtt.Connection
	Db         *sql.DB
}

func (api *Api) InitializeAPI() {
	http.HandleFunc("GET /devices", api.getDevices)
	http.HandleFunc("GET /devices/{deviceId}/events", api.deviceEvents)
	http.HandleFunc("POST /devices/{deviceId}/action", api.sendMessageToDevice)
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

func (api *Api) deviceEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	deviceId := r.PathValue("deviceId")
	c := make(chan string)
	go api.MqttClient.Subscribe(fmt.Sprintf("/device/%s/event", deviceId), 0, func(topic string, message string) {
		c <- message
	})
	for {
		select {
		case deviceMessage := <-c:
			fmt.Fprintf(w, deviceMessage)
			w.(http.Flusher).Flush()
		case <-r.Context().Done():
			return
		}
	}
}
