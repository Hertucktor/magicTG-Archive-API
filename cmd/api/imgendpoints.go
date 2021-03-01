package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/imgHandler"
	"magicTGArchive/internal/pkg/mongodb"
	"net/http"
)

func returnSingleImg(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Endpoint Hit: returnSingleImg")
	var dbCollName = "imgInfo"

	vars := mux.Vars(r)
	setName := vars["setName"]

	client, ctx, cancelCtx, err := mongodb.CreateClient()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: creating client\n")
	}

	imgInfo, err := imgHandler.SingleSetImg(setName, dbCollName, client, ctx)
	if err != nil {
		log.Error().Timestamp().Err(err)
		w.WriteHeader(400)
		_,_ = w.Write([]byte("The image link you requested is not in db"))
	}

	imgLinkByte, err := json.Marshal(imgInfo.ImgLink)
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't marshal")
		w.WriteHeader(500)
		return
	}

	if _, err = w.Write(imgLinkByte); err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't write url to user")
		w.WriteHeader(500)
		return
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Error().Timestamp().Err(err).Msg("Error: closing client\n")
		}
		cancelCtx()
	}()
}