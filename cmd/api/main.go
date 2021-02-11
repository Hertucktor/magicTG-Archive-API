package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
)
var Item = Card{
	Name:          "",
	ManaCost:      "",
	Cmc:           0,
	Colors:        nil,
	ColorIdentity: nil,
	Type:          "",
	Supertypes:    nil,
	Types:         nil,
	Subtypes:      nil,
	Rarity:        "",
	Set:           "",
	SetName:       "",
	Text:          "",
	Flavor:        "",
	Artist:        "",
	Number:        "",
	Power:         "",
	Toughness:     "",
	Layout:        "",
	Multiverseid:  0,
	ImageURL:      "",
	Rulings:       nil,
	ForeignNames:  nil,
	Printings:     nil,
	OriginalText:  "",
	OriginalType:  "",
	Legalities:    nil,
	ID:            "",
	Variations:    nil,
	Quantity:      0,
}

func main() {
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
	http.HandleFunc("/", homePage)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Panic().Timestamp().Err(err).Msg("Panic: problem with TCP network connection")
	}
}

func returnCard(w http.ResponseWriter, r *http.Request) error{
	log.Info().Msg("Endpoint Hit: returnCard")
	if err := json.NewEncoder(w).Encode(Item); err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: problem with encoding struct to json")
		return err
	}
	return nil
}


