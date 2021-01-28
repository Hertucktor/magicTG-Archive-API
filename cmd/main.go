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
	filter := importer.FilterSelector(language)

	card, err := importer.RequestCardInfo(cardName, filter)
	if err != nil {
		log.Fatal().Err(err)
	}

	for _, card := range card.Cards{
		fmt.Println(card.Name)
	}
}
