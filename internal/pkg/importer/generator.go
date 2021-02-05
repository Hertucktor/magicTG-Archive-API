package importer

import (
	"fmt"
	"strings"
)

func URLByCardNameAndLanguage(lang string, cardName string) string{
	var foreignLang string
	var URL string
	cardName = strings.ReplaceAll(cardName, " ", "+")

	if lang == "de" {
		foreignLang = "&language=german"
	} else {
		foreignLang = ""
	}

	tempURL := fmt.Sprint("https://api.magicthegathering.io/v1/cards?name="+ cardName + foreignLang)
	URL = strings.ReplaceAll(tempURL, "\n", "")

	return URL
}

func URLForMultiverserID(multiverseID string) (URL string) {
	URL = fmt.Sprint("https://api.magicthegathering.io/v1/cards/"+ multiverseID)

	return
}