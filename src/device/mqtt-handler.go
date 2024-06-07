package device

import (
	"fmt"
	"github.com/montediogo/home/src/mqtt"
)

type MqttHandler struct {
	Connection mqtt.Connection
}

func (mqttHandler *MqttHandler) InitializeMqttHandler() {
	mqttHandler.Connection.Subscribe("/device/+/event", 0, mqttHandler.handleEventMessage)
}

func (mqttHandler *MqttHandler) handleEventMessage(topic, message string) {
	fmt.Printf("message received on topic: %s, %s \n", topic, message)
}
