package main

import (
	"time"

	"github.com/Rouret/golangProject/internal/config"
	"github.com/Rouret/golangProject/internal/random"
	"github.com/Rouret/golangProject/internal/sensor"
	"github.com/Rouret/mqtt.golang"
)

func main() {
	config := config.GetConfig()
	
	mqtt.Setup(mqtt.LibConfiguration{
		IsPersistent: true,
		BrokerUrl: config.BrokerUrl,
    	BrokerPort: config.BrokerPort,
		ID: config.ID,
	})

	mqtt.Connect()
	
	for range time.Tick(time.Second * time.Duration(config.DelayMessage)) {
		message := sensor.Message{
			IdCapteur: config.ID,
			IATA:      config.IATA,
			TypeValue: config.ValueType,
			Value:     random.GetRandomFloat(10, 50),
			Timestamp: time.Now().Unix(),
		}
		mqtt.Send(config.Topic,config.QOS,message,true)
	}
}
