package model

// Catalog represents topic for english lesson model
type Catalog struct {
	Base
	Name       string   `json:"name,omitempty"`
	Title      string   `json:"title,omitempty"`
	Content    string   `json:"content,omitempty"`
	MenuOrder  int8     `json:"menu_order,omitempty"`
	Slug       string   `json:"slug,omitempty"`
	LanguageID int8     `json:"language_id"`
	Language   Language `json:"language"`
	Tags       string   `json:"tags"`
	LocalizeID string   `json:"localize_id"`

	Products []Product `pg:"fk:catalog_id" json:"products,omitempty"`
}
