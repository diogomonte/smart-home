package device

import (
	"database/sql"
)

type Device struct {
	Id         int
	DeviceId   string
	DeviceType string
	Status     string
}

func CreateDevice(db *sql.DB, device Device) error {
	_, err := db.Exec("NSERT INTO devices (device_id, device_type, status) VALUES ($1, $2, $3);", device.DeviceId, device.DeviceType, device.Status)
	if err != nil {
		return err
	}
	return nil
}

func ListDevices(db *sql.DB) ([]Device, error) {
	rows, err := db.Query("SELECT id, device_id, device_type, status FROM devices")
	if err != nil {
		return nil, err
	}

	var devices []Device

	for rows.Next() {
		var id int
		var deviceId, deviceType string
		err := rows.Scan(&id, &deviceId, &deviceType)
		if err != nil {
			return nil, err
		}
		device := Device{
			Id:         id,
			DeviceId:   deviceId,
			DeviceType: deviceType,
		}
		devices = append(devices, device)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return devices, nil
}
