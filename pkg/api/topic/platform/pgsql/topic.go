package pgsql

import (
	"net/http"
	"strings"

	"github.com/go-pg/pg"

	"../../../../model"
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
	ErrAlreadyExists = echo.NewHTTPError(http.StatusInternalServerError, "Topicname  already exists.")
)

// Create creates a new topic on database
func (u *Topic) Create(db orm.DB, usr model.Topic) (*model.Topic, error) {
	var topic = new(model.Topic)
	err := db.Model(topic).Where("lower(topicname) = ? and deleted_at is null",
		strings.ToLower(usr.Topicname)).Select()

	if err != nil && err != pg.ErrNoRows {
		return nil, ErrAlreadyExists

	}

	if err := db.Insert(&usr); err != nil {
		return nil, err
	}
	return &usr, nil
}

// View returns single topic by ID
func (u *Topic) View(db orm.DB, id int) (*model.Topic, error) {
	var topic = new(model.Topic)
	sql := `SELECT "topic".*
	FROM "topics" AS "topic"
	WHERE ("topic"."id" = ? and deleted_at is null)`
	_, err := db.QueryOne(topic, sql, id)
	if err != nil {
		return nil, err
	}

	return topic, nil
}

// Update updates topic's contact info
func (u *Topic) Update(db orm.DB, topic *model.Topic) error {
	return db.Update(topic)
}

// List returns list of all topics retrievable for the current topic, depending on role
func (u *Topic) List(db orm.DB, qp *model.ListQuery, p *model.Pagination) ([]model.Topic, error) {
	var topics []model.Topic
	q := db.Model(&topics).Column("topic.*").Limit(p.Limit).Offset(p.Offset).Where("deleted_at is null").Order("topic.id desc")
	if qp != nil {
		q.Where(qp.Query, qp.ID)
	}
	if err := q.Select(); err != nil {
		return nil, err
	}
	return topics, nil
}

// Delete sets deleted_at for a topic
func (u *Topic) Delete(db orm.DB, topic *model.Topic) error {
	return db.Delete(topic)
}
