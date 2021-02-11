package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
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
	log.Info().Msgf("Starting API on port:", port)

	myRouter := mux.NewRouter().StrictSlash(true)
	//Interface for UI
	myRouter.HandleFunc("/", homePage)

	//CRUD Operations
	myRouter.HandleFunc("/article", createNewCardEntry).Methods(http.MethodPost)
	myRouter.HandleFunc("/article/{id}", returnSingleCardEntry).Methods(http.MethodGet)
	myRouter.HandleFunc("/articles", returnAllCardEntries).Methods(http.MethodGet)


	if err := http.ListenAndServe(port, myRouter); err != nil {
		log.Panic().Timestamp().Err(err).Msg("Panic: problem with TCP network connection")
	}
}
func createNewCardEntry(w http.ResponseWriter, r *http.Request) {
	var article Article
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: problem with reading request body")
	}

	if err = json.Unmarshal(reqBody, &article);err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't unmarshal reqBody json into article struct")
	}

	Articles = append(Articles, article)

	if err = json.NewEncoder(w).Encode(Articles); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: problem with writing json encoded struct http.ResponseWriter")
	}

}

func returnAllCardEntries(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Endpoint Hit: returnAllCardEntries")

	if err := json.NewEncoder(w).Encode(Articles); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: problem with writing json encoded struct http.ResponseWriter")
	}
}

func returnSingleCardEntry(w http.ResponseWriter, r *http.Request){
	log.Info().Msg("Endpoint Hit: returnSingleCardEntry")

	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			if err := json.NewEncoder(w).Encode(article); err != nil {
				log.Fatal().Timestamp().Err(err).Msg("Fatal: problem with writing json encoded struct http.ResponseWriter")
			}
		}
	}
}