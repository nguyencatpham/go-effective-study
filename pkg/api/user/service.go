package user

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
	"gitlab.com/nguyencatpham/go-effective-study/pkg/api/user/platform/pgsql"
	"gitlab.com/nguyencatpham/go-effective-study/pkg/utl/model"
)

// Service represents user application interface
type Service interface {
	Create(echo.Context, model.User) (*model.User, error)
	List(echo.Context, *model.Pagination) ([]model.User, error)
	View(echo.Context, int) (*model.User, error)
	Delete(echo.Context, int) error
	Update(echo.Context, *Update) (*model.User, error)
}

// New creates new user application service
func New(db *pg.DB, udb UDB, rbac RBAC, sec Securer) *User {
	return &User{db: db, udb: udb, rbac: rbac, sec: sec}
}

// Initialize initalizes User application service with defaults
func Initialize(db *pg.DB, rbac RBAC, sec Securer) *User {
	return New(db, pgsql.NewUser(), rbac, sec)
}

// User represents user application service
type User struct {
	db   *pg.DB
	udb  UDB
	rbac RBAC
	sec  Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents user repository interface
type UDB interface {
	Create(orm.DB, model.User) (*model.User, error)
	View(orm.DB, int) (*model.User, error)
	List(orm.DB, *model.ListQuery, *model.Pagination) ([]model.User, error)
	Update(orm.DB, *model.User) error
	Delete(orm.DB, *model.User) error
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	User(echo.Context) *model.AuthUser
	EnforceUser(echo.Context, int) error
	AccountCreate(echo.Context, model.AccessRole, int, int) error
	IsLowerRole(echo.Context, model.AccessRole) error
}
