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

func createKeyId(message Models.Message)  string {
	return message.IATA + ":" + message.TypeValue + ":" + strconv.Itoa(message.IdCapteur) + ":" + strconv.FormatInt(message.Timestamp, 10) 
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func extractDataFromRedis(redisConnection *redis.Client,keyFilter string) Models.Messages {
	var messages Models.Messages
	keys, err := redisConnection.Do("KEYS", keyFilter).Result()
	HandleError(err)
	
	for _, key := range keys.([]interface{}) {

		var message Models.Message

		reply, err := redisConnection.Do("GET", key.(string)).String() //Get l'objet de la clé 'key'
		HandleError(err)

		if err := json.Unmarshal([]byte(reply), &message); err != nil { //Transforme le json de la value dans valeur dans la variable message
			HandleError(err)
		}
		messages = append(messages, message) // Ajoute le messages au reste des messages
	}

	return messages
}


func FindAllMessages() Models.Messages {
	println("FindAll() method")

	redisConnection := redisConnect()
	defer redisConnection.Close()
	
	return extractDataFromRedis(redisConnection,"*")
}

func FindAllMessageByAirportId(IATA string) Models.Messages {
	redisConnection := redisConnect()
	defer redisConnection.Close()
	
	return extractDataFromRedis(redisConnection,IATA + ":*")
}

func FindAllMessageByAirportIdAndValueType(IATA string,valueType string) Models.Messages {
	redisConnection := redisConnect()
	defer redisConnection.Close()
	
	return extractDataFromRedis(redisConnection,IATA + ":" + valueType + ":*")
}

func CreateMessage(message Models.Message) {
	
	redisConnection := redisConnect()
	defer redisConnection.Close()
	
	jsonMessage, err := json.Marshal(message)
	HandleError(err)
	
	_, err = redisConnection.Do("SET", createKeyId(message), jsonMessage).Result()
	HandleError(err)
}

