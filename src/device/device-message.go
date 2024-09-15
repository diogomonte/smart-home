package device

import "encoding/json"

type DeviceType string

const (
	DeviceTypePlant = "plant"
)

type MessageHeader struct {
	MessageId  string `json:"message_id"`
	DeviceType string `json:"device_type"`
	DeviceId   string `json:"device_id"`
}

type Message struct {
	Header MessageHeader `json:"header"`
	Body   interface{}   `json:"body"`
}

type ConnectedMessageBody struct {
	IP string `json:"ip"`
}

type TemperatureMessageBody struct {
	Temperature float32 `json:"temperature"`
	SunLight    float32 `json:"sun_light"`
}

type PlantMessageBody struct {
	SoilMoisture float32 `json:"moisture"`
}

func ParseMessage(message string) (Message, error) {
	msg := Message{}
	err := json.Unmarshal([]byte(message), &msg)
	return msg, err
}
