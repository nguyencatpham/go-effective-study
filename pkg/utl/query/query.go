package query

import (
	"github.com/labstack/echo"
	"gitlab.com/nguyencatpham/go-effective-study/pkg/utl/model"
)

// List prepares data for list queries
func List(u *model.AuthUser) (*model.ListQuery, error) {
	switch true {
	case u.Role <= model.AdminRole: // user is SuperAdmin or Admin
		return nil, nil
	case u.Role == model.CompanyAdminRole:
		return &model.ListQuery{Query: "company_id = ?", ID: u.CompanyID}, nil
	case u.Role == model.LocationAdminRole:
		return &model.ListQuery{Query: "location_id = ?", ID: u.LocationID}, nil
	default:
		return nil, echo.ErrForbidden
	}
}
