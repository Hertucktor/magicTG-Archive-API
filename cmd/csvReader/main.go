package main

import (
	"github.com/rs/zerolog/log"
)

var dbCollName = "imgInfo"
var filePath = "./csv/mtgSetIcons.csv"

type ImgCollection struct {
	Imgs []Img
}
type Img struct {
	SetName string
	ImgLink string
}

func main() {
	imgInfos, err := ConvertCSVEntriesIntoStruct()
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't ")
	}

	for _, imgInfo := range imgInfos.Imgs{
		if err := InsertImgInfo(imgInfo, dbCollName); err != nil {
			log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't insert ImgData into db")
		}
	}

}