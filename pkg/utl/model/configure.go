package model

// Configure represents topic for english lesson model
type Configure struct {
	ID    int    `json:"id"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
