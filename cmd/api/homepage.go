package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
)

//TODO: serves UI
func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: homePage")

	if _, err := fmt.Fprintf(w, "Welcome to the HomePage!"); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't serve string on homepage")
	}
}