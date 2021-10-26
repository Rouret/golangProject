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

func RedisConnect() *redis.Client {
	//c, err := redis.Dial("tcp", ":6379")

	clientR := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })


	//HandleError(err)
	return clientR
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}


func FindAll() Models.Messages {
	println("FindAll() method")

	var messages Models.Messages

	c := RedisConnect()
	defer c.Close()
	
	println("FindAll() before do")

	//Redigo returns everithing as type interface{}
	keys, err := c.Do("KEYS", "message:*").Result()
	//keys, err := redis.Scan(c.Do("KEYS", "message:*"))
	//keys, err := redis.Strings(c.Do("KEYS", "*"))
	HandleError(err)
	
	println("FindAll() after do ", keys)

	for _, k := range keys.([]interface{}) {
		
		var message Models.Message
		println("FindAll() boucle do before get ")
		//reply, err := c.Do("GET", k.([]byte)).Result()
		reply, err := c.Do("GET", k.(string)).String() //test et converti si type String
		HandleError(err) //erreur
		println("FindAll() boucle do after get ")
		
		//[]byte(reply) fonction pour convertir en byte
		if err := json.Unmarshal([]byte(reply), &message); err != nil {
			panic(err)
		}
		messages = append(messages, message)
	}

	println("FindAll() fin")
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

	c := RedisConnect()
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
	
	c := RedisConnect()
	defer c.Close()
	
	b, err := json.Marshal(m)
	HandleError(err)
	
	// Save JSON blob to Redis
	reply, err := c.Do("SET", "message:" + strconv.Itoa(m.IdCapteur), b).Result()
	HandleError(err)
	
	fmt.Println("POST ", reply)
}

