package device

import "encoding/json"

type DeviceMessageHeader struct {
	MessageId  string `json:"message_id"`
	DeviceType string `json:"device_type"`
	DeviceId   string `json:"device_id"`
}

type DeviceMessage struct {
	Header DeviceMessageHeader `json:"header"`
	Body   interface{}         `json:"body"`
}

type ConnectedMessageBody struct {
	IP string `json:"ip"`
}

type TemperatureMessageBody struct {
	Temperature float32 `json:"temperature"`
	SunLight    float32 `json:"sun_light"`
}

func ParseMessage(message string) (DeviceMessage, error) {
	msg := DeviceMessage{}
	err := json.Unmarshal([]byte(message), &msg)
	return msg, err
}
