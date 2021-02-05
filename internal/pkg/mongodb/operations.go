package mongodb

import (
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"magicTGArchive/internal/pkg/importer"
)

func InsertCard(cardInfo importer.Card) error {
	cardInfo.Quantity = 1

	client, ctx, err := CreateClient()
	if err != nil {
		log.Error().Err(err)
		return err
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Err(err)
		}
	}()

	collection := client.Database("Magic:The-Gathering-Archive").Collection("cards")

	insertResult, err := collection.InsertOne(ctx, cardInfo)
	if err != nil {
		log.Error().Err(err)
		return err
	}

	log.Info().Msgf("",insertResult.InsertedID)

	return err
}

func AllCardInfo() (bson.M, error){
	var cards bson.M

	client, ctx, err := CreateClient()
	if err != nil {
		log.Error().Err(err)
		return cards, err
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Err(err)
		}
	}()

	collection := client.Database("Magic:The-Gathering-Archive").Collection("cards")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Error().Err(err)
		return cards, err
	}

	defer func() {
		if err := cursor.Close(ctx); err != nil {
			log.Error().Err(err)
		}
	}()

	for cursor.Next(ctx) {
		if err = cursor.Decode(&cards); err != nil {
			log.Error().Err(err)
			return cards, err
		}
	}

	return cards, err
}

func SingleCardInfo(multiverseID string) ([]DBCards, error) {
	var cardInfoFiltered []DBCards
	filter := bson.M{"multiverseid": multiverseID}

	client, ctx, err := CreateClient()
	if err != nil {
		log.Error().Err(err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Err(err)
		}
	}()

	collection := client.Database("Magic:The-Gathering-Archive").Collection("cards")

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
		log.Error().Err(err)
		return err
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

	log.Info().Msgf("Amount of Documents removed:\n", result.DeletedCount)

	return err
}

func UpdateSingleCard(cardName string, cardQuantity int) error {
	newQuantity := cardQuantity+1

	client, ctx, err := CreateClient()
	if err != nil {
		log.Error().Err(err)
		return err
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Err(err)
		}
	}()

	collection := client.Database("Magic:The-Gathering-Archive").Collection("cards")

	result, err := collection.UpdateOne(
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

	log.Info().Msgf("Updated %v Documents!\n", result.ModifiedCount)

	return err
}