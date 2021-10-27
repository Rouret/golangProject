package persitence

import (
	"encoding/json"
	"strconv"

	Models "github.com/Rouret/golangProject/internal/models"
	"github.com/go-redis/redis"
)

func redisConnect() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func createKeyId(message Models.Message) string {
	return message.IATA + ":" + message.TypeValue + ":" + strconv.Itoa(message.IdCapteur) + ":" + strconv.FormatInt(message.Timestamp, 10)
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func extractDataFromRedis(redisClient *redis.Client, keyFilter string) Models.Messages {
	var messages Models.Messages
	keys, err := redisClient.Do("KEYS", keyFilter).Result()
	HandleError(err)

	for _, key := range keys.([]interface{}) {

		var message Models.Message

		reply, err := redisClient.Do("GET", key.(string)).String() //Get l'objet de la clé 'key'
		HandleError(err)

		if err := json.Unmarshal([]byte(reply), &message); err != nil { //Transforme le json de la value dans valeur dans la variable message
			HandleError(err)
		}
		messages = append(messages, message) // Ajoute le messages au reste des messages
	}

	return messages
}

func extractFloatFromRedis(redisClient *redis.Client, keyFilter string) []string {
	Average, err := redisClient.LRange(keyFilter, 0, 0).Result() //Get l'objet de la clé 'key'
	HandleError(err)
	//Get l'objet de la clé 'key'
	return Average

}

func FindAllMessages() Models.Messages {
	redisClient := redisConnect()
	defer redisClient.Close()
	return extractDataFromRedis(redisClient, "*")
}

func FindAllMessageByAirportId(IATA string) Models.Messages {
	redisClient := redisConnect()
	defer redisClient.Close()
	return extractDataFromRedis(redisClient, IATA+":*")
}

func FindAllMessageByAirportIdAndValueType(IATA string, valueType string) Models.Messages {
	redisClient := redisConnect()
	defer redisClient.Close()
	return extractDataFromRedis(redisClient, IATA+":"+valueType+":*")
}

func FindAverageValueByAirportIdValueTypeAndDateDay(IATA string, valueType string, dateDay string) []string {
	redisClient := redisConnect()
	defer redisClient.Close()
	return extractFloatFromRedis(redisClient, "MOY:"+IATA+":"+valueType+":"+dateDay)

}

func CreateMessage(message Models.Message) {
	redisClient := redisConnect()
	defer redisClient.Close()

	jsonMessage, err := json.Marshal(message)
	HandleError(err)

	_, err = redisClient.Do("SET", createKeyId(message), jsonMessage).Result()
	HandleError(err)
}
