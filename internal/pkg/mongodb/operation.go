package mongodb

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"magicTGArchive/internal/pkg/importer"
)
//FIXME: If insert of duplicates increase the quantity counter by +1
func InsertCard(cardInfo importer.Cards, client *mongo.Client, ctx context.Context) error {

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

func AllCardInfo(client *mongo.Client, ctx context.Context) error{
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

func SingleCardInfo(cardName string, client *mongo.Client, ctx context.Context) error {
	defer func() {
		err := client.Disconnect(ctx)
		log.Err(err)
	}()

	collection := client.Database("Magic:The-Gathering-Archive").Collection("cards")


	filter := bson.M{"name": cardName}
	filterCursor, err := collection.Find(ctx, filter)
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

func DeleteSingleCard(cardName string, client *mongo.Client, ctx context.Context) error {
	defer func() {
		err := client.Disconnect(ctx)
		log.Err(err)
	}()

	collection := client.Database("Magic:The-Gathering-Archive").Collection("cards")

	result, err := collection.DeleteOne(ctx, bson.M{"name": cardName})
	if err != nil {
		log.Error().Err(err)
		return err
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)

	return nil
}

func UpdateSingleCard(cardName string, client *mongo.Client, ctx context.Context) error {
	collection := client.Database("Magic:The-Gathering-Archive").Collection("cards")

	//quantity := findDocumentWithCardNameAndSafeQuantityToVariable

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"name": cardName},
		bson.D{
		{"$set", bson.D{{"quantity", 2}}},
		},
	)
	if err != nil {
		log.Error().Err(err)
		return err
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)


	return nil
}