package main

import (
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/config"
)

type ImgCollection struct {
	Imgs []Img
}
type Img struct {
	SetName string
	ImgLink string
}

func main() {
	var filePath = "./csv/mtgSetIcons.csv"
	conf, err := config.GetConfig("config.yml")
	if err != nil{
		log.Fatal().Err(err).Msg("")
	}

	imgInfos, err := ConvertCSVEntriesIntoStruct(filePath)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't ")
	}

	for _, imgInfo := range imgInfos.Imgs{
		if err = InsertImgInfo(imgInfo,conf.DBName, conf.DBCollectionSetimages); err != nil {
			log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't insert ImgData into db")
		}
	}

}