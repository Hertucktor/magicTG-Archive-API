package main

import (
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

	_, err = importer.RequestCardInfo(URL)
	if err != nil {
		log.Fatal().Err(err)
	}

	//for _, card := range card.Cards{
	//	fmt.Println(card.ImageURL)
	//}
}
