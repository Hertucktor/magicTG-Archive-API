package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"magicTGArchive/internal/pkg/config"
	"magicTGArchive/internal/pkg/mongodb"
	"net/http"
)

type RequestBody struct {
	Number string
	SetName string
}

func createNewCardEntryOnAllCardCollection(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Endpoint Hit: createNewCardEntryOnAllCardCollection")
	var requestBody RequestBody

	conf, err := config.GetConfig("config.yml")
	if err != nil{
		log.Fatal().Err(err).Msg("")
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("")
		w.WriteHeader(500)
		return
	}

	if err = json.Unmarshal(reqBody, &requestBody);err != nil {
		log.Fatal().Timestamp().Err(err).Msg("")
		w.WriteHeader(500)
		return
	}

	client, ctx, cancelCtx, err := mongodb.CreateClient()
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("")
	}

	cardInfo, err := ReadSingleCardEntry(requestBody.SetName, requestBody.Number, conf.DBName, conf.DBCollectionAllcards, client, ctx)
	if err != nil {
		w.WriteHeader(400)
		_,_ = w.Write([]byte("The card you requested is not in storage"))
		log.Fatal().Timestamp().Err(err).Msgf("Fatal: couldn't receive card set%v with number%v", requestBody.SetName, requestBody.Number)
		return
	}

	if err = InsertCard(cardInfo, conf.DBName, conf.DBCollectionMycards, client, ctx); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("")
		w.WriteHeader(500)
		return
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal().Timestamp().Err(err).Msg("Fatal: closing client\n")
		}
		cancelCtx()
	}()
}

//FIXME: paginate results or db ctx deadline will close connection
func returnAllCardEntriesFromAllCardCollection(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Endpoint Hit: returnAllCardEntriesFromAllCardCollection")

	conf, err := config.GetConfig("config.yml")
	if err != nil{
		log.Fatal().Err(err).Msg("")
	}

	client, ctx, cancelCtx, err := mongodb.CreateClient()
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("")
	}

	allCards, err := AllCards(conf.DBName, conf.DBCollectionAllcards, client, ctx)
	if err != nil {
		w.WriteHeader(400)
		if _, err = w.Write([]byte("The cards you requested are not in storage")); err != nil{
			log.Error().Timestamp().Err(err).Msg("")
		}
		log.Fatal().Timestamp().Err(err).Msg("")
		return
	}

	allCardsBytes, err := json.Marshal(allCards)
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("")
		w.WriteHeader(500)
		return
	}

	if _,err = w.Write(allCardsBytes); err != nil {
		log.Error().Timestamp().Err(err).Msg("")
		w.WriteHeader(500)
		return
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal().Timestamp().Err(err).Msg("Fatal: closing client\n")
		}
		cancelCtx()
	}()
}

func returnAllCardsBySetFromAllCardCollection(w http.ResponseWriter, r *http.Request){
	log.Info().Msg("Endpoint Hit: returnAllCardsBySetFromAllCardCollection")
	conf, err := config.GetConfig("config.yml")
	if err != nil{
		log.Fatal().Err(err).Msg("")
	}

	vars := mux.Vars(r)
	setName := vars["setName"]

	client, ctx, cancelCtx, err := mongodb.CreateClient()
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("")
	}

	//reads all entries by set name from allCards collection
	cardsBySet, err := AllCardsBySet(setName,conf.DBName, conf.DBCollectionAllcards, client, ctx)
	if err != nil {
		w.WriteHeader(400)
		if _, err = w.Write([]byte("The cards you requested are not in storage")); err != nil{
			log.Error().Err(err).Msg("")
		}
		log.Fatal().Timestamp().Err(err).Msg("")
		return
	}

	cardsBySetBytes, err := json.Marshal(cardsBySet)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("")
		w.WriteHeader(500)
		return
	}

	if _, err = w.Write(cardsBySetBytes); err != nil {
		log.Error().Timestamp().Err(err).Msg("")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal().Timestamp().Err(err).Msg("Fatal: closing client\n")
		}
		cancelCtx()
	}()
}

func readFromOwnCollection(w http.ResponseWriter, r *http.Request){
	log.Info().Msg("Endpoint Hit: readFromOwnCollection")

	vars := mux.Vars(r)
	setName := vars["setName"]
	number := vars["number"]

	conf, err := config.GetConfig("config.yml")
	if err != nil{
		log.Fatal().Err(err).Msg("")
	}

	client, ctx, cancelCtx, err := mongodb.CreateClient()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("")
	}

	cardResponse, err := ReadSingleCardEntry(setName, number,conf.DBName, conf.DBCollectionMycards, client, ctx)
	if err != nil {
		w.WriteHeader(400)
		if _, err = w.Write([]byte("The card you requested is not in storage")); err != nil{
			log.Error().Err(err).Msg("")
		}
		log.Fatal().Timestamp().Err(err).Msg("")
		return
	}

	response, err := json.Marshal(cardResponse)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("")
		w.WriteHeader(500)
		return
	}

	if _,err = w.Write(response); err != nil {
		log.Error().Timestamp().Err(err).Msg("")
		w.WriteHeader(500)
		return
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal().Timestamp().Err(err).Msg("Fatal: closing client\n")
		}
		cancelCtx()
	}()
}

func updateSingleCardFromOwnCollection(w http.ResponseWriter, r *http.Request){
	log.Info().Msg("Endpoint Hit: updateSingleCardFromOwnCollection")
	var card mongodb.Card

	vars := mux.Vars(r)
	setName := vars["setName"]
	number := vars["number"]

	conf, err := config.GetConfig("config.yml")
	if err != nil{
		log.Fatal().Err(err).Msg("")
	}

	client, ctx, cancelCtx, err := mongodb.CreateClient()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: creating client\n")
	}

	//reads one entry from myCards collection
	if _, err = ReadSingleCardEntry(setName, number, conf.DBName, conf.DBCollectionMycards, client, ctx); err != nil {
		w.WriteHeader(404)
		if _,err = w.Write([]byte("The card you requested is not in storage")); err != nil {
			log.Error().Timestamp().Err(err).Msg("")
		}
		log.Fatal().Timestamp().Err(err).Msg("")
		return
	}

	if err = UpdateSingelCard(setName, number, card.Quantity, conf.DBName, conf.DBCollectionMycards, client, ctx); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("")
		w.WriteHeader(500)
		return
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal().Timestamp().Err(err).Msg("Fatal: closing client\n")
		}
		cancelCtx()
	}()
}

func deleteSingleCardFromOwnCollection(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Endpoint Hit: deleteSingleCardFromOwnCollection")

	vars := mux.Vars(r)
	setName := vars["setName"]
	number := vars["number"]

	conf, err := config.GetConfig("config.yml")
	if err != nil{
		log.Fatal().Err(err).Msg("")
	}

	client, ctx, cancelCtx, err := mongodb.CreateClient()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: creating client\n")
	}

	if _, err = DeleteSingleCard(setName, number, conf.DBName, conf.DBCollectionMycards, client, ctx); err != nil {
		log.Error().Err(err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	if _,err = w.Write([]byte("Deletion successful!")); err != nil {
		log.Error().Err(err)
		w.WriteHeader(500)
		return
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal().Timestamp().Err(err).Msg("Fatal: closing client\n")
		}
		cancelCtx()
	}()
}