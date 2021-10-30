package main

import (
	"strconv"
	"time"

	"github.com/Rouret/golangProject/internal/config"
	"github.com/Rouret/golangProject/internal/random"
	"github.com/Rouret/golangProject/internal/sensor"
	"github.com/Rouret/mqtt.golang"
)

func main() {
	config := config.GetConfig()
	brokerPort := strconv.Itoa(config.BrokerPort)
	idClient := strconv.Itoa(config.ID)
	
	mqtt.Setup(mqtt.LibConfiguration{
		IsPersistent: true,
	})

	mqtt.Connect(config.BrokerUrl+":"+brokerPort, idClient)
	
	for range time.Tick(time.Second * time.Duration(config.DelayMessage)) {
		message := sensor.Message{
			IdCapteur: config.ID,
			IATA:      config.IATA,
			TypeValue: config.ValueType,
			Value:     random.GetRandomFloat(1000, 1030),
			Timestamp: time.Now().Unix(),
		}
		mqtt.Send(config.Topic,config.QOS,message,true)
	}
}
