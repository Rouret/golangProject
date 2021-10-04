package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Configuration struct {
	BrokerUrl    string
	BrokerPort   int
	QOS          int
	ID           int
	IATA         string
	ValueType    string
	Topic        string
	DelayMessage int
}

func GetConfig() Configuration {
	var config Configuration
	getEnvConfig()

	if len(os.Args) != 1 {
		fmt.Println("CONFIG BY CLI")
		fmt.Println("WARNING:  Si des paramètres ne sont pas renseignés, ils prennent les valeurs pas défaut.")
		config = getCLIConfig()
	} else {
		configENV, isOnePresent := getEnvConfig()
		if isOnePresent {
			fmt.Println("CONFIG BY ENV")
			config = configENV
		} else {
			fmt.Println("CONFIG BY CONF FILE")
			config = getConfigByFile()
		}
		d
	}
	return config
}

// Exemple : *.exe -host=test -port=123 -qos=0 -id=1234 -iata=TEST -type=TEST -topic=test -delay=1
// Exemple :  fonctionnel -host=tcp://localhost -port=1883 -qos=0 -id=123 -iata=CGT -type=TEMP -topic=temperature -delay=1
func getCLIConfig() Configuration {
	defaultConfig := getDefaultCongif()

	brokerUrl := flag.String("host", defaultConfig.BrokerUrl, "URL du broker")
	port := flag.Int("port", defaultConfig.BrokerPort, "Port du broker")
	qos := flag.Int("qos", defaultConfig.QOS, "QOS")
	id := flag.Int("id", defaultConfig.ID, "ID client")
	IATA := flag.String("iata", defaultConfig.IATA, "Code IATA")
	valueType := flag.String("type", defaultConfig.ValueType, "Type de la valeur")
	topic := flag.String("topic", defaultConfig.Topic, "Topic mqqt")
	delayMessage := flag.Int("delay", defaultConfig.DelayMessage, "Délai entre chaque message")

	flag.Parse()

	return Configuration{
		BrokerUrl:    *brokerUrl,
		BrokerPort:   *port,
		QOS:          *qos,
		ID:           *id,
		IATA:         *IATA,
		ValueType:    *valueType,
		Topic:        *topic,
		DelayMessage: *delayMessage,
	}
}

func getConfigByFile() Configuration {
	fileName := "./config.json"
	file, _ := os.Open(fileName)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Errorf("Fichier " + fileName + " est introuvable, la configuration par défaut est appliquée")
		configuration = getDefaultCongif()
	}
	return configuration
}

func getEnvConfig() (Configuration, bool) {
	isOnePresent := false
	defaultConfig := getDefaultCongif()
	reflectDefaultConfig := reflect.ValueOf(&defaultConfig)
	elm := reflectDefaultConfig.Elem()

	for i := 0; i < elm.NumField(); i++ {
		paramName := elm.Type().Field(i).Name
		paramType := elm.Type().Field(i).Type
		value, isPresent := os.LookupEnv(paramName)
		if isPresent {
			isOnePresent = true
			switch paramType.String() {
			case "string":
				elm.FieldByName(paramName).SetString(value)
				break
			case "int":
				intVar, _ := strconv.ParseInt(value, 10, 64)
				elm.FieldByName(paramName).SetInt(intVar)
				break
			}
		}
	}
	return elm.Interface().(Configuration), isOnePresent
}

func getDefaultCongif() Configuration {
	return Configuration{
		BrokerUrl:    "tcp://localhost",
		BrokerPort:   1883,
		QOS:          0,
		ID:           123,
		IATA:         "CGN",
		ValueType:    "TEMP",
		Topic:        "default",
		DelayMessage: 10,
	}
}
