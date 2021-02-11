package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
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
type ReqCard struct {
	Name string `json:"name"`
	SetName string `json:"setName"`
}

//var dbCollection = "myCardCollection"


func main() {
	handleRequests()
}
//TODO: serves UI
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
	myRouter.HandleFunc("/card", createNewCardEntry).Methods(http.MethodPost)
	//myRouter.HandleFunc("/card/name/{cardName}/set/name/{setName}", returnSingleCardEntry).Methods(http.MethodGet)
	/*myRouter.HandleFunc("/articles", returnAllCardEntries).Methods(http.MethodGet)
	myRouter.HandleFunc("/article/{id}", updateSingleCardEntry).Methods(http.MethodPut)
	myRouter.HandleFunc("/article/{id}", deleteSingleCardEntry).Methods(http.MethodDelete)*/

	if err := http.ListenAndServe(port, myRouter); err != nil {
		log.Panic().Timestamp().Err(err).Msg("Panic: problem with TCP network connection")
	}
}
//TODO: read out of allCards collection with reqBody params and then safes found card into collection myCards
func createNewCardEntry(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Endpoint Hit: createNewCardEntry")
	var card ReqCard

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: problem with reading request body")
	}
	if err = json.Unmarshal(reqBody, &card);err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't unmarshal reqBody json into article struct")
	}
	fmt.Fprint(w,card.SetName,card.Name)

	//card, _ := SingleCardInfo(card.Name,card.SetName, "allCards")


	/*
	Articles = append(Articles, article)

	if err = json.NewEncoder(w).Encode(Articles); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: problem with writing json encoded struct http.ResponseWriter")
	}*/

}
//TODO: Returns all cards from myCards collection
func returnAllCardEntries(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Endpoint Hit: returnAllCardEntries")

	if err := json.NewEncoder(w).Encode(Articles); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: problem with writing json encoded struct http.ResponseWriter")
	}
}
//TODO: Returns one card from myCards collection
func returnSingleCardEntry(w http.ResponseWriter, r *http.Request){
	log.Info().Msg("Endpoint Hit: returnSingleCardEntry")

	vars := mux.Vars(r)
	cardName := vars["cardName"]
	setName := vars["setName"]

	card, _ := SingleCardInfo(cardName,setName,"allCards")
	resp, _ := json.Marshal(card)
	_,_ = w.Write(resp)
	/*if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: Couldn't receive single card info from db")
	}
	_,_ =fmt.Fprint(w,singleCard)*/

}
//TODO: updates one card from myCards collection
func updateSingleCardEntry(w http.ResponseWriter, r *http.Request){
	log.Info().Msg("Endpoint Hit: updateSingleCardEntry")
	var updatedArticle Article
	vars := mux.Vars(r)
	id := vars["id"]
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: problem with reading request body")
	}
	if err = json.Unmarshal(reqBody, &updatedArticle);err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't unmarshal reqBody json into article struct")
	}

	for _,article := range Articles {
		if article.Id == id {
			article.Content = updatedArticle.Content
			article.Desc = updatedArticle.Desc
			article.Title = updatedArticle.Title
		}
	}

}
//TODO: deletes one card from myCards collection
func deleteSingleCardEntry(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Endpoint Hit: deleteSingleCardEntry")
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles{
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}