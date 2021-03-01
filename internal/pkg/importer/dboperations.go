package importer

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"magicTGArchive/internal/pkg/env"
)

func InsertImportCard(cardInfo Card, client *mongo.Client, ctx context.Context) error {
	var conf env.Conf
	var err error
	cardInfo.Quantity = 1

	if conf, err = env.ReceiveEnvVars(); err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive env vars")
		return err
	}

	collection := client.Database(conf.DbName).Collection(conf.DbCollAllCards)
	log.Info().Timestamp().Msgf("Successful: connected to collection:%v", collection.Name())

	insertResult, err := collection.InsertOne(ctx, cardInfo)
	if err != nil {
		log.Error().Timestamp().Err(err).Msgf("Error: couldn't insert into collection of db:\n", conf.DbCollAllCards, conf.DbName)
		return err
	}
	log.Info().Msgf("Success: insertion result: %v", insertResult.InsertedID)

	return err
}