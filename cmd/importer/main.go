package main

import (
	"fmt"
	"log"
	"magicTGArchive/internal/pkg/importer"
)

const URL = "https://api.magicthegathering.io/v1/cards?name=Archangel Avacyn"

func main(){
	filter := ""

	card,err := importer.RequestCardInfo(URL, filter)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(card)
}