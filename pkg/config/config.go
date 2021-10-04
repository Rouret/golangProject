package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	BrokerUrl  string
	BrokerPort int
	Delay      int
	QOS        int
	ID1        int
	ID2        int
	ID3        int
}

func GetConfig() Configuration {
	fmt.Println(len(os.Args))

	file, _ := os.Open("./config.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}
