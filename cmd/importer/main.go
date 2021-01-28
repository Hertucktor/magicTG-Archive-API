package main

import (
	"log"
	"magicTGArchive/internal/pkg/importer"
)

const URL = "https://api.magicthegathering.io/v1/cards"

func main(){
	filter := "?name=avacyn"

	err := importer.RequestCardInfo(URL, filter)
	if err != nil {
		log.Fatalln(err)
	}
}