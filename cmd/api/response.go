package main

type Card struct {
	Name          string         `json:"name,omitempty"`
	ManaCost      string         `json:"manaCost,omitempty"`
	Cmc           float64        `json:"cmc,omitempty"`
	Colors        []string       `json:"colors,omitempty"`
	ColorIdentity []string       `json:"colorIdentity,omitempty"`
	Type          string         `json:"type,omitempty"`
	Supertypes    []string  	 `json:"supertypes,omitempty"`
	Types         []string       `json:"types,omitempty"`
	Subtypes      []string       `json:"subtypes,omitempty"`
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
	Multiverseid  int            `json:"multiverseid,omitempty"`
	ImageURL      string         `json:"imageUrl,omitempty"`
	Rulings       []Rulings  	 `json:"rulings,omitempty"`
	ForeignNames  []ForeignNames `json:"foreignNames,omitempty"`
	Printings     []string       `json:"printings,omitempty"`
	OriginalText  string         `json:"originalText,omitempty"`
	OriginalType  string         `json:"originalType,omitempty"`
	Legalities    []Legalities   `json:"legalities,omitempty"`
	ID            string         `json:"id,omitempty"`
	Variations    []string       `json:"variations,omitempty"`
	Quantity	  int 			 `json:"quantity"`
}
type Rulings struct {
	Date string `json:"date,omitempty"`
	Text string `json:"text,omitempty"`
}
type ForeignNames struct {
	Name         string `json:"name,omitempty"`
	Text         string `json:"text,omitempty"`
	Flavor       string `json:"flavor,omitempty"`
	ImageURL     string `json:"imageUrl,omitempty"`
	Language     string `json:"language,omitempty"`
	Multiverseid int    `json:"multiverseid,omitempty"`
}
type Legalities struct {
	Format   string `json:"format,omitempty"`
	Legality string `json:"legality,omitempty"`
}