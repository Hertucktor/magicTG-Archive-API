package mongodb

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func CreateClient() (*mongo.Client, error) {
	//TODO: secure username and password
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://admin:admin@127.0.0.1:27017/Magic:The-Gathering-Archive"))
	if err != nil {
		log.Error().Err(err)
		return client, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Error().Err(err)
		return client, err
	}

	defer func() {
		err := client.Disconnect(ctx)
		log.Err(err)
	}()

	return client, err
}