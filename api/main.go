package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	
	initCreation()

	log.Fatal(http.ListenAndServe(":8080", router))
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

// Give us some seed data
// init() will be automatically launch upon running
func initCreation() {
	CreateMessage(Message{
		IdCapteur: 1,
		IATA: "AAA",
		TypeValue: "TEMP",
		Value: 15.6,
	})
	
	CreateMessage(Message{
		IdCapteur: 2,
		IATA: "AAA",
		TypeValue: "PRESS",
		Value: 24.2,
	})
}