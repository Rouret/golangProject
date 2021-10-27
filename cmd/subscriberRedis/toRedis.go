package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"strconv"
	"strings"
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

	keepAlive := make(chan os.Signal)
	signal.Notify(keepAlive, os.Interrupt, syscall.SIGINT)

	mqtt.Setup(mqtt.LibConfiguration{
		IsPersistent: true,
		BrokerUrl:    config.BrokerUrl,
		BrokerPort:   config.BrokerPort,
		ID:           config.ID,
	})

	mqtt.Connect()

	mqtt.Subscribe(config.Topic, byte(config.QOS), onReceive)

	log.Println("Subscribed")

	<-keepAlive

}

var onReceive paho.MessageHandler = func(client paho.Client, msg paho.Message) {
	//fmt.Printf("MSG: %s\n", msg.Payload())
	var info Data
	json.Unmarshal([]byte(msg.Payload()), &info)

	keyBrut := info.IATA + ":" + info.TypeValue + ":" + strconv.Itoa(info.IdCapteur) + ":" + strconv.Itoa(info.Timestamp)

	i, err := strconv.ParseInt(strconv.Itoa(info.Timestamp), 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	//	date := tm.Format("2006-01-02-03-04-05")

	//	keyTimestamp := info.IATA + ":" + info.TypeValue + ":" + date

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

	//clientR.RPush(keyTimestamp, info.Value).Err()
	//fmt.Println(keyTimestamp + "exist")

	Average, err := clientR.Get(keyAverage).Result()
	if err == redis.Nil {
		clientR.LPush(keyAverage, 1).Err()
		clientR.LPush(keyAverage, info.Value).Err()

		fmt.Println("creating " + keyAverage)
	} else {
		fmt.Println("keyAverage", Average)
		NewAverage, err := clientR.LRange(keyAverage, 0, 0).Result()
		if err == redis.Nil {
		}
		Count, err := clientR.LRange(keyAverage, 1, 1).Result()
		if err == redis.Nil {
		}

		avg, err := strconv.ParseFloat(arrayToString(NewAverage), 8)
		cnt, err := strconv.ParseFloat(arrayToString(Count), 8)

		//fmt.Println(avg)
		//fmt.Println(cnt)
		CalculatedAvg := ((cnt*avg + info.Value) / (cnt + 1))

		clientR.LSet(keyAverage, 0, math.Round(CalculatedAvg*1000)/1000).Err()
		clientR.LSet(keyAverage, 1, cnt+1).Err()

	}

}

func arrayToString(a []string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", "", -1), "[]")

}
