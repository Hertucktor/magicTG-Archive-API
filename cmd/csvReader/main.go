package main

import (
	"encoding/csv"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
)

//var dbCollName = "setImages"

func main() {
	var filePath = "./csv/mtgSetIcons.csv"
	/*var img csvReader.Img
	//var imgBaseURL = "https://media.magic.wizards.com/images/featured/"

	if err := csvReader.InsertSetImg(img, dbCollName); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't insert ImgData into db")
	}*/

	ReadCSV(filePath)

}

func ReadCSV(fileName string){
	csvFile, err := os.Open(fileName)

	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't open CSV")
	}

	reader := csv.NewReader(csvFile)

	records, _ := reader.ReadAll()
	//first index for column, second index for line entry
	fmt.Println(records[1][0])

}