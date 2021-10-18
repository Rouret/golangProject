package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

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
	mqtt.Subscribe(config.Topic, byte(config.ID), onReceive)

	log.Println("Subscribed")

	<-keepAlive

}

var onReceive paho.MessageHandler = func(client paho.Client, msg paho.Message) {
	var info Data

	json.Unmarshal([]byte(msg.Payload()), &info)
	log.Printf("Info Received from " + info.IATA + "\n")
	i, err := strconv.ParseInt(strconv.Itoa(info.Timestamp), 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	date := tm.Format("2006-01-02")

	csvName := info.IATA + "-" + date + "-" + info.TypeValue + ".csv"

	var csvData = [][]string{}

	if _, err := os.Stat(csvName); err == nil {
		csvData = [][]string{
			{strconv.Itoa(info.IdCapteur), info.IATA, info.TypeValue, fmt.Sprintf("%f", info.Value), tm.String()},
		}
	} else {
		csvData = [][]string{
			{"IDCapteur", "IATA", "TypeValue", "Value", "Timestamp"},
			{strconv.Itoa(info.IdCapteur), info.IATA, info.TypeValue, fmt.Sprintf("%f", info.Value), tm.String()},
		}
	}

	f, err := os.OpenFile(csvName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Println("err")
		return
	}
	w := csv.NewWriter(f)
	for _, csvRow := range csvData {
		_ = w.Write(csvRow)
	}
	w.Flush()

}
