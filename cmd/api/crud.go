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

func createNewCardEntry(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Endpoint Hit: createNewCardEntry")
	var reqCard ReqCard
	var card mongodb.DBCard

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: problem with reading request body")
	}

	if err = json.Unmarshal(reqBody, &reqCard);err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't unmarshal reqBody json into article struct")
	}

	results, err := mongodb.SingleCardInfo(reqCard.Name, reqCard.SetName, "allCards")
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't receive reqCard")
	}
	fmt.Println(results)

	response , err := json.Marshal(results[0])
	if err != nil {
		log.Fatal().Err(err)
	}

	if _,err = w.Write(response); err != nil {
		log.Fatal().Err(err)
	}

	if err = json.Unmarshal(response, &card); err != nil {
		log.Fatal().Err(err)
	}

	if err = mongodb.InsertCard(card,"myCards"); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't insert reqCard into db")
	}
}
//FIXME: paginate results or db will struggle over time
func returnAllCardEntries(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Endpoint Hit: returnAllCardEntries")

	results, err := mongodb.AllCardInfo("myCards")
	response , err := json.Marshal(results)
	if err != nil {
		log.Fatal().Err(err)
	}
	if _,err = w.Write(response); err != nil {
		log.Fatal().Err(err)
	}
}
//FIXME: Return only one card from myCards collection
func returnSingleCardEntry(w http.ResponseWriter, r *http.Request){
	log.Info().Msg("Endpoint Hit: returnSingleCardEntry")

	vars := mux.Vars(r)
	cardName := vars["cardName"]
	setName := vars["setName"]

	results, err := mongodb.SingleCardInfo(cardName, setName, "myCards")
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't receive reqCard")
	}

	response , err := json.Marshal(results)
	if err != nil {
		log.Fatal().Err(err)
	}

	if _,err = w.Write(response); err != nil {
		log.Fatal().Err(err)
	}

}

func updateSingleCardEntry(w http.ResponseWriter, r *http.Request){
	log.Info().Msg("Endpoint Hit: updateSingleCardEntry")
	var card mongodb.DBCard

	vars := mux.Vars(r)
	cardName := vars["cardName"]
	setName := vars["setName"]

	results, err := mongodb.SingleCardInfo(cardName, setName, "myCards")
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't receive reqCard")
	}

	response , err := json.Marshal(results[0])
	if err != nil {
		log.Fatal().Err(err)
	}

	if _,err = w.Write(response); err != nil {
		log.Fatal().Err(err)
	}

	if err = json.Unmarshal(response, &card); err != nil {
		log.Fatal().Err(err)
	}

	if err = mongodb.UpdateSingleCard(cardName,setName,card.Quantity,"myCards"); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't update card entry")
	}

}

func deleteSingleCardEntry(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Endpoint Hit: deleteSingleCardEntry")
	vars := mux.Vars(r)
	cardName := vars["cardName"]
	setName := vars["setName"]

	result, err := mongodb.DeleteSingleCard(cardName, setName, "myCards")
	if err != nil {
		log.Fatal().Err(err)
	}
	_,_ = fmt.Fprint(w, result)
}