package main

import (
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/mongodb"
)

func InsertImgInfo(imgInfo Img, dbName string, dbCollection string) error {
	client, ctx, cancelCtx, err := mongodb.CreateClient()
	if err != nil {
		return err
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Error().Timestamp().Err(err).Msg("Error: closing client\n")
		}
		cancelCtx()
	}()

	collection := client.Database(dbName).Collection(dbCollection)

	insertResult, err := collection.InsertOne(ctx, imgInfo)
	if err != nil {
		return err
	}

	log.Info().Msgf("Success: insertion result:\n", insertResult)

	return err
}
