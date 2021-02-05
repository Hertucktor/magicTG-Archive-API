package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/importer"
	"magicTGArchive/internal/pkg/mongodb"
)

func main() {
	/*language, cardName, setName, err := cli.ReadFromCLI()
	if err != nil {
		log.Fatal().Err(err)
	}*/

	var URL string
	var multiverseID int
	language := "en"
	cardName := "Quicksand"
	setName := "Worldwake"

	URL = importer.URLByCardNameAndLanguage(language, cardName)

	multipleCards, err := importer.RequestMultipleInfosOfOneCard(URL)
	if err != nil {
		log.Fatal().Err(err)
	}

	for _, card := range multipleCards.Cards{
		if card.SetName == setName {
			multiverseID = card.Multiverseid
		}
	}
	URL = importer.URLForMultiverserID(multiverseID)

	singleCard, err := importer.RequestOneCard(URL)
	fmt.Println(singleCard)

	if err = mongodb.InsertCard(singleCard.Card); err != nil {
		log.Fatal().Err(err)
	}

	/*
	if err = mongodb.AllCardInfo(client, ctx); err != nil{
		log.Fatal().Err(err)
	}
	if err = mongodb.SingleCardInfo(cardName, client, ctx); err != nil{
		log.Fatal().Err(err)
	}
	if err = mongodb.DeleteSingleCard(cardName, client, ctx); err != nil{
		log.Fatal().Err(err)
	}*/

}