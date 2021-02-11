package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
var Articles = []Article{
	{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
}
func main() {
	log.Info().Msg("Rest API v2.0 - Mux Routers")

	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request){
	if _, err := fmt.Fprintf(w, "Welcome to the HomePage!"); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't serve string on homepage")
	}
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests(){
	var port = "127.0.0.1:8080"
	log.Info().Msgf("Starting API on port:\n", port)

	myRouter := mux.NewRouter().StrictSlash(true)
	//Interface for UI
	myRouter.HandleFunc("/", homePage)

	myRouter.HandleFunc("/article/{id}", returnSingleCard)

	if err := http.ListenAndServe(port, myRouter); err != nil {
		log.Panic().Timestamp().Err(err).Msg("Panic: problem with TCP network connection")
	}
}

func returnSingleCard(w http.ResponseWriter, r *http.Request){
	log.Info().Msg("Endpoint Hit: returnSingleCard")

	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			if err := json.NewEncoder(w).Encode(article); err != nil {
				log.Fatal().Timestamp().Err(err).Msg("Fatal: problem with encoding struct to json")
			}
		}
	}

	/*if err := json.NewEncoder(w).Encode(Articles); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: problem with encoding struct to json")
	}*/

}