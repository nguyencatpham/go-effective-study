package transport

import (
	"github.com/nguyencatpham/go-effective-study/pkg/utl/model"
)

// Topic model response
// swagger:response topicResp
type swaggTopicResponse struct {
	// in:body
	Body struct {
		*model.Topic
	}
}

// Topics model response
// swagger:response topicListResp
type swaggTopicListResponse struct {
	// in:body
	Body struct {
		Topics []model.Topic `json:"topics"`
		Page   int           `json:"page"`
	}
}

// Topic create request
// swagger:model topicCreate
type createReq struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Type        model.TopicType `json:"type"`
}

// UpdateReq update request
// swagger:model topicUpdate
type updateReq struct {
	ID          int    `json:"-"`
	Name        string `json:"name" validate:"min=2"`
	Description string `json:"description,omitempty""`
	Type        int    `json:"type,omitempty"`
}

// Topic list request
// swagger:model list
type listResponse struct {
	Data       []model.Topic `json:"data"`
	Page       int           `json:"page"`
	TotalItems int           `json:"totalItems"`
}
