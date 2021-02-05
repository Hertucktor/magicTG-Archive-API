package database

import (
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/importer"
	"magicTGArchive/internal/pkg/mongodb"
)

func InsertDataset(cardInfo importer.APIResponseForOneCard) error {

	card, err := mongodb.SingleCardInfo(cardInfo.Card.Name)
	if err != nil {
		log.Error().Err(err)
		return err
	}

	if card.ID != "" {
		if err := mongodb.UpdateSingleCard(cardInfo.Card.Name, card.Quantity); err != nil {
			log.Error().Err(err)
			return err
		}
	} else {
		if err := mongodb.InsertCard(cardInfo.Card); err != nil {
			log.Error().Err(err)
			return err
		}
	}

	return err
}
