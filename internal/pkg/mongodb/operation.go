package mongodb

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"magicTGArchive/internal/pkg/importer"
)
//FIXME: Don't allow inserts of duplicates
func InsertCardInfo(cardInfo importer.APIResponse, client *mongo.Client, ctx context.Context) error {

	defer func() {
		err := client.Disconnect(ctx)
		log.Err(err)
	}()

	collection := client.Database("Magic:The-Gathering-Archive").Collection("cards")

	insertResult, err := collection.InsertOne(context.TODO(), cardInfo)
	if err != nil {
		log.Error().Err(err)
		return err
	}
	fmt.Println("Inserted card with ID:", insertResult.InsertedID)

	return err
}

func GetAllCardInfo(client *mongo.Client, ctx context.Context) error{
	defer func() {
		err := client.Disconnect(ctx)
		log.Err(err)
	}()

	collection := client.Database("Magic:The-Gathering-Archive").Collection("cards")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal().Err(err)
	}
	defer func() {
		err := cursor.Close(ctx)
		if err != nil {
			log.Error().Err(err)
		}
	}()
	for cursor.Next(ctx) {
		var episode bson.M
		if err = cursor.Decode(&episode); err != nil {
			log.Fatal().Err(err)
		}
		fmt.Println(episode)
	}

	return err
}
//FIXME: Receive One specific Document by filter
func GetFilteredSingleCardInfo(client *mongo.Client, ctx context.Context) error {
	defer func() {
		err := client.Disconnect(ctx)
		log.Err(err)
	}()

	collection := client.Database("Magic:The-Gathering-Archive").Collection("cards")

	filterCursor, err := collection.Find(ctx, bson.M{"name":"Quicksand"})
	if err != nil {
		log.Error().Err(err)
		return err
	}
	var cardInfoFiltered []bson.M
	if err = filterCursor.All(ctx, &cardInfoFiltered); err != nil {
		log.Error().Err(err)
		return err
	}

	fmt.Println(cardInfoFiltered)

	return nil
}