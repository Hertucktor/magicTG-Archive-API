package main

import (
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/env"
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
	envVars, err := env.ReceiveEnvVars()

	imgInfos, err := ConvertCSVEntriesIntoStruct(filePath)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't ")
	}

	for _, imgInfo := range imgInfos.Imgs{
		if err := InsertImgInfo(imgInfo, envVars.DbCollImgInfo); err != nil {
			log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't insert ImgData into db")
		}
	}

}