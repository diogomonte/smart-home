package mqtt

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"os"
	"testing"
	"time"
)

var connection Connection

func TestMain(m *testing.M) {
	ctx := context.Background()

	mqttContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "eclipse-mosquitto",
			ExposedPorts: []string{"1883/tcp"},
			WaitingFor:   wait.ForListeningPort("1883/tcp").WithStartupTimeout(45 * time.Second),
			Files: []testcontainers.ContainerFile{
				{
					HostFilePath:      "/Users/diogo/.go/src/github.com/montediogo/home-automation/mosquitto/config/mosquitto.conf",
					ContainerFilePath: "/mosquitto/config/mosquitto.conf",
					FileMode:          0644,
				},
			},
		},
		Started: true,
	})

	if err != nil {
		log.Fatalf("Failed to start container: %s", err)
	}

	defer mqttContainer.Terminate(ctx)
	port, _ := mqttContainer.MappedPort(ctx, "1883")
	host, _ := mqttContainer.Host(ctx)

	brokerURL := fmt.Sprintf("tcp://%s:%s", host, port.Port())
	c, err := Connect("test", brokerURL)
	if err != nil {
		log.Fatal(err)
	}

	connection = c

	code := m.Run()
	os.Exit(code)
}

type MqttCallbackSpy struct {
	receivedTopic   string
	receivedMessage string
}

func (spy *MqttCallbackSpy) callback(topic string, message string) {
	spy.receivedTopic = topic
	spy.receivedMessage = message
}

func TestMqttClient(t *testing.T) {
	callbackSpy := new(MqttCallbackSpy)

	connection.Subscribe("home/device/+/event", 0, callbackSpy.callback)

	var jsonString = `{"header":{"message_id":"123"},"body":{"light":"on"}}`
	err := connection.Publish("home/device/1/event", jsonString)
	if err != nil {
		t.Fatal(err)
	}

	for callbackSpy.receivedTopic == "" {
		time.Sleep(5 * time.Millisecond)

		if callbackSpy.receivedTopic != "home/device/1/event" {
			t.Errorf("Expected received topic ´home/device/1/event´. Got '%s'", callbackSpy.receivedTopic)
		}

		if callbackSpy.receivedMessage != `{"header":{"message_id":"123"},"body":{"light":"on"}}` {
			t.Errorf(`Expected message "{"header":{"message_id":"123"},"body":{"light":"on"}}". Got '%s'`, callbackSpy.receivedMessage)
		}
	}
}
