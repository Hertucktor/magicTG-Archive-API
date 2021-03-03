package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/vitali-fedulov/images"
)

func main() {
	uploadedImgPath := "./temp/images/upload/ModernHorizons.png"
	compareImgPath := "./temp/images/compare/Core20.png"


	if err := selectImgToCompare(uploadedImgPath, compareImgPath); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't compare images")
	}
}

func selectImgToCompare(uploadedImg string, compareImg string) error{
	// Open photos.
	imgA, err := images.Open(uploadedImg)
	if err != nil {
		log.Error().Timestamp().Err(err)
	}
	imgB, err := images.Open(compareImg)
	if err != nil {
		log.Error().Timestamp().Err(err)
	}

	// Calculate hashes and image sizes.
	hashA, imgSizeA := images.Hash(imgA)
	hashB, imgSizeB := images.Hash(imgB)

	// Image comparison.
	if images.Similar(hashA, hashB, imgSizeA, imgSizeB) {
		fmt.Println("Images are similar.")
	} else {
		fmt.Println("Images are distinct.")
	}

	return nil
}
