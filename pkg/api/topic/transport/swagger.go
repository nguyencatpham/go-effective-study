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

// Topic update request
// swagger:model topicUpdate
type updateReq struct {
	ID        int     `json:"-"`
	FirstName *string `json:"first_name,omitempty" validate:"omitempty,min=2"`
	LastName  *string `json:"last_name,omitempty" validate:"omitempty,min=2"`
	Mobile    *string `json:"mobile,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	Address   *string `json:"address,omitempty"`
}

// Topic list request
// swagger:model list
type listResponse struct {
	Data       []model.Topic `json:"data"`
	Page       int           `json:"page"`
	TotalItems int           `json:"totalItems"`
}
