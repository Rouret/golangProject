package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Configuration struct {
	BrokerUrl  string
	BrokerPort int
	QOS        int
	ID         int
	IATA 	   string
	ValueType  string
	Topic      string
	DelayMessage int
}

func GetConfig() Configuration {
	var config Configuration;

	if len(os.Args) != 1 {
		fmt.Println("CONFIG BY CLI")
		fmt.Println("WARNING:  Si des paramètres ne sont pas renseignés, ils prennent les valeurs pas défaut.")
		config = getCLIConfig()
	}else{
		fmt.Println("CONFIG BY CONF FILE")
		config = getConfigByFile()
	}
	return config
}
// Exemple : *.exe -host=test -port=123 -qos=0 -id=1234 -iata=TEST -type=TEST -topic=test -delay=1
// Exemple :  fonctionnel -host=tcp://localhost -port=1883 -qos=0 -id=123 -iata=CGT -type=TEMP -topic=temperature -delay=1
func getCLIConfig() Configuration {
	brokerUrl := flag.String("host", "tcp://localhost", "URL du broker")
	port := flag.Int("port", 1883, "Port du broker")
	qos := flag.Int("qos", 0, "QOS")
	id := flag.Int("id", 123, "ID client")
	IATA := flag.String("iata", "CGN", "Code IATA")
	valueType := flag.String("type", "TEMP", "Type de la valeur")
	topic := flag.String("topic", "default", "Topic mqqt")
	delayMessage := flag.Int("delay",10,"Délai entre chaque message")

	flag.Parse()

	return Configuration{
		BrokerUrl:*brokerUrl,
		BrokerPort:*port,
		QOS:*qos,
		ID:*id,
		IATA:*IATA,
		ValueType:*valueType,
		Topic:*topic,
		DelayMessage:*delayMessage,
	}
}

// func getEnvConfig() Configuration{

// }

func getConfigByFile() Configuration {
	file, _ := os.Open("./config.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}