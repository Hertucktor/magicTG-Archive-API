package importer

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"magicTGArchive/internal/pkg/config"
	"magicTGArchive/internal/pkg/mongodb"
	"time"
)

func InsertImportCard(cardInfo mongodb.Card, client *mongo.Client, ctx context.Context, conf config.Config) error {
	cardInfo.Quantity = 1
	cardInfo.Created = time.Now().String()

	collection := client.Database(conf.DBName).Collection(conf.DBCollectionAllcards)
	log.Info().Timestamp().Msgf("Successful: connected to collection:%v", collection.Name())

	insertResult, err := collection.InsertOne(ctx, cardInfo)
	if err != nil {
		log.Error().Timestamp().Err(err).Msgf("Error: couldn't insert into collection of db:\n", conf.DBCollectionAllcards, conf.DBName)
		return err
	}
	log.Info().Msgf("Success: insertion result: %v", insertResult.InsertedID)

	return err
}

func FindCard (setName string, number string, client *mongo.Client, ctx context.Context, conf config.Config) (bool, error) {
	var readFilter = bson.M{"setName": setName, "number": number}
	var card mongodb.Card

	collection := client.Database(conf.DBName).Collection("allCards")

	_ = collection.FindOne(ctx, readFilter).Decode(&card)

	if card.Number != "" {
		return true, nil
	}

	return false, nil
}
//FIXME: update complete dataset plus modified
func UpdateSingleCard(card mongodb.Card, setName string, number string, client *mongo.Client, ctx context.Context, conf config.Config) error {
	opts := options.Update().SetUpsert(true)
	filter := bson.M{"setName": setName, "number":number}
	modifiedDate := time.Now().String()
	update := bson.D{{"$set", bson.D{{"modified", modifiedDate}}}}

	collection := client.Database(conf.DBName).Collection(conf.DBCollectionAllcards)
	log.Info().Timestamp().Msgf("Success: created collection:\n", collection)

	updateResult, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Error().Timestamp().Err(err).Msgf("Error: updating the quantity of a card in collection of db: \n", conf.DBCollectionAllcards, conf.DBName)
		return err
	}

	log.Info().Timestamp().Msgf("Success: Updated Documents!\n", updateResult)

	return err
}