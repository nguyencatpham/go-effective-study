// Package topic contains topic application services
package topic

import (
	"github.com/labstack/echo"
	"gitlab.com/nguyencatpham/go-effective-study/pkg/utl/model"
	"gitlab.com/nguyencatpham/go-effective-study/pkg/utl/structs"
)

// Create creates a new topic account
func (u *Topic) Create(c echo.Context, req model.Topic) (*model.Topic, error) {
	return u.udb.Create(u.db, req)
}

// List returns list of topics
func (u *Topic) List(c echo.Context, p *model.Pagination) ([]model.Topic, error) {
	q := &model.ListQuery{}
	return u.udb.List(u.db, q, p)
}

// View returns single topic
func (u *Topic) View(c echo.Context, id int) (*model.Topic, error) {
	return u.udb.View(u.db, id)
}

// Delete deletes a topic
func (u *Topic) Delete(c echo.Context, id int) error {
	topic, err := u.udb.View(u.db, id)
	if err != nil {
		return err
	}
	return u.udb.Delete(u.db, topic)
}

// Update contains topic's information used for updating
type Update struct {
	ID        int
	FirstName *string
	LastName  *string
	Mobile    *string
	Phone     *string
	Address   *string
}

// Update updates topic's contact information
func (u *Topic) Update(c echo.Context, req *Update) (*model.Topic, error) {

	topic, err := u.udb.View(u.db, req.ID)
	if err != nil {
		return nil, err
	}

	structs.Merge(topic, req)
	if err := u.udb.Update(u.db, topic); err != nil {
		return nil, err
	}

	return topic, nil
}
