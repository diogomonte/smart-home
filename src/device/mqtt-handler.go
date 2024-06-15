package device

import (
	"database/sql"
	"fmt"
	device "github.com/montediogo/home/src/device/registry"
	"github.com/montediogo/home/src/mqtt"
	"log"
)

type MqttHandler struct {
	Connection mqtt.Connection
	Db         *sql.DB
}

func (mqttHandler *MqttHandler) InitializeMqttHandler() {
	mqttHandler.Connection.Subscribe("/device/+/event", 0, mqttHandler.handleEventMessage)
	mqttHandler.Connection.Subscribe("/device/+/connected", 0, mqttHandler.handleConnectedEvent)
}

func (mqttHandler *MqttHandler) handleEventMessage(topic, message string) {
	fmt.Printf("message received on topic: %s, %s \n", topic, message)
}

func (mqttHandler *MqttHandler) handleConnectedEvent(topic, message string) {
	fmt.Printf("message received on topic: %s, %s \n", topic, message)
	deviceMessage, err := ParseMessage(message)
	if err != nil {
		log.Fatal("error parsing connected message", err)
		return
	}

	savedDevice, err := device.FindDevice(mqttHandler.Db, deviceMessage.Header.DeviceId)
	if err != nil {
		log.Fatal("error finding device", err)
		return
	}

	if savedDevice == nil {
		newDevice := device.Device{
			DeviceId:   deviceMessage.Header.DeviceId,
			DeviceType: deviceMessage.Header.DeviceType,
			Status:     device.Online,
		}
		err := device.CreateDevice(mqttHandler.Db, newDevice)
		if err != nil {
			log.Fatal("error saving new connected device", err)
			return
		}
	}
}
