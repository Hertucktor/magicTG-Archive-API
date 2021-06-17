package api

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"magicTGArchive/internal/pkg/mongodb"
)

type Img struct {
	ID string `json:"id" bson:"_id"`
	ImgLink string `json:"imgLink" bson:"imglink"`
	SetName string `json:"setName" bson:"setname"`
}

func InsertCard(cardInfo mongodb.Card,dbName string, dbCollection string, client *mongo.Client, ctx context.Context) error {
	collection := client.Database(dbName).Collection(dbCollection)

	insertResult, err := collection.InsertOne(ctx, cardInfo)
	if err != nil {
		return err
	}

	log.Info().Msgf("Success: insertion result:\n", insertResult)

	return err
}

func AllCards(dbName string, dbCollection string, client *mongo.Client, ctx context.Context) ([]mongodb.Card, error){
	var filter = bson.M{}
	var cards []mongodb.Card

	collection := client.Database(dbName).Collection(dbCollection)

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return cards, err
	}

	if err = cursor.All(ctx, &cards); err != nil {
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

func AllCardsBySet(setName string, dbName string, dbCollection string, client *mongo.Client, ctx context.Context)([]mongodb.Card, error){
	var filter = bson.M{"setName": setName}
	var cards []mongodb.Card

	collection := client.Database(dbName).Collection(dbCollection)

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return cards, err
	}

	if err = cursor.All(ctx, &cards); err != nil {
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

func ReadSingleCardEntry(setName string, number string,dbName string, dbCollection string, client *mongo.Client, ctx context.Context) (mongodb.Card, error) {
	var readFilter = bson.M{"setName": setName, "number": number}
	var card mongodb.Card

	collection := client.Database(dbName).Collection(dbCollection)

	if err := collection.FindOne(ctx, readFilter).Decode(&card); err != nil {
		return mongodb.Card{}, err
	}

	return card, nil
}

func DeleteSingleCard(setName string, number string,dbName string, dbCollection string, client *mongo.Client, ctx context.Context) (*mongo.DeleteResult, error) {
	var deleteFilter = bson.M{"setName": setName, "number": number}

	collection := client.Database(dbName).Collection(dbCollection)

	deleteResult, err := collection.DeleteOne(ctx, deleteFilter)
	if err != nil {
		return deleteResult, err
	}

	log.Info().Timestamp().Msgf("Success: Result after successful deletion:%v", deleteResult)

	return deleteResult, err
}

func UpdateSingelCard(setName string, number string, cardQuantity int, dbName string, dbCollection string, client *mongo.Client, ctx context.Context) error {
	var newQuantity = cardQuantity+1
	var updateFilter = bson.M{"setName": setName, "number":number}
	var updateSet = bson.D{
		{"$set", bson.D{{"quantity", newQuantity}}},
	}

	collection := client.Database(dbName).Collection(dbCollection)

	updateResult, err := collection.UpdateOne(ctx, updateFilter, updateSet)
	if err != nil {
		return err
	}

	log.Info().Timestamp().Msgf("Success: Updated Documents!\n", updateResult)

	return err
}

//CRUD OPERATIONS FOR IMAGE
func SingleSetImg(setName string,dbName string, dbCollection string, client *mongo.Client, ctx context.Context) (Img, error) {
	var readFilter = bson.M{"setName": setName}
	var imgInfo Img

	collection := client.Database(dbName).Collection(dbCollection)

	if err := collection.FindOne(ctx, readFilter).Decode(&imgInfo); err != nil {
		return Img{}, err
	}

	return imgInfo, nil
}