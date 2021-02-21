package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"magicTGArchive/internal/pkg/imgHandler"
	"net/http"
)

var dbCollName = "setImages"

func returnSingleImg(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Endpoint Hit: returnSingleImg")
	var img imgHandler.Img
	var readResult bson.M
	var err error
	var imgBaseURL = "https://media.magic.wizards.com/images/featured/"

	vars := mux.Vars(r)
	setName := vars["setName"]

	readResult, err = imgHandler.SingleSetImg(setName, dbCollName)

	marshalled, err := json.Marshal(readResult)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't marshal")
	}

	if err = json.Unmarshal(marshalled, &img); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't unmarshal")
	}

	requestURL := imgBaseURL + img.PicName + "." + img.Extension

	marshalled, err = json.Marshal(requestURL)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't marshal")
	}

	if _, err = w.Write(marshalled); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't write url to user")
	}
}

func uploadImg(w http.ResponseWriter, r *http.Request){
	log.Info().Msgf("Endpoint Hit: uploadImg with method:", r.Method)

	fmt.Fprintf(w, "Uploading File")
}