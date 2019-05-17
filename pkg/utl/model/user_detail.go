package model

import "time"

// UserDetail detail represents topic for english lesson model
type UserDetail struct {
	Base
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	UserID        int         `json:"user_id,omitempty"`
	User          User        `json:"user,omitempty"`
	TopicDetailID int         `json:"topic_detail_id,omitempty"`
	TopicDetail   TopicDetail `json:"topic_detail,omitempty"`
	Point         int         `sql:", default:0" json:"point,omitempty"`
	StudyCount    int         `sql:", default:0" json:"study_count,omitempty"`
	RecentStudy   time.Time   `json:"recent_study,omitempty"`
}
