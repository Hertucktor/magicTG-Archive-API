package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

func setupRoutes(){
	var port = ":8080"
	log.Info().Msgf("Starting API on port %v", port)

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	//Status calls
	status := myRouter.PathPrefix("/status").Subrouter()
	status.HandleFunc("/alive",statusAlive).Methods(http.MethodGet)
	status.HandleFunc("/check",statusCheck).Methods(http.MethodGet)

	//CRUD Operations for card info
	api := myRouter.PathPrefix("/api").Subrouter()
	api.Use(mux.CORSMethodMiddleware(api), corsOriginMiddleware)
	api.HandleFunc("/card", createNewCardEntry).Methods(http.MethodPost)
	api.HandleFunc("/cards", returnAllCardEntries).Methods(http.MethodGet) //From allCards Coll
	api.HandleFunc("/cards/set-names/{setName}", returnAllCardsBySet).Methods(http.MethodGet) //From allCards Coll
	api.HandleFunc("/cards/collector-numbers/{number}/set-names/{setName}", returnSingleCardEntry).Methods(http.MethodGet) //From myCards Coll
	api.HandleFunc("/cards/collector-number/{number}/set-names/{setName}", updateSingleCardEntry).Methods(http.MethodPut)
	api.HandleFunc("/cards/collector-number/{number}/set-names/{setName}", deleteSingleCardEntry).Methods(http.MethodDelete)

	//API Operations for img info
	img := myRouter.PathPrefix("/img").Subrouter()
	img.HandleFunc("/set-names/{setName}", returnSingleImg).Methods(http.MethodGet)
	img.HandleFunc("/upload", uploadImg).Methods(http.MethodPost)

	//Serve UI
	staticDir := "/static/"
	myRouter.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	//Open http connection
	if err := http.ListenAndServe(port, myRouter); err != nil {
		log.Panic().Timestamp().Err(err).Msg("Panic: problem with TCP network connection")
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Header().Add("Content-Type", "application/json")
	_, err := w.Write([]byte("\"There is nothing here\""))
	if err != nil {
		log.Err(err).Msg("Cannot write http 404 response body")
	}
}