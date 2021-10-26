package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	mux "github.com/julienschmidt/httprouter"
)

//Afficher un string
func Index(w http.ResponseWriter, r *http.Request, _ mux.Params) {
    fmt.Fprintf(w, "<h1>Hello, welcome to my blog</h1>")
}

//Afficher all
func MessageIndex(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	
	messages := FindAll()  // TODO

	if err := json.NewEncoder(w).Encode(messages); err != nil {
			panic(err)
	}
}

//Afficher un élément par l'id
func MessageShow(w http.ResponseWriter, r *http.Request, p mux.Params) {
	id, err := strconv.Atoi(p.ByName("messageId")) //conversion de l'id récupéré en une variable integer
	HandleError(err)

	message := FindMessage(id)   // We'll work on this
	
	//Encoder le résultat en json
	if err := json.NewEncoder(w).Encode(message); err != nil {
			panic(err)
	}
}

func MessageCreate(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	var message Message
	
	// Read request body and close it
	body, err := ioutil.ReadAll(io.LimitReader(r.Body,1048576))
	HandleError(err)
	defer r.Body.Close()

	// Save JSON to Message struct
	if err := json.Unmarshal(body, &message); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)

		if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
		}
	}
	
	CreateMessage(message)   // We'll work on this
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

