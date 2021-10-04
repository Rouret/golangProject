package main

import (
	"encoding/json"
	"strconv"
	"time"

	"foo.org/myapp/pkg/config"
	"foo.org/myapp/pkg/mqtt"
	"foo.org/myapp/pkg/random"
)

type Message struct {
	IdCapteur int
	IATA      string
	TypeValue string
	Value     float32
	Timestamp int64
}

func main() {
	config := config.GetConfig()
	brokerPort := strconv.Itoa(config.BrokerPort)
	idClient := strconv.Itoa(config.ID)
	client := mqtt.Connect(config.BrokerUrl+":"+brokerPort, idClient)
	client.Connect().Wait()

	for range time.Tick(time.Second * time.Duration(config.DelayMessage)) {
		message, _ := json.Marshal(Message{
			IdCapteur: config.ID,
			IATA:      config.IATA,
			TypeValue: config.ValueType,
			Value:     random.GetRandomFloat(20, 25),
			Timestamp: time.Now().Unix(),
		})
		token := client.Publish(config.Topic, byte(config.QOS), false, message)
		token.Wait()
	}
}
