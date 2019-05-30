package pgsql

import (
	"net/http"
	"strings"

	"github.com/go-pg/pg"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/helper"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/model"

	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
)

// NewTopic returns a new topic database instance
func NewTopic() *Topic {
	return &Topic{}
}

// Topic represents the client for topic table
type Topic struct{}

// Custom errors
var (
	ErrAlreadyExists = echo.NewHTTPError(http.StatusInternalServerError, "Topic name  already exists.")
)

// Create creates a new topic on database
func (self *Topic) Create(db orm.DB, topicModel model.Topic) (*model.Topic, error) {
	var topic = new(model.Topic)
	err := db.Model(topic).
		Where("lower(name) = ?", strings.
			ToLower(topicModel.Name)).
		Select()

	if err != nil && err != pg.ErrNoRows {
		return nil, ErrAlreadyExists

	}

	if err := db.Insert(&topicModel); err != nil {
		return nil, err
	}
	return &topicModel, nil
}

// View returns single topic by ID
func (self *Topic) View(db orm.DB, id int) (*model.Topic, error) {
	var topic = new(model.Topic)
	sql := `SELECT "topic".*
	FROM "topics" AS "topic"
	WHERE ("topic"."id" = ?)`
	_, err := db.QueryOne(topic, sql, id)
	if err == pg.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return topic, nil
}

// Update updates topic's contact info
func (self *Topic) Update(db orm.DB, topic *model.Topic) error {
	return db.Update(topic)
}

// List returns list of all topics retrievable for the current topic, depending on role
func (self *Topic) List(db orm.DB, qp *model.FilterQuery, p *model.Pagination) ([]model.Topic, int, error) {
	var topics []model.Topic
	q := db.Model(&topics).
		Column("topic.*").
		Where("deleted_at is null").
		Limit(p.Limit).
		Offset(p.Offset).
		Order("topic.id desc")
	if qp != nil && qp.Query != "" {
		q.Where(qp.Query, qp.Params...)
	}
	totalItems, err := q.Count()
	if err != nil {
		return nil, 0, helper.HandleError("web:topic:list.problem:query.error")
	}
	if err := q.Limit(p.Limit).Offset(p.Offset).Order("topic.created_at desc").Select(); err != nil {
		return nil, 0, helper.HandleError("web:topic:list.problem:query.error")
	}

	return topics, totalItems, nil
}

// Delete sets deleted_at for a topic
func (self *Topic) Delete(db orm.DB, topic *model.Topic) error {
	return db.Delete(topic)
}
