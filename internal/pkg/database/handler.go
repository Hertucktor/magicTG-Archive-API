package database

import (
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/importer"
	"magicTGArchive/internal/pkg/mongodb"
)

func InsertDataset(cardInfo importer.Card, multiverseID string) error {

	cards, err := mongodb.SingleCardInfo(multiverseID)
	if err != nil {
		log.Error().Err(err)
		return err
	}
	if cards != nil {
		if err := mongodb.InsertCard(cardInfo); err != nil {
			log.Error().Err(err)
			return err
		}
	}

	return err
}
