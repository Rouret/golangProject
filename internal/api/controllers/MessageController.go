package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	Persitence "github.com/Rouret/golangProject/internal/api/persistence"
	Models "github.com/Rouret/golangProject/internal/models"
	mux "github.com/julienschmidt/httprouter"
)

func prepareResponseWriter(responseWriter http.ResponseWriter) http.ResponseWriter {
	responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return responseWriter
}

func GetAllMessages(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	log.Println("GetAllMessages requested")

	w = prepareResponseWriter(w)
	w.WriteHeader(http.StatusOK)

	messages := Persitence.FindAllMessages()

	if err := json.NewEncoder(w).Encode(messages); err != nil {
		panic(err)
	}
}

func GetAllMessageByAirportId(w http.ResponseWriter, r *http.Request, p mux.Params) {

	w = prepareResponseWriter(w)

	IATA := p.ByName("iata") //conversion de l'id récupéré en une variable integer

	log.Println("GetAllMessageByAirportId requested (IATA=" + IATA + ")")

	messages := Persitence.FindAllMessageByAirportId(IATA) // We'll work on this

	//Encoder le résultat en json
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		panic(err)
	}
}

func GetAllMessageByAirportIdAndValueType(w http.ResponseWriter, r *http.Request, p mux.Params) {
	w = prepareResponseWriter(w)

	IATA := p.ByName("iata") //conversion de l'id récupéré en une variable integer
	valueType := p.ByName("type")

	log.Println("GetAllMessageByAirportId requested (IATA=" + IATA + ", ValuType=" + valueType + ")")

	messages := Persitence.FindAllMessageByAirportIdAndValueType(IATA, valueType) // We'll work on this

	//Encoder le résultat en json
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		panic(err)
	}
}
func GetAverageValueByAirportIdValueTypeAndDateDay(w http.ResponseWriter, r *http.Request, p mux.Params){
	w = prepareResponseWriter(w)

	IATA := p.ByName("iata")
	valueType := p.ByName("type")
	dateDay := p.ByName("dateDay")

	log.Println("GetAverageValueByAirportIdValueTypeAndDateDay requested (IATA=" + IATA + ", ValueType=" + valueType + ", DateDay=" + dateDay + ")")

	messages := Persitence.FindAverageValueByAirportIdValueTypeAndDateDay(IATA, valueType, dateDay) // We'll work on this

	//Encoder le résultat en json
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		panic(err)
	}
}



func GetAverageValueByAirportIdValueTypeAndDateHour(w http.ResponseWriter, r *http.Request, p mux.Params) {
	w = prepareResponseWriter(w)

	IATA := p.ByName("iata")
	valueType := p.ByName("type")
	dateHour := p.ByName("dateHour")

	log.Println("GetAverageValueByAirportIdValueTypeAndDateHour requested (IATA=" + IATA + ", ValueType=" + valueType + ", DateHour=" + dateHour + ")")

	messages := Persitence.FindAverageValueByAirportIdValueTypeAndDateHour(IATA, valueType, dateHour) // We'll work on this

	//Encoder le résultat en json
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		panic(err)
	}
}

func GetAllAirportIds(w http.ResponseWriter, r *http.Request, _ mux.Params){
	w = prepareResponseWriter(w)

	log.Println("GetAllAirportIds requested")

	messages := Persitence.FindAllAirportIds() 

	if err := json.NewEncoder(w).Encode(messages); err != nil {
		panic(err)
	}
}

func CreateMessage(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	w = prepareResponseWriter(w)

	var message Models.Message

	// Read request body and close it
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	Persitence.HandleError(err)
	defer r.Body.Close()

	// Save JSON to Message struct
	if err := json.Unmarshal(body, &message); err != nil {
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	Persitence.CreateMessage(message) // We'll work on this
	w.WriteHeader(http.StatusCreated)
}
