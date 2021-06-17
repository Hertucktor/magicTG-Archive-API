package setNames

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"magicTGArchive/internal/pkg/config"
)

func BuildClient(conf config.Config)(*mongo.Client, context.Context, error){

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://"+conf.DBUser+":"+conf.DBPass+"@"+conf.DBPort+"/"+conf.DBName))
	if err != nil {
		log.Error().Err(err)
		return client, nil, err
	}

	ctx := context.Background()
	if err = client.Connect(ctx); err != nil {
		log.Error().Err(err)
		return client, ctx, err
	}

	return client, ctx, err
}
