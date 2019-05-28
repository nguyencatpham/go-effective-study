package model

// PostStatus represents type of topic
type PostStatus int

const (

	// Darft is a new word need to learn
	Darft PostStatus = 100
	// Pending is a new word need to learn
	Pending PostStatus = 110
	// Publish is a new word need to learn
	Publish PostStatus = 120
	// Inherit is a new word need to learn
	Inherit PostStatus = 130
)

// PostType represents type of topic
type PostType int

const (

	// Page is a new word need to learn
	Page PostType = 100
	// Revision is a new word need to learn
	Revision PostType = 110
	// Blog is a new word need to learn
	Blog PostType = 120
	// Attachment is a new word need to learn
	Attachment PostType = 130
)

// Post represents topic for english lesson model
type Post struct {
	Base
	Author         int        `json:"user_id,omitempty"`
	ParentID       int        `json:"parent_id,omitempty"`
	Name           string     `json:"name,omitempty"`
	Title          string     `json:"title,omitempty"`
	Content        string     `json:"content,omitempty"`
	Status         PostStatus `sql:"default:100" json:"status,omitempty"`
	AllowComment   bool       `json:"allow_comment,omitempty"`
	Password       string     `json:"password,omitempty"`
	GUID           string     `json:"guid,omitempty"`
	MenuOrder      int8       `json:"menu_order,omitempty"`
	Type           PostType   `sql:"default:100" json:"type,omitempty"`
	MineType       string     `json:"mine_type,omitempty"`
	CommentCount   int16      `json:"comment_count,omitempty"`
	Slug           string     `json:"slug,omitempty"`
	LanguageID     int8       `json:"language_id"`
	Language       Language   `json:"language"`
	CategoryID     int        `json:"category_id"`
	Category       Category   `json:"category"`
	Tags           string     `json:"tags"`
	LocalizePostID string     `json:"localize_post_id"`

	PostMetas []PostMeta `pg:"fk:post_id" json:"post_metas,omitempty"`
}
