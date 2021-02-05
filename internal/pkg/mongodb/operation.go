package mongodb

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"magicTGArchive/internal/pkg/importer"
)
//FIXME: If insert of duplicates increase the quantity counter by +1
func InsertCard(cardInfo importer.Card) error {
	/*var cardName string*/


	client, ctx, err := CreateClient()
	if err != nil {
		log.Fatal().Err(err)
	}

	defer func() {
		err := client.Disconnect(ctx)
		log.Err(err)
	}()

	collection := client.Database("Magic:The-Gathering-Archive").Collection("cards")
	cardInfo.Quantity = 1
	insertResult, err := collection.InsertOne(ctx, cardInfo)
	if err != nil {
		log.Error().Err(err)
		return err
	}
	log.Info().Msgf("",insertResult.InsertedID)

	/*results, err := SingleCardInfo(cardInfo.Name)
	if err != nil {
		log.Error().Err(err)
		return err
	}
	//When there is no document with the given name, insert new document
	if results == nil {
		//log.Info().Msgf("no document found with name: ", cardInfo.Name)
		collection := client.Database("Magic:The-Gathering-Archive").Collection("cards")
		insertResult, err := collection.InsertOne(context.TODO(), cardInfo)
		if err != nil {
			log.Error().Err(err)
			return err
		}
		log.Info().Msgf("Inserted card with ID:", insertResult.InsertedID)
	//When there is a document with the given name, update new document with card quantity +1
	}else {
		for _, card := range results{
			cardName = card.Name
			cardQuantity = card.Quantity
		}
		if err := UpdateSingleCard(cardName, cardQuantity); err != nil{
			log.Error().Err(err)
			return err
		}
	}*/

	return err
}

func AllCardInfo() error{
	client, ctx, err := CreateClient()
	if err != nil {
		log.Fatal().Err(err)
	}
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

func SingleCardInfo(cardName string) ([]DBCards, error) {
	var cardInfoFiltered []DBCards

	client, ctx, err := CreateClient()
	if err != nil {
		log.Fatal().Err(err)
	}

	defer func() {
		err := client.Disconnect(ctx)
		log.Err(err)
	}()

	collection := client.Database("Magic:The-Gathering-Archive").Collection("cards")


	filter := bson.M{"name": cardName}
	filterCursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Error().Err(err)
		return cardInfoFiltered, err
	}

	if err = filterCursor.All(ctx, &cardInfoFiltered); err != nil {
		log.Error().Err(err)
		return cardInfoFiltered, err
	}

	return cardInfoFiltered, err
}

func DeleteSingleCard(cardName string) error {
	client, ctx, err := CreateClient()
	if err != nil {
		log.Fatal().Err(err)
	}

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

func UpdateSingleCard(cardName string, cardQuantity int) error {
	newQuantity := cardQuantity+1

	client, ctx, err := CreateClient()
	if err != nil {
		log.Fatal().Err(err)
	}

	defer func() {
		err := client.Disconnect(ctx)
		log.Err(err)
	}()

	collection := client.Database("Magic:The-Gathering-Archive").Collection("cards")

	_, err = collection.UpdateOne(
		ctx,
		bson.M{"name": cardName},
		bson.D{
			{"$set", bson.D{{"quantity", newQuantity}}},
		},
	)
	if err != nil {
		log.Error().Err(err)
		return err
	}
	/*
		log.Info().Msgf("Updated %v Documents!\n", result.ModifiedCount)*/

	return nil
}