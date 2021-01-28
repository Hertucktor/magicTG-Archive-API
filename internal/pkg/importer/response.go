package importer

type Card struct {
	Name string `json:"name"`
	ManaCost string `json:"manaCost"`
	Cmc float32 `json:"cmc"`
	Colors []string `json:"colors"`
	ColorIdentity []string `json:"colorIdentity"`
	Type string `json:"type"`
	SuperTypes []string `json:"supertypes"`
	Types []string `json:"types"`
	SubTypes []string `json:"subtypes"`
	Rarity string `json:"rarity"`
	Set string `json:"set"`
	SetName  string            `json:"setName"`
	Text string                `json:"text"`
	Artist string              `json:"artist"`
	Number string              `json:"number"`
	Power string               `json:"power"`
	Toughness string           `json:"toughness"`
	Layout string              `json:"layout"`
	MultiverseId int           `json:"multiverseid"`
	ImageURL string            `json:"imageUrl"`
	Watermark string           `json:"watermark"`
	Rulings []Ruling           `json:"rulings"`
	ForeignNames []ForeignName `json:"foreignNames"`
	Printings []string         `json:"printings"`
	OriginalText string        `json:"originalText"`
	OriginalType string        `json:"originalType"`
	Legalities []Legal         `json:"legalities"`
	ID string                  `json:"id"`
}

type Ruling struct {
	Date string `json:"date"`
	Text string `json:"text"`
}

type ForeignName struct {
	Name string `json:"name"`
	Text string `json:"text"`
	Flavor string `json:"flavor"`
	ImageURL string `json:"imageUrl"`
	Language string `json:"language"`
	MultiverseID string `json:"multiverseid"`
}

type Legal struct {
	Format string `json:"format"`
	Legality string `json:"legality"`
}

