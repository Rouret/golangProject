package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"foo.org/myapp/pkg/config"
	"foo.org/myapp/pkg/mqtt"
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
	config := config.GetConfig()
	client := mqtt.Connect(config.BrokerUrl+":"+strconv.Itoa(config.BrokerPort), strconv.Itoa(config.ID3))
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
