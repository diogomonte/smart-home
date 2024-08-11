package device

import (
	"database/sql"
)

type DeviceStatus string

const (
	Online  DeviceStatus = "online"
	Offline DeviceStatus = "offline"
	Unknown DeviceStatus = "unknown"
)

type Device struct {
	Id         int          `json:"id"`
	DeviceId   string       `json:"deviceId"`
	DeviceType string       `json:"deviceType"`
	Status     DeviceStatus `json:"status"`
}

func CreateDevice(db *sql.DB, device Device) error {
	_, err := db.Exec("INSERT INTO devices (device_id, device_type, status) VALUES (?, ?, ?);", device.DeviceId, device.DeviceType, device.Status)
	if err != nil {
		return err
	}
	return nil
}

func FindDevice(db *sql.DB, deviceId string) (*Device, error) {
	rows, err := db.Query("SELECT id, device_id, device_type, status from devices where device_id = ?;", deviceId)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		device, err := rowToDevice(rows)
		if err != nil {
			return nil, err
		}
		return &device, nil
	}
	return nil, nil
}

func ListDevices(db *sql.DB) ([]Device, error) {
	rows, err := db.Query("SELECT id, device_id, device_type, status FROM devices")
	if err != nil {
		return nil, err
	}

	var devices []Device

	for rows.Next() {
		device, err := rowToDevice(rows)
		if err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return devices, nil
}

func rowToDevice(rows *sql.Rows) (Device, error) {
	var id int
	var deviceId, deviceType string
	var status DeviceStatus
	err := rows.Scan(&id, &deviceId, &deviceType, &status)
	if err != nil {
		return Device{}, err
	}
	device := Device{
		Id:         id,
		DeviceId:   deviceId,
		DeviceType: deviceType,
		Status:     status,
	}
	return device, nil
}
