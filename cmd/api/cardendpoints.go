package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"magicTGArchive/internal/pkg/mongodb"
	"net/http"
)

type ReqCard struct {
	Number string
	SetName string
}

func createNewCardEntry(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Endpoint Hit: createNewCardEntry")
	var reqCard ReqCard
	var card mongodb.Card

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Fatal: problem with reading request body")
	}

	if err = json.Unmarshal(reqBody, &reqCard);err != nil {
		log.Error().Timestamp().Err(err).Msg("Fatal: couldn't unmarshal reqBody json into article struct")
	}

	client, ctx, cancelCtx, err := mongodb.CreateClient()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: creating client\n")
	}

	//read from allCards collection
	results, err := SingleCardInfo(reqCard.SetName, reqCard.Number, "allCards", client, ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_,_ = w.Write([]byte("The card you requested is not in storage"))
		log.Error().Timestamp().Err(err).Msg("Fatal: couldn't receive reqCard for create new entry")
		return
	}

	response , err := json.Marshal(results)
	if err != nil {
		log.Error().Err(err)
	}

	if _,err = w.Write(response); err != nil {
		log.Error().Err(err)
	}

	if err = json.Unmarshal(response, &card); err != nil {
		log.Error().Err(err)
	}
	//insert into myCards collection
	if err = InsertCard(card,"myCards", client, ctx); err != nil {
		log.Error().Timestamp().Err(err).Msg("Fatal: couldn't insert reqCard into db")
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
	results, err := AllCards("allCards", client, ctx)

	response , err := json.Marshal(results)
	if err != nil {
		log.Error().Timestamp().Err(err)
	}

	if _,err = w.Write(response); err != nil {
		log.Error().Timestamp().Err(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal().Timestamp().Err(err).Msg("Fatal: closing client\n")
		}
		cancelCtx()
	}()
}

func returnAllCardsBySet(w http.ResponseWriter, r *http.Request){
	log.Info().Msg("Endpoint Hit: returnSingleCardEntry")

	vars := mux.Vars(r)
	setName := vars["setName"]

	client, ctx, cancelCtx, err := mongodb.CreateClient()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: creating client\n")
	}

	//reads all entries by set name from allCards collection
	results, err := AllCardsBySet(setName, "allCards", client, ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_,_ = w.Write([]byte("The card you requested is not in storage"))
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive reqCard for return single card")
		return
	}

	response , err := json.Marshal(results)
	if err != nil {
		log.Error().Timestamp().Err(err)
	}
	w.WriteHeader(200)
	if _,err = w.Write(response); err != nil {
		log.Error().Timestamp().Err(err)
	}

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
		w.WriteHeader(http.StatusBadRequest)
		_,_ = w.Write([]byte("The card you requested is not in storage"))
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive reqCard for return single card")
		return
	}

	response, err := json.Marshal(cardResponse)

	if _,err = w.Write(response); err != nil {
		log.Error().Timestamp().Err(err)
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
	results, err := SingleCardInfo(setName, number, "myCards", client, ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_,_ = w.Write([]byte("The card you requested is not in storage"))
		log.Error().Timestamp().Err(err).Msg("Fatal: couldn't receive reqCard for update single card")
		return
	}

	response , err := json.Marshal(results)
	if err != nil {
		log.Error().Err(err)
	}

	if _,err = w.Write(response); err != nil {
		log.Error().Err(err)
	}

	if err = json.Unmarshal(response, &card); err != nil {
		log.Error().Err(err)
	}
	//update one entry in myCards collection
	if err = UpdateSingleCard(setName, number, card.Quantity,"myCards", client, ctx); err != nil {
		log.Error().Timestamp().Err(err).Msg("Fatal: couldn't update card entry")
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
	//reads one entry from myCards collection
	result, err := DeleteSingleCard(setName, number, "myCards", client, ctx)
	if err != nil {
		log.Error().Err(err)
		return
	}

	_,_ = fmt.Fprint(w, result)

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal().Timestamp().Err(err).Msg("Fatal: closing client\n")
		}
		cancelCtx()
	}()
}