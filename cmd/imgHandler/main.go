package main

import (
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/imgHandler"
)

var dbCollName = "setImages"

func main() {
	var img imgHandler.Img
	//var imgBaseURL = "https://media.magic.wizards.com/images/featured/"


	if err := imgHandler.InsertSetImg(img, dbCollName); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't insert ImgData into db")
	}
}

func readCSV(){

}