package model

// PostMeta represents topic for english lesson model
type PostMeta struct {
	Base
	PostID int   `json:"post_id,omitempty"`
	Post   Post  `json:"post,omitempty"`
	Key    int8  `json:"key,omitempty"`
	Value  int16 `json:"value,omitempty"`
}
