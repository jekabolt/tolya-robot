package schemas

var Genders = map[string]int{
	"male":   0,
	"female": 1,
}

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
	Premium
	Classic
	AvantGarde
)

// clo types
const (
	Outwear = iota
	Tops
	Bottoms
	Outerwear
	Footwear
	Accessories
)

type User struct {
	Gender                 int
	TopSizes               []int
	BottomSizes            []int
	ShoeSizes              []int
	StyleConcepts          []int
	FavoriteTypesOfClothes []int
	Seen                   bool
}
