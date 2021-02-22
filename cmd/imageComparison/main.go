package main

import (
	"fmt"
	"github.com/vitali-fedulov/images"
)

func main() {
	// Open photos.
	imgA, err := images.Open("./temp/images/compare/ModernHorizons.png")
	if err != nil {
		panic(err)
	}
	imgB, err := images.Open("./temp/images/compare/Core20.png")
	if err != nil {
		panic(err)
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
}
