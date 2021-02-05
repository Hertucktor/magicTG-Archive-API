package mongodb

type DBCard struct {
	ID			  string 		`bson:"_id,omitempty"`
	Name          string         `bson:"name,omitempty"`
	ManaCost      string         `bson:"manaCost,omitempty"`
	Cmc           float64        `bson:"cmc,omitempty"`
	Colors        []string       `bson:"colors,omitempty"`
	ColorIdentity []string       `bson:"colorIdentity,omitempty"`
	Type          string         `bson:"type,omitempty"`
	Supertypes    []string  	 `bson:"supertypes,omitempty"`
	Types         []string       `bson:"types,omitempty"`
	Subtypes      []string       `bson:"subtypes,omitempty"`
	Rarity        string         `bson:"rarity,omitempty"`
	Set           string         `bson:"set,omitempty"`
	SetName       string         `bson:"setName,omitempty"`
	Text          string         `bson:"text,omitempty"`
	Flavor        string         `bson:"flavor,omitempty"`
	Artist        string         `bson:"artist,omitempty"`
	Number        string         `bson:"number,omitempty"`
	Power         string         `bson:"power,omitempty"`
	Toughness     string         `bson:"toughness,omitempty"`
	Layout        string         `bson:"layout,omitempty"`
	Multiverseid  int            `bson:"multiverseid,omitempty"`
	ImageURL      string         `bson:"imageUrl,omitempty"`
	Rulings       []Rulings  	 `bson:"rulings,omitempty"`
	ForeignNames  []ForeignNames `bson:"foreignNames,omitempty"`
	Printings     []string       `bson:"printings,omitempty"`
	OriginalText  string         `bson:"originalText,omitempty"`
	OriginalType  string         `bson:"originalType,omitempty"`
	Legalities    []Legalities   `bson:"legalities,omitempty"`
	CardID            string     `bson:"cardID,omitempty"`
	Quantity	  int 			 `bson:"quantity"`
}
type Rulings struct {
	Date string `bson:"date"`
	Text string `bson:"text"`
}
type ForeignNames struct {
	Name         string `bson:"name"`
	Text         string `bson:"text"`
	Flavor       string `bson:"flavor"`
	ImageURL     string `bson:"imageUrl"`
	Language     string `bson:"language"`
	Multiverseid int    `bson:"multiverseid"`
}
type Legalities struct {
	Format   string `bson:"format"`
	Legality string `bson:"legality"`
}