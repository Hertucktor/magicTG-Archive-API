package main

import (
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/config"
	"magicTGArchive/internal/pkg/csvReader"
)

func main() {
	var filePath = "./csv/mtgSetIcons.csv"
	conf, err := config.GetConfig("config.yml")
	if err != nil{
		log.Fatal().Err(err).Msg("")
	}

	csvReader.TransferCSVDataToDatabase(filePath, conf)

}