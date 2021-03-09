package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"magicTGArchive/internal/pkg/env"
	"magicTGArchive/internal/pkg/mongodb"
	"time"
)

type SetNameInfo struct {
	SetName  []string `json:"setName" bson:"setName"`
	Created  string   `json:"created" bson:"created"`
	Modified string   `json:"modified" bson:"modified"`
}

func main() {
	returnAllSetName()
}

func returnAllSetName(){
	log.Info().Msg("Endpoint Hit: returnAllCardsBySet")

	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't receive env vars")
	}

	client, ctx, err := buildClient(conf)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("v: couldn't build client")
	}

	setNames, err := allSetNames(client, ctx, conf)
	
	if err = insertSetNames(setNames, client, ctx); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't insert set names into collection")
	}
}

func allSetNames(client *mongo.Client, ctx context.Context, conf env.Conf)([]string, error){
	var filter = bson.M{}
	var cards []mongodb.Card

	collection := client.Database(conf.DbName).Collection(conf.DbCollAllCards)
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

func buildClient(conf env.Conf)(*mongo.Client, context.Context, error){

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://"+conf.DbUser+":"+conf.DbPass+"@"+conf.DbPort+"/"+conf.DbName))
	if err != nil {
		log.Error().Err(err)
		return client, nil, err
	}

	ctx := context.TODO()
	if err = client.Connect(ctx); err != nil {
		log.Error().Err(err)
		return client, ctx, err
	}

	return client, ctx, err
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

func insertSetNames(setNames []string, client *mongo.Client, ctx context.Context) error {
	conf, err := env.ReceiveEnvVars()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't receive env vars")
		return err
	}

	var newSetNames = SetNameInfo{
		SetName:  setNames,
		Created:  time.Now().String(),
		Modified: "",
	}

	collection := client.Database(conf.DbName).Collection("setNames")
	log.Info().Timestamp().Msgf("Successful: created collection:\n", collection)

	//FIXME: Upsert logic or directly with Mongos Query language -> this is a hack
	if err = collection.Drop(ctx); err != nil {
		log.Error().Timestamp().Err(err)
		return err
	}

	_, err = collection.InsertOne(ctx, newSetNames)
	if err != nil {
		log.Error().Timestamp().Err(err).Msgf("Error: couldn't insert into collection of db: %v\n", conf.DbName)
		return err
	}

	return err
}