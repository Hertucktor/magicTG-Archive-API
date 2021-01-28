package importer

import (
	"fmt"
	"github.com/rs/zerolog/log"
)

const URL = "https://api.magicthegathering.io/v1/cards"

func HandleRequest(filter string) error{

	card,err := RequestCardInfo(URL, filter)
	if err != nil {
		log.Error().Err(err)
	}

	for _, card := range card.Cards{
		fmt.Println(card.ImageURL)
	}
	return err
}
