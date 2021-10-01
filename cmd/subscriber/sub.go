package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	broker "foo.org/myapp/pkg/mqtt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const TOPIC = "temperature"

var knt int

type Data struct {
	IdCapteur int
	IATA      string
	TypeValue string
	Value     float64
	Timestamp int
}

func main() {
	knt = 0
	keepAlive := make(chan os.Signal)
	signal.Notify(keepAlive, os.Interrupt, syscall.SIGINT)
	//config := config.GetConfig()
	//client := mqtt.Connect(config.BrokerUrl+":"+strconv.Itoa(config.BrokerPort), strconv.Itoa(config.ID))
	client := broker.Connect("tcp://localhost:1883", "samir2")
	client.Connect().Wait()

	if token := client.Subscribe(TOPIC, 0, onReceive); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	fmt.Println("Subscribed")

	<-keepAlive

}

var onReceive mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	//fmt.Printf("MSG: %s\n", msg.Payload())
	var info Data
	json.Unmarshal([]byte(msg.Payload()), &info)
	fmt.Printf(info.IATA + "\n")

}
