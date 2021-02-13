package main

import (
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
//TODO: serves UI
func homePage(w http.ResponseWriter, r *http.Request){
	if _, err := fmt.Fprintf(w, "Welcome to the HomePage!"); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't serve string on homepage")
	}
	fmt.Println("Endpoint Hit: homePage")
}