package mongodb

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"magicTGArchive/internal/pkg/env"
	"time"
)

var dbName = "Magic:The-Gathering-Archive"

func CreateClient() (*mongo.Client, context.Context, context.CancelFunc, error) {
	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive env vars")
		return nil, nil, nil, err
	}

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://"+conf.DbUser+":"+conf.DbPass+"@"+conf.DbPort+"/"+conf.DbName))
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