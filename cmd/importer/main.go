package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"magicTGArchive/internal/pkg/env"
)


func main() {
	var page = 1
	var delimiter = 1

	client, ctx, err := ImporterClient()
	if err != nil {
		log.Fatal().Timestamp().Err(err)
	}

	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't receive env vars")
	}

	for delimiter != 0{
		requestAllCards, err := RequestAllCards(page)
		if err != nil {
			log.Fatal().Timestamp().Err(err).Msg("Fatal: API Request threw error")
		}

		for _, card := range requestAllCards.Cards{
			//If Card is in Database, update modified else insert card
			found, err := FindCard(card.SetName, card.Number, client, ctx, conf)
			if err != nil {
				log.Fatal().Timestamp().Err(err).Msgf("Fatal: problem reading card from database: %v",card)
			}
			if found != true{
				if err = InsertImportCard(card, client, ctx, conf); err != nil {
					log.Fatal().Timestamp().Err(err).Msgf("Fatal: couldn't insert dataset: %v",card)
				}
			}else {
				if err = UpdateSingleCard(card, card.SetName, card.Number, client, ctx, conf); err != nil{
					log.Fatal().Timestamp().Err(err).Msgf("Fatal: couldn't update dataset:\n",card)
				}
			}
		}
		page++
		log.Info().Timestamp().Msgf("Request page number:%v one page = 100 cards", page)
		delimiter = len(requestAllCards.Cards)
	}
}

func ImporterClient() (*mongo.Client, context.Context, error) {
	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive env vars")
		return nil, nil, err
	}

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://"+conf.DbUser+":"+conf.DbPass+"@"+conf.DbPort+"/"+conf.DbName))
	if err != nil {
		log.Error().Err(err)
		return client, nil, err
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Error().Err(err)
		return client, ctx, err
	}

	return client, ctx, err
}