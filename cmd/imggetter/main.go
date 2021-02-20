package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"magicTGArchive/internal/pkg/imggetter"
)

var dbCollName = "setImages"

func main() {
	var img imggetter.Img
	var err error
	var imgBaseURL = "https://media.magic.wizards.com/images/featured/"
	var setName = "Modern Horizons"

	/*if err := InsertSetImg(imgAttr, dbCollName); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't insert ImgData into db")
	}*/

	if img, err = ReturnImgUrl(setName); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't return Img info")
	}

	requestURL := imgBaseURL + img.PicName + "." + img.Extension

	fmt.Println(requestURL)
}

func ReturnImgUrl(setName string) (imggetter.Img, error) {
	var attributes imggetter.Img
	var singleResult bson.M
	var err error

	if singleResult, err = imggetter.SingleSetImg(setName, dbCollName); err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't read attributes from db")
		return attributes, err
	}

	marshalled , err := json.Marshal(singleResult)
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't marshal")
		return attributes, err
	}

	if err = json.Unmarshal(marshalled, &attributes); err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't unmarshal")
		return attributes, err
	}

	return attributes, err
}