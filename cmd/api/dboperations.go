package main

import (
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"magicTGArchive/internal/pkg/env"
	"magicTGArchive/internal/pkg/mongodb"
)

func InsertCard(cardInfo mongodb.DBCard, dbCollection string) error {
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

	insertResult, err := collection.InsertOne(ctx, cardInfo)
	if err != nil {
		log.Error().Timestamp().Err(err).Msgf("Error: couldn't insert into collection of db:\n", dbCollection, conf.DbName)
		return err
	}

	log.Info().Msgf("Success: insertion result:\n", insertResult)

	return err
}

func AllCardInfo(dbCollection string) ([]bson.M, error){
	var filter = bson.M{}
	var card bson.M
	var cards []bson.M

	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive env vars")
		return nil, err
	}

	client, ctx, cancelCtx, err := mongodb.CreateClient()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: Creating Client\n")
		return cards, err
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
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
		if err = cursor.Close(ctx); err != nil {
			log.Error().Timestamp().Err(err).Msgf("Error: couldn't close cursor:\n", cursor)
		}
		log.Info().Msg("Closed cursor:")
	}()

	for cursor.Next(ctx) {

		if err = cursor.Decode(&card); err != nil {
			log.Error().Timestamp().Err(err).Msgf("Error: couldn't decode data into interface:\n")
			return cards, err
		}
		cards = append(cards, card)
	}

	return cards, err
}

func SingleCardInfo(setName string, number string, dbCollection string) ([]bson.M, error) {
	var databaseResponse []bson.M
	var readFilter = bson.M{"setname": setName, "number": number}

	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive env vars")
		return databaseResponse, err
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

	cursor, err := collection.Find(ctx, readFilter)
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: cursor couldn't be created\n")
		return databaseResponse, err
	}

	defer func() {
		if err = cursor.Close(ctx); err != nil {
			log.Error().Timestamp().Err(err).Msgf("Error: couldn't close cursor", cursor)
		}
		log.Info().Msg("Success: Closed cursor\n")
	}()

	if err = cursor.All(ctx, &databaseResponse); err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: problem with the cursor\n")
		return databaseResponse, err
	}

	return databaseResponse, err
}

func DeleteSingleCard(setName string, number string, dbCollection string) (*mongo.DeleteResult, error) {
	var deleteResult *mongo.DeleteResult
	var deleteFilter = bson.M{"setname": setName, "number": number}
	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive env vars")
		return deleteResult, err
	}

	client, ctx, cancelCtx, err := mongodb.CreateClient()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: Creating Client\n")
		return deleteResult, err
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Error().Timestamp().Err(err).Msg("Error: closing client\n")
		}
		cancelCtx()
	}()

	collection := client.Database(conf.DbName).Collection(dbCollection)
	log.Info().Timestamp().Msgf("Success: created collection:\n", collection)


	deleteResult, err = collection.DeleteOne(ctx, deleteFilter)
	if err != nil {
		log.Error().Timestamp().Err(err).Msgf("Error: couldn't delete document with given deleteFilter\n")
		return deleteResult, err
	}
	log.Info().Timestamp().Msgf("Success: Result after successful deletion:\n", deleteResult)

	return deleteResult, err
}

func UpdateSingleCard(setName string, number string, cardQuantity int, dbCollection string) error {
	var newQuantity = cardQuantity+1
	var updateFilter = bson.M{"setname": setName, "number":number}
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
		if err = client.Disconnect(ctx); err != nil {
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