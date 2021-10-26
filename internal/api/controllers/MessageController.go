package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	Persitence "github.com/Rouret/golangProject/internal/api/persistence"
	Models "github.com/Rouret/golangProject/internal/models"
	mux "github.com/julienschmidt/httprouter"
)

func prepareResponseWriter(responseWriter http.ResponseWriter) http.ResponseWriter{
	responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return responseWriter
}

func GetAllMessages(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	w = prepareResponseWriter(w);
	w.WriteHeader(http.StatusOK)
	
	messages := Persitence.FindAllMessages()

	if err := json.NewEncoder(w).Encode(messages); err != nil {
			panic(err)
	}
}

func GetAllMessageByAirportId(w http.ResponseWriter, r *http.Request, p mux.Params) {
	w = prepareResponseWriter(w);

	IATA := p.ByName("iata") //conversion de l'id récupéré en une variable integer
	
	message := Persitence.FindAllMessageByAirportId(IATA)   // We'll work on this
	
	//Encoder le résultat en json
	if err := json.NewEncoder(w).Encode(message); err != nil {
			panic(err)
	}
}

func GetAllMessageByAirportIdAndValueType(w http.ResponseWriter, r *http.Request, p mux.Params) {
	w = prepareResponseWriter(w);

	IATA := p.ByName("iata") //conversion de l'id récupéré en une variable integer
	valueType := p.ByName("type")

	message := Persitence.FindAllMessageByAirportIdAndValueType(IATA,valueType)   // We'll work on this
	
	//Encoder le résultat en json
	if err := json.NewEncoder(w).Encode(message); err != nil {
			panic(err)
	}
}

func CreateMessage(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	w = prepareResponseWriter(w);

	var message Models.Message
	
	// Read request body and close it
	body, err := ioutil.ReadAll(io.LimitReader(r.Body,1048576))
	Persitence.HandleError(err)
	defer r.Body.Close()

	// Save JSON to Message struct
	if err := json.Unmarshal(body, &message); err != nil {
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
		}
	}
	
	Persitence.CreateMessage(message)   // We'll work on this
	w.WriteHeader(http.StatusCreated)
}

