package main

import (
	"encoding/csv"
	"github.com/rs/zerolog/log"
	"os"
)

func ReadCSV(fileName string) ([][]string, error){
	csvFile, err := os.Open(fileName)
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't open CSV")
	}

	reader := csv.NewReader(csvFile)

	records, err := reader.ReadAll()
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't read from")
	}

	return records,err
}

func ConvertCSVEntriesIntoStruct(filePath string) (ImgCollection, error){
	var img Img
	var imgColl ImgCollection

	entries, err := ReadCSV(filePath)
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't read csv file")
		return ImgCollection{}, err
	}

	//first index for row, second index for column in csv file
	for row:=1;row<len(entries);row++{

		img.SetName = entries[row][0]
		img.ImgLink = entries[row][1]
		imgColl.Imgs = append(imgColl.Imgs, img)

	}
	return imgColl, err
}
