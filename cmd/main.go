package main

import (
	"fmt"
	"github.com/montediogo/home/src/db"
	"github.com/montediogo/home/src/db/migrator"
	"github.com/montediogo/home/src/device"
	"github.com/montediogo/home/src/mqtt"
	"log"
	"os"
)

func main() {
	fmt.Println("Initializing home automation app")
	dataSource := fmt.Sprintf("user:password@tcp(%s:3306)/home-automation?multiStatements=true", getEnv("DB_HOST", "localhost"))
	mqttHost := fmt.Sprintf("tcp://%s:1883", getEnv("MQTT_HOST", "localhost"))

	databaseConnection := db.Connect("mysql", dataSource)
	mqttConnection, err := mqtt.Connect("service", mqttHost)
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
		Db:         databaseConnection,
	}
	deviceMqtt.InitializeMqttHandler()

	fmt.Println("Home automation app is up and running")
	deviceHttp := device.Api{
		MqttClient: mqttConnection,
		Db:         databaseConnection,
	}
	deviceHttp.InitializeAPI()
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
