package importer

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"reflect"
)

/*
RequestCardInfo receives a response with type *http.Response from
the official mtg api containing all available card detail for one card
specified with the multiverseID
Returning the response and an error
*/
func RequestCardInfo() (int, string, error) {
	//var customError Error
	var card Card
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
	err := UnmarshalJSON(body,&card)
	if err != nil {
		log.Err(err)
		return
	}


	//err = json.Unmarshal(body, &card)
	//if err != nil {
	//	log.Err(err)
	//	fmt.Println(card)
	//	return card, resp.StatusCode, err
	//}

	return resp.StatusCode,"", err
}

func UnmarshalJSON(data []byte, v interface{}) error{
	err := json.Unmarshal(data, &v)
	log.Err(err)
	return err
}
