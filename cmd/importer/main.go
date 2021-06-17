package main

import (
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/config"
	"magicTGArchive/internal/pkg/importer"
)


func main() {
	var page = 1
	var delimiter = 1
	configFile := "config.yml"

	conf, err := config.GetConfig(configFile)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("")
	}

	importer.ImportCardsIntoDatabase(conf, page, delimiter)

}
