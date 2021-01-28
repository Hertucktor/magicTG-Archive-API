package importer

func FilterSelector (lang string) string{
	var filter string
	if lang == "de" {
		filter = "&language=german"
	} else {
		return ""
	}

	return filter
}

/*func HandleRequest(filter string) error{

	card,err := RequestCardInfo(URL, filter)
	if err != nil {
		log.Error().Err(err)
	}

	for _, card := range card.Cards{
		fmt.Println(card.ImageURL)
	}
	return err
}*/
