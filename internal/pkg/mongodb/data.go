package mongodb

type Card struct {
	ID			  string 		 `json:"id,omitempty" bson:"_id,omitempty"`
	Name          string         `json:"name,omitempty" bson:"name,omitempty"`
	ManaCost      string         `json:"manaCost,omitempty" bson:"manaCost,omitempty"`
	Cmc           float64        `json:"cmc,omitempty" bson:"cmc,omitempty"`
	Colors        []string       `json:"colors,omitempty" bson:"colors,omitempty"`
	ColorIdentity []string       `json:"colorIdentity,omitempty" bson:"colorIdentity,omitempty"`
	Type          string         `json:"type,omitempty" bson:"type,omitempty"`
	Supertypes    []string  	 `json:"supertypes,omitempty" bson:"supertypes,omitempty"`
	Types         []string       `json:"types,omitempty" bson:"types,omitempty"`
	Subtypes      []string       `json:"subtypes,omitempty" bson:"subtypes,omitempty"`
	Rarity        string         `json:"rarity,omitempty" bson:"rarity,omitempty"`
	Set           string         `json:"set,omitempty" bson:"set,omitempty"`
	SetName       string         `json:"setname,omitempty" bson:"setname,omitempty"`
	Text          string         `json:"text,omitempty" bson:"text,omitempty"`
	Flavor        string         `json:"flavor,omitempty" bson:"flavor,omitempty"`
	Artist        string         `json:"artist,omitempty" bson:"artist,omitempty"`
	Number        string         `json:"number,omitempty" bson:"number,omitempty"`
	Power         string         `json:"power,omitempty" bson:"power,omitempty"`
	Toughness     string         `json:"toughness,omitempty" bson:"toughness,omitempty"`
	Layout        string         `json:"layout,omitempty" bson:"layout,omitempty"`
	Multiverseid  int            `json:"multiverseid,omitempty" bson:"multiverseid,omitempty"`
	ImageURL      string         `json:"imageUrl,omitempty" bson:"imageUrl,omitempty"`
	Rulings       []Rulings  	 `json:"rulings,omitempty" bson:"rulings,omitempty"`
	ForeignNames  []ForeignNames `json:"foreignNames,omitempty" bson:"foreignNames,omitempty"`
	Printings     []string       `json:"printings,omitempty" bson:"printings,omitempty"`
	OriginalText  string         `json:"originalText,omitempty" bson:"originalText,omitempty"`
	OriginalType  string         `json:"originalType,omitempty" bson:"originalType,omitempty"`
	Legalities    []Legalities   `json:"legalities,omitempty" bson:"legalities,omitempty"`
	CardID        string         `json:"cardID,omitempty" bson:"cardID,omitempty"`
	Quantity	  int 			 `json:"quantity" bson:"quantity"`
}
type Rulings struct {
	Date 		string 			 `json:"date,omitempty" bson:"date"`
	Text 		string 			 `json:"text,omitempty" bson:"text"`
}
type ForeignNames struct {
	Name         string 		 `json:"name,omitempty" bson:"name"`
	Text         string 		 `json:"text,omitempty" bson:"text"`
	Flavor       string 		 `json:"flavor,omitempty" bson:"flavor"`
	ImageURL     string 		 `json:"imageUrl,omitempty" bson:"imageUrl"`
	Language     string 		 `json:"language,omitempty" bson:"language"`
	Multiverseid int    		 `json:"multiverseid,omitempty" bson:"multiverseid"`
}
type Legalities struct {
	Format   	string 			 `json:"format,omitempty" bson:"format"`
	Legality 	string 			 `json:"legality,omitempty" bson:"legality"`
}