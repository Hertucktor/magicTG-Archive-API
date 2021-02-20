package main

import (
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/importer"
)
var delimiter = 1

func main() {
	var page int

	for delimiter != 0{
		resp, err := importer.RequestAllCards(page)
		if err != nil {
			log.Fatal().Timestamp().Err(err).Msg("Fatal: API Request threw error")
		}

		for _, card := range resp.Cards{
			if err = importer.InsertImportCard(card); err != nil {
				log.Fatal().Timestamp().Err(err).Msgf("Error: couldn't insert dataset:\n",card)
			}
		}
		page++
		log.Info().Timestamp().Msgf("", page)
		delimiter = len(resp.Cards)
	}
}