package main

import (
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/cli"
	"magicTGArchive/internal/pkg/importer"
)

func main() {
	language,cardName, err := cli.ReadFromCLI()
	if err != nil {
		log.Fatal().Err(err)
	}
	filter := importer.FilterSelector(language)


}
