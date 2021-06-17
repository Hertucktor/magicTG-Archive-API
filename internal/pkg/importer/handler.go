package importer

import (
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/config"
	"time"
)

func ImportCardsIntoDatabase(conf config.Config, page, delimiter int) error{
	client, ctx, err := ImporterClient(conf.DBUser, conf.DBPass, conf.DBPort, conf.DBName)
	if err != nil {
		return err
	}

	log.Info().Msgf("Start of import: %v", time.Now().Unix())
	for delimiter != 0{
		requestAllCards, err := RequestAllCards(page)
		if err != nil {
			return err
		}

		for _, card := range requestAllCards.Cards{
			//If Card is in Database, update modified else insert card
			found, err := FindCard(card.SetName, card.Number, client, ctx, conf)
			if err != nil {
				return err
			}
			if found != true{
				if err = InsertImportCard(card, client, ctx, conf); err != nil {
					return err
				}
			}else {
				if err = UpdateSingleCard(card, card.SetName, card.Number, client, ctx, conf); err != nil{
					return err
				}
			}
		}

		logPageImpression(page)

		increadePageImpression(page)

		delimiter = len(requestAllCards.Cards)
	}
	log.Info().Msgf("End of import: %v", time.Now().Unix())
	return err
}

func logPageImpression(page int){
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

func increadePageImpression(page int) (increasedPage int) {
	increasedPage = page + 1
	return
}
