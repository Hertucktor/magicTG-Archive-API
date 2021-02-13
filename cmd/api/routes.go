package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

func handleRequests(){
	var port = "127.0.0.1:8080"
	log.Info().Msgf("Starting API on port:", port)

	myRouter := mux.NewRouter().StrictSlash(true)

	//Status calls
	status := myRouter.PathPrefix("/status").Subrouter()
	status.HandleFunc("/alive",statusAlive).Methods(http.MethodGet)
	status.HandleFunc("/check",statusCheck).Methods(http.MethodGet)

	//Interface for UI
	ui := myRouter.PathPrefix("/").Subrouter()
	ui.HandleFunc("/", homePage)

	//CRUD Operations
	api := myRouter.PathPrefix("/api").Subrouter()
	api.HandleFunc("/card", createNewCardEntry).Methods(http.MethodPost)
	api.HandleFunc("/card/all", returnAllCardEntries).Methods(http.MethodGet)
	api.HandleFunc("/card/number/{number}/set/name/{setName}", returnSingleCardEntry).Methods(http.MethodGet)
	api.HandleFunc("/card/number/{number}/set/name/{setName}", updateSingleCardEntry).Methods(http.MethodPut)
	api.HandleFunc("/card/number/{number}/set/name/{setName}", deleteSingleCardEntry).Methods(http.MethodDelete)

	//Open http connection
	if err := http.ListenAndServe(port, myRouter); err != nil {
		log.Panic().Timestamp().Err(err).Msg("Panic: problem with TCP network connection")
	}
}