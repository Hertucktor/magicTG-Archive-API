package importer

import (
	"fmt"
	"github.com/rs/zerolog/log"
)

const URL = "https://api.magicthegathering.io/v1/cards"

func HandleRequest(){
	filter := "?name=Archangel Avacyn"

	card,err := RequestCardInfo(URL, filter)
	if err != nil {
		log.Fatal().Err(err)
	}

	for _, card := range card.Cards{
		fmt.Println(card.ImageURL)
	}
}
