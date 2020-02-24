package schemas

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

var Genders = map[string]int{
	"male":   0,
	"female": 1,
}

var DBName = "users"
var ConsumersCollectionName = "consumers"
var JoinedCollectionName = "joined"

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
	ChatID                 string `json:"chatID" bson:"chatID"`
	LAT                    int    `json:"lat" bson:"gender"`
	Gender                 []int  `json:"gender" bson:"gender"`
	TopSizes               []int  `json:"topSizes" bson:"topSizes"`
	BottomSizes            []int  `json:"bottomSizes" bson:"bottomSizes"`
	ShoeSizes              []int  `json:"shoeSizes" bson:"shoeSizes"`
	StyleConcepts          []int  `json:"styleConcepts" bson:"styleConcepts"`
	FavoriteTypesOfClothes []int  `json:"favoriteTypesOfClothes" bson:"favoriteTypesOfClothes"`
}

type TGUser struct {
	User      *tgbotapi.User `json:"user" bson:"user"`
	Submitted bool           `json:"submitted" bson:"chatID"`
	ChatID    int64          `json:"chatID" bson:"chatID"`
}
