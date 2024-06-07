package device

type DeviceMessageHeader struct {
	MessageId string
	DeviceId  string
}

type DeviceMessage struct {
	Header DeviceMessageHeader `json:"header"`
	Body   interface{}         `json:"body"`
}
