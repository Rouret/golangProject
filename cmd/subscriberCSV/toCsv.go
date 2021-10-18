package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

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
	var info Data

	json.Unmarshal([]byte(msg.Payload()), &info)
	fmt.Printf("Info Received from " + info.IATA + "\n")
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
		fmt.Println("err")
		return
	}
	w := csv.NewWriter(f)
	for _, csvRow := range csvData {
		_ = w.Write(csvRow)
	}
	w.Flush()

}
