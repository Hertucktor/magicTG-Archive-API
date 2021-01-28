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
func InsertCardInfo(cardInfo importer.MTGResponse, client *mongo.Client, ctx context.Context) error {

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
	var cards []bson.M
	if err = cursor.All(ctx, &cards); err != nil {
		log.Fatal().Err(err)
	}
	fmt.Println(cards)

	return err
}