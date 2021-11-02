package main

import (
	"log"
	"net/http"
	"time"

	Controllers "github.com/Rouret/golangProject/internal/api/controllers"
	Persitence "github.com/Rouret/golangProject/internal/api/persistence"
	Router "github.com/Rouret/golangProject/internal/api/router"
	Models "github.com/Rouret/golangProject/internal/models"
	"github.com/rs/cors"
)

func main() {
	//Create and Register the routes
	router := Router.NewRouter(getRoutes())
	
	handler := cors.Default().Handler(router)

	//ListenAndServe rejte une erreur si il y a un probléme
	log.Fatal(http.ListenAndServe(":8000", handler))
}

func getRoutes() Models.Routes {
	return Models.Routes{
		Models.Route{
			Method: "GET",
			Path: "/messages",
			Handle: Controllers.GetAllMessages,
		},
		Models.Route{
			Method: "GET",
			Path: "/airports/:iata",
			Handle: Controllers.GetAllMessageByAirportId,
		},
		Models.Route{
			Method: "GET",
			Path: "/airports/:iata/type/:type",
			Handle: Controllers.GetAllMessageByAirportIdAndValueType,
		},
		Models.Route{
			Method: "GET",
			Path: "/airports/:iata/type/:type/dateDay/:dateDay/moy",
			Handle: Controllers.GetAverageValueByAirportIdValueTypeAndDateDay,
		},
		Models.Route{
			Method: "GET",
			Path: "/airports/:iata/type/:type/dateHour/:dateHour/moy",
			Handle: Controllers.GetAverageValueByAirportIdValueTypeAndDateHour,
		},
		Models.Route{
			Method: "GET",
			Path: "/airports",
			Handle: Controllers.GetAllAirportIds,
		},
		Models.Route{
			Method: "POST",
			Path: "/messages",
			Handle: Controllers.CreateMessage,
		},
	}
}


func testCreationMessages() {
	Persitence.CreateMessage(Models.Message{
		IdCapteur: 1,
		IATA: "AAA",
		TypeValue: "TEMP",
		Value: 15.6,
		Timestamp: time.Now().Unix(),
	})
	
	Persitence.CreateMessage(Models.Message{
		IdCapteur: 2,
		IATA: "GGG",
		TypeValue: "PRESS",
		Value: 24.2,
		Timestamp: time.Now().Unix(),
	})
}