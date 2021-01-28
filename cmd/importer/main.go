package main

import "magicTGArchive/internal/pkg/importer"

func main() {
	filter := "?name=Archangel Avacyn"
	_ = importer.HandleRequest(filter)
}