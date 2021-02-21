package main

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	var port = ":8081"
	log.Info().Msgf("Serving static files on port %v", port)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal().Err(err)
	}
}
