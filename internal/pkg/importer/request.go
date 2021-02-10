package importer

import (
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)

/*
RequestAllCards receives a response with type *http.Response from
the mtg api containing 100 cards
Returning the response and an error
*/
func RequestAllCards(page string) (APIResponseForMultipleCards, error) {
	var response APIResponseForMultipleCards

	resp, err := http.Get("https://api.magicthegathering.io/v1/cards?page="+page)
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: problem with http GET request\n")
		return response, err
	}

	log.Info().Timestamp().Msgf("HTTP GET REQUEST TO https://api.magicthegathering.io/v1/cards?page=\n",page)

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Error().Timestamp().Err(err).Msg("Error: could't close response body\n")
		}
	}()

	if resp.StatusCode != 200{
		err := errors.New("Http statuscode != 200")
		log.Error().Timestamp().Err(err).Msgf("Error: Http status code:\n", resp.StatusCode)
		return response, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't read from response body\n")
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't unmarshal body into MTGDevAPIResponse struct\n")
		return response, err
	}

	return response, err
}