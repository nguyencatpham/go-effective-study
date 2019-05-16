// Package topic contains topic application services
package topic

import (
	"github.com/labstack/echo"
	"gitlab.com/nguyencatpham/go-effective-study/pkg/utl/query"
	"gitlab.com/nguyencatpham/go-effective-study/pkg/utl/structs"
)

// Create creates a new topic account
func (u *Topic) Create(c echo.Context, req gorsk.Topic) (*gorsk.Topic, error) {
	if err := u.rbac.AccountCreate(c, req.RoleID, req.CompanyID, req.LocationID); err != nil {
		return nil, err
	}
	req.Password = u.sec.Hash(req.Password)
	return u.udb.Create(u.db, req)
}

// List returns list of topics
func (u *Topic) List(c echo.Context, p *gorsk.Pagination) ([]gorsk.Topic, error) {
	au := u.rbac.Topic(c)
	q, err := query.List(au)
	if err != nil {
		return nil, err
	}
	return u.udb.List(u.db, q, p)
}

// View returns single topic
func (u *Topic) View(c echo.Context, id int) (*gorsk.Topic, error) {
	if err := u.rbac.EnforceTopic(c, id); err != nil {
		return nil, err
	}
	return u.udb.View(u.db, id)
}

// Delete deletes a topic
func (u *Topic) Delete(c echo.Context, id int) error {
	topic, err := u.udb.View(u.db, id)
	if err != nil {
		return err
	}
	if err := u.rbac.IsLowerRole(c, topic.Role.AccessLevel); err != nil {
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
func (u *Topic) Update(c echo.Context, req *Update) (*gorsk.Topic, error) {
	if err := u.rbac.EnforceTopic(c, req.ID); err != nil {
		return nil, err
	}

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
