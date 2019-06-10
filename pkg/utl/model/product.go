package model

// ProductStatus represents type of topic
type ProductStatus string

const (

	// Active is a new word need to learn
	Active ProductStatus = "ACTIVE"
)

// Product represents topic for english lesson model
type Product struct {
	Base
	AlternateName int           `json:"alternateName,omitempty"`
	Name          string        `json:"name,omitempty"`
	Description   string        `json:"description,omitempty"`
	ImageURL      string        `json:"imageUrl,omitempty"`
	Note          bool          `json:"note,omitempty"`
	NumOfArticles string        `json:"numOfArticles,omitempty"`
	NumOfImages   string        `json:"numOfImages,omitempty"`
	Price         int8          `json:"price,omitempty"`
	Quantity      int           `sql:"default:0" json:"quantity,omitempty"`
	RegularPrice  string        `json:"regularPrice,omitempty"`
	SaleOff       int16         `json:"saleOff,omitempty"`
	Sku           string        `json:"sku,omitempty"`
	Size          int8          `json:"size"`
	Status        ProductStatus `json:"status"`
	ThumbnailSrc  int           `json:"thumbnailSrc"`
	Unit          int           `json:"unit"`
	Tags          string        `json:"tags"`
	Slug          string        `json:"slug"`
	CatalogID     int           `json:"catalog_id"`
	Catalog       Catalog       `json:"catalog"`
}
