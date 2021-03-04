package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"magicTGArchive/internal/pkg/env"
	"magicTGArchive/internal/pkg/mongodb"
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
	api.HandleFunc("/cards/set-names", returnAllSetName).Methods(http.MethodGet)

	//API Operations for img info
	img := myRouter.PathPrefix("/img").Subrouter()
	img.HandleFunc("/set-names/{setName}", returnSingleImg).Methods(http.MethodGet)

	//Routes for uploads
	upload := myRouter.PathPrefix("/uploads").Subrouter()
	upload.HandleFunc("/img", uploadImg).Methods(http.MethodPost)

	//Serve UI
	staticDir := "/static/"
	myRouter.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	//Open http connection
	if err := http.ListenAndServe(port, myRouter); err != nil {
		log.Panic().Timestamp().Err(err).Msg("Panic: problem with TCP network connection")
	}
}

func returnAllSetName(w http.ResponseWriter, r *http.Request){
	log.Info().Msg("Endpoint Hit: returnAllCardsBySet")

	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't receive env vars")
	}

	client, ctx, err := buildClient(conf)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("v: couldn't build client")
	}

	setNames, err := allSetNames(client, ctx, conf)
	if err != nil {
		w.WriteHeader(500)
		_,_ = w.Write([]byte("set names couldn't been found"))
		log.Error().Timestamp().Err(err).Msg("Error: couldn't return set names")
		return
	}

	setNamesBytes, err := json.Marshal(setNames)
	if err != nil {
		log.Error().Timestamp().Err(err)
		w.WriteHeader(500)
		return
	}

	if _,err = w.Write(setNamesBytes); err != nil {
		log.Error().Timestamp().Err(err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)

}

func allSetNames(client *mongo.Client, ctx context.Context, conf env.Conf)([]string, error){
	var filter = bson.M{}
	var cards []mongodb.Card

	collection := client.Database(conf.DbName).Collection(conf.DbCollAllCards)
	log.Info().Timestamp().Msgf("Successful: created collection:\n", collection)

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Error().Timestamp().Err(err).Msgf("Error: ")
		return nil, err
	}

	if err = cursor.All(ctx, &cards); err != nil {
		log.Error().Timestamp().Err(err).Msgf("Error: couldn't decode data into interface:\n")
		return nil, err
	}

	setNames := sortSetNames(cards)

	defer func() {
		if err = cursor.Close(ctx); err != nil {
			log.Error().Timestamp().Err(err).Msgf("Error: couldn't close cursor:%v", cursor.Current)
		}
		log.Info().Timestamp().Msg("Gracefully closed cursor")
	}()

	return setNames, err
}

func buildClient(conf env.Conf)(*mongo.Client, context.Context, error){

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://"+conf.DbUser+":"+conf.DbPass+"@"+conf.DbPort+"/"+conf.DbName))
	if err != nil {
		log.Error().Err(err)
		return client, nil, err
	}

	ctx := context.TODO()
	if err = client.Connect(ctx); err != nil {
		log.Error().Err(err)
		return client, ctx, err
	}

	return client, ctx, err
}

func sortSetNames(cards []mongodb.Card) ([]string){
	var setNames []string
	for _, card := range cards {
		if !stringInSlice(card.SetName, setNames) {
			setNames = append(setNames, card.SetName)
		}
	}
	return setNames
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Header().Add("Content-Type", "application/json")
	_, err := w.Write([]byte("\"There is nothing here\""))
	if err != nil {
		log.Err(err).Msg("Cannot write http 404 response body")
	}
}