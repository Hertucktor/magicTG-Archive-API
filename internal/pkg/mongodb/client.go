package mongodb

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var dbUser = "admin"
var dbPass = "admin"
var dbPort = "127.0.0.1:27017"
var dbName = "Magic:The-Gathering-Archive"

func CreateClient() (*mongo.Client, context.Context, context.CancelFunc, error) {
	//TODO: secure username and password
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://"+dbUser+":"+dbPass+"@"+dbPort+"/"+dbName))
	if err != nil {
		log.Error().Err(err)
		return client,nil, nil, err
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Error().Err(err)
		return client, ctx, cancelFunc, err
	}

	return client, ctx, cancelFunc, err
}