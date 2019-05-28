package model

// Media represents topic for english lesson model
type Media struct {
	ID             int    `json:"id"`
	Name           string `json:"name,omitempty"`
	MimeType       string `json:"mimeType,omitempty"`
	SeoFilename    string `json:"seo_file_name,omitempty"`
	AltAttribute   string `json:"alt,omitempty"`
	TitleAttribute bool   `json:"title,omitempty"`
	URL            bool   `json:"url,omitempty"`
}
