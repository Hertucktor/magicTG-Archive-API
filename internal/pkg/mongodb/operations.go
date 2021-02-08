package mongodb

import (
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"magicTGArchive/internal/pkg/env"
	"magicTGArchive/internal/pkg/importer"
)

var dbCollection = "myCardCollection"

func InsertCard(cardInfo importer.Card) error {
	cardInfo.Quantity = 1

	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive env vars")
		return err
	}

	client, ctx, cancelCtx, err := CreateClient()
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

func AllCardInfo() (bson.M, error){
	var filter = bson.M{}
	var cards bson.M

	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive env vars")
		return nil, err
	}

	client, ctx, cancelCtx, err := CreateClient()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: Creating Client\n")
		return cards, err
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error().Timestamp().Err(err).Msg("Error: closing client\n")
		}
		cancelCtx()
	}()

	collection := client.Database(conf.DbName).Collection(dbCollection)
	log.Info().Timestamp().Msgf("Successful: created collection:\n", collection)

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Error().Timestamp().Err(err).Msgf("Error: ")
		return cards, err
	}

	defer func() {
		if err := cursor.Close(ctx); err != nil {
			log.Error().Timestamp().Err(err).Msgf("Error: couldn't close cursor:\n", cursor)
		}
		log.Info().Msgf("Closed cursor:", cursor)
	}()

	for cursor.Next(ctx) {
		if err = cursor.Decode(&cards); err != nil {
			log.Error().Timestamp().Err(err).Msgf("Error: couldn't decode data into interface:\n", cards)
			return cards, err
		}
	}

	return cards, err
}

func SingleCardInfo(cardName string) (DBCard, error) {
	var cardInfoFiltered []DBCard
	var singleCard DBCard
	var filter = bson.M{"name": cardName}

	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive env vars")
		return singleCard, err
	}

	client, ctx, cancelCtx, err := CreateClient()
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

func DeleteSingleCard(cardName string) error {
	var filter = bson.M{"name": cardName}

	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive env vars")
		return err
	}

	client, ctx, cancelCtx, err := CreateClient()
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
	log.Info().Timestamp().Msgf("Success: created collection:\n", collection)

	deleteResult, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Error().Timestamp().Err(err).Msgf("Error: couldn't delete document with given filter\n",filter)
		return err
	}

	log.Info().Timestamp().Msgf("Success: Result after successful deletion:\n", deleteResult)

	return err
}

func UpdateSingleCard(cardName string, cardQuantity int) error {
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

	client, ctx, cancelCtx, err := CreateClient()
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