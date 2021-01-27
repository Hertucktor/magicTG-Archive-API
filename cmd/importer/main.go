package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"magicTGArchive/internal/pkg/importer"
	"net/http"
	"reflect"
)

func main(){
	_, _, _ = ImportCardInfo()
	//if err != nil {
	//	log.Fatal().Err(err)
	//}
	//fmt.Println(err)
	//fmt.Println(code)

}
/*
ImportCardInfo receives a response with type *http.Response from
the official mtg api containing all available card detail for one card
specified with the multiverseID
Returning the response and an error
 */
func ImportCardInfo() (int, string, error) {
	var customError importer.Error
	//var card importer.Card
	URL := "https://api.magicthegathering.io/v/cards/"
	multiverseID := "386616"


	resp, err := http.Get(URL+multiverseID)
	if err != nil {
		log.Print(err)
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Print(err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Err(err)
	}

	_ = json.Unmarshal(body, &customError)
	fmt.Println(reflect.TypeOf(customError.Error))

	//err = json.Unmarshal(body, &card)
	//if err != nil {
	//	log.Err(err)
	//	fmt.Println(card)
	//	return card, resp.StatusCode, err
	//}

	return resp.StatusCode,"", err
}
