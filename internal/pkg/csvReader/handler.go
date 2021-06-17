package csvReader

import (
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/config"
)

func TransferCSVDataToDatabase(filePath string, conf config.Config) {
	imgInfos, err := ConvertCSVEntriesIntoStruct(filePath)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("")
	}

	for _, imgInfo := range imgInfos.Imgs{
		if err = InsertImgInfo(imgInfo,conf.DBName, conf.DBCollectionSetimages); err != nil {
			log.Fatal().Timestamp().Err(err).Msg("")
		}
	}
}
