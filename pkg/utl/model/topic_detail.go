package model

// TopicDetail represents topic for english lesson model
type TopicDetail struct {
	Base
	Name        string `json:"name"`
	Description string `json:"description"`
	TopicID     int    `json:"topic_id,omitempty"`
	Topic       Topic  `json:"topic,omitempty"`
}
