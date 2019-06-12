package product

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
	"github.com/nguyencatpham/go-effective-study/pkg/api/product/platform/pgsql"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/model"
)

// Service represents product application interface
type Service interface {
	Create(echo.Context, model.Product) (*model.Product, error)
	List(echo.Context, *model.Pagination, []string) ([]model.Product, int, error)
	View(echo.Context, []string) (*model.Product, error)
	Delete(echo.Context, int) error
	Update(echo.Context, model.UpdateReq) (*model.Product, error)
}

// New creates new product application service
func New(db *pg.DB, udb UDB, rbac RBAC, sec Securer) *Product {
	return &Product{db: db, udb: udb, rbac: rbac, sec: sec}
}

// Initialize initalizes Product application service with defaults
func Initialize(db *pg.DB, rbac RBAC, sec Securer) *Product {
	return New(db, pgsql.NewProduct(), rbac, sec)
}

// Product represents product application service
type Product struct {
	db   *pg.DB
	udb  UDB
	rbac RBAC
	sec  Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents product repository interface
type UDB interface {
	Create(orm.DB, model.Product) (*model.Product, error)
	View(orm.DB, *model.FilterQuery) (*model.Product, error)
	List(orm.DB, *model.FilterQuery, *model.Pagination) ([]model.Product, int, error)
	Update(orm.DB, *model.Product) error
	Delete(orm.DB, *model.Product) error
}

// RBAC represents role-based-access-control interface
type RBAC interface {
}
