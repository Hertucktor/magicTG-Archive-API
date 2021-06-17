package main

import (
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/imageComparator"
)

func main() {
	uploadedImgPath := "./temp/images/upload/ModernHorizons.png"
	compareImgPath := "./temp/images/compare/Core20.png"


	if err := imageComparator.SelectImgToCompare(uploadedImgPath, compareImgPath); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't compare images")
	}
}
