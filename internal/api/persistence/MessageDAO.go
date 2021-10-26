package persitence

import (
	"encoding/json"
	"fmt"
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


func FindAllMessages() Models.Messages {
	println("FindAll() method")
	var messages Models.Messages

	redisConnection := redisConnect()
	defer redisConnection.Close()
	
	keys, err := redisConnection.Do("KEYS", "*").Result()
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

func FindMessage(id int) Models.Message {
	
	var message Models.Message

	c := redisConnect()
	defer c.Close()
	fmt.Println("GET  FindMessage id avant do")
	reply, err := c.Do("GET", "message:" + strconv.Itoa(id)).String()
	//reply,err := c.Get("message:" + strconv.Itoa(id)).Result()
	HandleError(err)
	
	fmt.Println("GET OK FindMessage id")
	fmt.Println(reply)
	
	//TODO passer la chaîne JSON en String -> struct Message
	/*data := []byte(reply)
	err2 := json.Unmarshal(data, &message)
	if(err2 != nil) {
		log.Fatal(err2)
	}
	fmt.Println(message)*/

	//message = make(map[byte][]Message)
	errStruct := json.Unmarshal([]byte(reply), &message)
	if errStruct != nil {
		panic(errStruct)
	}

	fmt.Printf("\n\n json object:::: %v", message)


	/*if err = json.Unmarshal(reply.([]byte), &message); err != nil {
		panic(err)
	}*/

	return message
}

func CreateMessage(message Models.Message) {
	
	redisConnection := redisConnect()
	defer redisConnection.Close()
	
	jsonMessage, err := json.Marshal(message)
	HandleError(err)
	
	_, err = redisConnection.Do("SET", createKeyId(message), jsonMessage).Result()
	HandleError(err)
}

