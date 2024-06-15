package device

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
