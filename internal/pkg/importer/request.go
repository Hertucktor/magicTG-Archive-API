package importer

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)

const URL = "https://api.magicthegathering.io/v1/cards"

/*
RequestCardInfo receives a response with type *http.Response from
the official mtg api containing all available card detail for one card
specified with the multiverseID
Returning the response and an error
*/
func RequestCardInfo(cardName string, filterLang string) (MTGResponse, error) {
	var response MTGResponse
	nameFilter := "?name=" + cardName

	resp, err := http.Get(URL+ nameFilter +filterLang)
	if err != nil {
		log.Print(err)
		return response, err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Print(err)
		}
	}()

	if resp.StatusCode != 200{
		log.Error().Msg(resp.Status)
		return response, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Err(err)
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Err(err)
		return response, err
	}

	return response, err
}