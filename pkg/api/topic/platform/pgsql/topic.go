package pgsql

import (
	"net/http"
	"strings"

	"github.com/go-pg/pg"

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
	ErrAlreadyExists = echo.NewHTTPError(http.StatusInternalServerError, "Topicname or email already exists.")
)

// Create creates a new topic on database
func (u *Topic) Create(db orm.DB, usr gorsk.Topic) (*gorsk.Topic, error) {
	var topic = new(gorsk.Topic)
	err := db.Model(topic).Where("lower(topicname) = ? or lower(email) = ? and deleted_at is null",
		strings.ToLower(usr.Topicname), strings.ToLower(usr.Email)).Select()

	if err != nil && err != pg.ErrNoRows {
		return nil, ErrAlreadyExists

	}

	if err := db.Insert(&usr); err != nil {
		return nil, err
	}
	return &usr, nil
}

// View returns single topic by ID
func (u *Topic) View(db orm.DB, id int) (*gorsk.Topic, error) {
	var topic = new(gorsk.Topic)
	sql := `SELECT "topic".*, "role"."id" AS "role__id", "role"."access_level" AS "role__access_level", "role"."name" AS "role__name"
	FROM "topics" AS "topic" LEFT JOIN "roles" AS "role" ON "role"."id" = "topic"."role_id"
	WHERE ("topic"."id" = ? and deleted_at is null)`
	_, err := db.QueryOne(topic, sql, id)
	if err != nil {
		return nil, err
	}

	return topic, nil
}

// Update updates topic's contact info
func (u *Topic) Update(db orm.DB, topic *gorsk.Topic) error {
	return db.Update(topic)
}

// List returns list of all topics retrievable for the current topic, depending on role
func (u *Topic) List(db orm.DB, qp *gorsk.ListQuery, p *gorsk.Pagination) ([]gorsk.Topic, error) {
	var topics []gorsk.Topic
	q := db.Model(&topics).Column("topic.*", "Role").Limit(p.Limit).Offset(p.Offset).Where("deleted_at is null").Order("topic.id desc")
	if qp != nil {
		q.Where(qp.Query, qp.ID)
	}
	if err := q.Select(); err != nil {
		return nil, err
	}
	return topics, nil
}

// Delete sets deleted_at for a topic
func (u *Topic) Delete(db orm.DB, topic *gorsk.Topic) error {
	return db.Delete(topic)
}
