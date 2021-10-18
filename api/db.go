package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

var currentMessageId int

func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	HandleError(err)
	return c
}



func FindAll() Messages {
	
	var messages Messages

	c := RedisConnect()
	defer c.Close()
	
	//Redigo returns everithing as type interface{}
	keys, err := c.Do("KEYS", "message:*")
	HandleError(err)
	
	for _, k := range keys.([]interface{}) {
		
		var message Message
		
		reply, err := c.Do("GET", k.([]byte))
		HandleError(err)
		
		if err := json.Unmarshal(reply.([]byte), &message); err != nil {
			panic(err)
		}
		messages = append(messages, message)
	}
	return messages
}

func FindMessage(id int) Message {
	
	var message Message

	c := RedisConnect()
	defer c.Close()
	
	reply, err := c.Do("GET", "message:" + strconv.Itoa(id))
	HandleError(err)
	
	fmt.Println("GET OK")
	
	if err = json.Unmarshal(reply.([]byte), &message); err != nil {
		panic(err)
	}
	return message
}

func CreateMessage(m Message) {
	
	currentMessageId += 1
	
	m.IdCapteur = currentMessageId
	m.Timestamp = time.Now().Unix() //Unix() to convert into int
	
	c := RedisConnect()
	defer c.Close()
	
	b, err := json.Marshal(m)
	HandleError(err)
	
	// Save JSON blob to Redis
	reply, err := c.Do("SET", "post:" + strconv.Itoa(m.IdCapteur), b)
	HandleError(err)
	
	fmt.Println("GET ", reply)
}