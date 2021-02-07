package database

import (
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/importer"
	"magicTGArchive/internal/pkg/mongodb"
)

func InsertDataset(cardInfo importer.APIResponseForOneCard) error {

	card, err := mongodb.SingleCardInfo(cardInfo.Card.Name)
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error reading from the database")
		return err
	}

	if card.ID != "" {
		if err := mongodb.UpdateSingleCard(cardInfo.Card.Name, card.Quantity); err != nil {
			log.Error().Timestamp().Err(err).Msg("Error updating the database")
			return err
		}
	} else {
		if err := mongodb.InsertCard(cardInfo.Card); err != nil {
			log.Error().Timestamp().Err(err).Msg("Error inserting in the database")
			return err
		}
	}

	return err
}
