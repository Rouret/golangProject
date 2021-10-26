package main

import (
	"log"
	"net/http"

	Controllers "github.com/Rouret/golangProject/internal/api/controllers"
	Persitence "github.com/Rouret/golangProject/internal/api/persistence"
	Router "github.com/Rouret/golangProject/internal/api/router"
	Models "github.com/Rouret/golangProject/internal/models"
)

func main() {
	//Create and Register the routes
	router := Router.NewRouter(getRoutes())
	
	testCreationMessage()

	//ListenAndServe rejte une erreur si il y a un probléme
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getRoutes() Models.Routes {
	return Models.Routes{
		Models.Route{
			Method: "GET",
			Path: "/",
			Handle: Controllers.Index,
		},
		Models.Route{
			Method: "GET",
			Path: "/messages",
			Handle: Controllers.MessageIndex,
		},
		Models.Route{
			Method: "GET",
			Path: "/messages/:id",
			Handle: Controllers.MessageShow,
		},
		Models.Route{
			Method: "POST",
			Path: "/messages",
			Handle: Controllers.MessageCreate,
		},
	}
}

// Give us some seed data
// init() will be automatically launch upon running
func testCreationMessage() {
	Persitence.CreateMessage(Models.Message{
		IdCapteur: 1,
		IATA: "AAA",
		TypeValue: "TEMP",
		Value: 15.6,
	})
	
	Persitence.CreateMessage(Models.Message{
		IdCapteur: 2,
		IATA: "AAA",
		TypeValue: "PRESS",
		Value: 24.2,
	})
}