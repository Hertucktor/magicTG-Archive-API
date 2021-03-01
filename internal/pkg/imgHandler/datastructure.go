package imgHandler

type Img struct {
	ID string `json:"id" bson:"_id"`
	ImgLink string `json:"imgLink" bson:"imglink"`
	SetName string `json:"setName" bson:"setname"`
}