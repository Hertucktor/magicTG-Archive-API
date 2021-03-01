package imgHandler

import (
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"magicTGArchive/internal/pkg/env"
	"magicTGArchive/internal/pkg/mongodb"
)

type databaseResponse struct {
	ID string `bson:"_id"`
	ImgLink string `bson:"imglink"`
	SetName string `bson:"setname"`
}

func SingleSetImg(setName string, dbCollection string) (databaseResponse, error) {
	var dbresp databaseResponse
	var readFilter = bson.M{"setname": setName}

	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive env vars")
		return dbresp, err
	}

	client, ctx, cancelCtx, err := mongodb.CreateClient()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: creating client\n")
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Error().Timestamp().Err(err).Msg("Error: closing client\n")
		}
		cancelCtx()
	}()

	collection := client.Database(conf.DbName).Collection(dbCollection)
	log.Info().Timestamp().Msgf("Success: created collection:\n", collection.Name())

	if err = collection.FindOne(ctx, readFilter).Decode(&dbresp); err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't find document")
	}

	return dbresp, err
}