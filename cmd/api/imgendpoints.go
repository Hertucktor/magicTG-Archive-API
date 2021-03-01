package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/imgHandler"
	"net/http"
)

var dbCollName = "setImages"

func returnSingleImg(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Endpoint Hit: returnSingleImg")

	vars := mux.Vars(r)
	setName := vars["setName"]

	imgInfo, err := imgHandler.SingleSetImg(setName, dbCollName)
	if err != nil {
		log.Error().Timestamp().Err(err)

	}

	marshalled, err := json.Marshal(imgInfo.ImgLink)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't marshal")
	}


	if _, err = w.Write(marshalled); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't write url to user")
	}
}