package main

import (
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"

	mqttClient "github.com/nunocgoncalves/mqtt-transformer/mqtt"
)

func main() {
	// MQTT client setup
	centralUniqueID := uuid.New()
	clientid := "mqtt-" + centralUniqueID.String()
	mqttBroker := os.Getenv("MQTT_BROKER")
	mqttPort := os.Getenv("MQTT_PORT")
	mqttUser := os.Getenv("MQTT_USER")
	mqttPassword := os.Getenv("MQTT_PASSWORD")
  mqttTopic := "shellies/plug-master1-server/relay/0/power"

	mqttInstance := mqttClient.NewClient(clientid, mqttBroker, mqttPort, mqttUser, mqttPassword)
	if err := mqttInstance.Connect(); err != nil {
		log.Fatalf("Failed to connect to MQTT broker: %v", err)
	}
	defer mqttInstance.Disconnect()

	if err := mqttInstance.Subscribe(mqttTopic, handleMessage); err != nil {
		log.Fatalf("Failed to subscribe to topic %s: %v", mqttTopic, err)
	}
}

func handleMessage(client mqtt.Client, msg mqtt.Message) {	
	log.Printf("New message: %+v", msg)
  client.Publish("city0/servers/master1/power", 0, false, msg)
}

