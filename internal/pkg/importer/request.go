package importer

import (
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)

/*
RequestMultipleInfosOfOneCard receives a response with type *http.Response from
the official mtg api containing all available card detail for one card
specified with the multiverseID
Returning the response and an error
*/
func RequestMultipleInfosOfOneCard(URL string) (APIResponseForMultipleCards, error) {
	var response APIResponseForMultipleCards

	resp, err := http.Get(URL)
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

	if response.Cards != nil{

	}

	return response, err
}

func RequestOneCard(URL string) (APIResponseForOneCard, error) {
	var response APIResponseForOneCard

	resp, err := http.Get(URL)
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

func RequestAllCards(page string) (MTGDevAPIResponse, error) {
	var response MTGDevAPIResponse

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