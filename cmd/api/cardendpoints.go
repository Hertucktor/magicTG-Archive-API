package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"magicTGArchive/internal/pkg/mongodb"
	"net/http"
)

type RequestBody struct {
	Number string
	SetName string
}

func createNewCardEntry(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Endpoint Hit: createNewCardEntry")
	var requestBody RequestBody

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Fatal: problem with reading request requestBody")
		w.WriteHeader(500)
		return
	}

	if err = json.Unmarshal(reqBody, &requestBody);err != nil {
		log.Error().Timestamp().Err(err).Msg("Fatal: couldn't unmarshal reqBody json into article struct")
		w.WriteHeader(500)
		return
	}

	client, ctx, cancelCtx, err := mongodb.CreateClient()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: creating client\n")
	}

	//read from allCards collection
	cardInfo, err := SingleCardInfo(requestBody.SetName, requestBody.Number, "allCards", client, ctx)
	if err != nil {
		w.WriteHeader(400)
		_,_ = w.Write([]byte("The card you requested is not in storage"))
		log.Error().Timestamp().Err(err).Msgf("Fatal: couldn't receive card set%v with number%v", requestBody.SetName, requestBody.Number)
		return
	}

	//insert into myCards collection
	if err = InsertCard(cardInfo,"myCards", client, ctx); err != nil {
		log.Error().Timestamp().Err(err).Msg("Fatal: couldn't insert requestBody into db")
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
func returnAllCardEntries(w http.ResponseWriter, r *http.Request) {

	log.Info().Msg("Endpoint Hit: returnAllCardEntries")

	client, ctx, cancelCtx, err := mongodb.CreateClient()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: creating client\n")
	}

	//read all entries out of allCards collection
	allCards, err := AllCards("allCards", client, ctx)
	if err != nil {
		w.WriteHeader(400)
		_,_ = w.Write([]byte("The cards you requested are not in storage"))
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive cards")
		return
	}

	allCardsBytes, err := json.Marshal(allCards)
	if err != nil {
		log.Error().Timestamp().Err(err)
		w.WriteHeader(500)
		return
	}

	if _,err = w.Write(allCardsBytes); err != nil {
		log.Error().Timestamp().Err(err)
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

func returnAllCardsBySet(w http.ResponseWriter, r *http.Request){
	log.Info().Msg("Endpoint Hit: returnAllCardsBySet")

	vars := mux.Vars(r)
	setName := vars["setName"]

	client, ctx, cancelCtx, err := mongodb.CreateClient()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: creating client\n")
	}

	//reads all entries by set name from allCards collection
	cardsBySet, err := AllCardsBySet(setName, "allCards", client, ctx)
	if err != nil {
		w.WriteHeader(400)
		_,_ = w.Write([]byte("The cards you requested are not in storage"))
		log.Error().Timestamp().Err(err).Msgf("Error: couldn't return cards from set: %v",setName)
		return
	}

	cardsBySetBytes, err := json.Marshal(cardsBySet)
	if err != nil {
		log.Error().Timestamp().Err(err)
		w.WriteHeader(500)
		return
	}

	if _,err = w.Write(cardsBySetBytes); err != nil {
		log.Error().Timestamp().Err(err)
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

func returnSingleCardEntry(w http.ResponseWriter, r *http.Request){
	log.Info().Msg("Endpoint Hit: returnSingleCardEntry")

	vars := mux.Vars(r)
	setName := vars["setName"]
	number := vars["number"]

	client, ctx, cancelCtx, err := mongodb.CreateClient()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: creating client\n")
	}

	//reads one entry from myCards collection
	cardResponse, err := SingleCardInfo(setName, number, "myCards", client, ctx)
	if err != nil {
		w.WriteHeader(400)
		_,_ = w.Write([]byte("The card you requested is not in storage"))
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive reqCard for return single card")
		return
	}

	response, err := json.Marshal(cardResponse)
	if err != nil {
		log.Error().Timestamp().Err(err)
		w.WriteHeader(500)
		return
	}

	if _,err = w.Write(response); err != nil {
		log.Error().Timestamp().Err(err)
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

func updateSingleCardEntry(w http.ResponseWriter, r *http.Request){
	log.Info().Msg("Endpoint Hit: updateSingleCardEntry")
	var card mongodb.Card

	vars := mux.Vars(r)
	setName := vars["setName"]
	number := vars["number"]

	client, ctx, cancelCtx, err := mongodb.CreateClient()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: creating client\n")
	}

	//reads one entry from myCards collection
	_, err = SingleCardInfo(setName, number, "myCards", client, ctx)
	if err != nil {
		w.WriteHeader(404)
		_,_ = w.Write([]byte("The card you requested is not in storage"))
		log.Error().Timestamp().Err(err).Msg("Fatal: couldn't receive reqCard for update single card")
		return
	}

	//update one entry in myCards collection
	if err = UpdateSingleCard(setName, number, card.Quantity,"myCards", client, ctx); err != nil {
		log.Error().Timestamp().Err(err).Msg("Fatal: couldn't update card entry")
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

func deleteSingleCardEntry(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Endpoint Hit: deleteSingleCardEntry")
	vars := mux.Vars(r)
	setName := vars["setName"]
	number := vars["number"]

	client, ctx, cancelCtx, err := mongodb.CreateClient()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: creating client\n")
	}
	//deletes one entry from myCards collection
	_, err = DeleteSingleCard(setName, number, "myCards", client, ctx)
	if err != nil {
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