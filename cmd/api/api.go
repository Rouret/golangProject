package main

import (
	"log"
	"net/http"
	"time"

	Controllers "github.com/Rouret/golangProject/internal/api/controllers"
	Persitence "github.com/Rouret/golangProject/internal/api/persistence"
	Router "github.com/Rouret/golangProject/internal/api/router"
	Models "github.com/Rouret/golangProject/internal/models"
)

func main() {
	//Create and Register the routes
	router := Router.NewRouter(getRoutes())

	testCreationMessages()

	//ListenAndServe rejte une erreur si il y a un probléme
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getRoutes() Models.Routes {
	return Models.Routes{
		Models.Route{
			Method: "GET",
<<<<<<< Updated upstream
			Path: "/messages",
			Handle: Controllers.GetAllMessages,//OK
		},
		Models.Route{
			Method: "GET",
			Path: "/airport/:iata",
			Handle: Controllers.GetAllMessageByAirportId,//OK
=======
			Path:   "/messages",
			Handle: Controllers.GetAllMessages,
		},
		Models.Route{
			Method: "GET",
			Path:   "/airport/:iata",
			Handle: Controllers.GetAllMessageByAirportId,
>>>>>>> Stashed changes
		},
		Models.Route{
			Method: "GET",
			Path:   "/airport/:iata/type/:type",
			Handle: Controllers.GetAllMessageByAirportIdAndValueType,
		},
		Models.Route{
			Method: "GET",
			Path:   "/airport/:iata/type/:type/date/:dateDay/moy",
			Handle: Controllers.GetAverageValueByAirportIdValueTypeAndDateDay,
		},
		Models.Route{
<<<<<<< Updated upstream
=======
			Method: "GET",
			Path:   "/airport",
			Handle: Controllers.GetAllAirportIds,
		},
		Models.Route{
>>>>>>> Stashed changes
			Method: "POST",
			Path:   "/messages",
			Handle: Controllers.CreateMessage,
		},
	}
}

func testCreationMessages() {
	Persitence.CreateMessage(Models.Message{
		IdCapteur: 1,
		IATA:      "AAA",
		TypeValue: "TEMP",
		Value:     15.6,
		Timestamp: time.Now().Unix(),
	})

	Persitence.CreateMessage(Models.Message{
		IdCapteur: 2,
<<<<<<< Updated upstream
		IATA: "AAA",
=======
		IATA:      "GGG",
>>>>>>> Stashed changes
		TypeValue: "PRESS",
		Value:     24.2,
		Timestamp: time.Now().Unix(),
	})
}
