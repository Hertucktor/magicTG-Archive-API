package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/cli"
	"magicTGArchive/internal/pkg/importer"
)

func main() {

	language, cardName, err := cli.ReadFromCLI()
	if err != nil {
		log.Fatal().Err(err)
	}

	URL := importer.URLGenerator(language, cardName)
	fmt.Println(URL)
	cardInfo, err := importer.RequestCardInfo(URL)
	if err != nil {
		log.Fatal().Err(err)
	}

	for _, cards := range cardInfo.Cards{
		for _, names := range cards.ForeignNames {
			fmt.Println(names.Language)
		}
	}
	/*
	client, ctx, err := mongodb.CreateClient()
	if err != nil {
		log.Fatal().Err(err)
	}


	if err = mongodb.InsertCardInfo(cardInfo, client, ctx); err != nil {
		log.Fatal().Err(err)
	}

	if err = mongodb.GetAllCardInfo(client, ctx); err != nil{
		log.Fatal().Err(err)
	}

	if err = mongodb.GetFilteredSingleCardInfo(client, ctx); err != nil{
		log.Fatal().Err(err)
	}*/
}