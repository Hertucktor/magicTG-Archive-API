package database

//FIXME: For user input validation and quantity increasment CRUCIAL!
/*func InsertDataset(cardInfo importer.APIResponseForMultipleCards) error {

	card, err := mongodb.SingleCardInfo(cardInfo.Card.Name)
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error reading from the database")
		return err
	}

	if card.ID != "" {
		if err := mongodb.UpdateSingleCard(cardInfo.Cards, card.Quantity); err != nil {
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
}*/
