package main

import (
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
	"github.com/go-redis/redis"

	paho "github.com/eclipse/paho.mqtt.golang"
)

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
	fmt.Printf("MSG: %s\n", msg.Payload())
	var info Data
	json.Unmarshal([]byte(msg.Payload()), &info)

	keyBrut := strconv.Itoa(info.IdCapteur) + ":" + info.IATA + ":" + info.TypeValue + ":" + strconv.Itoa(info.Timestamp)

	i, err := strconv.ParseInt(strconv.Itoa(info.Timestamp), 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	date := tm.Format("2006-01-02-03-04-05")

	keyTimestamp := info.IATA + ":" + info.TypeValue + ":" + date

	dateDay := tm.Format("2006-01-02")

	keyAverage := "MOY:" + info.IATA + ":" + info.TypeValue + ":" + dateDay

	clientR := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	json, err := json.Marshal(Data{IdCapteur: info.IdCapteur, IATA: info.IATA, TypeValue: info.TypeValue, Value: info.Value, Timestamp: info.Timestamp})
	if err != nil {
		fmt.Println(err)
	}

	clientR.Set(keyBrut, json, 0).Err()

	clientR.RPush(keyTimestamp, info.Value, 0).Err()

	Average, err := clientR.Get(keyAverage).Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Average)

	/*val, err := clientR.Get(keyBrut).Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)

	fmt.Println(date)*/

}
