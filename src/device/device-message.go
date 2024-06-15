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

func ParseMessage(message string) (DeviceMessage, error) {
	msg := DeviceMessage{}
	err := json.Unmarshal([]byte(message), &msg)
	return msg, err
}
