package model

// Language represents topic for english lesson model
type Language struct {
	ID            int    `json:"id"`
	Name          string `json:"name,omitempty"`
	Culture       string `json:"culture,omitempty"`
	UniqueSeoCode string `json:"unique_seo_code,omitempty"`
	FlagImageURL  string `json:"flag_image_url,omitempty"`
	Rtl           bool   `json:"rl,omitempty"`
	IsEnable      bool   `json:"is_enable,omitempty"`
	Posts         []Post `pg:"fk:language_id"  json:"post,omitempty"`
}
