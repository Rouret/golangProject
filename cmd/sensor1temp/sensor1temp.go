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

const TOPIC = "temperature"
const DELAY = 10

func main() {
	config := config.GetConfig()
	client := mqtt.Connect(config.BrokerUrl+":"+strconv.Itoa(config.BrokerPort), strconv.Itoa(config.ID))
	client.Connect().Wait()

	for range time.Tick(time.Second * DELAY) {
		message, _ := json.Marshal(Message{
			IdCapteur: 1,
			IATA:      "AAA",
			TypeValue: "TEMP",
			Value:     random.GetRandomFloat(20, 25),
			Timestamp: time.Now().Unix(),
		})
		token := client.Publish(TOPIC, byte(config.QOS), false, message)
		token.Wait()
	}
}
