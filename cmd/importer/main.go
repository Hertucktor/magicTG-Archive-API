package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/importer"
	"strconv"
)

func main() {
	var page int
	var delimiter = 1
	var collectionName = "allCards"

	for delimiter != 0{
		resp, err := importer.RequestAllCards(strconv.Itoa(page))
		if err != nil {
			log.Fatal().Timestamp().Err(err).Msg("Fatal: API Request threw error")
		}

		for _, card := range resp.Cards{
			if err := importer.InsertCard(card, collectionName); err != nil {
				log.Fatal().Timestamp().Err(err).Msgf("Error: couldn't insert dataset:\n",card)
			}
		}
		page++
		fmt.Println(page)
		delimiter = len(resp.Cards)
	}
}
