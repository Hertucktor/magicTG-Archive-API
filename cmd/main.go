package main

import (
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/cli"
	"magicTGArchive/internal/pkg/importer"
	"magicTGArchive/internal/pkg/mongodb"
)

func main() {
	language, cardName, err := cli.ReadFromCLI()
	if err != nil {
		log.Fatal().Err(err)
	}

	URL := importer.URLGenerator(language, cardName)

	cardInfo, err := importer.RequestCardInfo(URL)
	if err != nil {
		log.Fatal().Err(err)
	}

	client, err := mongodb.CreateClient()
	if err != nil {
		log.Fatal().Err(err)
	}

	err = mongodb.InsertCardInfo(cardInfo, client)
	if err != nil {
		log.Fatal().Err(err)
	}

	//for _, card := range card.Cards{
	//	fmt.Println(card.ImageURL)
	//}
}
