package mockdb

import (
	"github.com/go-pg/pg/orm"
	"gitlab.com/nguyencatpham/go-effective-study/pkg/utl/model"
)

// User database mock
type User struct {
	CreateFn         func(orm.DB, model.User) (*model.User, error)
	ViewFn           func(orm.DB, int) (*model.User, error)
	FindByUsernameFn func(orm.DB, string) (*model.User, error)
	FindByTokenFn    func(orm.DB, string) (*model.User, error)
	ListFn           func(orm.DB, *model.ListQuery, *model.Pagination) ([]model.User, error)
	DeleteFn         func(orm.DB, *model.User) error
	UpdateFn         func(orm.DB, *model.User) error
}

// Create mock
func (u *User) Create(db orm.DB, usr model.User) (*model.User, error) {
	return u.CreateFn(db, usr)
}

// View mock
func (u *User) View(db orm.DB, id int) (*model.User, error) {
	return u.ViewFn(db, id)
}

// FindByUsername mock
func (u *User) FindByUsername(db orm.DB, uname string) (*model.User, error) {
	return u.FindByUsernameFn(db, uname)
}

// FindByToken mock
func (u *User) FindByToken(db orm.DB, token string) (*model.User, error) {
	return u.FindByTokenFn(db, token)
}

// List mock
func (u *User) List(db orm.DB, lq *model.ListQuery, p *model.Pagination) ([]model.User, error) {
	return u.ListFn(db, lq, p)
}

// Delete mock
func (u *User) Delete(db orm.DB, usr *model.User) error {
	return u.DeleteFn(db, usr)
}

// Update mock
func (u *User) Update(db orm.DB, usr *model.User) error {
	return u.UpdateFn(db, usr)
}
