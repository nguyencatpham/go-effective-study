package transport

import (
	"github.com/nguyencatpham/go-effective-study/pkg/utl/model"
)

// Product model response
// swagger:response productResp
type swaggProductResponse struct {
	// in:body
	Body struct {
		*model.Product
	}
}

// Products model response
// swagger:response productListResp
type swaggProductListResponse struct {
	// in:body
	Body struct {
		Products []model.Product `json:"products"`
		Page     int             `json:"page"`
	}
}

// Product create request
// swagger:model productCreate
type createReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// UpdateReq update request
// swagger:model productUpdate
type updateReq struct {
	ID          int    `json:"-"`
	Name        string `json:"name" validate:"min=2"`
	Description string `json:"description,omitempty""`
	Type        int    `json:"type,omitempty"`
}

// Product list request
// swagger:model list
type listResponse struct {
	Data       []model.Product `json:"data"`
	Page       int             `json:"page"`
	TotalItems int             `json:"totalItems"`
}
