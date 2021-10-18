package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Rouret/golangProject/internal/config"
	"github.com/Rouret/mqtt.golang"

	paho "github.com/eclipse/paho.mqtt.golang"
)

const TOPIC = "temperature"

type Data struct {
	IdCapteur int
	IATA      string
	TypeValue string
	Value     float64
	Timestamp int
}

func main() {
	config := config.GetConfig()
	brokerPort := strconv.Itoa(config.BrokerPort)
	idClient := strconv.Itoa(config.ID)
	
	keepAlive := make(chan os.Signal)
	signal.Notify(keepAlive, os.Interrupt, syscall.SIGINT)

	mqtt.Setup(mqtt.LibConfiguration{
		IsPersistent: true,
	})

	mqtt.Connect(config.BrokerUrl+":"+brokerPort, idClient)
	mqtt.Subscribe(config.Topic, byte(config.QOS), onReceive)

	log.Println("Subscribed")

	<-keepAlive

}

var onReceive paho.MessageHandler = func(client paho.Client, msg paho.Message) {
	//fmt.Printf("MSG: %s\n", msg.Payload())
	var info Data
	json.Unmarshal([]byte(msg.Payload()), &info)
	log.Println(info.IATA + "\n")

}
