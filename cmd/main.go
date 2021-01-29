package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/cli"
)

func main() {
	language, cardName, err := cli.ReadFromCLI()
	if err != nil {
		log.Fatal().Err(err)
	}

	fmt.Println(language, cardName)

	/*URL := importer.URLGenerator(language, cardName)

	cardInfo, err := importer.RequestCardInfo(URL)
	if err != nil {
		log.Fatal().Err(err)
	}

	client, ctx, err := mongodb.CreateClient()
	if err != nil {
		log.Fatal().Err(err)
	}

	err = mongodb.InsertCardInfo(cardInfo, client, ctx)
	if err != nil {
		log.Fatal().Err(err)
	}

	if err = mongodb.GetAllCardInfo(client, ctx); err != nil{
		log.Fatal().Err(err)
	}

	if err = mongodb.GetFilteredSingleCardInfo(client, ctx); err != nil{
		log.Fatal().Err(err)
	}*/
}