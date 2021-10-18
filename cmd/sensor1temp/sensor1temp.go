package main

import (
	"strconv"
	"time"

	"github.com/Rouret/golangProject/internal/config"
	"github.com/Rouret/golangProject/internal/random"
	"github.com/Rouret/mqtt.golang"
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
	
	mqtt.Setup(mqtt.LibConfiguration{
		IsPersistent: true,
	})

	mqtt.Connect(config.BrokerUrl+":"+brokerPort, idClient)
	
	for range time.Tick(time.Second * time.Duration(config.DelayMessage)) {
		message := Message{
			IdCapteur: config.ID,
			IATA:      config.IATA,
			TypeValue: config.ValueType,
			Value:     random.GetRandomFloat(20, 25),
			Timestamp: time.Now().Unix(),
		}
		mqtt.Send(config.Topic,config.QOS,message,true)
	}
}
