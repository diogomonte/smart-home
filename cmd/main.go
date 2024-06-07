package main

import (
	"embed"
	"github.com/montediogo/home/src/db"
	"github.com/montediogo/home/src/device"
	"github.com/montediogo/home/src/migrator"
	"github.com/montediogo/home/src/mqtt"
	"log"
)

var MigrationsFS embed.FS

func main() {
	databaseConnection := db.Connect("mysql", "user:password@tcp(localhost:3306)/home-automation?multiStatements=true")
	mqttConnection, err := mqtt.Connect("service", "tcp://localhost:1883")
	if err != nil {
		log.Fatal("error connecting to mqtt broker", err)
	}

	m := migrator.Migrator{}
	err = m.RunMigrations(databaseConnection)
	if err != nil {
		panic(err)
	}

	deviceMqtt := device.MqttHandler{
		Connection: mqttConnection,
	}
	deviceMqtt.InitializeMqttHandler()

	deviceHttp := device.Api{
		MqttClient: mqttConnection,
		Db:         databaseConnection,
	}
	deviceHttp.InitializeAPI()
}
