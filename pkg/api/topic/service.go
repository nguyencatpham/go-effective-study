package topic

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
	"github.com/nguyencatpham/go-effective-study/pkg/api/topic/platform/pgsql"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/model"
)

// Service represents topic application interface
type Service interface {
	Create(echo.Context, model.Topic) (*model.Topic, error)
	List(echo.Context, *model.Pagination) ([]model.Topic, error)
	View(echo.Context, int) (*model.Topic, error)
	Delete(echo.Context, int) error
	Update(echo.Context, *Update) (*model.Topic, error)
}

// New creates new topic application service
func New(db *pg.DB, udb UDB, rbac RBAC, sec Securer) *Topic {
	return &Topic{db: db, udb: udb, rbac: rbac, sec: sec}
}

// Initialize initalizes Topic application service with defaults
func Initialize(db *pg.DB, rbac RBAC, sec Securer) *Topic {
	return New(db, pgsql.NewTopic(), rbac, sec)
}

// Topic represents topic application service
type Topic struct {
	db   *pg.DB
	udb  UDB
	rbac RBAC
	sec  Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents topic repository interface
type UDB interface {
	Create(orm.DB, model.Topic) (*model.Topic, error)
	View(orm.DB, int) (*model.Topic, error)
	List(orm.DB, *model.ListQuery, *model.Pagination) ([]model.Topic, error)
	Update(orm.DB, *model.Topic) error
	Delete(orm.DB, *model.Topic) error
}

// RBAC represents role-based-access-control interface
type RBAC interface {
}
