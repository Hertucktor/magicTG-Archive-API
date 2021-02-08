package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/importer"
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

	for delimiter != 0{
		resp, err := importer.RequestAllCards(strconv.Itoa(page))
		if err != nil {
			log.Fatal().Timestamp().Err(err).Msg("Fatal: API Request threw error")
		}
		page++
		fmt.Println(page)
		delimiter = len(resp.Cards)
	}


}