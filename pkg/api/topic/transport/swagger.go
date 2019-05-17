package transport

import (
	"gitlab.com/nguyencatpham/go-effective-study/pkg/utl/model"
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
