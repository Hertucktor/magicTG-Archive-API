package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/importer"
	"magicTGArchive/internal/pkg/mongodb"
	"strconv"
)

func main() {
	/*var URL string
	var multiverseID string

	language, cardName, setName, err := cli.ReadFromCLI()
	if err != nil {
		log.Fatal().Err(err)
	}

	URL = importer.URLByCardNameAndLanguage(language, cardName)
	fmt.Println(URL)

	multipleCards, err := importer.RequestMultipleInfosOfOneCard(URL)
	if err != nil {
		log.Fatal().Err(err)
	}

	for _, card := range multipleCards.Cards{
		if card.SetName == setName {
			multiverseID = strconv.Itoa(card.Multiverseid)
		}
	}
	URL = importer.URLForMultiverserID(multiverseID)

	singleCard, err := importer.RequestOneCard(URL)
	if err != nil {
		log.Fatal().Err(err)
	}

	if err := database.InsertDataset(singleCard); err != nil {
		log.Fatal().Err(err)

	}*/
	var page int
	var delimiter = 1
	var collectionName = "allCards"

	for delimiter != 0{
		resp, err := importer.RequestAllCards(strconv.Itoa(page))
		if err != nil {
			log.Fatal().Timestamp().Err(err).Msg("Fatal: API Request threw error")
		}

		for _, card := range resp.Cards{
			if err := mongodb.InsertCard(card, collectionName); err != nil {
				log.Fatal().Timestamp().Err(err).Msgf("Error: couldn't insert dataset:\n",card)
			}
		}
		page++
		fmt.Println(page)
		delimiter = len(resp.Cards)
	}


}