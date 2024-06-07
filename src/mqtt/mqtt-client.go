package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"net/url"
	"time"
)

type Connection interface {
	Subscribe(topic string, qos byte, callback func(string, string))
	Publish(topic string, message string) error
}

type mqttConnection struct {
	client mqtt.Client
}

func (c mqttConnection) Subscribe(topic string, qos byte, callback func(string, string)) {
	token := c.client.Subscribe(topic, qos, func(client mqtt.Client, message mqtt.Message) {
		go callback(message.Topic(), string(message.Payload()))
	})
	if token.Error() != nil {
		log.Fatal("error subscribing to topic", token.Error())
	}
}

func (c mqttConnection) Publish(topic string, message string) error {
	token := c.client.Publish(topic, 0, false, message)
	return token.Error()
}

func newClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(uri.String())
	opts.SetClientID(clientId)
	return opts
}

func Connect(clientId, uri string) (Connection, error) {
	mqttUrl, err := url.Parse(uri)
	if err != nil {
		log.Fatalf("Cannot parse mqtt string url: %s", uri)
	}
	client := mqtt.NewClient(newClientOptions(clientId, mqttUrl))
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}

	if err := token.Error(); err != nil {
		return nil, err
	}

	return mqttConnection{
		client: client,
	}, nil
}
