package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/config"
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
		log.Fatal().Timestamp().Err(err).Msg("")
	}

	conf, err := config.GetConfig("config.yml")
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("")
	}

	imgInfo, err := SingleSetImg(setName,conf.DBName, dbCollName, client, ctx)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("")
		w.WriteHeader(400)
		if _, err = w.Write([]byte("The image link you requested is not in db")); err != nil {
			log.Fatal().Timestamp().Err(err).Msg("")
		}
	}

	imgLinkByte, err := json.Marshal(imgInfo.ImgLink)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("")
		w.WriteHeader(500)
		return
	}

	if _, err = w.Write(imgLinkByte); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("")
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