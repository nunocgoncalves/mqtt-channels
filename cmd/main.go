package main

import (
	"crypto/tls"
	"log"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"

	"github.com/nunocgoncalves/mqtt-transformer/internal/mqtt"
)

type MQTTClientHolder struct {
	SourceClient      mqtt.Client
	DestinationClient mqtt.Client
}

func main() {
	// TODO: Refactor to use dynamic config
	sourceConfig := mqtt.ClientConfig{
		ClientID:  uuid.New().String(),
		Broker:    os.Getenv("SOURCE_MQTT_BROKER"),
		Port:      os.Getenv("SOURCE_MQTT_PORT"),
		Username:  os.Getenv("SOURCE_MQTT_USER"),
		Password:  os.Getenv("SOURCE_MQTT_PASSWORD"),
		UseTLS:    false,
		TLSConfig: nil,
	}

	destinationConfig := mqtt.ClientConfig{
		ClientID:  uuid.New().String(),
		Broker:    os.Getenv("DESTINATION_MQTT_BROKER"),
		Port:      os.Getenv("DESTINATION_MQTT_PORT"),
		Username:  os.Getenv("DESTINATION_MQTT_USER"),
		Password:  os.Getenv("DESTINATION_MQTT_PASSWORD"),
		UseTLS:    true,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}

	sourceClient := mqtt.NewClient(sourceConfig)
	if err := sourceClient.Connect(); err != nil {
		log.Fatalf("Failed to connect to source MQTT broker: %v", err)
	}
	defer sourceClient.Disconnect()

	destinationClient := mqtt.NewClient(destinationConfig)
	if err := destinationClient.Connect(); err != nil {
		log.Fatalf("Failed to connect to destination MQTT broker: %v", err)
	}
	defer destinationClient.Disconnect()

	// Initialize MQTTClientHolder
	holder := MQTTClientHolder{
		SourceClient:      sourceClient,
		DestinationClient: destinationClient,
	}

	// TODO: Refactor to handle subcriptions dynamically
	// Subscribe using the holder's method
	if err := sourceClient.Subscribe("topic", holder.handleMessage); err != nil {
		log.Fatalf("Failed to subscribe to source topic: %v", err)
	}

	select {}
}

// TODO: Refactor to handle publications dynamically
func (m *MQTTClientHolder) handleMessage(client MQTT.Client, msg MQTT.Message) {
	log.Printf("New message: %+v", msg.Payload())
	// Use your interface's method to publish
	if err := m.DestinationClient.Publish("topic", 0, false, msg.Payload()); err != nil {
		log.Printf("Failed to publish message: %v", err)
	}
}
