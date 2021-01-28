package mongodb

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"magicTGArchive/internal/pkg/importer"
)

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