package transport

import (
	"github.com/nguyencatpham/go-effective-study/pkg/utl/model"
)

// User model response
// swagger:response userResp
type swaggUserResponse struct {
	// in:body
	Body struct {
		*model.User
	}
}

// Users model response
// swagger:response userListResp
type swaggUserListResponse struct {
	// in:body
	Body struct {
		Users []model.User `json:"users"`
		Page  int          `json:"page"`
	}
}
