package imageComparator

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/vitali-fedulov/images"
)

func SelectImgToCompare(uploadedImg string, compareImg string) error{
	imgA, err := images.Open(uploadedImg)
	if err != nil {
		log.Error().Timestamp().Err(err)
	}
	imgB, err := images.Open(compareImg)
	if err != nil {
		log.Error().Timestamp().Err(err)
	}

	hashA, imgSizeA := images.Hash(imgA)
	hashB, imgSizeB := images.Hash(imgB)

	if images.Similar(hashA, hashB, imgSizeA, imgSizeB) {
		fmt.Println("Images are similar.")
	} else {
		fmt.Println("Images are distinct.")
	}

	return nil
}
