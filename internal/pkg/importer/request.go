package importer

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)



/*
RequestCardInfo receives a response with type *http.Response from
the official mtg api containing all available card detail for one card
specified with the multiverseID
Returning the response and an error
*/
func RequestCardInfo(URL string, filter string) error {
	//var customError Error
	var card Card

	resp, err := http.Get(URL+filter)
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
	err = UnmarshalJSON(body,&card)
	if err != nil {
		log.Err(err)
		return err
	}


	//err = json.Unmarshal(body, &card)
	//if err != nil {
	//	log.Err(err)
	//	fmt.Println(card)
	//	return card, resp.StatusCode, err
	//}

	return err
}

func UnmarshalJSON(data []byte, v interface{}) error{
	err := json.Unmarshal(data, &v)
	log.Err(err)
	return err
}
