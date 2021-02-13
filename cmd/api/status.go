package main

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

func statusAlive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("\"OK\"")); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't write status alive message back to user")
	}
}

func statusCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	body := struct {
		ResponseCode       int    `json:"-"`
	}{
		ResponseCode:       http.StatusOK,
	}

	marshalledObject, err := json.Marshal(body)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Could not marshal body")
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte("something bad happened, please contact the administrator")); err != nil {
			log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't write error message back to user")
		}
		return
	}

	w.WriteHeader(body.ResponseCode)
	if _, err = w.Write(marshalledObject); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't write info response back to user")
	}
}