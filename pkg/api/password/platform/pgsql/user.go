package pgsql

import (
	"github.com/go-pg/pg/orm"
	"gitlab.com/nguyencatpham/go-effective-study/pkg/utl/model"
)

// NewUser returns a new user database instance
func NewUser() *User {
	return &User{}
}

// User represents the client for user table
type User struct{}

// View returns single user by ID
func (u *User) View(db orm.DB, id int) (*model.User, error) {
	user := &model.User{Base: model.Base{ID: id}}
	err := db.Select(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update updates user's info
func (u *User) Update(db orm.DB, user *model.User) error {
	return db.Update(user)
}
