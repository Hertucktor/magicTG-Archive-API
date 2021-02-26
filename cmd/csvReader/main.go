package main

import (
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/env"
	"magicTGArchive/internal/pkg/mongodb"
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

func InsertImgInfo(imgInfo Img, dbCollection string) error {
	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive env vars")
		return err
	}

	client, ctx, cancelCtx, err := mongodb.CreateClient()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: Creating Client\n")
		return err
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Error().Timestamp().Err(err).Msg("Error: closing client\n")
		}
		cancelCtx()
	}()

	collection := client.Database(conf.DbName).Collection(dbCollection)
	log.Info().Timestamp().Msgf("Successful: created collection:\n", collection)

	insertResult, err := collection.InsertOne(ctx, imgInfo)
	if err != nil {
		log.Error().Timestamp().Err(err).Msgf("Error: couldn't insert into collection of db:\n", dbCollection, conf.DbName)
		return err
	}

	log.Info().Msgf("Success: insertion result:\n", insertResult)

	return err
}