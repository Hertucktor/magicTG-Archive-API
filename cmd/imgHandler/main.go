package main

import (
	"encoding/csv"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"os"
)

//var dbCollName = "setImages"

func main() {
	/*var img imgHandler.Img
	//var imgBaseURL = "https://media.magic.wizards.com/images/featured/"

	if err := imgHandler.InsertSetImg(img, dbCollName); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't insert ImgData into db")
	}*/


	ReadCSV()

}

func ReadCSV(){
	fmt.Println("s")
	csvFile, err := os.Open("./csv/mtgSetIcons.csv")
	fmt.Println(csvFile)
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't open CSV")
	}

	r := csv.NewReader(csvFile)
	fmt.Println(r)
	end := false
	for end == true  {
		record, err := r.Read()
		if err == io.EOF{
			end = true
		}
		if err != nil {
			log.Fatal().Timestamp().Err(err)
		}

		fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
	}

}