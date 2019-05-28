package model

// Category represents topic for english lesson model
type Category struct {
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

	Posts []Post `pg:"fk:category_id" json:"posts,omitempty"`
}
