package main

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Info().Msg("Listening on :8081...")

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal().Err(err)
	}
}


