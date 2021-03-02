package main

type MultipleCards struct {
	Cards []Card `json:"cards"`
}
type Card struct {
	Name          string         `json:"name,omitempty"`
	ManaCost      string         `json:"manaCost,omitempty"`
	Cmc           float64        `json:"cmc,omitempty"`
	Colors        []string       `json:"colors,omitempty"`
	ColorIdentity []string       `json:"colorIdentity,omitempty"`
	Type          string         `json:"type,omitempty"`
	Supertypes    []string  	 `json:"superTypes,omitempty"`
	Types         []string       `json:"types,omitempty"`
	Subtypes      []string       `json:"subTypes,omitempty"`
	Rarity        string         `json:"rarity,omitempty"`
	Set           string         `json:"set,omitempty"`
	SetName       string         `json:"setName,omitempty"`
	Text          string         `json:"text,omitempty"`
	Flavor        string         `json:"flavor,omitempty"`
	Artist        string         `json:"artist,omitempty"`
	Number        string         `json:"number,omitempty"`
	Power         string         `json:"power,omitempty"`
	Toughness     string         `json:"toughness,omitempty"`
	Layout        string         `json:"layout,omitempty"`
	Multiverseid  int            `json:"multiverseID,omitempty"`
	ImageURL      string         `json:"imageURL,omitempty"`
	Rulings       []Rulings      `json:"rulings,omitempty"`
	ForeignNames  []ForeignNames `json:"foreignNames,omitempty"`
	Printings     []string       `json:"printings,omitempty"`
	OriginalText  string         `json:"originalText,omitempty"`
	OriginalType  string         `json:"originalType,omitempty"`
	Legalities    []Legalities   `json:"legalities,omitempty"`
	ID            string         `json:"id,omitempty"`
	Variations    []string       `json:"variations,omitempty"`
	Quantity	  int            `json:"quantity"`
}
type Rulings struct {
	Date string `json:"date"`
	Text string `json:"text"`
}
type ForeignNames struct {
	Name         string `json:"name"`
	Text         string `json:"text"`
	Flavor       string `json:"flavor"`
	ImageURL     string `json:"imageURL"`
	Language     string `json:"language"`
	Multiverseid int    `json:"multiverseID"`
}
type Legalities struct {
	Format   string `json:"format"`
	Legality string `json:"legality"`
}