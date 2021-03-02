package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"magicTGArchive/internal/pkg/env"
	"magicTGArchive/internal/pkg/mongodb"
)

type Img struct {
	ID string `json:"id" bson:"_id"`
	ImgLink string `json:"imgLink" bson:"imglink"`
	SetName string `json:"setName" bson:"setname"`
}

//CRUD OPERATIONS FOR CARD
func InsertCard(cardInfo mongodb.Card, dbCollection string, client *mongo.Client, ctx context.Context) error {
	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive env vars")
		return err
	}

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

func AllCards(dbCollection string, client *mongo.Client, ctx context.Context) ([]mongodb.Card, error){
	var filter = bson.M{}
	var cards []mongodb.Card

	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive env vars")
		return nil, err
	}

	collection := client.Database(conf.DbName).Collection(dbCollection)
	log.Info().Timestamp().Msgf("Successful: created collection:\n", collection)

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Error().Timestamp().Err(err).Msgf("Error: ")
		return cards, err
	}

	if err = cursor.All(ctx, &cards); err != nil {
		log.Error().Timestamp().Err(err).Msgf("Error: couldn't decode data into interface:\n")
		return cards, err
	}

	defer func() {
		if err = cursor.Close(ctx); err != nil {
			log.Error().Timestamp().Err(err).Msgf("Error: couldn't close cursor:%v", cursor.Current)
		}
		log.Info().Timestamp().Msg("Gracefully closed cursor")
	}()

	return cards, err
}

func AllCardsBySet(setName string, dbCollection string, client *mongo.Client, ctx context.Context)([]mongodb.Card, error){
	var filter = bson.M{"setname": setName}
	var cards []mongodb.Card

	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive env vars")
		return nil, err
	}

	collection := client.Database(conf.DbName).Collection(dbCollection)
	log.Info().Timestamp().Msgf("Successful: created collection:\n", collection)

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Error().Timestamp().Err(err).Msgf("Error: ")
		return cards, err
	}

	if err = cursor.All(ctx, &cards); err != nil {
		log.Error().Timestamp().Err(err).Msgf("Error: couldn't decode data into interface:\n")
		return cards, err
	}

	defer func() {
		if err = cursor.Close(ctx); err != nil {
			log.Error().Timestamp().Err(err).Msgf("Error: couldn't close cursor:%v", cursor.Current)
		}
		log.Info().Msg("Gracefully closed cursor")
	}()

	return cards, err
}

func SingleCardInfo(setName string, number string, dbCollection string, client *mongo.Client, ctx context.Context) (mongodb.Card, error) {
	var readFilter = bson.M{"setname": setName, "number": number}
	var card mongodb.Card

	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive env vars")
		return card, err
	}

	collection := client.Database(conf.DbName).Collection(dbCollection)
	log.Info().Timestamp().Msgf("Success: created collection: %v", collection.Name())

	if err = collection.FindOne(ctx, readFilter).Decode(&card); err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't find single")
	}

	return card, err
}

func DeleteSingleCard(setName string, number string, dbCollection string, client *mongo.Client, ctx context.Context) (*mongo.DeleteResult, error) {
	var deleteResult *mongo.DeleteResult
	var deleteFilter = bson.M{"setname": setName, "number": number}
	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive env vars")
		return deleteResult, err
	}

	collection := client.Database(conf.DbName).Collection(dbCollection)
	log.Info().Timestamp().Msgf("Success: created collection:%v", collection)


	deleteResult, err = collection.DeleteOne(ctx, deleteFilter)
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't delete document with given deleteFilter")
		return deleteResult, err
	}
	log.Info().Timestamp().Msgf("Success: Result after successful deletion:%v", deleteResult)

	return deleteResult, err
}

func UpdateSingleCard(setName string, number string, cardQuantity int, dbCollection string, client *mongo.Client, ctx context.Context) error {
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

//CRUD OPERATIONS FOR IMAGE
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