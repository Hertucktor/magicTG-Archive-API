package importer

import (
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"magicTGArchive/internal/pkg/env"
	"magicTGArchive/internal/pkg/mongodb"
)

func InsertCard(cardInfo Cards, dbCollection string) error {
	cardInfo.Quantity = 1

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
		if err := client.Disconnect(ctx); err != nil {
			log.Error().Timestamp().Err(err).Msg("Error: closing client\n")
		}
		cancelCtx()
	}()

	collection := client.Database(conf.DbName).Collection(dbCollection)
	log.Info().Timestamp().Msgf("Successful: created collection:\n", collection)

	insertResult, err := collection.InsertOne(ctx, cardInfo)
	if err != nil {
		log.Error().Timestamp().Err(err).Msgf("Error: couldn't insert into collection of db:\n", dbCollection, conf.DbName)
		return err
	}

	log.Info().Msgf("Success: insertion result:\n", insertResult)

	return err
}

func SingleCardInfo(cardName string, dbCollection string) (mongodb.DBCard, error) {
	var cardInfoFiltered []mongodb.DBCard
	var singleCard mongodb.DBCard
	var filter = bson.M{"name": cardName}

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
	log.Info().Timestamp().Msgf("Success: created collection:\n", collection)

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

	for _, card := range cardInfoFiltered {
		singleCard = card
	}

	return singleCard, err
}

func UpdateSingleCard(cardName string, cardQuantity int, dbCollection string) error {
	var newQuantity = cardQuantity+1
	var updateFilter = bson.M{"name": cardName}
	var updateSet = bson.D{
		{"$set", bson.D{{"quantity", newQuantity}}},
	}

	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive env vars")
		return err
	}

	client, ctx, cancelCtx, err := mongodb.CreateClient()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: Creating client\n")
		return err
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error().Timestamp().Err(err).Msg("Error: closing client\n")
		}
		cancelCtx()
	}()

	collection := client.Database(conf.DbName).Collection(dbCollection)
	log.Info().Timestamp().Msgf("Success: created collection:\n", collection)

	updateResult, err := collection.UpdateOne(ctx, updateFilter, updateSet)
	if err != nil {
		log.Error().Timestamp().Err(err).Msgf("Error: updating the quantity of a card in collection of db:\n", dbCollection, conf.DbName)
		return err
	}

	log.Info().Timestamp().Msgf("Success: Updated Documents!\n", updateResult)

	return err
}