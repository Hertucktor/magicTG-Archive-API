package importer

import (
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/env"
	"magicTGArchive/internal/pkg/mongodb"
)

func InsertImportCard(cardInfo Card) error {
	var conf env.Conf
	var err error
	cardInfo.Quantity = 1

	if conf, err = env.ReceiveEnvVars(); err != nil {
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

	collection := client.Database(conf.DbName).Collection(conf.DbCollAllCards)
	log.Info().Timestamp().Msgf("Successful: created collection:\n", collection)

	insertResult, err := collection.InsertOne(ctx, cardInfo)
	if err != nil {
		log.Error().Timestamp().Err(err).Msgf("Error: couldn't insert into collection of db:\n", conf.DbCollAllCards, conf.DbName)
		return err
	}
	log.Info().Msgf("Success: insertion result:\n", insertResult)

	return err
}