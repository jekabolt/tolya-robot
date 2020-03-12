package schemas

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Genders = map[string]int{
	"male":   0,
	"female": 1,
}

var DBName = "users"
var ConsumersCollectionName = "consumers"
var JoinedCollectionName = "joined"

const (
	Top = iota
	Bottom
	Shoe
)

// top sizes
const (
	XXS = iota
	XS
	S
	M
	XL
	XXL
	XXXL
)

// style concepts
const (
	CasualMassMarket = iota
	StreetWear
	Classic
	AvantGarde
)

// clo types
const (
	Outwear = iota
	Tops
	Bottoms
	Footwear
	Accessories
)

type Consumer struct {
	ID                     primitive.ObjectID `json:"_id" bson:"_id"`
	ChatID                 string             `json:"chatID" bson:"chatID"`
	LAT                    int                `json:"lat" bson:"lat"`
	Gender                 int                `json:"gender" bson:"gender"`
	TopSizes               []int              `json:"topSizes" bson:"topSizes"`
	BottomSizes            []int              `json:"bottomSizes" bson:"bottomSizes"`
	ShoeSizes              []int              `json:"shoeSizes" bson:"shoeSizes"`
	StyleConcepts          []int              `json:"styleConcepts" bson:"styleConcepts"`
	FavoriteTypesOfClothes []int              `json:"favoriteTypesOfClothes" bson:"favoriteTypesOfClothes"`
}

type ConsumerMarshal struct {
	ID                     primitive.ObjectID `json:"_id" bson:"_id"`
	ChatID                 string             `json:"chatID" bson:"chatID"`
	LAT                    int                `json:"lat" bson:"lat"`
	Gender                 []int              `json:"gender" bson:"gender"`
	TopSizes               []int              `json:"topSizes" bson:"topSizes"`
	BottomSizes            []int              `json:"bottomSizes" bson:"bottomSizes"`
	ShoeSizes              []int              `json:"shoeSizes" bson:"shoeSizes"`
	StyleConcepts          []int              `json:"styleConcepts" bson:"styleConcepts"`
	FavoriteTypesOfClothes []int              `json:"favoriteTypesOfClothes" bson:"favoriteTypesOfClothes"`
}

type TGUser struct {
	User      *tgbotapi.User `json:"user" bson:"user"`
	Submitted bool           `json:"submitted" bson:"chatID"`
	ChatID    int64          `json:"chatID" bson:"chatID"`
}

type Post struct {
	PostSelect
	Title        string `json:"title" bson:"title"`
	Price        string `json:"price" bson:"price"`
	DiscountRate string `json:"discountRate" bson:"discountRate"`
	AboutText    string `json:"aboutText" bson:"aboutText"`
	Hashtags     string `json:"hashtags" bson:"hashtags"`
	Link         string `json:"link" bson:"link"`
	Image        string `json:"image" bson:"image"`
}

type PostSelect struct {
	Gender       int   `json:"gender" bson:"gender"`
	TypeOfCloth  int   `json:"typeOfCloth" bson:"typeOfCloth"`
	StyleConcept int   `json:"styleConcept" bson:"styleConcept"`
	TopSizes     []int `json:"topSizes" bson:"topSizes"`
	BottomSizes  []int `json:"bottomSizes" bson:"bottomSizes"`
	ShoeSizes    []int `json:"shoeSizes" bson:"shoeSizes"`
}
