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

	client, ctx, err := mongodb.CreateClient()
	if err != nil {
		log.Fatal().Err(err)
	}

	for _,cards := range cardInfo.Cards{
		/*cards.Quantity = 1
		if err = mongodb.InsertCard(cards, client, ctx); err != nil {
			log.Fatal().Err(err)
		}*/
		if err = mongodb.UpdateSingleCard(cards.Name, client, ctx); err != nil {
			log.Fatal().Err(err)
		}
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