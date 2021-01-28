package main

import (
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/cli"
)

func main(){
	_, err := cli.ReadFromCLI()
	if err != nil {
		log.Fatal().Err(err)
	}
}