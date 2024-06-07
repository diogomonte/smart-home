package device

import (
	"database/sql"
)

type Device struct {
	Id         int
	DeviceId   string
	DeviceType string
}

func ListDevices(db *sql.DB) ([]Device, error) {
	rows, err := db.Query("SELECT id, device_id, device_type FROM devices")
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
