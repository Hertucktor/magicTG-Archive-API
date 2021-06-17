package main

import (
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/config"
	"magicTGArchive/internal/pkg/setNames"
)

func main() {
	configFile := "config.yml"

	conf, err := config.GetConfig(configFile)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't receive env vars")
	}

	setNames.ReturnAllSetName(conf)
}
