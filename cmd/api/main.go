package main

import (
	"log"
	"net/http"

	"foo.org/myapp/api"
)

func main() {
	router := api.NewRouter()
	
	initCreation()

	log.Fatal(http.ListenAndServe(":8080", router))
}



// Give us some seed data
// init() will be automatically launch upon running
func initCreation() {
	api.CreateMessage(api.Message{
		IdCapteur: 1,
		IATA: "AAA",
		TypeValue: "TEMP",
		Value: 15.6,
	})
	
	api.CreateMessage(api.Message{
		IdCapteur: 2,
		IATA: "AAA",
		TypeValue: "PRESS",
		Value: 24.2,
	})
}