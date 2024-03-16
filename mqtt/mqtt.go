package mqtt

import (
	"crypto/tls"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Client interface {
	Connect() error
	Subscribe(topic string, callback mqtt.MessageHandler) error
	Publish(topic string, qos byte, retained bool, payload interface{}) error
	Disconnect()
  IsConnected() bool
}

type client struct {
	mqttClient mqtt.Client
}

func NewClient(clientid, broker string, port string, username, password string) Client {
	opts := mqtt.NewClientOptions()
	opts.SetClientID(clientid)
	opts.AddBroker(fmt.Sprintf("tls://%s:%s", broker, port))
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetTLSConfig(&tls.Config{InsecureSkipVerify: true}) // Adjust this based on your TLS configuration

	opts.OnConnect = func(c mqtt.Client) {
		log.Println("Connected to MQTT broker")
	}
	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		log.Printf("Connection lost: %v. Attempting to reconnect...\n", err)
	}

	return &client{mqttClient: mqtt.NewClient(opts)}
}

func (c *client) Connect() error {
	if token := c.mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (c *client) Subscribe(topic string, callback mqtt.MessageHandler) error {
	if token := c.mqttClient.Subscribe(topic, 0, callback); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (c *client) Publish(topic string, qos byte, retained bool, payload interface{}) error {
	token := c.mqttClient.Publish(topic, qos, retained, payload)
	token.Wait()
	return token.Error()
}

func (c *client) Disconnect() {
	c.mqttClient.Disconnect(250)
}

func (c *client) IsConnected() bool {
  return c.mqttClient.IsConnected()
}

