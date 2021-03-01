package imgHandler

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"magicTGArchive/internal/pkg/env"
)

func SingleSetImg(setName string, dbCollection string, client *mongo.Client, ctx context.Context) (Img, error) {
	var readFilter = bson.M{"setname": setName}
	var imgInfo Img

	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive env vars")
		return imgInfo, err
	}

	collection := client.Database(conf.DbName).Collection(dbCollection)
	log.Info().Timestamp().Msgf("Success: created collection: %v", collection.Name())

	if err = collection.FindOne(ctx, readFilter).Decode(&imgInfo); err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't find document")
	}

	return imgInfo, err
}