package persitence

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	Models "github.com/Rouret/golangProject/internal/models"
	"github.com/go-redis/redis"
)

var currentMessageId int

func redisConnect() *redis.Client {
	return redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })
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
	
	keys, err := redisConnection.Do("KEYS", "message:*").Result()
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

/*func FindAll() Messages {
	
	var messages Messages

	c := RedisConnect()
	defer c.Close()
	
	//Redigo returns everithing as type interface{}
	keys, err := c.Get("message:*").Result()
	HandleError(err)
	
	for _, k := range keys.([]interface{}) {
		
		var message Message
		
		reply, err := c.Do("GET", k.([]byte)).Result()
		HandleError(err)
		
		if err := json.Unmarshal(reply.([]byte), &message); err != nil {
			panic(err)
		}
		messages = append(messages, message)
	}
	return messages
}*/


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

func CreateMessage(m Models.Message) {
	
	currentMessageId += 1
	
	m.IdCapteur = currentMessageId
	m.Timestamp = time.Now().Unix() //Unix() to convert into int
	
	c := redisConnect()
	defer c.Close()
	
	b, err := json.Marshal(m)
	HandleError(err)
	
	// Save JSON blob to Redis
	reply, err := c.Do("SET", "message:" + strconv.Itoa(m.IdCapteur), b).Result()
	HandleError(err)
	
	fmt.Println("POST ", reply)
}

