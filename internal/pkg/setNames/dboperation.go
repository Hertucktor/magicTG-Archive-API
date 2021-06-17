package setNames

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"magicTGArchive/internal/pkg/config"
	"magicTGArchive/internal/pkg/mongodb"
	"time"
)

type SetNameInfo struct {
	SetName  []string `json:"setName" bson:"setName"`
	Created  string   `json:"created" bson:"created"`
	Modified string   `json:"modified" bson:"modified"`
}

func ReturnAllSetName(conf config.Config){
	log.Info().Msg("Endpoint Hit: returnAllCardsBySet")

	client, ctx, err := BuildClient(conf)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("v: couldn't build client")
	}

	setNames, err := allSetNames(client, ctx, conf)

	if err = insertSetNames(setNames, client, ctx, conf); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't insert set names into collection")
	}
}

func allSetNames(client *mongo.Client, ctx context.Context, conf config.Config)([]string, error){
	var filter = bson.M{}
	var cards []mongodb.Card

	collection := client.Database(conf.DBName).Collection(conf.DBCollectionAllcards)
	log.Info().Timestamp().Msgf("Successful: created collection:\n", collection)

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Error().Timestamp().Err(err).Msgf("Error: ")
		return nil, err
	}

	if err = cursor.All(ctx, &cards); err != nil {
		log.Error().Timestamp().Err(err).Msgf("Error: couldn't decode data into interface:\n")
		return nil, err
	}

	setNames := sortSetNames(cards)

	defer func() {
		if err = cursor.Close(ctx); err != nil {
			log.Error().Timestamp().Err(err).Msgf("Error: couldn't close cursor:%v", cursor.Current)
		}
		log.Info().Timestamp().Msg("Gracefully closed cursor")
	}()

	return setNames, err
}

func insertSetNames(setNames []string, client *mongo.Client, ctx context.Context, conf config.Config) error {

	var newSetNames = SetNameInfo{
		SetName:  setNames,
		Created:  time.Now().String(),
		Modified: "",
	}

	collection := client.Database(conf.DBName).Collection(conf.DBCollectionSetNames)
	log.Info().Timestamp().Msgf("Successful: created collection:\n", collection)

	//FIXME: Upsert logic or directly with Mongos Query language -> this is a hack
	if err := collection.Drop(ctx); err != nil {
		log.Error().Timestamp().Err(err)
		return err
	}

	_, err := collection.InsertOne(ctx, newSetNames)
	if err != nil {
		log.Error().Timestamp().Err(err).Msgf("Error: couldn't insert into collection of db: %v\n", conf.DBName)
		return err
	}

	return err
}

func sortSetNames(cards []mongodb.Card) ([]string){
	var setNames []string
	for _, card := range cards {
		if !stringInSlice(card.SetName, setNames) {
			setNames = append(setNames, card.SetName)
		}
	}
	return setNames
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
