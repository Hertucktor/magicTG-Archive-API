package main

import (
	"fmt"
	"github.com/rs/zerolog/log"

	"net/http"
)

func main(){
	response, err := ImportCardInfo()
	if err != nil {
		log.Err(err)
	}
	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)

}
/*
ImportCardInfo receives a response with type *http.Response from
the official mtg api containing all available card detail for one card
specified with the multiverseID
 */
func ImportCardInfo() (*http.Response, error) {
	URL := "https://api.magicthegathering.io/v1/cards/"
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

	return resp, err
}
