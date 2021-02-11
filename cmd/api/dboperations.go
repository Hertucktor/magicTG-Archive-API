package main

import (
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"magicTGArchive/internal/pkg/env"
	"magicTGArchive/internal/pkg/mongodb"
)

func SingleCardInfo(cardName string, setName string, dbCollection string) (mongodb.DBCard, error) {
	var cardInfoFiltered []mongodb.DBCard
	var singleCard mongodb.DBCard
	var filter = bson.M{"name": cardName, "setname": setName}

	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive env vars")
		return singleCard, err
	}

	client, ctx, cancelCtx, err := mongodb.CreateClient()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: Creating Client\n")
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error().Timestamp().Err(err).Msg("Error: closing client\n")
		}
		cancelCtx()
	}()

	collection := client.Database(conf.DbName).Collection(dbCollection)
	log.Info().Timestamp().Msgf("Success: created collection:\n", collection.Name())

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: cursor couldn't be created\n")
		return singleCard, err
	}

	defer func() {
		if err := cursor.Close(ctx); err != nil {
			log.Error().Timestamp().Err(err).Msgf("Error: couldn't close cursor", cursor)
		}
		log.Info().Msg("Success: Closed cursor\n")
	}()

	if err = cursor.All(ctx, &cardInfoFiltered); err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: problem with the cursor\n")
		return singleCard, err
	}
	log.Info().Timestamp().Msgf("", singleCard)
	for _, card := range cardInfoFiltered {
		singleCard = card
	}

	return singleCard, err
}
