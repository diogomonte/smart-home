package device

type DeviceMessageHeader struct {
	MessageId  string `json:"message_id"`
	DeviceType string `json:"device_type"`
}

type DeviceMessage struct {
	Header DeviceMessageHeader `json:"header"`
	Body   interface{}         `json:"body"`
}
