package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
)

type ReqCard struct {
	Name string
	SetName string
}

func main() {
	handleRequests()
}

func statusAlive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("\"OK\""))
}
func statusCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	body := struct {
		ResponseCode       int    `json:"-"`
	}{
		ResponseCode:       http.StatusOK,
	}

	marshalledObject, err := json.Marshal(body)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Could not marshal body")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("something bad happened, please contact the administrator"))
		return
	}

	w.WriteHeader(body.ResponseCode)
	_, _ = w.Write(marshalledObject)
}
//TODO: serves UI
func homePage(w http.ResponseWriter, r *http.Request){
	if _, err := fmt.Fprintf(w, "Welcome to the HomePage!"); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't serve string on homepage")
	}
	fmt.Println("Endpoint Hit: homePage")
}