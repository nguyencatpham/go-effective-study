package model

// TopicType represents type of topic
type TopicType int

const (
	// Vocabulary is a new word need to learn
	Vocabulary TopicType = 100

	// Reading is a new topic need to read
	Reading TopicType = 110
)

// Topic represents topic for english lesson model
type Topic struct {
	Base
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Type         TopicType     `json:"type"`
	TopicDetails []TopicDetail `pg:"fk:topic_id" json:"topic_details,omitempty"`
}

// UpdateReq contains topic's information used for updating
type UpdateReq struct {
	ID          int    `json:"-"`
	Name        string `json:"name" validate:"min=2"`
	Description string `json:"description,omitempty"`
	Type        int    `json:"type,omitempty"`
}
