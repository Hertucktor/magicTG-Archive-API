package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"magicTGArchive/internal/pkg/config"
	"time"
)


func main() {
	var page = 1
	var delimiter = 1

	configFile := "config.yml"
	conf, err := config.GetConfig(configFile)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("")
	}


	client, ctx, err := ImporterClient(conf.DBUser, conf.DBPass, conf.DBPort, conf.DBName)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("")
	}

	log.Info().Msgf("Start of import: %v", time.Now().Unix())
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

		pageImpression(page)
		//increments page counter for paginated api request by 1
		page++
		//when requestAllCards.Cards == empty/0 the import is completed
		delimiter = len(requestAllCards.Cards)
	}
	log.Info().Msgf("End of import: %v", time.Now().Unix())
}

/*
pageImpression logs the current iteration of the page counter
used in func RequestAllCards
 */
func pageImpression(page int){
	switch page {
	case 100:
		log.Info().Msgf("Reached page: %v, at time: %v",page, time.Now().Unix())
	case 200:
		log.Info().Msgf("Reached page: %v, at time: %v",page, time.Now().Unix())
	case 300:
		log.Info().Msgf("Reached page: %v, at time: %v",page, time.Now().Unix())
	case 400:
		log.Info().Msgf("Reached page: %v, at time: %v",page, time.Now().Unix())
	case 500:
		log.Info().Msgf("Reached page: %v, at time: %v",page, time.Now().Unix())
	default:
		log.Info().Timestamp().Msgf("Request page number:%v one page = 100 cards", page)
	}
}

func ImporterClient(dbUser, dbPass, dbPort, dbName string) (*mongo.Client, context.Context, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://"+dbUser+":"+dbPass+"@"+dbPort+"/"+dbName))
	if err != nil {
		return client, nil, err
	}

	ctx := context.Background()
	if err = client.Connect(ctx);err != nil {
		return client, ctx, err
	}

	return client, ctx, err
}