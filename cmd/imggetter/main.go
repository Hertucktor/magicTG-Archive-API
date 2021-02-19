package main

import "fmt"

func main() {
	imgAttr := Img{
		PicName:   "i0vHvg5vVd",
		Extension: "png",
	}

	var imgBaseURL = "https://media.magic.wizards.com/images/featured/"

	var requestURL = imgBaseURL + imgAttr.PicName + "." + imgAttr.Extension

	fmt.Println(requestURL)



}
