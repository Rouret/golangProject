package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

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

	keepAlive := make(chan os.Signal)
	signal.Notify(keepAlive, os.Interrupt, syscall.SIGINT)
	//config := config.GetConfig()
	mqtt.Setup(mqtt.LibConfiguration{
		IsPersistent: true,
	})

	mqtt.Connect("tcp://localhost:1883", "samir2")
	mqtt.Subscribe(TOPIC, 0, onReceive)

	fmt.Println("Subscribed")

	<-keepAlive

}

var onReceive paho.MessageHandler = func(client paho.Client, msg paho.Message) {
	//fmt.Printf("MSG: %s\n", msg.Payload())
	var info Data
	json.Unmarshal([]byte(msg.Payload()), &info)
	fmt.Printf(info.IATA + "\n")

}
