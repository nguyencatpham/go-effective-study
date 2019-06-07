package model

// Tag represents topic for english lesson model
type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}
