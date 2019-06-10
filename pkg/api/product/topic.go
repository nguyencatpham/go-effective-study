// Package topic contains topic application services
package topic

import (
	"log"

	"github.com/labstack/echo"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/model"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/structs"
)

// Create creates a new topic account
func (u *Topic) Create(c echo.Context, req model.Topic) (*model.Topic, error) {
	log.Println("aaaaaaaaaa")
	return u.udb.Create(u.db, req)
}

// List returns list of topics
func (u *Topic) List(c echo.Context, p *model.Pagination, query []string) ([]model.Topic, int, error) {
	// TODO: []string to filter query
	q := &model.FilterQuery{}
	log.Println("Params", query)
	return u.udb.List(u.db, q, p)
}

// View returns single topic
func (u *Topic) View(c echo.Context, query []string) (*model.Topic, error) {
	q := &model.FilterQuery{
		// Query:  "topic.id=?",
		// Params: []interface{}{query},
	}
	return u.udb.View(u.db, q)
}

// Delete deletes a topic
func (u *Topic) Delete(c echo.Context, id int) error {
	q := &model.FilterQuery{Query: "topic.id=?", Params: []interface{}{id}}
	topic, err := u.udb.View(u.db, q)
	if err != nil {
		return err
	}
	return u.udb.Delete(u.db, topic)
}

// Update updates topic's contact information
func (u *Topic) Update(c echo.Context, req model.UpdateReq) (*model.Topic, error) {
	q := &model.FilterQuery{Query: "topic.id=?", Params: []interface{}{req.ID}}
	topic, err := u.udb.View(u.db, q)
	if err != nil {
		return nil, err
	}

	structs.Merge(topic, req)
	if err := u.udb.Update(u.db, topic); err != nil {
		return nil, err
	}

	return topic, nil
}
