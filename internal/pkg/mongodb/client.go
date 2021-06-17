package mongodb

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"magicTGArchive/internal/pkg/config"
	"time"
)

func CreateClient() (*mongo.Client, context.Context, context.CancelFunc, error) {
	conf, err := config.GetConfig("config.yml")
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive env vars")
		return nil, nil, nil, err
	}

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://"+conf.DBUser+":"+conf.DBPass+"@"+conf.DBPort+"/"+conf.DBName))
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