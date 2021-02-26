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
