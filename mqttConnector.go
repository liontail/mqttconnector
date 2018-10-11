package mqttconnector

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttConnector struct {
	Username string
	Password string
	MqttURL  string
	Client   mqtt.Client
}

func (mc *MqttConnector) Connect() error {
	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s", mc.MqttURL))
	opts.SetUsername(mc.Username)
	opts.SetPassword(mc.Password)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	mc.Client = client
	for !token.WaitTimeout(3 * time.Second) {
		return errors.New("Connection MQTT Timeout")
	}
	if err := token.Error(); err != nil {
		return err
	}
	fmt.Printf("MQTT %s Connected", mc.MqttURL)
	return nil
}

func (mc *MqttConnector) ListenTo(topic string, f func(Message)) {
	mc.Client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		// fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
		message := Message{}
		if err := json.Unmarshal(msg.Payload(), &message); err != nil {
			log.Panic(err)
		}
		message.Topic = msg.Topic()
		f(message)
	})
}
